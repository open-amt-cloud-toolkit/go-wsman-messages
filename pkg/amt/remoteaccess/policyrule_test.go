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

func TestAMT_RemoteAccessPolicyRule(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	// client := wsmantesting.MockClient{
	// 	PackageUnderTest: "amt/general",
	// }
	client := wsman.NewClient("http://localhost:16992/wsman", "admin", "Intel123!", true)

	elementUnderTest := NewPolicyRuleWithClient(wsmanMessageCreator, client)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			extraHeader  string
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
					//client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				BodyRule{
					XMLName: xml.Name{Space: "", Local: ""},
					RemotePolicyRule: RemotePolicyRule{
					
					},
				},
			},
			//ENUMERATES
			/*{
				"should create a valid AMT_RemoteAccessPolicyRule Enumerate wsman message", 
				"AMT_RemoteAccessPolicyRule", 
				wsmantesting.ENUMERATE, 
				wsmantesting.ENUMERATE_BODY, 
				"", 
				func() (ResponseRule, error) {
					//client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{},
			},*/
			//PULLS
			//{"should create a valid AMT_RemoteAccessPolicyRule Pull wsman message", "AMT_RemoteAccessPolicyRule", wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
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
				println(response.XMLOutput)
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.BodyRule)
			})
		}
	})
}
