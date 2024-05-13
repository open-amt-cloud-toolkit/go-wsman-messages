/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

const (
	CIMPhysicalMemory  string = "CIM_PhysicalMemory"
	CIMPhysicalPackage string = "CIM_PhysicalPackage"
	ValueNotFound      string = "Value not found in map"
)

const (
	MemoryTypeUnknown MemoryType = iota
	MemoryTypeOther
	MemoryTypeDRAM
	MemoryTypeSynchronousDRAM
	MemoryTypeCacheDRAM
	MemoryTypeEDO
	MemoryTypeEDRAM
	MemoryTypeVRAM
	MemoryTypeSRAM
	MemoryTypeRAM
	MemoryTypeROM
	MemoryTypeFlash
	MemoryTypeEEPROM
	MemoryTypeFEPROM
	MemoryTypeEPROM
	MemoryTypeCDRAM
	MemoryType3DRAM
	MemoryTypeSDRAM
	MemoryTypeSGRAM
	MemoryTypeRDRAM
	MemoryTypeDDR
	MemoryTypeDDR2
	MemoryTypeBRAM
	MemoryTypeFBDIMM
	MemoryTypeDDR3
	MemoryTypeFBD2
	MemoryTypeDDR4
	MemoryTypeLPDDR
	MemoryTypeLPDDR2
	MemoryTypeLPDDR3
	MemoryTypeLPDDR4
	MemoryTypeLogicalNonVolatileDevice
	MemoryTypeHBM
	MemoryTypeHBM2
	MemoryTypeDDR5
	MemoryTypeLPDDR5
	MemoryTypeHBM3
)

// memoryTypeMap is a map of the MemoryType enumeration.
var memoryTypeMap = map[MemoryType]string{
	MemoryTypeUnknown:                  "Unknown",
	MemoryTypeOther:                    "Other",
	MemoryTypeDRAM:                     "DRAM",
	MemoryTypeSynchronousDRAM:          "SynchronousDRAM",
	MemoryTypeCacheDRAM:                "CacheDRAM",
	MemoryTypeEDO:                      "EDO",
	MemoryTypeEDRAM:                    "EDRAM",
	MemoryTypeVRAM:                     "VRAM",
	MemoryTypeSRAM:                     "SRAM",
	MemoryTypeRAM:                      "RAM",
	MemoryTypeROM:                      "ROM",
	MemoryTypeFlash:                    "Flash",
	MemoryTypeEEPROM:                   "EEPROM",
	MemoryTypeFEPROM:                   "FEPROM",
	MemoryTypeEPROM:                    "EPROM",
	MemoryTypeCDRAM:                    "CDRAM",
	MemoryType3DRAM:                    "3DRAM",
	MemoryTypeSDRAM:                    "SDRAM",
	MemoryTypeSGRAM:                    "SGRAM",
	MemoryTypeRDRAM:                    "RDRAM",
	MemoryTypeDDR:                      "DDR",
	MemoryTypeDDR2:                     "DDR2",
	MemoryTypeBRAM:                     "BRAM",
	MemoryTypeFBDIMM:                   "FBDIMM",
	MemoryTypeDDR3:                     "DDR3",
	MemoryTypeFBD2:                     "FBD2",
	MemoryTypeDDR4:                     "DDR4",
	MemoryTypeLPDDR:                    "LPDDR",
	MemoryTypeLPDDR2:                   "LPDDR2",
	MemoryTypeLPDDR3:                   "LPDDR3",
	MemoryTypeLPDDR4:                   "LPDDR4",
	MemoryTypeLogicalNonVolatileDevice: "LogicalNon-VolatileDevice",
	MemoryTypeHBM:                      "HBM",
	MemoryTypeHBM2:                     "HBM2",
	MemoryTypeDDR5:                     "DDR5",
	MemoryTypeLPDDR5:                   "LPDDR5",
	MemoryTypeHBM3:                     "HBM3",
}

// String returns a human-readable string representation of the MemoryType enumeration.
func (e MemoryType) String() string {
	if s, ok := memoryTypeMap[e]; ok {
		return s
	}

	return ValueNotFound
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

// operationalStatusMap is a map of the OperationalStatus enumeration.
var operationalStatusMap = map[OperationalStatus]string{
	OperationalStatusUnknown:                 "Unknown",
	OperationalStatusOther:                   "Other",
	OperationalStatusOK:                      "OK",
	OperationalStatusDegraded:                "Degraded",
	OperationalStatusStressed:                "Stressed",
	OperationalStatusPredictiveFailure:       "PredictiveFailure",
	OperationalStatusError:                   "Error",
	OperationalStatusNonRecoverableError:     "NonRecoverableError",
	OperationalStatusStarting:                "Starting",
	OperationalStatusStopping:                "Stopping",
	OperationalStatusStopped:                 "Stopped",
	OperationalStatusInService:               "InService",
	OperationalStatusNoContact:               "NoContact",
	OperationalStatusLostCommunication:       "LostCommunication",
	OperationalStatusAborted:                 "Aborted",
	OperationalStatusDormant:                 "Dormant",
	OperationalStatusSupportingEntityInError: "SupportingEntityInError",
	OperationalStatusCompleted:               "Completed",
	OperationalStatusPowerMode:               "PowerMode",
	OperationalStatusRelocating:              "Relocating",
}

// String returns a human-readable string representation of the OperationalStatus enumeration.
func (e OperationalStatus) String() string {
	if s, ok := operationalStatusMap[e]; ok {
		return s
	}

	return ValueNotFound
}
