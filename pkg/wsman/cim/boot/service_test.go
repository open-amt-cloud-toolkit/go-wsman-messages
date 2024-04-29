/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/boot/service",
	}
	elementUnderTest := NewBootServiceWithClient(wsmanMessageCreator, &client)

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
				"should create and parse a valid cim_BootService Get call",
				CIM_BootService,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ServiceGetResponse: BootService{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService", Local: CIM_BootService},
						Name:                    "Intel(r) AMT Boot Service",
						CreationClassName:       "CIM_BootService",
						SystemName:              "Intel(r) AMT",
						SystemCreationClassName: "CIM_ComputerSystem",
						ElementName:             "Intel(r) AMT Boot Service",
						OperationalStatus:       []OperationalStatus{0},
						EnabledState:            2,
						RequestedState:          12,
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_BootService Enumerate call",
				CIM_BootService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "14000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_BootService Pull call",
				CIM_BootService,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						BootServiceItems: []BootService{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService", Local: CIM_BootService},
								Name:                    "Intel(r) AMT Boot Service",
								CreationClassName:       "CIM_BootService",
								SystemName:              "Intel(r) AMT",
								SystemCreationClassName: "CIM_ComputerSystem",
								ElementName:             "Intel(r) AMT Boot Service",
								OperationalStatus:       []OperationalStatus{0},
								EnabledState:            2,
								RequestedState:          12,
							},
						},
					},
				},
			},
			// SetBootConfigRole
			{
				"should handle error when making cim_BootService SetBootConfigRole wsman message",
				CIM_BootService,
				wsmantesting.SET_BOOT_CONFIG_ROLE,
				`<h:SetBootConfigRole_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService"><h:BootConfigSetting><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="InstanceID"></Selector></SelectorSet></ReferenceParameters></h:BootConfigSetting><h:Role>0</h:Role></h:SetBootConfigRole_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "SetBootConfigRole"
					return elementUnderTest.SetBootConfigRole(BootConfigSetting{}, 0)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SetBootConfigRole_OUTPUT: SetBootConfigRole_OUTPUT{
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

func TestNegativeService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/boot/service",
	}
	elementUnderTest := NewBootServiceWithClient(wsmanMessageCreator, &client)

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
				"should handle error when making cim_BootService Get call",
				CIM_BootService,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ServiceGetResponse: BootService{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService", Local: CIM_BootService},
						Name:                    "Intel(r) AMT Boot Service",
						CreationClassName:       "CIM_BootService",
						SystemName:              "Intel(r) AMT",
						SystemCreationClassName: "CIM_ComputerSystem",
						ElementName:             "Intel(r) AMT Boot Service",
						OperationalStatus:       []OperationalStatus{0},
						EnabledState:            2,
						RequestedState:          12,
					},
				},
			},
			//ENUMERATES
			{
				"should handle error when making cim_BootService Enumerate call",
				CIM_BootService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "14000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should handle error when making cim_BootService Pull wsman message",
				CIM_BootService,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						BootServiceItems: []BootService{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService", Local: CIM_BootService},
								Name:                    "Intel(r) AMT Boot Service",
								CreationClassName:       "CIM_BootService",
								SystemName:              "Intel(r) AMT",
								SystemCreationClassName: "CIM_ComputerSystem",
								ElementName:             "Intel(r) AMT Boot Service",
								OperationalStatus:       []OperationalStatus{0},
								EnabledState:            2,
								RequestedState:          12,
							},
						},
					},
				},
			},
			// SetBootConfigRole
			{
				"should handle error when making cim_BootService SetBootConfigRole wsman message",
				CIM_BootService,
				wsmantesting.SET_BOOT_CONFIG_ROLE,
				`<h:SetBootConfigRole_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService"><h:BootConfigSetting><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="InstanceID"></Selector></SelectorSet></ReferenceParameters></h:BootConfigSetting><h:Role>0</h:Role></h:SetBootConfigRole_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.SetBootConfigRole(BootConfigSetting{}, 0)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SetBootConfigRole_OUTPUT: SetBootConfigRole_OUTPUT{
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
