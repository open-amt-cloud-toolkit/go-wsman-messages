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

type ManagementService struct {
	base message.Base
}

// OUTPUTS
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                               xml.Name                         `xml:"Body"`
		AddTrustedRootCertificate_OUTPUT      AddTrustedRootCertificate_OUTPUT `xml:"AddTrustedRootCertificate_OUTPUT,omitempty"`
		AddCertificate_OUTPUT                 AddCertificate_OUTPUT            `xml:"AddCertificate_OUTPUT,omitempty"`
		AddKey_OUTPUT                         AddKey_OUTPUT                    `xml:"AddKey_OUTPUT,omitempty"`
		GenerateKeyPair_OUTPUT                GenerateKeyPair_OUTPUT           `xml:"GenerateKeyPair_OUTPUT,omitempty"`
		KeyManagementGetResponse              KeyManagementResponse            `xml:"AMT_PublicKeyManagementService,omitempty"`
		PublicKeyCertificateGetAndPutResponse PublicKeyCertificateResponse     `xml:"AMT_PublicKeyCertificate,omitempty"`
		EnumerateResponse                     common.EnumerateResponse
		PullResponse                          PullResponse
	}
	PullResponse struct {
		XMLName                   xml.Name                       `xml:"PullResponse,omitempty"`
		KeyManagementItems        []KeyManagementResponse        `xml:"Items>AMT_PublicKeyManagementService,omitempty"`
		PublicKeyCertificateItems []PublicKeyCertificateResponse `xml:"Items>AMT_PublicKeyCertificate,omitempty"`
	}
	KeyManagementResponse struct {
		XMLName                 xml.Name       `xml:"AMT_PublicKeyManagementService,omitempty"`
		CreationClassName       string         `xml:"CreationClassName,omitempty"`
		ElementName             string         `xml:"ElementName,omitempty"`
		EnabledDefault          EnabledDefault `xml:"EnabledDefault"`
		EnabledState            EnabledState   `xml:"EnabledState"`
		Name                    string         `xml:"Name,omitempty"`
		RequestedState          RequestedState `xml:"RequestedState"`
		SystemCreationClassName string         `xml:"SystemCreationClassName,omitempty"`
		SystemName              string         `xml:"SystemName,omitempty"`
	}
	PublicKeyCertificateResponse struct {
		XMLName                xml.Name `xml:"AMT_PublicKeyCertificate,omitempty"`
		ElementName            string   `xml:"ElementName,omitempty"`     // A user-friendly name for the object . . .
		InstanceID             string   `xml:"InstanceID,omitempty"`      // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		X509Certificate        string   `xml:"X509Certificate,omitempty"` // uint8[4100] // The X.509 Certificate blob.
		TrustedRootCertificate bool     `xml:"TrustedRootCertficate"`     // For root certificate [that were added by AMT_PublicKeyManagementService.AddTrustedRootCertificate()]this property will be true. FYI Certificate is spelled wrong comimg from AMT.
		Issuer                 string   `xml:"Issuer,omitempty"`          // The Issuer field of this certificate.
		Subject                string   `xml:"Subject,omitempty"`         // The Subject field of this certificate.
		ReadOnlyCertificate    bool     `xml:"ReadOnlyCertificate"`       // Indicates whether the certificate is an Intel AMT self-signed certificate. If True, the certificate cannot be deleted.
	}
	AddTrustedRootCertificate_OUTPUT struct {
		XMLName            xml.Name                   `xml:"AddTrustedRootCertificate_OUTPUT"`
		CreatedCertificate CreatedCertificateResponse `xml:"CreatedCertificate,omitempty"`
		ReturnValue        int                        `xml:"ReturnValue,omitempty"`
	}
	AddCertificate_OUTPUT struct {
		XMLName            xml.Name                   `xml:"AddCertificate_OUTPUT"`
		CreatedCertificate CreatedCertificateResponse `xml:"CreatedCertificate,omitempty"`
		ReturnValue        int                        `xml:"ReturnValue,omitempty"`
	}
	AddKey_OUTPUT struct {
		XMLName     xml.Name           `xml:"AddKey_OUTPUT,omitempty"`
		CreatedKey  CreatedKeyResponse `xml:"CreatedKey,omitempty"`
		ReturnValue int                `xml:"ReturnValue,omitempty"`
	}
	GenerateKeyPair_OUTPUT struct {
		XMLName     xml.Name        `xml:"GenerateKeyPair_OUTPUT,omitempty"`
		KeyPair     KeyPairResponse `xml:"KeyPair,omitempty"`
		ReturnValue int             `xml:"ReturnValue,omitempty"`
	}
	KeyPairResponse struct {
		XMLName             xml.Name                    `xml:"KeyPair,omitempty"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	CreatedKeyResponse struct {
		XMLName             xml.Name                    `xml:"CreatedKey,omitempty"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	CreatedCertificateResponse struct {
		XMLName             xml.Name                    `xml:"CreatedCertificate,omitempty"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	ReferenceParametersResponse struct {
		XMLName     xml.Name `xml:"ReferenceParameters,omitempty"`
		ResourceURI string
		SelectorSet SelectorSetResponse `xml:"SelectorSet,omitempty"`
	}
	SelectorSetResponse struct {
		XMLName   xml.Name           `xml:"SelectorSet,omitempty"`
		Selectors []SelectorResponse `xml:"Selector"`
	}
	SelectorResponse struct {
		XMLName xml.Name `xml:"Selector,omitempty"`
		Name    string   `xml:"Name,attr"`
		Text    string   `xml:",chardata"`
	}

	EnabledDefault int
	EnabledState   int
	RequestedState int
)

// INPUTS
// Request Types
type (
	PublicKeyCertificateRequest struct {
		XMLName                xml.Name `xml:"h:AMT_PublicKeyCertificate"`
		H                      string   `xml:"xmlns:h,attr"`
		ElementName            string   `xml:"h:ElementName"`            // A user-friendly name for the object . . .
		InstanceID             string   `xml:"h:InstanceID"`             // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		X509Certificate        string   `xml:"h:X509Certificate"`        // uint8[4100] // The X.509 Certificate blob.
		TrustedRootCertificate bool     `xml:"h:TrustedRootCertificate"` // For root certificate [that were added by AMT_PublicKeyManagementService.AddTrustedRootCertificate()]this property will be true.
		Issuer                 string   `xml:"h:Issuer"`                 // The Issuer field of this certificate.
		Subject                string   `xml:"h:Subject"`                // The Subject field of this certificate.
		ReadOnlyCertificate    bool     `xml:"h:ReadOnlyCertificate"`    // Indicates whether the certificate is an Intel AMT self-signed certificate. If True, the certificate cannot be deleted.
	}
	AddCertificate_INPUT struct {
		XMLName         xml.Name `xml:"h:AddCertificate_INPUT"`
		H               string   `xml:"xmlns:h,attr"`
		CertificateBlob string   `xml:"h:CertificateBlob"`
	}
	AddTrustedRootCertificate_INPUT struct {
		XMLName         xml.Name `xml:"h:AddTrustedRootCertificate_INPUT"`
		H               string   `xml:"xmlns:h,attr"`
		CertificateBlob string   `xml:"h:CertificateBlob"`
	}
	AddKey_INPUT struct {
		XMLName xml.Name `xml:"h:AddKey_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		KeyBlob string   `xml:"h:KeyBlob"`
	}
	GenerateKeyPair_INPUT struct {
		XMLName      xml.Name     `xml:"h:GenerateKeyPair_INPUT"`
		H            string       `xml:"xmlns:h,attr"`
		KeyAlgorithm KeyAlgorithm `xml:"h:KeyAlgorithm"`
		KeyLength    KeyLength    `xml:"h:KeyLength"`
	}
	PKCS10Request struct {
		XMLName                      xml.Name         `xml:"h:GeneratePKCS10RequestEx_INPUT"`
		H                            string           `xml:"xmlns:h,attr"`
		KeyPair                      string           `xml:"h:KeyPair"`
		NullSignedCertificateRequest string           `xml:"h:NullSignedCertificateRequest"`
		SigningAlgorithm             SigningAlgorithm `xml:"h:SigningAlgorithm"`
	}
	SigningAlgorithm int
	KeyAlgorithm     int
	KeyLength        int
)
