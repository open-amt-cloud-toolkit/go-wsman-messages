/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package chassis

const CIM_Chassis string = "CIM_Chassis"

// ChassisPackageType values
const (
	Unknown ChassisPackageType = iota
	Other
	SMBIOSReserved1
	Desktop
	LowProfileDesktop
	PizzaBox
	MiniTower
	Tower
	Portable
	LapTop
	Notebook
	HandHeld
	DockingStation
	AllInOne
	SubNotebook
	SpaceSaving
	LunchBox
	MainSystemChassis
	ExpansionChassis
	SubChassis
	BusExpansionChassis
	PeripheralChassis
	StorageChassis
	SMBIOSReserved2
	SealedCasePC
	SMBIOSReserved3
	CompactPCI
	AdvancedTCA
	BladeEnclosure
	SMBIOSReserved4
	Tablet
	Convertible
	Detachable
	IoTGateway
	EmbeddedPC
	MiniPC
	StickPC
)

// chassisPackageTypeToString is a map of ChassisPackageType values to their string representations.
var chassisPackageTypeToString = map[ChassisPackageType]string{
	Unknown:             "Unknown",
	Other:               "Other",
	SMBIOSReserved1:     "SMBIOSReserved1",
	Desktop:             "Desktop",
	LowProfileDesktop:   "LowProfileDesktop",
	PizzaBox:            "PizzaBox",
	MiniTower:           "MiniTower",
	Tower:               "Tower",
	Portable:            "Portable",
	LapTop:              "Laptop",
	Notebook:            "Notebook",
	HandHeld:            "Handheld",
	DockingStation:      "DockingStation",
	AllInOne:            "AllInOne",
	SubNotebook:         "SubNotebook",
	SpaceSaving:         "SpaceSaving",
	LunchBox:            "LunchBox",
	MainSystemChassis:   "MainSystemChassis",
	ExpansionChassis:    "ExpansionChassis",
	SubChassis:          "SubChassis",
	BusExpansionChassis: "BusExpansionChassis",
	PeripheralChassis:   "PeripheralChassis",
	StorageChassis:      "StorageChassis",
	SMBIOSReserved2:     "SMBIOSReserved2",
	SealedCasePC:        "SealedCasePC",
	SMBIOSReserved3:     "SMBIOSReserved3",
	CompactPCI:          "CompactPCI",
	AdvancedTCA:         "AdvancedTCA",
	BladeEnclosure:      "BladeEnclosure",
	SMBIOSReserved4:     "SMBIOSReserved4",
	Tablet:              "Tablet",
	Convertible:         "Convertible",
	Detachable:          "Detachable",
	IoTGateway:          "IoTGateway",
	EmbeddedPC:          "EmbeddedPC",
	MiniPC:              "MiniPC",
	StickPC:             "StickPC",
}

// String returns the string representation of a ChassisPackageType value.
func (c ChassisPackageType) String() string {
	if s, ok := chassisPackageTypeToString[c]; ok {
		return s
	}
	return "Value not found in map"
}

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

// operationalStatusMap is a map of the OperationalStatus enumeration
var operationalStatusMap = map[OperationalStatus]string{
	OperationalStatusUnknown:                 "Unknown",
	OperationalStatusOther:                   "Other",
	OperationalStatusOK:                      "OK",
	OperationalStatusDegraded:                "Degraded",
	OperationalStatusStressed:                "Stressed",
	OperationalStatusPredictiveFailure:       "Predictive Failure",
	OperationalStatusError:                   "Error",
	OperationalStatusNonRecoverableError:     "Non-Recoverable Error",
	OperationalStatusStarting:                "Starting",
	OperationalStatusStopping:                "Stopping",
	OperationalStatusStopped:                 "Stopped",
	OperationalStatusInService:               "In Service",
	OperationalStatusNoContact:               "No Contact",
	OperationalStatusLostCommunication:       "Lost Communication",
	OperationalStatusAborted:                 "Aborted",
	OperationalStatusDormant:                 "Dormant",
	OperationalStatusSupportingEntityInError: "Supporting Entity In Error",
	OperationalStatusCompleted:               "Completed",
	OperationalStatusPowerMode:               "Power Mode",
	OperationalStatusRelocating:              "Relocating",
}

// String returns a human-readable string representation of the OperationalStatus enumeration
func (e OperationalStatus) String() string {
	if s, ok := operationalStatusMap[e]; ok {
		return s
	}
	return "Value not found in map"
}

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

// packageTypeMap is a map of the PackageType enumeration
var packageTypeMap = map[PackageType]string{
	PackageTypeUnknown:               "Unknown",
	PackageTypeOther:                 "Other",
	PackageTypeRack:                  "Rack",
	PackageTypeChassisFrame:          "ChassisFrame",
	PackageTypeCrossConnectBackplane: "CrossConnectBackplane",
	PackageTypeContainerFrameSlot:    "ContainerFrameSlot",
	PackageTypePowerSupply:           "PowerSupply",
	PackageTypeFan:                   "Fan",
	PackageTypeSensor:                "Sensor",
	PackageTypeModuleCard:            "ModuleCard",
	PackageTypePortConnector:         "PortConnector",
	PackageTypeBattery:               "Battery",
	PackageTypeProcessor:             "Processor",
	PackageTypeMemory:                "Memory",
	PackageTypePowerSourceGenerator:  "PowerSourceGenerator",
	PackageTypeStorageMediaPackage:   "StorageMediaPackage",
	PackageTypeBlade:                 "Blade",
	PackageTypeBladeExpansion:        "BladeExpansion",
}

// String returns a human-readable string representation of the PackageType enumeration
func (e PackageType) String() string {
	if s, ok := packageTypeMap[e]; ok {
		return s
	}
	return "Value not found in map"
}
