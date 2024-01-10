/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package card

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewCard returns a new instance of the Card struct.
func NewCardWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_Card, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (card Package) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: card.base.Get(nil),
		},
	}

	err = card.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return

}

// Enumerates the instances of this class
func (card Package) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: card.base.Enumerate(),
		},
	}

	err = card.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return

}

// Pulls instances of this class, following an Enumerate operation
func (card Package) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: card.base.Pull(enumerationContext),
		},
	}
	err = card.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
