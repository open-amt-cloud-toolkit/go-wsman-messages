/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Service struct {
	base message.Base
}

type RequestedState int

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
		XMLName                   xml.Name                  `xml:"Body"`
		RequestStateChange_OUTPUT RequestStateChange_OUTPUT `xml:"RequestStateChange_OUTPUT"`
		GetResponse               UserResponse
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse
	}
	UserResponse struct {
		XMLName                 xml.Name     `xml:"AMT_UserInitiatedConnectionService"`
		CreationClassName       string       `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName             string       `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledState            EnabledState `xml:"EnabledState"`                      // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		Name                    string       `xml:"Name,omitempty"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		SystemCreationClassName string       `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		SystemName              string       `xml:"SystemName,omitempty"`              // The Name of the scoping System.
	}
	PullResponse struct {
		XMLName   xml.Name       `xml:"PullResponse"`
		UserItems []UserResponse `xml:"Items>AMT_UserInitiatedConnectionService"`
	}

	// The state requested for the element. This information will be placed into the RequestedState property of the instance if the return code of the RequestStateChange method is 0 ('Completed with No Error'), 3 ('Timeout'), or 4096 (0x1000) ('Job Started'). Refer to the description of the EnabledState and RequestedState properties for the detailed explanations of the RequestedState values.
	//
	// ValueMap={2, 3, 4, 6, 7, 8, 9, 10, 11, .., 32768, 32769, 32770, 32771, 32772..65535}
	//
	// Values={Enabled, Disabled, Shut Down, Offline, Test, Defer, Quiesce, Reboot, Reset, DMTF Reserved, All Interfaces disabled, BIOS Interface enabled, OS Interface enabled, BIOS and OS Interfaces enabled, Vendor Reserved}
	RequestStateChange_OUTPUT struct {
		XMLName     xml.Name `xml:"RequestStateChange_OUTPUT"`
		ReturnValue int      `xml:"ReturnValue"`
	}
	//EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states. For example, shutting down (value=4) and starting (value=10) are transient states between enabled and disabled. The following text briefly summarizes the various enabled and disabled states:
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
	// Quiesce (9) indicates that the element is enabled but in a restricted mode. The behavior of the element is similar to the Enabled state, but it processes only a restricted set of commands. All other requests are queued.
	//
	// Starting (10) indicates that the element is in the process of going to an Enabled state. New requests are queued.
	//
	// The supported values are 32768-32771.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11..32767, 32768, 32769, 32770, 32771, 32772..65535}
	//
	// Values={Unknown, Other, Enabled, Disabled, Shutting Down, Not Applicable, Enabled but Offline, In Test, Deferred, Quiesce, Starting, DMTF Reserved, All Interfaces disabled, BIOS Interface enabled, OS Interface enabled, BIOS and OS Interfaces enabled, Vendor Reserved}
	EnabledState int
)
