/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publicprivate

import (
	"encoding/xml"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_PublicPrivateKeyPair = "AMT_PublicPrivateKeyPair"

type PullResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  wsman.Header
	Body    PullResponseBody
}

type PullResponseBody struct {
	PullResponse PullResponse
}

type PullResponse struct {
	Items         []PublicPrivateKeyPair `xml:"Items>AMT_PublicPrivateKeyPair"`
	EndOfSequence string
}

type PublicPrivateKeyPair struct {
	ElementName string // A user-friendly name for the object . . .
	InstanceID  string // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
	//DERKey      [1210]uint8 // RSA Key encoded as DES PKCS#1.
	DERKey string
}

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
