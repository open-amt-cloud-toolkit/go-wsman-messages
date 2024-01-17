/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewAlarmClockOccurrence returns a new instance of the AlarmClockOccurrence struct.
func NewAlarmClockOccurrenceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Occurrence {
	return Occurrence{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPS_AlarmClockOccurrence, client),
	}
}

// Get retrieves the representation of the instance
func (occurrence Occurrence) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: occurrence.base.Get(nil),
		},
	}
	err = occurrence.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Delete removes a the specified instance
func (occurrence Occurrence) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "Name", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: occurrence.base.Delete(selector),
		},
	}
	err = occurrence.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (occurrence Occurrence) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: occurrence.base.Enumerate(),
		},
	}
	err = occurrence.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pulls instances of this class, following an Enumerate operation
func (occurrence Occurrence) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: occurrence.base.Pull(enumerationContext),
		},
	}
	err = occurrence.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
