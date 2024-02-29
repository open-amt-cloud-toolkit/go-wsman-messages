/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package WifiPort facilitates communication with Intel® AMT devices where WiFiPort is the logical representation of wireless local area network communications hardware that conforms to the IEEE 802.11 series of specifications.
//
// It embodies properties at the lowest layers of a network stack, such as the antennas used for transmission and reception, the address permanently embedded into the hardware, and the operational bandwidth of the device.
package wifi

import (
	"encoding/xml"
	"errors"
	"strconv"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewWiFiPort returns a new instance of the WiFiPort struct.
func NewWiFiPortWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Port {
	return Port{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_WiFiPort, client),
		client: client,
	}
}

// RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
func (port Port) RequestStateChange(requestedState int) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: port.base.RequestStateChange(methods.GenerateAction(CIM_WiFiPort, "RequestStateChange"), requestedState),
		},
	}

	err = port.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	if response.Body.RequestStateChange_OUTPUT.ReturnValue != 0 {
		err = errors.New("RequestStateChange failed with return code " + strconv.Itoa(response.Body.RequestStateChange_OUTPUT.ReturnValue))
	}
	return
}

// Get retrieves the representation of the instance
func (port Port) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: port.base.Get(nil),
		},
	}

	err = port.base.Execute(response.Message)
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
func (port Port) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: port.base.Enumerate(),
		},
	}

	err = port.base.Execute(response.Message)
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
func (port Port) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: port.base.Pull(enumerationContext),
		},
	}
	err = port.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
