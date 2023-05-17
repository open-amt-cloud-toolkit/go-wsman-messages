/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type EndpointSettings struct {
	base wsman.Base
}

const CIM_WiFiEndpoint = "CIM_WiFiEndpoint"
const CIM_WiFiEndpointSettings = "CIM_WiFiEndpointSettings"

// NewWiFiEndpointSettings returns a new instance of the WiFiEndpointSettings struct.
func NewWiFiEndpointSettings(wsmanMessageCreator *wsman.WSManMessageCreator) EndpointSettings {
	return EndpointSettings{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_WiFiEndpointSettings)),
	}
}

// Get retrieves the representation of the instance
func (b EndpointSettings) Get() string {
	return b.base.Get(nil)
}

func (b EndpointSettings) Enumerate() string {
	return b.base.Enumerate()
}
func (b EndpointSettings) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
func (b EndpointSettings) Delete(handle string) string {
	selector := wsman.Selector{Name: "Name", Value: handle}
	return b.base.Delete(selector)
}
