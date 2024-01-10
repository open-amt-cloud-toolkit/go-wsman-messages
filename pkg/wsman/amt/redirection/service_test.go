/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package redirection

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
	"github.com/stretchr/testify/assert"
)

func TestPositiveAMT_RedirectionService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/redirectionservice",
	}
	elementUnderTest := NewRedirectionServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_RedirectionService Tests", func(t *testing.T) {
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
				"should create a valid AMT_RedirectionService Get wsman message",
				AMT_RedirectionService,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: RedirectionResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_RedirectionService), Local: AMT_RedirectionService},
						CreationClassName:       AMT_RedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_RedirectionService Enumerate wsman message",
				AMT_RedirectionService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					if elementUnderTest.base.WSManMessageCreator == nil {
						print("Error")
					}
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D3000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_RedirectionService Pull wsman message",
				AMT_RedirectionService,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						RedirectionItems: []RedirectionResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_RedirectionService), Local: AMT_RedirectionService},
								CreationClassName:       AMT_RedirectionService,
								ElementName:             "Intel(r) AMT Redirection Service",
								EnabledState:            32771,
								ListenerEnabled:         true,
								Name:                    "Intel(r) AMT Redirection Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			//PUTS
			{
				"should create a valid AMT_RedirectionService Put wsman message",
				AMT_RedirectionService,
				wsmantesting.PUT,
				"<h:AMT_RedirectionService xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService\"><h:CreationClassName>AMT_RedirectionService</h:CreationClassName><h:ElementName>Intel(r) AMT Redirection Service</h:ElementName><h:EnabledState>32771</h:EnabledState><h:ListenerEnabled>true</h:ListenerEnabled><h:Name>Intel(r) AMT Redirection Service</h:Name><h:SystemCreationClassName>CIM_ComputerSystem</h:SystemCreationClassName><h:SystemName>Intel(r) AMT</h:SystemName></h:AMT_RedirectionService>",
				func() (Response, error) {
					client.CurrentMessage = "Put"
					redirectionRequest := RedirectionRequest{
						CreationClassName:       AMT_RedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					}
					return elementUnderTest.Put(redirectionRequest)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: RedirectionResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_RedirectionService), Local: AMT_RedirectionService},
						CreationClassName:       AMT_RedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//REQUEST STATE CHANGE
			{
				"should create a valid AMT_RedirectionService Request State Change wsman message",
				AMT_RedirectionService,
				fmt.Sprintf("%s%s/%s", message.AMTSchema, AMT_RedirectionService, "RequestStateChange"),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService\"><h:RequestedState>32771</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "RequestStateChange"
					return elementUnderTest.RequestStateChange(EnableIDERAndSOL)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: RequestStateChange_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_RedirectionService), Local: "RequestStateChange_OUTPUT"},
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
func TestNegativeAMT_RedirectionService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/redirectionservice",
	}
	elementUnderTest := NewRedirectionServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_RedirectionService Tests", func(t *testing.T) {
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
				"should create a valid AMT_RedirectionService Get wsman message",
				AMT_RedirectionService,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: RedirectionResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_RedirectionService), Local: AMT_RedirectionService},
						CreationClassName:       AMT_RedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_RedirectionService Enumerate wsman message",
				AMT_RedirectionService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					if elementUnderTest.base.WSManMessageCreator == nil {
						print("Error")
					}
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D3000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_RedirectionService Pull wsman message",
				AMT_RedirectionService,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						RedirectionItems: []RedirectionResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_RedirectionService), Local: AMT_RedirectionService},
								CreationClassName:       AMT_RedirectionService,
								ElementName:             "Intel(r) AMT Redirection Service",
								EnabledState:            32771,
								ListenerEnabled:         true,
								Name:                    "Intel(r) AMT Redirection Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			//PUTS
			{
				"should create a valid AMT_RedirectionService Put wsman message",
				AMT_RedirectionService,
				wsmantesting.PUT,
				"<h:AMT_RedirectionService xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService\"><h:CreationClassName>AMT_RedirectionService</h:CreationClassName><h:ElementName>Intel(r) AMT Redirection Service</h:ElementName><h:EnabledState>32771</h:EnabledState><h:ListenerEnabled>true</h:ListenerEnabled><h:Name>Intel(r) AMT Redirection Service</h:Name><h:SystemCreationClassName>CIM_ComputerSystem</h:SystemCreationClassName><h:SystemName>Intel(r) AMT</h:SystemName></h:AMT_RedirectionService>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					redirectionRequest := RedirectionRequest{
						CreationClassName:       AMT_RedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					}
					return elementUnderTest.Put(redirectionRequest)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: RedirectionResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_RedirectionService), Local: AMT_RedirectionService},
						CreationClassName:       AMT_RedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//REQUEST STATE CHANGE
			{
				"should create a valid AMT_RedirectionService Request State Change wsman message",
				AMT_RedirectionService,
				fmt.Sprintf("%s%s/%s", message.AMTSchema, AMT_RedirectionService, "RequestStateChange"),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService\"><h:RequestedState>32771</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.RequestStateChange(EnableIDERAndSOL)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: RequestStateChange_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_RedirectionService), Local: "RequestStateChange_OUTPUT"},
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
