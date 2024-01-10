/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package authorization

const (
	AMT_AuthorizationService  string = "AMT_AuthorizationService"
	EnumerateUserAclEntries   string = "EnumerateUserAclEntries"
	GetUserAclEntryEx         string = "GetUserAclEntryEx"
	UpdateUserAclEntryEx      string = "UpdateUserAclEntryEx"
	RemoveUserAclEntry        string = "RemoveUserAclEntry"
	GetAdminAclEntry          string = "GetAdminAclEntry"
	GetAdminAclEntryStatus    string = "GetAdminAclEntryStatus"
	GetAdminNetAclEntryStatus string = "GetAdminNetAclEntryStatus"
	SetAclEnabledState        string = "SetAclEnabledState"
	GetAclEnabledState        string = "GetAclEnabledState"
	SetAdminAclEntryEx        string = "SetAdminAclEntryEx"
	AddUserAclEntryEx         string = "AddUserAclEntryEx"
)

const (
	PTStatusSuccess                 PTStatus = 0
	PTStatusInternalError           PTStatus = 1
	PTStatusInvalidName             PTStatus = 12
	PTStatusNotPermitted            PTStatus = 16
	PTStatusMaxLimitReached         PTStatus = 23
	PTStatusInvalidIndex            PTStatus = 35
	PTStatusFlashWriteLimitExceeded PTStatus = 38
	PTStatusInvalidHandle           PTStatus = 2053
	PTStatusInvalidPassword         PTStatus = 2054
	PTStatusInvalidRealm            PTStatus = 2055
	AMTStatusDuplicate              PTStatus = 2058
	PTStatusMaxKerbDomainReached    PTStatus = 2065
	PTStatusAuditFail               PTStatus = 2075
)

const (
	LocalAccessOnly AccessPermission = iota
	NetworkAccessOnly
	LocalAndNetworkAccess
)

const (
	InvalidRealm RealmValues = iota
	ReservedRealm0
	RedirectionRealm
	PTAdministrationRealm
	HardwareAssetRealm
	RemoteControlRealm
	StorageRealm
	EventManagerRealm
	StorageAdminRealm
	AgentPresenceLocalRealm
	AgentPresenceRemoteRealm
	CircuitBreakerRealm
	NetworkTimeRealm
	GeneralInfoRealm
	FirmwareUpdateRealm
	EITRealm
	LocalUN
	EndpointAccessControlRealm
	EndpointAccessControlAdminRealm
	EventLogReaderRealm
	AuditLogRealm
	ACLRealm
	ReservedRealm1
	ReservedRealm2
	LocalSystemRealm
)

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledbutOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)

const (
	RequestedStateUnknown       RequestedState = 0
	RequestedStateEnabled       RequestedState = 2
	RequestedStateDisabled      RequestedState = 3
	RequestedStateShutDown      RequestedState = 4
	RequestedStateNoChange      RequestedState = 5
	RequestedStateOffline       RequestedState = 6
	RequestedStateTest          RequestedState = 7
	RequestedStateDeferred      RequestedState = 8
	RequestedStateQuiesce       RequestedState = 9
	RequestedStateReboot        RequestedState = 10
	RequestedStateReset         RequestedState = 11
	RequestedStateNotApplicable RequestedState = 12
)
