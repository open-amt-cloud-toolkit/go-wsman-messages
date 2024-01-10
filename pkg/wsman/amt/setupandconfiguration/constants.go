/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

const (
	AMT_SetupAndConfigurationService string = "AMT_SetupAndConfigurationService"
	CommitChanges                    string = "CommitChanges"
	Unprovision                      string = "Unprovision"
	SetMEBxPassword                  string = "SetMEBxPassword"
	GetUuid                          string = "GetUuid"
)

const (
	RequestedStateEnabled RequestedState = iota + 2
	RequestedStateDisabled
	RequestedStateShutDown
	RequestedStateNoChange
	RequestedStateOffline
	RequestedStateTest
	RequestedStateDeferred
	RequestedStateQuiesce
	RequestedStateReboot
	RequestedStateReset
	RequestedStateNotApplicable
	RequestedStateUnknown RequestedState = 0
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

const (
	AdminControlMode  ProvisioningModeValue  = 1
	ClientControlMode ProvisioningModeValue  = 4
	PreProvisioning   ProvisioningStateValue = 0
	InProvisioning    ProvisioningStateValue = 1
	PostProvisioning  ProvisioningStateValue = 2
)
