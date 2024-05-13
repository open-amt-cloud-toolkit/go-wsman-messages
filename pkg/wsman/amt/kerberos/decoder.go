/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kerberos

const (
	AMTKerberosSettingData  string = "AMT_KerberosSettingData"
	SetCredentialCacheState string = "SetCredentialCacheState"
	GetCredentialCacheState string = "GetCredentialCacheState"
	ValueNotFound           string = "Value not found in map"
)

const (
	EncryptionAlgorithmRC4EncryptionAndHMACAuthentication EncryptionAlgorithm = iota
)

const (
	ReturnValueSuccess ReturnValue = iota
	ReturnValueNotSupported
	ReturnValueUnspecifiedError
	ReturnValueUnsupportedParameter
	ReturnValueNoKerberosRealmName
	ReturnValueNoServicePrincipalName
	ReturnValueNoServicePrincipalProtocol
	ReturnValueNoKeyVersion
	ReturnValueNoEncryptionAlgorithm
	ReturnValueNoMasterKey
	ReturnValueNoMaximumClockTolerance
	ReturnValueNoKerberosEnabled
	ReturnValueNoPassphrase
	ReturnValueNoSalt
	ReturnValueNoIterationCount
	ReturnValueNoSupportedEncryptionAlgorithms
	ReturnValueNoConfiguredEncryptionAlgorithms
)

// returnValueString is a map of ReturnValue to string. These values are estimated return values from the AMT device.
var returnValueString = map[ReturnValue]string{
	ReturnValueSuccess:                          "Success",
	ReturnValueNotSupported:                     "NotSupported",
	ReturnValueUnspecifiedError:                 "UnspecifiedError",
	ReturnValueUnsupportedParameter:             "UnsupportedParameter",
	ReturnValueNoKerberosRealmName:              "NoKerberosRealmName",
	ReturnValueNoServicePrincipalName:           "NoServicePrincipalName",
	ReturnValueNoServicePrincipalProtocol:       "NoServicePrincipalProtocol",
	ReturnValueNoKeyVersion:                     "NoKeyVersion",
	ReturnValueNoEncryptionAlgorithm:            "NoEncryptionAlgorithm",
	ReturnValueNoMasterKey:                      "NoMasterKey",
	ReturnValueNoMaximumClockTolerance:          "NoMaximumClockTolerance",
	ReturnValueNoKerberosEnabled:                "NoKerberosEnabled",
	ReturnValueNoPassphrase:                     "NoPassphrase",
	ReturnValueNoSalt:                           "NoSalt",
	ReturnValueNoIterationCount:                 "NoIterationCount",
	ReturnValueNoSupportedEncryptionAlgorithms:  "NoSupportedEncryptionAlgorithms",
	ReturnValueNoConfiguredEncryptionAlgorithms: "NoConfiguredEncryptionAlgorithms",
}

// ConvertReturnValueToString returns the string representation of the ReturnValue value.
func (r ReturnValue) String() string {
	if value, exists := returnValueString[r]; exists {
		return value
	}

	return ValueNotFound
}

const (
	ServicePrincipalProtocolHTTPProtocolDefinition ServicePrincipalProtocol = iota
	ServicePrincipalProtocolHTTPSProtocolDefinition
	ServicePrincipalProtocolSOLAndIDERProtocolDefinition
	ServicePrincipalProtocolSOLAndIDERProtocolDefinitionUsingSSL
)

// servicePrincipalProtocolToString is a map of ServicePrincipalProtocol to string.
var servicePrincipalProtocolToString = map[ServicePrincipalProtocol]string{
	ServicePrincipalProtocolHTTPProtocolDefinition:               "HTTPProtocolDefinition",
	ServicePrincipalProtocolHTTPSProtocolDefinition:              "HTTPSProtocolDefinition",
	ServicePrincipalProtocolSOLAndIDERProtocolDefinition:         "SOL&IDERProtocolDefinition",
	ServicePrincipalProtocolSOLAndIDERProtocolDefinitionUsingSSL: "SOL&IDERProtocolDefinitionUsingSSL",
}

// String returns the string representation of the ServicePrincipalProtocol value.
func (s ServicePrincipalProtocol) String() string {
	if value, exists := servicePrincipalProtocolToString[s]; exists {
		return value
	}

	return ValueNotFound
}

const (
	SupportedEncryptionAlgorithmsRC4HMAC SupportedEncryptionAlgorithms = iota
	SupportedEncryptionAlgorithmsAES128CTSHMACSHA196
	SupportedEncryptionAlgorithmsAES256CTSHMACSHA196
)

// supportedEncryptionAlgorithmsToString is a map of SupportedEncryptionAlgorithms to string.
var supportedEncryptionAlgorithmsToString = map[SupportedEncryptionAlgorithms]string{
	SupportedEncryptionAlgorithmsRC4HMAC:             "RC4-HMAC",
	SupportedEncryptionAlgorithmsAES128CTSHMACSHA196: "AES128-CTS-HMAC-SHA1-96",
	SupportedEncryptionAlgorithmsAES256CTSHMACSHA196: "AES256-CTS-HMAC-SHA1-96",
}

// String returns the string representation of the SupportedEncryptionAlgorithms value.
func (s SupportedEncryptionAlgorithms) String() string {
	if value, exists := supportedEncryptionAlgorithmsToString[s]; exists {
		return value
	}

	return ValueNotFound
}

const (
	ConfiguredEncryptionAlgorithmsRC4HMAC ConfiguredEncryptionAlgorithms = iota
	ConfiguredEncryptionAlgorithmsAES128CTSHMACSHA196
	ConfiguredEncryptionAlgorithmsAES256CTSHMACSHA196
)

// configuredEncryptionAlgorithmsToString is a map of ConfiguredEncryptionAlgorithms to string.
var configuredEncryptionAlgorithmsToString = map[ConfiguredEncryptionAlgorithms]string{
	ConfiguredEncryptionAlgorithmsRC4HMAC:             "RC4-HMAC",
	ConfiguredEncryptionAlgorithmsAES128CTSHMACSHA196: "AES128-CTS-HMAC-SHA1-96",
	ConfiguredEncryptionAlgorithmsAES256CTSHMACSHA196: "AES256-CTS-HMAC-SHA1-96",
}

// String returns the string representation of the ConfiguredEncryptionAlgorithms value.
func (c ConfiguredEncryptionAlgorithms) String() string {
	if value, exists := configuredEncryptionAlgorithmsToString[c]; exists {
		return value
	}

	return ValueNotFound
}
