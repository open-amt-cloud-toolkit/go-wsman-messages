/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"strconv"
	"strings"
	"time"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

const AMT_AlarmClockService = "AMT_AlarmClockService"

type AlarmClockOccurrence struct {
	models.ManagedElement
	InstanceID         string    `json:"InstanceID"`
	StartTime          time.Time `json:"StartTime"`
	Interval           int       `json:"Interval,omitempty"`
	DeleteOnCompletion bool      `json:"DeleteOnCompletion"`
}

type Service struct {
	base wsman.Base
}

func NewService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{base: wsman.NewBase(wsmanMessageCreator, string(AMT_AlarmClockService))}
}

func (acs Service) Get() string {
	return acs.base.Get(nil)
}
func (acs Service) Enumerate() string {
	return acs.base.Enumerate()
}
func (acs Service) Pull(enumerationContext string) string {
	return acs.base.Pull(enumerationContext)
}

func (acs Service) AddAlarm(alarmClockOccurrence AlarmClockOccurrence) string {
	header := acs.base.WSManMessageCreator.CreateHeader(string(actions.AddAlarm), string(AMT_AlarmClockService), nil, "", "")
	startTime := alarmClockOccurrence.StartTime.UTC().Format(time.RFC3339Nano)
	startTime = strings.Split(startTime, ".")[0] + "Z"

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
