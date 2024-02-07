/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportconfiguration

const (
	AMT_WiFiPortConfigurationService string = "AMT_WiFiPortConfigurationService"
	AddWiFiSettings                  string = "AddWiFiSettings"
)

const (
	LocalSyncDisabled LocalProfileSynchronizationEnabled = 0
	UnrestrictedSync  LocalProfileSynchronizationEnabled = 3
)

const (
	RelaxedPolicy NoHostCsmeSoftwarePolicy = iota
	AggressivePolicy
	Reserved
)

const (
	HealthStateUnknown             HealthState = 0
	HealthStateOK                  HealthState = 5
	HealthStateDegradedWarning     HealthState = 10
	HealthStateMinorFailure        HealthState = 15
	HealthStateMajorFailure        HealthState = 20
	HealthStateCriticalFailure     HealthState = 25
	HealthStateNonRecoverableError HealthState = 30
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
