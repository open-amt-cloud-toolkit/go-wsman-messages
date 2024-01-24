/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewTLSProtocolEndpointCollectionWithClient instantiates a new ProtocolEndpointCollection
func NewTLSProtocolEndpointCollectionWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) ProtocolEndpointCollection {
	return ProtocolEndpointCollection{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_TLSProtocolEndpointCollection, client),
	}
}

// Get retrieves the representation of the instance
func (collection ProtocolEndpointCollection) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: collection.base.Get(nil),
		},
	}
	// send the message to AMT
	err = collection.base.Execute(response.Message)
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
func (collection ProtocolEndpointCollection) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: collection.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = collection.base.Execute(response.Message)
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
func (collection ProtocolEndpointCollection) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: collection.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = collection.base.Execute(response.Message)
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
