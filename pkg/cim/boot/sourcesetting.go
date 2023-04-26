/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

const CIM_BootSourceSetting = "CIM_BootSourceSetting"

type SourceSetting struct {
	base wsman.Base
}

// NewBootSourceSetting returns a new instance of the BootSourceSetting struct.
func NewBootSourceSetting(wsmanMessageCreator *wsman.WSManMessageCreator) SourceSetting {
	return SourceSetting{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_BootSourceSetting)),
	}
}
func (b SourceSetting) Get() string {
	return b.base.Get(nil)
}

func (b SourceSetting) Enumerate() string {
	return b.base.Enumerate()
}
func (b SourceSetting) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
