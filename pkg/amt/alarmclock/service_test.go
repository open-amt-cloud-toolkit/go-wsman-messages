/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_AlarmClockService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := wsman.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewService(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_AlarmClockService Get wsman message", "AMT_AlarmClockService", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_AlarmClockService Enumerate wsman message", "AMT_AlarmClockService", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_AlarmClockService Pull wsman message", "AMT_AlarmClockService", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			//AddAlarm
			{"should create a valid AMT_AlarmClockService AddAlarm wsman message", "AMT_AlarmClockService", string(actions.AddAlarm), `<p:AddAlarm_INPUT xmlns:p="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService"><p:AlarmTemplate><s:InstanceID xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">Instance</s:InstanceID><s:ElementName xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">Alarm instance name</s:ElementName><s:StartTime xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence"><p:Datetime xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">2022-12-31T23:59:00Z</p:Datetime></s:StartTime><s:Interval xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence"><p:Interval xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">P1DT23H59M</p:Interval></s:Interval><s:DeleteOnCompletion xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">true</s:DeleteOnCompletion></p:AlarmTemplate></p:AddAlarm_INPUT>`, func() string {

				startTime := "2022-12-31T23:59:00Z"
				minutes := 59
				hours := 23
				days := 1
				interval := minutes + hours*60 + days*1440

				startTimeFormatted, _ := time.Parse(time.RFC3339, startTime)
				return elementUnderTest.AddAlarm(AlarmClockOccurrence{
					InstanceID:         "Instance",
					StartTime:          startTimeFormatted,
					ElementName:        "Alarm instance name",
					Interval:           interval,
					DeleteOnCompletion: true,
				})
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
