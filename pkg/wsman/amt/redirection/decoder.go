/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package redirection

const (
	AMT_RedirectionService string = "AMT_RedirectionService"
	RequestStateChange     string = "RequestStateChange"
)

const (
	DisableIDERAndSOL       RequestedState = 32768
	EnableIDERAndDisableSOL RequestedState = 32769
	EnableSOLAndDisableIDER RequestedState = 32770
	EnableIDERAndSOL        RequestedState = 32771
)

const (
	Unknown EnabledState = iota
	Other
	Enabled
	Disabled
	ShuttingDown
	NotApplicable
	EnabledButOffline
	InTest
	Deferred
	Quiesce
	Starting
	DMTFReserved
	IDERAndSOLAreDisabled         EnabledState = 32768
	IDERIsEnabledAndSOLIsDisabled EnabledState = 32769
	SOLIsEnabledAndIDERIsDisabled EnabledState = 32770
	IDERAndSOLAreEnabled          EnabledState = 32771
)

var enabledStateToString = map[EnabledState]string{
	Unknown:                       "Unknown",
	Other:                         "Other",
	Enabled:                       "Enabled",
	Disabled:                      "Disabled",
	ShuttingDown:                  "ShuttingDown",
	NotApplicable:                 "NotApplicable",
	EnabledButOffline:             "EnabledButOffline",
	InTest:                        "InTest",
	Deferred:                      "Deferred",
	Quiesce:                       "Quiesce",
	Starting:                      "Starting",
	DMTFReserved:                  "DMTFReserved",
	IDERAndSOLAreDisabled:         "IDERAndSOLAreDisabled",
	IDERIsEnabledAndSOLIsDisabled: "IDERIsEnabledAndSOLIsDisabled",
	SOLIsEnabledAndIDERIsDisabled: "SOLIsEnabledAndIDERIsDisabled",
	IDERAndSOLAreEnabled:          "IDERAndSOLAreEnabled",
}

func (es EnabledState) String() string {
	if v, ok := enabledStateToString[es]; ok {
		return v
	}
	return "Value not found in map"
}

const (
	CompletedWithNoError              ReturnValue = 0
	NotSupported                      ReturnValue = 1
	UnknownOrUnspecified              ReturnValue = 2
	CannotCompleteWithinTimeoutPeriod ReturnValue = 3
	Failed                            ReturnValue = 4
	InvalidParameter                  ReturnValue = 5
	InUse                             ReturnValue = 6
	MethodParametersCheckedJobStarted ReturnValue = 4096
	InvalidStateTransition            ReturnValue = 4097
	UseOfTimeoutParameterNotSupported ReturnValue = 4098
	Busy                              ReturnValue = 4099
)

// returnValueToString is a map of ReturnValue values to string
var returnValueToString = map[ReturnValue]string{
	CompletedWithNoError:              "CompletedWithNoError",
	NotSupported:                      "NotSupported",
	UnknownOrUnspecified:              "UnknownOrUnspecified",
	CannotCompleteWithinTimeoutPeriod: "CannotCompleteWithinTimeoutPeriod",
	Failed:                            "Failed",
	InvalidParameter:                  "InvalidParameter",
	InUse:                             "InUse",
	MethodParametersCheckedJobStarted: "MethodParametersCheckedJobStarted",
	InvalidStateTransition:            "InvalidStateTransition",
	UseOfTimeoutParameterNotSupported: "UseOfTimeoutParameterNotSupported",
	Busy:                              "Busy",
}

// String returns a string representation of ReturnValue
func (rv ReturnValue) String() string {
	if v, ok := returnValueToString[rv]; ok {
		return v
	}
	return "Value not found in map"
}
