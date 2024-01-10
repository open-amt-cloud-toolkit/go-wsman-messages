/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package credential

import (
	"encoding/xml"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
	"github.com/stretchr/testify/assert"
)

func TestPositiveCIMCredentialContext(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/credential/context",
	}
	elementUnderTest := NewContextWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//ENUMERATES
			{
				"should create and parse a valid cim_IEEE8021xSettings Enumerate call",
				CIM_CredentialContext,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "19000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_IEEE8021xSettings Pull call",
				CIM_CredentialContext,
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
						Items: []CredentialContext{
							{
								ElementInContext: models.AssociationReference{
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: models.ReferenceParmetersNoNamespace{
										ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
										SelectorSet: []models.SelectorNoNamespace{
											{
												Name:  "InstanceID",
												Value: "Intel(r) AMT Certificate: Handle: 0",
											},
										},
									},
								},
								ElementProvidingContext: models.AssociationReference{
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: models.ReferenceParmetersNoNamespace{
										ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSProtocolEndpointCollection",
										SelectorSet: []models.SelectorNoNamespace{
											{
												Name:  "ElementName",
												Value: "TLSProtocolEndpoint Instances Collection",
											},
										},
									},
								},
							},
						},
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

func TestNegativeCIMCredentialContext(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/credential/context",
	}
	elementUnderTest := NewContextWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//ENUMERATES
			{
				"should handle error when cim_IEEE8021xSettings Enumerate call",
				CIM_CredentialContext,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "19000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should handle error when cim_IEEE8021xSettings Pull call",
				CIM_CredentialContext,
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
						Items: []CredentialContext{
							{
								ElementInContext: models.AssociationReference{
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: models.ReferenceParmetersNoNamespace{
										ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
										SelectorSet: []models.SelectorNoNamespace{
											{
												Name:  "InstanceID",
												Value: "Intel(r) AMT Certificate: Handle: 0",
											},
										},
									},
								},
								ElementProvidingContext: models.AssociationReference{
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: models.ReferenceParmetersNoNamespace{
										ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSProtocolEndpointCollection",
										SelectorSet: []models.SelectorNoNamespace{
											{
												Name:  "ElementName",
												Value: "TLSProtocolEndpoint Instances Collection",
											},
										},
									},
								},
							},
						},
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
