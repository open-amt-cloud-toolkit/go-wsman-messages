/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package system

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveCIMComputerSystemPackage(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/system/packaging",
	}
	elementUnderTest := NewSystemPackageWithClient(wsmanMessageCreator, &client)

	t.Run("cim_ComputerSystemPackage Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_SystemPackaging Enumerate call",
				CIM_SystemPackaging,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "E3020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_SystemPackaging Pull call",
				CIM_SystemPackaging,
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
						SystemPackageItems: []SystemPackage{
							{
								Antecedent: Antecedent{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystemPackage", Local: "Antecedent"},
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: ReferenceParameters{
										XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
										ResourceURI: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chassis",
										SelectorSet: SelectorSet{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selector: []Selector{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "CreationClassName",
													Value:   "CIM_Chassis",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "Tag",
													Value:   "CIM_Chassis",
												},
											},
										},
									},
								},
								Dependent: Dependent{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystemPackage", Local: "Dependent"},
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: ReferenceParameters{
										XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
										ResourceURI: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem",
										SelectorSet: SelectorSet{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selector: []Selector{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "CreationClassName",
													Value:   "CIM_ComputerSystem",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "Name",
													Value:   "ManagedSystem",
												},
											},
										},
									},
								},
								PlatformGUID: "FA4EC1D8F1B1EA11BE6DC6D21730D800",
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
func TestNegativeCIMComputerSystemPackage(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/system/packaging",
	}
	elementUnderTest := NewSystemPackageWithClient(wsmanMessageCreator, &client)

	t.Run("cim_ComputerSystemPackage Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_SystemPackaging Enumerate call",
				CIM_SystemPackaging,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "E3020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_SystemPackaging Pull call",
				CIM_SystemPackaging,
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
						SystemPackageItems: []SystemPackage{
							{
								Antecedent: Antecedent{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystemPackage", Local: "Antecedent"},
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: ReferenceParameters{
										ResourceURI: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chassis",
										SelectorSet: SelectorSet{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selector: []Selector{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "CreationClassName",
													Value:   "CIM_Chassis",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "Tag",
													Value:   "CIM_Chassis",
												},
											},
										},
									},
								},
								Dependent: Dependent{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystemPackage", Local: "Dependent"},
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: ReferenceParameters{
										ResourceURI: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem",
										SelectorSet: SelectorSet{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selector: []Selector{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "CreationClassName",
													Value:   "CIM_ComputerSystem",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "Name",
													Value:   "ManagedSystem",
												},
											},
										},
									},
								},
								PlatformGUID: "FA4EC1D8F1B1EA11BE6DC6D21730D800",
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
