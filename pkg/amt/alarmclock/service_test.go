/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
	"github.com/stretchr/testify/assert"
)

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_AlarmClockService><g:CreationClassName>AMT_AlarmClockService</g:CreationClassName><g:ElementName>Intel(r) AMT Alarm Clock Service</g:ElementName><g:Name>Intel(r) AMT Alarm Clock Service</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_AlarmClockService>`
)

func TestAMT_AlarmClockService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/alarmclock",
	} // wsman.NewClient("http://localhost:16992/wsman", "admin", "P@ssw0rd", true)
	elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)
	// enumerationId := ""
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
			{"should create a valid AMT_AlarmClockService Get wsman message",
				"AMT_AlarmClockService",
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					AlarmClockService: AlarmClockService{
						CreationClassName:       "AMT_AlarmClockService",
						ElementName:             "Intel(r) AMT Alarm Clock Service",
						Name:                    "Intel(r) AMT Alarm Clock Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
					},
				},
			},
			//ENUMERATES
			{"should create a valid AMT_AlarmClockService Enumerate wsman message",
				"AMT_AlarmClockService",
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "8A000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{"should create a valid AMT_AlarmClockService Pull wsman message",
				"AMT_AlarmClockService",
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					PullResponse: PullResponse{
						Items: []Item{
							{
								AlarmClockService: AlarmClockService{
									Name:                    "Intel(r) AMT Alarm Clock Service",
									CreationClassName:       "AMT_AlarmClockService",
									SystemName:              "ManagedSystem",
									SystemCreationClassName: "CIM_ComputerSystem",
									ElementName:             "Intel(r) AMT Alarm Clock Service",
								},
							},
						},
					},
				},
			},
			//AddAlarm
			{"should create a valid AMT_AlarmClockService AddAlarm wsman message",
				"AMT_AlarmClockService",
				string(actions.AddAlarm),
				`<p:AddAlarm_INPUT xmlns:p="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService"><p:AlarmTemplate><s:InstanceID xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">Instance</s:InstanceID><s:ElementName xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">Alarm instance name</s:ElementName><s:StartTime xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence"><p:Datetime xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">2022-12-31T23:59:00Z</p:Datetime></s:StartTime><s:Interval xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence"><p:Interval xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">P1DT23H59M</p:Interval></s:Interval><s:DeleteOnCompletion xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">true</s:DeleteOnCompletion></p:AlarmTemplate></p:AddAlarm_INPUT>`,
				func() (Response, error) {
					currentMessage = "AddAlarm"
					startTime := "2022-12-31T23:59:00Z"
					minutes := 59
					hours := 23
					days := 1
					interval := minutes + hours*60 + days*1440

					startTimeFormatted, _ := time.Parse(time.RFC3339, startTime)
					return elementUnderTest.AddAlarm(AlarmClockOccurrence{
						InstanceID:         "Instance",
						StartTime:          startTimeFormatted,
						ElementName:        "Alarm instance name",
						Interval:           interval,
						DeleteOnCompletion: true,
					})
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					AddAlarmOutput: AddAlarmOutput{
						AlarmClock: AlarmClock{
							Address: "default",
							ReferenceParameters: models.ReferenceParameters_OUTPUT{
								ResourceURI: "",
								SelectorSet: models.SelectorSet_OUTPUT{
									XMLName: xml.Name{
										Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
										Local: "SelectorSet",
									},
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
