/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package service

import "testing"

func TestAvailableRequestedPowerStates_String(t *testing.T) {
	tests := []struct {
		state    AvailableRequestedPowerStates
		expected string
	}{
		{AvailableRequestedPowerStatesOther, "Other"},
		{AvailableRequestedPowerStatesOn, "On"},
		{AvailableRequestedPowerStatesSleepLight, "SleepLight"},
		{AvailableRequestedPowerStatesSleepDeep, "SleepDeep"},
		{AvailableRequestedPowerStatesPowerCycleSoft, "PowerCycleSoft"},
		{AvailableRequestedPowerStatesOffHard, "OffHard"},
		{AvailableRequestedPowerStatesHibernate, "Hibernate"},
		{AvailableRequestedPowerStatesOffSoft, "OffSoft"},
		{AvailableRequestedPowerStatesPowerCycleHard, "PowerCycleHard"},
		{AvailableRequestedPowerStatesMasterBusReset, "MasterBusReset"},
		{AvailableRequestedPowerStatesDiagnosticInterrupt, "DiagnosticInterrupt"},
		{AvailableRequestedPowerStatesPowerOffSoftGraceful, "PowerOffSoftGraceful"},
		{AvailableRequestedPowerStatesPowerOffHardGraceful, "PowerOffHardGraceful"},
		{AvailableRequestedPowerStatesMasterBusResetGraceful, "MasterBusResetGraceful"},
		{AvailableRequestedPowerStatesPowerCycleSoftGraceful, "PowerCycleSoftGraceful"},
		{AvailableRequestedPowerStatesPowerCycleHardGraceful, "PowerCycleHardGraceful"},
		{AvailableRequestedPowerStates(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestPowerState_String(t *testing.T) {
	tests := []struct {
		state    PowerState
		expected string
	}{
		{PowerStateOther, "Other"},
		{PowerStateOn, "On"},
		{PowerStateSleepLight, "SleepLight"},
		{PowerStateSleepDeep, "SleepDeep"},
		{PowerStatePowerCycleSoft, "PowerCycleSoft"},
		{PowerStateOffHard, "OffHard"},
		{PowerStateHibernate, "Hibernate"},
		{PowerStateOffSoft, "OffSoft"},
		{PowerStatePowerCycleHard, "PowerCycleHard"},
		{PowerStateMasterBusReset, "MasterBusReset"},
		{PowerStateDiagnosticInterruptNMI, "DiagnosticInterruptNMI"},
		{PowerStatePowerOffSoftGraceful, "PowerOffSoftGraceful"},
		{PowerStatePowerOffHardGraceful, "PowerOffHardGraceful"},
		{PowerStateMasterBusResetGraceful, "MasterBusResetGraceful"},
		{PowerStatePowerCycleSoftGraceful, "PowerCycleSoftGraceful"},
		{PowerStatePowerCycleHardGraceful, "PowerCycleHardGraceful"},
		{PowerStateDiagnosticInterruptINIT, "DiagnosticInterruptINIT"},
		{PowerState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
