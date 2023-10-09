/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ethernetport

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_EthernetPortSettings(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	//client := MockClient{} // wsman.NewClient("http://localhost:16992/wsman", "admin", "P@ssw0rd", true)
	//elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)
	// enumerationId := ""
	client := wsman.NewClient("http://localhost:16992/wsman", "admin", "Intel123!", true)
	elementUnderTest := NewEthernetPortSettingsWithClient(wsmanMessageCreator, client)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc 	 	func() (Response, error)
			expectedResponse 	interface{}
		}{
			//GETS
			// {
			// 	"should create a valid AMT_EthernetPortSettings Get wsman message", 
			// 	"AMT_EthernetPortSettings", 
			// 	wsmantesting.GET, 
			// 	"", 
			// 	func() (Response, error) {
			// 		//client.CurrentMessage = "Get"
			// 		return elementUnderTest.Get()
			// 	},
			// 	Body{
			// 		XMLName: xml.Name{Space: "", Local: ""},
			// 		EthernetPort: EthernetPort{

			// 		},
			// 	},
			// },

		
			//ENUMERATES
			{
				"should create a valid AMT_EthernetPortSettings Enumerate wsman message", 
				"AMT_EthernetPortSettings", 
				wsmantesting.ENUMERATE, 
				wsmantesting.ENUMERATE_BODY, 
				func() (Response, error) {
					//client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				}, 
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "76000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			// {"should create a valid AMT_EthernetPortSettings Pull wsman message", "AMT_EthernetPortSettings", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
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
