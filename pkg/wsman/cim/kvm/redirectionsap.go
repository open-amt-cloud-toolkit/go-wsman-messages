/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kvm

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/actions"
)

const CIM_KVMRedirectionSAP = "CIM_KVMRedirectionSAP"

type RedirectionSAP struct {
	base message.Base
}

// NewKVMRedirectionSAP returns a new instance of the KVMRedirectionSAP struct.
func NewKVMRedirectionSAP(wsmanMessageCreator *message.WSManMessageCreator) RedirectionSAP {
	return RedirectionSAP{
		base: message.NewBase(wsmanMessageCreator, string(CIM_KVMRedirectionSAP)),
	}
}

// RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
func (k RedirectionSAP) RequestStateChange(requestedState int) string {
	return k.base.RequestStateChange(actions.RequestStateChange(string(CIM_KVMRedirectionSAP)), requestedState)
}

// Get retrieves the representation of the instance
func (b RedirectionSAP) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b RedirectionSAP) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b RedirectionSAP) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
