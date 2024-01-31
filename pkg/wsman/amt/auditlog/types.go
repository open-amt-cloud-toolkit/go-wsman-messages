/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package auditlog

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Service struct {
	base message.Base
}

// INPUTS
// Request Types
type readRecords_INPUT struct {
	XMLName    xml.Name `xml:"h:ReadRecords_INPUT"`
	H          string   `xml:"xmlns:h,attr"`
	StartIndex int      `xml:"h:StartIndex" json:"StartIndex"`
}

// OUTPUTS
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName             xml.Name `xml:"Body"`
		EnumerateResponse   common.EnumerateResponse
		GetResponse         AuditLog
		PullResponse        PullResponse
		ReadRecordsResponse ReadRecords_OUTPUT
	}
	PullResponse struct {
		XMLName       xml.Name   `xml:"PullResponse"`
		AuditLogItems []AuditLog `xml:"Items>AMT_AuditLog"`
	}

	AuditLog struct {
		XMLName                xml.Name        `xml:"AMT_AuditLog"`
		OverwritePolicy        OverwritePolicy `xml:"OverwritePolicy,omitempty"`        // OverwritePolicy is an integer enumeration that indicates whether the log, represented by the CIM_Log subclasses, can overwrite its entries.Unknown (0) indicates the log's overwrite policy is unknown
		CurrentNumberOfRecords int             `xml:"CurrentNumberOfRecords,omitempty"` // Current number of records in the Log
		MaxNumberOfRecords     int             `xml:"MaxNumberOfRecords,omitempty"`     // Maximum number of records that can be captured in the Log
		ElementName            string          `xml:"ElementName,omitempty"`            // A user-friendly name for the object
		EnabledState           EnabledState    `xml:"EnabledState,omitempty"`           // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element
		RequestedState         RequestedState  `xml:"RequestedState,omitempty"`         // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested
		PercentageFree         int             `xml:"PercentageFree,omitempty"`         // Indicates the percentage of free space in the storage dedicated to the audit log
		Name                   string          `xml:"Name,omitempty"`                   // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed
		TimeOfLastRecord       Datetime        `xml:"TimeOfLastRecord"`                 // Time stamp of the most recent entry in the log if such an entry exists
		AuditState             int             `xml:"AuditState,omitempty"`             // State of log
		MaxAllowedAuditors     int             `xml:"MaxAllowedAuditors,omitempty"`     // Maximum number of auditors allowed
		StoragePolicy          StoragePolicy   `xml:"StoragePolicy,omitempty"`          // AuditLog storage policy
		MinDaysToKeep          int             `xml:"MinDaysToKeep,omitempty"`          // Minimum number of days to keep records in the AuditLog
	}

	Datetime struct {
		Datetime string `xml:"Datetime,omitempty"`
	}

	ReadRecords_OUTPUT struct {
		XMLName          xml.Name `xml:"ReadRecords_OUTPUT,omitempty"`
		TotalRecordCount int      `xml:"TotalRecordCount,omitempty"` // The total number of records in the log.
		RecordsReturned  int      `xml:"RecordsReturned,omitempty"`  // The number of records returned + content of 10 records from the start index.
		EventRecords     []string `xml:"EventRecords,omitempty"`     // Notice: the values of this array are actually base64 encoded values. A list of event records.
		ReturnValue      int      `xml:"ReturnValue,omitempty"`      // ValueMap={0, 1, 2, 35} Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_NOT_READY, PT_STATUS_INVALID_INDEX}
	}
)

type (
	// OverwritePolicy is an integer enumeration that indicates whether the log, represented by the CIM_Log subclasses, can overwrite its entries.
	//
	// Unknown (0) indicates the log's overwrite policy is unknown.
	//
	// Wraps when Full (2) indicates that the log overwrites its entries with new entries when the log has reached its maximum capacity.
	//
	// Never Overwrites (7) indicates that the log never overwrites its entries by the new entries.
	//
	// Value of 32768 means "Parital (restricted) rollover" - only old events (under certain threshold, which can be set using SetStoragePolicy) will be overwritten.
	//
	// This is a read-only property.
	//
	// ValueMap={0, 2, 7, .., 32768..65535}
	//
	// Values={Unknown, Wraps When Full, Never Overwrites, DMTF Reserved, Vendor Reserved}
	OverwritePolicy int
	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states. For example, shutting down (value=4) and starting (value=10) are transient states between enabled and disabled.
	//
	// The following text briefly summarizes the various enabled and disabled states:
	//
	// Enabled (2) indicates that the element is or could be executing commands, will process any queued commands, and queues new requests.
	//
	// Disabled (3) indicates that the element will not execute commands and will drop any new requests.
	//
	// Shutting Down (4) indicates that the element is in the process of going to a Disabled state.
	//
	// Not Applicable (5) indicates the element does not support being enabled or disabled.
	//
	// Enabled but Offline (6) indicates that the element might be completing commands, and will drop any new requests.
	//
	// Test (7) indicates that the element is in a test state.
	//
	// Deferred (8) indicates that the element might be completing commands, but will queue any new requests.
	//
	// Quiesce (9) indicates that the element is enabled but in a restricted mode.
	//
	// Starting (10) indicates that the element is in the process of going to an Enabled state. New requests are queued.
	//
	// Value 6 ("Enabled but Offline") can be recieved also if the Audit Log is in locked state.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11..32767, 32768..65535}
	//
	// Values={Unknown, Other, Enabled, Disabled, Shutting Down, Not Applicable, Enabled but Offline, In Test, Deferred, Quiesce, Starting, DMTF Reserved, Vendor Reserved}
	EnabledState int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested. The actual state of the element is represented by EnabledState. This property is provided to compare the last requested and current enabled or disabled states. Note that when EnabledState is set to 5 ("Not Applicable"), then this property has no meaning. Refer to the EnabledState property description for explanations of the values in the RequestedState enumeration.
	//
	// "Unknown" (0) indicates the last requested state for the element is unknown.
	//
	// Note that the value "No Change" (5) has been deprecated in lieu of indicating the last requested state is "Unknown" (0). If the last requested or desired state is unknown, RequestedState should have the value "Unknown" (0), but may have the value "No Change" (5).Offline (6) indicates that the element has been requested to transition to the Enabled but Offline EnabledState.	It should be noted that there are two new values in RequestedState that build on the statuses of EnabledState. These are "Reboot" (10) and "Reset" (11). Reboot refers to doing a "Shut Down" and then moving to an "Enabled" state. Reset indicates that the element is first "Disabled" and then "Enabled". The distinction between requesting "Shut Down" and "Disabled" should also be noted. Shut Down requests an orderly transition to the Disabled state, and might involve removing power, to completely erase any existing state. The Disabled state requests an immediate disabling of the element, such that it will not execute or accept any commands or processing requests.
	//
	// This property is set as the result of a method invocation (such as Start or StopService on CIM_Service), or can be overridden and defined as WRITEable in a subclass. The method approach is considered superior to a WRITEable property, because it allows an explicit invocation of the operation and the return of a result code.
	//
	// If knowledge of the last RequestedState is not supported for the EnabledLogicalElement, the property shall be NULL or have the value 12 "Not Applicable".
	//
	// ValueMap={0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, .., 32768..65535}
	//
	// Values={Unknown, Enabled, Disabled, Shut Down, No Change, Offline, Test, Deferred, Quiesce, Reboot, Reset, Not Applicable, DMTF Reserved, Vendor Reserved}
	RequestedState int
	// AuditLog storage policy. The available policies are: "NO_ROLL_OVER" - No roll-over in storage. Old events will not be overwritten. "ROLL_OVER" - Full rollover. Any old record will be overwritten. "RESTRICTED_ROLL_OVER" - Parital (restricted) rollover. Only old events (under certain threshold, which can be set using SetStoragePolicy) will be overwritten. If not specified default is "ROLL_OVER" unless FW was upgraded from AMT 5.0 and then "NO_ROLL_OVER" will be used as default.
	//
	// Additional Notes: 'StoragePolicy' is only supported in Intel AMT Release 5.1 and later releases.
	//
	// ValueMap={0, 1, 2}
	//
	// Values={NO_ROLL_OVER, ROLL_OVER, RESTRICTED_ROLL_OVER}
	StoragePolicy int
)
