/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_PublicKeyCertificate = "AMT_PublicKeyCertificate"

type PublicKeyCertificate struct {
	ElementName           string      // A user-friendly name for the object . . .
	InstanceID            string      // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
	X509Certificate       [4100]uint8 // uint8[4100] // The X.509 Certificate blob.
	TrustedRootCertficate bool        // For root certificate [that were added by AMT_PublicKeyManagementService.AddTrustedRootCertificate()]this property will be true.
	Issuer                string      // The Issuer field of this certificate.
	Subject               string      // The Subject field of this certificate.
	ReadOnlyCertificate   bool        // Indicates whether the certificate is an Intel AMT self-signed certificate. If True, the certificate cannot be deleted.
}
type Certificate struct {
	base wsman.Base
}

func NewPublicKeyCertificate(wsmanMessageCreator *wsman.WSManMessageCreator) Certificate {
	return Certificate{
		base: wsman.NewBase(wsmanMessageCreator, AMT_PublicKeyCertificate),
	}
}

// Get retrieves the representation of the instance
func (PublicKeyCertificate Certificate) Get() string {
	return PublicKeyCertificate.base.Get(nil)
}

// Enumerates the instances of this class
func (PublicKeyCertificate Certificate) Enumerate() string {
	return PublicKeyCertificate.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (PublicKeyCertificate Certificate) Pull(enumerationContext string) string {
	return PublicKeyCertificate.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (PublicKeyCertificate Certificate) Put(publicKeyCertificate PublicKeyCertificate) string {
	return PublicKeyCertificate.base.Put(publicKeyCertificate, false, nil)
}

// Delete removes a the specified instance
func (PublicKeyCertificate Certificate) Delete(instanceID string) string {
	selector := wsman.Selector{Name: "InstanceID", Value: instanceID}
	return PublicKeyCertificate.base.Delete(selector)
}
