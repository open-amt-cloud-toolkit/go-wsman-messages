/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package redirection

const (
	AMT_RedirectionService string = "AMT_RedirectionService"
	RequestStateChange     string = "RequestStateChange"
)

const (
	Unknown EnabledState = iota
	Other
	Enabled
	Disabled
	ShuttingDown
	NotApplicable
	EnabledButOffline
	InTest
	Deferred
	Quiesce
	Starting
	DMTFReserved
	IDERAndSOLAreDisabled         = 32768
	IDERIsEnabledAndSOLIsDisabled = 32769
	SOLIsEnabledAndIDERIsDisabled = 32770
	IDERAndSOLAreEnabled          = 32771
)

const (
	DisableIDERAndSOL       RequestedState = 32768
	EnableIDERAndDisableSOL RequestedState = 32769
	EnableSOLAndDisableIDER RequestedState = 32770
	EnableIDERAndSOL        RequestedState = 32771
)
