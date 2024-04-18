/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

// INPUTS Constants
const (
	AMT_AlarmClockService string = "AMT_AlarmClockService"
	AddAlarm              string = "AddAlarm"
)

const (
	Success ReturnValue = iota
)

// returnValueToString is a map of ReturnValue values to their string representations
var returnValueToString = map[ReturnValue]string{
	Success: "Success",
}

// String returns the string representation of the ReturnValue value
func (r ReturnValue) String() string {
	if value, exists := returnValueToString[r]; exists {
		return value
	}

	return "Value not found in map"
}
