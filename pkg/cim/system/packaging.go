/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package system

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type Packaging struct {
	base wsman.Base
}

const CIM_SystemPackaging = "CIM_SystemPackaging"

// NewSystemPackaging returns a new instance of the SystemPackaging struct.
func NewSystemPackaging(wsmanMessageCreator *wsman.WSManMessageCreator) Packaging {
	return Packaging{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_SystemPackaging)),
	}
}

// Get retrieves the representation of the instance
func (b Packaging) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Packaging) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Packaging) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
