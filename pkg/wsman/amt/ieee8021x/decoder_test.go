/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import "testing"

func TestAuthenticationProtocol_String(t *testing.T) {
	tests := []struct {
		state    AuthenticationProtocol
		expected string
	}{
		{AuthenticationProtocolTLS, "TLS"},
		{AuthenticationProtocolTTLSMSCHAPv2, "TTLSMSCHAPv2"},
		{AuthenticationProtocolPEAPMSCHAPv2, "PEAPMSCHAPv2"},
		{AuthenticationProtocolEAPGTC, "EAPGTC"},
		{AuthenticationProtocolEAPFASTMSCHAPv2, "EAPFASTMSCHAPv2"},
		{AuthenticationProtocolEAPFASTGTC, "EAPFASTGTC"},
		{AuthenticationProtocolEAPFASTTLS, "EAPFASTTLS"},
		{AuthenticationProtocol(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestServerCertificateNameComparison_String(t *testing.T) {
	tests := []struct {
		state    ServerCertificateNameComparison
		expected string
	}{
		{FullName, "FullName"},
		{DomainSuffix, "DomainSuffix"},
		{ServerCertificateNameComparison(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
