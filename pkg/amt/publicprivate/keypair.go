/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publicprivate

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_PublicPrivateKeyPair = "AMT_PublicPrivateKeyPair"

type KeyPair struct {
	base wsman.Base
}

func NewPublicPrivateKeyPair(wsmanMessageCreator *wsman.WSManMessageCreator) KeyPair {
	return KeyPair{
		base: wsman.NewBase(wsmanMessageCreator, AMT_PublicPrivateKeyPair),
	}
}

// Get retrieves the representation of the instance
func (PublicPrivateKeyPair KeyPair) Get() string {
	return PublicPrivateKeyPair.base.Get(nil)
}

// Enumerates the instances of this class
func (PublicPrivateKeyPair KeyPair) Enumerate() string {
	return PublicPrivateKeyPair.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (PublicPrivateKeyPair KeyPair) Pull(enumerationContext string) string {
	return PublicPrivateKeyPair.base.Pull(enumerationContext)
}
func (PublicPrivateKeyPair KeyPair) Delete(instanceID string) string {
	selector := wsman.Selector{Name: "InstanceID", Value: instanceID}
	return PublicPrivateKeyPair.base.Delete(selector)
}
