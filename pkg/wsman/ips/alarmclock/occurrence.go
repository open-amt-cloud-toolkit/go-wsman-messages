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
func (occurrence Occurrence) Get(alarmName string) (response Response, err error) {
	selector := message.Selector{
		Name:  "Name",
		Value: alarmName,
	}
	response = Response{
		Message: &client.Message{
			XMLInput: occurrence.base.Get(&selector),
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

// Enumerate returns an enumeration context which is used in a subsequent Pull call
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

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
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
