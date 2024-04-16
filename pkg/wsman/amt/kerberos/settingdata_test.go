/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kerberos

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: KerberosSettingDataResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"RealmName\":\"\",\"ServicePrincipalName\":null,\"ServicePrincipalProtocol\":null,\"KeyVersion\":0,\"EncryptionAlgorithm\":0,\"MasterKey\":null,\"MaximumClockTolerance\":0,\"KrbEnabled\":false,\"Passphrase\":\"\",\"Salt\":\"\",\"IterationCount\":0,\"SupportedEncryptionAlgorithms\":null,\"ConfiguredEncryptionAlgorithms\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"KerberosSettingDataItems\":null},\"GetCredentialCacheState_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Enabled\":false,\"ReturnValue\":0},\"SetCredentialCacheState_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: KerberosSettingDataResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    realmname: \"\"\n    serviceprincipalname: []\n    serviceprincipalprotocol: []\n    keyversion: 0\n    encryptionalgorithm: 0\n    masterkey: []\n    maximumclocktolerance: 0\n    krbenabled: false\n    passphrase: \"\"\n    salt: \"\"\n    iterationcount: 0\n    supportedencryptionalgorithms: []\n    configuredencryptionalgorithms: []\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    kerberossettingdataitems: []\ngetcredentialcachestate_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    enabled: false\n    returnvalue: 0\nsetcredentialcachestate_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_KerberosSettingData(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/kerberos",
	}
	elementUnderTest := NewKerberosSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_KerberosSettingData Tests", func(t *testing.T) {
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
				"should create a valid AMT_KerberosSettingData Get wsman message",
				AMT_KerberosSettingData,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: KerberosSettingDataResponse{
						XMLName:                       xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_KerberosSettingData), Local: AMT_KerberosSettingData},
						ElementName:                   "Intel(r) AMT: Kerberos Settings",
						InstanceID:                    "Intel (r) AMT: Kerberos Settings",
						KrbEnabled:                    false,
						SupportedEncryptionAlgorithms: []SupportedEncryptionAlgorithms{0, 1, 2},
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_KerberosSettingData Enumerate wsman message",
				AMT_KerberosSettingData,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "61080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_KerberosSettingData Pull wsman message",
				AMT_KerberosSettingData,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						KerberosSettingDataItems: []KerberosSettingDataResponse{
							{
								XMLName:                       xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_KerberosSettingData), Local: AMT_KerberosSettingData},
								ElementName:                   "Intel(r) AMT: Kerberos Settings",
								InstanceID:                    "Intel (r) AMT: Kerberos Settings",
								KrbEnabled:                    false,
								SupportedEncryptionAlgorithms: []SupportedEncryptionAlgorithms{0, 1, 2},
							},
						},
					},
				},
			},
			// GET CREDENTIAL CACHE STATE
			{
				"should return a valid amt_KerberosSettingData GetCredentialCacheState wsman message",
				AMT_KerberosSettingData,
				fmt.Sprintf("%s%s/%s", message.AMTSchema, AMT_KerberosSettingData, "GetCredentialCacheState"),
				`<h:GetCredentialCacheState_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData"></h:GetCredentialCacheState_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "GetCredentialCacheState"
					return elementUnderTest.GetCredentialCacheState()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetCredentialCacheState_OUTPUT: GetCredentialCacheState_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_KerberosSettingData), Local: "GetCredentialCacheState_OUTPUT"},
						Enabled:     false,
						ReturnValue: 0,
					},
				},
			},

			// SET CREDENTIAL CACHE STATE
			// {
			// 	"should return a valid amt_KerberosSettingData SetCredentialCacheState wsman message",
			// 	AMT_KerberosSettingData,
			// 	fmt.Sprintf("%s%s/%s", message.AMTSchema, AMT_KerberosSettingData, "SetCredentialCacheState"),
			// 	`<h:SetCredentialCacheState_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData"></h:SetCredentialCacheState_INPUT>`,
			// 	func() (Response, error) {
			// 		client.CurrentMessage = "SetCredentialCacheState"
			// 		return elementUnderTest.SetCredentialCacheState(true)
			// 	},
			// 	Body{},
			// },
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
}
func TestNegativeAMT_KerberosSettingData(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/kerberos",
	}
	elementUnderTest := NewKerberosSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_KerberosSettingData Tests", func(t *testing.T) {
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
				"should create a valid AMT_KerberosSettingData Get wsman message",
				AMT_KerberosSettingData,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: KerberosSettingDataResponse{
						XMLName:                       xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_KerberosSettingData), Local: AMT_KerberosSettingData},
						ElementName:                   "Intel(r) AMT: Kerberos Settings",
						InstanceID:                    "Intel (r) AMT: Kerberos Settings",
						KrbEnabled:                    false,
						SupportedEncryptionAlgorithms: []SupportedEncryptionAlgorithms{0, 1, 2},
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_KerberosSettingData Enumerate wsman message",
				AMT_KerberosSettingData,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "61080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_KerberosSettingData Pull wsman message",
				AMT_KerberosSettingData,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						KerberosSettingDataItems: []KerberosSettingDataResponse{
							{
								XMLName:                       xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_KerberosSettingData), Local: AMT_KerberosSettingData},
								ElementName:                   "Intel(r) AMT: Kerberos Settings",
								InstanceID:                    "Intel (r) AMT: Kerberos Settings",
								KrbEnabled:                    false,
								SupportedEncryptionAlgorithms: []SupportedEncryptionAlgorithms{0, 1, 2},
							},
						},
					},
				},
			},
			// GET CREDENTIAL CACHE STATE
			{
				"should return a valid amt_KerberosSettingData GetCredentialCacheState wsman message",
				AMT_KerberosSettingData,
				fmt.Sprintf("%s%s/%s", message.AMTSchema, AMT_KerberosSettingData, "GetCredentialCacheState"),
				`<h:GetCredentialCacheState_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData"></h:GetCredentialCacheState_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.GetCredentialCacheState()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetCredentialCacheState_OUTPUT: GetCredentialCacheState_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_KerberosSettingData), Local: "GetCredentialCacheState_OUTPUT"},
						Enabled:     false,
						ReturnValue: 0,
					},
				},
			},

			// SET CREDENTIAL CACHE STATE
			// {
			// 	"should return a valid amt_KerberosSettingData SetCredentialCacheState wsman message",
			// 	AMT_KerberosSettingData,
			// 	fmt.Sprintf("%s%s/%s", message.AMTSchema, AMT_KerberosSettingData, "SetCredentialCacheState"),
			// 	`<h:SetCredentialCacheState_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData"></h:SetCredentialCacheState_INPUT>`,
			// 	func() (Response, error) {
			// 		client.CurrentMessage = "SetCredentialCacheState"
			// 		return elementUnderTest.SetCredentialCacheState(true)
			// 	},
			// 	Body{},
			// },
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)

			})
		}
	})
}
