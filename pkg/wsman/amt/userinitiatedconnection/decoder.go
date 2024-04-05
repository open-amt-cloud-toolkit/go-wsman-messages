/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

// INPUTS constants

const (
	AMT_UserInitiatedConnectionService string = "AMT_UserInitiatedConnectionService"
)

const (
	AllInterfacesDisabled      RequestedState = 32768
	BIOSInterfaceEnabled       RequestedState = 32769
	OSInterfaceEnabled         RequestedState = 32770
	BIOSandOSInterfacesEnabled RequestedState = 32771
)

const (
	EnabledStateUnknown                    EnabledState = 0
	EnabledStateOther                      EnabledState = 1
	EnabledStateEnabled                    EnabledState = 2
	EnabledStateDisabled                   EnabledState = 3
	EnabledStateShuttingDown               EnabledState = 4
	EnabledStateNotApplicable              EnabledState = 5
	EnabledStateEnabledButOffline          EnabledState = 6
	EnabledStateInTest                     EnabledState = 7
	EnabledStateDeferred                   EnabledState = 8
	EnabledStateQuiesce                    EnabledState = 9
	EnabledStateStarting                   EnabledState = 10
	EnabledStateAllInterfacesDisabled      EnabledState = 32768
	EnabledStateBIOSInterfaceEnabled       EnabledState = 32769
	EnabledStateOSInterfaceEnabled         EnabledState = 32770
	EnabledStateBIOSAndOSInterfacesEnabled EnabledState = 32771
)

// enabledStateToString is a map of EnabledState value to string
var enabledStateToString = map[EnabledState]string{
	EnabledStateUnknown:                    "Unknown",
	EnabledStateOther:                      "Other",
	EnabledStateEnabled:                    "Enabled",
	EnabledStateDisabled:                   "Disabled",
	EnabledStateShuttingDown:               "ShuttingDown",
	EnabledStateNotApplicable:              "NotApplicable",
	EnabledStateEnabledButOffline:          "EnabledButOffline",
	EnabledStateInTest:                     "InTest",
	EnabledStateDeferred:                   "Deferred",
	EnabledStateQuiesce:                    "Quiesce",
	EnabledStateStarting:                   "Starting",
	EnabledStateAllInterfacesDisabled:      "AllInterfacesDisabled",
	EnabledStateBIOSInterfaceEnabled:       "BIOSInterfaceEnabled",
	EnabledStateOSInterfaceEnabled:         "OSInterfaceEnabled",
	EnabledStateBIOSAndOSInterfacesEnabled: "BIOSAndOSInterfacesEnabled",
}

// String returns the string representation of the EnabledState value
func (e EnabledState) String() string {
	if value, exists := enabledStateToString[e]; exists {
		return value
	}
	return "Value not found in map"
}

const (
	ReturnValueCompletedWithNoError              ReturnValue = 0
	ReturnValueNotSupported                      ReturnValue = 1
	ReturnValueUnknownOrUnspecified              ReturnValue = 2
	ReturnValueCannotCompleteWithinTimeout       ReturnValue = 3
	ReturnValueFailed                            ReturnValue = 4
	ReturnValueInvalidParameter                  ReturnValue = 5
	ReturnValueInUse                             ReturnValue = 6
	ReturnValueMethodParametersCheckedJobStarted ReturnValue = 4096
	ReturnValueInvalidStateTransition            ReturnValue = 4097
	ReturnValueUseOfTimeoutParameterNotSupported ReturnValue = 4098
	ReturnValueBusy                              ReturnValue = 4099
)

// returnValueToString is a map of ReturnValue value to string
var returnValueToString = map[ReturnValue]string{
	ReturnValueCompletedWithNoError:              "CompletedWithNoError",
	ReturnValueNotSupported:                      "NotSupported",
	ReturnValueUnknownOrUnspecified:              "UnknownOrUnspecified",
	ReturnValueCannotCompleteWithinTimeout:       "CannotCompleteWithinTimeout",
	ReturnValueFailed:                            "Failed",
	ReturnValueInvalidParameter:                  "InvalidParameter",
	ReturnValueInUse:                             "InUse",
	ReturnValueMethodParametersCheckedJobStarted: "MethodParametersCheckedJobStarted",
	ReturnValueInvalidStateTransition:            "InvalidStateTransition",
	ReturnValueUseOfTimeoutParameterNotSupported: "UseOfTimeoutParameterNotSupported",
	ReturnValueBusy:                              "Busy",
}

// String returns the string representation of the ReturnValue value
func (r ReturnValue) String() string {
	if value, exists := returnValueToString[r]; exists {
		return value
	}
	return "Value not found in map"
}
