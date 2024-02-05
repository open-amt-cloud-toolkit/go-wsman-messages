/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveAMT_TLSProtocolEndpointCollection(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/tls/protocolendpointcollection",
	}
	elementUnderTest := NewTLSProtocolEndpointCollectionWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TLSProtocolEndpointCollection Tests", func(t *testing.T) {
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
				"should create a valid AMT_TLSProtocolEndpointCollection Get wsman message",
				AMT_TLSProtocolEndpointCollection,
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ProtocolEndpointCollectionGetResponse: ProtocolEndpointCollectionResponse{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSProtocolEndpointCollection), Local: AMT_TLSProtocolEndpointCollection},
						ElementName: "TLSProtocolEndpoint Instances Collection",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_TLSProtocolEndpointCollection Enumerate wsman message",
				AMT_TLSProtocolEndpointCollection,
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
						EnumerationContext: "8B080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_TLSProtocolEndpointCollection Pull wsman message",
				AMT_TLSProtocolEndpointCollection,
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
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						ProtocolEndpointCollectionItems: []ProtocolEndpointCollectionResponse{
							{
								XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSProtocolEndpointCollection), Local: AMT_TLSProtocolEndpointCollection},
								ElementName: "TLSProtocolEndpoint Instances Collection",
							},
						},
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
func TestNegativeAMT_TLSProtocolEndpointCollection(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/tls/protocolendpointcollection",
	}
	elementUnderTest := NewTLSProtocolEndpointCollectionWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TLSProtocolEndpointCollection Tests", func(t *testing.T) {
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
				"should create a valid AMT_TLSProtocolEndpointCollection Get wsman message",
				AMT_TLSProtocolEndpointCollection,
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ProtocolEndpointCollectionGetResponse: ProtocolEndpointCollectionResponse{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSProtocolEndpointCollection), Local: AMT_TLSProtocolEndpointCollection},
						ElementName: "TLSProtocolEndpoint Instances Collection",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_TLSProtocolEndpointCollection Enumerate wsman message",
				AMT_TLSProtocolEndpointCollection,
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
						EnumerationContext: "8B080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_TLSProtocolEndpointCollection Pull wsman message",
				AMT_TLSProtocolEndpointCollection,
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
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						ProtocolEndpointCollectionItems: []ProtocolEndpointCollectionResponse{
							{
								XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_TLSProtocolEndpointCollection), Local: AMT_TLSProtocolEndpointCollection},
								ElementName: "TLSProtocolEndpoint Instances Collection",
							},
						},
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
