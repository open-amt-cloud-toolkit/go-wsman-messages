/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/common"
)

const AMT_PublicKeyManagementService = "AMT_PublicKeyManagementService"

type ( 
	Response struct {
		*wsman.Message 
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName							 xml.Name 					  `xml:"Body"`
		AddTrustedRootCertificate_OUTPUT AddTrustedCertificate_OUTPUT `xml:"AddTrustedRootCertificate_OUTPUT,omitempty"`
		AddTrustedCertificate_OUTPUT     AddTrustedCertificate_OUTPUT `xml:"AddCertificate_OUTPUT,omitempty"`
		AddKey_OUTPUT                    AddKey_OUTPUT                `xml:"AddKey_OUTPUT,omitempty"`
		KeyManagement 					 KeyManagement 				  `xml:"AMT__PublicKeyManagement"`

		EnumerateResponse common.EnumerateResponse
	}

	KeyManagement struct {
		CreationClassName		string 
		ElementName				string 
		EnabledDefault			int 
		EnabledState			int 
		Name					string 
		RequestedState			int 
		SystemCreationClassName	string 
		SystemName		 		string 
	}
)
type AddTrustedCertificate_OUTPUT struct {
	CreatedCertificate CreatedCertificate `xml:"CreatedCertificate"`
	ReturnValue        int
}

type CreatedCertificate struct {
	Address             string                            `xml:"Address,omitempty"`
	ReferenceParameters models.ReferenceParameters_OUTPUT `xml:"ReferenceParameters,omitempty"`
}
type AddCertificate_INPUT struct {
	XMLName         xml.Name `xml:"h:AddCertificate_INPUT"`
	H               string   `xml:"xmlns:h,attr"`
	CertificateBlob string   `xml:"h:CertificateBlob"`
}
type AddTrustedRootCertificate_INPUT struct {
	XMLName         xml.Name `xml:"h:AddTrustedRootCertificate_INPUT"`
	H               string   `xml:"xmlns:h,attr"`
	CertificateBlob string   `xml:"h:CertificateBlob"`
}
type AddKey_OUTPUT struct {
	CreatedKey  CreatedCertificate `xml:"CreatedKey"`
	ReturnValue int
}
type AddKey_INPUT struct {
	XMLName xml.Name `xml:"h:AddKey_INPUT"`
	H       string   `xml:"xmlns:h,attr"`
	KeyBlob []byte   `xml:"h:KeyBlob"`
}
type GenerateKeyPair_INPUT struct {
	XMLName      xml.Name     `xml:"h:GenerateKeyPair_INPUT"`
	H            string       `xml:"xmlns:h,attr"`
	KeyAlgorithm KeyAlgorithm `xml:"h:KeyAlgorithm"`
	KeyLength    KeyLength    `xml:"h:KeyLength"`
}

type KeyAlgorithm int

const (
	RSA KeyAlgorithm = 0
)

type KeyLength int

const (
	KeyLength2048 KeyLength = 2048
)

type PKCS10Request struct {
	XMLName                      xml.Name         `xml:"h:GeneratePKCS10RequestEx_INPUT"`
	H                            string           `xml:"xmlns:h,attr"`
	KeyPair                      string           `xml:"h:KeyPair"`
	NullSignedCertificateRequest string           `xml:"h:NullSignedCertificateRequest"`
	SigningAlgorithm             SigningAlgorithm `xml:"h:SigningAlgorithm"`
}
type SigningAlgorithm int

const (
	SHA1RSA SigningAlgorithm = iota
	SHA256RSA
)

type ManagementService struct {
	base message.Base
	client wsman.WSManClient 
}

func NewPublicKeyManagementServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client wsman.WSManClient) ManagementService {
	return ManagementService{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_PublicKeyManagementService, client),
		client: client, 
	}
}

func NewPublicKeyManagementService(wsmanMessageCreator *message.WSManMessageCreator) ManagementService {
	return ManagementService{
		base: message.NewBase(wsmanMessageCreator, AMT_PublicKeyManagementService),
	}
}

// Get retrieves the representation of the instance
func (PublicKeyManagementService ManagementService) Get() (response Response, err error) {

	response = Response{
		Message: &wsman.Message{
			XMLInput: PublicKeyManagementService.base.Get(nil),
		},
	}

	// send the message to AMT
	err = PublicKeyManagementService.base.Execute(response.Message)
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
func (PublicKeyManagementService ManagementService) Enumerate() (response Response, err error) {

	response = Response{
		Message: &wsman.Message{
			XMLInput: PublicKeyManagementService.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = PublicKeyManagementService.base.Execute(response.Message)
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
func (PublicKeyManagementService ManagementService) Pull(enumerationContext string) string {
	return PublicKeyManagementService.base.Pull(enumerationContext)
}

// Delete removes a the specified instance
func (PublicKeyManagementService ManagementService) Delete(instanceID string) string {
	selector := message.Selector{Name: "InstanceID", Value: instanceID}
	return PublicKeyManagementService.base.Delete(selector)
}
func (p ManagementService) AddCertificate(certificateBlob string) string {
	header := p.base.WSManMessageCreator.CreateHeader(string(actions.AddCertificate), AMT_PublicKeyManagementService, nil, "", "")
	certificate := AddCertificate_INPUT{CertificateBlob: certificateBlob}
	body := p.base.WSManMessageCreator.CreateBody("AddCertificate_INPUT", AMT_PublicKeyManagementService, &certificate)

	return p.base.WSManMessageCreator.CreateXML(header, body)
}

func (p ManagementService) AddTrustedRootCertificate(certificateBlob string) string {
	header := p.base.WSManMessageCreator.CreateHeader(string(actions.AddTrustedRootCertificate), AMT_PublicKeyManagementService, nil, "", "")
	trustedRootCert := AddTrustedRootCertificate_INPUT{CertificateBlob: certificateBlob}
	body := p.base.WSManMessageCreator.CreateBody("AddTrustedRootCertificate_INPUT", AMT_PublicKeyManagementService, &trustedRootCert)

	return p.base.WSManMessageCreator.CreateXML(header, body)
}

func (p ManagementService) GenerateKeyPair(keyPairParameters GenerateKeyPair_INPUT) string {
	header := p.base.WSManMessageCreator.CreateHeader(string(actions.GenerateKeyPair), AMT_PublicKeyManagementService, nil, "", "")
	body := p.base.WSManMessageCreator.CreateBody("GenerateKeyPair_INPUT", AMT_PublicKeyManagementService, &keyPairParameters)

	return p.base.WSManMessageCreator.CreateXML(header, body)
}

func (p ManagementService) GeneratePKCS10RequestEx(pkcs10Request PKCS10Request) string {
	header := p.base.WSManMessageCreator.CreateHeader(string(actions.GeneratePKCS10RequestEx), AMT_PublicKeyManagementService, nil, "", "")
	body := p.base.WSManMessageCreator.CreateBody("GeneratePKCS10RequestEx_INPUT", AMT_PublicKeyManagementService, &pkcs10Request)

	return p.base.WSManMessageCreator.CreateXML(header, body)
}

// AddKey adds a new certificate key to the Intel(R) AMT CertStore.
// A key cannot be removed if its corresponding certificate is referenced (for example, used by TLS, 802.1X, or EAC).
// After the method succeeds, a new instance of AMT_PublicPrivateKeyPair will be created.
// Possible return values are: PT_STATUS_SUCCESS(0), PT_STATUS_INTERNAL_ERROR(1), PT_STATUS_MAX_LIMIT_REACHED(23),
// PT_STATUS_FLASH_WRITE_LIMIT_EXCEEDED(38), PT_STATUS_DUPLICATE(2068), PT_STATUS_INVALID_KEY(2062).
func (p ManagementService) AddKey(keyBlob []byte) string {
	header := p.base.WSManMessageCreator.CreateHeader(string(actions.AddKey), AMT_PublicKeyManagementService, nil, "", "")
	params := &AddKey_INPUT{
		KeyBlob: keyBlob,
	}
	body := p.base.WSManMessageCreator.CreateBody("AddKey_INPUT", AMT_PublicKeyManagementService, params)

	return p.base.WSManMessageCreator.CreateXML(header, body)
}
