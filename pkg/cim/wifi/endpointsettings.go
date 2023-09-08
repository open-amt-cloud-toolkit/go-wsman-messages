/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
)

type EndpointSettings struct {
	base message.Base
}

type PullResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  message.Header
	Body    PullResponseBody
}

type PullResponseBody struct {
	PullResponse PullResponse
}

type PullResponse struct {
	Items         []CIMWiFiEndpointSettings `xml:"Items>CIM_WiFiEndpointSettings"`
	EndOfSequence string                    `xml:"EndOfSequence"`
}

type CIMWiFiEndpointSettings struct {
	AuthenticationMethod int
	BSSType              int
	ElementName          string
	EncryptionMethod     int
	InstanceID           string
	Priority             int
	SSID                 string
}

const CIM_WiFiEndpoint = "CIM_WiFiEndpoint"
const CIM_WiFiEndpointSettings = "CIM_WiFiEndpointSettings"

// NewWiFiEndpointSettings returns a new instance of the WiFiEndpointSettings struct.
func NewWiFiEndpointSettings(wsmanMessageCreator *message.WSManMessageCreator) EndpointSettings {
	return EndpointSettings{
		base: message.NewBase(wsmanMessageCreator, string(CIM_WiFiEndpointSettings)),
	}
}

// Get retrieves the representation of the instance
func (b EndpointSettings) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b EndpointSettings) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b EndpointSettings) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}

// Delete removes a the specified instance
func (b EndpointSettings) Delete(handle string) string {
	selector := message.Selector{Name: "InstanceID", Value: handle}
	return b.base.Delete(selector)
}
