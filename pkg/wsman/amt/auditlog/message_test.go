/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package auditlog

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveAMT_AuditLog(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/auditlog",
	}
	elementUnderTest := NewAuditLogWithClient(wsmanMessageCreator, &client)

	t.Run("amt_AuditLog Tests", func(t *testing.T) {
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
				"should create a valid AMT_AuditLog Get wsman message",
				AMT_AuditLog,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AuditLog{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "AMT_AuditLog"},
						OverwritePolicy:        OverwritePolicyWrapsWhenFull,
						CurrentNumberOfRecords: 161,
						MaxNumberOfRecords:     0,
						ElementName:            "Intel(r) AMT Audit Log",
						EnabledState:           EnabledStateEnabled,
						RequestedState:         RequestedStateEnabled,
						PercentageFree:         92,
						Name:                   "Intel(r) AMT:Audit Log",
						TimeOfLastRecord: Datetime{
							Datetime: "2024-01-03T00:44:35Z",
						},
						AuditState:         16,
						MaxAllowedAuditors: 1,
						StoragePolicy:      StoragePolicyRollOver,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_AuditLog Enumerate wsman message",
				AMT_AuditLog,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "92070000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_AuditLog Pull wsman message",
				AMT_AuditLog,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						AuditLogItems: []AuditLog{
							{
								XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "AMT_AuditLog"},
								OverwritePolicy:        OverwritePolicyWrapsWhenFull,
								CurrentNumberOfRecords: 162,
								MaxNumberOfRecords:     0,
								ElementName:            "Intel(r) AMT Audit Log",
								EnabledState:           EnabledStateEnabled,
								RequestedState:         RequestedStateEnabled,
								PercentageFree:         92,
								Name:                   "Intel(r) AMT:Audit Log",
								TimeOfLastRecord: Datetime{
									Datetime: "2024-01-03T00:45:41Z",
								},
								AuditState:         16,
								MaxAllowedAuditors: 1,
								StoragePolicy:      StoragePolicyRollOver,
							},
						},
					},
				},
			},
			//READ RECORDS
			{
				"should create a valid AMT_AuditLog Read Records wsman message",
				AMT_AuditLog,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog/ReadRecords`,
				`<h:ReadRecords_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog"><h:StartIndex>1</h:StartIndex></h:ReadRecords_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "ReadRecords"
					return elementUnderTest.ReadRecords(1)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ReadRecordsResponse: ReadRecords_OUTPUT{
						XMLName:          xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "ReadRecords_OUTPUT"},
						TotalRecordCount: 2,
						RecordsReturned:  2,
						EventRecords:     []string{"ABMAAAI/9M1uAgAQAAwAAAAoBZkADAAAACQFlA==", "ABMAAAI/9M1vAgAQAAwAAAAoBZkADAAAACQFlA=="},
						ReturnValue:      0,
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
func TestNegativeAMT_AuditLog(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/auditlog",
	}
	elementUnderTest := NewAuditLogWithClient(wsmanMessageCreator, &client)

	t.Run("amt_AuditLog Tests", func(t *testing.T) {
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
				"should create a valid AMT_AuditLog Get wsman message",
				AMT_AuditLog,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AuditLog{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "AMT_AuditLog"},
						OverwritePolicy:        OverwritePolicyWrapsWhenFull,
						CurrentNumberOfRecords: 161,
						MaxNumberOfRecords:     0,
						ElementName:            "Intel(r) AMT Audit Log",
						EnabledState:           EnabledStateEnabled,
						RequestedState:         RequestedStateEnabled,
						PercentageFree:         92,
						Name:                   "Intel(r) AMT:Audit Log",
						TimeOfLastRecord: Datetime{
							Datetime: "2024-01-03T00:44:35Z",
						},
						AuditState:         16,
						MaxAllowedAuditors: 1,
						StoragePolicy:      StoragePolicyRollOver,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_AuditLog Enumerate wsman message",
				AMT_AuditLog,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "92070000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_AuditLog Pull wsman message",
				AMT_AuditLog,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						AuditLogItems: []AuditLog{
							{
								XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "AMT_AuditLog"},
								OverwritePolicy:        OverwritePolicyWrapsWhenFull,
								CurrentNumberOfRecords: 162,
								MaxNumberOfRecords:     0,
								ElementName:            "Intel(r) AMT Audit Log",
								EnabledState:           EnabledStateEnabled,
								RequestedState:         RequestedStateEnabled,
								PercentageFree:         92,
								Name:                   "Intel(r) AMT:Audit Log",
								TimeOfLastRecord: Datetime{
									Datetime: "2024-01-03T00:45:41Z",
								},
								AuditState:         16,
								MaxAllowedAuditors: 1,
								StoragePolicy:      StoragePolicyRollOver,
							},
						},
					},
				},
			},
			//READ RECORDS
			{
				"should create a valid AMT_AuditLog Read Records wsman message",
				AMT_AuditLog,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog/ReadRecords`,
				`<h:ReadRecords_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog"><h:StartIndex>1</h:StartIndex></h:ReadRecords_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.ReadRecords(1)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ReadRecordsResponse: ReadRecords_OUTPUT{
						XMLName:          xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "ReadRecords_OUTPUT"},
						TotalRecordCount: 2,
						RecordsReturned:  2,
						EventRecords:     []string{"ABMAAAI/9M1uAgAQAAwAAAAoBZkADAAAACQFlA==", "ABMAAAI/9M1vAgAQAAwAAAAoBZkADAAAACQFlA=="},
						ReturnValue:      0,
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
