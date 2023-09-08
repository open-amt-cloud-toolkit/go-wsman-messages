/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
)

type Chassis struct {
	base message.Base
}

const CIM_Chassis = "CIM_Chassis"

// NewChassis returns a new instance of the Chassis struct.
func NewChassis(wsmanMessageCreator *message.WSManMessageCreator) Chassis {
	return Chassis{
		base: message.NewBase(wsmanMessageCreator, string(CIM_Chassis)),
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
