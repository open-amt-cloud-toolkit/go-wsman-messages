/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publicprivate

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestAMT_PublicPrivateKeyPair(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewPublicPrivateKeyPair(wsmanMessageCreator)

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
			{"should create a valid AMT_PublicPrivateKeyPair Get wsman message", AMT_PublicPrivateKeyPair, wsmantesting.GET, "", "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_PublicPrivateKeyPair Enumerate wsman message", AMT_PublicPrivateKeyPair, wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, "", elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_PublicPrivateKeyPair Pull wsman message", AMT_PublicPrivateKeyPair, wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			//DELETE
			{"should create a valid AMT_PublicPrivateKeyPair Delete wsman message", AMT_PublicPrivateKeyPair, wsmantesting.DELETE, "", "<w:SelectorSet><w:Selector Name=\"InstanceID\">instanceID123</w:Selector></w:SelectorSet>", func() string { return elementUnderTest.Delete("instanceID123") }},
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
