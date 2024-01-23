/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publicprivate

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveAMT_PublicPrivateKeyPair(t *testing.T) {
	messageID := 0
	resourceUriBase := message.AMTSchema
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/publicprivate",
	}
	elementUnderTest := NewPublicPrivateKeyPairWithClient(wsmanMessageCreator, &client)

	t.Run("amt_PublicPrivateKeyPair Tests", func(t *testing.T) {
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
				"should create a valid AMT_PublicPrivateKeyPair Get wsman message",
				AMT_PublicPrivateKeyPair, wsmantesting.GET,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Key: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get(0)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: PublicPrivateKeyPair{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicPrivateKeyPair), Local: AMT_PublicPrivateKeyPair},
						ElementName: "Intel(r) AMT Key",
						InstanceID:  "Intel(r) AMT Key: Handle: 0",
						DERKey:      "MIIBCgKCAQEA4y00wezZ1XwsSITMvqeYf61tgfVhlGbBVwq9Au0BaEgofPFCLuWMnKaTnMhUlJEGaeB2y6F8qjId0xMwLtNY6XWhmMoCP0R+ymgClT0treqtYp2zL1QPK1R04KTgF0KZh247oQpPGnB2nIe7PKCjPaY8BfOyBC6eNLeWUVIOA5TLL0gSTuk8y3iaadKo+LoWBaH/WDrIJ21Dzn6yU3zGueA8tphPH7yXaOJuNiijOUYZjVT7J0Ia8qMxUv1CrbfL2+N0lrcCG/E4f0QF1XgoCJnwIHdYaNhWzKVhfh2TTZIxJo8bXngckNOLzdYM35hUq98CxPiMSO8+G7J8RZaobQIDAQAB",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_PublicPrivateKeyPair Enumerate wsman message",
				AMT_PublicPrivateKeyPair,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "56080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_PublicPrivateKeyPair Pull wsman message",
				AMT_PublicPrivateKeyPair,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						PublicPrivateKeyPairItems: []PublicPrivateKeyPair{
							{
								XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicPrivateKeyPair), Local: AMT_PublicPrivateKeyPair},
								ElementName: "Intel(r) AMT Key",
								InstanceID:  "Intel(r) AMT Key: Handle: 0",
								DERKey:      "MIIBCgKCAQEA4y00wezZ1XwsSITMvqeYf61tgfVhlGbBVwq9Au0BaEgofPFCLuWMnKaTnMhUlJEGaeB2y6F8qjId0xMwLtNY6XWhmMoCP0R+ymgClT0treqtYp2zL1QPK1R04KTgF0KZh247oQpPGnB2nIe7PKCjPaY8BfOyBC6eNLeWUVIOA5TLL0gSTuk8y3iaadKo+LoWBaH/WDrIJ21Dzn6yU3zGueA8tphPH7yXaOJuNiijOUYZjVT7J0Ia8qMxUv1CrbfL2+N0lrcCG/E4f0QF1XgoCJnwIHdYaNhWzKVhfh2TTZIxJo8bXngckNOLzdYM35hUq98CxPiMSO8+G7J8RZaobQIDAQAB",
							},
							{
								XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicPrivateKeyPair), Local: AMT_PublicPrivateKeyPair},
								ElementName: "Intel(r) AMT Key",
								InstanceID:  "Intel(r) AMT Key: Handle: 1",
								DERKey:      "MIIBCgKCAQEAvMgYL2FyGuHOVvwYgjABqRlJ8j8LhMo2OCU1HU2WvDN3NoLmjAh2XmBS6ic5IjIc4VtjL7S8ImKP8+PSye9nxf+lv33AqcGsvQFcUuJ5gLTnYzrmqVk6XTcHf1qtvHEmVoykTV6bN7BQx0eTejTjhw3Ro6HZBMyStaTGIKjC9HLQySV6SnFGbrjdNZZoCYsaT8dVetn23npeses9f6dZT5K3IgpA13NcdJioS71uppjIcg8dXpcxA4QKgHLmmELPN9JLbywMvcCuU+xMDceWQlFld9ohmr8NiwgebLyVCh/Q+O+jkQT43snNolyTGLRWQFR4M6DT5fdgXivoFhzMcwIDAQAB",
							},
						},
					},
				},
			},
			//DELETE
			{
				"should create a valid AMT_PublicPrivateKeyPair Delete wsman message",
				AMT_PublicPrivateKeyPair,
				wsmantesting.DELETE,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Key: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Delete"
					return elementUnderTest.Delete("Intel(r) AMT Key: Handle: 0")
				},
				Body{XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"}},
			},
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
func TestNegativeAMT_PublicPrivateKeyPair(t *testing.T) {
	messageID := 0
	resourceUriBase := message.AMTSchema
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/publicprivate",
	}
	elementUnderTest := NewPublicPrivateKeyPairWithClient(wsmanMessageCreator, &client)

	t.Run("amt_PublicPrivateKeyPair Tests", func(t *testing.T) {
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
				"should create a valid AMT_PublicPrivateKeyPair Get wsman message",
				AMT_PublicPrivateKeyPair, wsmantesting.GET,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Key: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get(0)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: PublicPrivateKeyPair{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicPrivateKeyPair), Local: AMT_PublicPrivateKeyPair},
						ElementName: "Intel(r) AMT Key",
						InstanceID:  "Intel(r) AMT Key: Handle: 0",
						DERKey:      "MIIBCgKCAQEA4y00wezZ1XwsSITMvqeYf61tgfVhlGbBVwq9Au0BaEgofPFCLuWMnKaTnMhUlJEGaeB2y6F8qjId0xMwLtNY6XWhmMoCP0R+ymgClT0treqtYp2zL1QPK1R04KTgF0KZh247oQpPGnB2nIe7PKCjPaY8BfOyBC6eNLeWUVIOA5TLL0gSTuk8y3iaadKo+LoWBaH/WDrIJ21Dzn6yU3zGueA8tphPH7yXaOJuNiijOUYZjVT7J0Ia8qMxUv1CrbfL2+N0lrcCG/E4f0QF1XgoCJnwIHdYaNhWzKVhfh2TTZIxJo8bXngckNOLzdYM35hUq98CxPiMSO8+G7J8RZaobQIDAQAB",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_PublicPrivateKeyPair Enumerate wsman message",
				AMT_PublicPrivateKeyPair,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "56080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_PublicPrivateKeyPair Pull wsman message",
				AMT_PublicPrivateKeyPair,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						PublicPrivateKeyPairItems: []PublicPrivateKeyPair{
							{
								XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicPrivateKeyPair), Local: AMT_PublicPrivateKeyPair},
								ElementName: "Intel(r) AMT Key",
								InstanceID:  "Intel(r) AMT Key: Handle: 0",
								DERKey:      "MIIBCgKCAQEA4y00wezZ1XwsSITMvqeYf61tgfVhlGbBVwq9Au0BaEgofPFCLuWMnKaTnMhUlJEGaeB2y6F8qjId0xMwLtNY6XWhmMoCP0R+ymgClT0treqtYp2zL1QPK1R04KTgF0KZh247oQpPGnB2nIe7PKCjPaY8BfOyBC6eNLeWUVIOA5TLL0gSTuk8y3iaadKo+LoWBaH/WDrIJ21Dzn6yU3zGueA8tphPH7yXaOJuNiijOUYZjVT7J0Ia8qMxUv1CrbfL2+N0lrcCG/E4f0QF1XgoCJnwIHdYaNhWzKVhfh2TTZIxJo8bXngckNOLzdYM35hUq98CxPiMSO8+G7J8RZaobQIDAQAB",
							},
							{
								XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMT_PublicPrivateKeyPair), Local: AMT_PublicPrivateKeyPair},
								ElementName: "Intel(r) AMT Key",
								InstanceID:  "Intel(r) AMT Key: Handle: 1",
								DERKey:      "MIIBCgKCAQEAvMgYL2FyGuHOVvwYgjABqRlJ8j8LhMo2OCU1HU2WvDN3NoLmjAh2XmBS6ic5IjIc4VtjL7S8ImKP8+PSye9nxf+lv33AqcGsvQFcUuJ5gLTnYzrmqVk6XTcHf1qtvHEmVoykTV6bN7BQx0eTejTjhw3Ro6HZBMyStaTGIKjC9HLQySV6SnFGbrjdNZZoCYsaT8dVetn23npeses9f6dZT5K3IgpA13NcdJioS71uppjIcg8dXpcxA4QKgHLmmELPN9JLbywMvcCuU+xMDceWQlFld9ohmr8NiwgebLyVCh/Q+O+jkQT43snNolyTGLRWQFR4M6DT5fdgXivoFhzMcwIDAQAB",
							},
						},
					},
				},
			},
			//DELETE
			{
				"should create a valid AMT_PublicPrivateKeyPair Delete wsman message",
				AMT_PublicPrivateKeyPair,
				wsmantesting.DELETE,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Key: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Delete("Intel(r) AMT Key: Handle: 0")
				},
				Body{XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"}},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
