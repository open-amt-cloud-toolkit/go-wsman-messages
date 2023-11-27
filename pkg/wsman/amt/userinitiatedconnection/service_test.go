/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

type MockClient struct {
}

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_RedirectionService><g:CreationClassName>AMT_RedirectionService</g:CreationClassName><g:ElementName>Intel(r) AMT Redirection Service</g:ElementName><g:Name>Intel(r) AMT Redirection Service</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_RedirectionService`
)

var currentMessage = ""

func (c *MockClient) Post(msg string) ([]byte, error) {
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/amt/userinitiatedconnection/" + strings.ToLower(currentMessage) + ".xml")
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
func TestAMT_UserInitiatedConnectionService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := MockClient{}
	elementUnderTest := NewUserInitiatedConnectionServiceWithClient(wsmanMessageCreator, &client)
	elementUnderTest1 := NewUserInitiatedConnectionService(wsmanMessageCreator)
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
				"should create a valid AMT_UserInitiatedConnectionService Get wsman message",
				"AMT_UserInitiatedConnectionService",
				wsmantesting.GET,
				"",
				func() (Response, error) {
					currentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					User: User{
						CreationClassName:       "AMT_UserInitiatedConnectionService",
						ElementName:             "Intel(r) AMT User Initiated Connection Service",
						EnabledState:            32771,
						Name:                    "Intel(r) AMT User Initiated Connection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_UserInitiatedConnectionService Enumerate wsman message",
				"AMT_UserInitiatedConnectionService",
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					currentMessage = "Enumerate"
					if elementUnderTest1.base.WSManMessageCreator == nil {
						print("Error")
					}
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
			{
				"should create a valid AMT_UserInitiatedConnectionService Pull wsman message", 
				"AMT_UserInitiatedConnectionService", 
				wsmantesting.PULL, 
				wsmantesting.PULL_BODY, 
				func() (Response, error) {
					currentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponse: PullResponse{
						Items: []Item{
							{
								User: User{
									CreationClassName:       "AMT_UserInitiatedConnectionService",
									ElementName:             "Intel(r) AMT User Initiated Connection Service",
									EnabledState:            32771,
									Name:                    "Intel(r) AMT User Initiated Connection Service",
									SystemCreationClassName: "CIM_ComputerSystem",
									SystemName:              "Intel(r) AMT",
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
			{
				"should create an invalid AMT_UserInitiatedConnectionService Pull wsman message", 
				"AMT_UserInitiatedConnectionService", 
				wsmantesting.PULL, 
				wsmantesting.PULL_BODY,
				"", 
				func() (Response, error) {
					currentMessage = "Error"
					response, err := elementUnderTest.Pull("")
					return response, err
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponse: PullResponse{
						Items: []Item{
							{
								User: User{
									CreationClassName:       "AMT_UserInitiatedConnectionService",
									ElementName:             "Intel(r) AMT User Initiated Connection Service",
									EnabledState:            32771,
									Name:                    "Intel(r) AMT User Initiated Connection Service",
									SystemCreationClassName: "CIM_ComputerSystem",
									SystemName:              "Intel(r) AMT",
								},
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.NotEqual(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
