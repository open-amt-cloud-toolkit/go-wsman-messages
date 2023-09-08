/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package optin

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestIPS_OptInService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewOptInService(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid IPS_OptInService Get wsman message", "IPS_OptInService", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid IPS_OptInService Enumerate wsman message", "IPS_OptInService", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid IPS_OptInService Pull wsman message", "IPS_OptInService", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},

			// SEND_OPT_IN_CODE
			{"should create a valid IPS_OptInService send opt in code wsman message", "IPS_OptInService", wsmantesting.SEND_OPT_IN_CODE, `<h:SendOptInCode_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"><h:OptInCode>1</h:OptInCode></h:SendOptInCode_INPUT>`, func() string {
				return elementUnderTest.SendOptInCode(1)
			}},

			// START_OPT_IN
			{"should create a valid IPS_OptInService start opt in code wsman message", "IPS_OptInService", wsmantesting.START_OPT_IN, `<h:StartOptIn_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"></h:StartOptIn_INPUT>`, func() string {
				return elementUnderTest.StartOptIn()
			}},

			// CANCEL_OPT_IN
			{"should create a valid IPS_OptInService cancel opt in code wsman message", "IPS_OptInService", wsmantesting.CANCEL_OPT_IN, `<h:CancelOptIn_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"></h:CancelOptIn_INPUT>`, func() string {
				return elementUnderTest.CancelOptIn()
			}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				correctResponse := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response := test.responseFunc()
				if response != correctResponse {
					assert.Equal(t, correctResponse, response)
				}
			})
		}
	})
}
