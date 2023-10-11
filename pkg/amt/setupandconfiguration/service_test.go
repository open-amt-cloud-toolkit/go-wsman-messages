/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
	//"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/common"
)

func TestAMT_SetupAndConfigurationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	//client := MockClient{} // wsman.NewClient("http://localhost:16992/wsman", "admin", "P@ssw0rd", true)
	//elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)
	// enumerationId := ""
	client := wsman.NewClient("http://localhost:16992/wsman", "admin", "Intel123!", true)
	elementUnderTest := NewSetupAndConfigurationServiceWithClient(wsmanMessageCreator, client)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_SetupAndConfigurationService Get wsman message", 
				"AMT_SetupAndConfigurationService", 
				wsmantesting.GET, "", 
				func() (Response, error) {
					//currentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					Setup: Setup{
						CreationClassName: "",
						ElementName: "",
						EnabledState: 0,
						Name: "",
						PasswordModel: 0,
						ProvisioningMode: 0,
						ProvisioningServerOTP: "",
						ProvisioningState: 0,
						RequestedState: 0,
						SystemCreationClassName: "",
						SystemName: "",
						ZeroTouchConfigurationEnabled: false,
					},
				},
			},
			//ENUMERATES
			// {
			// 	"should create a valid AMT_SetupAndConfigurationService Enumerate wsman message", 
			// 	"AMT_SetupAndConfigurationService", 
			// 	wsmantesting.ENUMERATE, 
			// 	wsmantesting.ENUMERATE_BODY, 
			// 	func() (Response, error) {
			// 		//client.CurrentMessage = "Enumerate"
			// 		return elementUnderTest.Enumerate()
			// 	}, 
			// 	Body{
			// 		XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
			// 		EnumerateResponse: common.EnumerateResponse{
			// 			EnumerationContext: "5C000000-0000-0000-0000-000000000000",
			// 		},
			// 	},
			// },
			//PULLS
			// {"should create a valid AMT_SetupAndConfigurationService Pull wsman message", "AMT_SetupAndConfigurationService", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			// {"should create a valid AMT_SetupAndConfigurationService CommitChanges wsman message", "AMT_SetupAndConfigurationService", string(actions.CommitChanges), `<h:CommitChanges_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"></h:CommitChanges_INPUT>`, elementUnderTest.CommitChanges},
			// {"should create a valid AMT_SetupAndConfigurationService GetUuid wsman message", "AMT_SetupAndConfigurationService", string(actions.GetUuid), `<h:GetUuid_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"></h:GetUuid_INPUT>`, elementUnderTest.GetUuid},
			// {"should create a valid AMT_SetupAndConfigurationService SetMEBxPassword wsman message", "AMT_SetupAndConfigurationService", string(actions.SetMEBxPassword), `<h:SetMEBxPassword_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"><h:Password>P@ssw0rd</h:Password></h:SetMEBxPassword_INPUT>`, func() string { return elementUnderTest.SetMEBXPassword("P@ssw0rd") }},
			// {"should create a valid AMT_SetupAndConfigurationService Unprovision wsman message", "AMT_SetupAndConfigurationService", string(actions.Unprovision), `<h:Unprovision_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"><h:ProvisioningMode>1</h:ProvisioningMode></h:Unprovision_INPUT>`, func() string { return elementUnderTest.Unprovision(1) }},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				println(response.XMLOutput)
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
