/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mediaaccess

const CIM_MediaAccessDevice string = "CIM_MediaAccessDevice"

const (
	CapabilitiesValuesUnknown CapabilitiesValues = iota
	CapabilitiesValuesOther
	CapabilitiesValuesSequentialAccess
	CapabilitiesValuesRandomAccess
	CapabilitiesValuesSupportsWriting
	CapabilitiesValuesEncryption
	CapabilitiesValuesCompression
	CapabilitiesValuesSupportsRemoveableMedia
	CapabilitiesValuesManualCleaning
	CapabilitiesValuesAutomaticCleaning
	CapabilitiesValuesSMARTNotification
	CapabilitiesValuesSupportsDualSidedMedia
	CapabilitiesValuesPredismountEjectNotRequired
)

const (
	EnabledDefaultEnabled           EnabledDefault = 2
	EnabledDefaultDisabled          EnabledDefault = 3
	EnabledDefaultNotApplicable     EnabledDefault = 5
	EnabledDefaultEnabledbutOffline EnabledDefault = 6
	EnabledDefaultNoDefault         EnabledDefault = 7
	EnabledDefaultQuiesce           EnabledDefault = 9
	EnabledDefaultDMTFReserved      EnabledDefault = 32768
	EnabledDefaultVendorReserved    EnabledDefault = 65535
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
	EnabledStateDMTFReserved
	EnabledStateVendorReserved
)

const (
	OperationalStatusUnknown OperationalStatus = iota
	OperationalStatusOther
	OperationalStatusOK
	OperationalStatusDegraded
	OperationalStatusStressed
	OperationalStatusPredictiveFailure
	OperationalStatusError
	OperationalStatusNonRecoverableError
	OperationalStatusStarting
	OperationalStatusStopping
	OperationalStatusStopped
	OperationalStatusInService
	OperationalStatusNoContact
	OperationalStatusLostCommunication
	OperationalStatusAborted
	OperationalStatusDormant
	OperationalStatusSupportingEntityinError
	OperationalStatusCompleted
	OperationalStatusPowerMode
	OperationalStatusRelocating
	OperationalStatusDMTFReserved
	OperationalStatusVendorReserved
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
	RequestedStateDMTFReserved
	RequestedStateVendorReserved
	RequestedStateUnknown RequestedState = 0
)

const (
	SecurityValuesOther SecurityValues = iota + 1
	SecurityValuesUnknown
	SecurityValuesNone
	SecurityValuesReadOnly
	SecurityValuesLockedOut
	SecurityValuesBootBypass
	SecurityValuesBootBypassAndReadOnly
)
