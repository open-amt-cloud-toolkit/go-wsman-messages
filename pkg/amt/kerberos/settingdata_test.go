/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kerberos

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_KerberosSettingData(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewKerberosSettingData(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_KerberosSettingData Get wsman message", "AMT_KerberosSettingData", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_KerberosSettingData Enumerate wsman message", "AMT_KerberosSettingData", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_KerberosSettingData Pull wsman message", "AMT_KerberosSettingData", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			// KERBEROS SETTING DATA
			// GET CREDENTIAL CACHE STATE
			{"should return a valid amt_KerberosSettingData GetCredentialCacheState wsman message", "AMT_KerberosSettingData", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData/GetCredentialCacheState`, `<h:GetCredentialCacheState_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData"></h:GetCredentialCacheState_INPUT>`, func() string {
				return elementUnderTest.GetCredentialCacheState()
			}},

			// GET CREDENTIAL CACHE STATE
			// {"should return a valid amt_KerberosSettingData SetCredentialCacheState wsman message", "AMT_KerberosSettingData", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData/SetCredentialCacheState`, `<h:SetCredentialCacheState_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_KerberosSettingData"></h:SetCredentialCacheState_INPUT>`, func() string {
			// 	return elementUnderTest.SetCredentialCacheState(true)
			// }},
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
