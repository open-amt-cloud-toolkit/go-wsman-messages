/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package service

const (
	CIMServiceAvailableToElement string = "CIM_ServiceAvailableToElement"
	ValueNotFound                string = "Value not found in map"
)

const (
	AvailableRequestedPowerStatesOther AvailableRequestedPowerStates = iota + 1
	AvailableRequestedPowerStatesOn
	AvailableRequestedPowerStatesSleepLight
	AvailableRequestedPowerStatesSleepDeep
	AvailableRequestedPowerStatesPowerCycleSoft
	AvailableRequestedPowerStatesOffHard
	AvailableRequestedPowerStatesHibernate
	AvailableRequestedPowerStatesOffSoft
	AvailableRequestedPowerStatesPowerCycleHard
	AvailableRequestedPowerStatesMasterBusReset
	AvailableRequestedPowerStatesDiagnosticInterrupt
	AvailableRequestedPowerStatesPowerOffSoftGraceful
	AvailableRequestedPowerStatesPowerOffHardGraceful
	AvailableRequestedPowerStatesMasterBusResetGraceful
	AvailableRequestedPowerStatesPowerCycleSoftGraceful
	AvailableRequestedPowerStatesPowerCycleHardGraceful
)

// availableRequestedPowerStatesMap is a map of the AvailableRequestedPowerStates enumeration.
var availableRequestedPowerStatesMap = map[AvailableRequestedPowerStates]string{
	AvailableRequestedPowerStatesOther:                  "Other",
	AvailableRequestedPowerStatesOn:                     "On",
	AvailableRequestedPowerStatesSleepLight:             "SleepLight",
	AvailableRequestedPowerStatesSleepDeep:              "SleepDeep",
	AvailableRequestedPowerStatesPowerCycleSoft:         "PowerCycleSoft",
	AvailableRequestedPowerStatesOffHard:                "OffHard",
	AvailableRequestedPowerStatesHibernate:              "Hibernate",
	AvailableRequestedPowerStatesOffSoft:                "OffSoft",
	AvailableRequestedPowerStatesPowerCycleHard:         "PowerCycleHard",
	AvailableRequestedPowerStatesMasterBusReset:         "MasterBusReset",
	AvailableRequestedPowerStatesDiagnosticInterrupt:    "DiagnosticInterrupt",
	AvailableRequestedPowerStatesPowerOffSoftGraceful:   "PowerOffSoftGraceful",
	AvailableRequestedPowerStatesPowerOffHardGraceful:   "PowerOffHardGraceful",
	AvailableRequestedPowerStatesMasterBusResetGraceful: "MasterBusResetGraceful",
	AvailableRequestedPowerStatesPowerCycleSoftGraceful: "PowerCycleSoftGraceful",
	AvailableRequestedPowerStatesPowerCycleHardGraceful: "PowerCycleHardGraceful",
}

// String returns a human-readable string representation of the AvailableRequestedPowerStates enumeration.
func (e AvailableRequestedPowerStates) String() string {
	if s, ok := availableRequestedPowerStatesMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	PowerStateOther PowerState = iota + 1
	PowerStateOn
	PowerStateSleepLight
	PowerStateSleepDeep
	PowerStatePowerCycleSoft
	PowerStateOffHard
	PowerStateHibernate
	PowerStateOffSoft
	PowerStatePowerCycleHard
	PowerStateMasterBusReset
	PowerStateDiagnosticInterruptNMI
	PowerStatePowerOffSoftGraceful
	PowerStatePowerOffHardGraceful
	PowerStateMasterBusResetGraceful
	PowerStatePowerCycleSoftGraceful
	PowerStatePowerCycleHardGraceful
	PowerStateDiagnosticInterruptINIT
)

// powerStateMap is a map of the PowerState enumeration.
var powerStateMap = map[PowerState]string{
	PowerStateOther:                   "Other",
	PowerStateOn:                      "On",
	PowerStateSleepLight:              "SleepLight",
	PowerStateSleepDeep:               "SleepDeep",
	PowerStatePowerCycleSoft:          "PowerCycleSoft",
	PowerStateOffHard:                 "OffHard",
	PowerStateHibernate:               "Hibernate",
	PowerStateOffSoft:                 "OffSoft",
	PowerStatePowerCycleHard:          "PowerCycleHard",
	PowerStateMasterBusReset:          "MasterBusReset",
	PowerStateDiagnosticInterruptNMI:  "DiagnosticInterruptNMI",
	PowerStatePowerOffSoftGraceful:    "PowerOffSoftGraceful",
	PowerStatePowerOffHardGraceful:    "PowerOffHardGraceful",
	PowerStateMasterBusResetGraceful:  "MasterBusResetGraceful",
	PowerStatePowerCycleSoftGraceful:  "PowerCycleSoftGraceful",
	PowerStatePowerCycleHardGraceful:  "PowerCycleHardGraceful",
	PowerStateDiagnosticInterruptINIT: "DiagnosticInterruptINIT",
}

// String returns a human-readable string representation of the PowerState enumeration.
func (e PowerState) String() string {
	if s, ok := powerStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}
