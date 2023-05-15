/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type Chassis struct {
	base wsman.Base
}

const CIM_Chassis = "CIM_Chassis"

// NewChassis returns a new instance of the Chassis struct.
func NewChassis(wsmanMessageCreator *wsman.WSManMessageCreator) Chassis {
	return Chassis{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_Chassis)),
	}
}

// Get retrieves the representation of the instance
func (b Chassis) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Chassis) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Chassis) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
