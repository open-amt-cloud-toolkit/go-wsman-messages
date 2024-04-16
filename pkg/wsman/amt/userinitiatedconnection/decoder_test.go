/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

import "testing"

func TestEnabledStateString(t *testing.T) {
	tests := []struct {
		state    EnabledState
		expected string
	}{
		{EnabledStateUnknown, "Unknown"},
		{EnabledStateOther, "Other"},
		{EnabledStateEnabled, "Enabled"},
		{EnabledStateDisabled, "Disabled"},
		{EnabledStateShuttingDown, "ShuttingDown"},
		{EnabledStateNotApplicable, "NotApplicable"},
		{EnabledStateEnabledButOffline, "EnabledButOffline"},
		{EnabledStateInTest, "InTest"},
		{EnabledStateDeferred, "Deferred"},
		{EnabledStateQuiesce, "Quiesce"},
		{EnabledStateStarting, "Starting"},
		{EnabledStateAllInterfacesDisabled, "AllInterfacesDisabled"},
		{EnabledStateBIOSInterfaceEnabled, "BIOSInterfaceEnabled"},
		{EnabledStateOSInterfaceEnabled, "OSInterfaceEnabled"},
		{EnabledStateBIOSAndOSInterfacesEnabled, "BIOSAndOSInterfacesEnabled"},
		{EnabledState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestReturnValueString(t *testing.T) {
	tests := []struct {
		state    ReturnValue
		expected string
	}{
		{ReturnValueCompletedWithNoError, "CompletedWithNoError"},
		{ReturnValueNotSupported, "NotSupported"},
		{ReturnValueUnknownOrUnspecified, "UnknownOrUnspecified"},
		{ReturnValueCannotCompleteWithinTimeout, "CannotCompleteWithinTimeout"},
		{ReturnValueFailed, "Failed"},
		{ReturnValueInvalidParameter, "InvalidParameter"},
		{ReturnValueInUse, "InUse"},
		{ReturnValueMethodParametersCheckedJobStarted, "MethodParametersCheckedJobStarted"},
		{ReturnValueInvalidStateTransition, "InvalidStateTransition"},
		{ReturnValueUseOfTimeoutParameterNotSupported, "UseOfTimeoutParameterNotSupported"},
		{ReturnValueBusy, "Busy"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
