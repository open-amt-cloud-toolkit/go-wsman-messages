/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_PublicKeyManagementService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := wsman.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewPublicKeyManagementService(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_PublicKeyManagementService Get wsman message", "AMT_PublicKeyManagementService", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_PublicKeyManagementService Enumerate wsman message", "AMT_PublicKeyManagementService", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_PublicKeyManagementService Pull wsman message", "AMT_PublicKeyManagementService", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},

			// PUBLIC KEY MANAGEMENT SERVICE
			{"should return a valid amt_PublicKeyManagementService AddTrustedRootCertificate wsman message", "AMT_PublicKeyManagementService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddTrustedRootCertificate`, fmt.Sprintf(`<h:AddTrustedRootCertificate_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:CertificateBlob>%s</h:CertificateBlob></h:AddTrustedRootCertificate_INPUT>`, wsmantesting.TrustedRootCert), func() string {
				return elementUnderTest.AddTrustedRootCertificate(wsmantesting.TrustedRootCert)
			}},

			{"should return a valid amt_PublicKeyManagementService GenerateKeyPair wsman message", "AMT_PublicKeyManagementService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GenerateKeyPair`, `<h:GenerateKeyPair_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyAlgorithm>0</h:KeyAlgorithm><h:KeyLength>2048</h:KeyLength></h:GenerateKeyPair_INPUT>`, func() string {
				params := GenerateKeyPair_INPUT{
					KeyAlgorithm: 0,
					KeyLength:    2048,
				}
				return elementUnderTest.GenerateKeyPair(params)
			}},

			{"should return a valid amt_PublicKeyManagementService AddCertificate wsman message", "AMT_PublicKeyManagementService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddCertificate`, fmt.Sprintf(`<h:AddCertificate_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:CertificateBlob>%s</h:CertificateBlob></h:AddCertificate_INPUT>`, wsmantesting.TrustedRootCert), func() string {
				return elementUnderTest.AddCertificate(wsmantesting.TrustedRootCert)
			}},

			{"should return a valid amt_PublicKeyManagementService GeneratePKCS10RequestEx wsman message", "AMT_PublicKeyManagementService", "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GeneratePKCS10RequestEx", `<h:GeneratePKCS10RequestEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyPair>test</h:KeyPair><h:NullSignedCertificateRequest>reallylongcertificateteststring</h:NullSignedCertificateRequest><h:SigningAlgorithm>1</h:SigningAlgorithm></h:GeneratePKCS10RequestEx_INPUT>`, func() string {
				pkcs10Request := PKCS10Request{
					KeyPair:                      "test",
					NullSignedCertificateRequest: "reallylongcertificateteststring",
					SigningAlgorithm:             1,
				}
				return elementUnderTest.GeneratePKCS10RequestEx(pkcs10Request)
			}},
			{"should return a valid amt_PublicKeyManagementService AddKey wsman message", "AMT_PublicKeyManagementService", "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddKey", `<h:AddKey_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyBlob>privatekey</h:KeyBlob></h:AddKey_INPUT>`, func() string {
				cert := []byte("privatekey")
				return elementUnderTest.AddKey(cert)
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
