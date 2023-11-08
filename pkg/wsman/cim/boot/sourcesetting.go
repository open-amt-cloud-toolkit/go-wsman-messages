/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
)

const CIM_BootSourceSetting = "CIM_BootSourceSetting"

type SourceSetting struct {
	base message.Base
}

// NewBootSourceSetting returns a new instance of the BootSourceSetting struct.
func NewBootSourceSetting(wsmanMessageCreator *message.WSManMessageCreator) SourceSetting {
	return SourceSetting{
		base: message.NewBase(wsmanMessageCreator, string(CIM_BootSourceSetting)),
	}
}

// Get retrieves the representation of the instance
func (b SourceSetting) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b SourceSetting) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b SourceSetting) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
