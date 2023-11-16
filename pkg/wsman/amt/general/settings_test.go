/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package general

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
	"github.com/stretchr/testify/assert"
)

type MockClient struct {
}

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_GeneralSettings><g:CreationClassName>AMT_GeneralSettings</g:CreationClassName><g:ElementName>Intel(r) AMT General Settings</g:ElementName><g:Name>Intel(r) AMT General Settings</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_GeneralSettings>`
)

var currentMessage = ""

func (c *MockClient) Post(msg string) ([]byte, error) {
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/amt/general/" + strings.ToLower(currentMessage) + ".xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer xmlFile.Close()
	// read file into string
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	// strip carriage returns and new line characters
	xmlData = []byte(strings.ReplaceAll(string(xmlData), "\r\n", ""))

	// Simulate a successful response for testing.
	return []byte(xmlData), nil
}
func TestAMT_GeneralSettings(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)

	client := MockClient{}
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
				"AMT_GeneralSettings",
				wsmantesting.GET,
				"",
				func() (Response, error) {
					currentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					AMTGeneralSettings: GeneralSettings{
						XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings", Local: "AMT_GeneralSettings"},
						SettingData:                   models.SettingData{InstanceID: "Intel(r) AMT: General Settings", ManagedElement: models.ManagedElement{ElementName: ""}},
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
					currentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "14000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_GeneralSettings Pull wsman message",
				"AMT_GeneralSettings",
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					currentMessage = "Pull"
					//wsmantesting.EnumerationContext
					//return elementUnderTest.Pull("14000000-0000-0000-0000-000000000000")
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponse: PullResponse{
						Items: []Item{
							{
								AMTGeneralSettings: GeneralSettings{
									XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings", Local: "AMT_GeneralSettings"},
									SettingData:                   models.SettingData{InstanceID: "Intel(r) AMT: General Settings", ManagedElement: models.ManagedElement{ElementName: ""}},
									AMTNetworkEnabled:             1,
									DDNSPeriodicUpdateInterval:    1440,
									DDNSTTL:                       900,
									DDNSUpdateByDHCPServerEnabled: true,
									DDNSUpdateEnabled:             false,
									DHCPSyncRequiresHostname:      0, //Intel SDK documentation missing
									DHCPv6ConfigurationTimeout:    0,
									DigestRealm:                   "Digest:6EE8C61BA74893E059F032EA919D699E",
									IdleWakeTimeout:               65535,
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
