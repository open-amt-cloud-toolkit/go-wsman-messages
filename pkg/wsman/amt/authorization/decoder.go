/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package authorization

const ValueNotFound string = "Value not found in map"

// INPUTS Constants.
const (
	AMTAuthorizationService   string = "AMT_AuthorizationService"
	EnumerateUserACLEntries   string = "EnumerateUserAclEntries"
	GetUserACLEntryEx         string = "GetUserAclEntryEx"
	UpdateUserACLEntryEx      string = "UpdateUserAclEntryEx"
	RemoveUserACLEntry        string = "RemoveUserAclEntry"
	GetAdminACLEntry          string = "GetAdminAclEntry"
	GetAdminACLEntryStatus    string = "GetAdminAclEntryStatus"
	GetAdminNetACLEntryStatus string = "GetAdminNetAclEntryStatus"
	SetACLEnabledState        string = "SetAclEnabledState"
	GetACLEnabledState        string = "GetAclEnabledState"
	SetAdminACLEntryEx        string = "SetAdminAclEntryEx"
	AddUserACLEntryEx         string = "AddUserAclEntryEx"
)

const (
	AccessPermissionLocalAccessOnly AccessPermission = iota
	AccessPermissionNetworkAccessOnly
	AccessPermissionLocalAndNetworkAccess
)

const (
	RealmValuesInvalidRealm RealmValues = iota
	RealmValuesReservedRealm0
	RealmValuesRedirectionRealm
	RealmValuesPTAdministrationRealm
	RealmValuesHardwareAssetRealm
	RealmValuesRemoteControlRealm
	RealmValuesStorageRealm
	RealmValuesEventManagerRealm
	RealmValuesStorageAdminRealm
	RealmValuesAgentPresenceLocalRealm
	RealmValuesAgentPresenceRemoteRealm
	RealmValuesCircuitBreakerRealm
	RealmValuesNetworkTimeRealm
	RealmValuesGeneralInfoRealm
	RealmValuesFirmwareUpdateRealm
	RealmValuesEITRealm
	RealmValuesLocalUN
	RealmValuesEndpointAccessControlRealm
	RealmValuesEndpointAccessControlAdminRealm
	RealmValuesEventLogReaderRealm
	RealmValuesAuditLogRealm
	RealmValuesACLRealm
	RealmValuesReservedRealm1
	RealmValuesReservedRealm2
	RealmValuesLocalSystemRealm
)

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)

// enabledStateStrings is a map of EnabledState values to their string representations.
var enabledStateToString = map[EnabledState]string{
	EnabledStateUnknown:           "Unknown",
	EnabledStateOther:             "Other",
	EnabledStateEnabled:           "Enabled",
	EnabledStateDisabled:          "Disabled",
	EnabledStateShuttingDown:      "ShuttingDown",
	EnabledStateNotApplicable:     "NotApplicable",
	EnabledStateEnabledButOffline: "EnabledButOffline",
	EnabledStateInTest:            "InTest",
	EnabledStateDeferred:          "Deferred",
	EnabledStateQuiesce:           "Quiesce",
	EnabledStateStarting:          "Starting",
}

// String returns the string representation of an EnabledState value.
func (e EnabledState) String() string {
	if value, exists := enabledStateToString[e]; exists {
		return value
	}

	return ValueNotFound
}

const (
	RequestedStateUnknown RequestedState = iota
	RequestedStateEnabled
	RequestedStateDisabled
	RequestedStateShutDown
	RequestedStateNoChange
	RequestedStateOffline
	RequestedStateTest
	RequestedStateDeferred
	RequestedStateQuiesce
	RequestedStateReboot
	RequestedStateReset
	RequestedStateNotApplicable
)

// RequestedStateStrings is a map of RequestedState values to their string representations.
var requestedStateToString = map[RequestedState]string{
	RequestedStateUnknown:       "Unknown",
	RequestedStateEnabled:       "Enabled",
	RequestedStateDisabled:      "Disabled",
	RequestedStateShutDown:      "ShutDown",
	RequestedStateNoChange:      "NoChange",
	RequestedStateOffline:       "Offline",
	RequestedStateTest:          "Test",
	RequestedStateDeferred:      "Deferred",
	RequestedStateQuiesce:       "Quiesce",
	RequestedStateReboot:        "Reboot",
	RequestedStateReset:         "Reset",
	RequestedStateNotApplicable: "NotApplicable",
}

// String returns the string representation of a RequestedState value.
func (r RequestedState) String() string {
	if value, exists := requestedStateToString[r]; exists {
		return value
	}

	return ValueNotFound
}

const (
	PTStatusSuccess                 ReturnValue = 0
	PTStatusInternalError           ReturnValue = 1
	PTStatusInvalidName             ReturnValue = 12
	PTStatusNotPermitted            ReturnValue = 16
	PTStatusMaxLimitReached         ReturnValue = 23
	PTStatusInvalidIndex            ReturnValue = 35
	PTStatusFlashWriteLimitExceeded ReturnValue = 38
	PTStatusInvalidHandle           ReturnValue = 2053
	PTStatusInvalidPassword         ReturnValue = 2054
	PTStatusInvalidRealm            ReturnValue = 2055
	AMTStatusDuplicate              ReturnValue = 2058
	PTStatusMaxKerbDomainReached    ReturnValue = 2065
	PTStatusAuditFail               ReturnValue = 2075
)

// returnValuesToString is a map of return values to their string representation.
var returnValuesToString = map[ReturnValue]string{
	PTStatusSuccess:                 "Success",
	PTStatusInternalError:           "InternalError",
	PTStatusInvalidName:             "InvalidName",
	PTStatusNotPermitted:            "NotPermitted",
	PTStatusMaxLimitReached:         "MaxLimitReached",
	PTStatusInvalidIndex:            "InvalidIndex",
	PTStatusFlashWriteLimitExceeded: "FlashWriteLimitExceeded",
	PTStatusInvalidHandle:           "InvalidHandle",
	PTStatusInvalidPassword:         "InvalidPassword",
	PTStatusInvalidRealm:            "InvalidRealm",
	AMTStatusDuplicate:              "Duplicate",
	PTStatusMaxKerbDomainReached:    "MaxKerbDomainReached",
	PTStatusAuditFail:               "AuditFail",
}

// String returns the string representation of the return value.
func (r ReturnValue) String() string {
	if value, exists := returnValuesToString[r]; exists {
		return value
	}

	return ValueNotFound
}
