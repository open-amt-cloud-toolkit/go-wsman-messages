/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbasedsetup

const (
	IPSHostBasedSetupService string = "IPS_HostBasedSetupService"
	Setup                    string = "Setup"
	AdminSetup               string = "AdminSetup"
	AddNextCertInChain       string = "AddNextCertInChain"
	UpgradeClientToAdmin     string = "UpgradeClientToAdmin"
	ValueNotFound            string = "Value not found in map"
)

const (
	AdminPassEncryptionTypeNone AdminPassEncryptionType = iota
	AdminPassEncryptionTypeOther
	AdminPassEncryptionTypeHTTPDigestMD5A1
)

// adminPassEncryptionTypeToString is a map of AdminPassEncryptionType value to string.
var adminPassEncryptionTypeToString = map[AdminPassEncryptionType]string{
	AdminPassEncryptionTypeNone:            "None",
	AdminPassEncryptionTypeOther:           "Other",
	AdminPassEncryptionTypeHTTPDigestMD5A1: "HTTPDigestMD5A1",
}

// String returns a human-readable string representation of the AdminPassEncryptionType enumeration.
func (a AdminPassEncryptionType) String() string {
	if s, ok := adminPassEncryptionTypeToString[a]; ok {
		return s
	}

	return ValueNotFound
}

const (
	SigningAlgorithmNone SigningAlgorithm = iota
	SigningAlgorithmOther
	SigningAlgorithmRSASHA2256
)

// signingAlgorithmToString is a map of SigningAlgorithm value to string.
var signingAlgorithmToString = map[SigningAlgorithm]string{
	SigningAlgorithmNone:       "None",
	SigningAlgorithmOther:      "Other",
	SigningAlgorithmRSASHA2256: "RSASHA2256",
}

// String returns a human-readable string representation of the SigningAlgorithm enumeration.
func (s SigningAlgorithm) String() string {
	if s, ok := signingAlgorithmToString[s]; ok {
		return s
	}

	return ValueNotFound
}

const (
	NotProvisioned CurrentControlMode = iota
	Client
	Admin
)

// currentControlModeToString is a map of CurrentControlMode value to string.
var currentControlModeToString = map[CurrentControlMode]string{
	NotProvisioned: "NotProvisioned",
	Client:         "Client",
	Admin:          "Admin",
}

// String returns a human-readable string representation of the CurrentControlMode enumeration.
func (c CurrentControlMode) String() string {
	if s, ok := currentControlModeToString[c]; ok {
		return s
	}

	return ValueNotFound
}

const (
	CertChainStatusNotStarted CertChainStatus = iota
	CertChainStatusChainInProgress
	CertChainStatusChainComplete
)

// certChainStatusToString is a map of CertChainStatus value to string.
var certChainStatusToString = map[CertChainStatus]string{
	CertChainStatusNotStarted:      "NotStarted",
	CertChainStatusChainInProgress: "ChainInProgress",
	CertChainStatusChainComplete:   "ChainComplete",
}

// String returns a human-readable string representation of the CertChainStatus enumeration.
func (c CertChainStatus) String() string {
	if s, ok := certChainStatusToString[c]; ok {
		return s
	}

	return ValueNotFound
}

const (
	AllowedControlModesNotProvisioned AllowedControlModes = iota
	AllowedControlModesClient
	AllowedControlModesAdmin
)

// allowedControlModesToString is a map of AllowedControlModes value to string.
var allowedControlModesToString = map[AllowedControlModes]string{
	AllowedControlModesNotProvisioned: "NotProvisioned",
	AllowedControlModesClient:         "Client",
	AllowedControlModesAdmin:          "Admin",
}

// String returns a human-readable string representation of the AllowedControlModes enumeration.
func (a AllowedControlModes) String() string {
	if s, ok := allowedControlModesToString[a]; ok {
		return s
	}

	return ValueNotFound
}

const (
	SetupReturnValueSuccess SetupReturnValue = iota
	SetupReturnValueInternalError
	SetupReturnValueInvalidState
	SetupReturnValueInvalidParam
	SetupReturnValueMethodDisabled
	SetupReturnValueAuthFailed
	SetupReturnValueFlashWriteLimitExceeded
)

// setupReturnValueToString is a map of ReturnValue value to string.
var setupReturnValueToString = map[SetupReturnValue]string{
	SetupReturnValueSuccess:                 "Success",
	SetupReturnValueInternalError:           "InternalError",
	SetupReturnValueInvalidState:            "InvalidState",
	SetupReturnValueInvalidParam:            "InvalidParam",
	SetupReturnValueMethodDisabled:          "MethodDisabled",
	SetupReturnValueAuthFailed:              "AuthFailed",
	SetupReturnValueFlashWriteLimitExceeded: "FlashWriteLimitExceeded",
}

// String returns a human-readable string representation of the ReturnValue enumeration.
func (s SetupReturnValue) String() string {
	if s, ok := setupReturnValueToString[s]; ok {
		return s
	}

	return ValueNotFound
}

const (
	AddNextCertInChainReturnValueSuccess AddNextCertInChainReturnValue = iota
	AddNextCertInChainReturnValueInvalidParam
	AddNextCertInChainReturnValueInternalError
	AddNextCertInChainReturnValueInvalidState
	AddNextCertInChainReturnValueCertVerifyFailed
	AddNextCertInChainReturnValueCertChainLengthExceeded
)

// addNextCertInChainReturnValueToString is a map of ReturnValue value to string.
var addNextCertInChainReturnValueToString = map[AddNextCertInChainReturnValue]string{
	AddNextCertInChainReturnValueSuccess:                 "Success",
	AddNextCertInChainReturnValueInvalidParam:            "InvalidParam",
	AddNextCertInChainReturnValueInternalError:           "InternalError",
	AddNextCertInChainReturnValueInvalidState:            "InvalidState",
	AddNextCertInChainReturnValueCertVerifyFailed:        "CertVerifyFailed",
	AddNextCertInChainReturnValueCertChainLengthExceeded: "CertChainLengthExceeded",
}

// String returns a human-readable string representation of the ReturnValue enumeration.
func (a AddNextCertInChainReturnValue) String() string {
	if s, ok := addNextCertInChainReturnValueToString[a]; ok {
		return s
	}

	return ValueNotFound
}

const (
	AdminSetupReturnValueSuccess AdminSetupReturnValue = iota
	AdminSetupReturnValueInternalError
	AdminSetupReturnValueInvalidState
	AdminSetupReturnValueInvalidParam
	_
	AdminSetupReturnValueAuthFailed
	AdminSetupReturnValueFlashWriteLimitExceeded
)

// adminSetupReturnValueToString is a map of ReturnValue value to string.
var adminSetupReturnValueToString = map[AdminSetupReturnValue]string{
	AdminSetupReturnValueSuccess:                 "Success",
	AdminSetupReturnValueInternalError:           "InternalError",
	AdminSetupReturnValueInvalidState:            "InvalidState",
	AdminSetupReturnValueInvalidParam:            "InvalidParam",
	AdminSetupReturnValueAuthFailed:              "AuthFailed",
	AdminSetupReturnValueFlashWriteLimitExceeded: "FlashWriteLimitExceeded",
}

// String returns a human-readable string representation of the ReturnValue enumeration.
func (a AdminSetupReturnValue) String() string {
	if s, ok := adminSetupReturnValueToString[a]; ok {
		return s
	}

	return ValueNotFound
}
