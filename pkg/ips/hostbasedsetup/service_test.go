/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbasedsetup

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestIPS_HostBasedSetupService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := wsman.NewWSManMessageCreator(resourceUriBase)
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
			{"should create a valid IPS_HostBasedSetupService AdminSetup wsman message", "IPS_HostBasedSetupService", wsmantesting.ADMIN_SETUP, fmt.Sprintf(`<h:AdminSetup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword><h:McNonce>%s</h:McNonce><h:SigningAlgorithm>%d</h:SigningAlgorithm><h:DigitalSignature>%s</h:DigitalSignature></h:AdminSetup_INPUT>`, wsmantesting.AdminPassEncryptionType, wsmantesting.AdminPassword, wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature), func() string {
				return elementUnderTest.AdminSetup(wsmantesting.AdminPassEncryptionType, wsmantesting.AdminPassword, wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature)
			}},

			//ADD NEXT CERT IN CHAIN
			{"should create a valid IPS_HostBasedSetupService Setup wsman message", "IPS_HostBasedSetupService", wsmantesting.SETUP, fmt.Sprintf(`<h:Setup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword></h:Setup_INPUT>`, wsmantesting.AdminPassEncryptionType, wsmantesting.AdminPassword), func() string {
				return elementUnderTest.Setup(wsmantesting.AdminPassEncryptionType, wsmantesting.AdminPassword)
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
