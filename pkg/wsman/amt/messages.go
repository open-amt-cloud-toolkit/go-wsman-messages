/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package amt

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/alarmclock"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/auditlog"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/authorization"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/boot"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/environmentdetection"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/ethernetport"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/general"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/ieee8021x"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/kerberos"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/managementpresence"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/messagelog"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/mps"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/publickey"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/publicprivate"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/redirection"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/remoteaccess"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/setupandconfiguration"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/timesynchronization"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/tls"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/userinitiatedconnection"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/wifiportconfiguration"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

type Messages struct {
	wsmanMessageCreator             *message.WSManMessageCreator
	AlarmClockService               alarmclock.Service
	AuditLog                        auditlog.Service
	AuthorizationService            authorization.AuthorizationService
	BootCapabilities                boot.Capabilities
	BootSettingData                 boot.SettingData
	EnvironmentDetectionSettingData environmentdetection.SettingData
	EthernetPortSettings            ethernetport.Settings
	GeneralSettings                 general.Settings
	IEEE8021xCredentialContext      ieee8021x.CredentialContext
	IEEE8021xProfile                ieee8021x.Profile
	KerberosSettingData             kerberos.SettingData
	ManagementPresenceRemoteSAP     managementpresence.RemoteSAP
	MessageLog                      messagelog.MessageLog
	MPSUsernamePassword             mps.UsernamePassword
	PublicKeyCertificate            publickey.Certificate
	PublicKeyManagementService      publickey.ManagementService
	PublicPrivateKeyPair            publicprivate.KeyPair
	RedirectionService              redirection.Service
	RemoteAccessPolicyAppliesToMPS  remoteaccess.PolicyAppliesToMPS
	RemoteAccessPolicyRule          remoteaccess.PolicyRule
	RemoteAccessService             remoteaccess.Service
	SetupAndConfigurationService    setupandconfiguration.Service
	TimeSynchronizationService      timesynchronization.Service
	TLSCredentialContext            tls.CredentialContext
	TLSProtocolEndpointCollection   tls.Collection
	TLSSettingData                  tls.SettingData
	UserInitiatedConnectionService  userinitiatedconnection.Service
	WiFiPortConfigurationService    wifiportconfiguration.Service
}

func NewMessages(client client.WSMan) Messages {
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	m := Messages{
		wsmanMessageCreator: wsmanMessageCreator,
	}
	m.AlarmClockService = alarmclock.NewServiceWithClient(wsmanMessageCreator, client)
	m.AuditLog = auditlog.NewAuditLogWithClient(wsmanMessageCreator, client)
	m.AuthorizationService = authorization.NewServiceWithClient(wsmanMessageCreator, client)
	m.BootCapabilities = boot.NewBootCapabilitiesWithClient(wsmanMessageCreator, client)
	m.BootSettingData = boot.NewBootSettingDataWithClient(wsmanMessageCreator, client)
	m.EnvironmentDetectionSettingData = environmentdetection.NewEnvironmentDetectionSettingDataWithClient(wsmanMessageCreator, client)
	m.EthernetPortSettings = ethernetport.NewEthernetPortSettingsWithClient(wsmanMessageCreator, client)
	m.GeneralSettings = general.NewGeneralSettingsWithClient(wsmanMessageCreator, client)
	m.IEEE8021xCredentialContext = ieee8021x.NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator, client)
	m.IEEE8021xProfile = ieee8021x.NewIEEE8021xProfileWithClient(wsmanMessageCreator, client)
	m.KerberosSettingData = kerberos.NewKerberosSettingDataWithClient(wsmanMessageCreator, client)
	m.ManagementPresenceRemoteSAP = managementpresence.NewManagementPresenceRemoteSAPWithClient(wsmanMessageCreator, client)
	m.MessageLog = messagelog.NewMessageLogWithClient(wsmanMessageCreator, client)
	m.MPSUsernamePassword = mps.NewMPSUsernamePasswordWithClient(wsmanMessageCreator, client)
	m.PublicKeyCertificate = publickey.NewPublicKeyCertificateWithClient(wsmanMessageCreator, client)
	m.PublicKeyManagementService = publickey.NewPublicKeyManagementServiceWithClient(wsmanMessageCreator, client)
	m.PublicPrivateKeyPair = publicprivate.NewPublicPrivateKeyPairWithClient(wsmanMessageCreator, client)
	m.RedirectionService = redirection.NewRedirectionServiceWithClient(wsmanMessageCreator, client)
	m.RemoteAccessPolicyAppliesToMPS = remoteaccess.NewRemoteAccessPolicyAppliesToMPSWithClient(wsmanMessageCreator, client)
	m.RemoteAccessPolicyRule = remoteaccess.NewPolicyRuleWithClient(wsmanMessageCreator, client)
	m.RemoteAccessService = remoteaccess.NewRemoteAccessServiceWithClient(wsmanMessageCreator, client)
	m.SetupAndConfigurationService = setupandconfiguration.NewSetupAndConfigurationServiceWithClient(wsmanMessageCreator, client)
	m.TimeSynchronizationService = timesynchronization.NewTimeSynchronizationServiceWithClient(wsmanMessageCreator, client)
	m.TLSCredentialContext = tls.NewTLSCredentialContextWithClient(wsmanMessageCreator, client)
	m.TLSProtocolEndpointCollection = tls.NewTLSProtocolEndpointCollectionWithClient(wsmanMessageCreator, client)
	m.TLSSettingData = tls.NewTLSSettingDataWithClient(wsmanMessageCreator, client)
	m.UserInitiatedConnectionService = userinitiatedconnection.NewUserInitiatedConnectionServiceWithClient(wsmanMessageCreator, client)
	m.WiFiPortConfigurationService = wifiportconfiguration.NewWiFiPortConfigurationServiceWithClient(wsmanMessageCreator, client)
	return m
}
