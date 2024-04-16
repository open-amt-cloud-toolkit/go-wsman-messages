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
		EnabledState           int             `xml:"EnabledState,omitempty"`           // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element
		RequestedState         int             `xml:"RequestedState,omitempty"`         // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested
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

	// OverwritePolicy is an integer enumeration that indicates whether the log, represented by the CIM_Log subclasses, can overwrite its entries.
	OverwritePolicy int

	// StoragePolicy is an integer enumeration that indicates the storage policy of the log.
	StoragePolicy int
)
