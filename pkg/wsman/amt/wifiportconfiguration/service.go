/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package wifiportconfiguration facilitiates communication with Intel® AMT devices to provides management of the Wi-Fi network interfaces associated with a Wi-Fi network port.
package wifiportconfiguration

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/wifi"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewWiFiPortConfigurationServiceWithClient instantiates a new Service
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

// Enumerate returns an enumeration context which is used in a subsequent Pull call
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

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
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

	if response.Body.WiFiPortConfigurationService.LocalProfileSynchronizationEnabled == 0 {
		err = errors.New("failed to enable wifi local profile synchronization")
	}
	return
}

// AddWiFiSettings atomically creates an instance of CIM_WifiEndpointSettings from the embedded instance parameter
// and optionally an instance of CIM_IEEE8021xSettings from the embedded instance parameter (if provided),
// associates the CIM_WiFiEndpointSettings instance with the referenced instance of CIM_WiFiEndpoint using
// an instance of CIM_ElementSettingData optionally associates the newly created or referenced by parameter
// instance of CIM_IEEE8021xSettings with the instance of CIM_WiFiEndpointSettings using an instance of CIM_ConcreteComponent
// and optionally associates the referenced instance of AMT_PublicKeyCertificate (if provided) with the instance of
// CIM_IEEE8021xSettings (if provided) using an instance of CIM_CredentialContext.
//
// Additional Notes:
//
// 1) 'AddWiFiSettings' in Intel AMT Release 6.0 and later releases is permitted only to 'ADMIN_SECURITY_ADMINISTRATION_REALM' and 'ADMIN_SECURITY_LOCAL_SYSTEM_REALM '
//
// 2) When selecting the value EAP-TLS or EAP-FAST/TLS in AuthenticationProtocol property in IEEE8021xSettings - ClientCredential is mandatory.
//
// ValueMap={0, 1, 2, 3, 4, .., 32768..65535}
//
// Values={Completed with No Error, Not Supported, Failed, Invalid Parameter, Invalid Reference, Method Reserved, Vendor Specific}
func (service Service) AddWiFiSettings(wifiEndpointSettings wifi.WiFiEndpointSettingsRequest, ieee8021xSettingsInput models.IEEE8021xSettings, wifiEndpoint, clientCredential, caCredential string) (response Response, err error) {
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_WiFiPortConfigurationService, AddWiFiSettings), AMT_WiFiPortConfigurationService, nil, "", "")
	input := AddWiFiSettings_INPUT{
		WifiEndpoint: WiFiEndpoint{
			Address: "/wsman",
			ReferenceParameters: ReferenceParameters{
				H:           "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
				ResourceURI: fmt.Sprintf("%s%s", message.CIMSchema, wifi.CIM_WiFiEndpoint),
				SelectorSet: SelectorSet{
					H: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
					Selector: []Selector{
						{
							H:     "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
							Name:  "Name",
							Value: wifiEndpoint,
						},
					},
				},
			},
		},
		WiFiEndpointSettings: wifiEndpointSettings,
	}
	input.WiFiEndpointSettings.H = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings"
	if wifiEndpointSettings.AuthenticationMethod == wifi.AuthenticationMethod_WPA_IEEE8021x ||
		wifiEndpointSettings.AuthenticationMethod == wifi.AuthenticationMethod_WPA2_IEEE8021x {
		input.IEEE8021xSettings = ieee8021xSettingsInput
		input.IEEE8021xSettings.H = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_IEEE8021xSettings"
		input.CACredential = &CACredentialRequest{
			H:       "http://schemas.xmlsoap.org/ws/2004/08/addressing",
			Address: "default",
			ReferenceParameters: ReferenceParameters{
				H:           "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
				ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
				SelectorSet: SelectorSet{
					H: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
					Selector: []Selector{
						{
							H:     "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
							Name:  "InstanceID",
							Value: caCredential,
						},
					},
				},
			},
		}
		if clientCredential != "" {
			input.ClientCredential = &ClientCredentialRequest{
				H:       "http://schemas.xmlsoap.org/ws/2004/08/addressing",
				Address: "default",
				ReferenceParameters: ReferenceParameters{
					H:           "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
					ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
					SelectorSet: SelectorSet{
						H: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
						Selector: []Selector{
							{
								H:     "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
								Name:  "InstanceID",
								Value: clientCredential,
							},
						},
					},
				},
			}
		}
	}

	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(AddWiFiSettings), AMT_WiFiPortConfigurationService, &input)
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
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
	if response.Body.AddWiFiSettings_OUTPUT.ReturnValue != 0 {
		err = fmt.Errorf("AddWiFiSettings_OUTPUT.ReturnValue: %d", response.Body.AddWiFiSettings_OUTPUT.ReturnValue)
	}

	return
}

// TODO: Add UpdateWiFiSettings
// TODO: Add DeleteAllITProfiles
// TODO: Add DeleteAllUserProfiles
// TODO: Add SetApplicationRequestedRfKill
