/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package amt

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/alarmclock"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/auditlog"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/authorization"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/boot"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/environmentdetection"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/ethernetport"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/general"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/ieee8021x"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/kerberos"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/managementpresence"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/messagelog"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/mps"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/publickey"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/publicprivate"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/redirection"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/remoteaccess"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/setupandconfiguration"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/timesynchronization"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/tls"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/userinitiatedconnection"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/wifiportconfiguration"
)

type Messages struct {
	wsmanMessageCreator             *wsman.WSManMessageCreator
	AlarmClockService               alarmclock.Service
	AuditLog                        auditlog.AuditLog
	AuthorizationService            authorization.AuthorizationService
	BootCapabilities                boot.BootCapabilities
	BootSettingData                 boot.SettingData
	EnvironmentDetectionSettingData environmentdetection.SettingData
	EthernetPortSettings            ethernetport.Settings
	GeneralSettings                 general.Settings
	IEEE8021xCredentialContext      ieee8021x.CredentialContext
	IEEE8021xProfile                ieee8021x.Profile
	KerberosSettingData             kerberos.KerberosSettingData
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
	TLSSettingData                  tls.SettingData
	UserInitiatedConnectionService  userinitiatedconnection.Service
	WiFiPortConfigurationService    wifiportconfiguration.Service
}

func NewMessages() Messages {
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := wsman.NewWSManMessageCreator(resourceUriBase)
	m := Messages{
		wsmanMessageCreator: wsmanMessageCreator,
	}
	m.AlarmClockService = alarmclock.NewService(wsmanMessageCreator)
	m.AuditLog = auditlog.NewAuditLog(wsmanMessageCreator)
	m.AuthorizationService = authorization.NewAuthorizationService(wsmanMessageCreator)
	m.BootCapabilities = boot.NewBootCapabilities(wsmanMessageCreator)
	m.BootSettingData = boot.NewBootSettingData(wsmanMessageCreator)
	m.EnvironmentDetectionSettingData = environmentdetection.NewEnvironmentDetectionSettingData(wsmanMessageCreator)
	m.EthernetPortSettings = ethernetport.NewEthernetPortSettings(wsmanMessageCreator)
	m.GeneralSettings = general.NewGeneralSettings(wsmanMessageCreator)
	m.IEEE8021xCredentialContext = ieee8021x.NewIEEE8021xCredentialContext(wsmanMessageCreator)
	m.IEEE8021xProfile = ieee8021x.NewIEEE8021xProfile(wsmanMessageCreator)
	m.KerberosSettingData = kerberos.NewKerberosSettingData(wsmanMessageCreator)
	m.ManagementPresenceRemoteSAP = managementpresence.NewManagementPresenceRemoteSAP(wsmanMessageCreator)
	m.MessageLog = messagelog.NewMessageLog(wsmanMessageCreator)
	m.MPSUsernamePassword = mps.NewMPSUsernamePassword(wsmanMessageCreator)
	m.PublicKeyCertificate = publickey.NewPublicKeyCertificate(wsmanMessageCreator)
	m.PublicKeyManagementService = publickey.NewPublicKeyManagementService(wsmanMessageCreator)
	m.PublicPrivateKeyPair = publicprivate.NewPublicPrivateKeyPair(wsmanMessageCreator)
	m.RedirectionService = redirection.NewRedirectionService(wsmanMessageCreator)
	m.RemoteAccessPolicyAppliesToMPS = remoteaccess.NewRemoteAccessPolicyAppliesToMPS(wsmanMessageCreator)
	m.RemoteAccessPolicyRule = remoteaccess.NewRemoteAccessPolicyRule(wsmanMessageCreator)
	m.RemoteAccessService = remoteaccess.NewRemoteAccessService(wsmanMessageCreator)
	m.SetupAndConfigurationService = setupandconfiguration.NewSetupAndConfigurationService(wsmanMessageCreator)
	m.TimeSynchronizationService = timesynchronization.NewTimeSynchronizationService(wsmanMessageCreator)
	m.TLSCredentialContext = tls.NewTLSCredentialContext(wsmanMessageCreator)
	m.TLSSettingData = tls.NewTLSSettingData(wsmanMessageCreator)
	m.UserInitiatedConnectionService = userinitiatedconnection.NewUserInitiatedConnectionService(wsmanMessageCreator)
	m.WiFiPortConfigurationService = wifiportconfiguration.NewWiFiPortConfigurationService(wsmanMessageCreator)
	return m
}
