/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package power

import (
	"encoding/xml"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
	"github.com/stretchr/testify/assert"
)

const (
	RequestPowerStateChange_BODY = "<h:RequestPowerStateChange_INPUT xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService\"><h:PowerState>8</h:PowerState><h:ManagedElement><Address xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\"><ResourceURI xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem</ResourceURI><SelectorSet xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\"><Selector Name=\"CreationClassName\">CIM_ComputerSystem</Selector><Selector Name=\"Name\">ManagedSystem</Selector></SelectorSet></ReferenceParameters></h:ManagedElement></h:RequestPowerStateChange_INPUT>"
)

func TestPositiveCIMPowerManagementService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/power/managementservice",
	}
	elementUnderTest := NewPowerManagementServiceWithClient(wsmanMessageCreator, &client)

	t.Run("cim_PowerManagementService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			{
				"Should issue a valid cim_PowerManagementService RequestPowerStateChange call",
				CIM_PowerManagementService,
				methods.GenerateAction(CIM_PowerManagementService, RequestPowerStateChange),
				RequestPowerStateChange_BODY,
				func() (Response, error) {
					client.CurrentMessage = "RequestPowerStateChange"
					var powerState PowerState = PowerOffHard
					return elementUnderTest.RequestPowerStateChange(powerState)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestPowerStateChangeResponse: PowerActionResponse{
						ReturnValue: 0,
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Get call",
				CIM_PowerManagementService,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: PowerManagementService{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService", Local: "CIM_PowerManagementService"},
						CreationClassName:       "CIM_PowerManagementService",
						ElementName:             "Intel(r) AMT Power Management Service",
						EnabledState:            5,
						Name:                    "Intel(r) AMT Power Management Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Enumerate call",
				CIM_PowerManagementService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "DB020000-0000-0000-0000-000000000000",
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Pull call",
				CIM_PowerManagementService,
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
						PowerManagementServiceItems: []PowerManagementService{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService", Local: "CIM_PowerManagementService"},
								CreationClassName:       "CIM_PowerManagementService",
								ElementName:             "Intel(r) AMT Power Management Service",
								EnabledState:            5,
								Name:                    "Intel(r) AMT Power Management Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
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

func TestNegativeCIMPowerManagementService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/power/managementservice",
	}
	elementUnderTest := NewPowerManagementServiceWithClient(wsmanMessageCreator, &client)

	t.Run("cim_PowerManagementService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			{
				"Should issue a valid cim_PowerManagementService RequestPowerStateChange call",
				CIM_PowerManagementService,
				methods.GenerateAction(CIM_PowerManagementService, RequestPowerStateChange),
				RequestPowerStateChange_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					var powerState PowerState = PowerOffHard
					return elementUnderTest.RequestPowerStateChange(powerState)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestPowerStateChangeResponse: PowerActionResponse{
						ReturnValue: 0,
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Get call",
				CIM_PowerManagementService,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: PowerManagementService{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService", Local: "CIM_PowerManagementService"},
						CreationClassName:       "CIM_PowerManagementService",
						ElementName:             "Intel(r) AMT Power Management Service",
						EnabledState:            5,
						Name:                    "Intel(r) AMT Power Management Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Enumerate call",
				CIM_PowerManagementService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "DB020000-0000-0000-0000-000000000000",
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Pull call",
				CIM_PowerManagementService,
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
						PowerManagementServiceItems: []PowerManagementService{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService", Local: "CIM_PowerManagementService"},
								CreationClassName:       "CIM_PowerManagementService",
								ElementName:             "Intel(r) AMT Power Management Service",
								EnabledState:            5,
								Name:                    "Intel(r) AMT Power Management Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
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
