/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

func NewBootConfigSettingWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) ConfigSetting {
	return ConfigSetting{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_BootConfigSetting, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (configSetting ConfigSetting) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: configSetting.base.Get(nil),
		},
	}

	err = configSetting.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (configSetting ConfigSetting) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: configSetting.base.Enumerate(),
		},
	}

	err = configSetting.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pulls instances of this class, following an Enumerate operation
func (configSetting ConfigSetting) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: configSetting.base.Pull(enumerationContext),
		},
	}
	err = configSetting.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (configSetting ConfigSetting) ChangeBootOrder(source Source) (response Response, err error) {
	header := configSetting.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(CIM_BootConfigSetting, ChangeBootOrder), CIM_BootConfigSetting, nil, "", "")
	body := fmt.Sprintf(`<Body><h:ChangeBootOrder_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting"><h:Source><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="InstanceID">%s</Selector></SelectorSet></ReferenceParameters></h:Source></h:ChangeBootOrder_INPUT></Body>`, source)
	response = Response{
		Message: &client.Message{
			XMLInput: configSetting.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = configSetting.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
