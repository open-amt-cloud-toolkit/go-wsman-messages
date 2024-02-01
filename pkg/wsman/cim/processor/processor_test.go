/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package processor

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveCIMProcessor(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/processor",
	}
	elementUnderTest := NewProcessorWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Processor Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Processor Get wsman call",
				CIM_Processor,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Processor", Local: "CIM_Processor"},
						CPUStatus:               1,
						CreationClassName:       "CIM_Processor",
						CurrentClockSpeed:       2400,
						DeviceID:                "CPU 0",
						ElementName:             "Managed System CPU",
						EnabledState:            2,
						ExternalBusClockSpeed:   100,
						Family:                  198,
						HealthState:             0,
						MaxClockSpeed:           8300,
						OperationalStatus:       0,
						OtherFamilyDescription:  "",
						RequestedState:          12,
						Role:                    "Central",
						Stepping:                "13",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
						UpgradeMethod:           52,
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_Processor Enumerate wsman call",
				CIM_Processor,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D9020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_Processor Pull wsman call",
				CIM_Processor,
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
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Processor", Local: "CIM_Processor"},
								CPUStatus:               1,
								CreationClassName:       "CIM_Processor",
								CurrentClockSpeed:       2400,
								DeviceID:                "CPU 0",
								ElementName:             "Managed System CPU",
								EnabledState:            2,
								ExternalBusClockSpeed:   100,
								Family:                  198,
								HealthState:             0,
								MaxClockSpeed:           8300,
								OperationalStatus:       0,
								OtherFamilyDescription:  "",
								RequestedState:          12,
								Role:                    "Central",
								Stepping:                "13",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
								UpgradeMethod:           52,
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
func TestNegativeCIMProcessor(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/processor",
	}
	elementUnderTest := NewProcessorWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Processor Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Processor Get wsman call",
				CIM_Processor,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Processor", Local: "CIM_Processor"},
						CPUStatus:               1,
						CreationClassName:       "CIM_Processor",
						CurrentClockSpeed:       2400,
						DeviceID:                "CPU 0",
						ElementName:             "Managed System CPU",
						EnabledState:            2,
						ExternalBusClockSpeed:   100,
						Family:                  198,
						HealthState:             0,
						MaxClockSpeed:           8300,
						OperationalStatus:       0,
						OtherFamilyDescription:  "",
						RequestedState:          12,
						Role:                    "Central",
						Stepping:                "13",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
						UpgradeMethod:           52,
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_Processor Enumerate wsman call",
				CIM_Processor,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D9020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_Processor Pull wsman call",
				CIM_Processor,
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
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Processor", Local: "CIM_Processor"},
								CPUStatus:               1,
								CreationClassName:       "CIM_Processor",
								CurrentClockSpeed:       2400,
								DeviceID:                "CPU 0",
								ElementName:             "Managed System CPU",
								EnabledState:            2,
								ExternalBusClockSpeed:   100,
								Family:                  198,
								HealthState:             0,
								MaxClockSpeed:           8300,
								OperationalStatus:       0,
								OtherFamilyDescription:  "",
								RequestedState:          12,
								Role:                    "Central",
								Stepping:                "13",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
								UpgradeMethod:           52,
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
