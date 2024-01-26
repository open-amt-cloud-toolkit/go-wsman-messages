/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// Instantiates a new Boot Setting Data service
func NewBootSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SettingData {
	return SettingData{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_BootSettingData, client),
	}
}

// Get retrieves the representation of the instance
func (settingData SettingData) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Get(nil),
		},
	}
	// send the message to AMT
	err = settingData.base.Execute(response.Message)
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

// Enumerate returns an enumeration context which is used in a subsequent Pull call
func (settingData SettingData) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = settingData.base.Execute(response.Message)
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

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (settingData SettingData) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = settingData.base.Execute(response.Message)
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

// Put will change properties of the selected instance
func (settingData SettingData) Put(bootSettingData BootSettingDataRequest) (response Response, err error) {
	bootSettingData.H = fmt.Sprintf("%s%s", message.AMTSchema, AMT_BootSettingData)
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Put(bootSettingData, false, nil),
		},
	}
	// send the message to AMT
	err = settingData.base.Execute(response.Message)
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
