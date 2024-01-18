/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package kerberos

const (
	AMT_KerberosSettingData string = "AMT_KerberosSettingData"
	SetCredentialCacheState string = "SetCredentialCacheState"
	GetCredentialCacheState string = "GetCredentialCacheState"
)

const (
	ServicePrincipalProtocolHTTPProtocoldefinition ServicePrincipalProtocol = iota
	ServicePrincipalProtocolHTTPSProtocoldefinition
	ServicePrincipalProtocolSOLAndIDERprotocoldefinition
	ServicePrincipalProtocolSOLAndIDERprotocoldefinitionUsingSSL
)

const (
	EncryptionAlgorithmRC4EncryptionAndHMACAuthentication EncryptionAlgorithm = iota
)

const (
	SupportedEncryptionAlgorithmsRC4HMAC SupportedEncryptionAlgorithms = iota
	SupportedEncryptionAlgorithmsAES128CTSHMACSHA196
	SupportedEncryptionAlgorithmsAES256CTSHMACSHA196
)

const (
	ConfiguredEncryptionAlgorithmsRC4HMAC SupportedEncryptionAlgorithms = iota
	ConfiguredEncryptionAlgorithmsAES128CTSHMACSHA196
	ConfiguredEncryptionAlgorithmsAES256CTSHMACSHA196
)
