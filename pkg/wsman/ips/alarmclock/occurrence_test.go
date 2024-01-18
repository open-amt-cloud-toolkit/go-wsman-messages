/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveIPS_AlarmClockOccurrence(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/alarmclock",
	}
	elementUnderTest := NewAlarmClockOccurrenceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_AlarmClockOccurrence Tests", func(t *testing.T) {
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
				"should create a valid ips_AlarmClockOccurrence Get wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.GET,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">testalarm</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get("testalarm")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AlarmClockOccurrence{
						XMLName:            xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPS_AlarmClockOccurrence), Local: IPS_AlarmClockOccurrence},
						ElementName:        "testalarm",
						InstanceID:         "testalarm",
						StartTime:          "testdatetime",
						Interval:           "0",
						DeleteOnCompletion: true,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid IPS_AlarmClockOccurrence Enumerate wsman message",
				"IPS_AlarmClockOccurrence",
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
						EnumerationContext: "9C0A0000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid IPS_AlarmClockOccurrence Pull wsman message",
				"IPS_AlarmClockOccurrence",
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
						Items: []AlarmClockOccurrence{
							{
								XMLName:            xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPS_AlarmClockOccurrence), Local: IPS_AlarmClockOccurrence},
								ElementName:        "testalarm",
								InstanceID:         "testalarm",
								StartTime:          "testdatetime",
								Interval:           "0",
								DeleteOnCompletion: true,
							},
						},
					},
				},
			},
			// DELETE
			{
				"should create a valid ips_AlarmClockOccurrence Delete wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.DELETE,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">testalarm</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Delete"
					return elementUnderTest.Delete("testalarm")
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
func TestNegativeIPS_AlarmClockOccurrence(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/alarmclock",
	}
	elementUnderTest := NewAlarmClockOccurrenceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_AlarmClockOccurrence Tests", func(t *testing.T) {
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
				"should create a valid ips_AlarmClockOccurrence Get wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.GET,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">testalarm</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get("testalarm")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AlarmClockOccurrence{
						XMLName:            xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPS_AlarmClockOccurrence), Local: IPS_AlarmClockOccurrence},
						ElementName:        "testalarm",
						InstanceID:         "testalarm",
						StartTime:          "testdatetime",
						Interval:           "0",
						DeleteOnCompletion: true,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid IPS_AlarmClockOccurrence Enumerate wsman message",
				"IPS_AlarmClockOccurrence",
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
						EnumerationContext: "9C0A0000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid IPS_AlarmClockOccurrence Pull wsman message",
				"IPS_AlarmClockOccurrence",
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
						Items: []AlarmClockOccurrence{
							{
								XMLName:            xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPS_AlarmClockOccurrence), Local: IPS_AlarmClockOccurrence},
								ElementName:        "testalarm",
								InstanceID:         "testalarm",
								StartTime:          "testdatetime",
								Interval:           "0",
								DeleteOnCompletion: true,
							},
						},
					},
				},
			},
			// DELETE
			{
				"should create a valid ips_AlarmClockOccurrence Delete wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.DELETE,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">testalarm</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Delete("testalarm")
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
