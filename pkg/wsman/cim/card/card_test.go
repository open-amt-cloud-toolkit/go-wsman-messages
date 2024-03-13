/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package card

import (
	"encoding/xml"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CardItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PackageResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CanBeFRUed\":false,\"CreationClassName\":\"\",\"ElementName\":\"\",\"Manufacturer\":\"\",\"Model\":\"\",\"OperationalStatus\":0,\"PackageType\":0,\"SerialNumber\":\"\",\"Tag\":\"\",\"Version\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    carditems: []\nenumerateresponse:\n    enumerationcontext: \"\"\npackageresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    canbefrued: false\n    creationclassname: \"\"\n    elementname: \"\"\n    manufacturer: \"\"\n    model: \"\"\n    operationalstatus: 0\n    packagetype: 0\n    serialnumber: \"\"\n    tag: \"\"\n    version: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMCard(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/card",
	}
	elementUnderTest := NewCardWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Card Tests", func(t *testing.T) {
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
				"should create a valid cim_Card Get wsman message",
				CIM_Card,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Card", Local: "CIM_Card"},
						CanBeFRUed:        true,
						CreationClassName: "CIM_Card",
						ElementName:       "Managed System Base Board",
						Manufacturer:      "Intel Corporation",
						Model:             "NUC9V7QNB",
						OperationalStatus: 0,
						PackageType:       9,
						SerialNumber:      "KNQN0221020W",
						Tag:               "CIM_Card",
						Version:           "K47180-402",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid cim_Card Enumerate wsman message",
				CIM_Card,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CF020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid cim_Card Pull wsman message",
				CIM_Card,
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
						CardItems: []PackageResponse{
							{
								XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Card", Local: "CIM_Card"},
								CanBeFRUed:        true,
								CreationClassName: "CIM_Card",
								ElementName:       "Managed System Base Board",
								Manufacturer:      "Intel Corporation",
								Model:             "NUC9V7QNB",
								OperationalStatus: 0,
								PackageType:       9,
								SerialNumber:      "KNQN0221020W",
								Tag:               "CIM_Card",
								Version:           "K47180-402",
							},
						},
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

func TestNegativeCIMCard(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/card",
	}
	elementUnderTest := NewCardWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Card* Tests", func(t *testing.T) {
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
				"should create a valid cim_Card Get wsman message",
				CIM_Card,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Card", Local: "CIM_Card"},
						CanBeFRUed:        true,
						CreationClassName: "CIM_Card",
						ElementName:       "Managed System Base Board",
						Manufacturer:      "Intel Corporation",
						Model:             "NUC9V7QNB",
						OperationalStatus: 0,
						PackageType:       9,
						SerialNumber:      "KNQN0221020W",
						Tag:               "CIM_Card",
						Version:           "K47180-402",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid cim_Card Enumerate wsman message",
				CIM_Card,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CF020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid cim_Card Pull wsman message",
				CIM_Card,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						CardItems: []PackageResponse{
							{
								XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Card", Local: "CIM_Card"},
								CanBeFRUed:        true,
								CreationClassName: "CIM_Card",
								ElementName:       "Managed System Base Board",
								Manufacturer:      "Intel Corporation",
								Model:             "NUC9V7QNB",
								OperationalStatus: 0,
								PackageType:       9,
								SerialNumber:      "KNQN0221020W",
								Tag:               "CIM_Card",
								Version:           "K47180-402",
							},
						},
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
