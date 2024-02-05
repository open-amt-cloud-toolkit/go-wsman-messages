/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kvm

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type RedirectionSAP struct {
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
		XMLName           xml.Name          `xml:"Body"`
		GetResponse       KVMRedirectionSAP `xml:"CIM_KVMRedirectionSAP"`
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse `xml:"PullResponse"`
	}

	KVMRedirectionSAP struct {
		CreationClassName       string                                `xml:"CreationClassName"`        // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName             string                                `xml:"ElementName"`              // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		Name                    string                                `xml:"Name"`                     // The Name property uniquely identifies the ServiceAccessPoint and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		SystemCreationClassName string                                `xml:"SystemCreationClassName"`  // The CreationClassName of the scoping System.
		SystemName              string                                `xml:"SystemName"`               // The Name of the scoping System.
		EnabledState            EnabledState                          `xml:"EnabledState,omitempty"`   // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          KVMRedirectionSAPRequestedStateInputs `xml:"RequestedState,omitempty"` // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		KVMProtocol             KVMRedirectionSAPKVMProtocol          `xml:"KVMProtocol,omitempty"`    // An enumeration specifying the type of the KVM stream supported on this SAP.
	}

	Time struct {
		DateTime string `xml:"Datetime"`
	}

	PullResponse struct {
		XMLName xml.Name            `xml:"PullResponse"`
		Items   []KVMRedirectionSAP `xml:"Items>CIM_KVMRedirectionSAP"`
	}
)

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
// -"Unknown" (0) indicates the last requested state for the element is unknown.
//
// Note that the value "No Change" (5) has been deprecated in lieu of indicating the last requested state is "Unknown" (0).
//
// If the last requested or desired state is unknown, RequestedState should have the value "Unknown" (0), but may have the value "No Change" (5).
//
// - Offline (6) indicates that the element has been requested to transition to the Enabled but Offline EnabledState.
//
// It should be noted that there are two new values in RequestedState that build on the statuses of EnabledState. These are "Reboot" (10) and "Reset" (11).
//
// - Reboot refers to doing a "Shut Down" and then moving to an "Enabled" state.
//
// - Reset indicates that the element is first "Disabled" and then "Enabled".
//
// The distinction between requesting "Shut Down" and "Disabled" should also be noted. Shut Down requests an orderly transition to the Disabled state, and might involve removing power, to completely erase any existing state. The Disabled state requests an immediate disabling of the element, such that it will not execute or accept any commands or processing requests.
//
// This property is set as the result of a method invocation (such as Start or StopService on CIM_Service), or can be overridden and defined as WRITEable in a subclass. The method approach is considered superior to a WRITEable property, because it allows an explicit invocation of the operation and the return of a result code.
//
// If knowledge of the last RequestedState is not supported for the EnabledLogicalElement, the property shall be NULL or have the value 12 "Not Applicable".
type KVMRedirectionSAPRequestedStateInputs int

// An enumeration specifying the type of the KVM stream supported on this SAP. In some cases this may be a raw video steam, with the characters having no special meaning. However, in other cases it may support a protocol where some messages have a predefined structure.
//
// - 0 "Unknown" shall indicate the protocol is unknown.
//
// - 1 "Other" shall indicate the protocol is specified in OtherKVMProtocol.
//
// - 2 "Raw" shall indicate the protocol is a raw and uncompressed data stream. 3 "RDP" shall indicate the protocol is the Remote Desktop Protocol.
//
// - 4 "VNC" shall indicate the protocol is the VNC Protocol.
//
// ValueMap={0, 1, 2, 3, 4, 5..32767, 32768..65535}
//
// Values={Unknown, Other, Raw, RDP, VNC-RFB, DMTF Reserved, Vendor Specified}
type KVMRedirectionSAPKVMProtocol int
