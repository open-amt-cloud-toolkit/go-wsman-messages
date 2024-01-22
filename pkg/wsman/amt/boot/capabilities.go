/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package boot facilitates communication with Intel® AMT devices to access the boot capabilities and boot setting data.
//
// Capabilities reports what boot options that the Intel® AMT device supports.
//
// SettingData provides configuration-related and operational parameters for the boot service in the Intel® AMT device.  In order to activate these settings use [pkg/github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/power] RequestPowerStateChange().  Notice that you can't set certain values while others are enabled (for example: You can't set UseIDER or UseSOL if a CIM_BootSourceSetting is chosen).
package boot

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewBootCapabilitiesWithClient instantiates a new Boot Capabilities service
func NewBootCapabilitiesWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Capabilities {
	return Capabilities{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_BootCapabilities, client),
	}
}

// Get retrieves the representation of the instance
func (bootCapabilities Capabilities) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: bootCapabilities.base.Get(nil),
		},
	}
	// send the message to AMT
	err = bootCapabilities.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (bootCapabilities Capabilities) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: bootCapabilities.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = bootCapabilities.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pulls instances of this class, following an Enumerate operation
func (bootCapabilities Capabilities) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: bootCapabilities.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = bootCapabilities.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
