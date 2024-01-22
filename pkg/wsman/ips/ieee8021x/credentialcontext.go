/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewIEEE8021xCredentialContext returns a new instance of the IPS_8021xCredentialContext struct.
func NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) CredentialContext {
	return CredentialContext{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPS_8021xCredentialContext, client),
	}
}

// Get retrieves the representation of the instance
func (credentialContext CredentialContext) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Get(nil),
		},
	}
	err = credentialContext.base.Execute(response.Message)
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
func (credentialContext CredentialContext) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Enumerate(),
		},
	}
	err = credentialContext.base.Execute(response.Message)
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
func (credentialContext CredentialContext) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Pull(enumerationContext),
		},
	}
	err = credentialContext.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
