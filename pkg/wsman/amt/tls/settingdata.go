/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewTLSSettingDataWithClient instantiates a new SettingData
func NewTLSSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SettingData {
	return SettingData{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_TLSSettingData, client),
	}
}

// Get retrieves the representation of the instance
func (settingData SettingData) Get(instanceID string) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: instanceID,
	}
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Get(&selector),
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

// Enumerates the instances of this class
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

// Pulls instances of this class, following an Enumerate operation
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

// Put changes properties of the selected instance.
// The following properties must be included in any representation of SettingDataRequest:
//
// - ElementName(cannot be modified)
//
// - InstanceID (cannot be modified)
//
// - Enabled.
//
// This method will not modify the flash ("Enabled" property) until setupandconfiguration.CommitChanges() is issued and performed successfully.
func (settingData SettingData) Put(instanceID string, tlsSettingData SettingDataRequest) (response Response, err error) {
	tlsSettingData.H = fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSSettingData)
	selector := message.Selector{
		Name:  "InstanceID",
		Value: instanceID,
	}
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Put(tlsSettingData, true, &selector),
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
