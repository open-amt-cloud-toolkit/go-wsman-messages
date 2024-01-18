/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveCIMMemory(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/memory",
	}
	elementUnderTest := NewPhysicalMemoryWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Memory Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//ENUMERATES
			{
				"should create and parse a valid cim_PhysicalMemory Enumerate call",
				"CIM_PhysicalMemory",
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D6020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_PhysicalMemory Pull call",
				"CIM_PhysicalMemory",
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
						MemoryItems: []PhysicalMemory{
							{
								XMLName:                    xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PhysicalMemory", Local: "CIM_PhysicalMemory"},
								BankLabel:                  "BANK 0",
								Capacity:                   17179869184,
								ConfiguredMemoryClockSpeed: 2400,
								CreationClassName:          "CIM_PhysicalMemory",
								ElementName:                "Managed System Memory Chip",
								FormFactor:                 13,
								IsSpeedInMhz:               true,
								Manufacturer:               "859B",
								MaxMemorySpeed:             2400,
								MemoryType:                 26,
								PartNumber:                 "CT16G4SFD824A.M16FE ",
								SerialNumber:               "E0E8D670",
								Speed:                      0,
								Tag:                        "9876543210",
							},
							{
								XMLName:                    xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PhysicalMemory", Local: "CIM_PhysicalMemory"},
								BankLabel:                  "BANK 2",
								Capacity:                   17179869184,
								ConfiguredMemoryClockSpeed: 2400,
								CreationClassName:          "CIM_PhysicalMemory",
								ElementName:                "Managed System Memory Chip",
								FormFactor:                 13,
								IsSpeedInMhz:               true,
								Manufacturer:               "859B",
								MaxMemorySpeed:             2400,
								MemoryType:                 26,
								PartNumber:                 "CT16G4SFD824A.M16FE ",
								SerialNumber:               "E0E8D070",
								Speed:                      0,
								Tag:                        "9876543210 (#1)",
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
func TestNegativeCIMMemory(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/memory",
	}
	elementUnderTest := NewPhysicalMemoryWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Memory Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//ENUMERATES
			{
				"should create and parse a valid cim_PhysicalMemory Enumerate call",
				"CIM_PhysicalMemory",
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D6020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_PhysicalMemory Pull call",
				"CIM_PhysicalMemory",
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
						MemoryItems: []PhysicalMemory{
							{
								XMLName:                    xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PhysicalMemory", Local: "CIM_PhysicalMemory"},
								BankLabel:                  "BANK 0",
								Capacity:                   17179869184,
								ConfiguredMemoryClockSpeed: 2400,
								CreationClassName:          "CIM_PhysicalMemory",
								ElementName:                "Managed System Memory Chip",
								FormFactor:                 13,
								IsSpeedInMhz:               true,
								Manufacturer:               "859B",
								MaxMemorySpeed:             2400,
								MemoryType:                 26,
								PartNumber:                 "CT16G4SFD824A.M16FE ",
								SerialNumber:               "E0E8D670",
								Speed:                      0,
								Tag:                        "9876543210",
							},
							{
								XMLName:                    xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PhysicalMemory", Local: "CIM_PhysicalMemory"},
								BankLabel:                  "BANK 2",
								Capacity:                   17179869184,
								ConfiguredMemoryClockSpeed: 2400,
								CreationClassName:          "CIM_PhysicalMemory",
								ElementName:                "Managed System Memory Chip",
								FormFactor:                 13,
								IsSpeedInMhz:               true,
								Manufacturer:               "859B",
								MaxMemorySpeed:             2400,
								MemoryType:                 26,
								PartNumber:                 "CT16G4SFD824A.M16FE ",
								SerialNumber:               "E0E8D070",
								Speed:                      0,
								Tag:                        "9876543210 (#1)",
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
