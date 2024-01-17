/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package optin

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveIPS_OptInService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/optin",
	}
	elementUnderTest := NewOptInServiceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_OptInService Tests", func(t *testing.T) {
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
				"should create a valid IPS_OptInService Get wsman message",
				IPS_OptInService,
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{},
			},
			//ENUMERATES
			{
				"should create a valid IPS_OptInService Enumerate wsman message",
				IPS_OptInService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{},
			},
			//PULLS
			{
				"should create a valid IPS_OptInService Pull wsman message",
				IPS_OptInService,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{},
			},

			// SEND_OPT_IN_CODE
			{
				"should create a valid IPS_OptInService send opt in code wsman message",
				IPS_OptInService,
				wsmantesting.SEND_OPT_IN_CODE,
				`<h:SendOptInCode_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"><h:OptInCode>1</h:OptInCode></h:SendOptInCode_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = "SendOptInCode"
					return elementUnderTest.SendOptInCode(0)
				},
				Body{},
			},

			// START_OPT_IN
			{
				"should create a valid IPS_OptInService start opt in code wsman message",
				IPS_OptInService,
				wsmantesting.START_OPT_IN,
				`<h:StartOptIn_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"></h:StartOptIn_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = "StartOptIn"
					return elementUnderTest.StartOptIn()
				},
				Body{},
			},

			// CANCEL_OPT_IN
			{
				"should create a valid IPS_OptInService cancel opt in code wsman message",
				IPS_OptInService,
				wsmantesting.CANCEL_OPT_IN,
				`<h:CancelOptIn_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"></h:CancelOptIn_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = "CancelOptIn"
					return elementUnderTest.CancelOptIn()
				},
				Body{},
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
