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
	//"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_RemoteAccessPolicyAppliesToMPS(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	// client := wsmantesting.MockClient{
	// 	PackageUnderTest: "amt/general",
	// }
	client := wsman.NewClient("http://localhost:16992/wsman", "admin", "Intel123!", true)
	
	elementUnderTest := NewRemoteAccessPolicyAppliesToMPSWithClient(wsmanMessageCreator, client)
	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			extraHeader  string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Get wsman message", 
				"AMT_RemoteAccessPolicyAppliesToMPS", 
				wsmantesting.GET, 
				"", 
				"",
				func() (Response, error) {
					//client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					
				},
			},
			//ENUMERATES
			/*{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Enumerate wsman message", 
				"AMT_RemoteAccessPolicyAppliesToMPS", 
				wsmantesting.ENUMERATE, 
				wsmantesting.ENUMERATE_BODY, 
				"", 
				func() (Response, error) {
					//client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{},
			},*/
			//PULLS
			// {"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Pull wsman message", "AMT_RemoteAccessPolicyAppliesToMPS", wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
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
				println(response.XMLOutput)
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
