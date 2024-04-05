/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

const (
	AMT_PublicKeyCertificate       string = "AMT_PublicKeyCertificate"
	AMT_PublicKeyManagementService string = "AMT_PublicKeyManagementService"
	GeneratePKCS10RequestEx        string = "GeneratePKCS10RequestEx"
	AddTrustedRootCertificate      string = "AddTrustedRootCertificate"
	AddCertificate                 string = "AddCertificate"
	GenerateKeyPair                string = "GenerateKeyPair"
	AddKey                         string = "AddKey"
)

const (
	RSA KeyAlgorithm = 0
)

const (
	KeyLength2048 KeyLength = 2048
)

const (
	SHA1RSA SigningAlgorithm = iota
	SHA256RSA
)

const (
	EnabledDefaultEnabled           EnabledDefault = 2
	EnabledDefaultDisabled          EnabledDefault = 3
	EnabledDefaultNotApplicable     EnabledDefault = 5
	EnabledDefaultEnabledButOffline EnabledDefault = 6
	EnabledDefaultNoDefault         EnabledDefault = 7
	EnabledDefaultQuiesce           EnabledDefault = 9
)

// enabledDefaultToString is a map of EnabledDefault values to their string representations
var enabledDefaultToString = map[EnabledDefault]string{
	EnabledDefaultEnabled:           "Enabled",
	EnabledDefaultDisabled:          "Disabled",
	EnabledDefaultNotApplicable:     "NotApplicable",
	EnabledDefaultEnabledButOffline: "EnabledButOffline",
	EnabledDefaultNoDefault:         "NoDefault",
	EnabledDefaultQuiesce:           "Quiesce",
}

// String returns the string representation of the EnabledDefault value
func (e EnabledDefault) String() string {
	if value, exists := enabledDefaultToString[e]; exists {
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
)

// enabledStateToString is a map of EnabledState values to their string representations
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
}

// String returns the string representation of the EnabledState value
func (e EnabledState) String() string {
	if value, exists := enabledStateToString[e]; exists {
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

// operationalStatusToString is a map of OperationalStatus values to their string representations
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
	RequestedStateUnknown       RequestedState = 0
	RequestedStateEnabled       RequestedState = 2
	RequestedStateDisabled      RequestedState = 3
	RequestedStateShutDown      RequestedState = 4
	RequestedStateNoChange      RequestedState = 5
	RequestedStateOffline       RequestedState = 6
	RequestedStateTest          RequestedState = 7
	RequestedStateDeferred      RequestedState = 8
	RequestedStateQuiesce       RequestedState = 9
	RequestedStateReboot        RequestedState = 10
	RequestedStateReset         RequestedState = 11
	RequestedStateNotApplicable RequestedState = 12
)

// requestedStateToString is a map of RequestedState values to their string representations
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
}

// String returns the string representation of the RequestedState value
func (r RequestedState) String() string {
	if value, exists := requestedStateToString[r]; exists {
		return value
	}

	return "Value not found in map"
}

const (
	ReturnValueSuccess                 ReturnValue = 0
	ReturnValueInternalError           ReturnValue = 1
	ReturnValueNotPermitted            ReturnValue = 16
	ReturnValueMaxLimitReached         ReturnValue = 23
	ReturnValueInvalidParameter        ReturnValue = 36
	ReturnValueFlashWriteLimitExceeded ReturnValue = 38
	ReturnValueDuplicate               ReturnValue = 2058
	ReturnValueInvalidKeyLength        ReturnValue = 2062
	ReturnValueInvalidCert             ReturnValue = 2063
	ReturnValueUnsupported             ReturnValue = 2066
	ReturnValueOperationInProgress     ReturnValue = 2082
)

// returnValueToString is a map of ReturnValue values to their string representations
var returnValueToString = map[ReturnValue]string{
	ReturnValueSuccess:                 "Success",
	ReturnValueInternalError:           "InternalError",
	ReturnValueNotPermitted:            "NotPermitted",
	ReturnValueMaxLimitReached:         "MaxLimitReached",
	ReturnValueInvalidParameter:        "InvalidParameter",
	ReturnValueFlashWriteLimitExceeded: "FlashWriteLimitExceeded",
	ReturnValueDuplicate:               "Duplicate",
	ReturnValueInvalidKeyLength:        "InvalidKeyLength",
	ReturnValueInvalidCert:             "InvalidCertificate",
	ReturnValueUnsupported:             "Unsupported",
	ReturnValueOperationInProgress:     "OperationInProgress",
}

// String returns the string representation of the ReturnValue value
func (p ReturnValue) String() string {
	if value, exists := returnValueToString[p]; exists {
		return value
	}

	return "Value not found in map"
}
