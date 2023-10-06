/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package environmentdetection

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_EnvironmentDetectionSettingData(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	//client := MockClient{} // wsman.NewClient("http://localhost:16992/wsman", "admin", "P@ssw0rd", true)
	//elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)
	// enumerationId := ""
	client := wsman.NewClient("http://localhost:16992/wsman", "admin", "Intel123!", true)
	elementUnderTest := NewEnvironmentDetectionSettingDataWithClient(wsmanMessageCreator, client)
	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         		string
			method       		string
			action       		string
			body         		string
			responseFunc 	 	func() (Response, error)
			expectedResponse 	interface{}
		}{
			//GETS
			{
				"should create a valid AMT_EnvironmentDetectionSettingData Get wsman message", 
				"AMT_EnvironmentDetectionSettingData", 
				wsmantesting.GET, 
				"", 
				func() (Response, error) {
					//client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					DetectionSettingData: DetectionSettingData{
						DetectionStrings: "9fda2ad1-2d48-4853-8bd7-b09fdaf899de.com",
						ElementName: "Intel(r) AMT Environment Detection Settings",
						InstanceID:  "Intel(r) AMT Environment Detection Settings",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_EnvironmentDetectionSettingData Enumerate wsman message", 
				"AMT_EnvironmentDetectionSettingData", 
				wsmantesting.ENUMERATE, 
				wsmantesting.ENUMERATE_BODY, 
				func() (Response, error) {
					//client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				}, 
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "61000000-0000-0000-0000-000000000000",
					},
				},
			},
			// //PULLS
			// {"should create a valid AMT_EnvironmentDetectionSettingData Pull wsman message", "AMT_EnvironmentDetectionSettingData", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
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
