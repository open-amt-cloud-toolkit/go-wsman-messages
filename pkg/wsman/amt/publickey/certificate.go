/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package publickey facilitiates communication with Intel® AMT devices to access and configure Public Key Certificates and Public Key Management Service classes for AMT
//
// Certificate:
// This class represents a X.509 Certificate in the Intel® AMT CertStore. Instances of this class can be created using the AMT_PublicKeyManagementService.AddCertificate and AMT_PublicKeyManagementService.AddTrustedRootCertificate methods. A certificate cannot be deleted while it is being used by any service (TLS/EAC).
//
// Management Service:
// This service contains the information necessary to represent and manage the functionality provided by the Intel® AMT CertStore.
package publickey

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewPublicKeyCertificateWithClient instantiates a new Certificate
func NewPublicKeyCertificateWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Certificate {
	return Certificate{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_PublicKeyCertificate, client),
	}
}

// Get retrieves the representation of the instance
func (certificate Certificate) Get(handle int) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: fmt.Sprintf("Intel(r) AMT Certificate: Handle: %d", handle),
	}
	response = Response{
		Message: &client.Message{
			XMLInput: certificate.base.Get(&selector),
		},
	}
	// send the message to AMT
	err = certificate.base.Execute(response.Message)
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
func (certificate Certificate) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: certificate.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = certificate.base.Execute(response.Message)
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
func (certificate Certificate) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: certificate.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = certificate.base.Execute(response.Message)
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

// Put will change properties of the selected instance
func (certificate Certificate) Put(handle int, cert string) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: fmt.Sprintf("Intel(r) AMT Certificate: Handle: %d", handle),
	}
	publicKeyCertificate := PublicKeyCertificateRequest{}
	publicKeyCertificate.X509Certificate = cert
	publicKeyCertificate.H = fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicKeyCertificate)
	response = Response{
		Message: &client.Message{
			XMLInput: certificate.base.Put(publicKeyCertificate, true, &selector),
		},
	}
	// send the message to AMT
	err = certificate.base.Execute(response.Message)
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
func (certificate Certificate) Delete(instanceID string) (response Response, err error) {
	selector := message.Selector{Name: "InstanceID", Value: instanceID}
	response = Response{
		Message: &client.Message{
			XMLInput: certificate.base.Delete(selector),
		},
	}
	// send the message to AMT
	err = certificate.base.Execute(response.Message)
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
