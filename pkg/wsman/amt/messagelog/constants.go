/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package messagelog

const (
	AMT_MessageLog        string = "AMT_MessageLog"
	GetRecords            string = "GetRecords"
	PositionToFirstRecord string = "PositionToFirstRecord"
)

const (
	CharacterSetUnknown CharacterSet = iota
	CharacterSetOther
	CharacterSetASCII
	CharacterSetUnicode
	CharacterSetISO2022
	CharacterSetISO8859
	CharacterSetExtendedUNIXCode
	CharacterSetUTF8
	CharacterSetUCS2
	CharacterSetBitmappedData
	CharacterSetOctetString
	CharacterSetDefinedbyIndividualRecords
)

const (
	LastChangeUnknown LastChange = iota
	LastChangeAdd
	LastChangeDelete
	LastChangeModify
	LastChangeLogCleared
)

const (
	OverwritePolicyUnknown OverwritePolicy = iota
	OverwritePolicyOther
	OverwritePolicyWrapsWhenFull
	OverwritePolicyClearWhenNearFull
	OverwritePolicyOverwriteOutdatedWhenNeeded
	OverwritePolicyRemoveOutdatedRecords
	OverwritePolicyOverwriteSpecificRecords
	OverwritePolicyNeverOverwrite
)

const (
	CapabilitiesUnknown Capabilities = iota
	CapabilitiesOther
	CapabilitiesWriteRecordSupported
	CapabilitiesDeleteRecordSupported
	CapabilitiesCanMoveBackwardinLog
	CapabilitiesFreezeLogSupported
	CapabilitiesClearLogSupported
	CapabilitiesSupportsAddressingbyOrdinalRecordNumber
	CapabilitiesVariableLengthRecordsSupported
	CapabilitiesVariableFormatsforRecords
	CapabilitiesCanFlagRecordsforOverwrite
)

const (
	LogStateUnknown       LogState = 0
	LogStateNormal        LogState = 2
	LogStateErasing       LogState = 3
	LogStateNotApplicable LogState = 4
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
	HealthStateUnknown             HealthState = 0
	HealthStateOK                  HealthState = 5
	HealthStateDegradedWarning     HealthState = 10
	HealthStateMinorFailure        HealthState = 15
	HealthStateMajorFailure        HealthState = 20
	HealthStateCriticalFailure     HealthState = 25
	HealthStateNonRecoverableError HealthState = 30
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
