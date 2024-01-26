/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package software

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewSoftwareIdentity returns a new instance of the SoftwareIdentity struct.
func NewSoftwareIdentityWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Identity {
	return Identity{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_SoftwareIdentity, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (identity Identity) Get(selector Selector) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: identity.base.Get((*message.Selector)(&selector)),
		},
	}

	err = identity.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return

}

// Enumerate returns an enumeration context which is used in a subsequent Pull call
func (identity Identity) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: identity.base.Enumerate(),
		},
	}

	err = identity.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return

}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (identity Identity) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: identity.base.Pull(enumerationContext),
		},
	}
	err = identity.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
