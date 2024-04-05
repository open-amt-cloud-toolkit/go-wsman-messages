/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package processor

import "testing"

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

func TestHealthState_String(t *testing.T) {
	tests := []struct {
		state    HealthState
		expected string
	}{
		{HealthStateOK, "OK"},
		{HealthStateDegradedWarning, "DegradedWarning"},
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

func TestEnabledState_String(t *testing.T) {
	tests := []struct {
		state    EnabledState
		expected string
	}{
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

func TestUpgradeMethod_String(t *testing.T) {
	tests := []struct {
		state    UpgradeMethod
		expected string
	}{
		{UpgradeMethodOther, "Other"},
		{UpgradeMethodUnknown, "Unknown"},
		{UpgradeMethodDaughterBoard, "DaughterBoard"},
		{UpgradeMethodZIFSocket, "ZIFSocket"},
		{UpgradeMethodReplacementPiggyBack, "ReplacementPiggyBack"},
		{UpgradeMethodNone, "None"},
		{UpgradeMethodLIFSocket, "LIFSocket"},
		{UpgradeMethodSlot1, "Slot1"},
		{UpgradeMethodSlot2, "Slot2"},
		{UpgradeMethod370PinSocket, "370PinSocket"},
		{UpgradeMethodSlotA, "SlotA"},
		{UpgradeMethodSlotM, "SlotM"},
		{UpgradeMethodSocket423, "Socket423"},
		{UpgradeMethodSocketASocket462, "SocketASocket462"},
		{UpgradeMethodSocket478, "Socket478"},
		{UpgradeMethodSocket754, "Socket754"},
		{UpgradeMethodSocket940, "Socket940"},
		{UpgradeMethodSocket939, "Socket939"},
		{UpgradeMethodSocketmPGA604, "SocketmPGA604"},
		{UpgradeMethodSocketLGA771, "SocketLGA771"},
		{UpgradeMethodSocketLGA775, "SocketLGA775"},
		{UpgradeMethodSocketS1, "SocketS1"},
		{UpgradeMethodSocketAM2, "SocketAM2"},
		{UpgradeMethodSocketF1207, "SocketF1207"},
		{UpgradeMethodSocketLGA1366, "SocketLGA1366"},
		{UpgradeMethodSocketG34, "SocketG34"},
		{UpgradeMethodSocketAM3, "SocketAM3"},
		{UpgradeMethodSocketC32, "SocketC32"},
		{UpgradeMethodSocketLGA1156, "SocketLGA1156"},
		{UpgradeMethodSocketLGA1567, "SocketLGA1567"},
		{UpgradeMethodSocketPGA988A, "SocketPGA988A"},
		{UpgradeMethodSocketBGA1288, "SocketBGA1288"},
		{UpgradeMethodrPGA988B, "rPGA988B"},
		{UpgradeMethodBGA1023, "BGA1023"},
		{UpgradeMethodBGA1224, "BGA1224"},
		{UpgradeMethodLGA1155, "LGA1155"},
		{UpgradeMethodLGA1356, "LGA1356"},
		{UpgradeMethodLGA2011, "LGA2011"},
		{UpgradeMethodSocketFS1, "SocketFS1"},
		{UpgradeMethodSocketFS2, "SocketFS2"},
		{UpgradeMethodSocketFM1, "SocketFM1"},
		{UpgradeMethodSocketFM2, "SocketFM2"},
		{UpgradeMethodSocketLGA20113, "SocketLGA20113"},
		{UpgradeMethodSocketLGA13563, "SocketLGA13563"},
		{UpgradeMethodSocketLGA1150, "SocketLGA1150"},
		{UpgradeMethodSocketBGA1168, "SocketBGA1168"},
		{UpgradeMethodSocketBGA1234, "SocketBGA1234"},
		{UpgradeMethodSocketBGA1364, "SocketBGA1364"},
		{UpgradeMethodSocketAM4, "SocketAM4"},
		{UpgradeMethodSocketLGA1151, "SocketLGA1151"},
		{UpgradeMethodSocketBGA1356, "SocketBGA1356"},
		{UpgradeMethodSocketBGA1440, "SocketBGA1440"},
		{UpgradeMethodSocketBGA1515, "SocketBGA1515"},
		{UpgradeMethodSocketLGA36471, "SocketLGA36471"},
		{UpgradeMethodSocketSP3, "SocketSP3"},
		{UpgradeMethodSocketSP3r2, "SocketSP3r2"},
		{UpgradeMethodSocketLGA2066, "SocketLGA2066"},
		{UpgradeMethodSocketBGA1392, "SocketBGA1392"},
		{UpgradeMethodSocketBGA1510, "SocketBGA1510"},
		{UpgradeMethodSocketBGA1528, "SocketBGA1528"},
		{UpgradeMethodSocketLGA4189, "SocketLGA4189"},
		{UpgradeMethodSocketLGA1200, "SocketLGA1200"},
		{UpgradeMethodSocketLGA4677, "SocketLGA4677"},
		{UpgradeMethodSocketLGA1700, "SocketLGA1700"},
		{UpgradeMethodSocketBGA1744, "SocketBGA1744"},
		{UpgradeMethodSocketBGA1781, "SocketBGA1781"},
		{UpgradeMethodSocketBGA1211, "SocketBGA1211"},
		{UpgradeMethodSocketBGA2422, "SocketBGA2422"},
		{UpgradeMethodSocketLGA5773, "SocketLGA5773"},
		{UpgradeMethodSocketBGA5773, "SocketBGA5773"},
		{UpgradeMethodSocketAM5, "SocketAM5"},
		{UpgradeMethodSocketSP5, "SocketSP5"},
		{UpgradeMethodSocketSP6, "SocketSP6"},
		{UpgradeMethodSocketBGA883, "SocketBGA883"},
		{UpgradeMethodSocketBGA1190, "SocketBGA1190"},
		{UpgradeMethodSocketBGA4129, "SocketBGA4129"},
		{UpgradeMethodSocketLGA4710, "SocketLGA4710"},
		{UpgradeMethodSocketLGA7529, "SocketLGA7529"},
		{UpgradeMethodSocketBGA1964, "SocketBGA1964"},
		{UpgradeMethodSocketBGA1792, "SocketBGA1792"},
		{UpgradeMethodSocketBGA2049, "SocketBGA2049"},
		{UpgradeMethodSocketBGA2551, "SocketBGA2551"},
		{UpgradeMethodSocketLGA1851, "SocketLGA1851"},
		{UpgradeMethodSocketBGA2114, "SocketBGA2114"},
		{UpgradeMethodSocketBGA2833, "SocketBGA2833"},
		{UpgradeMethod(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestCPUStatus_String(t *testing.T) {
	tests := []struct {
		state    CPUStatus
		expected string
	}{
		{CPUStatusUnknown, "Unknown"},
		{CPUStatusCPUEnabled, "CPUEnabled"},
		{CPUStatusCPUDisabledByUser, "CPUDisabledByUser"},
		{CPUStatusCPUDisabledByBIOS, "CPUDisabledByBIOS"},
		{CPUStatusCPUIsIdle, "CPUIsIdle"},
		{CPUStatusOther, "Other"},
		{CPUStatus(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
