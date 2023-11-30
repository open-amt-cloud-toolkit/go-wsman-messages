/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"

	"github.com/google/uuid"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
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
		XMLName           xml.Name `xml:"Body"`
		Setup             Setup    `xml:"AMT_SetupAndConfigurationService"`
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
		GetUuid_OUTPUT    GetUuid_OUTPUT `xml:"GetUuid_OUTPUT"`
	}

	Setup struct {
		CreationClassName             string
		ElementName                   string
		EnabledState                  int
		Name                          string
		PasswordModel                 int
		ProvisioningMode              int
		ProvisioningServerOTP         string
		ProvisioningState             int
		RequestedState                int
		SystemCreationClassName       string
		SystemName                    string
		ZeroTouchConfigurationEnabled bool
	}
	PullResponse struct {
		Items []Item
	}
	Item struct {
		Setup Setup `xml:"AMT_SetupAndConfigurationService"`
	}

	GetUuid_OUTPUT struct {
		UUID string `xml:"UUID"`
	}
)
type UnprovisionResponse struct {
	XMLName xml.Name        `xml:"Envelope"`
	Header  message.Header  `xml:"Header"`
	Body    UnprovisionBody `xml:"Body"`
}

type UnprovisionBody struct {
	XMLName            xml.Name           `xml:"Body"`
	Unprovision_OUTPUT Unprovision_OUTPUT `xml:"Unprovision_OUTPUT"`
}

type Unprovision_OUTPUT struct {
	XMLName     xml.Name `xml:"Unprovision_OUTPUT"`
	ReturnValue int
}

type SetupAndConfigurationService struct {
	models.CredentialManagementService
	AMT_SetupAndConfigurationService struct {
		CreationClassName             string
		ElementName                   string
		EnabledState                  string
		Name                          string
		PasswordModel                 string
		ProvisioningMode              string
		ProvisioningServerOTP         string
		ProvisioningState             string
		RequestedState                string
		SystemCreationClassName       string
		SystemName                    string
		ZeroTouchConfigurationEnabled string
	}
}

func (w *Response) JSON() string {
	jsonOutput, err := json.Marshal(w.Body)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}

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

type Service struct {
	base   message.Base
	client client.WSMan
}

func NewSetupAndConfigurationService(wsmanMessageCreator *message.WSManMessageCreator) Service {
	return Service{
		base:   message.NewBase(wsmanMessageCreator, AMT_SetupAndConfigurationService),
		client: nil,
	}
}

func NewSetupAndConfigurationServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_SetupAndConfigurationService, client),
		client: client,
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

// // Put will change properties of the selected instance
// func (s Service) Put(setupAndConfigurationService SetupAndConfigurationService) string {
// 	return s.base.Put(setupAndConfigurationService, false, nil)
// }
// func (s Service) CommitChanges() string {
// 	header := s.base.WSManMessageCreator.CreateHeader(string(actions.CommitChanges), AMT_SetupAndConfigurationService, nil, "", "")
// 	body := s.base.WSManMessageCreator.CreateBody("CommitChanges_INPUT", AMT_SetupAndConfigurationService, nil)
// 	return s.base.WSManMessageCreator.CreateXML(header, body)
// }

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

// func (s Service) SetMEBXPassword(password string) string {
// 	header := s.base.WSManMessageCreator.CreateHeader(string(actions.SetMEBxPassword), AMT_SetupAndConfigurationService, nil, "", "")
// 	body := fmt.Sprintf(`<Body><h:SetMEBxPassword_INPUT xmlns:h="%s%s"><h:Password>%s</h:Password></h:SetMEBxPassword_INPUT></Body>`, s.base.WSManMessageCreator.ResourceURIBase, AMT_SetupAndConfigurationService, password)
// 	return s.base.WSManMessageCreator.CreateXML(header, body)
// }

// func (s Service) Unprovision(provisioningMode int) string {
// 	if provisioningMode == 0 {
// 		provisioningMode = 1
// 	}
// 	header := s.base.WSManMessageCreator.CreateHeader(string(actions.Unprovision), AMT_SetupAndConfigurationService, nil, "", "")
// 	body := fmt.Sprintf(`<Body><h:Unprovision_INPUT xmlns:h="%s%s"><h:ProvisioningMode>%d</h:ProvisioningMode></h:Unprovision_INPUT></Body>`, s.base.WSManMessageCreator.ResourceURIBase, AMT_SetupAndConfigurationService, provisioningMode)
// 	return s.base.WSManMessageCreator.CreateXML(header, body)
// }
