/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/actions"
)

type Port struct {
	base wsman.Base
}

const CIM_WiFiPort = "CIM_WiFiPort"

// NewWiFiPort returns a new instance of the WiFiPort struct.
func NewWiFiPort(wsmanMessageCreator *wsman.WSManMessageCreator) Port {
	return Port{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_WiFiPort)),
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
