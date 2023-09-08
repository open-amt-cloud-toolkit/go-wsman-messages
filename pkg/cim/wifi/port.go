/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/actions"
)

type Port struct {
	base message.Base
}
type RequestStateChangeResponse struct {
	XMLName xml.Name               `xml:"Envelope"`
	Header  message.Header         `xml:"Header"`
	Body    RequestStateChangeBody `xml:"Body"`
}
type RequestStateChangeBody struct {
	XMLName                   xml.Name                  `xml:"Body"`
	RequestStateChange_OUTPUT RequestStateChange_OUTPUT `xml:"RequestStateChange_OUTPUT"`
}
type RequestStateChange_OUTPUT struct {
	XMLName     xml.Name `xml:"RequestStateChange_OUTPUT"`
	ReturnValue int
}

const CIM_WiFiPort = "CIM_WiFiPort"

// NewWiFiPort returns a new instance of the WiFiPort struct.
func NewWiFiPort(wsmanMessageCreator *message.WSManMessageCreator) Port {
	return Port{
		base: message.NewBase(wsmanMessageCreator, string(CIM_WiFiPort)),
	}
}

// RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
func (w Port) RequestStateChange(requestedState int) string {
	return w.base.RequestStateChange(actions.RequestStateChange(string(CIM_WiFiPort)), requestedState)
}

// Get retrieves the representation of the instance
func (b Port) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Port) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Port) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
