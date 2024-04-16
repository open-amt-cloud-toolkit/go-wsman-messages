/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

const (
	AMT_IEEE8021xCredentialContext string = "AMT_8021xCredentialContext"
	AMT_IEEE8021xProfile           string = "AMT_8021XProfile"
)

const (
	AuthenticationProtocolTLS AuthenticationProtocol = iota
	AuthenticationProtocolTTLSMSCHAPv2
	AuthenticationProtocolPEAPMSCHAPv2
	AuthenticationProtocolEAPGTC
	AuthenticationProtocolEAPFASTMSCHAPv2
	AuthenticationProtocolEAPFASTGTC
	AuthenticationProtocolEAPFASTTLS
)

// authenticationProtocolString is a map of AuthenticationProtocol to their string representation
var authenticationProtocolString = map[AuthenticationProtocol]string{
	AuthenticationProtocolTLS:             "TLS",
	AuthenticationProtocolTTLSMSCHAPv2:    "TTLSMSCHAPv2",
	AuthenticationProtocolPEAPMSCHAPv2:    "PEAPMSCHAPv2",
	AuthenticationProtocolEAPGTC:          "EAPGTC",
	AuthenticationProtocolEAPFASTMSCHAPv2: "EAPFASTMSCHAPv2",
	AuthenticationProtocolEAPFASTGTC:      "EAPFASTGTC",
	AuthenticationProtocolEAPFASTTLS:      "EAPFASTTLS",
}

// String returns the string representation of the AuthenticationProtocol value
func (a AuthenticationProtocol) String() string {
	if value, exists := authenticationProtocolString[a]; exists {
		return value
	}

	return "Value not found in map"
}

const (
	FullName ServerCertificateNameComparison = iota
	DomainSuffix
)

// serverCertificateNameComparisonString is a map of ServerCertificateNameComparison to their string representation
var serverCertificateNameComparisonString = map[ServerCertificateNameComparison]string{
	FullName:     "FullName",
	DomainSuffix: "DomainSuffix",
}

// String returns the string representation of the ServerCertificateNameComparison value
func (s ServerCertificateNameComparison) String() string {
	if value, exists := serverCertificateNameComparisonString[s]; exists {
		return value
	}

	return "Value not found in map"
}
