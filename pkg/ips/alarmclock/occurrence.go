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

func (a Occurrence) Get() string {
	return a.base.Get(nil)
}

func (a Occurrence) Delete(selector *wsman.Selector) string {
	return a.base.Delete(selector)
}

func (a Occurrence) Enumerate() string {
	return a.base.Enumerate()
}

func (a Occurrence) Pull(enumerationContext string) string {
	return a.base.Pull(enumerationContext)
}
