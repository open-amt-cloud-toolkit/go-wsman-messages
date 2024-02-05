/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

const (
	CIM_BootConfigSetting string = "CIM_BootConfigSetting"
	CIM_BootSourceSetting string = "CIM_BootSourceSetting"
	CIM_BootService       string = "CIM_BootService"
	ChangeBootOrder       string = "ChangeBootOrder"
)

const (
	HardDrive             Source = "CIM:Hard-Disk:1"
	CD                    Source = "CIM:CD/DVD:1"
	PXE                   Source = "CIM:Network:1"
	OCR_UEFI_HTTPS        Source = "Intel(r)AMT:OCR-UEFI-Boot-Option-HTTPS:1"
	OCR_UEFI_BootOption1  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:1"
	OCR_UEFI_BootOption2  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:2"
	OCR_UEFI_BootOption3  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:3"
	OCR_UEFI_BootOption4  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:4"
	OCR_UEFI_BootOption5  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:5"
	OCR_UEFI_BootOption6  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:6"
	OCR_UEFI_BootOption7  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:7"
	OCR_UEFI_BootOption8  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:8"
	OCR_UEFI_BootOption9  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:9"
	OCR_UEFI_BootOption10 Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:10"
)

const (
	FailThroughSupportedUnknown FailThroughSupported = iota
	IsSupported
	NotSupported
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
