/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type Chip struct {
	base wsman.Base
}

const CIM_Chip = "CIM_Chip"

// NewChip returns a new instance of the Chip struct.
func NewChip(wsmanMessageCreator *wsman.WSManMessageCreator) Chip {
	return Chip{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_Chip)),
	}
}
func (b Chip) Get() string {
	return b.base.Get(nil)
}

func (b Chip) Enumerate() string {
	return b.base.Enumerate()
}
func (b Chip) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
