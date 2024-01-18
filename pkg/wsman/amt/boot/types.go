/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type SettingData struct {
	base message.Base
}

type Capabilities struct {
	base message.Base
}

// OUTPUTS
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                     xml.Name `xml:"Body"`
		BootSettingDataGetResponse  BootSettingDataResponse
		BootCapabilitiesGetResponse BootCapabilitiesResponse
		EnumerateResponse           common.EnumerateResponse
		PullResponse                PullResponse
	}
	PullResponse struct {
		XMLName               xml.Name                   `xml:"PullResponse"`
		BootSettingDataItems  []BootSettingDataResponse  `xml:"Items>AMT_BootSettingData"`
		BootCapabilitiesItems []BootCapabilitiesResponse `xml:"Items>AMT_BootCapabilities"`
	}
	BootSettingDataResponse struct {
		XMLName                  xml.Name          `xml:"AMT_BootSettingData"`
		InstanceID               string            `xml:"InstanceID,omitempty"`
		ElementName              string            `xml:"ElementName,omitempty"`
		OwningEntity             string            `xml:"OwningEntity,omitempty"`
		UseSOL                   bool              `xml:"UseSOL,omitempty"`
		UseSafeMode              bool              `xml:"UseSafeMode,omitempty"`
		ReflashBIOS              bool              `xml:"ReflashBIOS,omitempty"`
		BIOSSetup                bool              `xml:"BIOSSetup,omitempty"`
		BIOSPause                bool              `xml:"BIOSPause,omitempty"`
		LockPowerButton          bool              `xml:"LockPowerButton,omitempty"`
		LockResetButton          bool              `xml:"LockResetButton,omitempty"`
		LockKeyboard             bool              `xml:"LockKeyboard,omitempty"`
		LockSleepButton          bool              `xml:"LockSleepButton,omitempty"`
		UserPasswordBypass       bool              `xml:"UserPasswordBypass,omitempty"`
		ForcedProgressEvents     bool              `xml:"ForcedProgressEvents,omitempty"`
		FirmwareVerbosity        FirmwareVerbosity `xml:"FirmwareVerbosity,omitempty"`
		ConfigurationDataReset   bool              `xml:"ConfigurationDataReset,omitempty"`
		IDERBootDevice           IDERBootDevice    `xml:"IDERBootDevice,omitempty"`
		UseIDER                  bool              `xml:"UseIDER,omitempty"`
		EnforceSecureBoot        bool              `xml:"EnforceSecureBoot,omitempty"`
		BootMediaIndex           int               `xml:"BootMediaIndex,omitempty"`
		SecureErase              bool              `xml:"SecureErase,omitempty"`
		RSEPassword              string            `xml:"RSEPassword,omitempty"`
		OptionsCleared           bool              `xml:"OptionsCleared,omitempty"`           //readonly
		WinREBootEnabled         bool              `xml:"WinREBootEnabled,omitempty"`         //readonly
		UEFILocalPBABootEnabled  bool              `xml:"UEFILocalPBABootEnabled,omitempty"`  //readonly
		UEFIHTTPSBootEnabled     bool              `xml:"UEFIHTTPSBootEnabled,omitempty"`     //readonly
		SecureBootControlEnabled bool              `xml:"SecureBootControlEnabled,omitempty"` //readonly
		BootguardStatus          bool              `xml:"BootguardStatus,omitempty"`          //readonly
		BIOSLastStatus           []int             `xml:"BIOSLastStatus,omitempty"`           //readonly
		UEFIBootParametersArray  []int             `xml:"UEFIBootParametersArray,omitempty"`
		UEFIBootNumberOfParams   []int             `xml:"UEFIBootNumberOfParams,omitempty"`
		RPEEnabled               bool              `xml:"RPEEnabled,omitempty"`
		PlatformErase            bool              `xml:"PlatformErase,omitempty"`
	}

	BootCapabilitiesResponse struct {
		XMLName                            xml.Name `xml:"AMT_BootCapabilities"`
		InstanceID                         string   `xml:"InstanceID,omitempty"`
		ElementName                        string   `xml:"ElementName,omitempty"`
		IDER                               bool     `xml:"IDER,omitempty"`
		SOL                                bool     `xml:"SOL,omitempty"`
		BIOSReflash                        bool     `xml:"BIOSReflash,omitempty"`
		BIOSSetup                          bool     `xml:"BIOSSetup,omitempty"`
		BIOSPause                          bool     `xml:"BIOSPause,omitempty"`
		ForcePXEBoot                       bool     `xml:"ForcePXEBoot,omitempty"`
		ForceHardDriveBoot                 bool     `xml:"ForceHardDriveBoot,omitempty"`
		ForceHardDriveSafeModeBoot         bool     `xml:"ForceHardDriveSafeModeBoot,omitempty"`
		ForceDiagnosticBoot                bool     `xml:"ForceDiagnosticBoot,omitempty"`
		ForceCDorDVDBoot                   bool     `xml:"ForceCDorDVDBoot,omitempty"`
		VerbosityScreenBlank               bool     `xml:"VerbosityScreenBlank,omitempty"`
		PowerButtonLock                    bool     `xml:"PowerButtonLock,omitempty"`
		ResetButtonLock                    bool     `xml:"ResetButtonLock,omitempty"`
		KeyboardLock                       bool     `xml:"KeyboardLock,omitempty"`
		SleepButtonLock                    bool     `xml:"SleepButtonLock,omitempty"`
		UserPasswordBypass                 bool     `xml:"UserPasswordBypass,omitempty"`
		ForcedProgressEvents               bool     `xml:"ForcedProgressEvents,omitempty"`
		VerbosityVerbose                   bool     `xml:"VerbosityVerbose,omitempty"`
		VerbosityQuiet                     bool     `xml:"VerbosityQuiet,omitempty"`
		ConfigurationDataReset             bool     `xml:"ConfigurationDataReset,omitempty"`
		BIOSSecureBoot                     bool     `xml:"BIOSSecureBoot,omitempty"`
		SecureErase                        bool     `xml:"SecureErase,omitempty"`
		ForceWinREBoot                     bool     `xml:"ForceWinREBoot,omitempty"`
		ForceUEFILocalPBABoot              bool     `xml:"ForceUEFILocalPBABoot,omitempty"`
		ForceUEFIHTTPSBoot                 bool     `xml:"ForceUEFIHTTPSBoot,omitempty"`
		AMTSecureBootControl               bool     `xml:"AMTSecureBootControl,omitempty"`
		UEFIWiFiCoExistenceAndProfileShare bool     `xml:"UEFIWiFiCoExistenceAndProfileShare,omitempty"`
		PlatformErase                      int      `xml:"PlatformErase,omitempty"`
	}
)

type BootSettingDataRequest struct {
	XMLName                  xml.Name          `xml:"h:AMT_BootSettingData"`
	H                        string            `xml:"xmlns:h,attr"`
	InstanceID               string            `xml:"h:InstanceID"`
	ElementName              string            `xml:"h:ElementName"`
	OwningEntity             string            `xml:"h:OwningEntity"`
	UseSOL                   bool              `xml:"h:UseSOL"`
	UseSafeMode              bool              `xml:"h:UseSafeMode"`
	ReflashBIOS              bool              `xml:"h:ReflashBIOS"`
	BIOSSetup                bool              `xml:"h:BIOSSetup"`
	BIOSPause                bool              `xml:"h:BIOSPause"`
	LockPowerButton          bool              `xml:"h:LockPowerButton"`
	LockResetButton          bool              `xml:"h:LockResetButton"`
	LockKeyboard             bool              `xml:"h:LockKeyboard"`
	LockSleepButton          bool              `xml:"h:LockSleepButton"`
	UserPasswordBypass       bool              `xml:"h:UserPasswordBypass"`
	ForcedProgressEvents     bool              `xml:"h:ForcedProgressEvents"`
	FirmwareVerbosity        FirmwareVerbosity `xml:"h:FirmwareVerbosity"`
	ConfigurationDataReset   bool              `xml:"h:ConfigurationDataReset"`
	IDERBootDevice           IDERBootDevice    `xml:"h:IDERBootDevice"`
	UseIDER                  bool              `xml:"h:UseIDER"`
	EnforceSecureBoot        bool              `xml:"h:EnforceSecureBoot"`
	BootMediaIndex           int               `xml:"h:BootMediaIndex"`
	SecureErase              bool              `xml:"h:SecureErase"`
	RSEPassword              string            `xml:"h:RSEPassword"`
	OptionsCleared           bool              `xml:"h:OptionsCleared"`           //readonly
	WinREBootEnabled         bool              `xml:"h:WinREBootEnabled"`         //readonly
	UEFILocalPBABootEnabled  bool              `xml:"h:UEFILocalPBABootEnabled"`  //readonly
	UEFIHTTPSBootEnabled     bool              `xml:"h:UEFIHTTPSBootEnabled"`     //readonly
	SecureBootControlEnabled bool              `xml:"h:SecureBootControlEnabled"` //readonly
	BootguardStatus          bool              `xml:"h:BootguardStatus"`          //readonly
	BIOSLastStatus           []int             `xml:"h:BIOSLastStatus"`           //readonly
	UEFIBootParametersArray  []int             `xml:"h:UEFIBootParametersArray"`
	UEFIBootNumberOfParams   []int             `xml:"h:UEFIBootNumberOfParams"`
	RPEEnabled               bool              `xml:"h:RPEEnabled"`
	PlatformErase            bool              `xml:"h:PlatformErase"`
}

type FirmwareVerbosity int
type IDERBootDevice int
