/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportconfiguration

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

func NewWiFiPortConfigurationServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_WiFiPortConfigurationService, client),
	}
}

// Get retrieves the representation of the instance
func (service Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Get(nil),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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
func (service Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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
func (service Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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
func (service Service) Put(wiFiPortConfigurationService WiFiPortConfigurationServiceRequest) (response Response, err error) {
	//wiFiPortConfigurationService.XMLSchema = "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService"
	wiFiPortConfigurationService.H = fmt.Sprintf("%s%s", message.AMTSchema, AMT_WiFiPortConfigurationService)
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Put(wiFiPortConfigurationService, false, nil),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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

// AddWiFiSettings atomically creates instances and associates them based on the input parameters.
// func (service Service) AddWiFiSettings(wifiEndpointSettings wifi.WiFiEndpointSettings_INPUT, ieee8021xSettingsInput *models.IEEE8021xSettings, wifiEndpoint, clientCredential, caCredential string) (response Response, err error) {
// 	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_WiFiPortConfigurationService, AddWiFiSettings), AMT_WiFiPortConfigurationService, nil, "", "")
// 	input := AddWiFiSettings_INPUT{
// 		WifiEndpoint: WiFiEndpoint{
// 			Address: "/wsman",
// 			ReferenceParameters: ReferenceParameters{
// 				H:           "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
// 				ResourceURI: fmt.Sprintf("%s%s", message.CIMSchema, wifi.CIM_WiFiEndpoint),
// 				SelectorSet: SelectorSet{
// 					H: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
// 					Selector: []Selector{
// 						{
// 							H:     "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
// 							Name:  "Name",
// 							Value: wifiEndpoint,
// 						},
// 					},
// 				},
// 			},
// 		},
// 		WiFiEndpointSettings: wifiEndpointSettings,
// 	}
// 	input.WiFiEndpointSettings.H = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings"
// 	if ieee8021xSettingsInput != nil {
// 		input.IEEE8021xSettings = ieee8021xSettingsInput
// 		input.IEEE8021xSettings.H = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_IEEE8021xSettings"
// 		input.CACredential = &CACredentialRequest{
// 			H:       "http://schemas.xmlsoap.org/ws/2004/08/addressing",
// 			Address: "default",
// 			ReferenceParameters: ReferenceParameters{
// 				H:           "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
// 				ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
// 				SelectorSet: SelectorSet{
// 					H: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
// 					Selector: []Selector{
// 						{
// 							H:     "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
// 							Name:  "InstanceID",
// 							Value: caCredential,
// 						},
// 					},
// 				},
// 			},
// 		}
// 		if clientCredential != "" {
// 			input.ClientCredential = &ClientCredentialRequest{
// 				H:       "http://schemas.xmlsoap.org/ws/2004/08/addressing",
// 				Address: "default",
// 				ReferenceParameters: ReferenceParameters{
// 					H:           "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
// 					ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
// 					SelectorSet: SelectorSet{
// 						H: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
// 						Selector: []Selector{
// 							{
// 								H:     "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
// 								Name:  "InstanceID",
// 								Value: clientCredential,
// 							},
// 						},
// 					},
// 				},
// 			}
// 		}
// 	}

// 	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(AddWiFiSettings), AMT_WiFiPortConfigurationService, &input)
// 	response = Response{
// 		Message: &client.Message{
// 			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
// 		},
// 	}
// 	// send the message to AMT
// 	err = service.base.Execute(response.Message)
// 	if err != nil {
// 		return
// 	}
// 	// put the xml response into the go struct
// 	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// TODO: Add UpdateWiFiSettings
// TODO: Add DeleteAllITProfiles
// TODO: Add DeleteAllUserProfiles
// TODO: Add SetApplicationRequestedRfKill
