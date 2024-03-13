/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

const GetUuid_BODY = "<h:GetUuid_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService\"></h:GetUuid_INPUT>"

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: SetupAndConfigurationServiceResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"RequestedState\":0,\"EnabledState\":0,\"ElementName\":\"\",\"SystemCreationClassName\":\"\",\"SystemName\":\"\",\"CreationClassName\":\"\",\"Name\":\"\",\"ProvisioningMode\":0,\"ProvisioningState\":0,\"ZeroTouchConfigurationEnabled\":false,\"ProvisioningServerOTP\":\"\",\"ConfigurationServerFQDN\":\"\",\"PasswordModel\":0,\"DhcpDNSSuffix\":\"\",\"TrustedDNSSuffix\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"SetupAndConfigurationServiceItems\":null},\"GetUuid_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"UUID\":\"\"},\"Unprovision_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0},\"CommitChanges_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0},\"SetMEBxPassword_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: SetupAndConfigurationServiceResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    requestedstate: 0\n    enabledstate: 0\n    elementname: \"\"\n    systemcreationclassname: \"\"\n    systemname: \"\"\n    creationclassname: \"\"\n    name: \"\"\n    provisioningmode: 0\n    provisioningstate: 0\n    zerotouchconfigurationenabled: false\n    provisioningserverotp: \"\"\n    configurationserverfqdn: \"\"\n    passwordmodel: 0\n    dhcpdnssuffix: \"\"\n    trusteddnssuffix: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    setupandconfigurationserviceitems: []\ngetuuid_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    uuid: \"\"\nunprovision_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\ncommitchanges_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\nsetmebxpassword_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_SetupAndConfigurationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/setupandconfiguration",
	}
	elementUnderTest := NewSetupAndConfigurationServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_SetupAndConfiguration Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_SetupAndConfigurationService Get wsman message",
				AMT_SetupAndConfigurationService,
				wsmantesting.GET, "",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: SetupAndConfigurationServiceResponse{
						XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService", Local: "AMT_SetupAndConfigurationService"},
						CreationClassName:             AMT_SetupAndConfigurationService,
						ElementName:                   "Intel(r) AMT Setup and Configuration Service",
						EnabledState:                  5,
						Name:                          "Intel(r) AMT Setup and Configuration Service",
						PasswordModel:                 1,
						ProvisioningMode:              1,
						ProvisioningServerOTP:         "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
						ProvisioningState:             2,
						RequestedState:                12,
						SystemCreationClassName:       "CIM_ComputerSystem",
						SystemName:                    "Intel(r) AMT",
						ZeroTouchConfigurationEnabled: true,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_SetupAndConfigurationService Enumerate wsman message",
				AMT_SetupAndConfigurationService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D3000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_SetupAndConfigurationService Pull wsman message",
				AMT_SetupAndConfigurationService,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						SetupAndConfigurationServiceItems: []SetupAndConfigurationServiceResponse{
							{
								XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService", Local: "AMT_SetupAndConfigurationService"},
								CreationClassName:             AMT_SetupAndConfigurationService,
								ElementName:                   "Intel(r) AMT Setup and Configuration Service",
								EnabledState:                  5,
								Name:                          "Intel(r) AMT Setup and Configuration Service",
								PasswordModel:                 1,
								ProvisioningMode:              1,
								ProvisioningServerOTP:         "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
								ProvisioningState:             2,
								RequestedState:                12,
								SystemCreationClassName:       "CIM_ComputerSystem",
								SystemName:                    "Intel(r) AMT",
								ZeroTouchConfigurationEnabled: true,
							},
						},
					},
				},
			},
			//GetUuid
			{
				"should return a valid AMT_GetUuid response",
				AMT_SetupAndConfigurationService,
				methods.GenerateAction(AMT_SetupAndConfigurationService, GetUuid),
				GetUuid_BODY,
				func() (Response, error) {
					client.CurrentMessage = "getuuid"
					return elementUnderTest.GetUuid()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetUuid_OUTPUT: GetUuid_OUTPUT{
						XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService", Local: "GetUuid_OUTPUT"},
						UUID:    "E67jVdK/u2EXoIiu3XA36g==",
					},
				},
			},
			//CommitChanges
			{
				"should create a valid AMT_SetupAndConfigurationService CommitChanges wsman message",
				AMT_SetupAndConfigurationService,
				methods.GenerateAction(AMT_SetupAndConfigurationService, CommitChanges),
				`<h:CommitChanges_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"></h:CommitChanges_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "CommitChanges"
					return elementUnderTest.CommitChanges()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					CommitChanges_OUTPUT: CommitChanges_OUTPUT{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService", Local: "CommitChanges_OUTPUT"},
						ReturnValue: 0,
					},
				},
			},
			//SetMEBxPassword
			{
				"should create a valid AMT_SetupAndConfigurationService SetMEBxPassword wsman message",
				AMT_SetupAndConfigurationService,
				methods.GenerateAction(AMT_SetupAndConfigurationService, SetMEBxPassword),
				`<h:SetMEBxPassword_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"><h:Password>P@ssw0rd</h:Password></h:SetMEBxPassword_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "SetMEBxPassword"
					return elementUnderTest.SetMEBXPassword("P@ssw0rd")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SetMEBxPassword_OUTPUT: SetMEBxPassword_OUTPUT{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService", Local: "SetMEBxPassword_OUTPUT"},
						ReturnValue: 0,
					},
				},
			},
			//Unprovision
			{
				"should create a valid AMT_SetupAndConfigurationService Unprovision wsman message",
				AMT_SetupAndConfigurationService,
				methods.GenerateAction(AMT_SetupAndConfigurationService, Unprovision),
				`<h:Unprovision_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"><h:ProvisioningMode>1</h:ProvisioningMode></h:Unprovision_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "Unprovision"
					return elementUnderTest.Unprovision(1)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					Unprovision_OUTPUT: Unprovision_OUTPUT{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService", Local: "Unprovision_OUTPUT"},
						ReturnValue: 0,
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})

	t.Run("decodeUuid Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			responseFunc     func() (string, error, error)
			expectedResponse string
		}{{
			"should properly decode AMT GetUuid Response to a UUID string",
			func() (string, error, error) {
				client.CurrentMessage = "getuuid"
				response, err1 := elementUnderTest.GetUuid()
				uuid, err2 := response.DecodeUUID()
				return uuid, err1, err2
			},
			"55e3ae13-bfd2-61bb-17a0-88aedd7037ea",
		},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				response, err1, err2 := test.responseFunc()
				assert.NoError(t, err1)
				assert.NoError(t, err2)
				assert.Equal(t, test.expectedResponse, response)
			})
		}
	})

}

func TestNegativeAMT_SetupAndConfigurationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/setupandconfiguration",
	}
	elementUnderTest := NewSetupAndConfigurationServiceWithClient(wsmanMessageCreator, &client)
	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			{
				"should create an invalid AMT_SetupAndConfigurationService Pull wsman message",
				"AMT_EthernetPortSettings",
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					response, err := elementUnderTest.Pull("")
					return response, err
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						SetupAndConfigurationServiceItems: []SetupAndConfigurationServiceResponse{
							{
								XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService", Local: "AMT_SetupAndConfigurationService"},
								CreationClassName:             AMT_SetupAndConfigurationService,
								ElementName:                   "Intel(r) AMT Setup and Configuration Service",
								EnabledState:                  5,
								Name:                          "Intel(r) AMT Setup and Configuration Service",
								PasswordModel:                 1,
								ProvisioningMode:              1,
								ProvisioningServerOTP:         "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
								ProvisioningState:             2,
								RequestedState:                12,
								SystemCreationClassName:       "CIM_ComputerSystem",
								SystemName:                    "Intel(r) AMT",
								ZeroTouchConfigurationEnabled: true,
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.NotEqual(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
	t.Run("decodeUuid tests", func(t *testing.T) {
		tests := []struct {
			name             string
			responseFunc     func() (string, error)
			expectedResponse string
		}{{
			"should return error due to bad UUID returned",
			func() (string, error) {
				client.CurrentMessage = "getuuid-bad"
				response, _ := elementUnderTest.GetUuid()
				uuid, err := response.DecodeUUID()
				return uuid, err
			},
			"55e3ae13-bfd2-61bb-17a0-88aedd7037ea",
		},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.NotEqual(t, test.expectedResponse, response)
			})
		}
	})
}
