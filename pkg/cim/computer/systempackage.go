/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package computer

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

const CIM_ComputerSystemPackage = "CIM_ComputerSystemPackage"

type SystemPackage struct {
	base wsman.Base
}

// NewComputerSystemPackage returns a new instance of the ComputerSystemPackage struct.
func NewComputerSystemPackage(wsmanMessageCreator *wsman.WSManMessageCreator) SystemPackage {
	return SystemPackage{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_ComputerSystemPackage)),
	}
}

// Get retrieves the representation of the instance
func (b SystemPackage) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b SystemPackage) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b SystemPackage) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
