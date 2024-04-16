/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

import "testing"

func TestLocalTimeSyncEnabledString(t *testing.T) {
	tests := []struct {
		state    LocalTimeSyncEnabled
		expected string
	}{
		{LocalTimeSyncEnabledDefaultTrue, "LocalTimeSyncEnabledDefaultTrue"},
		{LocalTimeSyncEnabledConfiguredTrue, "LocalTimeSyncEnabledConfiguredTrue"},
		{LocalTimeSyncEnabledFalse, "LocalTimeSyncEnabledFalse"},
		{LocalTimeSyncEnabled(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestTimeSourceString(t *testing.T) {
	tests := []struct {
		state    TimeSource
		expected string
	}{
		{TimeSourceBiosRTC, "TimeSourceBiosRTC"},
		{TimeSourceConfigured, "TimeSourceConfigured"},
		{TimeSource(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestEnabledStateString(t *testing.T) {
	tests := []struct {
		state    EnabledState
		expected string
	}{
		{EnabledStateUnknown, "EnabledStateUnknown"},
		{EnabledStateOther, "EnabledStateOther"},
		{EnabledStateEnabled, "EnabledStateEnabled"},
		{EnabledStateDisabled, "EnabledStateDisabled"},
		{EnabledStateShuttingDown, "EnabledStateShuttingDown"},
		{EnabledStateNotApplicable, "EnabledStateNotApplicable"},
		{EnabledStateEnabledButOffline, "EnabledStateEnabledButOffline"},
		{EnabledStateInTest, "EnabledStateInTest"},
		{EnabledStateDeferred, "EnabledStateDeferred"},
		{EnabledStateQuiesce, "EnabledStateQuiesce"},
		{EnabledStateStarting, "EnabledStateStarting"},
		{EnabledState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestRequestedStateString(t *testing.T) {
	tests := []struct {
		state    RequestedState
		expected string
	}{
		{RequestedStateUnknown, "RequestedStateUnknown"},
		{RequestedStateEnabled, "RequestedStateEnabled"},
		{RequestedStateDisabled, "RequestedStateDisabled"},
		{RequestedStateShutDown, "RequestedStateShutDown"},
		{RequestedStateNoChange, "RequestedStateNoChange"},
		{RequestedStateOffline, "RequestedStateOffline"},
		{RequestedStateTest, "RequestedStateTest"},
		{RequestedStateDeferred, "RequestedStateDeferred"},
		{RequestedStateQuiesce, "RequestedStateQuiesce"},
		{RequestedStateReboot, "RequestedStateReboot"},
		{RequestedStateReset, "RequestedStateReset"},
		{RequestedStateNotApplicable, "RequestedStateNotApplicable"},
		{RequestedState(999), "Value not found in map"},
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
		{ReturnValueSuccess, "ReturnValueSuccess"},
		{ReturnValueInternalError, "ReturnValueInternalError"},
		{ReturnValueInvalidParameter, "ReturnValueInvalidParameter"},
		{ReturnValueFlashWriteLimitExceeded, "ReturnValueFlashWriteLimitExceeded"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
