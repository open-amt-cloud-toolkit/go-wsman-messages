package models

import (
	"time"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

type AMTAuthenticateObject struct {
	Nonce                []int
	UUID                 []string
	FQDN                 string
	FWVersion            string
	AMTSVN               int
	SignatureMechanism   AMTAuthenticateObjectSignatureMechanism
	Signature            []int
	LengthOfCertificates []int
	Certificates         []int
}

type AMTAuthenticateObjectSignatureMechanism int

const (
	AMTAuthenticateObjectSignatureMechanismTLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 AMTAuthenticateObjectSignatureMechanism = 0
)

type MPServer struct {
	AccessInfo string
	InfoFormat MPServerInfoFormat
	Port       int
	AuthMethod MPServerAuthMethod
	Username   string
	Password   string
	CommonName string
}

type MPServerInfoFormat uint8

const (
	MPServerInfoFormatIPv4Address MPServerInfoFormat = 3
	MPServerInfoFormatIPv6Address MPServerInfoFormat = 4
	MPServerInfoFormatFQDN        MPServerInfoFormat = 201
)

type MPServerAuthMethod uint8

const (
	MPServerAuthMethodMutualAuthentication           MPServerAuthMethod = 1
	MPServerAuthMethodUsernamePasswordAuthentication MPServerAuthMethod = 2
)

type SystemDefensePolicy struct {
	models.ManagedElement
	PolicyName            string
	PolicyPrecedence      int
	AntiSpoofingSupport   SystemDefensePolicyAntiSpoofingSupport
	FilterCreationHandles []int
	TxDefaultDrop         bool
	TxDefaultMatchEvent   bool
	TxDefaultCount        bool
	RxDefaultDrop         bool
	RxDefaultMatchEvent   bool
	RxDefaultCount        bool
}

type SystemDefensePolicyAntiSpoofingSupport uint8

const (
	SystemDefensePolicyAntiSpoofingSupportOff                             SystemDefensePolicyAntiSpoofingSupport = 0
	SystemDefensePolicyAntiSpoofingSupportEventOnMatch                    SystemDefensePolicyAntiSpoofingSupport = 1
	SystemDefensePolicyAntiSpoofingSupportCount                           SystemDefensePolicyAntiSpoofingSupport = 2
	SystemDefensePolicyAntiSpoofingSupportCountingEventOnMatch            SystemDefensePolicyAntiSpoofingSupport = 3
	SystemDefensePolicyAntiSpoofingSupportOnWithoutCountingOrEventOnMatch SystemDefensePolicyAntiSpoofingSupport = 4
)

type BootCapabilities struct {
	models.ManagedElement
	AMT_BootCapabilities struct {
		ElementName                        string
		InstanceID                         string
		IDER                               bool
		SOL                                bool
		BIOSReflash                        bool
		BIOSSetup                          bool
		BIOSPause                          bool
		ForcePXEBoot                       bool
		ForceHardDriveBoot                 bool
		ForceHardDriveSafeModeBoot         bool
		ForceDiagnosticBoot                bool
		ForceCDorDVDBoot                   bool
		VerbosityScreenBlank               bool
		PowerButtonLock                    bool
		ResetButtonLock                    bool
		KeyboardLock                       bool
		SleepButtonLock                    bool
		UserPasswordBypass                 bool
		ForcedProgressEvents               bool
		VerbosityVerbose                   bool
		VerbosityQuiet                     bool
		ConfigurationDataReset             bool
		BIOSSecureBoot                     bool
		SecureErase                        bool
		ForceWinREBoot                     bool
		ForceUEFILocalPBABoot              bool
		ForceUEFIHTTPSBoot                 bool
		AMTSecureBootControl               bool
		UEFIWiFiCoExistenceAndProfileShare bool
		PlatformErase                      int
	}
}

type MessageLog struct {
	models.MessageLog
}

type EVENT_DATA struct {
	DeviceAddress   int                     `json:"DeviceAddress,omitempty"`
	EventSensorType int                     `json:"EventSensorType,omitempty"`
	EventType       int                     `json:"EventType,omitempty"`
	EventOffset     int                     `json:"EventOffset,omitempty"`
	EventSourceType int                     `json:"EventSourceType,omitempty"`
	EventSeverity   EVENT_DATAEventSeverity `json:"EventSeverity,omitempty"`
	SensorNumber    int                     `json:"SensorNumber,omitempty"`
	Entity          int                     `json:"Entity,omitempty"`
	EntityInstance  int                     `json:"EntityInstance,omitempty"`
	EventData       []int                   `json:"EventData,omitempty"`
	TimeStamp       time.Time               `json:"TimeStamp,omitempty"`
}

type EVENT_DATAEventSeverity int

const (
	EVENT_DATAEventSeverityUnspecified    EVENT_DATAEventSeverity = 0
	EVENT_DATAEventSeverityMonitor        EVENT_DATAEventSeverity = 1
	EVENT_DATAEventSeverityInformation    EVENT_DATAEventSeverity = 2
	EVENT_DATAEventSeverityOK             EVENT_DATAEventSeverity = 4
	EVENT_DATAEventSeverityNonCritical    EVENT_DATAEventSeverity = 8
	EVENT_DATAEventSeverityCritical       EVENT_DATAEventSeverity = 16
	EVENT_DATAEventSeverityNonRecoverable EVENT_DATAEventSeverity = 32
)

type AuditLogReadRecords struct {
	ReadRecordsOutput struct {
		TotalRecordCount string
		RecordsReturned  string
		EventRecords     []string
		ReturnValue      string
	}
}

type TLSProtocolEndpointCollection struct {
	models.Collection
}

// IEEE8021xProfile represents the IEEE8021xProfile interface extending the CIM.Models.SettingData struct.

type AuthorizationService struct {
	models.Service
	AllowHttpQopAuthOnly int
}

// type SignatureMechanism int

// const (
// 	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 SignatureMechanism = 0
// 	Reserved                                SignatureMechanism = 1
// )
