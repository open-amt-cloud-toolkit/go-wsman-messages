/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

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
func (PublicKeyCertificate Certificate) Get() (response Response, err error) {

	response = Response{
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
func (PublicKeyCertificate Certificate) Enumerate() (response Response, err error) {

	response = Response{
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
func (PublicKeyCertificate Certificate) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: PublicKeyCertificate.base.Pull(enumerationContext),
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

// Put will change properties of the selected instance
// func (PublicKeyCertificate Certificate) Put(publicKeyCertificate PublicKeyCertificate) string {
// 	return PublicKeyCertificate.base.Put(publicKeyCertificate, false, nil)
// }

// Delete removes a the specified instance
// func (PublicKeyCertificate Certificate) Delete(instanceID string) string {
// 	selector := message.Selector{Name: "InstanceID", Value: instanceID}
// 	return PublicKeyCertificate.base.Delete(selector)
// }
