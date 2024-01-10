/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package credential

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewContext returns a new instance of the NewContext struct.
func NewContextWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Context {
	return Context{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_CredentialContext, client),
		client: client,
	}
}

// TODO: Figure out how to call GET requiring resourceURIs and Selectors

// Enumerate the instances of this class
func (context Context) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: context.base.Enumerate(),
		},
	}

	err = context.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return

}

// Pull instances of this class, following an Enumerate operation
func (context Context) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: context.base.Pull(enumerationContext),
		},
	}
	err = context.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
