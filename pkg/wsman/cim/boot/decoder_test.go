/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import "testing"

func TestFailThroughSupported_String(t *testing.T) {
	tests := []struct {
		state    FailThroughSupported
		expected string
	}{
		{FailThroughSupportedUnknown, "Unknown"},
		{FailThroughSupportedIsSupported, "IsSupported"},
		{FailThroughSupportedNotSupported, "NotSupported"},
		{FailThroughSupported(999), "Value not found in map"},
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
		{ReturnValueCompletedNoError, "CompletedNoError"},
		{ReturnValueNotSupported, "NotSupported"},
		{ReturnValueUnknownError, "UnknownError"},
		{ReturnValueBusy, "Busy"},
		{ReturnValueInvalidReference, "InvalidReference"},
		{ReturnValueInvalidParameter, "InvalidParameter"},
		{ReturnValueAccessDenied, "AccessDenied"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestOperationalStatus_String(t *testing.T) {
	tests := []struct {
		state    OperationalStatus
		expected string
	}{
		{OperationalStatusUnknown, "Unknown"},
		{OperationalStatusOther, "Other"},
		{OperationalStatusOK, "OK"},
		{OperationalStatusDegraded, "Degraded"},
		{OperationalStatusStressed, "Stressed"},
		{OperationalStatusPredictiveFailure, "PredictiveFailure"},
		{OperationalStatusError, "Error"},
		{OperationalStatusNonRecoverableError, "NonRecoverableError"},
		{OperationalStatusStarting, "Starting"},
		{OperationalStatusStopping, "Stopping"},
		{OperationalStatusStopped, "Stopped"},
		{OperationalStatusInService, "InService"},
		{OperationalStatusNoContact, "NoContact"},
		{OperationalStatusLostCommunication, "LostCommunication"},
		{OperationalStatusAborted, "Aborted"},
		{OperationalStatusDormant, "Dormant"},
		{OperationalStatusSupportingEntityInError, "SupportingEntityInError"},
		{OperationalStatusCompleted, "Completed"},
		{OperationalStatusPowerMode, "PowerMode"},
		{OperationalStatusRelocating, "Relocating"},
		{OperationalStatus(999), "Value not found in map"},
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
		{EnabledStateIntelOneClickRecoveryAndIntelRPEAreDisabledAndAllOtherBootOptionsAreEnabled, "IntelOneClickRecoveryAndIntelRPEAreDisabledAndAllOtherBootOptionsAreEnabled"},
		{EnabledStateIntelOneClickRecoveryIsEnabledAndIntelRPEIsDisabledAndAllOtherBootOptionsAreEnabled, "IntelOneClickRecoveryIsEnabledAndIntelRPEIsDisabledAndAllOtherBootOptionsAreEnabled"},
		{EnabledStateIntelRPEIsEnabledAndIntelOneClickRecoveryIsDisabledAndAllOtherBootOptionsAreEnabled, "IntelRPEIsEnabledAndIntelOneClickRecoveryIsDisabledAndAllOtherBootOptionsAreEnabled"},
		{EnabledStateIntelOneClickRecoveryAndIntelRPEAreEnabledAndAllOtherBootOptionsAreEnabled, "IntelOneClickRecoveryAndIntelRPEAreEnabledAndAllOtherBootOptionsAreEnabled"},
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
		{RequestedStateDisableIntelOneClickRecoveryAndIntelRPEAndEnableAllOtherBootOptions, "DisableIntelOneClickRecoveryAndIntelRPEAndEnableAllOtherBootOptions"},
		{RequestedStateDisableIntelRPEAndEnableIntelOneClickRecoveryAndAllOtherBootOptions, "DisableIntelRPEAndEnableIntelOneClickRecoveryAndAllOtherBootOptions"},
		{RequestedStateDisableIntelOneClickRecoveryAndEnableIntelRPEAndAllOtherBootOptions, "DisableIntelOneClickRecoveryAndEnableIntelRPEAndAllOtherBootOptions"},
		{RequestedStateEnableAllBootOptions, "EnableAllBootOptions"},
		{RequestedState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
