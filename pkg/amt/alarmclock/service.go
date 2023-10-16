/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// AMT Alarm Clock Service derived from Service and provides the ability to set an alarm time to turn the host computer system on. Setting an alarm time is done by calling "AddAlarm" method."NextAMTAlarmTime" and "AMTAlarmClockInterval" properties are deprecated and "AddAlarm" should be used instead.
package alarmclock

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"
	"time"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
)

const AMT_AlarmClockService = "AMT_AlarmClockService"

// INPUTS

// AlarmClockOccurrence represents a single alarm clock setting
type AlarmClockOccurrence struct {
	// Elementname is a user-friendly name for the object
	ElementName string `json:"ElementName"`
	// InstanceID is the instance key, set by the caller of AMT_AlarmClockService.AddAlarm.
	InstanceID string `json:"InstanceID"`
	// StartTime is the next time when the alarm is scheduled to be set.
	StartTime time.Time `json:"StartTime"`
	// Interval between occurrences of the alarm (0 if the alarm is scheduled to run once).
	Interval int `json:"Interval,omitempty"`
	// DeleteOnComplete if set to TRUE, the instance will be deleted by the FW when the alarm is completed
	DeleteOnCompletion bool `json:"DeleteOnCompletion"`
}

// OUTPUTS
type (
	Response struct {
		*wsman.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName           xml.Name          `xml:"Body"`
		AlarmClockService AlarmClockService `xml:"AMT_AlarmClockService"`
		EnumerateResponse common.EnumerateResponse
		AddAlarmOutput    AddAlarmOutput `xml:"AddAlarm_OUTPUT"`
		PullResponse      PullResponse
	}
	PullResponse struct {
		Items []Item
	}
	Item struct {
		AlarmClockService AlarmClockService `xml:"AMT_AlarmClockService"`
	}
	AlarmClockService struct {
		// The Name property uniquely identifies the Service and provides an indication of the functionality that is managed . . .
		Name string
		// CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance . . .
		CreationClassName string
		// The Name of the scoping System.
		SystemName string
		// The CreationClassName of the scoping System.
		SystemCreationClassName string
		// A user-friendly name for the object . . .
		ElementName string
		// Specifies the next AMT alarm time . . .
		NextAMTAlarmTime time.Time
		// Specifies the alarm time interval . . .
		AMTAlarmClockInterval time.Time
	}
	AddAlarmOutput struct {
		// A reference to the created instance of IPS_AlarmClockOccurrence.
		AlarmClock AlarmClock
		// Return code. 0 indicates success
		ReturnValue int
	}
	AlarmClock struct {
		// Reference address to the created instance of IPS_AlarmClockOccurrence
		Address             string
		ReferenceParameters models.ReferenceParameters_OUTPUT
	}
)

func (w *Response) JSON() string {
	jsonOutput, err := json.Marshal(w.Body)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}

type Service struct {
	base   message.Base
	client wsman.WSManClient
}

func NewService(wsmanMessageCreator *message.WSManMessageCreator) Service {
	return Service{
		base:   message.NewBase(wsmanMessageCreator, AMT_AlarmClockService),
		client: nil,
	}
}
func NewServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client wsman.WSManClient) Service {
	return Service{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_AlarmClockService, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (acs Service) Get() (response Response, err error) {

	response = Response{
		Message: &wsman.Message{
			XMLInput: acs.base.Get(nil),
		},
	}

	// send the message to AMT
	err = acs.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerates the instances of this class
func (acs Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &wsman.Message{
			XMLInput: acs.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = acs.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pulls instances of this class, following an Enumerate operation
func (acs Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &wsman.Message{
			XMLInput: acs.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = acs.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// AddAlarm creates an alarm that would wake the system at a given time.The method receives as input an embedded instance of type IPS_AlarmClockOccurrence, with the following fields set: StartTime, Interval, InstanceID, DeleteOnCompletion. Upon success, the method creates an instance of IPS_AlarmClockOccurrence which is associated with AlarmClockService.The method would fail if 5 instances or more of IPS_AlarmClockOccurrence already exist in the system.
func (acs Service) AddAlarm(alarmClockOccurrence AlarmClockOccurrence) (response Response, err error) {
	header := acs.base.WSManMessageCreator.CreateHeader(string(actions.AddAlarm), string(AMT_AlarmClockService), nil, "", "")
	startTime := alarmClockOccurrence.StartTime.UTC().Format(time.RFC3339Nano)
	startTime = strings.Split(startTime, ".")[0]

	var body strings.Builder
	body.WriteString(`<Body><p:AddAlarm_INPUT xmlns:p="`)
	body.WriteString(acs.base.WSManMessageCreator.ResourceURIBase)
	body.WriteString(`AMT_AlarmClockService"><p:AlarmTemplate><s:InstanceID xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">`)
	body.WriteString(alarmClockOccurrence.InstanceID)
	body.WriteString(`</s:InstanceID>`)

	if alarmClockOccurrence.ElementName != "" {
		body.WriteString(`<s:ElementName xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">`)
		body.WriteString(alarmClockOccurrence.ElementName)
		body.WriteString(`</s:ElementName>`)
	}

	body.WriteString(`<s:StartTime xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence"><p:Datetime xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">`)
	body.WriteString(startTime)
	body.WriteString(`</p:Datetime></s:StartTime>`)

	if alarmClockOccurrence.Interval != 0 {
		minutes := alarmClockOccurrence.Interval % 60
		hours := (alarmClockOccurrence.Interval / 60) % 24
		days := alarmClockOccurrence.Interval / 1440

		body.WriteString(`<s:Interval xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence"><p:Interval xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">P`)
		body.WriteString(strconv.Itoa(days))
		body.WriteString("DT")
		body.WriteString(strconv.Itoa(hours))
		body.WriteString("H")
		body.WriteString(strconv.Itoa(minutes))
		body.WriteString(`M</p:Interval></s:Interval>`)
	}

	body.WriteString(`<s:DeleteOnCompletion xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">`)
	body.WriteString(strconv.FormatBool(alarmClockOccurrence.DeleteOnCompletion))
	body.WriteString(`</s:DeleteOnCompletion></p:AlarmTemplate></p:AddAlarm_INPUT></Body>`)

	response = Response{
		Message: &wsman.Message{
			XMLInput: acs.base.WSManMessageCreator.CreateXML(header, body.String()),
		},
	}
	// send the message to AMT
	err = acs.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
