/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportconfiguration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_WiFiPortConfigurationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := wsman.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewWiFiPortConfigurationService(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_WiFiPortConfigurationService Get wsman message", "AMT_WiFiPortConfigurationService", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_WiFiPortConfigurationService Enumerate wsman message", "AMT_WiFiPortConfigurationService", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_WiFiPortConfigurationService Pull wsman message", "AMT_WiFiPortConfigurationService", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			// WIFI PORT CONFIGURATION SERVICE
			{"should return a valid amt_WiFiPortConfigurationService ADD_WIFI_SETTINGS wsman message", "AMT_WiFiPortConfigurationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService/AddWiFiSettings`, `<h:AddWiFiSettings_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService"><h:WiFiEndpoint><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpoint</w:ResourceURI><w:SelectorSet><w:Selector Name="Name">WiFi Endpoint 0</w:Selector></w:SelectorSet></a:ReferenceParameters></h:WiFiEndpoint><h:WiFiEndpointSettingsInput xmlns:q="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings"><q:ElementName>home</q:ElementName><q:InstanceID>Intel(r) AMT:WiFi Endpoint Settings home</q:InstanceID><q:AuthenticationMethod>6</q:AuthenticationMethod><q:EncryptionMethod>4</q:EncryptionMethod><q:SSID>admin</q:SSID><q:Priority>1</q:Priority><q:PSKPassPhrase>p&#39;ass&lt;&gt;&amp;&#34;code</q:PSKPassPhrase></h:WiFiEndpointSettingsInput></h:AddWiFiSettings_INPUT>`, func() string {
				wifiEndpointSettings := models.WiFiEndpointSettings{
					H:                    "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings", //todo: make more dynamic
					ElementName:          "home",
					InstanceID:           "Intel(r) AMT:WiFi Endpoint Settings home",
					AuthenticationMethod: 6,
					EncryptionMethod:     4,
					SSID:                 "admin",
					Priority:             1,
					PSKPassPhrase:        "p'ass<>&\"code",
				}
				return elementUnderTest.AddWiFiSettings(wifiEndpointSettings, nil, "WiFi Endpoint 0", "", "")
			}},

			{"should create a valid AMT_WiFiPortConfigurationService ADD_WIFI_SETTINGS 8021x wsman message", "AMT_WiFiPortConfigurationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService/AddWiFiSettings`, `<h:AddWiFiSettings_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService"><h:WiFiEndpoint><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpoint</w:ResourceURI><w:SelectorSet><w:Selector Name="Name">WiFi Endpoint 0</w:Selector></w:SelectorSet></a:ReferenceParameters></h:WiFiEndpoint><h:WiFiEndpointSettingsInput xmlns:q="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings"><q:ElementName>home</q:ElementName><q:InstanceID>Intel(r) AMT:WiFi Endpoint Settings home</q:InstanceID><q:AuthenticationMethod>6</q:AuthenticationMethod><q:EncryptionMethod>4</q:EncryptionMethod><q:SSID>admin</q:SSID><q:Priority>1</q:Priority><q:PSKPassPhrase>p&#39;ass&lt;&gt;&amp;&#34;code</q:PSKPassPhrase></h:WiFiEndpointSettingsInput><h:IEEE8021xSettingsInput xmlns:q="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_IEEE8021xSettings"><q:ElementName>wifi_8021x_profile</q:ElementName><q:AuthenticationProtocol>0</q:AuthenticationProtocol></h:IEEE8021xSettingsInput><h:ClientCredential><a:Address>default</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">handle 0</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ClientCredential><h:CACredential><a:Address>default</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">handle 1</w:Selector></w:SelectorSet></a:ReferenceParameters></h:CACredential></h:AddWiFiSettings_INPUT>`, func() string {
				wifiEndpointSettings := models.WiFiEndpointSettings{
					H:                    "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings",
					ElementName:          "home",
					InstanceID:           "Intel(r) AMT:WiFi Endpoint Settings home",
					AuthenticationMethod: 6,
					EncryptionMethod:     4,
					SSID:                 "admin",
					Priority:             1,
					PSKPassPhrase:        "p'ass<>&\"code",
				}
				ieee8021xSettingsInput := &models.IEEE8021xSettings{
					H:                      "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_IEEE8021xSettings",
					ElementName:            "wifi_8021x_profile",
					AuthenticationProtocol: 0,
				}
				clientCredential := "handle 0"
				caCredential := "handle 1"
				return elementUnderTest.AddWiFiSettings(wifiEndpointSettings, ieee8021xSettingsInput, "WiFi Endpoint 0", clientCredential, caCredential)
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
