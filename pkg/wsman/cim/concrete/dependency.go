/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package concrete

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewDependency returns a new instance of the NewDependency struct.
// should be NewDependency() because concrete is scoped already as package name.
func NewDependencyWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Dependency {
	return Dependency{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_ConcreteDependency, client),
		client: client,
	}
}

// TODO: Figure out how to call GET requiring resourceURIs and Selectors

// Enumerate the instances of this class
func (dependency Dependency) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: dependency.base.Enumerate(),
		},
	}

	err = dependency.base.Execute(response.Message)
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
func (dependency Dependency) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: dependency.base.Pull(enumerationContext),
		},
	}
	err = dependency.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
