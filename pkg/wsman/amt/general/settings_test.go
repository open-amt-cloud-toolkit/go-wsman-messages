/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package general

import (
	"encoding/xml"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
	"github.com/stretchr/testify/assert"
)

func TestPositiveAMT_GeneralSettings(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)

	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/general",
	}
	elementUnderTest := NewGeneralSettingsWithClient(wsmanMessageCreator, &client)
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
				AMT_GeneralSettings,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: GeneralSettingsResponse{
						XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings", Local: AMT_GeneralSettings},
						ElementName:                   "Intel(r) AMT: General Settings",
						InstanceID:                    "Intel(r) AMT: General Settings",
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
				AMT_GeneralSettings,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "14000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_GeneralSettings Pull wsman message",
				AMT_GeneralSettings,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						GeneralSettingsItems: []GeneralSettingsResponse{
							{
								XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings", Local: AMT_GeneralSettings},
								ElementName:                   "Intel(r) AMT: General Settings",
								InstanceID:                    "Intel(r) AMT: General Settings",
								AMTNetworkEnabled:             1,
								DDNSPeriodicUpdateInterval:    1440,
								DDNSTTL:                       900,
								DDNSUpdateByDHCPServerEnabled: true,
								DDNSUpdateEnabled:             false,
								DHCPv6ConfigurationTimeout:    0,
								DigestRealm:                   "Digest:6EE8C61BA74893E059F032EA919D699E",
								DomainName:                    "",
								HostOSFQDN:                    "",
								IdleWakeTimeout:               65535,
								NetworkInterfaceEnabled:       true,
								PingResponseEnabled:           true,
								PowerSource:                   0,
								PreferredAddressFamily:        0,
								PresenceNotificationInterval:  0,
								PrivacyLevel:                  0,
								RmcpPingResponseEnabled:       true,
								SharedFQDN:                    true,
								WsmanOnlyMode:                 false,
							},
						},
					},
				},
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
func TestNegativeAMT_GeneralSettings(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)

	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/general",
	}
	elementUnderTest := NewGeneralSettingsWithClient(wsmanMessageCreator, &client)
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
				AMT_GeneralSettings,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: GeneralSettingsResponse{
						XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings", Local: AMT_GeneralSettings},
						ElementName:                   "Intel(r) AMT: General Settings",
						InstanceID:                    "Intel(r) AMT: General Settings",
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
				AMT_GeneralSettings,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "14000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_GeneralSettings Pull wsman message",
				AMT_GeneralSettings,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						GeneralSettingsItems: []GeneralSettingsResponse{
							{
								XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings", Local: AMT_GeneralSettings},
								ElementName:                   "Intel(r) AMT: General Settings",
								InstanceID:                    "Intel(r) AMT: General Settings",
								AMTNetworkEnabled:             1,
								DDNSPeriodicUpdateInterval:    1440,
								DDNSTTL:                       900,
								DDNSUpdateByDHCPServerEnabled: true,
								DDNSUpdateEnabled:             false,
								DHCPv6ConfigurationTimeout:    0,
								DigestRealm:                   "Digest:6EE8C61BA74893E059F032EA919D699E",
								DomainName:                    "",
								HostOSFQDN:                    "",
								IdleWakeTimeout:               65535,
								NetworkInterfaceEnabled:       true,
								PingResponseEnabled:           true,
								PowerSource:                   0,
								PreferredAddressFamily:        0,
								PresenceNotificationInterval:  0,
								PrivacyLevel:                  0,
								RmcpPingResponseEnabled:       true,
								SharedFQDN:                    true,
								WsmanOnlyMode:                 false,
							},
						},
					},
				},
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
