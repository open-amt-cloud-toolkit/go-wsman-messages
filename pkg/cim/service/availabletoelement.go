/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package service

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"

type AvailableToElement struct {
	base message.Base
}

const CIM_ServiceAvailableToElement = "CIM_ServiceAvailableToElement"

// NewServiceAvailableToElement returns a new instance of the ServiceAvailableToElement struct.
func NewServiceAvailableToElement(wsmanMessageCreator *message.WSManMessageCreator) AvailableToElement {
	return AvailableToElement{
		base: message.NewBase(wsmanMessageCreator, string(CIM_ServiceAvailableToElement)),
	}
}

// Get retrieves the representation of the instance
func (b AvailableToElement) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b AvailableToElement) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b AvailableToElement) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
