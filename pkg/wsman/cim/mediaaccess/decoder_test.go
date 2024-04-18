/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mediaaccess

import "testing"

func TestSecurity_String(t *testing.T) {
	tests := []struct {
		state    Security
		expected string
	}{
		{Other, "Other"},
		{Unknown, "Unknown"},
		{None, "None"},
		{ReadOnly, "ReadOnly"},
		{LockedOut, "LockedOut"},
		{BootBypass, "BootBypass"},
		{BootBypassAndReadOnly, "BootBypassandReadOnly"},
		{Security(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestCapabilities_String(t *testing.T) {
	tests := []struct {
		state    Capabilities
		expected string
	}{
		{UnknownCapabilities, "Unknown"},
		{OtherCapabilities, "Other"},
		{SequentialAccess, "SequentialAccess"},
		{RandomAccess, "RandomAccess"},
		{SupportsWriting, "SupportsWriting"},
		{Encryption, "Encryption"},
		{Compression, "Compression"},
		{SupportsRemovableMedia, "SupportsRemovableMedia"},
		{ManualCleaning, "ManualCleaning"},
		{AutomaticCleaning, "AutomaticCleaning"},
		{SmartNotification, "SmartNotification"},
		{SupportsDualSidedMedia, "SupportsDualSidedMedia"},
		{PreDismountEjectNotRequired, "PreDismountEjectNotRequired"},
		{Capabilities(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestEnabledDefault_String(t *testing.T) {
	tests := []struct {
		state    EnabledDefault
		expected string
	}{
		{EnabledDefaultEnabled, "Enabled"},
		{EnabledDefaultDisabled, "Disabled"},
		{EnabledDefaultNotApplicable, "NotApplicable"},
		{EnabledDefaultEnabledButOffline, "EnabledButOffline"},
		{EnabledDefaultNoDefault, "NoDefault"},
		{EnabledDefaultQuiesce, "Quiesce"},
		{EnabledDefault(999), "Value not found in map"},
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
