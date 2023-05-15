/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package models

import (
	"encoding/xml"
	"time"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

type SettingData struct {
	ManagedElement
	InstanceID string `xml:"InstanceID,omitempty"` // MaxLen=256
}

type ManagedElement struct {
	XMLName     xml.Name `xml:"h:ManagedElement"`
	Caption     string   `xml:"h:Caption,omitempty"`     // MaxLen=64
	Description string   `xml:"h:Description,omitempty"` // MaxLen=256
	ElementName string   `xml:"h:ElementName,omitempty"` // MaxLen=256
}
type Collection struct {
	ManagedElement
}

type Role struct {
	Collection
	CreationClassName   string
	Name                string
	CommonName          string
	RoleCharacteristics []RoleCharacteristics
}

type RoleCharacteristics int

const (
	RoleCharacteristicsStatic RoleCharacteristics = 2
	RoleCharacteristicsOpaque RoleCharacteristics = 3
)

type MediaAccessDevice struct {
	LogicalDevice
	Capabilities []CapabilitiesValues // ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} Values={Unknown, Other, Sequential Access, Random Access, Supports Writing, Encryption, Compression, Supports Removeable Media, Manual Cleaning, Automatic Cleaning, SMART Notification, Supports Dual Sided Media, Predismount Eject Not Required} ArrayType=Indexed
	MaxMediaSize int
	Security     SecurityValues // ValueMap={1, 2, 3, 4, 5, 6, 7} Values={Other, Unknown, None, Read Only, Locked Out, Boot Bypass, Boot Bypass and Read Only}
}

type CapabilitiesValues int

const (
	CapabilitiesValuesUnknown CapabilitiesValues = iota
	CapabilitiesValuesOther
	CapabilitiesValuesSequentialAccess
	CapabilitiesValuesRandomAccess
	CapabilitiesValuesSupportsWriting
	CapabilitiesValuesEncryption
	CapabilitiesValuesCompression
	CapabilitiesValuesSupportsRemoveableMedia
	CapabilitiesValuesManualCleaning
	CapabilitiesValuesAutomaticCleaning
	CapabilitiesValuesSMARTNotification
	CapabilitiesValuesSupportsDualSidedMedia
	CapabilitiesValuesPredismountEjectNotRequired
)

type SecurityValues int

const (
	SecurityValuesOther SecurityValues = iota + 1
	SecurityValuesUnknown
	SecurityValuesNone
	SecurityValuesReadOnly
	SecurityValuesLockedOut
	SecurityValuesBootBypass
	SecurityValuesBootBypassAndReadOnly
)

type LogicalPort struct {
	LogicalDevice
	Speed            int
	MaxSpeed         int
	RequestedSpeed   int
	UsageRestriction UsageRestriction
	PortType         PortType
	OtherPortType    string
}
type UsageRestriction int

const (
	UsageRestrictionUnknown int = iota
	UsageRestrictionFrontEndOnly
	UsageRestrictionBackEndOnly
	UsageRestrictionNotRestricted
)

type PortType int

const (
	PortTypeUnknown int = iota
	PortTypeOther
	PortTypeNotApplicable
)

type NetworkPort struct {
	LogicalPort
	PortNumber                       int
	LinkTechnology                   LinkTechnology
	OtherLinkTechnology              string
	PermanentAddress                 string
	NetworkAddresses                 []string
	FullDuplex                       bool
	AutoSense                        bool
	SupportedMaximumTransmissionUnit int
	ActiveMaximumTransmissionUnit    int
}

type LinkTechnology int

const (
	LinkTechnologyUnknown LinkTechnology = iota
	LinkTechnologyOther
	LinkTechnologyEthernet
	LinkTechnologyIB
	LinkTechnologyFC
	LinkTechnologyFDDI
	LinkTechnologyATM
	LinkTechnologyTokenRing
	LinkTechnologyFrameRelay
	LinkTechnologyInfrared
	LinkTechnologyBlueTooth
	LinkTechnologyWirelessLAN
)

type EthernetPort struct {
	NetworkPort
}

type WiFiPort struct {
	NetworkPort
}

// TODO:fix types
type Dependency struct {
	Antecedent interface{}
	Dependent  interface{}
}

type SystemPackaging struct {
	Dependency
}

type ComputerSystemPackage struct {
	SystemPackaging
	PlatformGuid string //maxlen 40
}

type PhysicalElement struct {
	ManagedSystemElement
	Tag                  string // MaxLen=256
	CreationClassName    string // MaxLen=256
	Manufacturer         string // MaxLen=256
	Model                string // MaxLen=256
	Sku                  string // MaxLen=64
	SerialNumber         string // MaxLen=256
	Version              string // MaxLen=64
	PartNumber           string // MaxLen=256
	OtherIdentifyingInfo string // MaxLen=256
	PoweredOn            bool
	ManufactureDate      time.Time
	VendorEquipmentType  string // MaxLen=256
	UserTracking         string // MaxLen=256
	CanBeFRUed           bool
}

type PhysicalComponent struct {
	PhysicalElement
	RemovalConditions RemovalConditions
	Removable         bool
	Replaceable       bool
	HotSwappable      bool
}
type Chip struct {
	PhysicalComponent
	Tag               string
	CreationClassName string
	ElementName       string
	Manufacturer      string
	Version           string
	CanBeFRUed        bool
}

type PhysicalMemory struct {
	Chip
	FormFactor                 int
	MemoryType                 MemoryType
	Speed                      int
	Capacity                   int
	BankLabel                  string
	ConfiguredMemoryClockSpeed int
	IsSpeedInMhz               bool
	MaxMemorySpeed             int
}

type PhysicalPackage struct {
	PhysicalElement
	PackageType PackageType
}

type Card struct {
	PhysicalPackage
}

/**
 * Enabled:1 | Disabled:2 | Enabled For Debug:3
 */
type Enabled uint8

const (
	EnabledEnabled         Enabled = 1
	EnabledDisabled        Enabled = 2
	EnabledEnabledForDebug Enabled = 3
)

type PowerActionResponse struct {
	RequestPowerStateChange_OUTPUT wsman.ReturnValue
}
type PhysicalFrame struct {
	PhysicalPackage
	VendorCompatibilityStrings []string
	OtherPackageType           string
	Weight                     int
	Width                      int
	Depth                      int
	Height                     int
	RemovalConditions          RemovalConditions
	Removable                  bool
	Replaceable                bool
	HotSwappable               bool
	CableManagementStrategy    string
	ServicePhilosophy          ServicePhilosophy
	ServiceDescriptions        []string
	LockPresent                bool
	AudibleAlarm               bool
	VisibleAlarm               bool
	SecurityBreach             SecurityBreach
	BreachDescription          string
	IsLocked                   bool
}

type Chassis struct {
	PhysicalFrame
	ChassisPackageType ChassisPackageType
}

type LogicalElement struct {
	ManagedSystemElement
}

type SoftwareElement struct {
	LogicalElement
	Version               string
	SoftwareElementState  SoftwareElementState
	SoftwareElementId     string
	TargetOperatingSystem TargetOperatingSystem
	OtherTargetOs         string
	Manufacturer          string
	BuildNumber           string
	SerialNumber          string
	CodeSet               string
	IdentificationCode    string
	LanguageEdition       string
}
type BIOSElement struct {
	SoftwareElement
	PrimaryBIOS bool
	ReleaseDate time.Time
}
type BootSettingData struct {
	SettingData
	OwningEntity string // MaxLen=256
}

// SharedCredential represents a shared credential for a device.
type SharedCredential struct {
	InstanceID string
	RemoteID   string
	Secret     string
	Algorithm  string
	Protocol   string
}

// IEEE8021xSettings represents the IEEE 802.1x settings for a network interface.
type IEEE8021xSettings struct {
	XMLName                         xml.Name                        `xml:"h:IEEE8021xSettingsInput,omitempty"`
	H                               string                          `xml:"xmlns:q,attr"`
	ElementName                     string                          `xml:"q:ElementName,omitempty"`
	InstanceID                      string                          `xml:"q:InstanceID,omitempty"`
	AuthenticationProtocol          AuthenticationProtocol          `xml:"q:AuthenticationProtocol"`
	RoamingIdentity                 string                          `xml:"q:RoamingIdentity,omitempty"`
	ServerCertificateName           string                          `xml:"q:ServerCertificateName,omitempty"`
	ServerCertificateNameComparison ServerCertificateNameComparison `xml:"q:ServerCertificateNameComparison,omitempty"`
	Username                        string                          `xml:"q:Username,omitempty"`
	Password                        string                          `xml:"q:Password,omitempty"`
	Domain                          string                          `xml:"q:Domain,omitempty"`
	ProtectedAccessCredential       string                          `xml:"q:ProtectedAccessCredential,omitempty"`
	PACPassword                     string                          `xml:"q:PACPassword,omitempty"`
	PSK                             string                          `xml:"q:PSK,omitempty"`
}

type AuthenticationProtocol int

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

type WiFiEndpointSettings struct {
	XMLName xml.Name `xml:"h:WiFiEndpointSettingsInput"`
	H       string   `xml:"xmlns:q,attr"`
	// SettingData
	ElementName          string               `xml:"q:ElementName"`
	InstanceID           string               `xml:"q:InstanceID"`
	AuthenticationMethod AuthenticationMethod `xml:"q:AuthenticationMethod"`
	EncryptionMethod     EncryptionMethod     `xml:"q:EncryptionMethod"`
	SSID                 string               `xml:"q:SSID"` // Max Length 32
	Priority             int                  `xml:"q:Priority"`
	PSKPassPhrase        string               `xml:"q:PSKPassPhrase"` // Min Length 8 Max Length 63
	BSSType              BSSType              `xml:"q:BSSType,omitempty"`
	Keys                 []string             `xml:"q:Keys,omitempty"` // OctetString ArrayType=Indexed Max Length 256
	KeyIndex             int                  `xml:"q:KeyIndex,omitempty"`
	PSKValue             uint64               `xml:"q:PSKValue,omitempty"` // OctetString
}

// BootSourceSetting represents the boot source settings for a device.
type BootSourceSetting struct {
	SettingData
	ElementName          string
	InstanceID           BootConfigSettingInstanceID
	StructuredBootString string
	BIOSBootString       string
	BootString           string
	FailThroughSupported FailThroughSupported
}

type Service struct {
	EnabledLogicalElement
	SystemCreationClassName string
	SystemName              string
	CreationClassName       string
	PrimaryOwnerName        string
	PrimaryOwnerContact     string
	StartMode               string
	Started                 bool
}

type SecurityService struct {
	Service
}

type AuthenticationService struct {
	SecurityService
}
type NetworkPortConfigurationService struct {
	Service
}
type CredentialManagementService struct {
	AuthenticationService
	// InstanceID is an optional property that may be used to opaquely and uniquely identify an instance of this class within the scope of the instantiating Namespace . . .
	InstanceID string // MaxLen=256
}

type EnabledLogicalElement struct {
	LogicalElement
	EnabledState          EnabledLogicalElementEnabledState
	OtherEnabledState     string
	RequestedState        EnabledLogicalElementRequestedState
	EnabledDefault        EnabledLogicalElementEnabledDefault
	TimeOfLastStateChange time.Time // You may use "time.Time" from the standard library
}
type LogicalDevice struct {
	SystemCreationClassName     string                                           `json:"SystemCreationClassName,omitempty"`
	SystemName                  string                                           `json:"SystemName,omitempty"`
	CreationClassName           string                                           `json:"CreationClassName,omitempty"`
	DeviceId                    string                                           `json:"DeviceId,omitempty"`
	PowerManagementSupported    bool                                             `json:"PowerManagementSupported,omitempty"`
	PowerManagementCapabilities []LogicalDevicePowerManagementCapabilitiesValues `json:"PowerManagementCapabilities,omitempty"`
	Availability                LogicalDeviceAvailabilityValues                  `json:"Availability,omitempty"`
	StatusInfo                  LogicalDeviceStatusInfoValues                    `json:"StatusInfo,omitempty"`
	LastErrorCode               int                                              `json:"LastErrorCode,omitempty"`
	ErrorDescription            string                                           `json:"ErrorDescription,omitempty"`
	ErrorCleared                bool                                             `json:"ErrorCleared,omitempty"`
	OtherIdentifyingInfo        []string                                         `json:"OtherIdentifyingInfo,omitempty"`
	PowerOnHours                int                                              `json:"PowerOnHours,omitempty"`
	TotalPowerOnHours           int                                              `json:"TotalPowerOnHours,omitempty"`
	IdentifyingDescriptions     []string                                         `json:"IdentifyingDescriptions,omitempty"`
	AdditionalAvailability      []LogicalDeviceAdditionalAvailabilityValues      `json:"AdditionalAvailability,omitempty"`
	MaxQuiesceTime              int                                              `json:"MaxQuiesceTime,omitempty"`
}
type KVMRedirectionSAP struct {
	Name                    string
	CreationClassName       string
	SystemName              string
	SystemCreationClassName string
	ElementName             string
	EnabledState            KVMRedirectionSAPEnabledState
	RequestedState          KVMRedirectionSAPRequestedState
	KVMProtocol             KVMRedirectionSAPKVMProtocol
}
type KVMRedirectionSAPResponse struct {
	CIM_KVMRedirectionSAP KVMRedirectionSAP
}
type KVMRedirectionSAPEnabledState int

const (
	KVMRedirectionSAPEUnknown KVMRedirectionSAPEnabledState = iota
	KVMRedirectionSAPEOther
	KVMRedirectionSAPEEnabled
	KVMRedirectionSAPEDisabled
	KVMRedirectionSAPEShuttingDown
	KVMRedirectionSAPENotApplicable
	KVMRedirectionSAPEEnabledButOffline
	KVMRedirectionSAPENInTest
	KVMRedirectionSAPEDeferred
	KVMRedirectionSAPEQuiesce
	KVMRedirectionSAPEStarting
)

type KVMRedirectionSAPRequestedState int

const (
	KVMRedirectionSAPRUnknown KVMRedirectionSAPRequestedState = iota
	KVMRedirectionSAPREnabled
	KVMRedirectionSAPRDisabled
	KVMRedirectionSAPRShutDown
	KVMRedirectionSAPRNoChange
	KVMRedirectionSAPOffline
	KVMRedirectionSAPRTest
	KVMRedirectionSAPRDeferred
	KVMRedirectionSAPRQuiesce
	KVMRedirectionSAPRReboot
	KVMRedirectionSAPRReset
	KVMRedirectionSAPRNotApplicable
)

type KVMRedirectionSAPRequestedStateInputs int

const (
	KVMRedirectionSAPEEnabledInput KVMRedirectionSAPRequestedStateInputs = iota + 2
	KVMRedirectionSAPEDisabledInput
	KVMRedirectionSAPRShutDownInput
	KVMRedirectionSAPOfflineInput
	KVMRedirectionSAPRTestInput
	KVMRedirectionSAPEDeferredInput
	KVMRedirectionSAPEQuiesceInput
	KVMRedirectionSAPRRebootInput
	KVMRedirectionSAPRResetInput
)

type KVMRedirectionSAPKVMProtocol int

const (
	KVMRedirectionSAPKUnknown KVMRedirectionSAPKVMProtocol = iota
	KVMRedirectionSAPKOther
	KVMRedirectionSAPKRaw
	KVMRedirectionSAPKRDP
	KVMRedirectionSAPKVNC_RFB
)

type Job struct {
	LogicalElement
	InstanceId          string
	CommunicationStatus CommunicationStatus
	DetailedStatus      DetailedStatus
	OperatingStatus     OperatingStatus
	PrimaryStatus       PrimaryStatus
	JobStatus           string
	TimeSubmitted       time.Time
	ScheduledStartTime  time.Time
	StartTime           time.Time
	ElapsedTime         time.Time
	JobRunTimes         int
	RunMonth            RunMonth
	RunDay              int
	RunDayOfWeek        RunDayOfWeek
	RunStartInterval    time.Time
	LocalOrUtcTime      LocalOrUtcTime
	Notify              string
	Owner               string
	Priority            int
	PercentComplete     int
	DeleteOnCompletion  bool
	ErrorCode           int
	ErrorDescription    string
	RecoveryAction      RecoveryAction
	OtherRecoveryAction string
}
type ConcreteJob struct {
	Job
	UntilTime             time.Time
	JobState              ConcreteJobState
	TimeOfLastStateChange time.Time
	TimeBeforeRemoval     time.Time
}
type Processor struct {
	LogicalDevice
	Role                   string
	Family                 int
	OtherFamilyDescription string
	UpgradeMethod          UpgradeMethod
	MaxClockSpeed          int
	CurrentClockSpeed      int
	Stepping               string
	CPUStatus              CPUStatus
	ExternalBusClockSpeed  int
}
type UpgradeMethod int

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

type CPUStatus int

const (
	CPUStatusUnknown CPUStatus = iota
	CPUStatusEnabled
	CPUStatusDisabledByUser
	CPUStatusDisabledByBIOS
	CPUStatusIdle
	CPUStatusOther
)

type ConcreteJobState int

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

type CommunicationStatus int

const (
	UnknownCS CommunicationStatus = iota
	NotAvailableCS
	CommunicationOK
	LostCommunication
	NoContact
)

type DetailedStatus int

const (
	NotAvailableDS DetailedStatus = iota
	NoAdditionalInformation
	Stressed
	PredictiveFailure
	NonRecoverableError
	SupportingEntityInError
)

type OperatingStatus int

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

type PrimaryStatus int

const (
	UnknownPS PrimaryStatus = iota
	OK
	Degraded
	ErrorPS
)

type RunMonth int

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

type RunDayOfWeek int

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

type RecoveryAction int

const (
	UnknownRA RecoveryAction = iota
	Other
	DoNotContinue
	ContinueWithNextJob
	RerunJob
	RunRecoveryJob
)

type LocalOrUtcTime int

const (
	LocalTime LocalOrUtcTime = iota + 1
	UTCTime
)

type LogicalDevicePowerManagementCapabilitiesValues int

const (
	LogicalDevicePowerManagementCapabilitiesUnknown LogicalDevicePowerManagementCapabilitiesValues = iota
	LogicalDevicePowerManagementCapabilitiesNotSupported
	LogicalDevicePowerManagementCapabilitiesDisabled
	LogicalDevicePowerManagementCapabilitiesEnabled
	LogicalDevicePowerManagementCapabilitiesPowerSavingModesEnteredAutomatically
	LogicalDevicePowerManagementCapabilitiesPowerStateSettable
	LogicalDevicePowerManagementCapabilitiesPowerCyclingSupported
	LogicalDevicePowerManagementCapabilitiesTimedPowerOnSupported
)

type LogicalDeviceAvailabilityValues int

const (
	LogicalDeviceAvailabilityOther LogicalDeviceAvailabilityValues = iota + 1
	LogicalDeviceAvailabilityUnknown
	LogicalDeviceAvailabilityRunningFullPower
	LogicalDeviceAvailabilityWarning
	LogicalDeviceAvailabilityInTest
	LogicalDeviceAvailabilityNotApplicable
	LogicalDeviceAvailabilityPowerOff
	LogicalDeviceAvailabilityOffLine
	LogicalDeviceAvailabilityOffDuty
	LogicalDeviceAvailabilityDegraded
	LogicalDeviceAvailabilityNotInstalled
	LogicalDeviceAvailabilityInstallError
	LogicalDeviceAvailabilityPowerSaveUnknown
	LogicalDeviceAvailabilityPowerSaveLowPowerMode
	LogicalDeviceAvailabilityPowerSaveStandby
	LogicalDeviceAvailabilityPowerCycle
	LogicalDeviceAvailabilityPowerSaveWarning
	LogicalDeviceAvailabilityPaused
	LogicalDeviceAvailabilityNotReady
	LogicalDeviceAvailabilityNotConfigured
	LogicalDeviceAvailabilityQuiesced
)

type LogicalDeviceStatusInfoValues int

const (
	LogicalDeviceStatusInfoOther LogicalDeviceStatusInfoValues = iota + 1
	LogicalDeviceStatusInfoUnknown
	LogicalDeviceStatusInfoEnabled
	LogicalDeviceStatusInfoDisabled
	LogicalDeviceStatusInfoNotApplicable
)

type LogicalDeviceAdditionalAvailabilityValues int

const (
	LogicalDeviceAdditionalAvailabilityOther LogicalDeviceAdditionalAvailabilityValues = iota + 1
	LogicalDeviceAdditionalAvailabilityUnknown
	LogicalDeviceAdditionalAvailabilityRunningFullPower
	LogicalDeviceAdditionalAvailabilityWarning
	LogicalDeviceAdditionalAvailabilityInTest
	LogicalDeviceAdditionalAvailabilityNotApplicable
	LogicalDeviceAdditionalAvailabilityPowerOff
	LogicalDeviceAdditionalAvailabilityOffLine
	LogicalDeviceAdditionalAvailabilityOffDuty
	LogicalDeviceAdditionalAvailabilityDegraded
	LogicalDeviceAdditionalAvailabilityNotInstalled
	LogicalDeviceAdditionalAvailabilityInstallError
	LogicalDeviceAdditionalAvailabilityPowerSaveUnknown
	LogicalDeviceAdditionalAvailabilityPowerSaveLowPowerMode
	LogicalDeviceAdditionalAvailabilityPowerSaveStandby
	LogicalDeviceAdditionalAvailabilityPowerCycle
	LogicalDeviceAdditionalAvailabilityPowerSaveWarning
	LogicalDeviceAdditionalAvailabilityPaused
	LogicalDeviceAdditionalAvailabilityNotReady
	LogicalDeviceAdditionalAvailabilityNotConfigured
	LogicalDeviceAdditionalAvailabilityQuiesced
)

type Credential struct {
	ManagedElement
	Issued  time.Time // The date and time when the credential was issued. Default is current time
	Expires time.Time // The date and time when the credential expires (and is not appropriate for use for authentication/authorization). Default is '99991231235959.999999+999'
}

type CredentialContext struct {
	ElementInContext        Credential
	ElementProvidingContext ManagedElement
}

type ManagedSystemElement struct {
	ManagedElement
	InstallDate        time.Time
	Name               string
	OperationalStatus  ManagedSystemElementOperationalStatus
	StatusDescriptions []string
	Status             string
	HealthState        ManagedSystemElementHealthState
}

type ServiceAvailableToElement struct {
	ServiceProvided ServiceProvider
	UserOfService   UserOfService
}

type ServiceProvider struct {
	Address             string
	ReferenceParameters ReferenceParameters
}

type UserOfService struct {
	Address             string
	ReferenceParameters ReferenceParameters
}

type ReferenceParameters struct {
	ResourceURI string      `xml:"w:ResourceURI,omitempty"`
	SelectorSet SelectorSet `xml:"w:SelectorSet,omitempty"`
}

type ReferenceParameters_OUTPUT struct {
	ResourceURI string             `xml:"ResourceURI,omitempty"`
	SelectorSet SelectorSet_OUTPUT `xml:"SelectorSet,omitempty"`
}
type SelectorSet_OUTPUT struct {
	XMLName  xml.Name `xml:"SelectorSet,omitempty"`
	Selector []wsman.Selector_OUTPUT
}

type SelectorSet struct {
	XMLName  xml.Name `xml:"w:SelectorSet,omitempty"`
	Selector []wsman.Selector
}

type AssociatedPowerManagementService struct {
	ServiceAvailableToElement
	CIMAssociatedPowerManagementService CIMAssociatedPowerManagementServiceItem
}

type CIMAssociatedPowerManagementServiceItem struct {
	ServiceAvailableToElement
	AvailableRequestedPowerStates []string
	PowerState                    string
}

type SoftwareIdentity struct {
	LogicalElement
	CIMSoftwareIdentity []CIMSoftwareIdentityItem
}

type CIMSoftwareIdentityItem struct {
	LogicalElement
	InstanceID    string
	VersionString string
	IsEntity      bool
}

type Log struct {
	EnabledLogicalElement
	MaxNumberOfRecords     int
	CurrentNumberOfRecords int
	OverwritePolicy        TypesLogOverwritePolicy
	LogState               TypesLogLogState
}

type MessageLog struct {
	Log
	CreationClassName        string
	Capabilities             []TypesMessageLogCapabilitiesValues
	CapabilitiesDescriptions []string
	MaxLogSize               int
	SizeOfHeader             int
	HeaderFormat             string
	MaxRecordSize            int
	SizeOfRecordHeader       int
	RecordHeaderFormat       string
	OtherPolicyDescription   string
	TimeWhenOutdated         time.Time
	PercentageNearFull       int
	LastChange               TypesMessageLogLastChange
	TimeOfLastChange         time.Time
	RecordLastChanged        int
	IsFrozen                 bool
	CharacterSet             TypesMessageLogCharacterSet
}

type TypesLogOverwritePolicy int

const (
	TypesLogOverwritePolicyUnknown         TypesLogOverwritePolicy = 0
	TypesLogOverwritePolicyWrapsWhenFull   TypesLogOverwritePolicy = 2
	TypesLogOverwritePolicyNeverOverwrites TypesLogOverwritePolicy = 7
)

type TypesLogLogState int

const (
	TypesLogLogStateUnknown       TypesLogLogState = 0
	TypesLogLogStateNormal        TypesLogLogState = 2
	TypesLogLogStateErasing       TypesLogLogState = 3
	TypesLogLogStateNotApplicable TypesLogLogState = 4
)

type TypesMessageLogCapabilitiesValues int

const (
	TypesMessageLogCapabilitiesUnknown TypesMessageLogCapabilitiesValues = iota
	TypesMessageLogCapabilitiesOther
	TypesMessageLogCapabilitiesWriteRecordSupported
	TypesMessageLogCapabilitiesDeleteRecordSupported
	TypesMessageLogCapabilitiesCanMoveBackwardInLog
	TypesMessageLogCapabilitiesFreezeLogSupported
	TypesMessageLogCapabilitiesClearLogSupported
	TypesMessageLogCapabilitiesSupportsAddressingByOrdinalRecordNumber
	TypesMessageLogCapabilitiesVariableLengthRecordsSupported
	TypesMessageLogCapabilitiesVariableFormatsForRecords
	TypesMessageLogCapabilitiesCanFlagRecordsForOverwrite
)

type TypesMessageLogLastChange int

const (
	TypesMessageLogLastChangeUnknown TypesMessageLogLastChange = iota
	TypesMessageLogLastChangeAdd
	TypesMessageLogLastChangeDelete
	TypesMessageLogLastChangeModify
	TypesMessageLogLastChangeLogCleared
)

type TypesMessageLogCharacterSet int

const (
	TypesMessageLogCharacterSetUnknown TypesMessageLogCharacterSet = iota
	TypesMessageLogCharacterSetOther
	TypesMessageLogCharacterSetASCII
	TypesMessageLogCharacterSetUnicode
	TypesMessageLogCharacterSetISO2022
	TypesMessageLogCharacterSetISO8859
	TypesMessageLogCharacterSetExtendedUNIXCode
	TypesMessageLogCharacterSetUTF8
	TypesMessageLogCharacterSetUCS2
	TypesMessageLogCharacterSetBitmappedData
	TypesMessageLogCharacterSetOctetString
	TypesMessageLogCharacterSetDefinedByIndividualRecords
)

// Define the enums for the respective fields as types
type ManagedSystemElementOperationalStatus int

const (
	ManagedSystemElementOperationalStatusUnknown ManagedSystemElementOperationalStatus = iota
	ManagedSystemElementOperationalStatusOther
	ManagedSystemElementOperationalStatusOk
	ManagedSystemElementOperationalStatusDegraded
	ManagedSystemElementOperationalStatusStressed
	ManagedSystemElementOperationalStatusPredictiveFailure
	ManagedSystemElementOperationalStatusError
	ManagedSystemElementOperationalStatusNonRecoverableError
	ManagedSystemElementOperationalStatusStarting
	ManagedSystemElementOperationalStatusStopping
	ManagedSystemElementOperationalStatusStopped
	ManagedSystemElementOperationalStatusInService
	ManagedSystemElementOperationalStatusNoContact
	ManagedSystemElementOperationalStatusLostCommunication
	ManagedSystemElementOperationalStatusAborted
	ManagedSystemElementOperationalStatusDormant
	ManagedSystemElementOperationalStatusSupportingEntityInError
	ManagedSystemElementOperationalStatusCompleted
	ManagedSystemElementOperationalStatusPowerMode
	ManagedSystemElementOperationalStatusRelocating
)

type ManagedSystemElementHealthState int

const (
	ManagedSystemElementHealthStateUnknown               ManagedSystemElementHealthState = 0
	ManagedSystemElementHealthStateOk                    ManagedSystemElementHealthState = 5
	ManagedSystemElementHealthStateDegradedWarning       ManagedSystemElementHealthState = 10
	ManagedSystemElementHealthStateMinorFailure          ManagedSystemElementHealthState = 15
	ManagedSystemElementHealthStateMajorFailure          ManagedSystemElementHealthState = 20
	ManagedSystemElementHealthStateCriticalFailure       ManagedSystemElementHealthState = 25
	ManagedSystemElementHealthStateNonRecoverableFailure ManagedSystemElementHealthState = 30
)

// Define the enums for the respective fields as types
type EnabledLogicalElementEnabledState int

const (
	EnabledLogicalElementEnabledStateUnknown EnabledLogicalElementEnabledState = iota
	EnabledLogicalElementEnabledStateOther
	EnabledLogicalElementEnabledStateEnabled
	EnabledLogicalElementEnabledStateDisabled
	EnabledLogicalElementEnabledStateShuttingDown
	EnabledLogicalElementEnabledStateNotApplicable
	EnabledLogicalElementEnabledStateEnabledButOffline
	EnabledLogicalElementEnabledStateInTest
	EnabledLogicalElementEnabledStateDeferred
	EnabledLogicalElementEnabledStateQuiesce
	EnabledLogicalElementEnabledStateStarting
)

type EnabledLogicalElementRequestedState int

const (
	EnabledLogicalElementRequestedStateUnknown EnabledLogicalElementRequestedState = iota
	_
	EnabledLogicalElementRequestedStateEnabled
	EnabledLogicalElementRequestedStateDisable
	EnabledLogicalElementRequestedStateShutDown
	EnabledLogicalElementRequestedStateNoChange
	EnabledLogicalElementRequestedStateOffline
	EnabledLogicalElementRequestedStateTest
	EnabledLogicalElementRequestedStateDeferred
	EnabledLogicalElementRequestedStateQuiesce
	EnabledLogicalElementRequestedStateReboot
	EnabledLogicalElementRequestedStateReset
	EnabledLogicalElementRequestedStateNotApplicable
)

type EnabledLogicalElementEnabledDefault int

const (
	EnabledLogicalElementEnabledDefaultEnabled           EnabledLogicalElementEnabledDefault = 2
	EnabledLogicalElementEnabledDefaultDisabled          EnabledLogicalElementEnabledDefault = 3
	EnabledLogicalElementEnabledDefaultNotApplicable     EnabledLogicalElementEnabledDefault = 5
	EnabledLogicalElementEnabledDefaultEnabledButOffline EnabledLogicalElementEnabledDefault = 6
	EnabledLogicalElementEnabledDefaultNoDefault         EnabledLogicalElementEnabledDefault = 7
	EnabledLogicalElementEnabledDefaultQuiesce           EnabledLogicalElementEnabledDefault = 9
)

type BSSType int
type EncryptionMethod int
type AuthenticationMethod int
type BootConfigSettingInstanceID string
type FailThroughSupported int
type RemovalConditions int
type MemoryType int
type PackageType int
type ServicePhilosophy int
type SecurityBreach int
type ChassisPackageType int
type SoftwareElementState int
type TargetOperatingSystem int

// ServerCertificateNameComparison represents the ServerCertificateNameComparison type for IEEE8021xProfile.
type ServerCertificateNameComparison int

const (
	BSSType_Unknown        BSSType = 0
	BSSType_Independent    BSSType = 2
	BSSType_Infrastructure BSSType = 3

	EncryptionMethod_Other        EncryptionMethod = 1
	EncryptionMethod_WEP          EncryptionMethod = 2
	EncryptionMethod_TKIP         EncryptionMethod = 3
	EncryptionMethod_CCMP         EncryptionMethod = 4
	EncryptionMethod_None         EncryptionMethod = 5
	EncryptionMethod_DMTFReserved EncryptionMethod = 6

	AuthenticationMethod_Other          AuthenticationMethod = 1
	AuthenticationMethod_OpenSystem     AuthenticationMethod = 2
	AuthenticationMethod_SharedKey      AuthenticationMethod = 3
	AuthenticationMethod_WPA_PSK        AuthenticationMethod = 4
	AuthenticationMethod_WPA_IEEE8021x  AuthenticationMethod = 5
	AuthenticationMethod_WPA2_PSK       AuthenticationMethod = 6
	AuthenticationMethod_WPA2_IEEE8021x AuthenticationMethod = 7
	AuthenticationMethod_DMTFReserved   AuthenticationMethod = 8
	AuthenticationMethod_WPA3_SAE       AuthenticationMethod = 32768
	AuthenticationMethod_WPA3_OWE       AuthenticationMethod = 32769
	AuthenticationMethod_VendorReserved AuthenticationMethod = 32770
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

// PowerState represents the PowerState type in the PowerManagementService namespace.
