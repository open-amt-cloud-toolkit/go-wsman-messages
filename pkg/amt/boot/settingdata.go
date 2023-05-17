/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

const AMT_BootSettingData = "AMT_BootSettingData"

type BootSettingData struct {
	models.BootSettingData
	InstanceID               string
	ElementName              string
	UseSOL                   bool
	UseSafeMode              bool
	ReflashBIOS              bool
	BIOSSetup                bool
	BIOSPause                bool
	LockPowerButton          bool
	LockResetButton          bool
	LockKeyboard             bool
	LockSleepButton          bool
	UserPasswordBypass       bool
	ForcedProgressEvents     bool
	FirmwareVerbosity        FirmwareVerbosity
	ConfigurationDataReset   bool
	IDERBootDevice           IDERBootDevice
	UseIDER                  bool
	EnforceSecureBoot        bool
	BootMediaIndex           int
	SecureErase              bool
	RSEPassword              string
	WinREBootEnabled         bool  //readonly
	UEFILocalPBABootEnabled  bool  //readonly
	UEFIHTTPSBootEnabled     bool  //readonly
	SecureBootControlEnabled bool  //readonly
	BootguardStatus          bool  //readonly
	OptionsCleared           bool  //readonly
	BIOSLastStatus           []int //readonly
	UEFIBootParametersArray  []int
	UEFIBootNumberOfParams   []int
	RPEEnabled               bool
	PlatformErase            bool
}

type FirmwareVerbosity uint8

const (
	SystemDefault FirmwareVerbosity = iota
	QuietMinimal
	VerboseAll
	ScreenBlank
)

type IDERBootDevice uint8

const (
	FloppyBoot IDERBootDevice = iota
	CDBoot
)

type BootSettingDataResponse struct {
	AMT_BootSettingData BootSettingData
}

type SettingData struct {
	base wsman.Base
}

func NewBootSettingData(wsmanMessageCreator *wsman.WSManMessageCreator) SettingData {
	return SettingData{
		base: wsman.NewBase(wsmanMessageCreator, AMT_BootSettingData),
	}
}

// Get retrieves the representation of the instance
func (BootSettingData SettingData) Get() string {
	return BootSettingData.base.Get(nil)
}

// Enumerates the instances of this class
func (BootSettingData SettingData) Enumerate() string {
	return BootSettingData.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (BootSettingData SettingData) Pull(enumerationContext string) string {
	return BootSettingData.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (BootSettingData SettingData) Put(bootSettingData BootSettingData) string {
	return BootSettingData.base.Put(bootSettingData, false, nil)
}
