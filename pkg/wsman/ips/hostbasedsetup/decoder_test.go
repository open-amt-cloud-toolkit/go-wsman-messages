/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbasedsetup

import "testing"

func TestAdminPassEncryptionType_String(t *testing.T) {
	tests := []struct {
		state    AdminPassEncryptionType
		expected string
	}{
		{AdminPassEncryptionTypeNone, "None"},
		{AdminPassEncryptionTypeOther, "Other"},
		{AdminPassEncryptionTypeHTTPDigestMD5A1, "HTTPDigestMD5A1"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestSigningAlgorithm_String(t *testing.T) {
	tests := []struct {
		state    SigningAlgorithm
		expected string
	}{
		{SigningAlgorithmNone, "None"},
		{SigningAlgorithmOther, "Other"},
		{SigningAlgorithmRSASHA2256, "RSASHA2256"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestCurrentControlMode_String(t *testing.T) {
	tests := []struct {
		state    CurrentControlMode
		expected string
	}{
		{NotProvisioned, "NotProvisioned"},
		{Client, "Client"},
		{Admin, "Admin"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestCertChainStatus_String(t *testing.T) {
	tests := []struct {
		state    CertChainStatus
		expected string
	}{
		{CertChainStatusNotStarted, "NotStarted"},
		{CertChainStatusChainInProgress, "ChainInProgress"},
		{CertChainStatusChainComplete, "ChainComplete"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestAllowedControlModes_String(t *testing.T) {
	tests := []struct {
		state    AllowedControlModes
		expected string
	}{
		{AllowedControlModesNotProvisioned, "NotProvisioned"},
		{AllowedControlModesClient, "Client"},
		{AllowedControlModesAdmin, "Admin"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestSetupReturnValue_String(t *testing.T) {
	tests := []struct {
		state    SetupReturnValue
		expected string
	}{
		{SetupReturnValueSuccess, "Success"},
		{SetupReturnValueInternalError, "InternalError"},
		{SetupReturnValueInvalidState, "InvalidState"},
		{SetupReturnValueInvalidParam, "InvalidParam"},
		{SetupReturnValueMethodDisabled, "MethodDisabled"},
		{SetupReturnValueAuthFailed, "AuthFailed"},
		{SetupReturnValueFlashWriteLimitExceeded, "FlashWriteLimitExceeded"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestAddNextCertInChainReturnValue_String(t *testing.T) {
	tests := []struct {
		state    AddNextCertInChainReturnValue
		expected string
	}{
		{AddNextCertInChainReturnValueSuccess, "Success"},
		{AddNextCertInChainReturnValueInvalidParam, "InvalidParam"},
		{AddNextCertInChainReturnValueInternalError, "InternalError"},
		{AddNextCertInChainReturnValueInvalidState, "InvalidState"},
		{AddNextCertInChainReturnValueCertVerifyFailed, "CertVerifyFailed"},
		{AddNextCertInChainReturnValueCertChainLengthExceeded, "CertChainLengthExceeded"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestAdminSetupReturnValue_String(t *testing.T) {
	tests := []struct {
		state    AdminSetupReturnValue
		expected string
	}{
		{AdminSetupReturnValueSuccess, "Success"},
		{AdminSetupReturnValueInternalError, "InternalError"},
		{AdminSetupReturnValueInvalidState, "InvalidState"},
		{AdminSetupReturnValueInvalidParam, "InvalidParam"},
		{AdminSetupReturnValueAuthFailed, "AuthFailed"},
		{AdminSetupReturnValueFlashWriteLimitExceeded, "FlashWriteLimitExceeded"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
