package wifiportconfiguration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
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
			// {"should return a valid amt_WiFiPortConfigurationService ADD_WIFI_SETTINGS wsman message", "AMT_WiFiPortConfigurationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService/AddWiFiSettings`, `<h:AddWiFiSettings_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService"><h:WiFiEndpoint><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpoint</w:ResourceURI><w:SelectorSet><w:Selector Name="${selector.name}">${selector.value}</w:Selector></w:SelectorSet></a:ReferenceParameters></h:WiFiEndpoint><h:WiFiEndpointSettingsInput xmlns:q="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings"><h:ElementName>${wifiEndpointSettings.ElementName}</h:ElementName><h:InstanceID>${wifiEndpointSettings.InstanceID}</h:InstanceID><h:AuthenticationMethod>${wifiEndpointSettings.AuthenticationMethod}</h:AuthenticationMethod><h:EncryptionMethod>${wifiEndpointSettings.EncryptionMethod}</h:EncryptionMethod><h:SSID>${wifiEndpointSettings.SSID}</h:SSID><h:Priority>${wifiEndpointSettings.Priority}</h:Priority><h:PSKPassPhrase>p&apos;ass&lt;&gt;&amp;&quot;code</h:PSKPassPhrase></h:WiFiEndpointSettingsInput></h:AddWiFiSettings_INPUT>`, func() string {
			// 	ieee8021xSettingsInput := &cimModels.IEEE8021xSettings{}
			// 	var clientCredential string
			// 	var caCredential string
			// 	selector := wsman.Selector{
			// 		Name:  "Name",
			// 		Value: "WiFi Endpoint 0",
			// 	}
			// 	wifiEndpointSettings := cimModels.WiFiEndpointSettings{
			// 		ElementName:          "home",
			// 		InstanceID:           "Intel(r) AMT:WiFi Endpoint Settings home",
			// 		AuthenticationMethod: 6,
			// 		EncryptionMethod:     4,
			// 		SSID:                 "admin",
			// 		Priority:             1,
			// 		PSKPassPhrase:        `p\'ass<>&"code`,
			// 	}
			// 	return amtClass.WiFiPortConfigurationService.AddWiFiSettings(wifiEndpointSettings, selector, ieee8021xSettingsInput, clientCredential, caCredential)
			// }},

			// {"should create a valid AMT_WiFiPortConfigurationService Pulls wsman message", "AMT_WiFiPortConfigurationService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService/AddWiFiSettings`, `<h:AddWiFiSettings_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService"><h:WiFiEndpoint><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpoint</w:ResourceURI><w:SelectorSet><w:Selector Name="${selector.name}">${selector.value}</w:Selector></w:SelectorSet></a:ReferenceParameters></h:WiFiEndpoint><h:WiFiEndpointSettingsInput xmlns:q="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings"><h:ElementName>${wifiEndpointSettings.ElementName}</h:ElementName><h:InstanceID>${wifiEndpointSettings.InstanceID}</h:InstanceID><h:AuthenticationMethod>${wifiEndpointSettings.AuthenticationMethod}</h:AuthenticationMethod><h:EncryptionMethod>${wifiEndpointSettings.EncryptionMethod}</h:EncryptionMethod><h:SSID>${wifiEndpointSettings.SSID}</h:SSID><h:Priority>${wifiEndpointSettings.Priority}</h:Priority><h:PSKPassPhrase>p&apos;ass&lt;&gt;&amp;&quot;code</h:PSKPassPhrase></h:WiFiEndpointSettingsInput><h:ieee8021xSettingsInput xmlns:q="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_IEEE8021xSettings"><h:ElementName>wifi_8021x_profile</h:ElementName><h:AuthenticationProtocol>0</h:AuthenticationProtocol></h:ieee8021xSettingsInput><h:ClientCredential>handle 0</h:ClientCredential><h:CACredential>handle 1</h:CACredential></h:AddWiFiSettings_INPUT>`, func() string {
			// 	selector := wsman.Selector{
			// 		Name:  "Name",
			// 		Value: "WiFi Endpoint 0",
			// 	}
			// 	wifiEndpointSettings := cimModels.WiFiEndpointSettings{
			// 		ElementName:          "home",
			// 		InstanceID:           "Intel(r) AMT:WiFi Endpoint Settings home",
			// 		AuthenticationMethod: 6,
			// 		EncryptionMethod:     4,
			// 		SSID:                 "admin",
			// 		Priority:             1,
			// 		PSKPassPhrase:        "p'ass<>&\"code",
			// 	}
			// 	ieee8021xSettingsInput := &cimModels.IEEE8021xSettings{
			// 		ElementName:            "wifi_8021x_profile",
			// 		AuthenticationProtocol: 0,
			// 	}
			// 	clientCredential := "handle 0"
			// 	caCredential := "handle 1"
			// 	return amtClass.WiFiPortConfigurationService.AddWiFiSettings(wifiEndpointSettings, selector, ieee8021xSettingsInput, clientCredential, caCredential)
			// }},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				correctResponse := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.body)
				messageID++
				response := test.responseFunc()
				if response != correctResponse {
					assert.Equal(t, correctResponse, response)
				}
			})
		}
	})
}
