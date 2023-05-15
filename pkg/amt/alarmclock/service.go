/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// AMT Alarm Clock Service derived from Service and provides the ability to set an alarm time to turn the host computer system on. Setting an alarm time is done by calling "AddAlarm" method."NextAMTAlarmTime" and "AMTAlarmClockInterval" properties are deprecated and "AddAlarm" should be used instead.
package alarmclock

import (
	"strconv"
	"strings"
	"time"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
)

const AMT_AlarmClockService = "AMT_AlarmClockService"

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

type Service struct {
	base wsman.Base
}

func NewService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{base: wsman.NewBase(wsmanMessageCreator, string(AMT_AlarmClockService))}
}

// Get retrieves the representation of the instance
func (acs Service) Get() string {
	return acs.base.Get(nil)
}

// Enumerates the instances of this class
func (acs Service) Enumerate() string {
	return acs.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (acs Service) Pull(enumerationContext string) string {
	return acs.base.Pull(enumerationContext)
}

// AddAlarm creates an alarm that would wake the system at a given time.The method receives as input an embedded instance of type IPS_AlarmClockOccurrence, with the following fields set: StartTime, Interval, InstanceID, DeleteOnCompletion. Upon success, the method creates an instance of IPS_AlarmClockOccurrence which is associated with AlarmClockService.The method would fail if 5 instances or more of IPS_AlarmClockOccurrence already exist in the system.
func (acs Service) AddAlarm(alarmClockOccurrence AlarmClockOccurrence) string {
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

	return acs.base.WSManMessageCreator.CreateXML(header, body.String())
}
