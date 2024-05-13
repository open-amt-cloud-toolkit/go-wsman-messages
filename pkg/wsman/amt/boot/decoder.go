/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

// INPUTS Constants.
const (
	AMTBootSettingData  string = "AMT_BootSettingData"
	AMTBootCapabilities string = "AMT_BootCapabilities"
	ValueNotFound       string = "Value not found in map"
)

const (
	SystemDefault FirmwareVerbosity = iota
	QuietMinimal
	VerboseAll
	ScreenBlank
)

const (
	FloppyBoot IDERBootDevice = iota
	CDBoot
)

// FirmwareVerbosityToString is a map of FirmwareVerbosity values to their string representations.
var firmwareVerbosityToString = map[FirmwareVerbosity]string{
	SystemDefault: "SystemDefault",
	QuietMinimal:  "QuietMinimal",
	VerboseAll:    "VerboseAll",
	ScreenBlank:   "ScreenBlank",
}

// String returns the string representation of the FirmwareVerbosity value.
func (f FirmwareVerbosity) String() string {
	if value, exists := firmwareVerbosityToString[f]; exists {
		return value
	}

	return ValueNotFound
}

// iderBootDeviceToString is a map of IDERBootDevice values to their string representations.
var iderBootDeviceToString = map[IDERBootDevice]string{
	FloppyBoot: "FloppyBoot",
	CDBoot:     "CDBoot",
}

// String returns the string representation of the IDERBootDevice value.
func (i IDERBootDevice) String() string {
	if value, exists := iderBootDeviceToString[i]; exists {
		return value
	}

	return ValueNotFound
}
