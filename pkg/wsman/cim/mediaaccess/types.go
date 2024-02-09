/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mediaaccess

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Device struct {
	base   message.Base
	client client.WSMan
}

// Response Types
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
	}

	PullResponse struct {
		XMLName            xml.Name            `xml:"PullResponse"`
		MediaAccessDevices []MediaAccessDevice `xml:"Items>CIM_MediaAccessDevice"`
	}

	// TODO: Capabilities and OperationalStatus can return multiple items with the same tag, need to handle this
	MediaAccessDevice struct {
		Capabilities            CapabilitiesValues `xml:"Capabilities,omitempty"`  // ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} Values={Unknown, Other, Sequential Access, Random Access, Supports Writing, Encryption, Compression, Supports Removeable Media, Manual Cleaning, Automatic Cleaning, SMART Notification, Supports Dual Sided Media, Predismount Eject Not Required} ArrayType=Indexed
		CreationClassName       string             `xml:"CreationClassName"`       // CreationClassName indicates the name of the class or the subclass used in the creation of an instance.
		DeviceID                string             `xml:"DeviceID"`                // An address or other identifying information to uniquely name the LogicalDevice.
		ElementName             string             `xml:"ElementName"`             // This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information.
		EnabledDefault          EnabledDefault     `xml:"EnabledDefault"`          // An enumerated value indicating an administrator's default or startup configuration for the Enabled State of an element.
		EnabledState            EnabledState       `xml:"EnabledState"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states.
		MaxMediaSize            int                `xml:"MaxMediaSize,omitempty"`  // Maximum size, in KBytes, of media supported by this Device.
		OperationalStatus       OperationalStatus  `xml:"OperationalStatus"`       // Indicates the current statuses of the element.
		RequestedState          RequestedState     `xml:"RequestedState"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		Security                SecurityValues     `xml:"Security,omitempty"`      // ValueMap={1, 2, 3, 4, 5, 6, 7} Values={Other, Unknown, None, Read Only, Locked Out, Boot Bypass, Boot Bypass and Read Only}
		SystemCreationClassName string             `xml:"SystemCreationClassName"` // The scoping System's CreationClassName.
		SystemName              string             `xml:"SystemName"`              // The scoping System's Name.
	}

	// Capabilities of the MediaAccessDevice. For example, the Device may support "Random Access", removeable media and "Automatic Cleaning". In this case, the values 3, 7 and 9 would be written to the array.
	//
	// - Several of the enumerated values require some explanation: 1) Value 11, Supports Dual Sided Media, distinguishes a Device that can access both sides of dual sided Media, from a Device that reads only a single side and requires the Media to be flipped; and,
	//
	// - 2) Value 12, Predismount Eject Not Required, indicates that Media does not have to be explicitly ejected from the Device before being accessed by a PickerElement.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//
	// Values={Unknown, Other, Sequential Access, Random Access, Supports Writing, Encryption, Compression, Supports Removeable Media, Manual Cleaning, Automatic Cleaning, SMART Notification, Supports Dual Sided Media, Predismount Eject Not Required}
	CapabilitiesValues int

	// An enumerated value indicating an administrator's default or startup configuration for the Enabled State of an element. By default, the element is "Enabled" (value=2).
	//
	// ValueMap={2, 3, 5, 6, 7, 9, .., 32768..65535}
	//
	// Values={Enabled, Disabled, Not Applicable, Enabled but Offline, No Default, Quiesce, DMTF Reserved, Vendor Reserved}
	EnabledDefault int

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
	EnabledState int

	// Indicates the current statuses of the element. Various operational statuses are defined. Many of the enumeration's values are self-explanatory. However, a few are not and are described here in more detail.
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
	RequestedState int

	// An enumeration indicating the operational security defined for the MediaAccessDevice. For example, information that the Device is "Read Only" (value=4) or "Boot Bypass" (value=6) can be described using this property.
	//
	// ValueMap={1, 2, 3, 4, 5, 6, 7}
	//
	// Values={Other, Unknown, None, Read Only, Locked Out, Boot Bypass, Boot Bypass and Read Only}
	SecurityValues int
)
