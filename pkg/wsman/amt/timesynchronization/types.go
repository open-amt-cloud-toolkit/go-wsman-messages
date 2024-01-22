/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Service struct {
	base message.Base
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
		XMLName                          xml.Name `xml:"Body"`
		GetResponse                      TimeSynchronizationServiceResponse
		EnumerateResponse                common.EnumerateResponse
		PullResponse                     PullResponse
		GetLowAccuracyTimeSynchResponse  GetLowAccuracyTimeSynchResponse
		SetHighAccuracyTimeSynchResponse SetHighAccuracyTimeSynchResponse
	}

	PullResponse struct {
		XMLName                         xml.Name                             `xml:"PullResponse"`
		TimeSynchronizationServiceItems []TimeSynchronizationServiceResponse `xml:"Items>AMT_TimeSynchronizationService"`
	}

	TimeSynchronizationServiceResponse struct {
		XMLName                 xml.Name             `xml:"AMT_TimeSynchronizationService"`
		Name                    string               `xml:"Name,omitempty"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		CreationClassName       string               `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		SystemName              string               `xml:"SystemName,omitempty"`              // The Name of the scoping System.
		SystemCreationClassName string               `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		ElementName             string               `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledState            EnabledState         `xml:"EnabledState,omitempty"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          RequestedState       `xml:"RequestedState,omitempty"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		LocalTimeSyncEnabled    LocalTimeSyncEnabled `xml:"LocalTimeSyncEnabled,omitempty"`    // Determines if user with LOCAL_SYSTEM_REALM permission can set the time.
		TimeSource              TimeSource           `xml:"TimeSource,omitempty"`              // Determines if RTC was set to UTC by any configuration SW.
	}

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
	// Determines if user with LOCAL_SYSTEM_REALM permission can set the time. The values are:
	//
	// - DEFAULT_TRUE - Time sync is enabled by default. Was not configured to enabled or disabled by the configuration SW.
	//
	// - CONFIGURED_TRUE - Time Sync is enabled and configured by configuration SW to TRUE. This option is required in order to differentiate between legacy configuration SW that do not support this setting and new SW that can configure it to TRUE.
	//
	// - FALSE - Time Sync is disabled.
	//
	// ValueMap={0, 1, 2, 3..}
	//
	// Values={DEFAULT_TRUE, CONFIGURED_TRUE, FALSE, RESERVED}
	LocalTimeSyncEnabled int
	// Determines if RTC was set to UTC by any configuration SW.
	//
	// ValueMap={0, 1, 2..}
	//
	// Values={BIOS_RTC, CONFIGURED, RESERVED}
	TimeSource int

	// ValueMap={0, 1}
	//
	// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR}
	GetLowAccuracyTimeSynchResponse struct {
		XMLName     xml.Name `xml:"GetLowAccuracyTimeSynch_OUTPUT"`
		Ta0         int64    `xml:"Ta0"`
		ReturnValue int      `xml:"ReturnValue"`
	}

	// ValueMap={0, 1, 36, 38}
	//
	// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_INVALID_PARAMETER, PT_STATUS_FLASH_WRITE_LIMIT_EXCEEDED}
	SetHighAccuracyTimeSynchResponse struct {
		XMLName     xml.Name `xml:"SetHighAccuracyTimeSynch_OUTPUT"`
		ReturnValue int      `xml:"ReturnValue"`
	}
)

// Request Types
type (
	// Ta0: The time value received from invoking GetLowAccuracyTimeSynch().
	//
	// Tm1: The remote client timestamp after getting a response from GetLowAccuracyTimeSynch().
	//
	// Tm2: The remote client timestamp obtained immediately prior to invoking this method.
	SetHighAccuracyTimeSynch_INPUT struct {
		XMLName xml.Name `xml:"h:SetHighAccuracyTimeSynch_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Ta0     int64    `xml:"h:Ta0"`
		Tm1     int64    `xml:"h:Tm1"`
		Tm2     int64    `xml:"h:Tm2"`
	}
)
