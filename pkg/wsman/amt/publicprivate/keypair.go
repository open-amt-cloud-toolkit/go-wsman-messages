/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publicprivate

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

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

// Enumerates the instances of this class
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

// Pulls instances of this class, following an Enumerate operation
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
func (keyPair KeyPair) Delete(handle int) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: fmt.Sprintf("Intel(r) AMT Key: Handle: %d", handle),
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
