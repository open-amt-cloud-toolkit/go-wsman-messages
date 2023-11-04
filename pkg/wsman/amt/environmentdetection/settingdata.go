/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package environmentdetection

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

const AMT_EnvironmentDetectionSettingData = "AMT_EnvironmentDetectionSettingData"

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName              xml.Name             `xml:"Body"`
		DetectionSettingData DetectionSettingData `xml:"AMT_EnvironmentDetectionSettingData"`

		EnumerateResponse common.EnumerateResponse
	}
	DetectionSettingData struct {
		DetectionStrings string
		ElementName      string
		InstanceID       string
	}
)

type EnvironmentDetectionSettingData struct {
	models.SettingData
	DetectionAlgorithm         DetectionAlgorithm
	DetectionStrings           []string
	DetectionIPv6LocalPrefixes []string
}

type DetectionAlgorithm uint8

const (
	LocalDomains DetectionAlgorithm = iota
	RemoteURLs
)

type SettingData struct {
	base   message.Base
	client client.WSManClient
}

func NewEnvironmentDetectionSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSManClient) SettingData {
	return SettingData{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_EnvironmentDetectionSettingData, client),
		client: client,
	}
}

func NewEnvironmentDetectionSettingData(wsmanMessageCreator *message.WSManMessageCreator) SettingData {
	return SettingData{
		base: message.NewBase(wsmanMessageCreator, AMT_EnvironmentDetectionSettingData),
	}
}

// Get retrieves the representation of the instance
func (sd SettingData) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: sd.base.Get(nil),
		},
	}

	// send the message to AMT
	err = sd.base.Execute(response.Message)
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
func (sd SettingData) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: sd.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = sd.base.Execute(response.Message)
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
// func (EnvironmentDetectionSettingData SettingData) Pull(enumerationContext string) string {
// 	return EnvironmentDetectionSettingData.base.Pull(enumerationContext)
// }

// // Put will change properties of the selected instance
// func (EnvironmentDetectionSettingData SettingData) Put(environmentDetectionSettingData EnvironmentDetectionSettingData) string {
// 	return EnvironmentDetectionSettingData.base.Put(environmentDetectionSettingData, false, nil)
// }
