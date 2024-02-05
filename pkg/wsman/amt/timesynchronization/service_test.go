/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveAMT_TimeSynchronizationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/timesynchronization",
	}
	elementUnderTest := NewTimeSynchronizationServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TimeSynchronizationService Tests", func(t *testing.T) {
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
				"should create a valid AMT_TimeSynchronizationService Get wsman message",
				AMT_TimeSynchronizationService,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: TimeSynchronizationServiceResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "AMT_TimeSynchronizationService"},
						CreationClassName:       "AMT_TimeSynchronizationService",
						ElementName:             "Intel(r) AMT Time Synchronization Service",
						EnabledState:            5,
						LocalTimeSyncEnabled:    0,
						Name:                    "Intel(r) AMT Time Synchronization Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
						TimeSource:              1,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_TimeSynchronizationService Enumerate wsman message",
				AMT_TimeSynchronizationService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "3B080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_TimeSynchronizationService Pull wsman message",
				AMT_TimeSynchronizationService,
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
						TimeSynchronizationServiceItems: []TimeSynchronizationServiceResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "AMT_TimeSynchronizationService"},
								CreationClassName:       "AMT_TimeSynchronizationService",
								ElementName:             "Intel(r) AMT Time Synchronization Service",
								EnabledState:            5,
								LocalTimeSyncEnabled:    0,
								Name:                    "Intel(r) AMT Time Synchronization Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
								TimeSource:              1,
							},
						},
					},
				},
			},
			{
				//GetLowAccuracyTimeSynch
				"should create a valid AMT_TimeSynchronizationService GetLowAccuracyTimeSynch wsman message",
				AMT_TimeSynchronizationService,
				methods.GenerateAction(AMT_TimeSynchronizationService, GetLowAccuracyTimeSynch),
				`<h:GetLowAccuracyTimeSynch_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService"></h:GetLowAccuracyTimeSynch_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "GetLowAccuracyTimeSynch"
					return elementUnderTest.GetLowAccuracyTimeSynch()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetLowAccuracyTimeSynchResponse: GetLowAccuracyTimeSynchResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "GetLowAccuracyTimeSynch_OUTPUT"},
						Ta0:         1704586865,
						ReturnValue: 0,
					},
				},
			},
			{
				//SetHighAccuracyTimeSynch
				"should create a valid AMT_TimeSynchronizationService SetHighAccuracyTimeSynch wsman message",
				AMT_TimeSynchronizationService,
				methods.GenerateAction(AMT_TimeSynchronizationService, SetHighAccuracyTimeSynch),
				"<h:SetHighAccuracyTimeSynch_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService\"><h:Ta0>1644240911</h:Ta0><h:Tm1>1644240943</h:Tm1><h:Tm2>1644240943</h:Tm2></h:SetHighAccuracyTimeSynch_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "SetHighAccuracyTimeSynch"
					return elementUnderTest.SetHighAccuracyTimeSynch(1644240911, 1644240943, 1644240943)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SetHighAccuracyTimeSynchResponse: SetHighAccuracyTimeSynchResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "SetHighAccuracyTimeSynch_OUTPUT"},
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
func TestNegativeAMT_TimeSynchronizationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/timesynchronization",
	}
	elementUnderTest := NewTimeSynchronizationServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TimeSynchronizationService Tests", func(t *testing.T) {
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
				"should create a valid AMT_TimeSynchronizationService Get wsman message",
				AMT_TimeSynchronizationService,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: TimeSynchronizationServiceResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "AMT_TimeSynchronizationService"},
						CreationClassName:       "AMT_TimeSynchronizationService",
						ElementName:             "Intel(r) AMT Time Synchronization Service",
						EnabledState:            5,
						LocalTimeSyncEnabled:    0,
						Name:                    "Intel(r) AMT Time Synchronization Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
						TimeSource:              1,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_TimeSynchronizationService Enumerate wsman message",
				AMT_TimeSynchronizationService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "3B080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_TimeSynchronizationService Pull wsman message",
				AMT_TimeSynchronizationService,
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
						TimeSynchronizationServiceItems: []TimeSynchronizationServiceResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "AMT_TimeSynchronizationService"},
								CreationClassName:       "AMT_TimeSynchronizationService",
								ElementName:             "Intel(r) AMT Time Synchronization Service",
								EnabledState:            5,
								LocalTimeSyncEnabled:    0,
								Name:                    "Intel(r) AMT Time Synchronization Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
								TimeSource:              1,
							},
						},
					},
				},
			},
			{
				//GetLowAccuracyTimeSynch
				"should create a valid AMT_TimeSynchronizationService GetLowAccuracyTimeSynch wsman message",
				AMT_TimeSynchronizationService,
				methods.GenerateAction(AMT_TimeSynchronizationService, GetLowAccuracyTimeSynch),
				`<h:GetLowAccuracyTimeSynch_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService"></h:GetLowAccuracyTimeSynch_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.GetLowAccuracyTimeSynch()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetLowAccuracyTimeSynchResponse: GetLowAccuracyTimeSynchResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "GetLowAccuracyTimeSynch_OUTPUT"},
						Ta0:         1704586865,
						ReturnValue: 0,
					},
				},
			},
			{
				//SetHighAccuracyTimeSynch
				"should create a valid AMT_TimeSynchronizationService SetHighAccuracyTimeSynch wsman message",
				AMT_TimeSynchronizationService,
				methods.GenerateAction(AMT_TimeSynchronizationService, SetHighAccuracyTimeSynch),
				"<h:SetHighAccuracyTimeSynch_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService\"><h:Ta0>1644240911</h:Ta0><h:Tm1>1644240943</h:Tm1><h:Tm2>1644240943</h:Tm2></h:SetHighAccuracyTimeSynch_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.SetHighAccuracyTimeSynch(1644240911, 1644240943, 1644240943)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SetHighAccuracyTimeSynchResponse: SetHighAccuracyTimeSynchResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "SetHighAccuracyTimeSynch_OUTPUT"},
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
