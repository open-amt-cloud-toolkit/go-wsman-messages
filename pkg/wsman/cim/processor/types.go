/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package processor

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Package struct {
	base   message.Base
	client client.WSMan
}

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name `xml:"Body"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
		PackageResponse   PackageResponse
	}

	PullResponse struct {
		XMLName      xml.Name          `xml:"PullResponse"`
		PackageItems []PackageResponse `xml:"Items>CIM_Processor"`
	}

	PackageResponse struct {
		XMLName                 xml.Name          `xml:"CIM_Processor"`
		DeviceID                string            `xml:"DeviceID,omitempty"`                // An address or other identifying information to uniquely name the LogicalDevice.
		CreationClassName       string            `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		SystemName              string            `xml:"SystemName,omitempty"`              // The scoping System's Name.
		SystemCreationClassName string            `xml:"SystemCreationClassName,omitempty"` // The scoping System's CreationClassName.
		ElementName             string            `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		OperationalStatus       OperationalStatus `xml:"OperationalStatus,omitempty"`       // Indicates the current statuses of the element.
		HealthState             HealthState       `xml:"HealthState,omitempty"`             // Indicates the current health of the element.
		EnabledState            EnabledState      `xml:"EnabledState,omitempty"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          RequestedState    `xml:"RequestedState,omitempty"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		Role                    string            `xml:"Role,omitempty"`                    // A free-form string that describes the role of the Processor, for example, "Central Processor" or "Math Processor".
		Family                  int               `xml:"Family,omitempty"`                  // The Processor family type. For example, values include "Pentium(R) processor with MMX(TM) technology" (value=14) and "68040" (value=96).
		OtherFamilyDescription  string            `xml:"OtherFamilyDescription,omitempty"`  // A string that describes the Processor Family type. It is used when the Family property is set to 1 ("Other"). This string should be set to NULL when the Family property is any value other than 1.
		UpgradeMethod           UpgradeMethod     `xml:"UpgradeMethod,omitempty"`           // CPU socket information that includes data on how this Processor can be upgraded (if upgrades are supported). This property is an integer enumeration.
		MaxClockSpeed           int               `xml:"MaxClockSpeed,omitempty"`           // The maximum speed (in MHz) of this Processor.
		CurrentClockSpeed       int               `xml:"CurrentClockSpeed,omitempty"`       // The current speed (in MHz) of this Processor.
		Stepping                string            `xml:"Stepping,omitempty"`                // Stepping is a free-form string that indicates the revision level of the Processor within the Processor.Family.
		CPUStatus               CPUStatus         `xml:"CPUStatus,omitempty"`               // The CPUStatus property that indicates the current status of the Processor.
		ExternalBusClockSpeed   int               `xml:"ExternalBusClockSpeed,omitempty"`   // The speed (in MHz) of the external bus interface (also known as the front side bus).
	}
)

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
type OperationalStatus int

// Indicates the current health of the element. This attribute expresses the health of this element but not necessarily that of its subcomponents. The possible values are 0 to 30, where 5 means the element is entirely healthy and 30 means the element is completely non-functional. The following continuum is defined:
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
type HealthState int

// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states. For example, shutting down (value=4) and starting (value=10) are transient states between enabled and disabled. The following text briefly summarizes the various enabled and disabled states:
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
type EnabledState int

// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested. The actual state of the element is represented by EnabledState. This property is provided to compare the last requested and current enabled or disabled states. Note that when EnabledState is set to 5 ("Not Applicable"), then this property has no meaning. Refer to the EnabledState property description for explanations of the values in the RequestedState enumeration.
//
// - "Unknown" (0) indicates the last requested state for the element is unknown.
//
// - Note that the value "No Change" (5) has been deprecated in lieu of indicating the last requested state is "Unknown" (0). If the last requested or desired state is unknown, RequestedState should have the value "Unknown" (0), but may have the value "No Change" (5).Offline (6) indicates that the element has been requested to transition to the Enabled but Offline EnabledState.
//
// - It should be noted that there are two new values in RequestedState that build on the statuses of EnabledState. These are "Reboot" (10) and "Reset" (11). Reboot refers to doing a "Shut Down" and then moving to an "Enabled" state. Reset indicates that the element is first "Disabled" and then "Enabled". The distinction between requesting "Shut Down" and "Disabled" should also be noted. Shut Down requests an orderly transition to the Disabled state, and might involve removing power, to completely erase any existing state. The Disabled state requests an immediate disabling of the element, such that it will not execute or accept any commands or processing requests.
//
// This property is set as the result of a method invocation (such as Start or StopService on CIM_Service), or can be overridden and defined as WRITEable in a subclass. The method approach is considered superior to a WRITEable property, because it allows an explicit invocation of the operation and the return of a result code.
//
// If knowledge of the last RequestedState is not supported for the EnabledLogicalElement, the property shall be NULL or have the value 12 "Not Applicable".
//
// ValueMap={0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, .., 32768..65535}
//
// Values={Unknown, Enabled, Disabled, Shut Down, No Change, Offline, Test, Deferred, Quiesce, Reboot, Reset, Not Applicable, DMTF Reserved, Vendor Reserved}
type RequestedState int

// CPU socket information that includes data on how this Processor can be upgraded (if upgrades are supported). This property is an integer enumeration.
//
// ValueMap={ "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72" , "73", "74", "75", "76", "77", "78", "79", "80" }
//
// Values={ "Other", "Unknown", "Daughter Board", "ZIF Socket", "Replacement/Piggy Back", "None", "LIF Socket", "Slot 1", "Slot 2", "370 Pin Socket", "Slot A", "Slot M", "Socket 423", "Socket A (Socket 462)", "Socket 478", "Socket 754", "Socket 940", "Socket 939", "Socket mPGA604", "Socket LGA771", "Socket LGA775", "Socket S1", "Socket AM2", "Socket F (1207)", "Socket LGA1366", "Socket G34", "Socket AM3", "Socket C32", "Socket LGA1156", "Socket LGA1567", "Socket PGA988A", "Socket BGA1288", "rPGA988B", "BGA1023", "BGA1224", "LGA1155", "LGA1356", "LGA2011", "Socket FS1", "Socket FS2", "Socket FM1", "Socket FM2", "Socket LGA2011-3", "Socket LGA1356-3", "Socket LGA1150", "Socket BGA1168", "Socket BGA1234", "Socket BGA1364", "Socket AM4", "Socket LGA1151", "Socket BGA1356", "Socket BGA1440", "Socket BGA1515", "Socket LGA3647-1", "Socket SP3", "Socket SP3r2", "Socket LGA2066", "Socket BGA1392", "Socket BGA1510", "Socket BGA1528", "Socket LGA4189", "Socket LGA1200", "Socket LGA4677", "Socket LGA1700", "Socket BGA1744", "Socket BGA1781", "Socket BGA1211", "Socket BGA2422", "Socket LGA5773", "Socket BGA5773", "Socket AM5", "Socket SP5", "Socket SP6", "Socket BGA883", "Socket BGA1190", "Socket BGA4129", "Socket LGA4710", "Socket LGA7529" }
type UpgradeMethod int

// The CPUStatus property that indicates the current status of the Processor. For example, the Processor might be disabled by the user (value=2), or disabled due to a POST error (value=3). Information in this property can be obtained from SMBIOS, the Type 4 structure, and the Status attribute.
//
// ValueMap={0, 1, 2, 3, 4, 7}
//
// Values={Unknown, CPU Enabled, CPU Disabled by User, CPU Disabled By BIOS (POST Error), CPU Is Idle, Other}
type CPUStatus int
