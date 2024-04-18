/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportconfiguration

import "testing"

func TestNoHostCsmeSoftwarePolicyString(t *testing.T) {
	tests := []struct {
		state    NoHostCsmeSoftwarePolicy
		expected string
	}{
		{RelaxedPolicy, "RelaxedPolicy"},
		{AggressivePolicy, "AggressivePolicy"},
		{NoHostCsmeSoftwarePolicy(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestHealthStateString(t *testing.T) {
	tests := []struct {
		state    HealthState
		expected string
	}{
		{HealthStateUnknown, "HealthStateUnknown"},
		{HealthStateOK, "HealthStateOK"},
		{HealthStateDegradedWarning, "HealthStateDegradedWarning"},
		{HealthStateMinorFailure, "HealthStateMinorFailure"},
		{HealthStateMajorFailure, "HealthStateMajorFailure"},
		{HealthStateCriticalFailure, "HealthStateCriticalFailure"},
		{HealthStateNonRecoverableError, "HealthStateNonRecoverableError"},
		{HealthState(999), "Value not found in map"},
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
		{RequestedStateUnknown, "RequestedStateUnknown"},
		{RequestedState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestLocalProfileSynchronizationEnabledString(t *testing.T) {
	tests := []struct {
		state    LocalProfileSynchronizationEnabled
		expected string
	}{
		{LocalSyncDisabled, "LocalSyncDisabled"},
		{LocalUserProfileSynchronizationEnabled, "LocalUserProfileSynchronizationEnabled"},
		{UnrestrictedSync, "UnrestrictedSync"},
		{LocalProfileSynchronizationEnabled(999), "Value not found in map"},
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
		{ReturnValueCompletedNoError, "ReturnValueCompletedNoError"},
		{ReturnValueNotSupported, "ReturnValueNotSupported"},
		{ReturnValueFailed, "ReturnValueFailed"},
		{ReturnValueInvalidParameter, "ReturnValueInvalidParameter"},
		{ReturnValueInvalidReference, "ReturnValueInvalidReference"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
