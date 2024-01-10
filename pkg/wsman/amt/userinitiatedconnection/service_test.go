/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveAMT_UserInitiatedConnectionService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/userinitiatedconnection",
	}
	elementUnderTest := NewUserInitiatedConnectionServiceWithClient(wsmanMessageCreator, &client)
	t.Run("amt_UserInitiatedConnectionService Tests", func(t *testing.T) {
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
				"should create a valid AMT_UserInitiatedConnectionService Get wsman message",
				AMT_UserInitiatedConnectionService,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: UserResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_UserInitiatedConnectionService), Local: AMT_UserInitiatedConnectionService},
						CreationClassName:       AMT_UserInitiatedConnectionService,
						ElementName:             "Intel(r) AMT User Initiated Connection Service",
						EnabledState:            32771,
						Name:                    "Intel(r) AMT User Initiated Connection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_UserInitiatedConnectionService Enumerate wsman message",
				AMT_UserInitiatedConnectionService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
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
				"should create a valid AMT_UserInitiatedConnectionService Pull wsman message",
				AMT_UserInitiatedConnectionService,
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
						UserItems: []UserResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_UserInitiatedConnectionService), Local: AMT_UserInitiatedConnectionService},
								CreationClassName:       AMT_UserInitiatedConnectionService,
								ElementName:             "Intel(r) AMT User Initiated Connection Service",
								EnabledState:            32771,
								Name:                    "Intel(r) AMT User Initiated Connection Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			//REQUEST STATE CHANGE
			{
				"should create a valid AMT_UserInitiatedConnectionService RequestStateChange wsman message",
				AMT_UserInitiatedConnectionService,
				fmt.Sprintf("%s%s/%s", message.AMTSchema, AMT_UserInitiatedConnectionService, "RequestStateChange"),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService\"><h:RequestedState>32771</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "Request"
					return elementUnderTest.RequestStateChange(32771)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: RequestStateChange_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_UserInitiatedConnectionService), Local: "RequestStateChange_OUTPUT"},
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
func TestNegativeAMT_UserInitiatedConnectionService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/userinitiatedconnection",
	}
	elementUnderTest := NewUserInitiatedConnectionServiceWithClient(wsmanMessageCreator, &client)
	t.Run("amt_UserInitiatedConnectionService Tests", func(t *testing.T) {
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
				"should create a valid AMT_UserInitiatedConnectionService Get wsman message",
				AMT_UserInitiatedConnectionService,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: UserResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_UserInitiatedConnectionService), Local: AMT_UserInitiatedConnectionService},
						CreationClassName:       AMT_UserInitiatedConnectionService,
						ElementName:             "Intel(r) AMT User Initiated Connection Service",
						EnabledState:            32771,
						Name:                    "Intel(r) AMT User Initiated Connection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_UserInitiatedConnectionService Enumerate wsman message",
				AMT_UserInitiatedConnectionService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
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
				"should create a valid AMT_UserInitiatedConnectionService Pull wsman message",
				AMT_UserInitiatedConnectionService,
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
						UserItems: []UserResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_UserInitiatedConnectionService), Local: AMT_UserInitiatedConnectionService},
								CreationClassName:       AMT_UserInitiatedConnectionService,
								ElementName:             "Intel(r) AMT User Initiated Connection Service",
								EnabledState:            32771,
								Name:                    "Intel(r) AMT User Initiated Connection Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			//REQUEST STATE CHANGE
			{
				"should create a valid AMT_UserInitiatedConnectionService RequestStateChange wsman message",
				AMT_UserInitiatedConnectionService,
				fmt.Sprintf("%s%s/%s", message.AMTSchema, AMT_UserInitiatedConnectionService, "RequestStateChange"),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService\"><h:RequestedState>32771</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.RequestStateChange(32771)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: RequestStateChange_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_UserInitiatedConnectionService), Local: "RequestStateChange_OUTPUT"},
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
