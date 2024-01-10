/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"encoding/xml"
	"time"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Service struct {
	base message.Base
}

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
	Interval int `json:"Interval"`
	// DeleteOnComplete if set to TRUE, the instance will be deleted by the FW when the alarm is completed
	DeleteOnCompletion bool `json:"DeleteOnCompletion"`
}

// OUTPUTS
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName           xml.Name          `xml:"Body"`
		GetResponse       AlarmClockService `xml:"AMT_AlarmClockService"`
		EnumerateResponse common.EnumerateResponse
		AddAlarmOutput    AddAlarmOutput `xml:"AddAlarm_OUTPUT"`
		PullResponse      PullResponse
	}
	PullResponse struct {
		XMLName                xml.Name            `xml:"PullResponse"`
		AlarmClockServiceItems []AlarmClockService `xml:"Items>AMT_AlarmClockService"`
	}
	AlarmClockService struct {
		XMLName xml.Name `xml:"AMT_AlarmClockService"`
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
		NextAMTAlarmTime string
		// Specifies the alarm time interval . . .
		AMTAlarmClockInterval string
	}
	AddAlarmOutput struct {
		XMLName xml.Name `xml:"AddAlarm_OUTPUT"`
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
