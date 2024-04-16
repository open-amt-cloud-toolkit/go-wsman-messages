/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kvm

const CIM_KVMRedirectionSAP string = "CIM_KVMRedirectionSAP"

const (
	RedirectionSAP_Enable  KVMRedirectionSAPRequestStateChangeInput = 2
	RedirectionSAP_Disable KVMRedirectionSAPRequestStateChangeInput = 3
)

// KVMProtocol constants
const (
	KVMProtocolUnknown KVMProtocol = iota
	KVMProtocolOther
	KVMProtocolRaw
	KVMProtocolRDP
	KVMProtocolVNCRFB
)

// kvmProtocolToString is a map of KVMProtocol value to string
var kvmProtocolToString = map[KVMProtocol]string{
	KVMProtocolUnknown: "Unknown",
	KVMProtocolOther:   "Other",
	KVMProtocolRaw:     "Raw",
	KVMProtocolRDP:     "RDP",
	KVMProtocolVNCRFB:  "VNC-RFB",
}

// String returns a human-readable string representation of the KVMProtocol enumeration
func (e KVMProtocol) String() string {
	if s, ok := kvmProtocolToString[e]; ok {
		return s
	}
	return "Value not found in map"
}

// ReturnValue constants
const (
	ReturnValueCompletedNoError ReturnValue = iota
	ReturnValueNotSupported
	ReturnValueUnknownError
	ReturnValueTimeout
	ReturnValueFailed
	ReturnValueInvalidParameter
	ReturnValueInUse
	ReturnValueMethodParametersChecked      ReturnValue = 4096
	ReturnValueInvalidStateTransition       ReturnValue = 4097
	ReturnValueTimeoutParameterNotSupported ReturnValue = 4098
	ReturnValueBusy                         ReturnValue = 4099
)

// returnValueToString is a map of ReturnValue value to string
var returnValueToString = map[ReturnValue]string{
	ReturnValueCompletedNoError:             "CompletedWithNoError",
	ReturnValueNotSupported:                 "NotSupported",
	ReturnValueUnknownError:                 "UnknownOrUnspecifiedError",
	ReturnValueTimeout:                      "CannotCompleteWithinTimeoutPeriod",
	ReturnValueFailed:                       "Failed",
	ReturnValueInvalidParameter:             "InvalidParameter",
	ReturnValueInUse:                        "InUse",
	ReturnValueMethodParametersChecked:      "MethodParametersChecked-JobStarted",
	ReturnValueInvalidStateTransition:       "InvalidStateTransition",
	ReturnValueTimeoutParameterNotSupported: "UseOfTimeoutParameterNotSupported",
	ReturnValueBusy:                         "Busy",
}

// String returns a human-readable string representation of the ReturnValue enumeration
func (e ReturnValue) String() string {
	if s, ok := returnValueToString[e]; ok {
		return s
	}
	return "Value not found in map"
}

// EnabledState constants
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

// enabledStateToString is a map of EnabledState value to string
var enabledStateToString = map[EnabledState]string{
	EnabledStateUnknown:           "Unknown",
	EnabledStateOther:             "Other",
	EnabledStateEnabled:           "Enabled",
	EnabledStateDisabled:          "Disabled",
	EnabledStateShuttingDown:      "ShuttingDown",
	EnabledStateNotApplicable:     "NotApplicable",
	EnabledStateEnabledButOffline: "EnabledButOffline",
	EnabledStateInTest:            "InTest",
	EnabledStateDeferred:          "Deferred",
	EnabledStateQuiesce:           "Quiesce",
	EnabledStateStarting:          "Starting",
}

// String returns a human-readable string representation of the EnabledState enumeration
func (e EnabledState) String() string {
	if s, ok := enabledStateToString[e]; ok {
		return s
	}
	return "Value not found in map"
}

// RequestedState constants
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

// requestedStateToString is a map of RequestedState value to string
var requestedStateToString = map[RequestedState]string{
	RequestedStateUnknown:       "Unknown",
	RequestedStateEnabled:       "Enabled",
	RequestedStateDisabled:      "Disabled",
	RequestedStateShutDown:      "ShutDown",
	RequestedStateNoChange:      "NoChange",
	RequestedStateOffline:       "Offline",
	RequestedStateTest:          "Test",
	RequestedStateDeferred:      "Deferred",
	RequestedStateQuiesce:       "Quiesce",
	RequestedStateReboot:        "Reboot",
	RequestedStateReset:         "Reset",
	RequestedStateNotApplicable: "NotApplicable",
}

// String returns a human-readable string representation of the RequestedState enumeration
func (e RequestedState) String() string {
	if s, ok := requestedStateToString[e]; ok {
		return s
	}
	return "Value not found in map"
}
