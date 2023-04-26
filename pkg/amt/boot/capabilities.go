/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_BootCapabilities = "AMT_BootCapabilities"

type BootCapabilities struct {
	ClassName string
	base      wsman.Base
}

func NewBootCapabilities(wsmanMessageCreator *wsman.WSManMessageCreator) BootCapabilities {
	return BootCapabilities{
		base: wsman.NewBase(wsmanMessageCreator, AMT_BootCapabilities),
	}
}
func (BootCapabilities BootCapabilities) Get() string {
	return BootCapabilities.base.Get(nil)
}
func (BootCapabilities BootCapabilities) Enumerate() string {
	return BootCapabilities.base.Enumerate()
}
func (BootCapabilities BootCapabilities) Pull(enumerationContext string) string {
	return BootCapabilities.base.Pull(enumerationContext)
}
