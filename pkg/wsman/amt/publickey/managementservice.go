/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// NewPublicKeyManagementServiceWithClient instantiates a new ManagementService
func NewPublicKeyManagementServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) ManagementService {
	return ManagementService{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_PublicKeyManagementService, client),
	}
}

// Get retrieves the representation of the instance
func (managementService ManagementService) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: managementService.base.Get(nil),
		},
	}
	// send the message to AMT
	err = managementService.base.Execute(response.Message)
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

// Enumerate returns an enumeration context which is used in a subsequent Pull call
func (managementService ManagementService) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: managementService.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = managementService.base.Execute(response.Message)
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

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (managementService ManagementService) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: managementService.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = managementService.base.Execute(response.Message)
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

// Delete removes a the specified instance
func (managementService ManagementService) Delete(instanceID string) (response Response, err error) {
	selector := message.Selector{Name: "InstanceID", Value: instanceID}
	response = Response{
		Message: &client.Message{
			XMLInput: managementService.base.Delete(selector),
		},
	}
	// send the message to AMT
	err = managementService.base.Execute(response.Message)
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

//unittest
// func TestCheckReturnValue(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		in   error
// 		item string
// 		want error
// 	}{
// 		{"TestNoError", 0, "item", nil},
// 		{"TestAlreadyExists", common.PT_STATUS_DUPLICATE, "item", utils.AmtPtStatusCodeBase + common.PT_STATUS_DUPLICATE},
// 		{"TestInvalidItem", common.PT_STATUS_INVALID_CERT, "item", utils.AmtPtStatusCodeBase + common.PT_STATUS_INVALID_CERT},
// 		{"TestNonZeroReturnCode", 2082, "item", utils.AmtPtStatusCodeBase + 2082},
// 	}

//		for _, tt := range tests {
//			t.Run(tt.name, func(t *testing.T) {
//				got := checkReturnValue(tt.in, tt.item)
//				assert.Equal(t, tt.want, got)
//			})
//		}
//	}
func checkReturnValue(rc int, item string) (err error) {
	if rc != common.PT_STATUS_SUCCESS {
		if rc == common.PT_STATUS_DUPLICATE {
			return errors.New(item + " already exists and must be removed before continuing")
		} else if rc == common.PT_STATUS_INVALID_CERT {
			return errors.New(item + " is invalid")
		} else {
			return errors.New(item + " non-zero return code: " + strconv.Itoa(rc))
		}
	}
	return nil
}

// This function adds new certificate to the Intel® AMT CertStore. A certificate cannot be removed if it is referenced (for example, used by TLS, 802.1X or EAC).
func (managementService ManagementService) AddCertificate(certificateBlob string) (response Response, err error) {
	header := managementService.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_PublicKeyManagementService, AddCertificate), AMT_PublicKeyManagementService, nil, "", "")
	certificate := AddCertificate_INPUT{
		H:               fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicKeyManagementService),
		CertificateBlob: certificateBlob,
	}
	body := managementService.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(AddCertificate), AMT_PublicKeyManagementService, &certificate)
	response = Response{
		Message: &client.Message{
			XMLInput: managementService.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = managementService.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	err = checkReturnValue(response.Body.AddCertificate_OUTPUT.ReturnValue, "Client Certificate")
	return
}

// This function adds new root certificate to the Intel® AMT CertStore. A certificate cannot be removed if it is referenced (for example, used by TLS, 802.1X or EAC).
func (managementService ManagementService) AddTrustedRootCertificate(certificateBlob string) (response Response, err error) {
	header := managementService.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_PublicKeyManagementService, AddTrustedRootCertificate), AMT_PublicKeyManagementService, nil, "", "")
	trustedRootCert := AddTrustedRootCertificate_INPUT{
		H:               fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicKeyManagementService),
		CertificateBlob: certificateBlob,
	}
	body := managementService.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(AddTrustedRootCertificate), AMT_PublicKeyManagementService, &trustedRootCert)

	response = Response{
		Message: &client.Message{
			XMLInput: managementService.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = managementService.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	err = checkReturnValue(response.Body.AddTrustedRootCertificate_OUTPUT.ReturnValue, "Root Certificate")

	return
}

// This API is used to generate a key in the FW
func (managementService ManagementService) GenerateKeyPair(keyAlgorithm KeyAlgorithm, keyLength KeyLength) (response Response, err error) {
	header := managementService.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_PublicKeyManagementService, GenerateKeyPair), AMT_PublicKeyManagementService, nil, "", "")
	generateKeyPair := GenerateKeyPair_INPUT{
		H:            fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicKeyManagementService),
		KeyAlgorithm: keyAlgorithm,
		KeyLength:    keyLength,
	}
	body := managementService.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GenerateKeyPair), AMT_PublicKeyManagementService, &generateKeyPair)
	response = Response{
		Message: &client.Message{
			XMLInput: managementService.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = managementService.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	err = checkReturnValue(response.Body.AddKey_OUTPUT.ReturnValue, "Private Key")

	return
}

// This API is used to create a PKCS#10 certificate signing request based on a key from the key store.
func (managementService ManagementService) GeneratePKCS10RequestEx(keyPair, nullSignedCertificateRequest string, signingAlgorithm SigningAlgorithm) (response Response, err error) {
	header := managementService.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_PublicKeyManagementService, GeneratePKCS10RequestEx), AMT_PublicKeyManagementService, nil, "", "")
	pkcs10Request := PKCS10Request{
		H: fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicKeyManagementService),
		KeyPair: KeyPair{
			Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
			ReferenceParameters: ReferenceParametersRequest{
				ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicPrivateKeyPair",
				SelectorSet: SelectorSetRequest{
					Selectors: []SelectorRequest{
						{
							Name: "InstanceID",
							Text: keyPair,
						},
					},
				},
			},
		},
		SigningAlgorithm:             signingAlgorithm,
		NullSignedCertificateRequest: nullSignedCertificateRequest,
	}
	body := managementService.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GeneratePKCS10RequestEx), AMT_PublicKeyManagementService, &pkcs10Request)
	response = Response{
		Message: &client.Message{
			XMLInput: managementService.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = managementService.base.Execute(response.Message)
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

// This function adds new certificate key to the Intel® AMT CertStore. A key cannot be removed if its corresponding certificate is referenced (for example, used by TLS, 802.1X or EAC).
// After the method succeeds, a new instance of AMT_PublicPrivateKeyPair will be created.
// Possible return values are: PT_STATUS_SUCCESS(0), PT_STATUS_INTERNAL_ERROR(1), PT_STATUS_MAX_LIMIT_REACHED(23),
// PT_STATUS_FLASH_WRITE_LIMIT_EXCEEDED(38), PT_STATUS_DUPLICATE(2068), PT_STATUS_INVALID_KEY(2062).
func (managementService ManagementService) AddKey(keyBlob string) (response Response, err error) {
	header := managementService.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_PublicKeyManagementService, AddKey), AMT_PublicKeyManagementService, nil, "", "")
	params := &AddKey_INPUT{
		H:       fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicKeyManagementService),
		KeyBlob: keyBlob,
	}
	body := managementService.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(AddKey), AMT_PublicKeyManagementService, params)
	response = Response{
		Message: &client.Message{
			XMLInput: managementService.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = managementService.base.Execute(response.Message)
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
