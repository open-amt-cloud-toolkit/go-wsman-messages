/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

const (
	AMT_PublicKeyCertificate       string = "AMT_PublicKeyCertificate"
	AMT_PublicKeyManagementService string = "AMT_PublicKeyManagementService"
	GeneratePKCS10RequestEx        string = "GeneratePKCS10RequestEx"
	AddTrustedRootCertificate      string = "AddTrustedRootCertificate"
	AddCertificate                 string = "AddCertificate"
	GenerateKeyPair                string = "GenerateKeyPair"
	AddKey                         string = "AddKey"
)

const (
	RSA KeyAlgorithm = 0
)

const (
	KeyLength2048 KeyLength = 2048
)

const (
	SHA1RSA SigningAlgorithm = iota
	SHA256RSA
)

const (
	EnabledDefaultEnabled           EnabledDefault = 2
	EnabledDefaultDisabled          EnabledDefault = 3
	EnabledDefaultNotApplicable     EnabledDefault = 5
	EnabledDefaultEnabledbutOffline EnabledDefault = 6
	EnabledDefaultNoDefault         EnabledDefault = 7
	EnabledDefaultQuiesce           EnabledDefault = 9
)

const (
	RequestedStateUnknown       RequestedState = 0
	RequestedStateEnabled       RequestedState = 2
	RequestedStateDisabled      RequestedState = 3
	RequestedStateShutDown      RequestedState = 4
	RequestedStateNoChange      RequestedState = 5
	RequestedStateOffline       RequestedState = 6
	RequestedStateTest          RequestedState = 7
	RequestedStateDeferred      RequestedState = 8
	RequestedStateQuiesce       RequestedState = 9
	RequestedStateReboot        RequestedState = 10
	RequestedStateReset         RequestedState = 11
	RequestedStateNotApplicable RequestedState = 12
)

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledbutOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)
