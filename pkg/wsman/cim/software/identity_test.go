/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package software

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveCIMSoftwareIdentity(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/software/identity",
	}
	elementUnderTest := NewSoftwareIdentityWithClient(wsmanMessageCreator, &client)

	t.Run("cim_SoftwareIdentity Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeaders     string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create and parse a valid cim_SoftwareIdentity Get call",
				CIM_SoftwareIdentity,
				wsmantesting.GET,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">AMTApps</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					selector := Selector{
						Name:  "InstanceID",
						Value: "AMTApps",
					}
					return elementUnderTest.Get(selector)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SoftwareIdentityResponse: SoftwareIdentity{
						XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
						InstanceID:    "AMTApps",
						IsEntity:      true,
						VersionString: "12.0.67",
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_SoftwareIdentity Enumerate call",
				CIM_SoftwareIdentity,
				wsmantesting.ENUMERATE,
				"",
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "E2020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_SoftwareIdentity Pull call",
				CIM_SoftwareIdentity,
				wsmantesting.PULL,
				"",
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						SoftwareIdentityItems: []SoftwareIdentity{
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Flash",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Netstack",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMTApps",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMT",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Sku",
								IsEntity:      true,
								VersionString: "16392",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "VendorID",
								IsEntity:      true,
								VersionString: "8086",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Build Number",
								IsEntity:      true,
								VersionString: "1579",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Recovery Version",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Recovery Build Num",
								IsEntity:      true,
								VersionString: "1579",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Legacy Mode",
								IsEntity:      true,
								VersionString: "False",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMT FW Core Version",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeaders, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
func TestNegativeCIMSoftwareIdentity(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/software/identity",
	}
	elementUnderTest := NewSoftwareIdentityWithClient(wsmanMessageCreator, &client)

	t.Run("cim_SoftwareIdentity Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeaders     string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create and parse a valid cim_SoftwareIdentity Get call",
				CIM_SoftwareIdentity,
				wsmantesting.GET,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">AMTApps</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					selector := Selector{
						Name:  "InstanceID",
						Value: "AMTApps",
					}
					return elementUnderTest.Get(selector)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SoftwareIdentityResponse: SoftwareIdentity{
						XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
						InstanceID:    "AMTApps",
						IsEntity:      true,
						VersionString: "12.0.67",
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_SoftwareIdentity Enumerate call",
				CIM_SoftwareIdentity,
				wsmantesting.ENUMERATE,
				"",
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "E2020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_SoftwareIdentity Pull call",
				CIM_SoftwareIdentity,
				wsmantesting.PULL,
				"",
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						SoftwareIdentityItems: []SoftwareIdentity{
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Flash",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Netstack",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMTApps",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMT",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Sku",
								IsEntity:      true,
								VersionString: "16392",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "VendorID",
								IsEntity:      true,
								VersionString: "8086",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Build Number",
								IsEntity:      true,
								VersionString: "1579",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Recovery Version",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Recovery Build Num",
								IsEntity:      true,
								VersionString: "1579",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Legacy Mode",
								IsEntity:      true,
								VersionString: "False",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMT FW Core Version",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeaders, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
