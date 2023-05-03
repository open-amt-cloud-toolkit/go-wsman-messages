/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_SetupAndConfigurationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := wsman.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewSetupAndConfigurationService(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_SetupAndConfigurationService Get wsman message", "AMT_SetupAndConfigurationService", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_SetupAndConfigurationService Enumerate wsman message", "AMT_SetupAndConfigurationService", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_SetupAndConfigurationService Pull wsman message", "AMT_SetupAndConfigurationService", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			{"should create a valid AMT_SetupAndConfigurationService CommitChanges wsman message", "AMT_SetupAndConfigurationService", string(actions.CommitChanges), `<h:CommitChanges_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"></h:CommitChanges_INPUT>`, elementUnderTest.CommitChanges},
			{"should create a valid AMT_SetupAndConfigurationService GetUuid wsman message", "AMT_SetupAndConfigurationService", string(actions.GetUuid), `<h:GetUuid_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"></h:GetUuid_INPUT>`, elementUnderTest.GetUuid},
			{"should create a valid AMT_SetupAndConfigurationService SetMEBxPassword wsman message", "AMT_SetupAndConfigurationService", string(actions.SetMEBxPassword), `<h:SetMEBxPassword_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"><h:Password>P@ssw0rd</h:Password></h:SetMEBxPassword_INPUT>`, func() string { return elementUnderTest.SetMEBXPassword("P@ssw0rd") }},
			{"should create a valid AMT_SetupAndConfigurationService Unprovision wsman message", "AMT_SetupAndConfigurationService", "Unprovision", `<h:Unprovision_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"><h:ProvisioningMode>1</h:ProvisioningMode></h:Unprovision_INPUT>`, func() string { return elementUnderTest.Unprovision(1) }},
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
