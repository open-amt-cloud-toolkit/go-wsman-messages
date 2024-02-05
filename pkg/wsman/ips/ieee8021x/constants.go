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
	EnabledWithCertificates    IEEE8021xSettingsEnabled = 2
	Disabled                   IEEE8021xSettingsEnabled = 3
	EnabledWithoutCertificates IEEE8021xSettingsEnabled = 6
)

const (
	AuthenticationProtocolEAPTLS AuthenticationProtocol = iota
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
