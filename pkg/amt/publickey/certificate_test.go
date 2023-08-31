/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_PublicKeyCertificate(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := wsman.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewPublicKeyCertificate(wsmanMessageCreator)
	putCert := PublicKeyCertificate{
		ElementName:           "",
		InstanceID:            "",
		X509Certificate:       "",
		TrustedRootCertficate: false,
		Issuer:                "",
		Subject:               "",
		ReadOnlyCertificate:   false,
	}
	expectedPutCertBody := `<PublicKeyCertificate><ElementName></ElementName><InstanceID></InstanceID><X509Certificate></X509Certificate><TrustedRootCertficate>false</TrustedRootCertficate><Issuer></Issuer><Subject></Subject><ReadOnlyCertificate>false</ReadOnlyCertificate></PublicKeyCertificate>`

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
			{"should create a valid AMT_PublicKeyCertificate Get wsman message", AMT_PublicKeyCertificate, wsmantesting.GET, "", "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_PublicKeyCertificate Enumerate wsman message", AMT_PublicKeyCertificate, wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, "", elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_PublicKeyCertificate Pull wsman message", AMT_PublicKeyCertificate, wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			//PUTS
			{"should create a valid AMT_PublicKeyCertificate Put wsman message", AMT_PublicKeyCertificate, wsmantesting.PUT, expectedPutCertBody, "", func() string { return elementUnderTest.Put(putCert) }},
			//DELETE
			{"should create a valid AMT_PublicKeyCertificate Delete wsman message", AMT_PublicKeyCertificate, wsmantesting.DELETE, "", "<w:SelectorSet><w:Selector Name=\"InstanceID\">instanceID123</w:Selector></w:SelectorSet>", func() string { return elementUnderTest.Delete("instanceID123") }},
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
