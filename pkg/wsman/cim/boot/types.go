/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type ConfigSetting struct {
	base   message.Base
	client client.WSMan
}

type SourceSetting struct {
	base   message.Base
	client client.WSMan
}

type Service struct {
	base   message.Base
	client client.WSMan
}

type Source string

// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                  xml.Name          `xml:"Body"`
		ConfigSettingGetResponse BootConfigSetting `xml:"CIM_BootConfigSetting"`
		SourceSettingGetResponse BootSourceSetting `xml:"CIM_BootSourceSetting"`
		ServiceGetResponse       BootService       `xml:"CIM_BootService"`
		EnumerateResponse        common.EnumerateResponse
		PullResponse             PullResponse           `xml:"PullResponse"`
		ChangeBootOrder_OUTPUT   ChangeBootOrder_OUTPUT `xml:"ChangeBootOrder_OUTPUT"`
	}

	BootConfigSetting struct {
		XMLName     xml.Name `xml:"CIM_BootConfigSetting"`
		InstanceID  string   `xml:"InstanceID"`  // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>. <LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance.	For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
		ElementName string   `xml:"ElementName"` // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
	}

	BootSourceSetting struct {
		XMLName              xml.Name             `xml:"CIM_BootSourceSetting"`
		ElementName          string               `xml:"ElementName"`          // he user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		InstanceID           string               `xml:"InstanceID"`           // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>. <LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance. For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
		StructuredBootString string               `xml:"StructuredBootString"` // A string identifying the boot source using the format "<OrgID>:<identifier>:<index>", in which neither <OrgID>, <identifier> or <index> contains a colon (":"). The value of <OrgID> is a copyrighted, trademarked or otherwise unique name that is owned by the entity defining the <identifier>, or is a registered ID that is assigned to the entity by a recognized global authority. For DMTF defined identifiers, the <OrgID> is set to 'CIM'. The <identifiers> are "Floppy", "Hard-Disk", "CD/DVD", "Network", "PCMCIA", "USB". The value of <index> shall be a non-zero integer.
		BIOSBootString       string               `xml:"BIOSBootString"`       // BIOS description of the boot option. UTF8 null-terminated string, relevant only to One-Click Recovery WinRE and PBA (pre-boot application) boot options.
		BootString           string               `xml:"BootString"`           // BIOS description of the EFI device path. UTF8 null-terminated string, relevant only to One-Click Recovery WinRE and PBA (pre-boot application) boot options.
		FailThroughSupported FailThroughSupported `xml:"FailThroughSupported"` // An enumeration indicating the behavior when the attempt to boot using the boot source fails (no media, timeout).
	}

	BootService struct {
		XMLName                 xml.Name            `xml:"CIM_BootService"`
		Name                    string              `xml:"Name"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		CreationClassName       string              `xml:"CreationClassName"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		SystemName              string              `xml:"SystemName"`              // The Name of the scoping System.
		SystemCreationClassName string              `xml:"SystemCreationClassName"` // The CreationClassName of the scoping System.
		ElementName             string              `xml:"ElementName"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		OperationalStatus       []OperationalStatus `xml:"OperationalStatus"`       // Indicates the current statuses of the element. Various operational statuses are defined.
		EnabledState            EnabledState        `xml:"EnabledState"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          RequestedState      `xml:"RequestedState"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	}

	PullResponse struct {
		BootSourceSettingItems []BootSourceSetting `xml:"Items>CIM_BootSourceSetting"`
		BootConfigSettingItems []BootConfigSetting `xml:"Items>CIM_BootConfigSetting"`
		BootServiceItems       []BootService       `xml:"Items>CIM_BootService"`
	}

	ChangeBootOrder_OUTPUT struct {
		ReturnValue int `xml:"ReturnValue"`
	}
)

type (
	// An enumeration indicating the behavior when the attempt to boot using the boot source fails (no media, timeout). The current values in the enumeration are:
	//
	// 0 = Unknown
	//
	// 1 = Is Supported
	//
	// 2 = Is Not Supported.
	//
	// A value of 1 (Is Supported) indicates that next boot source the boot order is used.
	// A value of 2 (Is Not Supported) indicates that the boot order is terminated and no other boot sources associated to the same CIM_BootConfigSetting are used).
	// The default is 1 (Is Supported).
	// In Intel (r) AMT the default value is 2 (Is Not Supported)
	//
	// ValueMap={0, 1, 2}
	//
	// Values={Unknown, Is Supported, Not Supported}
	FailThroughSupported int
	// Indicates the current statuses of the element. Various operational statuses are defined.
	// Many of the enumeration's values are self-explanatory.
	// However, a few are not and are described here in more detail.
	//
	// "Stressed" indicates that the element is functioning, but needs attention. Examples of "Stressed" states are overload, overheated, and so on.
	//
	// "Predictive Failure" indicates that an element is functioning nominally but predicting a failure in the near future.
	//
	// "In Service" describes an element being configured, maintained, cleaned, or otherwise administered.
	//
	// "No Contact" indicates that the monitoring system has knowledge of this element, but has never been able to establish communications with it.
	//
	// "Lost Communication" indicates that the ManagedSystem Element is known to exist and has been contacted successfully in the past, but is currently unreachable.
	//
	// "Stopped" and "Aborted" are similar, although the former implies a clean and orderly stop, while the latter implies an abrupt stop where the state and configuration of the element might need to be updated.
	//
	// "Dormant" indicates that the element is inactive or quiesced.
	//
	// "Supporting Entity in Error" indicates that this element might be "OK" but that another element, on which it is dependent, is in error. An example is a network service or endpoint that cannot function due to lower-layer networking problems.
	//
	// "Completed" indicates that the element has completed its operation. This value should be combined with either OK, Error, or Degraded so that a client can tell if the complete operation Completed with OK (passed), Completed with Error (failed), or Completed with Degraded (the operation finished, but it did not complete OK or did not report an error).
	//
	// "Power Mode" indicates that the element has additional power model information contained in the Associated PowerManagementService association.
	//
	// "Relocating" indicates the element is being relocated.
	//
	// OperationalStatus replaces the Status property on ManagedSystemElement to provide a consistent approach to enumerations, to address implementation needs for an array property, and to provide a migration path from today's environment to the future.
	// This change was not made earlier because it required the deprecated qualifier.
	// Due to the widespread use of the existing Status property in management applications, it is strongly recommended that providers or instrumentation provide both the Status and OperationalStatus properties.
	// Further, the first value of OperationalStatus should contain the primary status for the element.
	// When instrumented, Status (because it is single-valued) should also provide the primary status of the element.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, .., 0x8000..}
	//
	// Values={Unknown, Other, OK, Degraded, Stressed, Predictive Failure, Error, Non-Recoverable Error, Starting, Stopping, Stopped, In Service, No Contact, Lost Communication, Aborted, Dormant, Supporting Entity in Error, Completed, Power Mode, Relocating, DMTF Reserved, Vendor Reserved}
	OperationalStatus int
	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	// It can also indicate the transitions between these requested states.
	// For example, shutting down (value=4) and starting (value=10) are transient states between enabled and disabled.
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
	// Values of 32768 and 32769 (relevant from Intel CSME 15 onwards) indicate whether the Intel® One-Click Recovery feature is enabled (32769) or disabled (32768).
	// In Intel CSME 16 onwards, 32768 and 32769 have additional meaning, and 32771 and 32770 indicate whether Intel RPE is enabled or disabled.
	// See Qualifiers below for details.
	//
	// Default: 32769 (Intel One-Click Recovery enabled, Intel RPE is disabled, and all other boot options are enabled.
	// Note that Intel RPE is relevant from Intel CSME 16 onwards.)
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11..32767, 32768, 32769, 32770, 32771, 32772..65535}
	//
	// Values={ Unknown, Other, Enabled, Disabled, Shutting Down, Not Applicable, Enabled but Offline, In Test, Deferred, Quiesce, Starting, DMTF Reserved, Intel One-Click Recovery and Intel RPE are disabled and all other boot options are enabled, Intel One-Click Recovery is enabled and Intel RPE is disabled and all other boot options are enabled, Intel RPE is enabled and Intel One-Click Recovery is disabled and all other boot options are enabled, Intel One-Click Recovery and Intel RPE are enabled and all other boot options are enabled, Vendor Reserved }
	EnabledState int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	// The actual state of the element is represented by EnabledState.
	// This property is provided to compare the last requested and current enabled or disabled states.
	// Note that when EnabledState is set to 5 ("Not Applicable"), then this property has no meaning.
	// Refer to the EnabledState property description for explanations of the values in the RequestedState enumeration.
	//
	// "Unknown" (0) indicates the last requested state for the element is unknown.
	//
	// Note that the value "No Change" (5) has been deprecated in lieu of indicating the last requested state is "Unknown" (0).
	// If the last requested or desired state is unknown, RequestedState should have the value "Unknown" (0), but may have the value "No Change" (5).
	// Offline (6) indicates that the element has been requested to transition to the Enabled but Offline EnabledState.
	//
	// It should be noted that there are two new values in RequestedState that build on the statuses of EnabledState. These are "Reboot" (10) and "Reset" (11). Reboot refers to doing a "Shut Down" and then moving to an "Enabled" state. Reset indicates that the element is first "Disabled" and then "Enabled". The distinction between requesting "Shut Down" and "Disabled" should also be noted. Shut Down requests an orderly transition to the Disabled state, and might involve removing power, to completely erase any existing state. The Disabled state requests an immediate disabling of the element, such that it will not execute or accept any commands or processing requests.
	//
	// This property is set as the result of a method invocation (such as Start or StopService on CIM_Service), or can be overridden and defined as WRITEable in a subclass.
	// The method approach is considered superior to a WRITEable property, because it allows an explicit invocation of the operation and the return of a result code.
	//
	// If knowledge of the last RequestedState is not supported for the EnabledLogicalElement, the property shall be NULL or have the value 12 "Not Applicable".
	//
	// Values of 32768 and 32769 (relevant from Intel CSME 15 onwards) indicate whether the Intel® One-Click Recovery feature is enabled (32769) or disabled (32768). In Intel CSME 16 onwards, 32768 and 32769 have additional meaning, and 32771 and 32770 indicate whether Intel RPE is enabled or disabled. See Qualifiers below for details.
	//
	// Default: 32769 (Intel One-Click Recovery enabled, Intel RPE is disabled, and all other boot options are enabled.
	// Note that Intel RPE is relevant from Intel CSME 16 onwards.)
	//
	// ValueMap={0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, .., 32768, 32769, 32770, 32771, 32772..65535 }
	//
	// Values={ Unknown, Enabled, Disabled, Shut Down, No Change, Offline, Test, Deferred, Quiesce, Reboot, Reset, Not Applicable, DMTF Reserved, Disable Intel One-Click Recovery and Intel RPE and enable all other boot options, disable Intel RPE and enable Intel One-Click Recovery and all other boot options, disable Intel One-Click Recovery and enable Intel RPE and all other boot options, Enable all boot options, Vendor Reserved }
	RequestedState int
)
