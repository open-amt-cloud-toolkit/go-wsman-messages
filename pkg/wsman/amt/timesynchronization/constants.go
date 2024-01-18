/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

const (
	AMT_TimeSynchronizationService string = "AMT_TimeSynchronizationService"
	GetLowAccuracyTimeSynch        string = "GetLowAccuracyTimeSynch"
	SetHighAccuracyTimeSynch       string = "SetHighAccuracyTimeSynch"
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
	LocalTimeSyncEnabledDefaultTrue LocalTimeSyncEnabled = iota
	LocalTimeSyncEnabledConfiguredTrue
	LocalTimeSyncEnabledFalse
)

const (
	TimeSourceBiosRTC TimeSource = iota
	TimeSourceConfigured
)
