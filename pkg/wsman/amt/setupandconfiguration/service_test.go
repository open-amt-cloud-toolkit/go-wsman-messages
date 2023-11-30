/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

type MockClient struct {
}

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_RedirectionService><g:CreationClassName>AMT_RedirectionService</g:CreationClassName><g:ElementName>Intel(r) AMT Redirection Service</g:ElementName><g:Name>Intel(r) AMT Redirection Service</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_RedirectionService`
	GetUuid_BODY     = "<h:GetUuid_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService\"></h:GetUuid_INPUT>"
)

var currentMessage = ""

func (c *MockClient) Post(msg string) ([]byte, error) {
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/amt/setupandconfiguration/" + strings.ToLower(currentMessage) + ".xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer xmlFile.Close()
	// read file into string
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	// strip carriage returns and new line characters
	xmlData = []byte(strings.ReplaceAll(string(xmlData), "\r\n", ""))

	// Simulate a successful response for testing.
	return []byte(xmlData), nil
}
func TestAMT_SetupAndConfigurationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := MockClient{}
	elementUnderTest := NewSetupAndConfigurationServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_* Tests", func(t *testing.T) {
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
					currentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					Setup: Setup{
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
					currentMessage = "Enumerate"
					if elementUnderTest.base.WSManMessageCreator == nil {
						print("Error")
					}
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
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
					currentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponse: PullResponse{
						Items: []Item{
							{
								Setup: Setup{
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
			},
			//GetUuid
			{
				"should return a valid AMT_GetUuid response",
				AMT_SetupAndConfigurationService,
				methods.GenerateAction(AMT_SetupAndConfigurationService, GetUuid),
				GetUuid_BODY,
				func() (Response, error) {
					currentMessage = "getuuid"
					if elementUnderTest.base.WSManMessageCreator == nil {
						print("Error")
					}
					return elementUnderTest.GetUuid()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					GetUuid_OUTPUT: GetUuid_OUTPUT{
						UUID: "E67jVdK/u2EXoIiu3XA36g==",
					},
				},
			},
		}
		// {"should create a valid AMT_SetupAndConfigurationService CommitChanges wsman message", AMT_SetupAndConfigurationService, string(actions.CommitChanges), `<h:CommitChanges_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"></h:CommitChanges_INPUT>`, elementUnderTest.CommitChanges},
		// {"should create a valid AMT_SetupAndConfigurationService SetMEBxPassword wsman message", AMT_SetupAndConfigurationService, string(actions.SetMEBxPassword), `<h:SetMEBxPassword_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"><h:Password>P@ssw0rd</h:Password></h:SetMEBxPassword_INPUT>`, func() string { return elementUnderTest.SetMEBXPassword("P@ssw0rd") }},
		// {"should create a valid AMT_SetupAndConfigurationService Unprovision wsman message", AMT_SetupAndConfigurationService, string(actions.Unprovision), `<h:Unprovision_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"><h:ProvisioningMode>1</h:ProvisioningMode></h:Unprovision_INPUT>`, func() string { return elementUnderTest.Unprovision(1) }},

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
				currentMessage = "getuuid"
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
	client := MockClient{}
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
					currentMessage = "Error"
					response, err := elementUnderTest.Pull("")
					return response, err
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponse: PullResponse{
						Items: []Item{
							{
								Setup: Setup{
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
				currentMessage = "getuuid-bad"
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
