package remoteaccess

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_RemoteAccessPolicyRule = "AMT_RemoteAccessPolicyRule"

type RemoteAccessPolicyRule struct {
	Trigger        Trigger
	TunnelLifeTime int
	ExtendedData   string
}

type Trigger uint8

const (
	TriggerUserInitiated    Trigger = 0
	TriggerAlert            Trigger = 1
	TriggerPeriodic         Trigger = 2
	TriggerHomeProvisioning Trigger = 3
)

type PolicyRule struct {
	base wsman.Base
}

func NewRemoteAccessPolicyRule(wsmanMessageCreator *wsman.WSManMessageCreator) PolicyRule {
	return PolicyRule{
		base: wsman.NewBase(wsmanMessageCreator, AMT_RemoteAccessPolicyRule),
	}
}
func (RemoteAccessPolicyRule PolicyRule) Get() string {
	return RemoteAccessPolicyRule.base.Get(nil)
}
func (RemoteAccessPolicyRule PolicyRule) Enumerate() string {
	return RemoteAccessPolicyRule.base.Enumerate()
}
func (RemoteAccessPolicyRule PolicyRule) Pull(enumerationContext string) string {
	return RemoteAccessPolicyRule.base.Pull(enumerationContext)
}
func (RemoteAccessPolicyRule PolicyRule) Put(remoteAccessPolicyRule RemoteAccessPolicyRule) string {
	return RemoteAccessPolicyRule.base.Put(remoteAccessPolicyRule, false, nil)
}
func (RemoteAccessPolicyRule PolicyRule) Delete(selector *wsman.Selector) string {
	return RemoteAccessPolicyRule.base.Delete(selector)
}
