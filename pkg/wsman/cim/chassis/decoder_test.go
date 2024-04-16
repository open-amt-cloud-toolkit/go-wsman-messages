/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package chassis

import "testing"

func TestChassisPackageType_String(t *testing.T) {
	tests := []struct {
		state    ChassisPackageType
		expected string
	}{
		{Unknown, "Unknown"},
		{Other, "Other"},
		{SMBIOSReserved1, "SMBIOSReserved1"},
		{Desktop, "Desktop"},
		{LowProfileDesktop, "LowProfileDesktop"},
		{PizzaBox, "PizzaBox"},
		{MiniTower, "MiniTower"},
		{Tower, "Tower"},
		{Portable, "Portable"},
		{LapTop, "Laptop"},
		{Notebook, "Notebook"},
		{HandHeld, "Handheld"},
		{DockingStation, "DockingStation"},
		{AllInOne, "AllInOne"},
		{SubNotebook, "SubNotebook"},
		{SpaceSaving, "SpaceSaving"},
		{LunchBox, "LunchBox"},
		{MainSystemChassis, "MainSystemChassis"},
		{ExpansionChassis, "ExpansionChassis"},
		{SubChassis, "SubChassis"},
		{BusExpansionChassis, "BusExpansionChassis"},
		{PeripheralChassis, "PeripheralChassis"},
		{StorageChassis, "StorageChassis"},
		{SMBIOSReserved2, "SMBIOSReserved2"},
		{SealedCasePC, "SealedCasePC"},
		{SMBIOSReserved3, "SMBIOSReserved3"},
		{CompactPCI, "CompactPCI"},
		{AdvancedTCA, "AdvancedTCA"},
		{BladeEnclosure, "BladeEnclosure"},
		{SMBIOSReserved4, "SMBIOSReserved4"},
		{Tablet, "Tablet"},
		{Convertible, "Convertible"},
		{Detachable, "Detachable"},
		{IoTGateway, "IoTGateway"},
		{EmbeddedPC, "EmbeddedPC"},
		{MiniPC, "MiniPC"},
		{StickPC, "StickPC"},
		{ChassisPackageType(999), "Value not found in map"},
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
		{OperationalStatusPredictiveFailure, "Predictive Failure"},
		{OperationalStatusError, "Error"},
		{OperationalStatusNonRecoverableError, "Non-Recoverable Error"},
		{OperationalStatusStarting, "Starting"},
		{OperationalStatusStopping, "Stopping"},
		{OperationalStatusStopped, "Stopped"},
		{OperationalStatusInService, "In Service"},
		{OperationalStatusNoContact, "No Contact"},
		{OperationalStatusLostCommunication, "Lost Communication"},
		{OperationalStatusAborted, "Aborted"},
		{OperationalStatusDormant, "Dormant"},
		{OperationalStatusSupportingEntityInError, "Supporting Entity In Error"},
		{OperationalStatusCompleted, "Completed"},
		{OperationalStatusPowerMode, "Power Mode"},
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

func TestPackageType_String(t *testing.T) {
	tests := []struct {
		state    PackageType
		expected string
	}{
		{PackageTypeUnknown, "Unknown"},
		{PackageTypeOther, "Other"},
		{PackageTypeRack, "Rack"},
		{PackageTypeChassisFrame, "ChassisFrame"},
		{PackageTypeCrossConnectBackplane, "CrossConnectBackplane"},
		{PackageTypeContainerFrameSlot, "ContainerFrameSlot"},
		{PackageTypePowerSupply, "PowerSupply"},
		{PackageTypeFan, "Fan"},
		{PackageTypeSensor, "Sensor"},
		{PackageTypeModuleCard, "ModuleCard"},
		{PackageTypePortConnector, "PortConnector"},
		{PackageTypeBattery, "Battery"},
		{PackageTypeProcessor, "Processor"},
		{PackageTypeMemory, "Memory"},
		{PackageTypePowerSourceGenerator, "PowerSourceGenerator"},
		{PackageTypeStorageMediaPackage, "StorageMediaPackage"},
		{PackageTypeBlade, "Blade"},
		{PackageTypeBladeExpansion, "BladeExpansion"},
		{PackageType(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
