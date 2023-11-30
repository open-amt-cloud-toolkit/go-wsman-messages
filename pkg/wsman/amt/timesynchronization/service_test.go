/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestAMT_TimeSynchronizationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewTimeSynchronizationService(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_TimeSynchronizationService Get wsman message", AMT_TimeSynchronizationService, wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_TimeSynchronizationService Enumerate wsman message", AMT_TimeSynchronizationService, wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_TimeSynchronizationService Pull wsman message", AMT_TimeSynchronizationService, wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			{"should create a valid AMT_TimeSynchronizationService GetLowAccuracyTimeSynch wsman message", AMT_TimeSynchronizationService, methods.GenerateAction(AMT_TimeSynchronizationService, GetLowAccuracyTimeSynch), `<h:GetLowAccuracyTimeSynch_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService"></h:GetLowAccuracyTimeSynch_INPUT>`, elementUnderTest.GetLowAccuracyTimeSynch},
			{"should create a valid AMT_TimeSynchronizationService SetHighAccuracyTimeSynch wsman message", AMT_TimeSynchronizationService, methods.GenerateAction(AMT_TimeSynchronizationService, SetHighAccuracyTimeSynch), "<h:SetHighAccuracyTimeSynch_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService\"><h:Ta0>1644240911</h:Ta0><h:Tm1>1644240943</h:Tm1><h:Tm2>1644240943</h:Tm2></h:SetHighAccuracyTimeSynch_INPUT>", func() string { return elementUnderTest.SetHighAccuracyTimeSynch(1644240911, 1644240943, 1644240943) }},
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
