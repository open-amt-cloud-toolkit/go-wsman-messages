/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbasedsetup

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveIPS_HostBasedSetupService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/hostbasedsetup",
	}
	elementUnderTest := NewHostBasedSetupServiceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_HostBasedSetupService Tests", func(t *testing.T) {
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
				"should create a valid IPS_HostBasedSetupService Get wsman message",
				IPS_HostBasedSetupService,
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
				"should create a valid IPS_HostBasedSetupService Enumerate wsman message",
				IPS_HostBasedSetupService,
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
				"should create a valid IPS_HostBasedSetupService Pull wsman message",
				IPS_HostBasedSetupService,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{},
			},

			// ADD NEXT CERT IN CHAIN
			{
				"should create a valid IPS_HostBasedSetupService AddNextCertInChain wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.ADD_NEXT_CERT_IN_CHAIN,
				fmt.Sprintf(`<h:AddNextCertInChain_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NextCertificate>%s</h:NextCertificate><h:IsLeafCertificate>true</h:IsLeafCertificate><h:IsRootCertificate>false</h:IsRootCertificate></h:AddNextCertInChain_INPUT>`, wsmantesting.ClientCertificate),
				"",
				func() (Response, error) {
					client.CurrentMessage = "AddNextCertiInChain"
					return elementUnderTest.AddNextCertInChain(wsmantesting.ClientCertificate, true, false)
				},
				Body{},
			},

			// AdminSetup
			{
				"should create a valid IPS_HostBasedSetupService AdminSetup wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.ADMIN_SETUP,
				fmt.Sprintf(`<h:AdminSetup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword><h:McNonce>%s</h:McNonce><h:SigningAlgorithm>%d</h:SigningAlgorithm><h:DigitalSignature>%s</h:DigitalSignature></h:AdminSetup_INPUT>`, wsmantesting.AdminPassEncryptionType, "f73b2c17b1ecbd7a235ec37d66cbed71", wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature),
				"",
				func() (Response, error) {
					client.CurrentMessage = "AdminSetup"
					return elementUnderTest.AdminSetup(wsmantesting.AdminPassEncryptionType, wsmantesting.DigestRealm, wsmantesting.AdminPassword, wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature)
				},
				Body{},
			},

			// UpgradeToAdminSetup
			{
				"should create a valid IPS_HostBasedSetupService UpgradeToAdminSetup wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.UPGRADE_CLIENT_TO_ADMIN,
				fmt.Sprintf(`<h:UpgradeClientToAdmin_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:McNonce>%s</h:McNonce><h:SigningAlgorithm>%d</h:SigningAlgorithm><h:DigitalSignature>%s</h:DigitalSignature></h:UpgradeClientToAdmin_INPUT>`, wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature),
				"",
				func() (Response, error) {
					client.CurrentMessage = "UpgradeToAdminSetup"
					return elementUnderTest.UpgradeClientToAdmin(wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature)
				},
				Body{},
			},

			//Setup
			{
				"should create a valid IPS_HostBasedSetupService Setup wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.SETUP,
				fmt.Sprintf(`<h:Setup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword></h:Setup_INPUT>`, wsmantesting.AdminPassEncryptionType, "f73b2c17b1ecbd7a235ec37d66cbed71"),
				"",
				func() (Response, error) {
					client.CurrentMessage = "Setup"
					return elementUnderTest.Setup(wsmantesting.AdminPassEncryptionType, wsmantesting.DigestRealm, wsmantesting.AdminPassword)
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

func TestCreateMD5Hash(t *testing.T) {
	tests := []struct {
		adminPassword string
		digestRealm   string
		expected      string
	}{
		{"adminPassword1", "digestRealm1", "7eab95087308c968d56947a05e916d6b"},
		{"adminPassword2", "digestRealm2", "b404159c55fafd0b4a8e7d64833c7f26"},
	}
	for _, test := range tests {
		t.Run(test.digestRealm, func(t *testing.T) {
			result := createMD5Hash(test.adminPassword, test.digestRealm)
			assert.Equal(t, test.expected, result)
		})
	}
}
