/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package managementpresence

import (
	"encoding/xml"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
	"github.com/stretchr/testify/assert"
)

func TestPositiveAMT_ManagementPresenceRemoteSAP(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/managementpresence",
	}
	elementUnderTest := NewManagementPresenceRemoteSAPWithClient(wsmanMessageCreator, &client)
	t.Run("amt_ManagementPresenceRemoteSAP Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Get wsman message",
				AMT_ManagementPresenceRemoteSAP,
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: ManagementRemoteResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP", Local: "AMT_ManagementPresenceRemoteSAP"},
						AccessInfo:              "192.168.0.208",
						CN:                      "192.168.0.208",
						CreationClassName:       "AMT_ManagementPresenceRemoteSAP",
						ElementName:             "Intel(r) AMT:Management Presence Server",
						InfoFormat:              3,
						Name:                    "Intel(r) AMT:Management Presence Server 0",
						Port:                    4433,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Enumerate wsman message",
				AMT_ManagementPresenceRemoteSAP,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "C9000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Pull wsman message",
				AMT_ManagementPresenceRemoteSAP,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						ManagementRemoteItems: []ManagementRemoteResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP", Local: "AMT_ManagementPresenceRemoteSAP"},
								AccessInfo:              "192.168.10.196",
								CN:                      "192.168.10.196",
								CreationClassName:       "AMT_ManagementPresenceRemoteSAP",
								ElementName:             "Intel(r) AMT:Management Presence Server",
								InfoFormat:              3,
								Name:                    "Intel(r) AMT:Management Presence Server 0",
								Port:                    4433,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			//DELETE
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Delete wsman message",
				AMT_ManagementPresenceRemoteSAP,
				wsmantesting.DELETE,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">instanceID123</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Delete"
					return elementUnderTest.Delete("instanceID123")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
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
func TestNegativeAMT_ManagementPresenceRemoteSAP(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/managementpresence",
	}
	elementUnderTest := NewManagementPresenceRemoteSAPWithClient(wsmanMessageCreator, &client)
	t.Run("amt_ManagementPresenceRemoteSAP Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Get wsman message",
				AMT_ManagementPresenceRemoteSAP,
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: ManagementRemoteResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP", Local: "AMT_ManagementPresenceRemoteSAP"},
						AccessInfo:              "192.168.0.208",
						CN:                      "192.168.0.208",
						CreationClassName:       "AMT_ManagementPresenceRemoteSAP",
						ElementName:             "Intel(r) AMT:Management Presence Server",
						InfoFormat:              3,
						Name:                    "Intel(r) AMT:Management Presence Server 0",
						Port:                    4433,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Enumerate wsman message",
				AMT_ManagementPresenceRemoteSAP,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "C9000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Pull wsman message",
				AMT_ManagementPresenceRemoteSAP,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						ManagementRemoteItems: []ManagementRemoteResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP", Local: "AMT_ManagementPresenceRemoteSAP"},
								AccessInfo:              "192.168.10.196",
								CN:                      "192.168.10.196",
								CreationClassName:       "AMT_ManagementPresenceRemoteSAP",
								ElementName:             "Intel(r) AMT:Management Presence Server",
								InfoFormat:              3,
								Name:                    "Intel(r) AMT:Management Presence Server 0",
								Port:                    4433,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			//DELETE
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Delete wsman message",
				AMT_ManagementPresenceRemoteSAP,
				wsmantesting.DELETE,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">instanceID123</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Delete("instanceID123")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
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
