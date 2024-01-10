/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package chip

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveCIMChip(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/chip",
	}
	elementUnderTest := NewChipWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Chip Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Chip Get call",
				CIM_Chip, wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chip", Local: "CIM_Chip"},
						CanBeFRUed:        true,
						CreationClassName: "CIM_Chip",
						ElementName:       "Managed System Processor Chip",
						Manufacturer:      "Intel(R) Corporation",
						OperationalStatus: 0,
						Tag:               "CPU 0",
						Version:           "Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz",
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_Chip Enumerate call",
				CIM_Chip, wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D3020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_Chip Pull call",
				CIM_Chip, wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						ChipItems: []PackageResponse{
							{
								XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chip", Local: "CIM_Chip"},
								CanBeFRUed:        true,
								CreationClassName: "CIM_Chip",
								ElementName:       "Managed System Processor Chip",
								Manufacturer:      "Intel(R) Corporation",
								OperationalStatus: 0,
								Tag:               "CPU 0",
								Version:           "Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz",
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
func TestNegativeCIMChip(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/chip",
	}
	elementUnderTest := NewChipWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Chip Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Chip Get call",
				CIM_Chip, wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chip", Local: "CIM_Chip"},
						CanBeFRUed:        true,
						CreationClassName: "CIM_Chip",
						ElementName:       "Managed System Processor Chip",
						Manufacturer:      "Intel(R) Corporation",
						OperationalStatus: 0,
						Tag:               "CPU 0",
						Version:           "Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz",
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_Chip Enumerate call",
				CIM_Chip, wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D3020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_Chip Pull call",
				CIM_Chip, wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						ChipItems: []PackageResponse{
							{
								XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chip", Local: "CIM_Chip"},
								CanBeFRUed:        true,
								CreationClassName: "CIM_Chip",
								ElementName:       "Managed System Processor Chip",
								Manufacturer:      "Intel(R) Corporation",
								OperationalStatus: 0,
								Tag:               "CPU 0",
								Version:           "Intel(R) Core(TM) i7-9850H CPU @ 2.60GHz",
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
