/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package messagelog

import "testing"

func TestCapabilities_String(t *testing.T) {
	tests := []struct {
		state    Capabilities
		expected string
	}{
		{CapabilitiesUnknown, "Unknown"},
		{CapabilitiesOther, "Other"},
		{CapabilitiesWriteRecordSupported, "WriteRecordSupported"},
		{CapabilitiesDeleteRecordSupported, "DeleteRecordSupported"},
		{CapabilitiesCanMoveBackwardInLog, "CanMoveBackwardInLog"},
		{CapabilitiesFreezeLogSupported, "FreezeLogSupported"},
		{CapabilitiesClearLogSupported, "ClearLogSupported"},
		{CapabilitiesSupportsAddressingByOrdinalRecordNumber, "SupportsAddressingByOrdinalRecordNumber"},
		{CapabilitiesVariableLengthRecordsSupported, "VariableLengthRecordsSupported"},
		{CapabilitiesVariableFormatsForRecords, "VariableFormatsForRecords"},
		{CapabilitiesCanFlagRecordsForOverwrite, "CanFlagRecordsForOverwrite"},
		{Capabilities(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestCharacterSet_String(t *testing.T) {
	tests := []struct {
		state    CharacterSet
		expected string
	}{
		{CharacterSetUnknown, "Unknown"},
		{CharacterSetOther, "Other"},
		{CharacterSetASCII, "ASCII"},
		{CharacterSetUnicode, "Unicode"},
		{CharacterSetISO2022, "ISO2022"},
		{CharacterSetISO8859, "ISO8859"},
		{CharacterSetExtendedUNIXCode, "ExtendedUNIXCode"},
		{CharacterSetUTF8, "UTF8"},
		{CharacterSetUCS2, "UCS2"},
		{CharacterSetBitmappedData, "BitmappedData"},
		{CharacterSetOctetString, "OctetString"},
		{CharacterSetDefinedByIndividualRecords, "DefinedByIndividualRecords"},
		{CharacterSet(999), "Value not found in map"},
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

func TestHealthState_String(t *testing.T) {
	tests := []struct {
		state    HealthState
		expected string
	}{
		{HealthStateUnknown, "Unknown"},
		{HealthStateOK, "OK"},
		{HealthStateDegraded, "Degraded"},
		{HealthStateMinorFailure, "MinorFailure"},
		{HealthStateMajorFailure, "MajorFailure"},
		{HealthStateCriticalFailure, "CriticalFailure"},
		{HealthStateNonRecoverableError, "NonRecoverableError"},
		{HealthState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestLastChange_String(t *testing.T) {
	tests := []struct {
		state    LastChange
		expected string
	}{
		{LastChangeUnknown, "Unknown"},
		{LastChangeAdd, "Add"},
		{LastChangeDelete, "Delete"},
		{LastChangeModify, "Modify"},
		{LastChangeLogCleared, "LogCleared"},
		{LastChange(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestLogState_String(t *testing.T) {
	tests := []struct {
		state    LogState
		expected string
	}{
		{LogStateUnknown, "Unknown"},
		{LogStateNormal, "Normal"},
		{LogStateErasing, "Erasing"},
		{LogStateNotApplicable, "NotApplicable"},
		{LogState(999), "Value not found in map"},
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

func TestOverwritePolicy_String(t *testing.T) {
	tests := []struct {
		state    OverwritePolicy
		expected string
	}{
		{OverwritePolicyUnknown, "Unknown"},
		{OverwritePolicyOther, "Other"},
		{OverwritePolicyWrapWhenFull, "WrapWhenFull"},
		{OverwritePolicyClearWhenFull, "ClearWhenFull"},
		{OverwritePolicyOverwriteOutdatedWhenNeeded, "OverwriteOutdatedWhenNeeded"},
		{OverwritePolicyRemoveOutdatedRecords, "RemoveOutdatedRecords"},
		{OverwritePolicyOverwriteSpecificRecords, "OverwriteSpecificRecords"},
		{OverwritePolicyNeverOverwrite, "NeverOverwrite"},
		{OverwritePolicy(999), "Value not found in map"},
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

func TestGetRecordsReturnValue_String(t *testing.T) {
	tests := []struct {
		state    GetRecordsReturnValue
		expected string
	}{
		{GetRecordsReturnValueCompletedWithNoError, "CompletedWithNoError"},
		{GetRecordsReturnValueNotSupported, "NotSupported"},
		{GetRecordsReturnValueInvalidRecordPointed, "InvalidRecordPointed"},
		{GetRecordsReturnValueNoRecordExistsInLog, "NoRecordExistsInLog"},
		{GetRecordsReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestPositionToFirstRecordReturnValue_String(t *testing.T) {
	tests := []struct {
		state    PositionToFirstRecordReturnValue
		expected string
	}{
		{PositionToFirstRecordReturnValueCompletedWithNoError, "CompletedWithNoError"},
		{PositionToFirstRecordReturnValueNotSupported, "NotSupported"},
		{PositionToFirstRecordReturnValueNoRecordExists, "NoRecordExists"},
		{PositionToFirstRecordReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
