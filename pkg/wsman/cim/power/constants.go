/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package power

const (
	CIM_PowerManagementService string = "CIM_PowerManagementService"
	RequestPowerStateChange    string = "RequestPowerStateChange"
)

const (
	// Power On
	PowerOn PowerState = 2 // Verified Hardware Power On

	// Sleep - Light
	SleepLight PowerState = 3 // ?

	// Sleep - Deep
	SleepDeep PowerState = 4 // ?

	// Power Cycle (Off Soft)
	PowerCycleOffSoft PowerState = 6 // ?

	// Power Off - Hard
	PowerOffHard PowerState = 8 // Verfied Hardware Power Off

	// Hibernate
	Hibernate PowerState = 7 // ?

	// Power Off - Soft
	PowerOffSoft PowerState = 9 // ?

	// Power Cycle (Off Hard)
	PowerCycleOffHard PowerState = 5 // Verified Hardware Power Cycle (off then on)

	// Master Bus Reset
	MasterBusReset PowerState = 10 // Verified Hardware Reboot

	// Diagnostic Interrupt (NMI)
	DiagnosticInterruptNMI PowerState = 11 // ?

	// Power Off - Soft Graceful
	PowerOffSoftGraceful PowerState = 12 // ?

	// Power Off - Hard Graceful
	PowerOffHardGraceful PowerState = 13 // ?

	// Master Bus Reset Graceful
	MasterBusResetGraceful PowerState = 14 // ?

	// Power Cycle (Off - Soft Graceful)
	PowerCycleOffSoftGraceful PowerState = 15 // ?

	// Power Cycle (Off - Hard Graceful)
	PowerCycleOffHardGraceful PowerState = 16 // ?
)
