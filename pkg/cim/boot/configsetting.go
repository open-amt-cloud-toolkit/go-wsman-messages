/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/actions"
)

const CIM_BootConfigSetting = "CIM_BootConfigSetting"

type ConfigSetting struct {
	base wsman.Base
}

// NewBootConfigSetting returns a new instance of the BootConfigSetting struct.
func NewBootConfigSetting(wsmanMessageCreator *wsman.WSManMessageCreator) ConfigSetting {
	return ConfigSetting{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_BootConfigSetting)),
	}
}

func (b ConfigSetting) Get() string {
	return b.base.Get(nil)
}

func (b ConfigSetting) Enumerate() string {
	return b.base.Enumerate()
}
func (b ConfigSetting) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}

func (b ConfigSetting) ChangeBootOrder(source string) string {
	header := b.base.WSManMessageCreator.CreateHeader(string(actions.ChangeBootOrder), string(CIM_BootConfigSetting), nil, "", "")
	body := fmt.Sprintf(`<Body><h:ChangeBootOrder_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting"><h:Source><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="InstanceID">%s</Selector></SelectorSet></ReferenceParameters></h:Source></h:ChangeBootOrder_INPUT></Body>`, source)
	return b.base.WSManMessageCreator.CreateXML(header, body)
}
