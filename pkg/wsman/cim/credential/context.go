/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package credential facilitates communication with Intel® AMT devices in order to define a context (e.g., a System or Service) of a Credential.
//
// One example is a shared secret/ password which is defined within the context of an application (or Service).
//
// Generally, there is one scoping element for a Credential, however the multiplicities of the association allow a Credential to be scoped by more than one element.
//
// If this association is not instantiated for a Credential, that Credential is assumed to be scoped to the Namespace.
//
// This association may also be used to indicate that a Credential is valid in some other environment.
//
// For instance associating the Credential to a RemoteServiceAccessPoint would indicate that the Credential is used to access the remote service.
package credential

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
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
