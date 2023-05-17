/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type Card struct {
	base wsman.Base
}

const CIM_Card = "CIM_Card"

// NewCard returns a new instance of the Card struct.
func NewCard(wsmanMessageCreator *wsman.WSManMessageCreator) Card {
	return Card{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_Card)),
	}
}

// Get retrieves the representation of the instance
func (b Card) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Card) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Card) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
