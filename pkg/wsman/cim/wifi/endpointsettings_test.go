/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveCIMWifiEndpointSettings(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/wifi/endpointsettings",
	}
	elementUnderTest := NewWiFiEndpointSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//ENUMERATES
			{
				"should create and parse a valid cim_WiFiEndpointSettings Enumerate call",
				CIM_WiFiEndpointSettings,
				wsmantesting.ENUMERATE,
				"",
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "95040000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_WiFiEndpointSettings Pull call",
				CIM_WiFiEndpointSettings,
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
						EndpointSettingsItems: []WiFiEndpointSettings{
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings", Local: "CIM_WiFiEndpointSettings"},
								AuthenticationMethod: 2,
								BSSType:              3,
								ElementName:          "Test",
								EncryptionMethod:     2,
								InstanceID:           "Test",
								Priority:             1,
								SSID:                 "testSSID",
							},
						},
					},
				},
			},
			//DELETE
			{
				"should create and parse a valid cim_WiFiEndpointSettings Delete call",
				CIM_WiFiEndpointSettings,
				wsmantesting.DELETE,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">instanceID123</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Delete"
					return elementUnderTest.Delete("instanceID123")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: message.ReturnValue{
						ReturnValue: 0,
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
func TestNegativeCIMWifiEndpointSettings(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/wifi/endpointsettings",
	}
	elementUnderTest := NewWiFiEndpointSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//ENUMERATES
			{
				"should handle error when cim_WiFiEndpointSettings Enumerate call",
				CIM_WiFiEndpointSettings,
				wsmantesting.ENUMERATE,
				"",
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "95040000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should handle error when cim_WiFiEndpointSettings Pull call",
				CIM_WiFiEndpointSettings,
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
						EndpointSettingsItems: []WiFiEndpointSettings{
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings", Local: "CIM_WiFiEndpointSettings"},
								AuthenticationMethod: 2,
								BSSType:              3,
								ElementName:          "Test",
								EncryptionMethod:     2,
								InstanceID:           "Test",
								Priority:             1,
								SSID:                 "testSSID",
							},
						},
					},
				},
			},
			//DELETE
			{
				"should handle error when cim_WiFiEndpointSettings Delete call",
				CIM_WiFiEndpointSettings,
				wsmantesting.DELETE,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">instanceID123</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Delete("instanceID123")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: message.ReturnValue{
						ReturnValue: 0,
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
