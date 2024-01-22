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
	TLS AuthenticationProtocol = iota
	TTLS_MSCHAPv2
	PEAP_MSCHAPv2
	EAP_GTC
	EAPFAST_MSCHAPv2
	EAPFAST_GTC
	EAPFAST_TLS
)

const (
	FullName ServerCertificateNameComparison = iota
	DomainSuffix
)
