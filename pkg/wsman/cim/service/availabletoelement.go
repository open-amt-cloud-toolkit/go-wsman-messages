/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package service

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewServiceAvailableToElement returns a new instance of the ServiceAvailableToElement struct.
func NewServiceAvailableToElementWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) AvailableToElement {
	return AvailableToElement{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_ServiceAvailableToElement, client),
		client: client,
	}
}

// TODO Figure out how to call GET requiring resourceURIs and Selectors
// Get retrieves the representation of the instance.  No route

// Enumerate returns an enumeration context which is used in a subsequent Pull call
func (availableToElement AvailableToElement) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: availableToElement.base.Enumerate(),
		},
	}

	err = availableToElement.base.Execute(response.Message)
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
func (availableToElement AvailableToElement) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: availableToElement.base.Pull(enumerationContext),
		},
	}
	err = availableToElement.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
