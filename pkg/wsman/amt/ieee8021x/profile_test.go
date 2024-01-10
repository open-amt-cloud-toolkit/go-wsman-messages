/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveAMT_8021XProfile(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/ieee8021x/profile",
	}
	elementUnderTest := NewIEEE8021xProfileWithClient(wsmanMessageCreator, &client)

	t.Run("amt_8021XProfile Tests", func(t *testing.T) {
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
				"should create a valid AMT_8021XProfile Get wsman message",
				AMT_IEEE8021xProfile,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ProfileGetAndPutResponse: ProfileResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
						ElementName: "Intel(r) AMT 802.1x Profile",
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
						Enabled:     false,
						ActiveInS0:  true,
						PxeTimeout:  0,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_8021XProfile Enumerate wsman message",
				AMT_IEEE8021xProfile,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "04080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_8021XProfile Pull wsman message",
				AMT_IEEE8021xProfile,
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
						ProfileItems: []ProfileResponse{
							{
								XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
								ElementName: "Intel(r) AMT 802.1x Profile",
								InstanceID:  "Intel(r) AMT 802.1x Profile 0",
								Enabled:     false,
								ActiveInS0:  true,
								PxeTimeout:  0,
							},
						},
					},
				},
			},
			//PUTS
			{
				"should create a valid AMT_8021XProfile Put wsman message",
				AMT_IEEE8021xProfile,
				wsmantesting.PUT,
				"<h:AMT_8021XProfile xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile\"><h:ElementName>Intel(r) AMT 802.1x Profile</h:ElementName><h:InstanceID>Intel(r) AMT 802.1x Profile 0</h:InstanceID><h:Enabled>false</h:Enabled><h:ActiveInS0>false</h:ActiveInS0><h:AuthenticationProtocol>0</h:AuthenticationProtocol><h:ServerCertificateNameComparison>0</h:ServerCertificateNameComparison></h:AMT_8021XProfile>",
				func() (Response, error) {
					profileRequest := ProfileRequest{
						H:           "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile",
						ActiveInS0:  false,
						ElementName: "Intel(r) AMT 802.1x Profile",
						Enabled:     false,
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
					}
					client.CurrentMessage = "Put"
					return elementUnderTest.Put(profileRequest)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ProfileGetAndPutResponse: ProfileResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
						ElementName: "Intel(r) AMT 802.1x Profile",
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
						Enabled:     false,
						ActiveInS0:  false,
						PxeTimeout:  120,
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
func TestNegativeAMT_8021XProfile(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/ieee8021x/profile",
	}
	elementUnderTest := NewIEEE8021xProfileWithClient(wsmanMessageCreator, &client)

	t.Run("amt_8021XProfile Tests", func(t *testing.T) {
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
				"should create a valid AMT_8021XProfile Get wsman message",
				AMT_IEEE8021xProfile,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ProfileGetAndPutResponse: ProfileResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
						ElementName: "Intel(r) AMT 802.1x Profile",
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
						Enabled:     false,
						ActiveInS0:  true,
						PxeTimeout:  0,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_8021XProfile Enumerate wsman message",
				AMT_IEEE8021xProfile,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "04080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_8021XProfile Pull wsman message",
				AMT_IEEE8021xProfile,
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
						ProfileItems: []ProfileResponse{
							{
								XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
								ElementName: "Intel(r) AMT 802.1x Profile",
								InstanceID:  "Intel(r) AMT 802.1x Profile 0",
								Enabled:     false,
								ActiveInS0:  true,
								PxeTimeout:  0,
							},
						},
					},
				},
			},
			//PUTS
			{
				"should create a valid AMT_8021XProfile Put wsman message",
				AMT_IEEE8021xProfile,
				wsmantesting.PUT,
				"<h:AMT_8021XProfile xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile\"><h:ElementName>Intel(r) AMT 802.1x Profile</h:ElementName><h:InstanceID>Intel(r) AMT 802.1x Profile 0</h:InstanceID><h:Enabled>false</h:Enabled><h:ActiveInS0>false</h:ActiveInS0><h:AuthenticationProtocol>0</h:AuthenticationProtocol><h:ServerCertificateNameComparison>0</h:ServerCertificateNameComparison></h:AMT_8021XProfile>",
				func() (Response, error) {
					profileRequest := ProfileRequest{
						H:           "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile",
						ActiveInS0:  false,
						ElementName: "Intel(r) AMT 802.1x Profile",
						Enabled:     false,
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
					}
					client.CurrentMessage = "Error"
					return elementUnderTest.Put(profileRequest)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ProfileGetAndPutResponse: ProfileResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
						ElementName: "Intel(r) AMT 802.1x Profile",
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
						Enabled:     false,
						ActiveInS0:  false,
						PxeTimeout:  120,
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
