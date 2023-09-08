/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
)

type Memory struct {
	base message.Base
}

const CIM_PhysicalMemory = "CIM_PhysicalMemory"

// NewPhysicalMemory returns a new instance of the PhysicalMemory struct.
func NewPhysicalMemory(wsmanMessageCreator *message.WSManMessageCreator) Memory {
	return Memory{
		base: message.NewBase(wsmanMessageCreator, string(CIM_PhysicalMemory)),
	}
}

// Get retrieves the representation of the instance
func (b Memory) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Memory) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Memory) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
