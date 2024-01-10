package physical

const (
	CIM_PhysicalMemory  string = "CIM_PhysicalMemory"
	CIM_PhysicalPackage string = "CIM_PhysicalPackage"
)

const (
	MemoryTypeUnknown MemoryType = iota
	MemoryTypeOther
	MemoryTypeDRAM
	MemoryTypeSynchronous_DRAM
	MemoryTypeCache_DRAM
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
	MemoryTypeLogicalnonvolatiledevice
	MemoryTypeHBM
	MemoryTypeHBM2
	MemoryTypeDDR5
	MemoryTypeLPDDR5
	MemoryTypeHBM3
)
