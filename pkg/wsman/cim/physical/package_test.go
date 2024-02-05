/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/card"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveCIMPackage(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/package",
	}
	elementUnderTest := NewPhysicalPackageWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Package Tests", func(t *testing.T) {
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
				"should create a valid cim_PhysicalPackage Enumerate wsman message",
				CIM_PhysicalPackage,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D8020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid cim_PhysicalPackage Pull wsman message",
				CIM_PhysicalPackage,
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
						EnumerateResponse: common.EnumerateResponse{
							EnumerationContext: "D7020000-0000-0000-0000-000000000000",
						},
						PhysicalPackage: []card.PackageResponse{
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
func TestNegativeCIMPackage(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/package",
	}
	elementUnderTest := NewPhysicalPackageWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Package Tests", func(t *testing.T) {
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
				"should create a valid cim_PhysicalPackage Enumerate wsman message",
				CIM_PhysicalPackage,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D8020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid cim_PhysicalPackage Pull wsman message",
				CIM_PhysicalPackage,
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
						EnumerateResponse: common.EnumerateResponse{
							EnumerationContext: "D7020000-0000-0000-0000-000000000000",
						},
						PhysicalPackage: []card.PackageResponse{
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
