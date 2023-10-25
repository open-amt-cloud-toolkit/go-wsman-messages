/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package actions

import "fmt"

type Actions string

const (
	ReadRecords               Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog/ReadRecords"
	AddTrustedRootCertificate Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddTrustedRootCertificate"
	AddCertificate            Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddCertificate"
	GenerateKeyPair           Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GenerateKeyPair"
	AddKey                    Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddKey"
	AddMps                    Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddMpServer"
	AddRemoteAccessPolicyRule Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddRemoteAccessPolicyRule"
	SetBootConfigRole         Actions = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService/SetBootConfigRole"
	GetRecords                Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog/GetRecords"
	PositionToFirstRecord     Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog/PositionToFirstRecord"
	CommitChanges             Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService/CommitChanges"
	Unprovision               Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService/Unprovision"
	SetMEBxPassword           Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService/SetMEBxPassword"
	GetUuid                   Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService/GetUuid"
	SetAdminAclEntryEx        Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/SetAdminAclEntryEx"
	AddUserAclEntryEx         Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/AddUserAclEntryEx"
	GetLowAccuracyTimeSynch   Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService/GetLowAccuracyTimeSynch"
	SetHighAccuracyTimeSynch  Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService/SetHighAccuracyTimeSynch"
	AddWiFiSettings           Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService/AddWiFiSettings"
	AddAlarm                  Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService/AddAlarm"
	GeneratePKCS10RequestEx   Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GeneratePKCS10RequestEx"
	GetCredentialCacheState   Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData/GetCredentialCacheState"
	SetCredentialCacheState   Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData/SetCredentialCacheState"
	EnumerateUserAclEntries   Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/EnumerateUserAclEntries"
	GetUserAclEntryEx         Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetUserAclEntryEx"
	UpdateUserAclEntryEx      Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/UpdateUserAclEntryEx"
	RemoveUserAclEntry        Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/RemoveUserAclEntry"
	GetAdminAclEntry          Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAdminAclEntry"
	GetAdminAclEntryStatus    Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAdminAclEntryStatus"
	GetAdminNetAclEntryStatus Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAdminNetAclEntryStatus"
	SetAclEnabledState        Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/SetAclEnabledState"
	GetAclEnabledState        Actions = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAclEnabledState"
)

func RequestStateChange(className string) string {
	return fmt.Sprintf("http://schemas.dmtf.org/wbem/wscim/1/amt-schema/2/%s/RequestStateChange", className)
}
