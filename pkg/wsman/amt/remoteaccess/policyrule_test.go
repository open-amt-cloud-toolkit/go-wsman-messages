/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

type MockClientPolicy struct {
}

const (
	EnvelopeResponsePolicy = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBodyPolicy          = `<g:AMT_AuthorizationService><g:CreationClassName>AMT_RemoteAccessPolicyRule</g:CreationClassName><g:ElementName>Intel(r) AMT Remote Access Policy Rule</g:ElementName><g:Name>Intel(r) AMT Remote Access Policy Rule</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_RemoteAccessPolicyRule>`
)

func (c *MockClientPolicy) Post(msg string) ([]byte, error) {
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/amt/remoteaccess/policyrule/" + strings.ToLower(currentMessage) + ".xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer xmlFile.Close()
	// read file into string
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	// strip carriage returns and new line characters
	xmlData = []byte(strings.ReplaceAll(string(xmlData), "\r\n", ""))

	// Simulate a successful response for testing.
	return []byte(xmlData), nil
}
func TestAMT_RemoteAccessPolicyRule(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	clientPolicy := MockClientPolicy{}
	elementUnderTest := NewPolicyRuleWithClient(wsmanMessageCreator, &clientPolicy)
	elementUnderTest1 := NewRemoteAccessPolicyRule(wsmanMessageCreator)
	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (ResponseRule, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_RemoteAccessPolicyRule Get wsman message",
				"AMT_RemoteAccessPolicyRule",
				wsmantesting.GET,
				"",
				"",
				func() (ResponseRule, error) {
					currentMessage = "Get"
					return elementUnderTest.Get()
				},
				BodyRule{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					RemotePolicyRule: RemotePolicyRule{
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
			//ENUMERATES
			{
				"should create a valid AMT_RemoteAccessPolicyRule Enumerate wsman message",
				"AMT_RemoteAccessPolicyRule",
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (ResponseRule, error) {
					currentMessage = "Enumerate"
					if elementUnderTest1.base.WSManMessageCreator == nil {
						print("Error")
					}
					return elementUnderTest.Enumerate()
				},
				BodyRule{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D3000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_RemoteAccessPolicyRule Pull wsman message", 
				"AMT_RemoteAccessPolicyRule", 
				wsmantesting.PULL, 
				wsmantesting.PULL_BODY, 
				"", 
				func() (ResponseRule, error) {
					currentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				BodyRule{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponseRule: PullResponseRule{
						Items: []ItemRule{
							{
								RemotePolicyRule: RemotePolicyRule{
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
			},

			//DELETE
			//{"should create a valid AMT_RemoteAccessPolicyRule Delete wsman message", "AMT_RemoteAccessPolicyRule", wsmantesting.DELETE, "", "<w:SelectorSet><w:Selector Name=\"PolicyRuleName\">Instance</w:Selector></w:SelectorSet>", func() string {
			//return elementUnderTest.Delete("Instance")
			//}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.BodyRule)
			})
		}
	})
	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (ResponseRule, error)
			expectedResponse interface{}
		}{
			{
				"should create a invalid AMT_RemoteAccessPolicyRule Pull wsman message", 
				"AMT_RemoteAccessPolicyRule", 
				wsmantesting.PULL, 
				wsmantesting.PULL_BODY, 
				"", 
				func() (ResponseRule, error) {
					currentMessage = "Error"
					response, err := elementUnderTest.Pull("")
					return response, err
				},
				BodyRule{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponseRule: PullResponseRule{
						Items: []ItemRule{
							{
								RemotePolicyRule: RemotePolicyRule{
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
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.NotEqual(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.BodyRule)
			})
		}
	})
}
