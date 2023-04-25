package amt

import (
	"reflect"
	"testing"

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

func TestNewMessages(t *testing.T) {
	m := NewMessages()

	if m.wsmanMessageCreator == nil {
		t.Error("wsmanMessageCreator is not initialized")
	}
	if reflect.DeepEqual(m.AlarmClockService, alarmclock.Service{}) {
		t.Error("AlarmClockService is not initialized")
	}
	if reflect.DeepEqual(m.AuditLog, auditlog.AuditLog{}) {
		t.Error("AuditLog is not initialized")
	}
	if reflect.DeepEqual(m.AuthorizationService, authorization.AuthorizationService{}) {
		t.Error("AuthorizationService is not initialized")
	}
	if reflect.DeepEqual(m.BootCapabilities, boot.BootCapabilities{}) {
		t.Error("BootCapabilities is not initialized")
	}
	if reflect.DeepEqual(m.BootSettingData, boot.SettingData{}) {
		t.Error("BootSettingData is not initialized")
	}
	if reflect.DeepEqual(m.EnvironmentDetectionSettingData, environmentdetection.SettingData{}) {
		t.Error("EnvironmentDetectionSettingData is not initialized")
	}
	if reflect.DeepEqual(m.EthernetPortSettings, ethernetport.Settings{}) {
		t.Error("EthernetPortSettings is not initialized")
	}
	if reflect.DeepEqual(m.GeneralSettings, general.Settings{}) {
		t.Error("GeneralSettings is not initialized")
	}
	if reflect.DeepEqual(m.IEEE8021xCredentialContext, ieee8021x.CredentialContext{}) {
		t.Error("IEEE8021xCredentialContext is not initialized")
	}
	if reflect.DeepEqual(m.IEEE8021xProfile, ieee8021x.Profile{}) {
		t.Error("IEEE8021xProfile is not initialized")
	}
	if reflect.DeepEqual(m.KerberosSettingData, kerberos.KerberosSettingData{}) {
		t.Error("KerberosSettingData is not initialized")
	}
	if reflect.DeepEqual(m.ManagementPresenceRemoteSAP, managementpresence.RemoteSAP{}) {
		t.Error("ManagementPresenceRemoteSAP is not initialized")
	}
	if reflect.DeepEqual(m.MessageLog, messagelog.MessageLog{}) {
		t.Error("MessageLog is not initialized")
	}
	if reflect.DeepEqual(m.MPSUsernamePassword, mps.UsernamePassword{}) {
		t.Error("MPSUsernamePassword is not initialized")
	}
	if reflect.DeepEqual(m.PublicKeyCertificate, publickey.Certificate{}) {
		t.Error("PublicKeyCertificate is not initialized")
	}
	if reflect.DeepEqual(m.PublicKeyManagementService, publickey.ManagementService{}) {
		t.Error("PublicKeyManagementService is not initialized")
	}
	if reflect.DeepEqual(m.PublicPrivateKeyPair, publicprivate.KeyPair{}) {
		t.Error("PublicPrivateKeyPair is not initialized")
	}
	if reflect.DeepEqual(m.RedirectionService, redirection.Service{}) {
		t.Error("RedirectionService is not initialized")
	}
	if reflect.DeepEqual(m.RemoteAccessPolicyAppliesToMPS, remoteaccess.PolicyAppliesToMPS{}) {
		t.Error("RemoteAccessPolicyAppliesToMPS is not initialized")
	}
	if reflect.DeepEqual(m.RemoteAccessPolicyRule, remoteaccess.PolicyRule{}) {
		t.Error("RemoteAccessPolicyRule is not initialized")
	}
	if reflect.DeepEqual(m.RemoteAccessService, remoteaccess.Service{}) {
		t.Error("RemoteAccessService is not initialized")
	}
	if reflect.DeepEqual(m.SetupAndConfigurationService, setupandconfiguration.Service{}) {
		t.Error("SetupAndConfigurationService is not initialized")
	}
	if reflect.DeepEqual(m.TimeSynchronizationService, timesynchronization.Service{}) {
		t.Error("TimeSynchronizationService is not initialized")
	}
	if reflect.DeepEqual(m.TLSCredentialContext, tls.CredentialContext{}) {
		t.Error("TLSCredentialContext is not initialized")
	}
	if reflect.DeepEqual(m.TLSSettingData, tls.SettingData{}) {
		t.Error("TLSSettingData is not initialized")
	}
	if reflect.DeepEqual(m.UserInitiatedConnectionService, userinitiatedconnection.Service{}) {
		t.Error("UserInitiatedConnectionService is not initialized")
	}
	if reflect.DeepEqual(m.WiFiPortConfigurationService, wifiportconfiguration.Service{}) {
		t.Error("WiFiPortConfigurationService is not initialized")
	}
}
