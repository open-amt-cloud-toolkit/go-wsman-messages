/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbasedsetup

const (
	IPS_HostBasedSetupService string = "IPS_HostBasedSetupService"
	Setup                     string = "Setup"
	AdminSetup                string = "AdminSetup"
	AddNextCertInChain        string = "AddNextCertInChain"
	UpgradeClientToAdmin      string = "UpgradeClientToAdmin"
)

const (
	AdminPassEncryptionTypeNone AdminPassEncryptionType = iota
	AdminPassEncryptionTypeOther
	AdminPassEncryptionTypeHTTPDigestMD5A1
)

const (
	SigningAlgorithmNone SigningAlgorithm = iota
	SigningAlgorithmOther
	SigningAlgorithmRSASHA2256
)

const (
	NotProvisioned CurrentControlMode = iota
	Client
	Admin
)

const (
	CertChainStatusNotStarted CertChainStatus = iota
	CertChainStatusChainInProgress
	CertChainStatusChainComplete
)

const (
	AllowedControlModesNotProvisioned AllowedControlModes = iota
	AllowedControlModesClient
	AllowedControlModesAdmin
)

const (
	ReturnValueSuccess ReturnValue = iota
	ReturnValueInternalError
	ReturnValueInvalidState
	ReturnValueInvalidParam
	ReturnValueMethodDisabled
	ReturnValueAuthFailed
	ReturnValueFlashWriteLimitExceeded
)
