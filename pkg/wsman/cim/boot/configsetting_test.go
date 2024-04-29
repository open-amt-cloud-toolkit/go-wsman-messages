/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			ConfigSettingGetResponse: BootConfigSetting{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ConfigSettingGetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"InstanceID\":\"\",\"ElementName\":\"\"},\"SourceSettingGetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"StructuredBootString\":\"\",\"BIOSBootString\":\"\",\"BootString\":\"\",\"FailThroughSupported\":0},\"ServiceGetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Name\":\"\",\"CreationClassName\":\"\",\"SystemName\":\"\",\"SystemCreationClassName\":\"\",\"ElementName\":\"\",\"OperationalStatus\":null,\"EnabledState\":0,\"RequestedState\":0},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"BootSourceSettingItems\":null,\"BootConfigSettingItems\":null,\"BootServiceItems\":null},\"ChangeBootOrder_OUTPUT\":{\"ReturnValue\":0},\"SetBootConfigRole_OUTPUT\":{\"ReturnValue\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			ConfigSettingGetResponse: BootConfigSetting{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nconfigsettinggetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    instanceid: \"\"\n    elementname: \"\"\nsourcesettinggetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    structuredbootstring: \"\"\n    biosbootstring: \"\"\n    bootstring: \"\"\n    failthroughsupported: 0\nservicegetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    name: \"\"\n    creationclassname: \"\"\n    systemname: \"\"\n    systemcreationclassname: \"\"\n    elementname: \"\"\n    operationalstatus: []\n    enabledstate: 0\n    requestedstate: 0\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    bootsourcesettingitems: []\n    bootconfigsettingitems: []\n    bootserviceitems: []\nchangebootorder_output:\n    returnvalue: 0\nsetbootconfigrole_output:\n    returnvalue: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveConfigSetting(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/boot/configsetting",
	}
	elementUnderTest := NewBootConfigSettingWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_BootConfigSetting Get call",
				CIM_BootConfigSetting,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ConfigSettingGetResponse: BootConfigSetting{
						XMLName:     xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting", Local: CIM_BootConfigSetting},
						InstanceID:  "Intel(r) AMT: Boot Configuration 0",
						ElementName: "Intel(r) AMT: Boot Configuration",
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_BootConfigSetting Enumerate call",
				CIM_BootConfigSetting,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "14000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_BootConfigSetting Pull call",
				CIM_BootConfigSetting,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						BootConfigSettingItems: []BootConfigSetting{
							{
								XMLName:     xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting", Local: CIM_BootConfigSetting},
								InstanceID:  "Intel(r) AMT: Boot Configuration 0",
								ElementName: "Intel(r) AMT: Boot Configuration",
							},
						},
					},
				},
			},
			//Change Boot Order
			{
				"should create and parse a valid cim_BootConfigSetting ChangeBootOrder call",
				CIM_BootConfigSetting,
				methods.GenerateAction(CIM_BootConfigSetting, ChangeBootOrder),
				"<h:ChangeBootOrder_INPUT xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting\"><h:Source><Address xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\"><ResourceURI xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</ResourceURI><SelectorSet xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\"><Selector Name=\"InstanceID\">Intel(r) AMT: Force Hard-drive Boot</Selector></SelectorSet></ReferenceParameters></h:Source></h:ChangeBootOrder_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "ChangeBootOrder"
					return elementUnderTest.ChangeBootOrder(HardDrive)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ChangeBootOrder_OUTPUT: ChangeBootOrder_OUTPUT{
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
}

func TestNegativeConfigSetting(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/boot/configsetting",
	}
	elementUnderTest := NewBootConfigSettingWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
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
				"should handle error when cim_BootConfigSetting Get call",
				CIM_BootConfigSetting,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ConfigSettingGetResponse: BootConfigSetting{
						XMLName:     xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting", Local: CIM_BootConfigSetting},
						InstanceID:  "Intel(r) AMT: Boot Configuration 0",
						ElementName: "Intel(r) AMT: Boot Configuration",
					},
				},
			},
			//ENUMERATES
			{
				"should handle error when cim_BootConfigSetting Enumerate call",
				CIM_BootConfigSetting,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "14000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should handle error when cim_BootConfigSetting Pull call",
				CIM_BootConfigSetting,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						BootConfigSettingItems: []BootConfigSetting{
							{
								XMLName:     xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting", Local: CIM_BootConfigSetting},
								InstanceID:  "Intel(r) AMT: Boot Configuration 0",
								ElementName: "Intel(r) AMT: Boot Configuration",
							},
						},
					},
				},
			},
			//Change Boot Order
			{
				"should handle error when cim_BootConfigSetting ChangeBootOrder call",
				CIM_BootConfigSetting,
				methods.GenerateAction(CIM_BootConfigSetting, ChangeBootOrder),
				"<h:ChangeBootOrder_INPUT xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting\"><h:Source><Address xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\"><ResourceURI xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</ResourceURI><SelectorSet xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\"><Selector Name=\"InstanceID\">Intel(r) AMT: Force Hard-drive Boot</Selector></SelectorSet></ReferenceParameters></h:Source></h:ChangeBootOrder_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.ChangeBootOrder(HardDrive)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ChangeBootOrder_OUTPUT: ChangeBootOrder_OUTPUT{
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
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
