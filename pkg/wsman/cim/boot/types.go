/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type ConfigSetting struct {
	base   message.Base
	client client.WSMan
}

type SourceSetting struct {
	base   message.Base
	client client.WSMan
}

type Service struct {
	base   message.Base
	client client.WSMan
}

type Source string

// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                  xml.Name          `xml:"Body"`
		ConfigSettingGetResponse BootConfigSetting `xml:"CIM_BootConfigSetting"`
		SourceSettingGetResponse BootSourceSetting `xml:"CIM_BootSourceSetting"`
		ServiceGetResponse       BootService       `xml:"CIM_BootService"`
		EnumerateResponse        common.EnumerateResponse
		PullResponse             PullResponse           `xml:"PullResponse"`
		ChangeBootOrder_OUTPUT   ChangeBootOrder_OUTPUT `xml:"ChangeBootOrder_OUTPUT"`
	}

	BootConfigSetting struct {
		XMLName     xml.Name `xml:"CIM_BootConfigSetting"`
		InstanceID  string   `xml:"InstanceID"`
		ElementName string   `xml:"ElementName"`
	}

	BootSourceSetting struct {
		XMLName              xml.Name `xml:"CIM_BootSourceSetting"`
		ElementName          string   `xml:"ElementName"`
		InstanceID           string   `xml:"InstanceID"`
		StructuredBootString string   `xml:"StructuredBootString"`
		BIOSBootString       string   `xml:"BIOSBootString"`
		BootString           string   `xml:"BootString"`
		FailThroughSupported int      `xml:"FailThroughSupported"`
	}

	BootService struct {
		XMLName                 xml.Name                 `xml:"CIM_BootService"`
		Name                    string                   `xml:"Name"`
		CreationClassName       string                   `xml:"CreationClassName"`
		SystemName              string                   `xml:"SystemName"`
		SystemCreationClassName string                   `xml:"SystemCreationClassName"`
		ElementName             string                   `xml:"ElementName"`
		OperationalStatus       models.OperationalStatus `xml:"OperationalStatus"`
		EnabledState            models.EnabledState      `xml:"EnabledState"`
		RequestedState          models.RequestedState    `xml:"RequestedState"`
	}

	PullResponse struct {
		BootSourceSettingItems []BootSourceSetting `xml:"Items>CIM_BootSourceSetting"`
		BootConfigSettingItems []BootConfigSetting `xml:"Items>CIM_BootConfigSetting"`
		BootServiceItems       []BootService       `xml:"Items>CIM_BootService"`
	}

	ChangeBootOrder_OUTPUT struct {
		ReturnValue int `xml:"ReturnValue"`
	}
)
