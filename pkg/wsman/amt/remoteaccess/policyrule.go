/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
)

const AMT_RemoteAccessPolicyRule = "AMT_RemoteAccessPolicyRule"

type RemoteAccessPolicyRule struct {
	Trigger        Trigger
	TunnelLifeTime int
	ExtendedData   string
}

type Trigger uint8

const (
	UserInitiated Trigger = iota
	Alert
	Periodic
	HomeProvisioning
)

type PolicyRule struct {
	base message.Base
}

func NewRemoteAccessPolicyRule(wsmanMessageCreator *message.WSManMessageCreator) PolicyRule {
	return PolicyRule{
		base: message.NewBase(wsmanMessageCreator, AMT_RemoteAccessPolicyRule),
	}
}

// Get retrieves the representation of the instance
func (RemoteAccessPolicyRule PolicyRule) Get() string {
	return RemoteAccessPolicyRule.base.Get(nil)
}

// Enumerates the instances of this class
func (RemoteAccessPolicyRule PolicyRule) Enumerate() string {
	return RemoteAccessPolicyRule.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (RemoteAccessPolicyRule PolicyRule) Pull(enumerationContext string) string {
	return RemoteAccessPolicyRule.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (RemoteAccessPolicyRule PolicyRule) Put(remoteAccessPolicyRule RemoteAccessPolicyRule) string {
	return RemoteAccessPolicyRule.base.Put(remoteAccessPolicyRule, false, nil)
}

// Delete removes a the specified instance
func (RemoteAccessPolicyRule PolicyRule) Delete(handle string) string {
	selector := message.Selector{Name: "PolicyRuleName", Value: handle}
	return RemoteAccessPolicyRule.base.Delete(selector)
}
