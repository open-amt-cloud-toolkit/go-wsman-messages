/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type Package struct {
	base wsman.Base
}

const CIM_PhysicalPackage = "CIM_PhysicalPackage"

// NewPhysicalPackage returns a new instance of the PhysicalPackage struct.
func NewPhysicalPackage(wsmanMessageCreator *wsman.WSManMessageCreator) Package {
	return Package{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_PhysicalPackage)),
	}
}

// Get retrieves the representation of the instance
func (b Package) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Package) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Package) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
