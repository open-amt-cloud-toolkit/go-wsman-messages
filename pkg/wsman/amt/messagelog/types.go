/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package messagelog

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type MessageLog struct {
	base message.Base
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
		XMLName                       xml.Name `xml:"Body"`
		GetResponse                   MessageLogResponse
		EnumerateResponse             common.EnumerateResponse
		PullResponse                  PullResponse
		GetRecordsResponse            GetRecordsResponse
		PositionToFirstRecordResponse PositionToFirstRecordResponse
	}

	PullResponse struct {
		XMLName         xml.Name             `xml:"PullResponse"`
		MessageLogItems []MessageLogResponse `xml:"Items>AMT_MessageLog"`
	}

	MessageLogResponse struct {
		XMLName                xml.Name            `xml:"AMT_MessageLog"`
		Capabilities           []Capabilities      `xml:"Capabilities"`           // An array of integers indicating the Log capabilities. Information such as "Write Record Supported" (value= 2) or "Variable Length Records Supported" (8) is specified in this property.
		CharacterSet           CharacterSet        `xml:"CharacterSet"`           // An enumeration describing the character set used to record data in the individual Log entries. For example, the Log records may contain ASCII data (value=2), or be raw octet strings (value=10).
		CreationClassName      string              `xml:"CreationClassName"`      // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		CurrentNumberOfRecords int                 `xml:"CurrentNumberOfRecords"` // Current number of records in the Log.
		ElementName            string              `xml:"ElementName"`            // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledDefault         EnabledDefault      `xml:"EnabledDefault"`         // An enumerated value indicating an administrator's default or startup configuration for the Enabled State of an element. By default, the element is "Enabled" (value=2).
		EnabledState           EnabledState        `xml:"EnabledState"`           // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		HealthState            HealthState         `xml:"HealthState"`            // Indicates the current health of the element.
		IsFrozen               bool                `xml:"IsFrozen"`               // Boolean indicating that the Log is currently frozen and modifications are not allowed.
		LastChange             LastChange          `xml:"LastChange"`             // An enumeration describing the last change to the MessageLog.
		LogState               LogState            `xml:"LogState"`               // LogState is an integer enumeration that indicates the current state of a log represented by CIM_Log subclasses.
		MaxLogSize             int                 `xml:"MaxLogSize"`             // The maximum size, in bytes, to which the Log can grow. If there is no maximum, then MaxLogSize should be set to 0.
		MaxNumberOfRecords     int                 `xml:"MaxNumberOfRecords"`     // Maximum number of records that can be captured in the Log. If undefined, a value of zero should be specified.
		MaxRecordSize          int                 `xml:"MaxRecordSize"`          // Maximum size, in bytes, to which an individual Log entry (record) can grow - if the Capabilities array includes a value of 7 ("Variable Length Records Supported"). If the Capabilities array does not include a 7, then the Log only supports fixed length entries. The size of these entries is described by this property.
		Name                   string              `xml:"Name"`                   // The inherited Name serves as part of the key (a unique identifier) for the MessageLog instance.
		OperationalStatus      []OperationalStatus `xml:"OperationalStatus"`      // Indicates the current statuses of the element.
		OverwritePolicy        OverwritePolicy     `xml:"OverwritePolicy"`        // An enumeration describing the behavior of the Log, when it becomes full or near full.
		PercentageNearFull     int                 `xml:"PercentageNearFull"`     // If the OverwritePolicy is based on clearing records when the Log is near full (value=3), this property defines the record capacity (in percentage) that is considered to be 'near full'.
		RequestedState         RequestedState      `xml:"RequestedState"`         // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		SizeOfHeader           int                 `xml:"SizeOfHeader"`           // The size of the Log header, in bytes, if one is present. If there is no Log header, then this property should be set to 0. Headers may include general information about the Log such as the current number of records, time of last update, or a pointer to the location of the first Log entry. Note that this property is NOT the size of the header for an individual Log entry. The latter is described by the property, SizeOfRecordHeader.
		SizeOfRecordHeader     int                 `xml:"SizeOfRecordHeader"`     // The size of the header for the Log's individual entries, in bytes, if record headers are defined. If there are no record headers, then this property should be set to 0. Record headers may include information such as the type of the Log entry, the date/time that the entry was last updated, or a pointer to the start of optional data. Note that this property defines the header size for individual records in the Log, while the SizeOfHeader property describes the Log's overall header, typically located at the start of the MessageLog.
		Status                 string              `xml:"Status"`                 // A string indicating the current status of the object. This property is deprecated in lieu of OperationalStatus, which includes the same semantics in its enumeration.
	}

	// An array of integers indicating the Log capabilities. Information such as "Write Record Supported" (value= 2) or "Variable Length Records Supported" (8) is specified in this property.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//
	// Values={Unknown, Other, Write Record Supported, Delete Record Supported, Can Move Backward in Log, Freeze Log Supported, Clear Log Supported, Supports Addressing by Ordinal Record Number, Variable Length Records Supported, Variable Formats for Records, Can Flag Records for Overwrite}
	Capabilities int

	// An enumeration describing the character set used to record data in the individual Log entries. For example, the Log records may contain ASCII data (value=2), or be raw octet strings (value=10).
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	//
	// Values={Unknown, Other, ASCII, Unicode, ISO2022, ISO8859, Extended UNIX Code, UTF-8, UCS-2, Bitmapped Data, OctetString, Defined by Individual Records}
	CharacterSet int

	// An enumeration describing the last change to the MessageLog.
	//
	// ValueMap={0, 1, 2, 3, 4}
	//
	// Values={Unknown, Add, Delete, Modify, Log Cleared}
	LastChange int

	// LogState is an integer enumeration that indicates the current state of a log represented by CIM_Log subclasses. LogState is to be used in conjunction with the EnabledState property to fully describe the current state of the log.
	//
	// The following text briefly summarizes the various log states:
	//
	// - Unknown (0) indicates the state of the log is unknown.
	//
	// - Normal (2) indicates that the log is or could be executing logging commands, will process any queued log entries, and will queue new logging requests.
	//
	// - Erasing (3) indicates that the log is being erased.
	//
	// - Not Applicable (4) indicates the log does not support representing a log state.
	//
	// ValueMap={0, 2, 3, 4, .., 32768..65535}
	//
	// Values={Unknown, Normal, Erasing, Not Applicable, DMTF Reserved, Vendor Reserved}
	LogState int

	// An enumeration describing the behavior of the Log, when it becomes full or near full. For example, the Log may wrap (value=2) or may simply stop recording entries (value =7).
	//
	// Some of the property's possible values need further explanation:
	//
	// - 3="Clear When Near Full" indicates that all of the Log's entries will be deleted when a specified record capacity is reached. The capacity is specified in percentage, using the property, PercentageNearFull. 'Near Full' may be less than 100% if the Log takes time to clear, and a position should always be available for new records.
	//
	// - 4="Overwrite Outdated When Needed" describes that Log entries (timestamped later than the date/time specified in the property, TimeWhenOutdated) can be overwritten.
	//
	// - 5="Remove Outdated Records" specifies that records (timestamped later than the date/time specified in the property, TimeWhenOutdated) are logically and/or physically removed from the Log.
	//
	// - 6="Overwrite Specific Records" indicates that specially flagged records may be overwritten. This property only makes sense when the Capabilities array includes a value of 10, "Can Flag Records for Overwrite".
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7}
	//
	// Values={Unknown, Other, Wraps When Full, Clear When Near Full, Overwrite Outdated When Needed, Remove Outdated Records, Overwrite Specific Records, Never Overwrite}
	OverwritePolicy int

	// An enumerated value indicating an administrator's default or startup configuration for the Enabled State of an element. By default, the element is "Enabled" (value=2).
	//
	// ValueMap={2, 3, 5, 6, 7, 9, .., 32768..65535}
	//
	// Values={Enabled, Disabled, Not Applicable, Enabled but Offline, No Default, Quiesce, DMTF Reserved, Vendor Reserved}
	EnabledDefault int

	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states. For example, shutting down (value=4) and starting (value=10) are transient states between enabled and disabled.
	//
	// The following text briefly summarizes the various enabled and disabled states:
	//
	// - Enabled (2) indicates that the element is or could be executing commands, will process any queued commands, and queues new requests.
	//
	// - Disabled (3) indicates that the element will not execute commands and will drop any new requests.
	//
	// - Shutting Down (4) indicates that the element is in the process of going to a Disabled state.
	//
	// - Not Applicable (5) indicates the element does not support being enabled or disabled.
	//
	// - Enabled but Offline (6) indicates that the element might be completing commands, and will drop any new requests.
	//
	// - Test (7) indicates that the element is in a test state.
	//
	// - Deferred (8) indicates that the element might be completing commands, but will queue any new requests.
	//
	// - Quiesce (9) indicates that the element is enabled but in a restricted mode.
	//
	// - Starting (10) indicates that the element is in the process of going to an Enabled state. New requests are queued.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11..32767, 32768..65535}
	//
	// Values={Unknown, Other, Enabled, Disabled, Shutting Down, Not Applicable, Enabled but Offline, In Test, Deferred, Quiesce, Starting, DMTF Reserved, Vendor Reserved}
	EnabledState int

	// Indicates the current health of the element. This attribute expresses the health of this element but not necessarily that of its subcomponents. The possible values are 0 to 30, where 5 means the element is entirely healthy and 30 means the element is completely non-functional.
	//
	// The following continuum is defined:
	//
	// - "Non-recoverable Error" (30) - The element has completely failed, and recovery is not possible. All functionality provided by this element has been lost.
	//
	// - "Critical Failure" (25) - The element is non-functional and recovery might not be possible.
	//
	// - "Major Failure" (20) - The element is failing. It is possible that some or all of the functionality of this component is degraded or not working.
	//
	// - "Minor Failure" (15) - All functionality is available but some might be degraded.
	//
	// - "Degraded/Warning" (10) - The element is in working order and all functionality is provided. However, the element is not working to the best of its abilities. For example, the element might not be operating at optimal performance or it might be reporting recoverable errors.
	//
	// - "OK" (5) - The element is fully functional and is operating within normal operational parameters and without error.
	//
	// - "Unknown" (0) - The implementation cannot report on HealthState at this time.
	//
	// DMTF has reserved the unused portion of the continuum for additional HealthStates in the future.
	//
	// ValueMap={0, 5, 10, 15, 20, 25, 30, .., 32768..65535}
	//
	// Values={Unknown, OK, Degraded/Warning, Minor failure, Major failure, Critical failure, Non-recoverable error, DMTF Reserved, Vendor Specific}
	HealthState int

	// Indicates the current statuses of the element. Various operational statuses are defined. Many of the enumeration's values are self-explanatory. However, a few are not and are described here in more detail:
	//
	// - "Stressed" indicates that the element is functioning, but needs attention. Examples of "Stressed" states are overload, overheated, and so on.
	//
	// - "Predictive Failure" indicates that an element is functioning nominally but predicting a failure in the near future.
	//
	// - "In Service" describes an element being configured, maintained, cleaned, or otherwise administered.
	//
	// - "No Contact" indicates that the monitoring system has knowledge of this element, but has never been able to establish communications with it.
	//
	// - "Lost Communication" indicates that the ManagedSystem Element is known to exist and has been contacted successfully in the past, but is currently unreachable.
	//
	// - "Stopped" and "Aborted" are similar, although the former implies a clean and orderly stop, while the latter implies an abrupt stop where the state and configuration of the element might need to be updated.
	//
	// - "Dormant" indicates that the element is inactive or quiesced.
	//
	// - "Supporting Entity in Error" indicates that this element might be "OK" but that another element, on which it is dependent, is in error. An example is a network service or endpoint that cannot function due to lower-layer networking problems.
	//
	// - "Completed" indicates that the element has completed its operation. This value should be combined with either OK, Error, or Degraded so that a client can tell if the complete operation Completed with OK (passed), Completed with Error (failed), or Completed with Degraded (the operation finished, but it did not complete OK or did not report an error).
	//
	// - "Power Mode" indicates that the element has additional power model information contained in the Associated PowerManagementService association.
	//
	// - "Relocating" indicates the element is being relocated.
	//
	// OperationalStatus replaces the Status property on ManagedSystemElement to provide a consistent approach to enumerations, to address implementation needs for an array property, and to provide a migration path from today's environment to the future. This change was not made earlier because it required the deprecated qualifier. Due to the widespread use of the existing Status property in management applications, it is strongly recommended that providers or instrumentation provide both the Status and OperationalStatus properties. Further, the first value of OperationalStatus should contain the primary status for the element. When instrumented, Status (because it is single-valued) should also provide the primary status of the element.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, .., 0x8000..}
	//
	// Values={Unknown, Other, OK, Degraded, Stressed, Predictive Failure, Error, Non-Recoverable Error, Starting, Stopping, Stopped, In Service, No Contact, Lost Communication, Aborted, Dormant, Supporting Entity in Error, Completed, Power Mode, Relocating, DMTF Reserved, Vendor Reserved}
	OperationalStatus int

	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested. The actual state of the element is represented by EnabledState. This property is provided to compare the last requested and current enabled or disabled states.
	//
	// Note that when EnabledState is set to 5 ("Not Applicable"), then this property has no meaning. Refer to the EnabledState property description for explanations of the values in the RequestedState enumeration. "Unknown" (0) indicates the last requested state for the element is unknown.
	//
	// Note that the value "No Change" (5) has been deprecated in lieu of indicating the last requested state is "Unknown" (0). If the last requested or desired state is unknown, RequestedState should have the value "Unknown" (0), but may have the value "No Change" (5).  Offline (6) indicates that the element has been requested to transition to the Enabled but Offline EnabledState. It should be noted that there are two new values in RequestedState that build on the statuses of EnabledState. These are "Reboot" (10) and "Reset" (11). Reboot refers to doing a "Shut Down" and then moving to an "Enabled" state. Reset indicates that the element is first "Disabled" and then "Enabled". The distinction between requesting "Shut Down" and "Disabled" should also be noted. Shut Down requests an orderly transition to the Disabled state, and might involve removing power, to completely erase any existing state. The Disabled state requests an immediate disabling of the element, such that it will not execute or accept any commands or processing requests. This property is set as the result of a method invocation (such as Start or StopService on CIM_Service), or can be overridden and defined as WRITEable in a subclass. The method approach is considered superior to a WRITEable property, because it allows an explicit invocation of the operation and the return of a result code. If knowledge of the last RequestedState is not supported for the EnabledLogicalElement, the property shall be NULL or have the value 12 "Not Applicable".
	//
	// ValueMap={0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, .., 32768..65535}
	//
	// Values={Unknown, Enabled, Disabled, Shut Down, No Change, Offline, Test, Deferred, Quiesce, Reboot, Reset, Not Applicable, DMTF Reserved, Vendor Reserved}
	RequestedState int

	GetRecordsResponse struct {
		XMLName             xml.Name `xml:"GetRecords_OUTPUT"`
		IterationIdentifier int      `xml:"IterationIdentifier"` // An identifier for the iterator.
		NoMoreRecords       bool     `xml:"NoMoreRecords"`       // Indicates that there are no more records to read
		RecordArray         []string `xml:"RecordArray"`         // Array of records encoded as Base64
		ReturnValue         int      `xml:"ReturnValue"`         // ValueMap={0, 1, 2, 3} Values={Completed with No Error, Not Supported, Invalid record pointed, No record exists in log}
	}

	PositionToFirstRecordResponse struct {
		XMLName             xml.Name `xml:"PositionToFirstRecord_OUTPUT"`
		IterationIdentifier int      `xml:"IterationIdentifier"` // An identifier for the iterator.
		ReturnValue         int      `xml:"ReturnValue"`         // ValueMap={0, 1, 2} Values={Completed with No Error, Not Supported, No record exists}
	}
)

// INPUTS
type GetRecords_INPUT struct {
	XMLName             xml.Name `xml:"h:GetRecords_INPUT"`
	H                   string   `xml:"xmlns:h,attr"`
	IterationIdentifier int      `xml:"h:IterationIdentifier"` // An identifier for the iterator.
	MaxReadRecords      int      `xml:"h:MaxReadRecords"`      // Maximum number of records to read
}
