/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

const AMT_PublicKeyCertificate = "AMT_PublicKeyCertificate"

type (
	ResponseCert struct {
		*client.Message
		XMLName  xml.Name       `xml:"Envelope"`
		Header   message.Header `xml:"Header"`
		BodyCert BodyCert       `xml:"Body"`
	}
	BodyCert struct {
		XMLName xml.Name `xml:"Body"`
		KeyCert KeyCert  `xml:"AMT_PublicKeyCertificate"`

		EnumerateResponse common.EnumerateResponse
	}
	KeyCert struct {
		ElementName           string
		InstanceID            string
		X509Certificate       string
		TrustedRootCertficate bool
		Issuer                string
		Subject               string
	}
)
type PullResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  message.Header
	Body    PullResponseBody
}

type PullResponseBody struct {
	PullResponse PullResponse
}

type PullResponse struct {
	Items         []PublicKeyCertificate `xml:"Items>AMT_PublicKeyCertificate"`
	EndOfSequence string
}

type PublicKeyCertificate struct {
	ElementName           string // A user-friendly name for the object . . .
	InstanceID            string // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
	X509Certificate       string // uint8[4100] // The X.509 Certificate blob.
	TrustedRootCertficate bool   // For root certificate [that were added by AMT_PublicKeyManagementService.AddTrustedRootCertificate()]this property will be true.
	Issuer                string // The Issuer field of this certificate.
	Subject               string // The Subject field of this certificate.
	ReadOnlyCertificate   bool   // Indicates whether the certificate is an Intel AMT self-signed certificate. If True, the certificate cannot be deleted.
}

type Certificate struct {
	base   message.Base
	client client.WSMan
}

func NewPublicKeyCertificateWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Certificate {
	return Certificate{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_PublicKeyCertificate, client),
		client: client,
	}
}

func NewPublicKeyCertificate(wsmanMessageCreator *message.WSManMessageCreator) Certificate {
	return Certificate{
		base: message.NewBase(wsmanMessageCreator, AMT_PublicKeyCertificate),
	}
}

// Get retrieves the representation of the instance
func (PublicKeyCertificate Certificate) Get() (response ResponseCert, err error) {

	response = ResponseCert{
		Message: &client.Message{
			XMLInput: PublicKeyCertificate.base.Get(nil),
		},
	}

	// send the message to AMT
	err = PublicKeyCertificate.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerates the instances of this class
func (PublicKeyCertificate Certificate) Enumerate() (response ResponseCert, err error) {

	response = ResponseCert{
		Message: &client.Message{
			XMLInput: PublicKeyCertificate.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = PublicKeyCertificate.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pulls instances of this class, following an Enumerate operation
// func (PublicKeyCertificate Certificate) Pull(enumerationContext string) string {
// 	return PublicKeyCertificate.base.Pull(enumerationContext)
// }

// Put will change properties of the selected instance
// func (PublicKeyCertificate Certificate) Put(publicKeyCertificate PublicKeyCertificate) string {
// 	return PublicKeyCertificate.base.Put(publicKeyCertificate, false, nil)
// }

// Delete removes a the specified instance
// func (PublicKeyCertificate Certificate) Delete(instanceID string) string {
// 	selector := message.Selector{Name: "InstanceID", Value: instanceID}
// 	return PublicKeyCertificate.base.Delete(selector)
// }
