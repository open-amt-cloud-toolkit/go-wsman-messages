/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package messagelog

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_MessageLog(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := wsman.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewMessageLog(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_MessageLog Get wsman message", "AMT_MessageLog", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_MessageLog Enumerate wsman message", "AMT_MessageLog", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_MessageLog Pull wsman message", "AMT_MessageLog", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			// MESSAGE LOG
			// POSITION TO FIRST RECORD
			{"should return a valid amt_MessageLog PositionToFirstRecords wsman message", "AMT_MessageLog", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog/PositionToFirstRecord`, `<h:PositionToFirstRecord_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog" />`, func() string {
				return elementUnderTest.PositionToFirstRecord()
			}},
			// GET RECORDS
			{"should return a valid amt_MessageLog GetRecords wsman message", "AMT_MessageLog", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog/GetRecords`, `<h:GetRecords_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog"><h:IterationIdentifier>1</h:IterationIdentifier><h:MaxReadRecords>390</h:MaxReadRecords></h:GetRecords_INPUT>`, func() string {
				return elementUnderTest.GetRecords(1)
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
