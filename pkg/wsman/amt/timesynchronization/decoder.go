/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

const (
	AMTTimeSynchronizationService string = "AMT_TimeSynchronizationService"
	GetLowAccuracyTimeSynch       string = "GetLowAccuracyTimeSynch"
	SetHighAccuracyTimeSynch      string = "SetHighAccuracyTimeSynch"
	ValueNotFound                 string = "Value not found in map"
)

const (
	LocalTimeSyncEnabledDefaultTrue LocalTimeSyncEnabled = iota
	LocalTimeSyncEnabledConfiguredTrue
	LocalTimeSyncEnabledFalse
)

// localTimeSyncEnabledString is a map of LocalTimeSyncEnabled values to their string representations.
var localTimeSyncEnabledString = map[LocalTimeSyncEnabled]string{
	LocalTimeSyncEnabledDefaultTrue:    "LocalTimeSyncEnabledDefaultTrue",
	LocalTimeSyncEnabledConfiguredTrue: "LocalTimeSyncEnabledConfiguredTrue",
	LocalTimeSyncEnabledFalse:          "LocalTimeSyncEnabledFalse",
}

// String returns the string representation of the LocalTimeSyncEnabled value.
func (l LocalTimeSyncEnabled) String() string {
	if value, exists := localTimeSyncEnabledString[l]; exists {
		return value
	}

	return ValueNotFound
}

const (
	TimeSourceBiosRTC TimeSource = iota
	TimeSourceConfigured
)

// timeSourceString is a map of TimeSource values to their string representations.
var timeSourceString = map[TimeSource]string{
	TimeSourceBiosRTC:    "TimeSourceBiosRTC",
	TimeSourceConfigured: "TimeSourceConfigured",
}

// String returns the string representation of the TimeSource value.
func (t TimeSource) String() string {
	if value, exists := timeSourceString[t]; exists {
		return value
	}

	return ValueNotFound
}

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)

// enabledStateString is a map of EnabledState values to their string representations.
var enabledStateString = map[EnabledState]string{
	EnabledStateUnknown:           "EnabledStateUnknown",
	EnabledStateOther:             "EnabledStateOther",
	EnabledStateEnabled:           "EnabledStateEnabled",
	EnabledStateDisabled:          "EnabledStateDisabled",
	EnabledStateShuttingDown:      "EnabledStateShuttingDown",
	EnabledStateNotApplicable:     "EnabledStateNotApplicable",
	EnabledStateEnabledButOffline: "EnabledStateEnabledButOffline",
	EnabledStateInTest:            "EnabledStateInTest",
	EnabledStateDeferred:          "EnabledStateDeferred",
	EnabledStateQuiesce:           "EnabledStateQuiesce",
	EnabledStateStarting:          "EnabledStateStarting",
}

// String returns the string representation of the EnabledState value.
func (e EnabledState) String() string {
	if value, exists := enabledStateString[e]; exists {
		return value
	}

	return ValueNotFound
}

const (
	RequestedStateUnknown       RequestedState = 0
	RequestedStateEnabled       RequestedState = 2
	RequestedStateDisabled      RequestedState = 3
	RequestedStateShutDown      RequestedState = 4
	RequestedStateNoChange      RequestedState = 5
	RequestedStateOffline       RequestedState = 6
	RequestedStateTest          RequestedState = 7
	RequestedStateDeferred      RequestedState = 8
	RequestedStateQuiesce       RequestedState = 9
	RequestedStateReboot        RequestedState = 10
	RequestedStateReset         RequestedState = 11
	RequestedStateNotApplicable RequestedState = 12
)

// requestedStateString is a map of RequestedState values to their string representations.
var requestedStateString = map[RequestedState]string{
	RequestedStateUnknown:       "RequestedStateUnknown",
	RequestedStateEnabled:       "RequestedStateEnabled",
	RequestedStateDisabled:      "RequestedStateDisabled",
	RequestedStateShutDown:      "RequestedStateShutDown",
	RequestedStateNoChange:      "RequestedStateNoChange",
	RequestedStateOffline:       "RequestedStateOffline",
	RequestedStateTest:          "RequestedStateTest",
	RequestedStateDeferred:      "RequestedStateDeferred",
	RequestedStateQuiesce:       "RequestedStateQuiesce",
	RequestedStateReboot:        "RequestedStateReboot",
	RequestedStateReset:         "RequestedStateReset",
	RequestedStateNotApplicable: "RequestedStateNotApplicable",
}

// String returns the string representation of the RequestedState value.
func (r RequestedState) String() string {
	if value, exists := requestedStateString[r]; exists {
		return value
	}

	return ValueNotFound
}

const (
	ReturnValueSuccess                 ReturnValue = 0
	ReturnValueInternalError           ReturnValue = 1
	ReturnValueInvalidParameter        ReturnValue = 36
	ReturnValueFlashWriteLimitExceeded ReturnValue = 38
)

// returnValueString is a map of ReturnValue values to their string representations.
var returnValueString = map[ReturnValue]string{
	ReturnValueSuccess:                 "ReturnValueSuccess",
	ReturnValueInternalError:           "ReturnValueInternalError",
	ReturnValueInvalidParameter:        "ReturnValueInvalidParameter",
	ReturnValueFlashWriteLimitExceeded: "ReturnValueFlashWriteLimitExceeded",
}

// String returns the string representation of the ReturnValue value.
func (r ReturnValue) String() string {
	if value, exists := returnValueString[r]; exists {
		return value
	}

	return ValueNotFound
}
