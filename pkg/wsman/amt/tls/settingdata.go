/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/json"
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName    xml.Name   `xml:"Body"`
		TlsSetting TlsSetting `xml:"AMT_TLSSettingData"`

		EnumerateResponse common.EnumerateResponse
	}
	TlsSetting struct {
		AcceptNonSecureConnections bool
		ElementName                string
		Enabled                    bool
		InstanceID                 string
		MutualAuthentication       bool
	}
)

const AMT_TLSSettingData = "AMT_TLSSettingData"

type TLSSettingData struct {
	models.SettingData
	MutualAuthentication          bool
	Enabled                       bool
	TrustedCN                     string
	AcceptNonSecureConnections    bool
	NonSecureConnectionsSupported bool
}

func (w *Response) JSON() string {
	jsonOutput, err := json.Marshal(w.Body)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}

type SettingData struct {
	base   message.Base
	client client.WSManClient
}

func NewTLSSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSManClient) SettingData {
	return SettingData{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_TLSSettingData, client),
		client: client,
	}
}

func NewTLSSettingData(wsmanMessageCreator *message.WSManMessageCreator) SettingData {
	return SettingData{
		base: message.NewBase(wsmanMessageCreator, AMT_TLSSettingData),
	}
}

// Get retrieves the representation of the instance
func (TLSSettingData SettingData) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: TLSSettingData.base.Get(nil),
		},
	}

	// send the message to AMT
	err = TLSSettingData.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerates the instances of this class
func (TLSSettingData SettingData) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: TLSSettingData.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = TLSSettingData.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pulls instances of this class, following an Enumerate operation
// func (TLSSettingData SettingData) Pull(enumerationContext string) string {
// 	return TLSSettingData.base.Pull(enumerationContext)
// }

// Put will change properties of the selected instance
// func (TLSSettingData SettingData) Put(tlsSettingData TLSSettingData) string {
// 	return TLSSettingData.base.Put(tlsSettingData, false, nil)
// }
