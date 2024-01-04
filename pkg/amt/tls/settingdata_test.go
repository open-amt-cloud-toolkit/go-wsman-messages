/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_TLSSettingData(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewTLSSettingData(wsmanMessageCreator)
	data := TLSSettingData{
		ElementName:                `Intel(r) AMT LMS TLS Settings`,
		InstanceID:                 `Intel(r) AMT LMS TLS Settings`,
		AcceptNonSecureConnections: true,
		Enabled:                    true,
		MutualAuthentication:       true,
	}
	expectedPutSelector := `<w:SelectorSet><w:Selector Name="InstanceID">Intel(r) AMT LMS TLS Settings</w:Selector></w:SelectorSet>`
	expectedPutBody := `<h:AMT_TLSSettingData xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData">` +
		`<h:AcceptNonSecureConnections>true</h:AcceptNonSecureConnections>` +
		`<h:ElementName>Intel(r) AMT LMS TLS Settings</h:ElementName>` +
		`<h:Enabled>true</h:Enabled>` +
		`<h:InstanceID>Intel(r) AMT LMS TLS Settings</h:InstanceID>` +
		`<h:MutualAuthentication>true</h:MutualAuthentication>` +
		`</h:AMT_TLSSettingData>`

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			extraHeader  string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_TLSSettingData Get wsman message", "AMT_TLSSettingData", wsmantesting.GET, "", "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_TLSSettingData Enumerate wsman message", "AMT_TLSSettingData", wsmantesting.ENUMERATE, "", wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_TLSSettingData Pull wsman message", "AMT_TLSSettingData", wsmantesting.PULL, "", wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			//PUT
			{"should create a valid AMT_TLSSettingData Put wsman message", "AMT_TLSSettingData", wsmantesting.PUT, expectedPutSelector, expectedPutBody, func() string { return elementUnderTest.Put(data) }},
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
