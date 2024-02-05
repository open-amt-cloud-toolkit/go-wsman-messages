/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewIEEE8021xSettings returns a new instance of the IEEE8021xSettings struct.
func NewIEEE8021xSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_IEEE8021xSettings, client),
		client: client,
	}
}

// TODO: Figure out how to call GET requiring resourceURIs and Selectors

// Enumerate returns an enumeration context which is used in a subsequent Pull call
func (settings Settings) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Enumerate(),
		},
	}

	err = settings.base.Execute(response.Message)
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
func (settings Settings) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Pull(enumerationContext),
		},
	}
	err = settings.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
