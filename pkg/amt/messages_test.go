package amt

import (
	"fmt"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/publickey"
	"github.com/stretchr/testify/assert"
)

const (
	xmlHeader             = `<?xml version="1.0" encoding="utf-8"?>`
	envelope              = `<Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>`
	enumerationContext    = `AC070000-0000-0000-0000-000000000000`
	operationTimeout      = `PT60S`
	GET                   = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	ENUMERATE             = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate"
	PULL                  = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull"
	DELETE                = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete"
	ENUMERATE_BODY        = "<Enumerate xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" />"
	ADD_USER_ACL_ENTRY_EX = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/AddUserAclEntryEx"
	trustedRootCert       = "MIIEOzCCAqOgAwIBAgIDAZiFMA0GCSqGSIb3DQEBDAUAMD0xFzAVBgNVBAMTDk1QU1Jvb3QtNjE0ZDg4MRAwDgYDVQQKEwd1bmtub3duMRAwDgYDVQQGEwd1bmtub3duMCAXDTIwMDgyNTE4MzMzN1oYDzIwNTEwODI1MTgzMzM3WjA9MRcwFQYDVQQDEw5NUFNSb290LTYxNGQ4ODEQMA4GA1UEChMHdW5rbm93bjEQMA4GA1UEBhMHdW5rbm93bjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBAOi1jx9L8DG6gBPxd9gmJ6vqQC/F/TBMTJvb3ZAuRbDxUKnxZk3PafyNM6fO8QTL4qZVhvyGEZaIzVePrdJj31aZ93mNY2TJee3/DLRsJUIZHGFufBvi8pgQL+JjE9JmFD5/S2yciHIEVpKmXo1CbGmZGsnb8NRjaQVwB94pI1mg8JFMxyKzU/cUoCBfI+wmeMgBVdOJPNpH2zjC/GxwEFNQaxGe9GHmYbwoeiDeMPo75E/o+Gw6kJm429cuhJBC3KqHevAJj9V2nSUvoO0oxKqzLVkUYcjHEGYjxIvP6a6uo7x9llwfshJsBZ3PE5hucNdWS3dY3GeCqOwcaAQQIj2jULpZ/KlgVAdBK/o5QjE+IIQXCVK9USvktGzz7I5oH98zy8jCFStbGM7PQCo+DEnHn/SANmVbcy3hjzrXC8zf5dvmKiUb2eKnpv+z3FHsi64sVwFqBArB2ipcTM/qv4nEM6uLW1t+7+NB0OyaBmLktJrpb6af7z/EW1QuPIfTcQIDAQABo0IwQDAMBgNVHRMEBTADAQH/MBEGCWCGSAGG+EIBAQQEAwIABzAdBgNVHQ4EFgQUYU2IeTFqWXI1rG+JqZq8eVDO/LMwDQYJKoZIhvcNAQEMBQADggGBANoKIsFOn8/Lrb98DjOP+LUeopoU9KQ70ndreNqchrkPmM61V9IdD9OZiLr/7OY/rLGZwNvkhQYRPUa842Mqjfpr4YcV6HC0j6Zg0lcpxQ5eGGBkLb/teBcboi3sZcJvbCFUW2DJjhy7uqYxzE4eqSsKx5fEjp/wa6oNzNrgWRXyxQlaOo42RjXnOXS7sB0jPrgO0FClL1Xzif06kFHzzyJCVUqzNEJv0ynLgkpzCVdUUfoMM1RcKc3xJes5C0zg64ugj2R9e4VwJfn9W3+rlYS1So1q1jL8w+3qOM7lXyvr8Bdgc5BMvrOvHxzdOnpZmUEJkbKty62e8fYKN+WP7BrpxnzFQSzczX5S0uN4rn0rLO4wxVf2rtnTqIhKKYTsPMRBVEjpbRT1smzPPdINKu5l/Rz/zZS0b5I4yKJrkTYNgoPC/QSq8A9uXZxxQvj6x1bWZJVWywmaqYolEp8NaVHd+JYnlTmr4XpMHm01TPi1laowtY3ZepnKm8I55Ly0JA=="
)

var PULL_BODY = fmt.Sprintf(`<Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull>`, enumerationContext)

func TestAMT(t *testing.T) {
	messageID := 0
	amtClass := NewMessages()

	expectedResponse := func(method, action, body string) string {
		return fmt.Sprintf(`%s%s%s</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/%s</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>%s</w:OperationTimeout></Header><Body>%s</Body></Envelope>`, xmlHeader, envelope, action, method, messageID, operationTimeout, body)
	}

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_AlarmClockService Get wsman message", "AMT_AlarmClockService", GET, "", amtClass.AlarmClockService.Get},
			{"should create a valid AMT_AuditLog Get wsman message", "AMT_AuditLog", GET, "", amtClass.AuditLog.Get},
			{"should create a valid AMT_AuthorizationService Get wsman message", "AMT_AuthorizationService", GET, "", amtClass.AuthorizationService.Get},
			{"should create a valid AMT_BootCapabilities Get wsman message", "AMT_BootCapabilities", GET, "", amtClass.BootCapabilities.Get},
			{"should create a valid AMT_BootSettingData Get wsman message", "AMT_BootSettingData", GET, "", amtClass.BootSettingData.Get},
			{"should create a valid AMT_EnvironmentDetectionSettingData Get wsman message", "AMT_EnvironmentDetectionSettingData", GET, "", amtClass.EnvironmentDetectionSettingData.Get},
			{"should create a valid AMT_EthernetPortSettings Get wsman message", "AMT_EthernetPortSettings", GET, "", amtClass.EthernetPortSettings.Get},
			{"should create a valid AMT_GeneralSettings Get wsman message", "AMT_GeneralSettings", GET, "", amtClass.GeneralSettings.Get},
			{"should create a valid AMT_IEEE8021xCredentialContext Get wsman message", "AMT_8021xCredentialContext", GET, "", amtClass.IEEE8021xCredentialContext.Get},
			{"should create a valid AMT_IEEE8021xProfile Get wsman message", "AMT_8021XProfile", GET, "", amtClass.IEEE8021xProfile.Get},
			{"should create a valid AMT_KerberosSettingData Get wsman message", "AMT_KerberosSettingData", GET, "", amtClass.KerberosSettingData.Get},
			{"should create a valid AMT_ManagementPresenceRemoteSAP Get wsman message", "AMT_ManagementPresenceRemoteSAP", GET, "", amtClass.ManagementPresenceRemoteSAP.Get},
			{"should create a valid AMT_MessageLog Get wsman message", "AMT_MessageLog", GET, "", amtClass.MessageLog.Get},
			{"should create a valid AMT_MPSUsernamePassword Get wsman message", "AMT_MPSUsernamePassword", GET, "", amtClass.MPSUsernamePassword.Get},
			{"should create a valid AMT_PublicKeyCertificate Get wsman message", "AMT_PublicKeyCertificate", GET, "", amtClass.PublicKeyCertificate.Get},
			{"should create a valid AMT_PublicKeyManagementService Get wsman message", "AMT_PublicKeyManagementService", GET, "", amtClass.PublicKeyManagementService.Get},
			{"should create a valid AMT_PublicPrivateKeyPair Get wsman message", "AMT_PublicPrivateKeyPair", GET, "", amtClass.PublicPrivateKeyPair.Get},
			{"should create a valid AMT_RedirectionService Get wsman message", "AMT_RedirectionService", GET, "", amtClass.RedirectionService.Get},
			{"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Get wsman message", "AMT_RemoteAccessPolicyAppliesToMPS", GET, "", amtClass.RemoteAccessPolicyAppliesToMPS.Get},
			{"should create a valid AMT_RemoteAccessPolicyRule Get wsman message", "AMT_RemoteAccessPolicyRule", GET, "", amtClass.RemoteAccessPolicyRule.Get},
			{"should create a valid AMT_RemoteAccessService Get wsman message", "AMT_RemoteAccessService", GET, "", amtClass.RemoteAccessService.Get},
			{"should create a valid AMT_SetupAndConfigurationService Get wsman message", "AMT_SetupAndConfigurationService", GET, "", amtClass.SetupAndConfigurationService.Get},
			{"should create a valid AMT_TimeSynchronizationService Get wsman message", "AMT_TimeSynchronizationService", GET, "", amtClass.TimeSynchronizationService.Get},
			{"should create a valid AMT_TLSCredentialContext Get wsman message", "AMT_TLSCredentialContext", GET, "", amtClass.TLSCredentialContext.Get},
			{"should create a valid AMT_TLSSettingData Get wsman message", "AMT_TLSSettingData", GET, "", amtClass.TLSSettingData.Get},
			{"should create a valid AMT_UserInitiatedConnectionService Get wsman message", "AMT_UserInitiatedConnectionService", GET, "", amtClass.UserInitiatedConnectionService.Get},
			{"should create a valid AMT_WiFiPortConfigurationService Get wsman message", "AMT_WiFiPortConfigurationService", GET, "", amtClass.WiFiPortConfigurationService.Get},
			//ENUMERATES
			{"should create a valid AMT_AlarmClockService Enumerate wsman message", "AMT_AlarmClockService", ENUMERATE, ENUMERATE_BODY, amtClass.AlarmClockService.Enumerate},
			{"should create a valid AMT_AuditLog Enumerate wsman message", "AMT_AuditLog", ENUMERATE, ENUMERATE_BODY, amtClass.AuditLog.Enumerate},
			{"should create a valid AMT_AuthorizationService Enumerate wsman message", "AMT_AuthorizationService", ENUMERATE, ENUMERATE_BODY, amtClass.AuthorizationService.Enumerate},
			{"should create a valid AMT_BootCapabilities Enumerate wsman message", "AMT_BootCapabilities", ENUMERATE, ENUMERATE_BODY, amtClass.BootCapabilities.Enumerate},
			{"should create a valid AMT_BootSettingData Enumerate wsman message", "AMT_BootSettingData", ENUMERATE, ENUMERATE_BODY, amtClass.BootSettingData.Enumerate},
			{"should create a valid AMT_EnvironmentDetectionSettingData Enumerate wsman message", "AMT_EnvironmentDetectionSettingData", ENUMERATE, ENUMERATE_BODY, amtClass.EnvironmentDetectionSettingData.Enumerate},
			{"should create a valid AMT_EthernetPortSettings Enumerate wsman message", "AMT_EthernetPortSettings", ENUMERATE, ENUMERATE_BODY, amtClass.EthernetPortSettings.Enumerate},
			{"should create a valid AMT_GeneralSettings Enumerate wsman message", "AMT_GeneralSettings", ENUMERATE, ENUMERATE_BODY, amtClass.GeneralSettings.Enumerate},
			{"should create a valid AMT_IEEE8021xCredentialContext Enumerate wsman message", "AMT_8021xCredentialContext", ENUMERATE, ENUMERATE_BODY, amtClass.IEEE8021xCredentialContext.Enumerate},
			{"should create a valid AMT_IEEE8021xProfile Enumerate wsman message", "AMT_8021XProfile", ENUMERATE, ENUMERATE_BODY, amtClass.IEEE8021xProfile.Enumerate},
			{"should create a valid AMT_KerberosSettingData Enumerate wsman message", "AMT_KerberosSettingData", ENUMERATE, ENUMERATE_BODY, amtClass.KerberosSettingData.Enumerate},
			{"should create a valid AMT_ManagementPresenceRemoteSAP Enumerate wsman message", "AMT_ManagementPresenceRemoteSAP", ENUMERATE, ENUMERATE_BODY, amtClass.ManagementPresenceRemoteSAP.Enumerate},
			{"should create a valid AMT_MessageLog Enumerate wsman message", "AMT_MessageLog", ENUMERATE, ENUMERATE_BODY, amtClass.MessageLog.Enumerate},
			{"should create a valid AMT_MPSUsernamePassword Enumerate wsman message", "AMT_MPSUsernamePassword", ENUMERATE, ENUMERATE_BODY, amtClass.MPSUsernamePassword.Enumerate},
			{"should create a valid AMT_PublicKeyCertificate Enumerate wsman message", "AMT_PublicKeyCertificate", ENUMERATE, ENUMERATE_BODY, amtClass.PublicKeyCertificate.Enumerate},
			{"should create a valid AMT_PublicKeyManagementService Enumerate wsman message", "AMT_PublicKeyManagementService", ENUMERATE, ENUMERATE_BODY, amtClass.PublicKeyManagementService.Enumerate},
			{"should create a valid AMT_PublicPrivateKeyPair Enumerate wsman message", "AMT_PublicPrivateKeyPair", ENUMERATE, ENUMERATE_BODY, amtClass.PublicPrivateKeyPair.Enumerate},
			{"should create a valid AMT_RedirectionService Enumerate wsman message", "AMT_RedirectionService", ENUMERATE, ENUMERATE_BODY, amtClass.RedirectionService.Enumerate},
			{"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Enumerate wsman message", "AMT_RemoteAccessPolicyAppliesToMPS", ENUMERATE, ENUMERATE_BODY, amtClass.RemoteAccessPolicyAppliesToMPS.Enumerate},
			{"should create a valid AMT_RemoteAccessPolicyRule Enumerate wsman message", "AMT_RemoteAccessPolicyRule", ENUMERATE, ENUMERATE_BODY, amtClass.RemoteAccessPolicyRule.Enumerate},
			{"should create a valid AMT_RemoteAccessService Enumerate wsman message", "AMT_RemoteAccessService", ENUMERATE, ENUMERATE_BODY, amtClass.RemoteAccessService.Enumerate},
			{"should create a valid AMT_SetupAndConfigurationService Enumerate wsman message", "AMT_SetupAndConfigurationService", ENUMERATE, ENUMERATE_BODY, amtClass.SetupAndConfigurationService.Enumerate},
			{"should create a valid AMT_TimeSynchronizationService Enumerate wsman message", "AMT_TimeSynchronizationService", ENUMERATE, ENUMERATE_BODY, amtClass.TimeSynchronizationService.Enumerate},
			{"should create a valid AMT_TLSCredentialContext Enumerate wsman message", "AMT_TLSCredentialContext", ENUMERATE, ENUMERATE_BODY, amtClass.TLSCredentialContext.Enumerate},
			{"should create a valid AMT_TLSSettingData Enumerate wsman message", "AMT_TLSSettingData", ENUMERATE, ENUMERATE_BODY, amtClass.TLSSettingData.Enumerate},
			{"should create a valid AMT_UserInitiatedConnectionService Enumerate wsman message", "AMT_UserInitiatedConnectionService", ENUMERATE, ENUMERATE_BODY, amtClass.UserInitiatedConnectionService.Enumerate},
			{"should create a valid AMT_WiFiPortConfigurationService Enumerate wsman message", "AMT_WiFiPortConfigurationService", ENUMERATE, ENUMERATE_BODY, amtClass.WiFiPortConfigurationService.Enumerate},
			//PULLS
			{"should create a valid AMT_AlarmClockService Pulls wsman message", "AMT_AlarmClockService", PULL, PULL_BODY, func() string { return amtClass.AlarmClockService.Pull(enumerationContext) }},
			{"should create a valid AMT_AuditLog Pulls wsman message", "AMT_AuditLog", PULL, PULL_BODY, func() string { return amtClass.AuditLog.Pull(enumerationContext) }},
			{"should create a valid AMT_AuthorizationService Pulls wsman message", "AMT_AuthorizationService", PULL, PULL_BODY, func() string { return amtClass.AuthorizationService.Pull(enumerationContext) }},
			{"should create a valid AMT_BootCapabilities Pulls wsman message", "AMT_BootCapabilities", PULL, PULL_BODY, func() string { return amtClass.BootCapabilities.Pull(enumerationContext) }},
			{"should create a valid AMT_BootSettingData Pulls wsman message", "AMT_BootSettingData", PULL, PULL_BODY, func() string { return amtClass.BootSettingData.Pull(enumerationContext) }},
			{"should create a valid AMT_EnvironmentDetectionSettingData Pulls wsman message", "AMT_EnvironmentDetectionSettingData", PULL, PULL_BODY, func() string { return amtClass.EnvironmentDetectionSettingData.Pull(enumerationContext) }},
			{"should create a valid AMT_EthernetPortSettings Pulls wsman message", "AMT_EthernetPortSettings", PULL, PULL_BODY, func() string { return amtClass.EthernetPortSettings.Pull(enumerationContext) }},
			{"should create a valid AMT_GeneralSettings Pulls wsman message", "AMT_GeneralSettings", PULL, PULL_BODY, func() string { return amtClass.GeneralSettings.Pull(enumerationContext) }},
			{"should create a valid AMT_IEEE8021xCredentialContext Pulls wsman message", "AMT_8021xCredentialContext", PULL, PULL_BODY, func() string { return amtClass.IEEE8021xCredentialContext.Pull(enumerationContext) }},
			{"should create a valid AMT_IEEE8021xProfile Pulls wsman message", "AMT_8021XProfile", PULL, PULL_BODY, func() string { return amtClass.IEEE8021xProfile.Pull(enumerationContext) }},
			{"should create a valid AMT_KerberosSettingData Pulls wsman message", "AMT_KerberosSettingData", PULL, PULL_BODY, func() string { return amtClass.KerberosSettingData.Pull(enumerationContext) }},
			{"should create a valid AMT_ManagementPresenceRemoteSAP Pulls wsman message", "AMT_ManagementPresenceRemoteSAP", PULL, PULL_BODY, func() string { return amtClass.ManagementPresenceRemoteSAP.Pull(enumerationContext) }},
			{"should create a valid AMT_MessageLog Pulls wsman message", "AMT_MessageLog", PULL, PULL_BODY, func() string { return amtClass.MessageLog.Pull(enumerationContext) }},
			{"should create a valid AMT_MPSUsernamePassword Pulls wsman message", "AMT_MPSUsernamePassword", PULL, PULL_BODY, func() string { return amtClass.MPSUsernamePassword.Pull(enumerationContext) }},
			{"should create a valid AMT_PublicKeyCertificate Pulls wsman message", "AMT_PublicKeyCertificate", PULL, PULL_BODY, func() string { return amtClass.PublicKeyCertificate.Pull(enumerationContext) }},
			{"should create a valid AMT_PublicKeyManagementService Pulls wsman message", "AMT_PublicKeyManagementService", PULL, PULL_BODY, func() string { return amtClass.PublicKeyManagementService.Pull(enumerationContext) }},
			{"should create a valid AMT_PublicPrivateKeyPair Pulls wsman message", "AMT_PublicPrivateKeyPair", PULL, PULL_BODY, func() string { return amtClass.PublicPrivateKeyPair.Pull(enumerationContext) }},
			{"should create a valid AMT_RedirectionService Pulls wsman message", "AMT_RedirectionService", PULL, PULL_BODY, func() string { return amtClass.RedirectionService.Pull(enumerationContext) }},
			{"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Pulls wsman message", "AMT_RemoteAccessPolicyAppliesToMPS", PULL, PULL_BODY, func() string { return amtClass.RemoteAccessPolicyAppliesToMPS.Pull(enumerationContext) }},
			{"should create a valid AMT_RemoteAccessPolicyRule Pulls wsman message", "AMT_RemoteAccessPolicyRule", PULL, PULL_BODY, func() string { return amtClass.RemoteAccessPolicyRule.Pull(enumerationContext) }},
			{"should create a valid AMT_RemoteAccessService Pulls wsman message", "AMT_RemoteAccessService", PULL, PULL_BODY, func() string { return amtClass.RemoteAccessService.Pull(enumerationContext) }},
			{"should create a valid AMT_SetupAndConfigurationService Pulls wsman message", "AMT_SetupAndConfigurationService", PULL, PULL_BODY, func() string { return amtClass.SetupAndConfigurationService.Pull(enumerationContext) }},
			{"should create a valid AMT_TimeSynchronizationService Pulls wsman message", "AMT_TimeSynchronizationService", PULL, PULL_BODY, func() string { return amtClass.TimeSynchronizationService.Pull(enumerationContext) }},
			{"should create a valid AMT_TLSCredentialContext Pulls wsman message", "AMT_TLSCredentialContext", PULL, PULL_BODY, func() string { return amtClass.TLSCredentialContext.Pull(enumerationContext) }},
			{"should create a valid AMT_TLSSettingData Pulls wsman message", "AMT_TLSSettingData", PULL, PULL_BODY, func() string { return amtClass.TLSSettingData.Pull(enumerationContext) }},
			{"should create a valid AMT_UserInitiatedConnectionService Pulls wsman message", "AMT_UserInitiatedConnectionService", PULL, PULL_BODY, func() string { return amtClass.UserInitiatedConnectionService.Pull(enumerationContext) }},
			{"should create a valid AMT_WiFiPortConfigurationService Pulls wsman message", "AMT_WiFiPortConfigurationService", PULL, PULL_BODY, func() string { return amtClass.WiFiPortConfigurationService.Pull(enumerationContext) }},
			//READ RECORDS
			{"should create a valid AMT_AuditLog Read Records wsman message", "AMT_AuditLog", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog/ReadRecords`, `<h:ReadRecords_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog"><h:StartIndex>1</h:StartIndex></h:ReadRecords_INPUT>`, func() string {
				return amtClass.AuditLog.ReadRecords(1)
			}},
			// AUTHORIZATION SERVICE

			// ADD USER ACL ENTRY EX
			// Verify with Matt - Typescript is referring to wrong realm values
			// {"should return a valid amt_AuthorizationService ADD_USER_ACL_ENTRY_EX wsman message using digest", "AMT_AuthorizationService", ADD_USER_ACL_ENTRY_EX, fmt.Sprintf(`<h:AddUserAclEntryEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:DigestUsername>%s</h:DigestUsername><h:DigestPassword>%s</h:DigestPassword><h:AccessPermission>%d</h:AccessPermission><h:Realms>%d</h:Realms></h:AddUserAclEntryEx_INPUT>`, "test", "P@ssw0rd", 2, 3), func() string {
			// 	return amtClass.AuthorizationService.AddUserAclEntryEx(authorization.AccessPermissionLocalAndNetworkAccess, []authorization.RealmValues{authorization.RedirectionRealm}, "test", "P@ssw0rd", "")
			// }},
			// {"should return a valid amt_AuthorizationService ADD_USER_ACL_ENTRY_EX wsman message using kerberos", "AMT_AuthorizationService", ADD_USER_ACL_ENTRY_EX, fmt.Sprintf(`<h:AddUserAclEntryEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:KerberosUserSid>%d</h:KerberosUserSid><h:AccessPermission>%d</h:AccessPermission><h:Realms>%d3</h:Realms></h:AddUserAclEntryEx_INPUT>`, 64, 2, 3), func() string {
			// 	return amtClass.AuthorizationService.AddUserAclEntryEx(authorization.AccessPermissionLocalAndNetworkAccess, []authorization.RealmValues{authorization.RedirectionRealm}, "", "", "64")
			// }},
			// // Check how to verify for exceptions
			// // {"should throw an error if the digestUsername is longer than 16 when calling AddUserAclEntryEx", "", "", "", func() string {
			// // 	return amtClass.AuthorizationService.AddUserAclEntryEx(2, []models.RealmValues{models.RedirectionRealm}, "thisusernameistoolong", "test", "")
			// // }},
			// // To do
			// //'should throw an error if digest or kerberos credentials are not provided to AddUserAclEntryEx'
			// ENUMERATE USER ACL ENTRIES
			{"should return a valid amt_AuthorizationService EnumerateUserAclEntries wsman message when startIndex is undefined", "AMT_AuthorizationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/EnumerateUserAclEntries`, fmt.Sprintf(`<h:EnumerateUserAclEntries_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:StartIndex>%d</h:StartIndex></h:EnumerateUserAclEntries_INPUT>`, 1), func() string {
				var index int
				return amtClass.AuthorizationService.EnumerateUserAclEntries(index)
			}},
			{"should return a valid amt_AuthorizationService EnumerateUserAclEntries wsman message when startIndex is not 1", "AMT_AuthorizationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/EnumerateUserAclEntries`, fmt.Sprintf(`<h:EnumerateUserAclEntries_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:StartIndex>%d</h:StartIndex></h:EnumerateUserAclEntries_INPUT>`, 50), func() string {
				return amtClass.AuthorizationService.EnumerateUserAclEntries(50)
			}},
			// GET USER ACL ENTRY EX
			{"should return a valid amt_AuthorizationService GetUserAclEntryEx wsman message", "AMT_AuthorizationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetUserAclEntryEx`, `<h:GetUserAclEntryEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:Handle>1</h:Handle></h:GetUserAclEntryEx_INPUT>`, func() string {
				return amtClass.AuthorizationService.GetUserAclEntryEx(1)
			}},
			// UPDATE USER ACL ENTRY EX
			// {"should return a valid amt_AuthorizationService UpdateUserAclEntryEx wsman message using digest", "AMT_AuthorizationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/UpdateUserAclEntryEx`, `<h:GetUserAclEntryEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:Handle>1</h:Handle></h:GetUserAclEntryEx_INPUT>`, func() string {
			// 	return amtClass.AuthorizationService.UpdateUserAclEntryEx(1, 2, []authorization.RealmValues{authorization.RedirectionRealm}, "test", "test123!", "")
			// }},
			// {"should return a valid amt_AuthorizationService UpdateUserAclEntryEx wsman message using kerberos", "AMT_AuthorizationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/UpdateUserAclEntryEx`, `<h:UpdateUserAclEntryEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:Handle>1</h:Handle><h:KerberosUserSid>64</h:KerberosUserSid><h:AccessPermission>2</h:AccessPermission><h:Realms>3</h:Realms></h:UpdateUserAclEntryEx_INPUT>`, func() string {
			// 	return amtClass.AuthorizationService.UpdateUserAclEntryEx(1, 2, []authorization.RealmValues{authorization.RedirectionRealm}, "", "", "64")
			// }},
			// // should throw an error if digest or kerberos credentials are not provided to UpdateUserAclEntryEx
			// // should throw an error if the digestUsername is longer than 16 when calling UpdateUserAclEntryEx

			// REMOVE USER ACL ENTRY
			{"should return a valid amt_AuthorizationService RemoveUserAclEntry wsman message", "AMT_AuthorizationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/RemoveUserAclEntry`, `<h:RemoveUserAclEntry_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:Handle>1</h:Handle></h:RemoveUserAclEntry_INPUT>`, func() string {
				return amtClass.AuthorizationService.RemoveUserAclEntry(1)
			}},

			// GET ADMIN ACL ENTRY
			{"should return a valid amt_AuthorizationService GetAdminAclEntry wsman message", "AMT_AuthorizationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAdminAclEntry`, `<h:GetAdminAclEntry_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"></h:GetAdminAclEntry_INPUT>`, func() string {
				return amtClass.AuthorizationService.GetAdminAclEntry()
			}},

			// GET ADMIN ACL ENTRY STATUS
			{"should return a valid amt_AuthorizationService GetAdminAclEntry wsman message", "AMT_AuthorizationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAdminAclEntryStatus`, `<h:GetAdminAclEntryStatus_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"></h:GetAdminAclEntryStatus_INPUT>`, func() string {
				return amtClass.AuthorizationService.GetAdminAclEntryStatus()
			}},

			// // GET ADMIN NET ACL ENTRY STATUS
			{"should return a valid amt_AuthorizationService GetAdminNetAclEntryStatus wsman message", "AMT_AuthorizationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAdminNetAclEntryStatus`, `<h:GetAdminNetAclEntryStatus_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"></h:GetAdminNetAclEntryStatus_INPUT>`, func() string {
				return amtClass.AuthorizationService.GetAdminNetAclEntryStatus()
			}},

			// // GET ACL ENABLED STATE
			{"should return a valid amt_AuthorizationService GetAclEnabledState wsman message", "AMT_AuthorizationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAclEnabledState`, `<h:GetAclEnabledState_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:Handle>1</h:Handle></h:GetAclEnabledState_INPUT>`, func() string {
				return amtClass.AuthorizationService.GetAclEnabledState(1)
			}},

			// KERBEROS SETTING DATA
			// GET CREDENTIAL CACHE STATE
			{"should return a valid amt_KerberosSettingData GetCredentialCacheState wsman message", "AMT_KerberosSettingData", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData/GetCredentialCacheState`, `<h:GetCredentialCacheState_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData"></h:GetCredentialCacheState_INPUT>`, func() string {
				return amtClass.KerberosSettingData.GetCredentialCacheState()
			}},
			// GET CREDENTIAL CACHE STATE
			// {"should return a valid amt_KerberosSettingData SetCredentialCacheState wsman message", "AMT_KerberosSettingData", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData/SetCredentialCacheState`, `<h:SetCredentialCacheState_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData"></h:SetCredentialCacheState_INPUT>`, func() string {
			// 	return amtClass.KerberosSettingData.SetCredentialCacheState(true)
			// }},

			// MESSAGE LOG
			// POSITION TO FIRST RECORD
			{"should return a valid amt_MessageLog PositionToFirstRecords wsman message", "AMT_MessageLog", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog/PositionToFirstRecord`, `<h:PositionToFirstRecord_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog" />`, func() string {
				return amtClass.MessageLog.PositionToFirstRecord()
			}},
			// GET RECORDS
			{"should return a valid amt_MessageLog GetRecords wsman message", "AMT_MessageLog", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog/GetRecords`, `<h:GetRecords_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog"><h:IterationIdentifier>1</h:IterationIdentifier><h:MaxReadRecords>390</h:MaxReadRecords></h:GetRecords_INPUT>`, func() string {
				return amtClass.MessageLog.GetRecords(1)
			}},
			// PUBLIC KEY MANAGEMENT SERVICE
			{"should return a valid amt_PublicKeyManagementService AddTrustedRootCertificate wsman message", "AMT_PublicKeyManagementService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddTrustedRootCertificate`, fmt.Sprintf(`<h:AddTrustedRootCertificate_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:CertificateBlob>%s</h:CertificateBlob></h:AddTrustedRootCertificate_INPUT>`, trustedRootCert), func() string {
				return amtClass.PublicKeyManagementService.AddTrustedRootCertificate(trustedRootCert)
			}},

			{"should return a valid amt_PublicKeyManagementService GenerateKeyPair wsman message", "AMT_PublicKeyManagementService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GenerateKeyPair`, `<h:GenerateKeyPair_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyAlgorithm>0</h:KeyAlgorithm><h:KeyLength>2048</h:KeyLength></h:GenerateKeyPair_INPUT>`, func() string {
				params := publickey.GenerateKeyPair_INPUT{
					KeyAlgorithm: 0,
					KeyLength:    2048,
				}
				return amtClass.PublicKeyManagementService.GenerateKeyPair(params)
			}},

			{"should return a valid amt_PublicKeyManagementService AddCertificate wsman message", "AMT_PublicKeyManagementService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddCertificate`, fmt.Sprintf(`<h:AddCertificate_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:CertificateBlob>%s</h:CertificateBlob></h:AddCertificate_INPUT>`, trustedRootCert), func() string {
				return amtClass.PublicKeyManagementService.AddCertificate(trustedRootCert)
			}},

			{"should return a valid amt_PublicKeyManagementService GeneratePKCS10RequestEx wsman message", "AMT_PublicKeyManagementService", "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GeneratePKCS10RequestEx", `<h:GeneratePKCS10RequestEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyPair>test</h:KeyPair><h:NullSignedCertificateRequest>reallylongcertificateteststring</h:NullSignedCertificateRequest><h:SigningAlgorithm>1</h:SigningAlgorithm></h:GeneratePKCS10RequestEx_INPUT>`, func() string {
				pkcs10Request := publickey.PKCS10Request{
					KeyPair:                      "test",
					NullSignedCertificateRequest: "reallylongcertificateteststring",
					SigningAlgorithm:             1,
				}
				return amtClass.PublicKeyManagementService.GeneratePKCS10RequestEx(pkcs10Request)
			}},

			// WIFI PORT CONFIGURATION SERVICE
			// {"should return a valid amt_WiFiPortConfigurationService ADD_WIFI_SETTINGS wsman message", "AMT_WiFiPortConfigurationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService/AddWiFiSettings`, `<h:AddWiFiSettings_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService"><h:WiFiEndpoint><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpoint</w:ResourceURI><w:SelectorSet><w:Selector Name="${selector.name}">${selector.value}</w:Selector></w:SelectorSet></a:ReferenceParameters></h:WiFiEndpoint><h:WiFiEndpointSettingsInput xmlns:q="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings"><h:ElementName>${wifiEndpointSettings.ElementName}</h:ElementName><h:InstanceID>${wifiEndpointSettings.InstanceID}</h:InstanceID><h:AuthenticationMethod>${wifiEndpointSettings.AuthenticationMethod}</h:AuthenticationMethod><h:EncryptionMethod>${wifiEndpointSettings.EncryptionMethod}</h:EncryptionMethod><h:SSID>${wifiEndpointSettings.SSID}</h:SSID><h:Priority>${wifiEndpointSettings.Priority}</h:Priority><h:PSKPassPhrase>p&apos;ass&lt;&gt;&amp;&quot;code</h:PSKPassPhrase></h:WiFiEndpointSettingsInput></h:AddWiFiSettings_INPUT>`, func() string {
			// 	ieee8021xSettingsInput := &cimModels.IEEE8021xSettings{}
			// 	var clientCredential string
			// 	var caCredential string
			// 	selector := wsman.Selector{
			// 		Name:  "Name",
			// 		Value: "WiFi Endpoint 0",
			// 	}
			// 	wifiEndpointSettings := cimModels.WiFiEndpointSettings{
			// 		ElementName:          "home",
			// 		InstanceID:           "Intel(r) AMT:WiFi Endpoint Settings home",
			// 		AuthenticationMethod: 6,
			// 		EncryptionMethod:     4,
			// 		SSID:                 "admin",
			// 		Priority:             1,
			// 		PSKPassPhrase:        `p\'ass<>&"code`,
			// 	}
			// 	return amtClass.WiFiPortConfigurationService.AddWiFiSettings(wifiEndpointSettings, selector, ieee8021xSettingsInput, clientCredential, caCredential)
			// }},

			// {"should create a valid AMT_WiFiPortConfigurationService Pulls wsman message", "AMT_WiFiPortConfigurationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService/AddWiFiSettings`, `<h:AddWiFiSettings_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService"><h:WiFiEndpoint><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpoint</w:ResourceURI><w:SelectorSet><w:Selector Name="${selector.name}">${selector.value}</w:Selector></w:SelectorSet></a:ReferenceParameters></h:WiFiEndpoint><h:WiFiEndpointSettingsInput xmlns:q="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings"><h:ElementName>${wifiEndpointSettings.ElementName}</h:ElementName><h:InstanceID>${wifiEndpointSettings.InstanceID}</h:InstanceID><h:AuthenticationMethod>${wifiEndpointSettings.AuthenticationMethod}</h:AuthenticationMethod><h:EncryptionMethod>${wifiEndpointSettings.EncryptionMethod}</h:EncryptionMethod><h:SSID>${wifiEndpointSettings.SSID}</h:SSID><h:Priority>${wifiEndpointSettings.Priority}</h:Priority><h:PSKPassPhrase>p&apos;ass&lt;&gt;&amp;&quot;code</h:PSKPassPhrase></h:WiFiEndpointSettingsInput><h:ieee8021xSettingsInput xmlns:q="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_IEEE8021xSettings"><h:ElementName>wifi_8021x_profile</h:ElementName><h:AuthenticationProtocol>0</h:AuthenticationProtocol></h:ieee8021xSettingsInput><h:ClientCredential>handle 0</h:ClientCredential><h:CACredential>handle 1</h:CACredential></h:AddWiFiSettings_INPUT>`, func() string {
			// 	selector := wsman.Selector{
			// 		Name:  "Name",
			// 		Value: "WiFi Endpoint 0",
			// 	}
			// 	wifiEndpointSettings := cimModels.WiFiEndpointSettings{
			// 		ElementName:          "home",
			// 		InstanceID:           "Intel(r) AMT:WiFi Endpoint Settings home",
			// 		AuthenticationMethod: 6,
			// 		EncryptionMethod:     4,
			// 		SSID:                 "admin",
			// 		Priority:             1,
			// 		PSKPassPhrase:        "p'ass<>&\"code",
			// 	}
			// 	ieee8021xSettingsInput := &cimModels.IEEE8021xSettings{
			// 		ElementName:            "wifi_8021x_profile",
			// 		AuthenticationProtocol: 0,
			// 	}
			// 	clientCredential := "handle 0"
			// 	caCredential := "handle 1"
			// 	return amtClass.WiFiPortConfigurationService.AddWiFiSettings(wifiEndpointSettings, selector, ieee8021xSettingsInput, clientCredential, caCredential)
			// }},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				correctResponse := expectedResponse(test.method, test.action, test.body)
				messageID++
				response := test.responseFunc()
				if response != correctResponse {
					assert.Equal(t, correctResponse, response)
				}
			})
		}
	})
}
