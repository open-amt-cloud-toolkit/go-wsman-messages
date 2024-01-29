/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package card

const CIM_Card string = "CIM_Card"

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
	OperationalStatusSupportingEntityinError
	OperationalStatusCompleted
	OperationalStatusPowerMode
	OperationalStatusRelocating
)

const (
	PackageTypeUnknown PackageType = iota
	PackageTypeOther
	PackageTypeRack
	PackageTypeChassisFrame
	PackageTypeCrossConnectBackplane
	PackageTypeContainerFrameSlot
	PackageTypePowerSupply
	PackageTypeFan
	PackageTypeSensor
	PackageTypeModuleCard
	PackageTypePortConnector
	PackageTypeBattery
	PackageTypeProcessor
	PackageTypeMemory
	PackageTypePowerSourceGenerator
	PackageTypeStorageMediaPackage
	PackageTypeBlade
	PackageTypeBladeExpansion
)
