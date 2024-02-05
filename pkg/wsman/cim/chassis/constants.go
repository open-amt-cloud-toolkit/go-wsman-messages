/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package chassis

const CIM_Chassis string = "CIM_Chassis"

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

const (
	ChassisPackageTypeUnknown ChassisPackageType = iota
	ChassisPackageTypeOther
	ChassisPackageTypeSMBIOSReserved
	ChassisPackageTypeDesktop
	ChassisPackageTypeLowProfileDesktop
	ChassisPackageTypePizzaBox
	ChassisPackageTypeMiniTower
	ChassisPackageTypeTower
	ChassisPackageTypePortable
	ChassisPackageTypeLapTop
	ChassisPackageTypeNotebook
	ChassisPackageTypeHandHeld
	ChassisPackageTypeDockingStation
	ChassisPackageTypeAllinOne
	ChassisPackageTypeSubNotebook
	ChassisPackageTypeSpaceSaving
	ChassisPackageTypeLunchBox
	ChassisPackageTypeMainSystemChassis
	ChassisPackageTypeExpansionChassis
	ChassisPackageTypeSubChassis
	ChassisPackageTypeBusExpansionChassis
	ChassisPackageTypePeripheralChassis
	ChassisPackageTypeStorageChassis
	ChassisPackageTypeSealedCasePC
	ChassisPackageTypeCompactPCI
	ChassisPackageTypeAdvancedTCA
	ChassisPackageTypeBladeEnclosure
	ChassisPackageTypeTablet
	ChassisPackageTypeConvertible
	ChassisPackageTypeDetachable
	ChassisPackageTypeIoTGateway
	ChassisPackageTypeEmbeddedPC
	ChassisPackageTypeMiniPC
	ChassisPackageTypeStickPC
	ChassisPackageTypeDMTFReserved
	ChassisPackageTypeVendorReserved
)
