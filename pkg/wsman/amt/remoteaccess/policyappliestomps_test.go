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

type MockClientApply struct {
}

const (
	EnvelopeResponseApply = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBodyApply          = `<g:AMT_AuthorizationService><g:CreationClassName>AMT_AuthorizationService</g:CreationClassName><g:ElementName>Intel(r) AMT Authorization Service</g:ElementName><g:Name>Intel(r) AMT Alarm Clock Service</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_AuthorizationService>`
)

func (c *MockClientApply) Post(msg string) ([]byte, error) {
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/amt/remoteaccess/policyappliestomps/" + strings.ToLower(currentMessage) + ".xml")
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
func TestAMT_RemoteAccessPolicyAppliesToMPS(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := MockClientApply{}
	elementUnderTest := NewRemoteAccessPolicyAppliesToMPSWithClient(wsmanMessageCreator, &client)
	elementUnderTest1 := NewRemoteAccessPolicyAppliesToMPS(wsmanMessageCreator)
	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (ResponseApplies, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Get wsman message",
				"AMT_RemoteAccessPolicyAppliesToMPS",
				wsmantesting.GET,
				"",
				"",
				func() (ResponseApplies, error) {
					currentMessage = "Get"
					return elementUnderTest.Get()
				},
				BodyApplies{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PolicyApplies: PolicyApplies{
						CreationClassName:       "",
						Name:                    "",
						SystemCreationClassName: "",
						SystemName:              "",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Enumerate wsman message",
				"AMT_RemoteAccessPolicyAppliesToMPS",
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (ResponseApplies, error) {
					currentMessage = "Enumerate"
					if elementUnderTest1.base.WSManMessageCreator == nil {
						print("Error")
					}
					return elementUnderTest.Enumerate()
				},
				BodyApplies{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CE000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Pull wsman message", 
				"AMT_RemoteAccessPolicyAppliesToMPS", 
				wsmantesting.PULL, 
				wsmantesting.PULL_BODY, 
				"", 
				func() (ResponseApplies, error) {
					currentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				BodyApplies{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponseApplies: PullResponseApplies{
						Items: []ItemApplies{
							{
								PolicyApplies: PolicyApplies{
									CreationClassName:       "",
									Name:                    "",
									SystemCreationClassName: "",
									SystemName:              "",
								},
							},
						},
					},
				},
			},
			// {"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Put wsman message", "AMT_RemoteAccessPolicyAppliesToMPS", wsmantesting.PUT, `<h:AMT_RemoteAccessPolicyAppliesToMPS xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS"><h:PolicySet><h:Caption>test</h:Caption><h:Description>test</h:Description><h:ElementName>test</h:ElementName><h:CommonName>test</h:CommonName><h:PolicyKeywords>test</h:PolicyKeywords><h:PolicyDecisionStrategy>1</h:PolicyDecisionStrategy><h:PolicyRoles>test</h:PolicyRoles><h:Enabled>1</h:Enabled></h:PolicySet><h:ManagedElement><h:Caption>test</h:Caption><h:Description>test</h:Description><h:ElementName>test</h:ElementName></h:ManagedElement><h:OrderOfAccess>0</h:OrderOfAccess><h:MpsType>2</h:MpsType></h:AMT_RemoteAccessPolicyAppliesToMPS>`, "", func() string {
			// 	rapatmps := RemoteAccessPolicyAppliesToMPS{
			// 		PolicySetAppliesToElement: PolicySetAppliesToElement{
			// 			ManagedElement: models.ManagedElement{
			// 				Caption:     "test",
			// 				Description: "test",
			// 				ElementName: "test",
			// 			},
			// 			PolicySet: PolicySet{
			// 				Enabled:                1,
			// 				PolicyDecisionStrategy: PolicyDecisionStrategyFirstMatching,
			// 				PolicyRoles:            []string{"test"},
			// 				Policy: Policy{
			// 					ManagedElement: models.ManagedElement{
			// 						Caption:     "test",
			// 						Description: "test",
			// 						ElementName: "test",
			// 					},
			// 					CommonName:     "test",
			// 					PolicyKeywords: []string{"test"},
			// 				},
			// 			},
			// 		},
			// 		MPSType:       BothMPS,
			// 		OrderOfAccess: 0,
			// 	}

			// 	return elementUnderTest.Put(&rapatmps)
			// }},
			// //{"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Create wsman message", "AMT_RemoteAccessPolicyAppliesToMPS", wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			// {"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Delete wsman message", "AMT_RemoteAccessPolicyAppliesToMPS", wsmantesting.DELETE, "", "<w:SelectorSet><w:Selector Name=\"Name\">Instance</w:Selector></w:SelectorSet>", func() string {
			// 	return elementUnderTest.Delete("Instance")
			// }},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.BodyApplies)
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
			responseFunc     func() (ResponseApplies, error)
			expectedResponse interface{}
		}{
			{
				"should create a invalid AMT_RemoteAccessPolicyAppliesToMPS Pull wsman message", 
				"AMT_RemoteAccessPolicyAppliesToMPS", 
				wsmantesting.PULL, 
				wsmantesting.PULL_BODY, 
				"", 
				func() (ResponseApplies, error) {
					currentMessage = "Error"
					response, err := elementUnderTest.Pull("")
					return response, err
				},
				BodyApplies{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponseApplies: PullResponseApplies{
						Items: []ItemApplies{
							{
								PolicyApplies: PolicyApplies{
									CreationClassName:       "",
									Name:                    "",
									SystemCreationClassName: "",
									SystemName:              "",
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
				assert.NotEqual(t, test.expectedResponse, response.BodyApplies)
			})
		}
	})
}
