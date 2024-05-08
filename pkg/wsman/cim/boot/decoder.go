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
	SetBootConfigRole     string = "SetBootConfigRole"
)

const (
	HardDrive             Source = "Intel(r) AMT: Force Hard-drive Boot"
	CD                    Source = "Intel(r) AMT: Force CD/DVD Boot"
	PXE                   Source = "Intel(r) AMT: Force PXE Boot"
	OCR_UEFI_HTTPS        Source = "Intel(r) AMT: Force OCR UEFI HTTPS Boot"
	OCR_UEFI_BootOption1  Source = "Intel(r) AMT: Force OCR UEFI Boot Option 1"
	OCR_UEFI_BootOption2  Source = "Intel(r) AMT: Force OCR UEFI Boot Option 2"
	OCR_UEFI_BootOption3  Source = "Intel(r) AMT: Force OCR UEFI Boot Option 3"
	OCR_UEFI_BootOption4  Source = "Intel(r) AMT: Force OCR UEFI Boot Option 4"
	OCR_UEFI_BootOption5  Source = "Intel(r) AMT: Force OCR UEFI Boot Option 5"
	OCR_UEFI_BootOption6  Source = "Intel(r) AMT: Force OCR UEFI Boot Option 6"
	OCR_UEFI_BootOption7  Source = "Intel(r) AMT: Force OCR UEFI Boot Option 7"
	OCR_UEFI_BootOption8  Source = "Intel(r) AMT: Force OCR UEFI Boot Option 8"
	OCR_UEFI_BootOption9  Source = "Intel(r) AMT: Force OCR UEFI Boot Option 9"
	OCR_UEFI_BootOption10 Source = "Intel(r) AMT: Force OCR UEFI Boot Option 10"
)

const (
	FailThroughSupportedUnknown FailThroughSupported = iota
	FailThroughSupportedIsSupported
	FailThroughSupportedNotSupported
)

// FailThroughSupportedToString is a mapping of the FailThroughSupported value to a string
var FailThroughSupportedToString = map[FailThroughSupported]string{
	FailThroughSupportedUnknown:      "Unknown",
	FailThroughSupportedIsSupported:  "IsSupported",
	FailThroughSupportedNotSupported: "NotSupported",
}

// String returns the string representation of the FailThroughSupported value
func (f FailThroughSupported) String() string {
	if value, exists := FailThroughSupportedToString[f]; exists {
		return value
	}

	return "Value not found in map"
}

const (
	ReturnValueCompletedNoError ReturnValue = iota
	ReturnValueNotSupported
	ReturnValueUnknownError
	ReturnValueBusy
	ReturnValueInvalidReference
	ReturnValueInvalidParameter
	ReturnValueAccessDenied
)

// ReturnValueToString is a mapping of the ReturnValue value to a string
var ReturnValueToString = map[ReturnValue]string{
	ReturnValueCompletedNoError: "CompletedNoError",
	ReturnValueNotSupported:     "NotSupported",
	ReturnValueUnknownError:     "UnknownError",
	ReturnValueBusy:             "Busy",
	ReturnValueInvalidReference: "InvalidReference",
	ReturnValueInvalidParameter: "InvalidParameter",
	ReturnValueAccessDenied:     "AccessDenied",
}

// String returns the string representation of the ReturnValue value
func (r ReturnValue) String() string {
	if value, exists := ReturnValueToString[r]; exists {
		return value
	}

	return "Value not found in map"
}

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
	OperationalStatusSupportingEntityInError
	OperationalStatusCompleted
	OperationalStatusPowerMode
	OperationalStatusRelocating
)

// operationalStatusToString is a mapping of the OperationalStatus value to a string
var operationalStatusToString = map[OperationalStatus]string{
	OperationalStatusUnknown:                 "Unknown",
	OperationalStatusOther:                   "Other",
	OperationalStatusOK:                      "OK",
	OperationalStatusDegraded:                "Degraded",
	OperationalStatusStressed:                "Stressed",
	OperationalStatusPredictiveFailure:       "PredictiveFailure",
	OperationalStatusError:                   "Error",
	OperationalStatusNonRecoverableError:     "NonRecoverableError",
	OperationalStatusStarting:                "Starting",
	OperationalStatusStopping:                "Stopping",
	OperationalStatusStopped:                 "Stopped",
	OperationalStatusInService:               "InService",
	OperationalStatusNoContact:               "NoContact",
	OperationalStatusLostCommunication:       "LostCommunication",
	OperationalStatusAborted:                 "Aborted",
	OperationalStatusDormant:                 "Dormant",
	OperationalStatusSupportingEntityInError: "SupportingEntityInError",
	OperationalStatusCompleted:               "Completed",
	OperationalStatusPowerMode:               "PowerMode",
	OperationalStatusRelocating:              "Relocating",
}

// String returns the string representation of the OperationalStatus value
func (o OperationalStatus) String() string {
	if value, exists := operationalStatusToString[o]; exists {
		return value
	}

	return "Value not found in map"
}

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
	EnabledStateIntelOneClickRecoveryAndIntelRPEAreDisabledAndAllOtherBootOptionsAreEnabled         EnabledState = 32768
	EnabledStateIntelOneClickRecoveryIsEnabledAndIntelRPEIsDisabledAndAllOtherBootOptionsAreEnabled EnabledState = 32769
	EnabledStateIntelRPEIsEnabledAndIntelOneClickRecoveryIsDisabledAndAllOtherBootOptionsAreEnabled EnabledState = 32770
	EnabledStateIntelOneClickRecoveryAndIntelRPEAreEnabledAndAllOtherBootOptionsAreEnabled          EnabledState = 32771
)

// enabledStateToString is a mapping of the EnabledState value to a string
var enabledStateToString = map[EnabledState]string{
	EnabledStateUnknown:           "Unknown",
	EnabledStateOther:             "Other",
	EnabledStateEnabled:           "Enabled",
	EnabledStateDisabled:          "Disabled",
	EnabledStateShuttingDown:      "ShuttingDown",
	EnabledStateNotApplicable:     "NotApplicable",
	EnabledStateEnabledButOffline: "EnabledButOffline",
	EnabledStateInTest:            "InTest",
	EnabledStateDeferred:          "Deferred",
	EnabledStateQuiesce:           "Quiesce",
	EnabledStateStarting:          "Starting",
	EnabledStateIntelOneClickRecoveryAndIntelRPEAreDisabledAndAllOtherBootOptionsAreEnabled:         "IntelOneClickRecoveryAndIntelRPEAreDisabledAndAllOtherBootOptionsAreEnabled",
	EnabledStateIntelOneClickRecoveryIsEnabledAndIntelRPEIsDisabledAndAllOtherBootOptionsAreEnabled: "IntelOneClickRecoveryIsEnabledAndIntelRPEIsDisabledAndAllOtherBootOptionsAreEnabled",
	EnabledStateIntelRPEIsEnabledAndIntelOneClickRecoveryIsDisabledAndAllOtherBootOptionsAreEnabled: "IntelRPEIsEnabledAndIntelOneClickRecoveryIsDisabledAndAllOtherBootOptionsAreEnabled",
	EnabledStateIntelOneClickRecoveryAndIntelRPEAreEnabledAndAllOtherBootOptionsAreEnabled:          "IntelOneClickRecoveryAndIntelRPEAreEnabledAndAllOtherBootOptionsAreEnabled",
}

// String returns the string representation of the EnabledState value
func (e EnabledState) String() string {
	if value, exists := enabledStateToString[e]; exists {
		return value
	}

	return "Value not found in map"
}

const (
	RequestedStateUnknown                                                             RequestedState = 0
	RequestedStateEnabled                                                             RequestedState = 2
	RequestedStateDisabled                                                            RequestedState = 3
	RequestedStateShutDown                                                            RequestedState = 4
	RequestedStateNoChange                                                            RequestedState = 5
	RequestedStateOffline                                                             RequestedState = 6
	RequestedStateTest                                                                RequestedState = 7
	RequestedStateDeferred                                                            RequestedState = 8
	RequestedStateQuiesce                                                             RequestedState = 9
	RequestedStateReboot                                                              RequestedState = 10
	RequestedStateReset                                                               RequestedState = 11
	RequestedStateNotApplicable                                                       RequestedState = 12
	RequestedStateDisableIntelOneClickRecoveryAndIntelRPEAndEnableAllOtherBootOptions RequestedState = 32768
	RequestedStateDisableIntelRPEAndEnableIntelOneClickRecoveryAndAllOtherBootOptions RequestedState = 32769
	RequestedStateDisableIntelOneClickRecoveryAndEnableIntelRPEAndAllOtherBootOptions RequestedState = 32770
	RequestedStateEnableAllBootOptions                                                RequestedState = 32771
)

// requestedStateToString is a mapping of the RequestedState value to a string
var requestedStateToString = map[RequestedState]string{
	RequestedStateUnknown:       "Unknown",
	RequestedStateEnabled:       "Enabled",
	RequestedStateDisabled:      "Disabled",
	RequestedStateShutDown:      "ShutDown",
	RequestedStateNoChange:      "NoChange",
	RequestedStateOffline:       "Offline",
	RequestedStateTest:          "Test",
	RequestedStateDeferred:      "Deferred",
	RequestedStateQuiesce:       "Quiesce",
	RequestedStateReboot:        "Reboot",
	RequestedStateReset:         "Reset",
	RequestedStateNotApplicable: "NotApplicable",
	RequestedStateDisableIntelOneClickRecoveryAndIntelRPEAndEnableAllOtherBootOptions: "DisableIntelOneClickRecoveryAndIntelRPEAndEnableAllOtherBootOptions",
	RequestedStateDisableIntelRPEAndEnableIntelOneClickRecoveryAndAllOtherBootOptions: "DisableIntelRPEAndEnableIntelOneClickRecoveryAndAllOtherBootOptions",
	RequestedStateDisableIntelOneClickRecoveryAndEnableIntelRPEAndAllOtherBootOptions: "DisableIntelOneClickRecoveryAndEnableIntelRPEAndAllOtherBootOptions",
	RequestedStateEnableAllBootOptions:                                                "EnableAllBootOptions",
}

// String returns the string representation of the RequestedState value
func (r RequestedState) String() string {
	if value, exists := requestedStateToString[r]; exists {
		return value
	}

	return "Value not found in map"
}
