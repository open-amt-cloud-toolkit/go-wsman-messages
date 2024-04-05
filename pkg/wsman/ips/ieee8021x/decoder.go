/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

const (
	IPS_IEEE8021xSettings      string = "IPS_IEEE8021xSettings"
	IPS_8021xCredentialContext string = "IPS_8021xCredentialContext"
	SetCertificates            string = "SetCertificates"
)

const (
	EnabledWithCertificates    Enabled = 2
	Disabled                   Enabled = 3
	EnabledWithoutCertificates Enabled = 6
)

// enabledToString is a map of Enabled value to string
var enabledToString = map[Enabled]string{
	EnabledWithCertificates:    "EnabledWithCertificates",
	Disabled:                   "Disabled",
	EnabledWithoutCertificates: "EnabledWithoutCertificates",
}

// String returns a human-readable string representation of the Enabled enumeration
func (e Enabled) String() string {
	if s, ok := enabledToString[e]; ok {
		return s
	}
	return "Value not found in map"
}

const (
	AuthenticationProtocolEAPTLS int = iota
	AuthenticationProtocolEAPTTLS_MSCHAPv2
	AuthenticationProtocolPEAPv0_EAPMSCHAPv2
	AuthenticationProtocolPEAPv1_EAPGTC
	AuthenticationProtocolEAPFAST_MSCHAPv2
	AuthenticationProtocolEAPFAST_GTC
	AuthenticationProtocolEAP_MD5
	AuthenticationProtocolEAP_PSK
	AuthenticationProtocolEAP_SIM
	AuthenticationProtocolEAP_AKA
	AuthenticationProtocolEAPFAST_TLS
)

const (
	ReturnValueSuccess ReturnValue = iota
	ReturnValueInternalError
)

// returnValueToString is a map of ReturnValue value to string
var returnValueToString = map[ReturnValue]string{
	ReturnValueSuccess:       "Success",
	ReturnValueInternalError: "InternalError",
}

// String returns a human-readable string representation of the ReturnValue enumeration
func (r ReturnValue) String() string {
	if s, ok := returnValueToString[r]; ok {
		return s
	}
	return "Value not found in map"
}
