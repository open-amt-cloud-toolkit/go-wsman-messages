/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kerberos

import "testing"

func TestReturnValue_String(t *testing.T) {
	tests := []struct {
		state    ReturnValue
		expected string
	}{
		{ReturnValueSuccess, "Success"},
		{ReturnValueNotSupported, "NotSupported"},
		{ReturnValueUnspecifiedError, "UnspecifiedError"},
		{ReturnValueUnsupportedParameter, "UnsupportedParameter"},
		{ReturnValueNoKerberosRealmName, "NoKerberosRealmName"},
		{ReturnValueNoServicePrincipalName, "NoServicePrincipalName"},
		{ReturnValueNoServicePrincipalProtocol, "NoServicePrincipalProtocol"},
		{ReturnValueNoKeyVersion, "NoKeyVersion"},
		{ReturnValueNoEncryptionAlgorithm, "NoEncryptionAlgorithm"},
		{ReturnValueNoMasterKey, "NoMasterKey"},
		{ReturnValueNoMaximumClockTolerance, "NoMaximumClockTolerance"},
		{ReturnValueNoKerberosEnabled, "NoKerberosEnabled"},
		{ReturnValueNoPassphrase, "NoPassphrase"},
		{ReturnValueNoSalt, "NoSalt"},
		{ReturnValueNoIterationCount, "NoIterationCount"},
		{ReturnValueNoSupportedEncryptionAlgorithms, "NoSupportedEncryptionAlgorithms"},
		{ReturnValueNoConfiguredEncryptionAlgorithms, "NoConfiguredEncryptionAlgorithms"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestServicePrincipalProtocol_String(t *testing.T) {
	tests := []struct {
		state    ServicePrincipalProtocol
		expected string
	}{
		{ServicePrincipalProtocolHTTPProtocolDefinition, "HTTPProtocolDefinition"},
		{ServicePrincipalProtocolHTTPSProtocolDefinition, "HTTPSProtocolDefinition"},
		{ServicePrincipalProtocolSOLAndIDERProtocolDefinition, "SOL&IDERProtocolDefinition"},
		{ServicePrincipalProtocolSOLAndIDERProtocolDefinitionUsingSSL, "SOL&IDERProtocolDefinitionUsingSSL"},
		{ServicePrincipalProtocol(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestSupportedEncryptionAlgorithms_String(t *testing.T) {
	tests := []struct {
		state    SupportedEncryptionAlgorithms
		expected string
	}{
		{SupportedEncryptionAlgorithmsRC4HMAC, "RC4-HMAC"},
		{SupportedEncryptionAlgorithmsAES128CTSHMACSHA196, "AES128-CTS-HMAC-SHA1-96"},
		{SupportedEncryptionAlgorithmsAES256CTSHMACSHA196, "AES256-CTS-HMAC-SHA1-96"},
		{SupportedEncryptionAlgorithms(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestConfiguredEncryptionAlgorithms_String(t *testing.T) {
	tests := []struct {
		state    ConfiguredEncryptionAlgorithms
		expected string
	}{
		{ConfiguredEncryptionAlgorithmsRC4HMAC, "RC4-HMAC"},
		{ConfiguredEncryptionAlgorithmsAES128CTSHMACSHA196, "AES128-CTS-HMAC-SHA1-96"},
		{ConfiguredEncryptionAlgorithmsAES256CTSHMACSHA196, "AES256-CTS-HMAC-SHA1-96"},
		{ConfiguredEncryptionAlgorithms(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
