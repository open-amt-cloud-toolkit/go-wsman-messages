/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package boot

const (
	AMT_BootSettingData  string = "AMT_BootSettingData"
	AMT_BootCapabilities string = "AMT_BootCapabilities"
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
