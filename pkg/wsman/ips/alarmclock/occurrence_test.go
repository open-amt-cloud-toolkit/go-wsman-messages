/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveIPS_AlarmClockOccurrence(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/alarmclock",
	}
	elementUnderTest := NewAlarmClockOccurrenceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_AlarmClockOccurrence Tests", func(t *testing.T) {
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
				"should create a valid ips_AlarmClockOccurrence Get wsman message",
				"IPS_AlarmClockOccurrence",
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
				"should create a valid IPS_AlarmClockOccurrence Enumerate wsman message",
				"IPS_AlarmClockOccurrence",
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
				"should create a valid IPS_AlarmClockOccurrence Pull wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{},
			},
			// DELETE
			{
				"should create a valid ips_AlarmClockOccurrence Delete wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.DELETE,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">Instance</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Delete"
					return elementUnderTest.Delete("Instance")
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
