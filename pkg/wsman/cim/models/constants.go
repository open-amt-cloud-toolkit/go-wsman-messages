/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package models

// RoleCharacteristicsToString is a map of RoleCharacteristics value to string
var RoleCharacteristicsToString = map[int]string{
	2: "Static",
	3: "Opaque",
}

// ConvertRoleCharacteristicsToString returns the string representation of RoleCharacteristics
func ConvertRoleCharacteristicsToString(value int) string {
	if value, exist := RoleCharacteristicsToString[value]; exist {
		return value
	}
	return "Value not found in map"
}

// UsageRestrictionToString is a map of UsageRestriction value to string
var UsageRestrictionToString = map[int]string{
	0: "Unknown",
	1: "FrontEndOnly",
	2: "BackEndOnly",
	3: "NotRestricted",
}

// ConvertUsageRestrictionToString returns the string representation of UsageRestriction
func ConvertUsageRestrictionToString(value int) string {
	if value, exist := UsageRestrictionToString[value]; exist {
		return value
	}
	return "Value not found in map"
}

// PortTypeToString is a map of PortType value to string
var PortTypeToString = map[int]string{
	0: "Unknown",
	1: "Other",
	2: "Not Applicable",
}

// ConvertPortTypeToString returns the string representation of PortType
func ConvertPortTypeToString(value int) string {
	if value, exist := PortTypeToString[value]; exist {
		return value
	}
	return "Value not found in map"
}

// LinkTechnologyToString is a map of LinkTechnology value to string
var LinkTechnologyToString = map[int]string{
	0:  "Unknown",
	1:  "Other",
	2:  "Ethernet",
	3:  "IB",
	4:  "FC",
	5:  "FDDI",
	6:  "ATM",
	7:  "Token Ring",
	8:  "Frame Relay",
	9:  "Infrared",
	10: "Bluetooth",
	11: "Wireless LAN",
}

// ConvertLinkTechnologyToString returns the string representation of LinkTechnology
func ConvertLinkTechnologyToString(value int) string {
	if value, exist := LinkTechnologyToString[value]; exist {
		return value
	}
	return "Value not found in map"
}

// RemovalConditionsToString is a map of RemovalConditions value to string
var RemovalConditionsToString = map[int]string{
	0: "Unknown",
	2: "Not Applicable",
	3: "Removable When Off",
	4: "Removable When On or Off",
}

// ConvertRemovalConditionsToString returns the string representation of RemovalConditions
func ConvertRemovalConditionsToString(value int) string {
	if value, exist := RemovalConditionsToString[value]; exist {
		return value
	}
	return "Value not found in map"
}

// ServicePhilosophyToString is a map of ServicePhilosophy value to string
var ServicePhilosophyToString = map[int]string{
	0: "Unknown",
	1: "Other",
	2: "Service From Top",
	3: "Service From Front",
	4: "Service From Back",
	5: "Service From Side",
	6: "Sliding Trays",
	7: "Removable Sides",
	8: "Moveable",
}

// ConvertServicePhilosophyToString returns the string representation of ServicePhilosophy
func ConvertServicePhilosophyToString(value int) string {
	if value, exist := ServicePhilosophyToString[value]; exist {
		return value
	}
	return "Value not found in map"
}

// SecurityBreachToString is a map of SecurityBreach value to string
var SecurityBreachToString = map[int]string{
	1: "Other",
	2: "Unknown",
	3: "No Breach",
	4: "Breach Attempted",
	5: "Breach Successful",
}

// ConvertSecurityBreachToString returns the string representation of SecurityBreach
func ConvertSecurityBreachToString(value int) string {
	if value, exist := SecurityBreachToString[value]; exist {
		return value
	}
	return "Value not found in map"
}

// ChassisPackageTypeToString is a map of ChassisPackageType value to string
var ChassisPackageTypeToString = map[int]string{
	0:  "Unknown",
	1:  "Other",
	2:  "SMBIOS Reserved",
	3:  "Desktop",
	4:  "Low Profile Desktop",
	5:  "Pizza Box",
	6:  "Mini Tower",
	7:  "Tower",
	8:  "Portable",
	9:  "LapTop",
	10: "Notebook",
	11: "Hand Held",
	12: "Docking Station",
	13: "All in One",
	14: "Sub Notebook",
	15: "Space-Saving",
	16: "Lunch Box",
	17: "Main System Chassis",
	18: "Expansion Chassis",
	19: "SubChassis",
	20: "Bus Expansion Chassis",
	21: "Peripheral Chassis",
	22: "Storage Chassis",
	23: "SMBIOS Reserved",
	24: "Sealed-Case PC",
	25: "SMBIOS Reserved",
	26: "CompactPCI",
	27: "AdvancedTCA",
	28: "Blade Enclosure",
	29: "SMBIOS Reserved",
	30: "Tablet",
	31: "Convertible",
	32: "Detachable",
	33: "IoT Gateway",
	34: "Embedded PC",
	35: "Mini PC",
	36: "Stick PC",
}

// ConvertChassisPackageTypeToString returns the string representation of ChassisPackageType
func ConvertChassisPackageTypeToString(value int) string {
	if value, exist := ChassisPackageTypeToString[value]; exist {
		return value
	}
	return "Value not found in map"
}

// SoftwareElementStateToString is a map of SoftwareElementState value to string
var SoftwareElementStateToString = map[int]string{
	0: "Deployable",
	1: "Installable",
	2: "Executable",
	3: "Running",
}

// ConvertSoftwareElementStateToString returns the string representation of SoftwareElementState
func ConvertSoftwareElementStateToString(value int) string {
	if value, exist := SoftwareElementStateToString[value]; exist {
		return value
	}
	return "Value not found in map"
}

const (
	Enabled_Enabled Enabled = 1 + iota
	Enabled_Disabled
	Enabled_EnabledForDebug
)

const (
	AuthenticationProtocol_EAPTLS AuthenticationProtocol = iota
	AuthenticationProtocol_EAPTTLSMSCHAPv2
	AuthenticationProtocol_PEAPv0EAPMSCHAPv2
	AuthenticationProtocol_PEAPv1EAPGTC
	AuthenticationProtocol_EAPFASTMSCHAPv2
	AuthenticationProtocol_EAPFASTGTC
	AuthenticationProtocol_EAPMD5
	AuthenticationProtocol_EAPPSK
	AuthenticationProtocol_EAPSIM
	AuthenticationProtocol_EAPAKA
	AuthenticationProtocol_EAPFASTTLS
)

const (
	ConcreteJobState_New ConcreteJobState = 2 + iota
	ConcreteJobState_Starting
	ConcreteJobState_Running
	ConcreteJobState_Suspended
	ConcreteJobState_ShuttingDown
	ConcreteJobState_Completed
	ConcreteJobState_Terminated
	ConcreteJobState_Killed
	ConcreteJobState_Exception
	ConcreteJobState_Service
	ConcreteJobState_QueryPending
)
const (
	CommunicationStatus_UnknownCS CommunicationStatus = iota
	CommunicationStatus_NotAvailableCS
	CommunicationStatus_CommunicationOK
	CommunicationStatus_LostCommunication
	CommunicationStatus_NoContact
)
const (
	DetailedStatus_NotAvailableDS DetailedStatus = iota
	DetailedStatus_NoAdditionalInformation
	DetailedStatus_Stressed
	DetailedStatus_PredictiveFailure
	DetailedStatus_NonRecoverableError
	DetailedStatus_SupportingEntityInError
)
const (
	OperatingStatus_UnknownOS OperatingStatus = iota
	OperatingStatus_NotAvailableOS
	OperatingStatus_Servicing
	OperatingStatus_Starting
	OperatingStatus_Stopping
	OperatingStatus_Stopped
	OperatingStatus_Aborted
	OperatingStatus_Dormant
	OperatingStatus_Completed
	OperatingStatus_Migrating
	OperatingStatus_Emigrating
	OperatingStatus_Immigrating
	OperatingStatus_SnapShotting
	OperatingStatus_ShuttingDown
	OperatingStatus_InTest
	OperatingStatus_Transitioning
	OperatingStatus_InService
)
const (
	PrimaryStatus_UnknownPS PrimaryStatus = iota
	PrimaryStatus_OK
	PrimaryStatus_Degraded
	PrimaryStatus_ErrorPS
)
const (
	January RunMonth = iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
const (
	SaturdayNegative RunDayOfWeek = iota - 7
	FridayNegative
	ThursdayNegative
	WednesdayNegative
	TuesdayNegative
	MondayNegative
	SundayNegative
	ExactDayOfMonth
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
const (
	RecoveryAction_UnknownRA RecoveryAction = iota
	RecoveryAction_Other
	RecoveryAction_DoNotContinue
	RecoveryAction_ContinueWithNextJob
	RecoveryAction_RerunJob
	RecoveryAction_RunRecoveryJob
)
const (
	LocalTime LocalOrUtcTime = iota + 1
	UTCTime
)
const (
	PowerManagementCapabilities_Unknown PowerManagementCapabilitiesValues = iota
	PowerManagementCapabilities_NotSupported
	PowerManagementCapabilities_Disabled
	PowerManagementCapabilities_Enabled
	PowerManagementCapabilities_PowerSavingModesEnteredAutomatically
	PowerManagementCapabilities_PowerStateSettable
	PowerManagementCapabilities_PowerCyclingSupported
	PowerManagementCapabilities_TimedPowerOnSupported
)
const (
	Availability_Other AvailabilityValues = iota + 1
	Availability_Unknown
	Availability_RunningFullPower
	Availability_Warning
	Availability_InTest
	Availability_NotApplicable
	Availability_PowerOff
	Availability_OffLine
	Availability_OffDuty
	Availability_Degraded
	Availability_NotInstalled
	Availability_InstallError
	Availability_PowerSaveUnknown
	Availability_PowerSaveLowPowerMode
	Availability_PowerSaveStandby
	Availability_PowerCycle
	Availability_PowerSaveWarning
	Availability_Paused
	Availability_NotReady
	Availability_NotConfigured
	Availability_Quiesced
)
const (
	StatusInfo_Other StatusInfoValues = iota + 1
	StatusInfo_Unknown
	StatusInfo_Enabled
	StatusInfo_Disabled
	StatusInfo_NotApplicable
)
const (
	AdditionalAvailability_Other AdditionalAvailabilityValues = iota + 1
	AdditionalAvailability_Unknown
	AdditionalAvailability_RunningFullPower
	AdditionalAvailability_Warning
	AdditionalAvailability_InTest
	AdditionalAvailability_NotApplicable
	AdditionalAvailability_PowerOff
	AdditionalAvailability_OffLine
	AdditionalAvailability_OffDuty
	AdditionalAvailability_Degraded
	AdditionalAvailability_NotInstalled
	AdditionalAvailability_InstallError
	AdditionalAvailability_PowerSaveUnknown
	AdditionalAvailability_PowerSaveLowPowerMode
	AdditionalAvailability_PowerSaveStandby
	AdditionalAvailability_PowerCycle
	AdditionalAvailability_PowerSaveWarning
	AdditionalAvailability_Paused
	AdditionalAvailability_NotReady
	AdditionalAvailability_NotConfigured
	AdditionalAvailability_Quiesced
)

const (
	ForceHardDriveBoot       BootConfigSettingInstanceID = "Intel(r) AMT: Force Hard-drive Boot"
	ForceCD_DVDBoot          BootConfigSettingInstanceID = "Intel(r) AMT: Force CD/DVD Boot"
	ForcePXEBoot             BootConfigSettingInstanceID = "Intel(r) AMT: Force PXE Boot"
	ForceOCRUEFIHTTPSBoot    BootConfigSettingInstanceID = "Intel(r) AMT: Force OCR UEFI HTTPS Boot"
	ForceOCRUEFIBootOption1  BootConfigSettingInstanceID = "Intel(r) AMT: Force OCR UEFI Boot Option 1"
	ForceOCRUEFIBootOption2  BootConfigSettingInstanceID = "Intel(r) AMT: Force OCR UEFI Boot Option 2"
	ForceOCRUEFIBootOption3  BootConfigSettingInstanceID = "Intel(r) AMT: Force OCR UEFI Boot Option 3"
	ForceOCRUEFIBootOption4  BootConfigSettingInstanceID = "Intel(r) AMT: Force OCR UEFI Boot Option 4"
	ForceOCRUEFIBootOption5  BootConfigSettingInstanceID = "Intel(r) AMT: Force OCR UEFI Boot Option 5"
	ForceOCRUEFIBootOption6  BootConfigSettingInstanceID = "Intel(r) AMT: Force OCR UEFI Boot Option 6"
	ForceOCRUEFIBootOption7  BootConfigSettingInstanceID = "Intel(r) AMT: Force OCR UEFI Boot Option 7"
	ForceOCRUEFIBootOption8  BootConfigSettingInstanceID = "Intel(r) AMT: Force OCR UEFI Boot Option 8"
	ForceOCRUEFIBootOption9  BootConfigSettingInstanceID = "Intel(r) AMT: Force OCR UEFI Boot Option 9"
	ForceOCRUEFIBootOption10 BootConfigSettingInstanceID = "Intel(r) AMT: Force OCR UEFI Boot Option 10"
)
const (
	RemovalConditions_Unknown              RemovalConditions = 0
	RemovalConditions_NotApplicable        RemovalConditions = 2
	RemovalConditions_RemovableWhenOff     RemovalConditions = 3
	RemovalConditions_RemovableWhenOnOrOff RemovalConditions = 4
)

const (
	FailThroughSupportedUnknown FailThroughSupported = iota
	IsSupported
	NotSupported
)

const (
	OverwritePolicy_Unknown         OverwritePolicy = 0
	OverwritePolicy_WrapsWhenFull   OverwritePolicy = 2
	OverwritePolicy_NeverOverwrites OverwritePolicy = 7
)
const (
	LogState_Unknown       LogState = 0
	LogState_Normal        LogState = 2
	LogState_Erasing       LogState = 3
	LogState_NotApplicable LogState = 4
)
const (
	Capabilities_Unknown CapabilitiesValues = iota
	Capabilities_Other
	Capabilities_WriteRecordSupported
	Capabilities_DeleteRecordSupported
	Capabilities_CanMoveBackwardInLog
	Capabilities_FreezeLogSupported
	Capabilities_ClearLogSupported
	Capabilities_SupportsAddressingByOrdinalRecordNumber
	Capabilities_VariableLengthRecordsSupported
	Capabilities_VariableFormatsForRecords
	Capabilities_CanFlagRecordsForOverwrite
)
const (
	LastChange_Unknown LastChange = iota
	LastChange_Add
	LastChange_Delete
	LastChange_Modify
	LastChange_LogCleared
)
const (
	CharacterSet_Unknown CharacterSet = iota
	CharacterSet_Other
	CharacterSet_ASCII
	CharacterSet_Unicode
	CharacterSet_ISO2022
	CharacterSet_ISO8859
	CharacterSet_ExtendedUNIXCode
	CharacterSet_UTF8
	CharacterSet_UCS2
	CharacterSet_BitMappedData
	CharacterSet_OctetString
	CharacterSet_DefinedByIndividualRecords
)

const (
	UpgradeMethod_Other UpgradeMethod = iota + 1
	UpgradeMethod_Unknown
	UpgradeMethod_DaughterBoard
	UpgradeMethod_ZIFSocket
	UpgradeMethod_ReplacementPiggyBack
	UpgradeMethod_None
	UpgradeMethod_LIFSocket
	UpgradeMethod_Slot1
	UpgradeMethod_Slot2
	UpgradeMethod_Socket370Pin
	UpgradeMethod_SlotA
	UpgradeMethod_SlotM
	UpgradeMethod_Socket423
	UpgradeMethod_SocketASocket462
	UpgradeMethod_Socket478
	UpgradeMethod_Socket754
	UpgradeMethod_Socket940
	UpgradeMethod_Socket939
	UpgradeMethod_SocketmPGA604
	UpgradeMethod_SocketLGA771
	UpgradeMethod_SocketLGA775
	UpgradeMethod_SocketS1
	UpgradeMethod_SocketAM2
	UpgradeMethod_SocketF1207
	UpgradeMethod_SocketLGA1366
	UpgradeMethod_SocketG34
	UpgradeMethod_SocketAM3
	UpgradeMethod_SocketC32
	UpgradeMethod_SocketLGA1156
	UpgradeMethod_SocketLGA1567
	UpgradeMethod_SocketPGA988A
	UpgradeMethod_SocketBGA1288
	UpgradeMethod_rPGA988B
	UpgradeMethod_BGA1023
	UpgradeMethod_BGA1224
	UpgradeMethod_LGA1155
	UpgradeMethod_LGA1356
	UpgradeMethod_LGA2011
	UpgradeMethod_SocketFS1
	UpgradeMethod_SocketFS2
	UpgradeMethod_SocketFM1
	UpgradeMethod_SocketFM2
	UpgradeMethod_SocketLGA20113
	UpgradeMethod_SocketLGA13563
	UpgradeMethod_SocketLGA1150
	UpgradeMethod_SocketBGA1168
	UpgradeMethod_SocketBGA1234
	UpgradeMethod_SocketBGA1364
	UpgradeMethod_SocketAM4
	UpgradeMethod_SocketLGA1151
	UpgradeMethod_SocketBGA1356
	UpgradeMethod_SocketBGA1440
	UpgradeMethod_SocketBGA1515
	UpgradeMethod_SocketLGA36471
	UpgradeMethod_SocketSP3
	UpgradeMethod_SocketSP3r2
	UpgradeMethod_SocketLGA2066
	UpgradeMethod_SocketBGA1392
	UpgradeMethod_SocketBGA1510
	UpgradeMethod_SocketBGA1528
	UpgradeMethod_SocketLGA4189
	UpgradeMethod_SocketLGA1200
	UpgradeMethod_SocketLGA4677
	UpgradeMethod_SocketLGA1700
	UpgradeMethod_SocketBGA1744
	UpgradeMethod_SocketBGA1781
	UpgradeMethod_SocketBGA1211
	UpgradeMethod_SocketBGA2422
	UpgradeMethod_SocketLGA5773
	UpgradeMethod_SocketBGA5773
	UpgradeMethod_SocketAM5
	UpgradeMethod_SocketSP5
	UpgradeMethod_SocketSP6
	UpgradeMethod_SocketBGA883
	UpgradeMethod_SocketBGA1190
	UpgradeMethod_SocketBGA4129
	UpgradeMethod_SocketLGA4710
	UpgradeMethod_SocketLGA7529
	UpgradeMethod_SocketBGA1964
	UpgradeMethod_SocketBGA1792
	UpgradeMethod_SocketBGA2049
	UpgradeMethod_SocketBGA2551
	UpgradeMethod_SocketLGA1851
	UpgradeMethod_SocketBGA2114
	UpgradeMethod_SocketBGA2833
)

const (
	CPUStatus_Unknown CPUStatus = iota
	CPUStatus_Enabled
	CPUStatus_DisabledByUser
	CPUStatus_DisabledByBIOS
	CPUStatus_Idle
	CPUStatus_Other
)

const (
	OperationalStatus_Unknown OperationalStatus = iota
	OperationalStatus_Other
	OperationalStatus_OK
	OperationalStatus_Degraded
	OperationalStatus_Stressed
	OperationalStatus_PredictiveFailure
	OperationalStatus_Error
	OperationalStatus_NonRecoverableError
	OperationalStatus_Starting
	OperationalStatus_Stopping
	OperationalStatus_Stopped
	OperationalStatus_InService
	OperationalStatus_NoContact
	OperationalStatus_LostCommunication
	OperationalStatus_Aborted
	OperationalStatus_Dormant
	OperationalStatus_SupportingEntityinError
	OperationalStatus_Completed
	OperationalStatus_PowerMode
	OperationalStatus_Relocating
)

const (
	HealthState_Unknown             HealthState = 0
	HealthState_OK                  HealthState = 5
	HealthState_DegradedWarning     HealthState = 10
	HealthState_MinorFailure        HealthState = 15
	HealthState_MajorFailure        HealthState = 20
	HealthState_CriticalFailure     HealthState = 25
	HealthState_NonRecoverableError HealthState = 30
)

const (
	EnabledState_Unknown EnabledState = iota
	EnabledState_Other
	EnabledState_Enabled
	EnabledState_Disabled
	EnabledState_ShuttingDown
	EnabledState_NotApplicable
	EnabledState_EnabledbutOffline
	EnabledState_InTest
	EnabledState_Deferred
	EnabledState_Quiesce
	EnabledState_Starting
)

const (
	RequestedState_Enabled RequestedState = iota + 2
	RequestedState_Disabled
	RequestedState_ShutDown
	RequestedState_NoChange
	RequestedState_Offline
	RequestedState_Test
	RequestedState_Deferred
	RequestedState_Quiesce
	RequestedState_Reboot
	RequestedState_Reset
	RequestedState_NotApplicable
	RequestedState_Unknown RequestedState = 0
)

const (
	EnabledDefault_Enabled           EnabledDefault = 2
	EnabledDefault_Disabled          EnabledDefault = 3
	EnabledDefault_NotApplicable     EnabledDefault = 5
	EnabledDefault_EnabledbutOffline EnabledDefault = 6
	EnabledDefault_NoDefault         EnabledDefault = 7
	EnabledDefault_Quiesce           EnabledDefault = 9
)
