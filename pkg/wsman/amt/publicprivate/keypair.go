/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package publicprivate facilitiates communication with Intel® AMT devices to manage a public-private key in the Intel® AMT CertStore.
//
// Instances of this class can be created using the AMT_PublicKeyManagementService.AddKey method. You can't delete a key instance if it is used by some service (TLS/EAC).
package publicprivate

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewPublicPrivateKeyPairWithClient instantiates a new KeyPair
func NewPublicPrivateKeyPairWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) KeyPair {
	return KeyPair{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_PublicPrivateKeyPair, client),
	}
}

// Get retrieves the representation of the instance
func (keyPair KeyPair) Get(handle int) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: fmt.Sprintf("Intel(r) AMT Key: Handle: %d", handle),
	}
	response = Response{
		Message: &client.Message{
			XMLInput: keyPair.base.Get((*message.Selector)(&selector)),
		},
	}
	// send the message to AMT
	err = keyPair.base.Execute(response.Message)
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
func (keyPair KeyPair) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: keyPair.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = keyPair.base.Execute(response.Message)
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
func (keyPair KeyPair) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: keyPair.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = keyPair.base.Execute(response.Message)
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

// Deletes an instance of a key pair
func (keyPair KeyPair) Delete(handle string) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: handle,
	}
	response = Response{
		Message: &client.Message{
			XMLInput: keyPair.base.Delete(selector),
		},
	}
	// send the message to AMT
	err = keyPair.base.Execute(response.Message)
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
