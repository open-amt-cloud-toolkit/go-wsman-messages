/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package processor

const CIM_Processor string = "CIM_Processor"

const (
	OperationalStatusUnknown OperationalStatus = iota
	OperationalStatusOther
	OperationalStatusOK
	OperationalStatusDegraded
	OperationalStatusStressed
	OperationalStatusPredictiveFailure
	OperationalStatusError
	OperationalStatusNonRecoverableError
	OperationalStatusStarting
	OperationalStatusStopping
	OperationalStatusStopped
	OperationalStatusInService
	OperationalStatusNoContact
	OperationalStatusLostCommunication
	OperationalStatusAborted
	OperationalStatusDormant
	OperationalStatusSupportingEntityInError
	OperationalStatusCompleted
	OperationalStatusPowerMode
	OperationalStatusRelocating
)

const (
	HealthStateUnknown HealthState = iota
	HealthStateOK
	HealthStateDegraded
	HealthStateWarning
	HealthStateMinorFailure
	HealthStateMajorFailure
	HealthStateCriticalFailure
	HealthStateNonRecoverableError
	HealthStateDMTFReserved
	HealthStateVendorSpecific
)

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTestDeferred
	EnabledStateQuiesce
	EnabledStateStarting
	EnabledStateDMTFReserved
	EnabledStateVendorReserved
)

const (
	RequestedStateUnknown RequestedState = iota
	RequestedStateEnabled
	RequestedStateDisabled
	RequestedStateShutDown
	RequestedStateNoChange
	RequestedStateOffline
	RequestedStateTest
	RequestedStateDeferred
	RequestedStateQuiesce
	RequestedStateReboot
	RequestedStateReset
	RequestedStateNotApplicable
	RequestedStateDMTFReserved
	RequestedStateVendorReserved
)

const (
	UpgradeMethodUnknown UpgradeMethod = iota
	UpgradeMethodOther
	UpgradeMethodDaughterBoard
	UpgradeMethodZIFSocket
	UpgradeMethodReplacement
	UpgradeMethodPiggyBack
	UpgradeMethodNone
	UpgradeMethodLIFSocket
	UpgradeMethodSlot1
	UpgradeMethodSlot2
	UpgradeMethod370PinSocket
	UpgradeMethodSlotA
	UpgradeMethodSlotM
	UpgradeMethodSocket423
	UpgradeMethodSocketA
	UpgradeMethodSocket478
	UpgradeMethodSocket754
	UpgradeMethodSocket940
	UpgradeMethodSocket939
	UpgradeMethodSocketPGA604
	UpgradeMethodSocketLGA771
	UpgradeMethodSocketLGA775
	UpgradeMethodSocketS1
	UpgradeMethodSocketAM2
	UpgradeMethodSocketF
	UpgradeMethodSocketLGA1366
	UpgradeMethodSocketG34
	UpgradeMethodSocketAM3
	UpgradeMethodSocketC32
	UpgradeMethodSocketLGA1156
	UpgradeMethodSocketLGA1567
	UpgradeMethodSocketPGA988A
	UpgradeMethodSocketBGA1288
	UpgradeMethodPGA988B
	UpgradeMethodBGA1023
	UpgradeMethodBGA1224
	UpgradeMethodLGA1155
	UpgradeMethodLGA1356
	UpgradeMethodLGA2011
	UpgradeMethodSocketFS1
	UpgradeMethodSocketFS2
	UpgradeMethodSocketFM1
	UpgradeMethodSocketFM2
	UpgradeMethodSocketLGA2011
	UpgradeMethodSocketLGA1356
	UpgradeMethodSocketLGA1150
	UpgradeMethodSocketBGA1168
	UpgradeMethodSocketBGA1234
	UpgradeMethodSocketBGA1364
	UpgradeMethodSocketAM4
	UpgradeMethodSocketLGA1151
	UpgradeMethodSocketBGA1356
	UpgradeMethodSocketBGA1440
	UpgradeMethodSocketBGA1515
	UpgradeMethodSocketLGA3647
	UpgradeMethodSocketSP3
	UpgradeMethodSocketSP3r2
	UpgradeMethodSocketLGA2066
	UpgradeMethodSocketBGA1392
	UpgradeMethodSocketBGA1510
	UpgradeMethodSocketBGA1528
	UpgradeMethodSocketLGA4189
	UpgradeMethodSocketLGA1200
	UpgradeMethodSocketLGA4677
	UpgradeMethodSocketLGA1700
	UpgradeMethodSocketBGA1744
	UpgradeMethodSocketBGA1781
	UpgradeMethodSocketBGA1211
	UpgradeMethodSocketBGA2422
	UpgradeMethodSocketLGA5773
	UpgradeMethodSocketBGA5773
	UpgradeMethodSocketAM5
	UpgradeMethodSocketSP5
	UpgradeMethodSocketSP6
	UpgradeMethodSocketBGA883
	UpgradeMethodSocketBGA1190
	UpgradeMethodSocketBGA4129
	UpgradeMethodSocketLGA4710
	UpgradeMethodSocketLGA7529
)

const (
	CPUStatusUnknown CPUStatus = iota
	CPUStatusCPUEnabled
	CPUStatusCPUDisabledByUser
	CPUStatusCPUDisabledByBIOS
	CPUStatusCPUIsIdle
	CPUStatusOther
)
