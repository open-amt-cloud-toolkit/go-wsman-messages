/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

const (
	AMT_UserInitiatedConnectionService string = "AMT_UserInitiatedConnectionService"
)

const (
	AllInterfacesDisabled      RequestedState = 32768
	BIOSInterfaceEnabled       RequestedState = 32769
	OSInterfaceEnabled         RequestedState = 32770
	BIOSandOSInterfacesEnabled RequestedState = 32771
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
	EnabledStateAllInterfacesdisabled      EnabledState = 32768
	EnabledStateBIOSInterfaceenabled       EnabledState = 32769
	EnabledStateOSInterfaceenabled         EnabledState = 32770
	EnabledStateBIOSandOSInterfacesenabled EnabledState = 32771
)
