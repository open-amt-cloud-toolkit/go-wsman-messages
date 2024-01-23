/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveCIMWifiPort(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/wifi/port",
	}
	elementUnderTest := NewWiFiPortWithClient(wsmanMessageCreator, &client)

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
				"should create and parse a valid cim_WiFiPort Get call",
				CIM_WiFiPort,
				wsmantesting.GET,
				"", func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					WiFiPortGetResponse: WiFiPort{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiPort", Local: "CIM_WiFiPort"},
						CreationClassName:       "CIM_WiFiPort",
						DeviceID:                "WiFi Port 0",
						ElementName:             "WiFi Port 0",
						EnabledState:            3,
						HealthState:             5,
						LinkTechnology:          11,
						PermanentAddress:        "000000000000",
						PortType:                0,
						RequestedState:          3,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_WiFiPort Enumerate call",
				CIM_WiFiPort,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "22000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_WiFiPort Pull call",
				CIM_WiFiPort,
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
						WiFiPortItems: []WiFiPort{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiPort", Local: "CIM_WiFiPort"},
								CreationClassName:       "CIM_WiFiPort",
								DeviceID:                "WiFi Port 0",
								ElementName:             "WiFi Port 0",
								EnabledState:            3,
								HealthState:             5,
								LinkTechnology:          11,
								PermanentAddress:        "000000000000",
								PortType:                0,
								RequestedState:          3,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			{
				"should create and parse a valid cim_WiFiPort Request State Change call",
				CIM_WiFiPort,
				methods.RequestStateChange(CIM_WiFiPort),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiPort\"><h:RequestedState>3</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "RequestStateChange"
					return elementUnderTest.RequestStateChange(3)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: common.ReturnValue{
						XMLName:     xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiPort", Local: "RequestStateChange_OUTPUT"},
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
func TestNegativeCIMWifiPort(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/wifi/port",
	}
	elementUnderTest := NewWiFiPortWithClient(wsmanMessageCreator, &client)

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
				"should create and parse a valid cim_WiFiPort Get call",
				CIM_WiFiPort,
				wsmantesting.GET,
				"", func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					WiFiPortGetResponse: WiFiPort{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiPort", Local: "CIM_WiFiPort"},
						CreationClassName:       "CIM_WiFiPort",
						DeviceID:                "WiFi Port 0",
						ElementName:             "WiFi Port 0",
						EnabledState:            3,
						HealthState:             5,
						LinkTechnology:          11,
						PermanentAddress:        "000000000000",
						PortType:                0,
						RequestedState:          3,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_WiFiPort Enumerate call",
				CIM_WiFiPort,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "22000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_WiFiPort Pull call",
				CIM_WiFiPort,
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
						WiFiPortItems: []WiFiPort{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiPort", Local: "CIM_WiFiPort"},
								CreationClassName:       "CIM_WiFiPort",
								DeviceID:                "WiFi Port 0",
								ElementName:             "WiFi Port 0",
								EnabledState:            3,
								HealthState:             5,
								LinkTechnology:          11,
								PermanentAddress:        "000000000000",
								PortType:                0,
								RequestedState:          3,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			{
				"should create and parse a valid cim_WiFiPort Request State Change call",
				CIM_WiFiPort,
				methods.RequestStateChange(CIM_WiFiPort),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiPort\"><h:RequestedState>3</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.RequestStateChange(3)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: common.ReturnValue{
						XMLName:     xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiPort", Local: "RequestStateChange_OUTPUT"},
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
