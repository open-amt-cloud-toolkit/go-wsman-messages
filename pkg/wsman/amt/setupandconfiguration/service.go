/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"

	"github.com/google/uuid"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

func (w *Response) DecodeUUID() (amtUuid string, err error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(w.Body.GetUuid_OUTPUT.UUID)
	if err != nil {
		return
	}
	rearrangeBytes := []byte{
		decodedBytes[3], decodedBytes[2], decodedBytes[1], decodedBytes[0],
		decodedBytes[5], decodedBytes[4],
		decodedBytes[7], decodedBytes[6],
		decodedBytes[8], decodedBytes[9],
		decodedBytes[10], decodedBytes[11], decodedBytes[12], decodedBytes[13], decodedBytes[14], decodedBytes[15],
	}
	uuidSlice, err := uuid.FromBytes(rearrangeBytes)
	if err != nil {
		return
	}
	amtUuid = uuidSlice.String()
	return
}

func NewSetupAndConfigurationServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_SetupAndConfigurationService, client),
	}
}
func (s Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Get(nil),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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
func (s Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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
func (s Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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
func (s Service) Put(setupAndConfigurationService SetupAndConfigurationServiceRequest) (response Response, err error) {
	setupAndConfigurationService.H = fmt.Sprintf("%s%s", message.AMTSchema, AMT_SetupAndConfigurationService)
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Put(setupAndConfigurationService, false, nil),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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
func (s Service) CommitChanges() (response Response, err error) {
	header := s.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_SetupAndConfigurationService, CommitChanges), AMT_SetupAndConfigurationService, nil, "", "")
	body := s.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(CommitChanges), AMT_SetupAndConfigurationService, nil)
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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

// Gets the AMT UUID from the device
func (s Service) GetUuid() (response Response, err error) {
	header := s.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_SetupAndConfigurationService, GetUuid), AMT_SetupAndConfigurationService, nil, "", "")
	body := s.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetUuid), AMT_SetupAndConfigurationService, nil)
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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

func (s Service) SetMEBXPassword(password string) (response Response, err error) {
	header := s.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_SetupAndConfigurationService, SetMEBxPassword), AMT_SetupAndConfigurationService, nil, "", "")
	mebxPassword := MEBXPassword{
		Password: password,
	}
	body := s.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetMEBxPassword), AMT_SetupAndConfigurationService, &mebxPassword)
	// body := fmt.Sprintf(`<Body><h:SetMEBxPassword_INPUT xmlns:h="%s%s"><h:Password>%s</h:Password></h:SetMEBxPassword_INPUT></Body>`, s.base.WSManMessageCreator.ResourceURIBase, AMT_SetupAndConfigurationService, password)
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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

func (s Service) Unprovision(provisioningMode int) (response Response, err error) {
	if provisioningMode == 0 {
		provisioningMode = 1
	}
	pMode := ProvisioningMode{
		ProvisioningMode: provisioningMode,
	}
	header := s.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_SetupAndConfigurationService, Unprovision), AMT_SetupAndConfigurationService, nil, "", "")
	body := s.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(Unprovision), AMT_SetupAndConfigurationService, &pMode)
	// body := fmt.Sprintf(`<Body><h:Unprovision_INPUT xmlns:h="%s%s"><h:ProvisioningMode>%d</h:ProvisioningMode></h:Unprovision_INPUT></Body>`, s.base.WSManMessageCreator.ResourceURIBase, AMT_SetupAndConfigurationService, provisioningMode)
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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
