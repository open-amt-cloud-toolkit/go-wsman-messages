/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Port struct {
	base   message.Base
	client client.WSMan
}

type EndpointSettings struct {
	base message.Base
}

// OUTPUT
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                   xml.Name `xml:"Body"`
		WiFiPortGetResponse       WiFiPort
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse
		RequestStateChange_OUTPUT message.ReturnValue
	}

	PullResponse struct {
		XMLName               xml.Name               `xml:"PullResponse"`
		EndpointSettingsItems []WiFiEndpointSettings `xml:"Items>CIM_WiFiEndpointSettings"`
		WiFiPortItems         []WiFiPort             `xml:"Items>CIM_WiFiPort"`
	}

	WiFiEndpointSettings struct {
		XMLName              xml.Name `xml:"CIM_WiFiEndpointSettings"`
		AuthenticationMethod AuthenticationMethod
		BSSType              BSSType
		ElementName          string
		EncryptionMethod     EncryptionMethod
		InstanceID           string
		Priority             int
		SSID                 string
	}

	WiFiPort struct {
		XMLName                 xml.Name `xml:"CIM_WiFiPort"`
		LinkTechnology          LinkTechnology
		DeviceID                string
		CreationClassName       string
		SystemName              string
		SystemCreationClassName string
		ElementName             string
		HealthState             models.HealthState
		EnabledState            models.EnabledState
		RequestedState          models.RequestedState
		PortType                int
		PermanentAddress        string
	}
)

// INPUT
// Request Types
type (
	WiFiEndpointSettings_INPUT struct {
		XMLName              xml.Name `xml:"CIM_WiFiEndpointSettings"`
		H                    string   `xml:"xmlns:q,attr"`
		AuthenticationMethod AuthenticationMethod
		BSSType              BSSType
		ElementName          string
		EncryptionMethod     EncryptionMethod
		InstanceID           string
		Priority             int
		SSID                 string
	}
	WiFiEndpointSettingsInput struct {
		XMLName xml.Name `xml:"h:WiFiEndpointSettingsInput"`
		H       string   `xml:"xmlns:q,attr"`
		// SettingData
		ElementName          string               `xml:"q:ElementName,omitempty"`
		InstanceID           string               `xml:"q:InstanceID,omitempty"`
		AuthenticationMethod AuthenticationMethod `xml:"q:AuthenticationMethod,omitempty"`
		EncryptionMethod     EncryptionMethod     `xml:"q:EncryptionMethod,omitempty"`
		SSID                 string               `xml:"q:SSID,omitempty"` // Max Length 32
		Priority             int                  `xml:"q:Priority,omitempty"`
		PSKPassPhrase        string               `xml:"q:PSKPassPhrase,omitempty"` // Min Length 8 Max Length 63
		BSSType              BSSType              `xml:"q:BSSType,omitempty"`
		Keys                 []string             `xml:"q:Keys,omitempty"` // OctetString ArrayType=Indexed Max Length 256
		KeyIndex             int                  `xml:"q:KeyIndex,omitempty"`
		PSKValue             int                  `xml:"q:PSKValue,omitempty"` // OctetString
	}
)

type (
	LinkTechnology       int
	AuthenticationMethod int
	BSSType              int
	EncryptionMethod     int
)
