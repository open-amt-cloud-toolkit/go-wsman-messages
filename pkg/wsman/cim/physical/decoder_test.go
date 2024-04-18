/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import "testing"

func TestMemoryType_String(t *testing.T) {
	tests := []struct {
		state    MemoryType
		expected string
	}{
		{MemoryTypeUnknown, "Unknown"},
		{MemoryTypeOther, "Other"},
		{MemoryTypeDRAM, "DRAM"},
		{MemoryTypeSynchronous_DRAM, "SynchronousDRAM"},
		{MemoryTypeCache_DRAM, "CacheDRAM"},
		{MemoryTypeEDO, "EDO"},
		{MemoryTypeEDRAM, "EDRAM"},
		{MemoryTypeVRAM, "VRAM"},
		{MemoryTypeSRAM, "SRAM"},
		{MemoryTypeRAM, "RAM"},
		{MemoryTypeROM, "ROM"},
		{MemoryTypeFlash, "Flash"},
		{MemoryTypeEEPROM, "EEPROM"},
		{MemoryTypeFEPROM, "FEPROM"},
		{MemoryTypeEPROM, "EPROM"},
		{MemoryTypeCDRAM, "CDRAM"},
		{MemoryType3DRAM, "3DRAM"},
		{MemoryTypeSDRAM, "SDRAM"},
		{MemoryTypeSGRAM, "SGRAM"},
		{MemoryTypeRDRAM, "RDRAM"},
		{MemoryTypeDDR, "DDR"},
		{MemoryTypeDDR2, "DDR2"},
		{MemoryTypeBRAM, "BRAM"},
		{MemoryTypeFBDIMM, "FBDIMM"},
		{MemoryTypeDDR3, "DDR3"},
		{MemoryTypeFBD2, "FBD2"},
		{MemoryTypeDDR4, "DDR4"},
		{MemoryTypeLPDDR, "LPDDR"},
		{MemoryTypeLPDDR2, "LPDDR2"},
		{MemoryTypeLPDDR3, "LPDDR3"},
		{MemoryTypeLPDDR4, "LPDDR4"},
		{MemoryTypeLogicalNonVolatileDevice, "LogicalNon-VolatileDevice"},
		{MemoryTypeHBM, "HBM"},
		{MemoryTypeHBM2, "HBM2"},
		{MemoryTypeDDR5, "DDR5"},
		{MemoryTypeLPDDR5, "LPDDR5"},
		{MemoryTypeHBM3, "HBM3"},
		{MemoryType(999), "Value not found in map"},
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
