/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestCard(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewCard(wsmanMessageCreator)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid cim_Card Get wsman message", "CIM_Card", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid cim_Card Enumerate wsman message", "CIM_Card", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid cim_Card Pull wsman message", "CIM_Card", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
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
