/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package managementpresence

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/common"
	"github.com/stretchr/testify/assert"
	//"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

type MockClient struct {
}

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_ManagementPresence><g:CreationClassName>AMT_ManagementPresence</g:CreationClassName><g:ElementName>Intel(r) AMT Management Presence</g:ElementName><g:Name>Intel(r) AMT Management Presence</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_ManagementPresence>`
)

var currentMessage = ""

func (c *MockClient) Post(msg string) ([]byte, error) {
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/amt/managementpresence/" + strings.ToLower(currentMessage) + ".xml")
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
func TestAMT_ManagementPresenceRemoteSAP(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := MockClient{} // wsman.NewClient("http://localhost:16992/wsman", "admin", "P@ssw0rd", true)
	//elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)
	// enumerationId := ""
	//client := wsman.NewClient("http://localhost:16992/wsman", "admin", "Intel123!", true)
	elementUnderTest := NewManagementPresenceRemoteSAPWithClient(wsmanMessageCreator, &client)

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
				"should create a valid AMT_ManagementPresenceRemoteSAP Get wsman message",
				"AMT_ManagementPresenceRemoteSAP",
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					currentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					ManagementRemote: ManagementRemote{
						AccessInfo:              "",
						CN:                      "",
						CreationClassName:       "",
						ElementName:             "",
						InfoFormat:              0,
						Name:                    "",
						Port:                    0,
						SystemCreationClassName: "",
						SystemName:              "",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Enumerate wsman message",
				"AMT_ManagementPresenceRemoteSAP",
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
						EnumerationContext: "C9000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			// {"should create a valid AMT_ManagementPresenceRemoteSAP Pull wsman message", "AMT_ManagementPresenceRemoteSAP", wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			//DELETE
			// {"should create a valid AMT_ManagementPresenceRemoteSAP Delete wsman message", "AMT_ManagementPresenceRemoteSAP", wsmantesting.DELETE, "", "<w:SelectorSet><w:Selector Name=\"Name\">instanceID123</w:Selector></w:SelectorSet>", func() string { return elementUnderTest.Delete("instanceID123") }},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				//println(response.XMLOutput)
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
