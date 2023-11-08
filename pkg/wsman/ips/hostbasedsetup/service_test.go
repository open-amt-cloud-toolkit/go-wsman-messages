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

func TestIPS_HostBasedSetupService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewHostBasedSetupService(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid IPS_HostBasedSetupService Get wsman message", "IPS_HostBasedSetupService", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid IPS_HostBasedSetupService Enumerate wsman message", "IPS_HostBasedSetupService", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid IPS_HostBasedSetupService Pull wsman message", "IPS_HostBasedSetupService", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},

			// ADD NEXT CERT IN CHAIN
			{"should create a valid IPS_HostBasedSetupService AddNextCertInChain wsman message", "IPS_HostBasedSetupService", wsmantesting.ADD_NEXT_CERT_IN_CHAIN, fmt.Sprintf(`<h:AddNextCertInChain_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NextCertificate>%s</h:NextCertificate><h:IsLeafCertificate>true</h:IsLeafCertificate><h:IsRootCertificate>false</h:IsRootCertificate></h:AddNextCertInChain_INPUT>`, wsmantesting.ClientCertificate), func() string {
				return elementUnderTest.AddNextCertInChain(wsmantesting.ClientCertificate, true, false)
			}},

			// AdminSetup
			{"should create a valid IPS_HostBasedSetupService AdminSetup wsman message", "IPS_HostBasedSetupService", wsmantesting.ADMIN_SETUP, fmt.Sprintf(`<h:AdminSetup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword><h:McNonce>%s</h:McNonce><h:SigningAlgorithm>%d</h:SigningAlgorithm><h:DigitalSignature>%s</h:DigitalSignature></h:AdminSetup_INPUT>`, wsmantesting.AdminPassEncryptionType, "f73b2c17b1ecbd7a235ec37d66cbed71", wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature), func() string {
				return elementUnderTest.AdminSetup(wsmantesting.AdminPassEncryptionType, wsmantesting.DigestRealm, wsmantesting.AdminPassword, wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature)
			}},

			// UpgradeToAdminSetup
			{"should create a valid IPS_HostBasedSetupService UpgradeToAdminSetup wsman message", "IPS_HostBasedSetupService", wsmantesting.UPGRADE_CLIENT_TO_ADMIN, fmt.Sprintf(`<h:UpgradeClientToAdmin_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:McNonce>%s</h:McNonce><h:SigningAlgorithm>%d</h:SigningAlgorithm><h:DigitalSignature>%s</h:DigitalSignature></h:UpgradeClientToAdmin_INPUT>`, wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature), func() string {
				return elementUnderTest.UpgradeClientToAdmin(wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature)
			}},

			//Setup
			{"should create a valid IPS_HostBasedSetupService Setup wsman message", "IPS_HostBasedSetupService", wsmantesting.SETUP, fmt.Sprintf(`<h:Setup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword></h:Setup_INPUT>`, wsmantesting.AdminPassEncryptionType, "f73b2c17b1ecbd7a235ec37d66cbed71"), func() string {
				return elementUnderTest.Setup(wsmantesting.AdminPassEncryptionType, wsmantesting.DigestRealm, wsmantesting.AdminPassword)
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
