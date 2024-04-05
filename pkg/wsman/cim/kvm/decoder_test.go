/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kvm

import "testing"

func TestKVMProtocol_String(t *testing.T) {
	tests := []struct {
		state    KVMProtocol
		expected string
	}{
		{KVMProtocolUnknown, "Unknown"},
		{KVMProtocolOther, "Other"},
		{KVMProtocolRaw, "Raw"},
		{KVMProtocolRDP, "RDP"},
		{KVMProtocolVNCRFB, "VNC-RFB"},
		{KVMProtocol(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestReturnValue_String(t *testing.T) {
	tests := []struct {
		state    ReturnValue
		expected string
	}{
		{ReturnValueCompletedNoError, "CompletedWithNoError"},
		{ReturnValueNotSupported, "NotSupported"},
		{ReturnValueUnknownError, "UnknownOrUnspecifiedError"},
		{ReturnValueTimeout, "CannotCompleteWithinTimeoutPeriod"},
		{ReturnValueFailed, "Failed"},
		{ReturnValueInvalidParameter, "InvalidParameter"},
		{ReturnValueInUse, "InUse"},
		{ReturnValueMethodParametersChecked, "MethodParametersChecked-JobStarted"},
		{ReturnValueInvalidStateTransition, "InvalidStateTransition"},
		{ReturnValueTimeoutParameterNotSupported, "UseOfTimeoutParameterNotSupported"},
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

func TestEnabledState_String(t *testing.T) {
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
		{EnabledState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestRequestedState_String(t *testing.T) {
	tests := []struct {
		state    RequestedState
		expected string
	}{
		{RequestedStateUnknown, "Unknown"},
		{RequestedStateEnabled, "Enabled"},
		{RequestedStateDisabled, "Disabled"},
		{RequestedStateShutDown, "ShutDown"},
		{RequestedStateNoChange, "NoChange"},
		{RequestedStateOffline, "Offline"},
		{RequestedStateTest, "Test"},
		{RequestedStateDeferred, "Deferred"},
		{RequestedStateQuiesce, "Quiesce"},
		{RequestedStateReboot, "Reboot"},
		{RequestedStateReset, "Reset"},
		{RequestedStateNotApplicable, "NotApplicable"},
		{RequestedState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
