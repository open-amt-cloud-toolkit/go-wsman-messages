/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package authorization

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type AuthorizationService struct {
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
		XMLName           xml.Name `xml:"Body"`
		GetResponse       AuthorizationOccurrence
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	AuthorizationOccurrence struct {
		XMLName                 xml.Name       `xml:"AMT_AuthorizationService"`
		AllowHttpQopAuthOnly    int            `xml:"AllowHttpQopAuthOnly"`    // Indicates whether using the http "quality of protection" (qop) directive with value auth is allowed
		CreationClassName       string         `xml:"CreationClassName"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified. In Intel AMT Release 6.0 and later releases value is 'AMT_AuthorizationService'
		ElementName             string         `xml:"ElementName"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information.  Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledState            EnabledState   `xml:"EnabledState"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		Name                    string         `xml:"Name"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.  In Intel AMT Release 6.0 and later releases value is 'Intel® AMT Authorization Service'
		RequestedState          RequestedState `xml:"RequestedState"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		SystemCreationClassName string         `xml:"SystemCreationClassName"` // The CreationClassName of the scoping System. In Intel AMT Release 6.0 and later releases value is 'CIM_ComputerSystem'
		SystemName              string         `xml:"SystemName"`              // The Name of the scoping System.  In Intel AMT Release 6.0 and later releases value is 'Intel® AMT'
	}
	PullResponse struct {
		XMLName                      xml.Name                  `xml:"PullResponse"`
		AuthorizationOccurrenceItems []AuthorizationOccurrence `xml:"Items>AMT_AuthorizationService"`
	}
)
type AddUserAclEntry struct {
	XMLName          xml.Name         `xml:"h:AddUserAclEntryEx_INPUT"`
	H                string           `xml:"xmlns:h,attr"`
	Handle           int              `xml:"h:Handle,omitempty"`              // Contains a creation handle.
	DigestUsername   string           `xml:"h:DigestUsername"`                // Username for access control. Contains 7-bit ASCII characters. String length is limited to 16 characters. Username cannot be an empty string.
	DigestPassword   string           `xml:"h:DigestPassword"`                // An MD5 Hash of these parameters concatenated together (Username + ":" + DigestRealm + ":" + Password). The DigestRealm is a field in AMT_GeneralSettings
	AccessPermission AccessPermission `xml:"h:AccessPermission"`              // Indicates whether the User is allowed to access Intel® AMT from the Network or Local Interfaces. Note: this definition is restricted by the Default Interface Access Permissions of each Realm.
	Realms           []RealmValues    `xml:"h:Realms>h:RealmValue,omitempty"` // Array of interface names the ACL entry is allowed to access.
	KerberosUserSid  string           `xml:"h:KerberosUserSid"`               // Descriptor for user (SID) which is authenticated using the Kerberos Authentication. Byte array, specifying the Security Identifier (SID) according to the Kerberos specification. Current requirements imply that SID should be not smaller than 1 byte length and no longer than 28 bytes. SID length should also be a multiplicand of 4.
}
type UpdateUserAclEntry struct {
	XMLName          xml.Name         `xml:"h:UpdateUserAclEntry_INPUT"`
	H                string           `xml:"xmlns:h,attr"`
	Handle           int              `xml:"h:Handle,omitempty"`              // Contains a creation handle.
	DigestUsername   string           `xml:"h:DigestUsername"`                // Username for access control. Contains 7-bit ASCII characters. String length is limited to 16 characters. Username cannot be an empty string.
	DigestPassword   string           `xml:"h:DigestPassword"`                // An MD5 Hash of these parameters concatenated together (Username + ":" + DigestRealm + ":" + Password). The DigestRealm is a field in AMT_GeneralSettings
	AccessPermission AccessPermission `xml:"h:AccessPermission"`              // Indicates whether the User is allowed to access Intel® AMT from the Network or Local Interfaces. Note: this definition is restricted by the Default Interface Access Permissions of each Realm.
	Realms           []RealmValues    `xml:"h:Realms>h:RealmValue,omitempty"` // Array of interface names the ACL entry is allowed to access.
	KerberosUserSid  string           `xml:"h:KerberosUserSid"`               // Descriptor for user (SID) which is authenticated using the Kerberos Authentication. Byte array, specifying the Security Identifier (SID) according to the Kerberos specification. Current requirements imply that SID should be not smaller than 1 byte length and no longer than 28 bytes. SID length should also be a multiplicand of 4.
}

// ValueMap={0, 1, 2}
//
// Values={LocalAccessPermission, NetworkAccessPermission, AnyAccessPermission}
type AccessPermission int

// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, ..}
//
// Values={InvalidRealm, ReservedRealm0, RedirectionRealm, PTAdministrationRealm, HardwareAssetRealm, RemoteControlRealm, StorageRealm, EventManagerRealm, StorageAdminRealm, AgentPresenceLocalRealm, AgentPresenceRemoteRealm, CircuitBreakerRealm, NetworkTimeRealm, GeneralInfoRealm, FirmwareUpdateRealm, EITRealm, LocalUN, EndpointAccessControlRealm, EndpointAccessControlAdminRealm, EventLogReaderRealm, AuditLogRealm, ACLRealm, ReservedRealm1, ReservedRealm2, LocalSystemRealm, Reserved}
type RealmValues int

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
type EnabledState int

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
type RequestedState int

// ValueMap={0, 1, 12, 16, 23, 38, 2054, 2055, 2058, 2065, 2075}
//
// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_INVALID_NAME, PT_STATUS_NOT_PERMITTED, PT_STATUS_MAX_LIMIT_REACHED, PT_STATUS_FLASH_WRITE_LIMIT_EXCEEDED, PT_STATUS_INVALID_PASSWORD, PT_STATUS_INVALID_REALM, AMT_STATUS_DUPLICATE, PT_STATUS_MAX_KERB_DOMAIN_REACHED, PT_STATUS_AUDIT_FAIL}
type PTStatus int

// INPUTS
// Request Types
type (
	EnumerateUserAclEntries_INPUT struct {
		XMLName    xml.Name `xml:"h:EnumerateUserAclEntries_INPUT"`
		H          string   `xml:"xmlns:h,attr"`
		StartIndex int      `xml:"h:StartIndex"` // Indicates the first ACL entry to retrieve. To enumerate the entire list, an application sends this message with StartIndex set to 1.
	}

	GetAclEnabledState_INPUT struct {
		XMLName xml.Name `xml:"h:GetAclEnabledState_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"` // Specifies the ACL entry to fetch.
	}
	GetUserAclEntryEx_INPUT struct {
		XMLName xml.Name `xml:"h:GetUserAclEntryEx_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"` // Specifies the ACL entry to fetch.
	}
	RemoveUserAclEntry_INPUT struct {
		XMLName xml.Name `xml:"h:RemoveUserAclEntry_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"` // Specifies the ACL entry to be removed.
	}

	SetAclEnabledState_INPUT struct {
		XMLName xml.Name `xml:"h:SetAclEnabledState_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"`  // Specifies the ACL entry to update
		Enabled bool     `xml:"h:Enabled"` // Specifies the state of the ACL entry
	}

	SetAdminACLEntryEx_INPUT struct {
		XMLName        xml.Name `xml:"h:SetAdminACLEntryEx_INPUT"`
		H              string   `xml:"xmlns:h,attr"`
		Username       string   `xml:"h:Username"`       // Username for access control. Contains 7-bit ASCII characters. String length is limited to 16 characters. Username cannot be an empty string.
		DigestPassword string   `xml:"h:DigestPassword"` // An MD5 Hash of these parameters concatenated together (Username + ":" + DigestRealm + ":" + Password). The DigestRealm is a field in AMT_GeneralSettings
	}
)
