/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

const AMT_RemoteAccessPolicyAppliesToMPS = "AMT_RemoteAccessPolicyAppliesToMPS"

type RemoteAccessPolicyAppliesToMPS struct {
	models.PolicySetAppliesToElement
	OrderOfAccess int
	MpsType       MpsType
}

type MpsType int

const (
	ExternalMPS MpsType = iota
	InternalMPS
	BothMPS
)

type PolicyAppliesToMPS struct {
	base wsman.Base
}

func NewRemoteAccessPolicyAppliesToMPS(wsmanMessageCreator *wsman.WSManMessageCreator) PolicyAppliesToMPS {
	return PolicyAppliesToMPS{
		base: wsman.NewBase(wsmanMessageCreator, AMT_RemoteAccessPolicyAppliesToMPS),
	}
}
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Get() string {
	return RemoteAccessPolicyAppliesToMPS.base.Get(nil)
}
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Enumerate() string {
	return RemoteAccessPolicyAppliesToMPS.base.Enumerate()
}
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Pull(enumerationContext string) string {
	return RemoteAccessPolicyAppliesToMPS.base.Pull(enumerationContext)
}
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Put(remoteAccessPolicyAppliesToMPS RemoteAccessPolicyAppliesToMPS) string {
	return RemoteAccessPolicyAppliesToMPS.base.Put(remoteAccessPolicyAppliesToMPS, false, nil)
}
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Delete(selector *wsman.Selector) string {
	return RemoteAccessPolicyAppliesToMPS.base.Delete(selector)
}
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Create(remoteAccessPolicyAppliesToMPS RemoteAccessPolicyAppliesToMPS) string {
	return RemoteAccessPolicyAppliesToMPS.base.Create(remoteAccessPolicyAppliesToMPS, nil)
}
