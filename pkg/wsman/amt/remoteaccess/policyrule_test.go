/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

const (
	EnvelopeResponsePolicy = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBodyPolicy          = `<g:AMT_AuthorizationService><g:CreationClassName>AMT_RemoteAccessPolicyRule</g:CreationClassName><g:ElementName>Intel(r) AMT Remote Access Policy Rule</g:ElementName><g:Name>Intel(r) AMT Remote Access Policy Rule</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_RemoteAccessPolicyRule>`
)

func TestPositiveAMT_RemoteAccessPolicyRule(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/remoteaccess/policyrule",
	}
	elementUnderTest := NewPolicyRuleWithClient(wsmanMessageCreator, &client)
	t.Run("amt_* Tests", func(t *testing.T) {
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
				"should create a valid AMT_RemoteAccessPolicyRule Get wsman message",
				AMT_RemoteAccessPolicyRule,
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RemoteAccessPolicyRuleGetResponse: RemoteAccessPolicyRuleResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule", Local: "AMT_RemoteAccessPolicyRule"},
						CreationClassName:       AMT_RemoteAccessPolicyRule,
						ElementName:             "Inte(r) AMT:Remote Access Policy",
						ExtendedData:            "AAAAAAAAABk=",
						PolicyRuleName:          "Periodic",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
						Trigger:                 2,
						TunnelLifeTime:          0,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_RemoteAccessPolicyRule Enumerate wsman message",
				AMT_RemoteAccessPolicyRule,
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
						EnumerationContext: "D3000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_RemoteAccessPolicyRule Pull wsman message",
				AMT_RemoteAccessPolicyRule,
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
						RemotePolicyRuleItems: []RemoteAccessPolicyRuleResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule", Local: "AMT_RemoteAccessPolicyRule"},
								CreationClassName:       "AMT_RemoteAccessPolicyRule",
								ElementName:             "Inte(r) AMT:Remote Access Policy",
								ExtendedData:            "AAAAAAAAABk=",
								PolicyRuleName:          "Periodic",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
								Trigger:                 2,
								TunnelLifeTime:          0,
							},
						},
					},
				},
			},

			//DELETE
			{
				"should create a valid AMT_RemoteAccessPolicyRule Delete wsman message",
				AMT_RemoteAccessPolicyRule,
				wsmantesting.DELETE,
				"",
				"<w:SelectorSet><w:Selector Name=\"PolicyRuleName\">Instance</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Delete"
					return elementUnderTest.Delete("Instance")
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
func TestNegativeAMT_RemoteAccessPolicyRule(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/remoteaccess/policyrule",
	}
	elementUnderTest := NewPolicyRuleWithClient(wsmanMessageCreator, &client)
	t.Run("amt_* Tests", func(t *testing.T) {
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
				"should create a valid AMT_RemoteAccessPolicyRule Get wsman message",
				AMT_RemoteAccessPolicyRule,
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RemoteAccessPolicyRuleGetResponse: RemoteAccessPolicyRuleResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule", Local: "AMT_RemoteAccessPolicyRule"},
						CreationClassName:       AMT_RemoteAccessPolicyRule,
						ElementName:             "Inte(r) AMT:Remote Access Policy",
						ExtendedData:            "AAAAAAAAABk=",
						PolicyRuleName:          "Periodic",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
						Trigger:                 2,
						TunnelLifeTime:          0,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_RemoteAccessPolicyRule Enumerate wsman message",
				AMT_RemoteAccessPolicyRule,
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
						EnumerationContext: "D3000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_RemoteAccessPolicyRule Pull wsman message",
				AMT_RemoteAccessPolicyRule,
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
						RemotePolicyRuleItems: []RemoteAccessPolicyRuleResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule", Local: "AMT_RemoteAccessPolicyRule"},
								CreationClassName:       "AMT_RemoteAccessPolicyRule",
								ElementName:             "Inte(r) AMT:Remote Access Policy",
								ExtendedData:            "AAAAAAAAABk=",
								PolicyRuleName:          "Periodic",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
								Trigger:                 2,
								TunnelLifeTime:          0,
							},
						},
					},
				},
			},

			//DELETE
			{
				"should create a valid AMT_RemoteAccessPolicyRule Delete wsman message",
				AMT_RemoteAccessPolicyRule,
				wsmantesting.DELETE,
				"",
				"<w:SelectorSet><w:Selector Name=\"PolicyRuleName\">Instance</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Delete("Instance")
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
