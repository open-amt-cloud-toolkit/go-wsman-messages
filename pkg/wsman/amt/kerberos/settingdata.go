/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kerberos

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

func NewKerberosSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SettingData {
	return SettingData{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_KerberosSettingData, client),
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
func (settingData SettingData) GetCredentialCacheState() (response Response, err error) {
	header := settingData.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_KerberosSettingData, GetCredentialCacheState), AMT_KerberosSettingData, nil, "", "")
	body := settingData.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetCredentialCacheState), AMT_KerberosSettingData, nil)

	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.WSManMessageCreator.CreateXML(header, body),
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

// TODO: Current gets SOAP schema violation from AMT
func (settingData SettingData) SetCredentialCacheState(enabled bool) (response Response, err error) {
	credentialCasheState := SetCredentialCacheState_INPUT{
		H:       fmt.Sprintf("%s%s", message.AMTSchema, AMT_KerberosSettingData),
		Enabled: enabled,
	}
	header := settingData.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_KerberosSettingData, SetCredentialCacheState), AMT_KerberosSettingData, nil, "", "")
	body := settingData.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetCredentialCacheState), AMT_KerberosSettingData, credentialCasheState)

	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.WSManMessageCreator.CreateXML(header, body),
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
