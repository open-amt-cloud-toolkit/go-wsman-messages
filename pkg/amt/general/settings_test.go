/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package general

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_GeneralSettings(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	// client := wsmantesting.MockClient{
	// 	PackageUnderTest: "amt/general",
	// }
	client := wsman.NewClient("http://localhost:16992/wsman", "admin", "Intel123!", true)

	elementUnderTest := NewGeneralSettingsWithClient(wsmanMessageCreator, client)
	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_GeneralSettings Get wsman message",
				"AMT_GeneralSettings",
				wsmantesting.GET,
				"",
				func() (Response, error) {
					//client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					AMTGeneralSettings: GeneralSettings{
						XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings", Local: "AMT_GeneralSettings"},
						SettingData_OUTPUT:            models.SettingData_OUTPUT{InstanceID: "Intel(r) AMT: General Settings", ManagedElement_OUTPUT: models.ManagedElement_OUTPUT{ElementName: "Intel(r) AMT: General Settings"}},
						AMTNetworkEnabled:             1,
						DDNSPeriodicUpdateInterval:    1440,
						DDNSTTL:                       900,
						DDNSUpdateByDHCPServerEnabled: true,
						DDNSUpdateEnabled:             false,
						DHCPSyncRequiresHostname:      1, //Intel SDK documentation missing
						DHCPv6ConfigurationTimeout:    0,
						DigestRealm:                   "Digest:F3EB554784E729164447A89F60B641C5",
						DomainName:                    "Test Domain Name",
						HostName:                      "Test Host Name",
						HostOSFQDN:                    "Test Host OS FQDN",
						IdleWakeTimeout:               1,
						NetworkInterfaceEnabled:       true,
						PingResponseEnabled:           true,
						PowerSource:                   0,
						PreferredAddressFamily:        0,
						PresenceNotificationInterval:  0,
						PrivacyLevel:                  0,
						RmcpPingResponseEnabled:       true,
						SharedFQDN:                    true,
						ThunderboltDockEnabled:        0,
						WsmanOnlyMode:                 false,
					},
				},
			},

			//ENUMERATES
			{
				"should create a valid AMT_GeneralSettings Enumerate wsman message",
				"AMT_GeneralSettings",
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					//client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{},
			},
			// //PULLS
			{
				"should create a valid AMT_GeneralSettings Pull wsman message",
				"AMT_GeneralSettings",
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					//client.CurrentMessage = "Pull"
					//wsmantesting.EnumerationContext
					return elementUnderTest.Pull("14000000-0000-0000-0000-000000000000")
				},
				Body{},
			},
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
