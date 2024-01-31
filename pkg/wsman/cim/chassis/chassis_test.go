/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package chassis

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveCIMChassis(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/chassis",
	}
	elementUnderTest := NewChassisWithClient(wsmanMessageCreator, &client)

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
				"should create and parse a valid cim_Chassis Get call",
				CIM_Chassis,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:            xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chassis", Local: "CIM_Chassis"},
						ChassisPackageType: 0,
						CreationClassName:  "CIM_Chassis",
						ElementName:        "Managed System Chassis",
						Manufacturer:       "Intel(R) Client Systems",
						Model:              "NUC9V7QNX",
						OperationalStatus:  0,
						PackageType:        3,
						SerialNumber:       "JRQN0243007J",
						Tag:                "CIM_Chassis",
						Version:            "K47174-402",
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_Chassis Enumerate call",
				CIM_Chassis,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D1020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_Chassis Pull call",
				CIM_Chassis,
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
						PackageItems: []PackageResponse{
							{
								XMLName:            xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chassis", Local: "CIM_Chassis"},
								ChassisPackageType: 0,
								CreationClassName:  "CIM_Chassis",
								ElementName:        "Managed System Chassis",
								Manufacturer:       "Intel(R) Client Systems",
								Model:              "NUC9V7QNX",
								OperationalStatus:  0,
								PackageType:        3,
								SerialNumber:       "JRQN0243007J",
								Tag:                "CIM_Chassis",
								Version:            "K47174-402",
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

func TestNegativeCIMChassis(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/chassis",
	}
	elementUnderTest := NewChassisWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Chassis Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Chassis Get call",
				CIM_Chassis,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:            xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chassis", Local: "CIM_Chassis"},
						ChassisPackageType: 0,
						CreationClassName:  "CIM_Chassis",
						ElementName:        "Managed System Chassis",
						Manufacturer:       "Intel(R) Client Systems",
						Model:              "NUC9V7QNX",
						OperationalStatus:  0,
						PackageType:        3,
						SerialNumber:       "JRQN0243007J",
						Tag:                "CIM_Chassis",
						Version:            "K47174-402",
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_Chassis Enumerate call",
				CIM_Chassis,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D1020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_Chassis Pull call",
				CIM_Chassis,
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
						PackageItems: []PackageResponse{
							{
								XMLName:            xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chassis", Local: "CIM_Chassis"},
								ChassisPackageType: 0,
								CreationClassName:  "CIM_Chassis",
								ElementName:        "Managed System Chassis",
								Manufacturer:       "Intel(R) Client Systems",
								Model:              "NUC9V7QNX",
								OperationalStatus:  0,
								PackageType:        3,
								SerialNumber:       "JRQN0243007J",
								Tag:                "CIM_Chassis",
								Version:            "K47174-402",
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
