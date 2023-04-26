/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package bios

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type Element struct {
	base wsman.Base
}

const CIM_BiosElement = "CIM_BIOSElement"

// NewBIOSElement returns a new instance of the BIOSElement struct.
func NewBIOSElement(wsmanMessageCreator *wsman.WSManMessageCreator) Element {
	return Element{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_BiosElement)),
	}
}
func (b Element) Get() string {
	return b.base.Get(nil)
}

func (b Element) Enumerate() string {
	return b.base.Enumerate()
}
func (b Element) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
