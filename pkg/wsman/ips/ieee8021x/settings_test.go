/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveIPS_IEEE8021xSettings(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/ieee8021x/settings",
	}
	elementUnderTest := NewIEEE8021xSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("ips_IEEE8021xSettings Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid IPS_IEEE8021xSettings Get wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{},
			},
			//ENUMERATES
			{
				"should create a valid IPS_IEEE8021xSettings Enumerate wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{},
			},
			//PULLS
			{
				"should create a valid IPS_IEEE8021xSettings Pull wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{},
			},
			// SET CERTIFICATES
			{
				"should create a valid ips_IEEE8021xSettings set certificates wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.SET_CERTIFICATES,
				fmt.Sprintf(`<h:SetCertificates_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings"><h:ServerCertificateIssuer>%s</h:ServerCertificateIssuer><h:ClientCertificate>%s</h:ClientCertificate></h:SetCertificates_INPUT>`, wsmantesting.ServerCertificateIssuer, wsmantesting.ClientCertificate),
				"",
				func() (Response, error) {
					client.CurrentMessage = "SetCertificates"
					return elementUnderTest.SetCertificates(wsmantesting.ServerCertificateIssuer, wsmantesting.ClientCertificate)
				},
				Body{},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
