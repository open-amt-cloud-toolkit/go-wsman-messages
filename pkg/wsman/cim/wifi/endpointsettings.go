/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewWiFiEndpointSettings returns a new instance of the WiFiEndpointSettings struct.
func NewWiFiEndpointSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) EndpointSettings {
	return EndpointSettings{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIM_WiFiEndpointSettings, client),
	}
}

// TODO: Figure out how to call GET requiring resourceURIs and Selectors
// Get retrieves the representation of the instance

// Enumerate returns an enumeration context which is used in a subsequent Pull call
func (endpointSettings EndpointSettings) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: endpointSettings.base.Enumerate(),
		},
	}

	err = endpointSettings.base.Execute(response.Message)
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
func (endpointSettings EndpointSettings) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: endpointSettings.base.Pull(enumerationContext),
		},
	}
	err = endpointSettings.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Delete removes a the specified instance
func (endpointSettings EndpointSettings) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "InstanceID", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: endpointSettings.base.Delete(selector),
		},
	}

	err = endpointSettings.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
