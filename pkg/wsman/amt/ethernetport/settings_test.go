/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ethernetport

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
	"github.com/stretchr/testify/assert"
)

type MockClient struct {
}

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_EthernetPortSettings><g:CreationClassName>AMT_EthernetPortSettings</g:CreationClassName><g:ElementName>Intel(r) AMT Ethernet Port Settings</g:ElementName><g:Name>Intel(r) AMT Ethernet Port Settings</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_EthernetPortSettings>`
)

var currentMessage = ""

func (c *MockClient) Post(msg string) ([]byte, error) {
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/amt/ethernetport/" + strings.ToLower(currentMessage) + ".xml")
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
func TestAMT_EthernetPortSettings(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := MockClient{}
	elementUnderTest := NewEthernetPortSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("amt_* Tests", func(t *testing.T) {
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
				"should create a valid AMT_EthernetPortSettings Get wsman message",
				"AMT_EthernetPortSettings",
				wsmantesting.GET,
				"",
				"<w:SelectorSet><w:Selector Name=\"test\">test</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					currentMessage = "Get"
					selector := Selector{
						Name:  "test",
						Value: "test",
					}
					return elementUnderTest.Get(selector)
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					EthernetPort: EthernetPort{
						DHCPEnabled:            true,
						DefaultGateway:         "192.168.0.1",
						ElementName:            "Intel(r) AMT Ethernet Port Settings",
						InstanceID:             "Intel(r) AMT Ethernet Port Settings 0",
						IpSyncEnabled:          true,
						LinkIsUp:               true,
						LinkPolicy:             14,
						MACAddress:             "c8-d9-d2-7a-1e-33",
						PhysicalConnectionType: 0,
						PrimaryDNS:             "68.105.28.11",
						SecondaryDNS:           "68.105.29.11",
						SharedDynamicIP:        true,
						SharedMAC:              true,
						SharedStaticIp:         false,
						SubnetMask:             "255.255.255.0",
					},
				},
			},

			//ENUMERATES
			{
				"should create a valid AMT_EthernetPortSettings Enumerate wsman message",
				"AMT_EthernetPortSettings",
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (Response, error) {
					currentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "7700000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			// {"should create a valid AMT_EthernetPortSettings Pull wsman message", "AMT_EthernetPortSettings", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
