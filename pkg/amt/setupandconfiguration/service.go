/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

import (
	"encoding/json"
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
)

const AMT_SetupAndConfigurationService = "AMT_SetupAndConfigurationService"

type (
	Response struct {
		*wsman.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName xml.Name `xml:"Body"`
		Setup   Setup    `xml:"AMT_SetupAndConfigurationService"`

		EnumerateResponse common.EnumerateResponse
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

type Service struct {
	base   message.Base
	client wsman.WSManClient
}

func NewSetupAndConfigurationService(wsmanMessageCreator *message.WSManMessageCreator) Service {
	return Service{
		base:   message.NewBase(wsmanMessageCreator, AMT_SetupAndConfigurationService),
		client: nil,
	}
}

func NewSetupAndConfigurationServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client wsman.WSManClient) Service {
	return Service{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_SetupAndConfigurationService, client),
		client: client,
	}
}
func (s Service) Get() (response Response, err error) {
	response = Response{
		Message: &wsman.Message{
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
		Message: &wsman.Message{
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
// func (s Service) Pull(enumerationContext string) string {
// 	return s.base.Pull(enumerationContext)
// }

// // Put will change properties of the selected instance
// func (s Service) Put(setupAndConfigurationService SetupAndConfigurationService) string {
// 	return s.base.Put(setupAndConfigurationService, false, nil)
// }
// func (s Service) CommitChanges() string {
// 	header := s.base.WSManMessageCreator.CreateHeader(string(actions.CommitChanges), AMT_SetupAndConfigurationService, nil, "", "")
// 	body := s.base.WSManMessageCreator.CreateBody("CommitChanges_INPUT", AMT_SetupAndConfigurationService, nil)
// 	return s.base.WSManMessageCreator.CreateXML(header, body)
// }

// func (s Service) GetUuid() string {
// 	header := s.base.WSManMessageCreator.CreateHeader(string(actions.GetUuid), AMT_SetupAndConfigurationService, nil, "", "")
// 	body := s.base.WSManMessageCreator.CreateBody("GetUuid_INPUT", AMT_SetupAndConfigurationService, nil)
// 	return s.base.WSManMessageCreator.CreateXML(header, body)
// }

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
