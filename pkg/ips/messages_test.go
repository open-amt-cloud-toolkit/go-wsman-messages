package ips

import (
	"reflect"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/ips/alarmclock"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/ips/hostbasedsetup"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/ips/ieee8021x"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/ips/optin"
)

func TestNewMessages(t *testing.T) {
	m := NewMessages()

	if m.wsmanMessageCreator == nil {
		t.Error("wsmanMessageCreator is not initialized")
	}
	if reflect.DeepEqual(m.OptInService, optin.Service{}) {
		t.Error("AlarmClockService is not initialized")
	}
	if reflect.DeepEqual(m.HostBasedSetupService, hostbasedsetup.Service{}) {
		t.Error("AuditLog is not initialized")
	}
	if reflect.DeepEqual(m.AlarmClockOccurrence, alarmclock.Occurrence{}) {
		t.Error("AuthorizationService is not initialized")
	}
	if reflect.DeepEqual(m.IEEE8021xCredentialContext, ieee8021x.CredentialContext{}) {
		t.Error("BootCapabilities is not initialized")
	}
	if reflect.DeepEqual(m.IEEE8021xSettings, ieee8021x.Settings{}) {
		t.Error("BootSettingData is not initialized")
	}
}