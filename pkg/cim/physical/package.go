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
func (b Package) Get() string {
	return b.base.Get(nil)
}

func (b Package) Enumerate() string {
	return b.base.Enumerate()
}
func (b Package) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
