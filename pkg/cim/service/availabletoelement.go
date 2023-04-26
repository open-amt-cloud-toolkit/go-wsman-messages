/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package service

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type AvailableToElement struct {
	base wsman.Base
}

const CIM_ServiceAvailableToElement = "CIM_ServiceAvailableToElement"

// NewServiceAvailableToElement returns a new instance of the ServiceAvailableToElement struct.
func NewServiceAvailableToElement(wsmanMessageCreator *wsman.WSManMessageCreator) AvailableToElement {
	return AvailableToElement{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_ServiceAvailableToElement)),
	}
}
func (b AvailableToElement) Get() string {
	return b.base.Get(nil)
}

func (b AvailableToElement) Enumerate() string {
	return b.base.Enumerate()
}
func (b AvailableToElement) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
