/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

type MockClient struct {
}

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_PublicKeyCertificate><g:CreationClassName>AMT_PublicKeyCertificate</g:CreationClassName><g:ElementName>Intel(r) AMT Public Key Certificate</g:ElementName><g:Name>Intel(r) AMT Public Key Certificate</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_PublicKeyCertificate>`
)

var currentMessage = ""

func (c *MockClient) Post(msg string) ([]byte, error) {
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/amt/publickey/certificate/" + strings.ToLower(currentMessage) + ".xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer xmlFile.Close()
	// read file into string
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	// strip carriage returns and new line characters
	xmlData = []byte(strings.ReplaceAll(string(xmlData), "\r\n", ""))

	// Simulate a successful response for testing.
	return []byte(xmlData), nil
}
func TestAMT_PublicKeyCertificate(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := MockClient{}
	elementUnderTest := NewPublicKeyCertificateWithClient(wsmanMessageCreator, &client)
	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (ResponseCert, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_PublicKeyCertificate Get wsman message",
				AMT_PublicKeyCertificate,
				wsmantesting.GET,
				"",
				"",
				func() (ResponseCert, error) {
					currentMessage = "Get"
					return elementUnderTest.Get()
				},
				BodyCert{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					KeyCert: KeyCert{
						ElementName:           "Intel(r) AMT Certificate",
						InstanceID:            "Intel(r) AMT Certificate: Handle: 0",
						Issuer:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
						Subject:               "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
						TrustedRootCertficate: true,
						X509Certificate:       "MIIEOzCCAqOgAwIBAgIDAZMjMA0GCSqGSIb3DQEBDAUAMD0xFzAVBgNVBAMTDk1QU1Jvb3QtMGFmMWQ1MRAwDgYDVQQKEwd1bmtub3duMRAwDgYDVQQGEwd1bmtub3duMCAXDTIyMDkyNDEwNDUwOFoYDzIwNTMwOTI0MTA0NTA4WjA9MRcwFQYDVQQDEw5NUFNSb290LTBhZjFkNTEQMA4GA1UEChMHdW5rbm93bjEQMA4GA1UEBhMHdW5rbm93bjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBALz/oJNyWXlClSlteAieC8Uyd4A+tbn8b45k6LKiImhDmdz/xFo9xe0C9GNf7b42KVpg5WoH/sPhoClR9Tv5i1LnilT1SUir42fcm2NEV9dRcLsPd/RAQfz8u0D4zb3blnxE8isqzriNpG7kac35UidSr5ym8TZ3IwXx6JJuncGgfB0DFZADC/+dA74n3coykvWBYqLr6RI5pkAxvulkRlCsatJTJrvMUYJ51GI28jV56mIAc89sLrHqiSKCZBH9AcUrnZ/cB6ST/IikXpxy5wXBIvWT3VKVq75T/uIoCBEp5TLEn1EOYGqBBOCSQgmtmX7eVaB0s1+ppPW9w9a2zS45cHAtQ7tYvkkPv2dRhSzZdlk6HRXDP5wsF0aiflZCgbrjkq0SFC4e3Lo7XQX3FTNb0SOTZVTydupoMKkgJQTNlcosdu1ZzaIBl3eSkKkJZz2rUTssZC5tn9vcDd5vy3BzcGh5pvkgfAgN1sydqG7Ke1qCkNEzm11B/BsevatjjwIDAQABo0IwQDAMBgNVHRMEBTADAQH/MBEGCWCGSAGG+EIBAQQEAwIABzAdBgNVHQ4EFgQUCvHVQqerCid99eLApuLky9x6H5owDQYJKoZIhvcNAQEMBQADggGBAIzOyGV0hzsmH2biJlzwTZaHMxqS7boTFMkHw+KvzsI201tHqVmCoiQ8EHErBGLSoDOTDRgOUGOCA5XU5ie9OWupAGqKBSwIyAhmJMOzrzC4Gwpu8K1msoFJH30kx/V9purpbS3BRj0xfYXLa6IczbTg3E5IfTnZRJ9YuUtKQfI0P9c5U9CoKtddKn4+lRvOjFDoYfQGCJ7go3xjNCcGCVCjfkUhAVdbQ21DCRr6/YCZDWmjzZpL0p7UKF8roTiNuL/Z7gIXxch5HOmEWHY9uQ6K2MntuxAu0aK/mSD2kwmt/ECongdEGfUvhULLoPRQlQ2LnzcUQEgMECGQR5Yfy9jT0E8zdWDpc2tgVioNu6rEYKgp/GhG+sv7jv58pW82FRAV9xXtftW9+XDugC8tBJ6JHn0Q2v0QAflD2CEQVhWAY8bAqrbfTGUsaLfGL6kxV/qqssoMgLR8Whq96T5le/4XGhQpbCHWIlctD6MwbrsunIAeQKp1Sc3DosY7DLq1MQ==",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_PublicKeyCertificate Enumerate wsman message",
				AMT_PublicKeyCertificate,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (ResponseCert, error) {
					currentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				BodyCert{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CB000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			// {"should create a valid AMT_PublicKeyCertificate Pull wsman message", AMT_PublicKeyCertificate, wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			//PUTS
			// {"should create a valid AMT_PublicKeyCertificate Put wsman message", AMT_PublicKeyCertificate, wsmantesting.PUT, expectedPutCertBody, "", func() string { return elementUnderTest.Put(putCert) }},
			//DELETE
			// {"should create a valid AMT_PublicKeyCertificate Delete wsman message", AMT_PublicKeyCertificate, wsmantesting.DELETE, "", "<w:SelectorSet><w:Selector Name=\"InstanceID\">instanceID123</w:Selector></w:SelectorSet>", func() string { return elementUnderTest.Delete("instanceID123") }},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.BodyCert)
			})
		}
	})
}
