/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
	"github.com/stretchr/testify/assert"
)

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_TLSSettingData><g:CreationClassName>AMT_TLSSettingData</g:CreationClassName><g:ElementName>Intel(r) TLS Setting Data</g:ElementName><g:Name>Intel(r) AMT TLS Setting Data</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_TLSSettingData>`
)

func TestPositiveAMT_TLSSettingData(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/tls/settingdata",
	}
	elementUnderTest := NewTLSSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TLSSettingData Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_TLSSettingData Get wsman message",
				AMT_TLSSettingData,
				wsmantesting.GET,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT 802.3 TLS Settings</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get("Intel(r) AMT 802.3 TLS Settings")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SettingDataGetAndPutResponse: SettingDataResponse{
						XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSSettingData), Local: AMT_TLSSettingData},
						AcceptNonSecureConnections: false,
						ElementName:                "Intel(r) AMT 802.3 TLS Settings",
						Enabled:                    false,
						InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
						MutualAuthentication:       false,
					},
				},
			},

			//ENUMERATES
			{
				"should create a valid AMT_TLSSettingData Enumerate wsman message",
				AMT_TLSSettingData,
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
						EnumerationContext: "CA000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_TLSSettingData Pull wsman message",
				AMT_TLSSettingData,
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
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						SettingDataItems: []SettingDataResponse{
							{
								XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSSettingData), Local: AMT_TLSSettingData},
								AcceptNonSecureConnections: false,
								ElementName:                "Intel(r) AMT 802.3 TLS Settings",
								Enabled:                    false,
								InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
								MutualAuthentication:       false,
							},
							{
								XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSSettingData), Local: AMT_TLSSettingData},
								AcceptNonSecureConnections: true,
								ElementName:                "Intel(r) AMT LMS TLS Settings",
								Enabled:                    false,
								InstanceID:                 "Intel(r) AMT LMS TLS Settings",
								MutualAuthentication:       false,
							},
						},
					},
				},
			},

			//PUTS
			{
				"should create a valid AMT_TLSSettingData Put wsman message",
				AMT_TLSSettingData,
				wsmantesting.PUT,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT 802.3 TLS Settings</w:Selector></w:SelectorSet>",
				"<h:AMT_TLSSettingData xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData\"><h:ElementName>Intel(r) AMT 802.3 TLS Settings</h:ElementName><h:InstanceID>Intel(r) AMT 802.3 TLS Settings</h:InstanceID><h:MutualAuthentication>false</h:MutualAuthentication><h:Enabled>true</h:Enabled><h:AcceptNonSecureConnections>false</h:AcceptNonSecureConnections><h:NonSecureConnectionsSupported>false</h:NonSecureConnectionsSupported></h:AMT_TLSSettingData>",
				func() (Response, error) {
					client.CurrentMessage = "Put"
					tlsSettingData := SettingDataRequest{
						ElementName: "Intel(r) AMT 802.3 TLS Settings",
						InstanceID:  "Intel(r) AMT 802.3 TLS Settings",
						Enabled:     true,
					}
					return elementUnderTest.Put("Intel(r) AMT 802.3 TLS Settings", tlsSettingData)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SettingDataGetAndPutResponse: SettingDataResponse{
						XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSSettingData), Local: AMT_TLSSettingData},
						AcceptNonSecureConnections: false,
						ElementName:                "Intel(r) AMT 802.3 TLS Settings",
						Enabled:                    false,
						InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
						MutualAuthentication:       false,
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
func TestNegativeAMT_TLSSettingData(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/tls/settingdata",
	}
	elementUnderTest := NewTLSSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TLSSettingData Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_TLSSettingData Get wsman message",
				AMT_TLSSettingData,
				wsmantesting.GET,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT 802.3 TLS Settings</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get("Intel(r) AMT 802.3 TLS Settings")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SettingDataGetAndPutResponse: SettingDataResponse{
						XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSSettingData), Local: AMT_TLSSettingData},
						AcceptNonSecureConnections: false,
						ElementName:                "Intel(r) AMT 802.3 TLS Settings",
						Enabled:                    false,
						InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
						MutualAuthentication:       false,
					},
				},
			},

			//ENUMERATES
			{
				"should create a valid AMT_TLSSettingData Enumerate wsman message",
				AMT_TLSSettingData,
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
						EnumerationContext: "CA000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_TLSSettingData Pull wsman message",
				AMT_TLSSettingData,
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
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						SettingDataItems: []SettingDataResponse{
							{
								XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSSettingData), Local: AMT_TLSSettingData},
								AcceptNonSecureConnections: false,
								ElementName:                "Intel(r) AMT 802.3 TLS Settings",
								Enabled:                    false,
								InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
								MutualAuthentication:       false,
							},
							{
								XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSSettingData), Local: AMT_TLSSettingData},
								AcceptNonSecureConnections: true,
								ElementName:                "Intel(r) AMT LMS TLS Settings",
								Enabled:                    false,
								InstanceID:                 "Intel(r) AMT LMS TLS Settings",
								MutualAuthentication:       false,
							},
						},
					},
				},
			},

			//PUTS
			{
				"should create a valid AMT_TLSSettingData Put wsman message",
				AMT_TLSSettingData,
				wsmantesting.PUT,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT 802.3 TLS Settings</w:Selector></w:SelectorSet>",
				"<h:AMT_TLSSettingData xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData\"><h:ElementName>Intel(r) AMT 802.3 TLS Settings</h:ElementName><h:InstanceID>Intel(r) AMT 802.3 TLS Settings</h:InstanceID><h:MutualAuthentication>false</h:MutualAuthentication><h:Enabled>true</h:Enabled><h:AcceptNonSecureConnections>false</h:AcceptNonSecureConnections><h:NonSecureConnectionsSupported>false</h:NonSecureConnectionsSupported></h:AMT_TLSSettingData>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					tlsSettingData := SettingDataRequest{
						ElementName: "Intel(r) AMT 802.3 TLS Settings",
						InstanceID:  "Intel(r) AMT 802.3 TLS Settings",
						Enabled:     true,
					}
					return elementUnderTest.Put("Intel(r) AMT 802.3 TLS Settings", tlsSettingData)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SettingDataGetAndPutResponse: SettingDataResponse{
						XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSSettingData), Local: AMT_TLSSettingData},
						AcceptNonSecureConnections: false,
						ElementName:                "Intel(r) AMT 802.3 TLS Settings",
						Enabled:                    false,
						InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
						MutualAuthentication:       false,
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
