/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mediaaccess

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

const CIM_MediaAccessDevice = "CIM_MediaAccessDevice"

type Device struct {
	base wsman.Base
}

// NewMediaAccessDevice returns a new instance of the MediaAccessDevice struct.
func NewMediaAccessDevice(wsmanMessageCreator *wsman.WSManMessageCreator) Device {
	return Device{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_MediaAccessDevice)),
	}
}

// Get retrieves the representation of the instance
func (b Device) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Device) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Device) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
