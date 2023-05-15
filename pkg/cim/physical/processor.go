/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type Processor struct {
	base wsman.Base
}

const CIM_Processor = "CIM_Processor"

// NewProcessor returns a new instance of the Processor struct.
func NewProcessor(wsmanMessageCreator *wsman.WSManMessageCreator) Processor {
	return Processor{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_Processor)),
	}
}

// Get retrieves the representation of the instance
func (b Processor) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Processor) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Processor) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
