/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

type MockClient struct {
}

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_AuthorizationService><g:CreationClassName>AMT_RemoteAccessService</g:CreationClassName><g:ElementName>Intel(r) AMT Remote Access Service</g:ElementName><g:Name>Intel(r) AMT Remote Access Service</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT Remote Access Service>`
)

var currentMessage = ""

func (c *MockClient) Post(msg string) ([]byte, error) {
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/amt/remoteaccess/service/" + strings.ToLower(currentMessage) + ".xml")
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

func TestAMT_RemoteAccessService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := MockClient{}
	elementUnderTest := NewRemoteAccessServiceWithClient(wsmanMessageCreator, &client)
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
				"should create a valid AMT_RemoteAccessService Get wsman message",
				"AMT_RemoteAccessService",
				wsmantesting.GET,
				"",
				func() (Response, error) {
					currentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					RemoteAccess: RemoteAccess{
						CreationClassName:       "AMT_RemoteAccessService",
						ElementName:             "Intel(r) AMT Remote Access Service",
						Name:                    "Intel(r) AMT Remote Access Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_RemoteAccessService Enumerate wsman message",
				"AMT_RemoteAccessService",
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					currentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D3000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			// 	{"should create a valid AMT_RemoteAccessService Pull wsman message", "AMT_RemoteAccessService", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			// 	{"should create a valid AMT_RemoteAccessService AddMPS wsman message", "AMT_RemoteAccessService", string(actions.AddMps), fmt.Sprintf(`<h:AddMpServer_INPUT xmlns:h="%s%s"><h:AccessInfo>%s</h:AccessInfo><h:InfoFormat>%d</h:InfoFormat><h:Port>%d</h:Port><h:AuthMethod>%d</h:AuthMethod><h:Username>%s</h:Username><h:Password>%s</h:Password><h:CN>%s</h:CN></h:AddMpServer_INPUT>`, resourceUriBase, AMT_RemoteAccessService, "AccessInfo", 1, 2, 3, "Username", "Password", "CommonName"), func() string {
			// 		mpsServer := MPServer{
			// 			AccessInfo: "AccessInfo",
			// 			InfoFormat: 1,
			// 			Port:       2,
			// 			AuthMethod: 3,
			// 			Username:   "Username",
			// 			Password:   "Password",
			// 			CommonName: "CommonName",
			// 		}
			// 		return elementUnderTest.AddMPS(mpsServer)
			// 	}},
			// 	{"should create a valid AMT_RemoteAccessPolicyRule wsman message", "AMT_RemoteAccessService", string(actions.AddRemoteAccessPolicyRule), fmt.Sprintf(`<h:AddRemoteAccessPolicyRule_INPUT xmlns:h="%s%s"><h:Trigger>%d</h:Trigger><h:TunnelLifeTime>%d</h:TunnelLifeTime><h:ExtendedData>%s</h:ExtendedData><h:MpServer><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">%s%s</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="myselector">true</Selector></SelectorSet></ReferenceParameters></h:MpServer></h:AddRemoteAccessPolicyRule_INPUT>`, resourceUriBase, AMT_RemoteAccessService, 2, 0, "0300", "http://intel.com/wbem/wscim/1/amt-schema/1/", "AMT_ManagementPresenceRemoteSAP"), func() string {
			// 		remoteAccessPolicyRule := RemoteAccessPolicyRule{
			// 			Trigger:        2,
			// 			TunnelLifeTime: 0,
			// 			ExtendedData:   "0300",
			// 		}
			// 		selector := message.Selector{
			// 			Name:  "myselector",
			// 			Value: "true",
			// 		}
			// 		return elementUnderTest.AddRemoteAccessPolicyRule(remoteAccessPolicyRule, selector)
			// 	}},
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
