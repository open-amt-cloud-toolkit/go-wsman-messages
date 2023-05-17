/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type Occurrence struct {
	base wsman.Base
}

const IPS_AlarmClockOccurrence = "IPS_AlarmClockOccurrence"

// NewAlarmClockOccurrence returns a new instance of the AlarmClockOccurrence struct.
func NewAlarmClockOccurrence(wsmanMessageCreator *wsman.WSManMessageCreator) Occurrence {
	return Occurrence{
		base: wsman.NewBase(wsmanMessageCreator, IPS_AlarmClockOccurrence),
	}
}

// Get retrieves the representation of the instance
func (a Occurrence) Get() string {
	return a.base.Get(nil)
}

// Delete removes a the specified instance
func (a Occurrence) Delete(handle string) string {
	selector := wsman.Selector{Name: "Name", Value: handle}
	return a.base.Delete(selector)
}

// Enumerates the instances of this class
func (a Occurrence) Enumerate() string {
	return a.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (a Occurrence) Pull(enumerationContext string) string {
	return a.base.Pull(enumerationContext)
}
