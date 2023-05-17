/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestIPS_AlarmClockOccurrence(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := wsman.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewAlarmClockOccurrence(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			extraHeader  string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid ips_AlarmClockOccurrence Get wsman message", "IPS_AlarmClockOccurrence", wsmantesting.GET, "", "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid IPS_AlarmClockOccurrence Enumerate wsman message", "IPS_AlarmClockOccurrence", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, "", elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid IPS_AlarmClockOccurrence Pull wsman message", "IPS_AlarmClockOccurrence", wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			// DELETE
			{"should create a valid ips_AlarmClockOccurrence Delete wsman message", "IPS_AlarmClockOccurrence", wsmantesting.DELETE, "", "<w:SelectorSet><w:Selector Name=\"Name\">Instance</w:Selector></w:SelectorSet>", func() string {
				return elementUnderTest.Delete("Instance")
			}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				correctResponse := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response := test.responseFunc()
				if response != correctResponse {
					assert.Equal(t, correctResponse, response)
				}
			})
		}
	})
}
