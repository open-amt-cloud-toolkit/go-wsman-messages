/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

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

type SourceSettingMockClient struct{}

func (c *SourceSettingMockClient) Post(msg string) ([]byte, error) {
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/cim/boot/sourcesetting/" + strings.ToLower(currentMessage) + ".xml")
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

func TestPositiveSourceSetting(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := SourceSettingMockClient{}
	elementUnderTest := NewBootSourceSettingWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create and parse a valid cim_BootSourceSetting Get call",
				CIM_BootSourceSetting,
				wsmantesting.GET,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT: Force Hard-drive Boot</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					currentMessage = "Get"
					return elementUnderTest.Get("Intel(r) AMT: Force Hard-drive Boot")
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					SourceSettingGetResponse: BootSourceSetting{
						XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIM_BootSourceSetting},
						ElementName:          "Intel(r) AMT: Boot Source",
						InstanceID:           "Intel(r) AMT: Force Hard-drive Boot",
						StructuredBootString: "CIM:Hard-Disk:1",
						FailThroughSupported: 2,
					},
				},
			},
			//ENUMERATES
			{
				"should create and parse a valid cim_BootSourceSetting Enumerate call",
				CIM_BootSourceSetting,
				wsmantesting.ENUMERATE,
				"",
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
				"should create and parse a valid cim_BootSourceSetting Pull call",
				CIM_BootSourceSetting,
				wsmantesting.PULL,
				"",
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					currentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponse: PullResponse{
						BootSourceSettingItems: []BootSourceSetting{
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIM_BootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force Hard-drive Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:Hard-Disk:1",
								FailThroughSupported: 2,
							},
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIM_BootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force PXE Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:Network:1",
								FailThroughSupported: 2,
							},
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIM_BootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force CD/DVD Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:CD/DVD:1",
								FailThroughSupported: 2,
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedResponse := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}

func TestNegativeSourceSetting(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := SourceSettingMockClient{}
	elementUnderTest := NewBootSourceSettingWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should handle error when cim_BootSourceSetting Get call",
				CIM_BootSourceSetting,
				wsmantesting.GET,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT: Force Hard-drive Boot</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					currentMessage = "Error"
					return elementUnderTest.Get("Intel(r) AMT: Force Hard-drive Boot")
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					SourceSettingGetResponse: BootSourceSetting{
						XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIM_BootSourceSetting},
						ElementName:          "Intel(r) AMT: Boot Source",
						InstanceID:           "Intel(r) AMT: Force Hard-drive Boot",
						StructuredBootString: "CIM:Hard-Disk:1",
						FailThroughSupported: 2,
					},
				},
			},
			//ENUMERATES
			{
				"should handle error when cim_BootSourceSetting Enumerate call",
				CIM_BootSourceSetting,
				wsmantesting.ENUMERATE,
				"",
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					currentMessage = "Error"
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
				"should handle error when cim_BootSourceSetting Pull call",
				CIM_BootSourceSetting,
				wsmantesting.PULL,
				"",
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					currentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponse: PullResponse{
						BootSourceSettingItems: []BootSourceSetting{
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIM_BootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force Hard-drive Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:Hard-Disk:1",
								FailThroughSupported: 2,
							},
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIM_BootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force PXE Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:Network:1",
								FailThroughSupported: 2,
							},
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIM_BootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force CD/DVD Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:CD/DVD:1",
								FailThroughSupported: 2,
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedResponse := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedResponse, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
