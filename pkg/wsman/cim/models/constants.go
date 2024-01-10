/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package models

const (
	RoleCharacteristicsStatic RoleCharacteristics = 2
	RoleCharacteristicsOpaque RoleCharacteristics = 3
)

const (
	UsageRestrictionUnknown int = iota
	UsageRestrictionFrontEndOnly
	UsageRestrictionBackEndOnly
	UsageRestrictionNotRestricted
)

const (
	PortTypeUnknown int = iota
	PortTypeOther
	PortTypeNotApplicable
)

const (
	EnabledEnabled         Enabled = 1
	EnabledDisabled        Enabled = 2
	EnabledEnabledForDebug Enabled = 3
)

const (
	AuthenticationProtocolEAPTLS AuthenticationProtocol = iota
	AuthenticationProtocolEAPTTLS_MSCHAPv2
	AuthenticationProtocolPEAPv0_EAPMSCHAPv2
	AuthenticationProtocolPEAPv1_EAPGTC
	AuthenticationProtocolEAPFAST_MSCHAPv2
	AuthenticationProtocolEAPFAST_GTC
	AuthenticationProtocolEAP_MD5
	AuthenticationProtocolEAP_PSK
	AuthenticationProtocolEAP_SIM
	AuthenticationProtocolEAP_AKA
	AuthenticationProtocolEAPFAST_TLS
)

const (
	ConcreteJobStateNew          ConcreteJobState = 2
	ConcreteJobStateStarting     ConcreteJobState = 3
	ConcreteJobStateRunning      ConcreteJobState = 4
	ConcreteJobStateSuspended    ConcreteJobState = 5
	ConcreteJobStateShuttingDown ConcreteJobState = 6
	ConcreteJobStateCompleted    ConcreteJobState = 7
	ConcreteJobStateTerminated   ConcreteJobState = 8
	ConcreteJobStateKilled       ConcreteJobState = 9
	ConcreteJobStateException    ConcreteJobState = 10
	ConcreteJobStateService      ConcreteJobState = 11
	ConcreteJobStateQueryPending ConcreteJobState = 12
)
const (
	UnknownCS CommunicationStatus = iota
	NotAvailableCS
	CommunicationOK
	LostCommunication
	NoContact
)
const (
	NotAvailableDS DetailedStatus = iota
	NoAdditionalInformation
	Stressed
	PredictiveFailure
	NonRecoverableError
	SupportingEntityInError
)
const (
	UnknownOS OperatingStatus = iota
	NotAvailableOS
	Servicing
	Starting
	Stopping
	Stopped
	Aborted
	Dormant
	Completed
	Migrating
	Emigrating
	Immigrating
	Snapshotting
	ShuttingDown
	InTest
	Transitioning
	InService
)
const (
	UnknownPS PrimaryStatus = iota
	OK
	Degraded
	ErrorPS
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
	UnknownRA RecoveryAction = iota
	Other
	DoNotContinue
	ContinueWithNextJob
	RerunJob
	RunRecoveryJob
)
const (
	LocalTime LocalOrUtcTime = iota + 1
	UTCTime
)
const (
	PowerManagementCapabilitiesUnknown PowerManagementCapabilitiesValues = iota
	PowerManagementCapabilitiesNotSupported
	PowerManagementCapabilitiesDisabled
	PowerManagementCapabilitiesEnabled
	PowerManagementCapabilitiesPowerSavingModesEnteredAutomatically
	PowerManagementCapabilitiesPowerStateSettable
	PowerManagementCapabilitiesPowerCyclingSupported
	PowerManagementCapabilitiesTimedPowerOnSupported
)
const (
	AvailabilityOther AvailabilityValues = iota + 1
	AvailabilityUnknown
	AvailabilityRunningFullPower
	AvailabilityWarning
	AvailabilityInTest
	AvailabilityNotApplicable
	AvailabilityPowerOff
	AvailabilityOffLine
	AvailabilityOffDuty
	AvailabilityDegraded
	AvailabilityNotInstalled
	AvailabilityInstallError
	AvailabilityPowerSaveUnknown
	AvailabilityPowerSaveLowPowerMode
	AvailabilityPowerSaveStandby
	AvailabilityPowerCycle
	AvailabilityPowerSaveWarning
	AvailabilityPaused
	AvailabilityNotReady
	AvailabilityNotConfigured
	AvailabilityQuiesced
)
const (
	StatusInfoOther StatusInfoValues = iota + 1
	StatusInfoUnknown
	StatusInfoEnabled
	StatusInfoDisabled
	StatusInfoNotApplicable
)
const (
	AdditionalAvailabilityOther AdditionalAvailabilityValues = iota + 1
	AdditionalAvailabilityUnknown
	AdditionalAvailabilityRunningFullPower
	AdditionalAvailabilityWarning
	AdditionalAvailabilityInTest
	AdditionalAvailabilityNotApplicable
	AdditionalAvailabilityPowerOff
	AdditionalAvailabilityOffLine
	AdditionalAvailabilityOffDuty
	AdditionalAvailabilityDegraded
	AdditionalAvailabilityNotInstalled
	AdditionalAvailabilityInstallError
	AdditionalAvailabilityPowerSaveUnknown
	AdditionalAvailabilityPowerSaveLowPowerMode
	AdditionalAvailabilityPowerSaveStandby
	AdditionalAvailabilityPowerCycle
	AdditionalAvailabilityPowerSaveWarning
	AdditionalAvailabilityPaused
	AdditionalAvailabilityNotReady
	AdditionalAvailabilityNotConfigured
	AdditionalAvailabilityQuiesced
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
	OverwritePolicyUnknown         OverwritePolicy = 0
	OverwritePolicyWrapsWhenFull   OverwritePolicy = 2
	OverwritePolicyNeverOverwrites OverwritePolicy = 7
)
const (
	LogStateUnknown       LogState = 0
	LogStateNormal        LogState = 2
	LogStateErasing       LogState = 3
	LogStateNotApplicable LogState = 4
)
const (
	CapabilitiesUnknown CapabilitiesValues = iota
	CapabilitiesOther
	CapabilitiesWriteRecordSupported
	CapabilitiesDeleteRecordSupported
	CapabilitiesCanMoveBackwardInLog
	CapabilitiesFreezeLogSupported
	CapabilitiesClearLogSupported
	CapabilitiesSupportsAddressingByOrdinalRecordNumber
	CapabilitiesVariableLengthRecordsSupported
	CapabilitiesVariableFormatsForRecords
	CapabilitiesCanFlagRecordsForOverwrite
)
const (
	LastChangeUnknown LastChange = iota
	LastChangeAdd
	LastChangeDelete
	LastChangeModify
	LastChangeLogCleared
)
const (
	CharacterSetUnknown CharacterSet = iota
	CharacterSetOther
	CharacterSetASCII
	CharacterSetUnicode
	CharacterSetISO2022
	CharacterSetISO8859
	CharacterSetExtendedUNIXCode
	CharacterSetUTF8
	CharacterSetUCS2
	CharacterSetBitmappedData
	CharacterSetOctetString
	CharacterSetDefinedByIndividualRecords
)

const (
	UpgradeMethodOther UpgradeMethod = iota + 1
	UpgradeMethodUnknown
	UpgradeMethodDaughterBoard
	UpgradeMethodZIFSocket
	UpgradeMethodReplacementPiggyBack
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
	UpgradeMethodSocketmPGA604
	UpgradeMethodSocketLGA771
	UpgradeMethodSocketLGA775
	UpgradeMethodSocketS1
	UpgradeMethodSocketAM2
	UpgradeMethodSocketF1207
	UpgradeMethodSocketLGA1366
	UpgradeMethodSocketG34
	UpgradeMethodSocketAM3
	UpgradeMethodSocketC32
	UpgradeMethodSocketLGA1156
	UpgradeMethodSocketLGA1567
	UpgradeMethodSocketPGA988A
	UpgradeMethodSocketBGA1288
	UpgradeMethodrPGA988B
	UpgradeMethodBGA1023
	UpgradeMethodBGA1224
	UpgradeMethodLGA1155
	UpgradeMethodLGA1356
	UpgradeMethodLGA2011
	UpgradeMethodSocketFS1
	UpgradeMethodSocketFS2
	UpgradeMethodSocketFM1
	UpgradeMethodSocketFM2
	UpgradeMethodSocketLGA20113
	UpgradeMethodSocketLGA13563
	UpgradeMethodSocketLGA1150
	UpgradeMethodSocketBGA1168
	UpgradeMethodSocketBGA1234
	UpgradeMethodSocketBGA1364
	UpgradeMethodSocketAM4
	UpgradeMethodSocketLGA1151
	UpgradeMethodSocketBGA1356
	UpgradeMethodSocketBGA1440
	UpgradeMethodSocketBGA1515
	UpgradeMethodSocketLGA36471
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
	UpgradeMethodSocketLGA1211
	UpgradeMethodSocketLGA2422
	UpgradeMethodSocketLGA5773
	UpgradeMethodSocketBGA5773
)
const (
	CPUStatusUnknown CPUStatus = iota
	CPUStatusEnabled
	CPUStatusDisabledByUser
	CPUStatusDisabledByBIOS
	CPUStatusIdle
	CPUStatusOther
)

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
	HealthStateUnknown             HealthState = 0
	HealthStateOK                  HealthState = 5
	HealthStateDegradedWarning     HealthState = 10
	HealthStateMinorFailure        HealthState = 15
	HealthStateMajorFailure        HealthState = 20
	HealthStateCriticalFailure     HealthState = 25
	HealthStateNonRecoverableError HealthState = 30
)

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledbutOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)

const (
	RequestedStateEnabled RequestedState = iota + 2
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
	RequestedStateUnknown RequestedState = 0
)

const (
	EnabledDefaultEnabled           EnabledDefault = 2
	EnabledDefaultDisabled          EnabledDefault = 3
	EnabledDefaultNotApplicable     EnabledDefault = 5
	EnabledDefaultEnabledbutOffline EnabledDefault = 6
	EnabledDefaultNoDefault         EnabledDefault = 7
	EnabledDefaultQuiesce           EnabledDefault = 9
)
