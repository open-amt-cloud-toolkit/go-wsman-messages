/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package software

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type Identity struct {
	base wsman.Base
}

const CIM_SoftwareIdentity = "CIM_SoftwareIdentity"

// NewSoftwareIdentity returns a new instance of the SoftwareIdentity struct.
func NewSoftwareIdentity(wsmanMessageCreator *wsman.WSManMessageCreator) Identity {
	return Identity{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_SoftwareIdentity)),
	}
}
func (b Identity) Get() string {
	return b.base.Get(nil)
}

func (b Identity) Enumerate() string {
	return b.base.Enumerate()
}
func (b Identity) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
