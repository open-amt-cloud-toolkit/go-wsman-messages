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
	CapabilitiesUnknown Capabilities = iota
	CapabilitiesOther
	CapabilitiesWriteRecordSupported
	CapabilitiesDeleteRecordSupported
	CapabilitiesCanMoveBackwardInLog
	CapabilitiesFreezeLogSupported
	CapabilitiesClearLogSupported
	CapabilitiesSupportsAddressingByOrdinalRecordNumber
	CapabilitiesVariableLengthRecordsSupported
	CapabilitiesVariableFormatsForRecords
	CapabilitiesCanFlagRecordsForOverwrite
)

// capabilitiesString is a map of the capabilities to their string representation
var capabilitiesString = map[Capabilities]string{
	CapabilitiesUnknown:                                 "Unknown",
	CapabilitiesOther:                                   "Other",
	CapabilitiesWriteRecordSupported:                    "WriteRecordSupported",
	CapabilitiesDeleteRecordSupported:                   "DeleteRecordSupported",
	CapabilitiesCanMoveBackwardInLog:                    "CanMoveBackwardInLog",
	CapabilitiesFreezeLogSupported:                      "FreezeLogSupported",
	CapabilitiesClearLogSupported:                       "ClearLogSupported",
	CapabilitiesSupportsAddressingByOrdinalRecordNumber: "SupportsAddressingByOrdinalRecordNumber",
	CapabilitiesVariableLengthRecordsSupported:          "VariableLengthRecordsSupported",
	CapabilitiesVariableFormatsForRecords:               "VariableFormatsForRecords",
	CapabilitiesCanFlagRecordsForOverwrite:              "CanFlagRecordsForOverwrite",
}

// String returns the string representation of the capabilities
func (c Capabilities) String() string {
	if value, exists := capabilitiesString[c]; exists {
		return value
	}
	return "Value not found in map"
}

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
	CharacterSetDefinedByIndividualRecords
)

// characterSetString is a map of the character set to their string representation
var characterSetString = map[CharacterSet]string{
	CharacterSetUnknown:                    "Unknown",
	CharacterSetOther:                      "Other",
	CharacterSetASCII:                      "ASCII",
	CharacterSetUnicode:                    "Unicode",
	CharacterSetISO2022:                    "ISO2022",
	CharacterSetISO8859:                    "ISO8859",
	CharacterSetExtendedUNIXCode:           "ExtendedUNIXCode",
	CharacterSetUTF8:                       "UTF8",
	CharacterSetUCS2:                       "UCS2",
	CharacterSetBitmappedData:              "BitmappedData",
	CharacterSetOctetString:                "OctetString",
	CharacterSetDefinedByIndividualRecords: "DefinedByIndividualRecords",
}

// ConvertCharacterSetToString returns the string representation of the character set
func (c CharacterSet) String() string {
	if value, exists := characterSetString[c]; exists {
		return value
	}
	return "Value not found in map"
}

const (
	EnabledDefaultEnabled           EnabledDefault = 2
	EnabledDefaultDisabled          EnabledDefault = 3
	EnabledDefaultNotApplicable     EnabledDefault = 5
	EnabledDefaultEnabledButOffline EnabledDefault = 6
	EnabledDefaultNoDefault         EnabledDefault = 7
	EnabledDefaultQuiesce           EnabledDefault = 9
)

// enabledDefaultString is a map of the enabled default to their string representation
var enabledDefaultString = map[EnabledDefault]string{
	EnabledDefaultEnabled:           "Enabled",
	EnabledDefaultDisabled:          "Disabled",
	EnabledDefaultNotApplicable:     "NotApplicable",
	EnabledDefaultEnabledButOffline: "EnabledButOffline",
	EnabledDefaultNoDefault:         "NoDefault",
	EnabledDefaultQuiesce:           "Quiesce",
}

// String returns the string representation of the enabled default
func (e EnabledDefault) String() string {
	if value, exists := enabledDefaultString[e]; exists {
		return value
	}
	return "Value not found in map"
}

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)

// enabledStateString is a map of the enabled state to their string representation
var enabledStateString = map[EnabledState]string{
	EnabledStateUnknown:           "Unknown",
	EnabledStateOther:             "Other",
	EnabledStateEnabled:           "Enabled",
	EnabledStateDisabled:          "Disabled",
	EnabledStateShuttingDown:      "ShuttingDown",
	EnabledStateNotApplicable:     "NotApplicable",
	EnabledStateEnabledButOffline: "EnabledButOffline",
	EnabledStateInTest:            "InTest",
	EnabledStateDeferred:          "Deferred",
	EnabledStateQuiesce:           "Quiesce",
	EnabledStateStarting:          "Starting",
}

// String returns the string representation of the enabled state
func (e EnabledState) String() string {
	if value, exists := enabledStateString[e]; exists {
		return value
	}
	return "Value not found in map"
}

const (
	HealthStateUnknown             HealthState = 0
	HealthStateOK                  HealthState = 5
	HealthStateDegraded            HealthState = 10
	HealthStateMinorFailure        HealthState = 20
	HealthStateMajorFailure        HealthState = 25
	HealthStateCriticalFailure     HealthState = 30
	HealthStateNonRecoverableError HealthState = 35
)

// healthStateString is a map of the health state to their string representation
var healthStateString = map[HealthState]string{
	HealthStateUnknown:             "Unknown",
	HealthStateOK:                  "OK",
	HealthStateDegraded:            "Degraded",
	HealthStateMinorFailure:        "MinorFailure",
	HealthStateMajorFailure:        "MajorFailure",
	HealthStateCriticalFailure:     "CriticalFailure",
	HealthStateNonRecoverableError: "NonRecoverableError",
}

// String returns the string representation of the health state
func (h HealthState) String() string {
	if value, exists := healthStateString[h]; exists {
		return value
	}
	return "Value not found in map"
}

const (
	LastChangeUnknown LastChange = iota
	LastChangeAdd
	LastChangeDelete
	LastChangeModify
	LastChangeLogCleared
)

// lastChangeString is a map of the last change to their string representation
var lastChangeString = map[LastChange]string{
	LastChangeUnknown:    "Unknown",
	LastChangeAdd:        "Add",
	LastChangeDelete:     "Delete",
	LastChangeModify:     "Modify",
	LastChangeLogCleared: "LogCleared",
}

// String returns the string representation of the last change
func (l LastChange) String() string {
	if value, exists := lastChangeString[l]; exists {
		return value
	}
	return "Value not found in map"
}

const (
	LogStateUnknown       LogState = 0
	LogStateNormal        LogState = 2
	LogStateErasing       LogState = 3
	LogStateNotApplicable LogState = 5
)

// logStateString is a map of the log state to their string representation
var logStateString = map[LogState]string{
	LogStateUnknown:       "Unknown",
	LogStateNormal:        "Normal",
	LogStateErasing:       "Erasing",
	LogStateNotApplicable: "NotApplicable",
}

// String returns the string representation of the log state
func (l LogState) String() string {
	if value, exists := logStateString[l]; exists {
		return value
	}
	return "Value not found in map"
}

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
	OperationalStatusSupportingEntityInError
	OperationalStatusCompleted
	OperationalStatusPowerMode
	OperationalStatusRelocating
)

// operationalStatusString is a map of the operational status to their string representation
var operationalStatusString = map[OperationalStatus]string{
	OperationalStatusUnknown:                 "Unknown",
	OperationalStatusOther:                   "Other",
	OperationalStatusOK:                      "OK",
	OperationalStatusDegraded:                "Degraded",
	OperationalStatusStressed:                "Stressed",
	OperationalStatusPredictiveFailure:       "PredictiveFailure",
	OperationalStatusError:                   "Error",
	OperationalStatusNonRecoverableError:     "NonRecoverableError",
	OperationalStatusStarting:                "Starting",
	OperationalStatusStopping:                "Stopping",
	OperationalStatusStopped:                 "Stopped",
	OperationalStatusInService:               "InService",
	OperationalStatusNoContact:               "NoContact",
	OperationalStatusLostCommunication:       "LostCommunication",
	OperationalStatusAborted:                 "Aborted",
	OperationalStatusDormant:                 "Dormant",
	OperationalStatusSupportingEntityInError: "SupportingEntityInError",
	OperationalStatusCompleted:               "Completed",
	OperationalStatusPowerMode:               "PowerMode",
	OperationalStatusRelocating:              "Relocating",
}

// String returns the string representation of the operational status
func (o OperationalStatus) String() string {
	if value, exists := operationalStatusString[o]; exists {
		return value
	}
	return "Value not found in map"
}

const (
	OverwritePolicyUnknown OverwritePolicy = iota
	OverwritePolicyOther
	OverwritePolicyWrapWhenFull
	OverwritePolicyClearWhenFull
	OverwritePolicyOverwriteOutdatedWhenNeeded
	OverwritePolicyRemoveOutdatedRecords
	OverwritePolicyOverwriteSpecificRecords
	OverwritePolicyNeverOverwrite
)

// overwritePolicyString is a map of the overwrite policy to their string representation
var overwritePolicyString = map[OverwritePolicy]string{
	OverwritePolicyUnknown:                     "Unknown",
	OverwritePolicyOther:                       "Other",
	OverwritePolicyWrapWhenFull:                "WrapWhenFull",
	OverwritePolicyClearWhenFull:               "ClearWhenFull",
	OverwritePolicyOverwriteOutdatedWhenNeeded: "OverwriteOutdatedWhenNeeded",
	OverwritePolicyRemoveOutdatedRecords:       "RemoveOutdatedRecords",
	OverwritePolicyOverwriteSpecificRecords:    "OverwriteSpecificRecords",
	OverwritePolicyNeverOverwrite:              "NeverOverwrite",
}

// String returns the string representation of the overwrite policy
func (o OverwritePolicy) String() string {
	if value, exists := overwritePolicyString[o]; exists {
		return value
	}
	return "Value not found in map"
}

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

// requestedStateString is a map of the requested state to their string representation
var requestedStateString = map[RequestedState]string{
	RequestedStateUnknown:       "Unknown",
	RequestedStateEnabled:       "Enabled",
	RequestedStateDisabled:      "Disabled",
	RequestedStateShutDown:      "ShutDown",
	RequestedStateNoChange:      "NoChange",
	RequestedStateOffline:       "Offline",
	RequestedStateTest:          "Test",
	RequestedStateDeferred:      "Deferred",
	RequestedStateQuiesce:       "Quiesce",
	RequestedStateReboot:        "Reboot",
	RequestedStateReset:         "Reset",
	RequestedStateNotApplicable: "NotApplicable",
}

// String returns the string representation of the requested state
func (r RequestedState) String() string {
	if value, exists := requestedStateString[r]; exists {
		return value
	}
	return "Value not found in map"
}

const (
	GetRecordsReturnValueCompletedWithNoError GetRecordsReturnValue = 0
	GetRecordsReturnValueNotSupported         GetRecordsReturnValue = 1
	GetRecordsReturnValueInvalidRecordPointed GetRecordsReturnValue = 2
	GetRecordsReturnValueNoRecordExistsInLog  GetRecordsReturnValue = 3
)

// getRecordsReturnValueString is a map of the GetRecordsReturnValue to their string representation
var getRecordsReturnValueString = map[GetRecordsReturnValue]string{
	GetRecordsReturnValueCompletedWithNoError: "CompletedWithNoError",
	GetRecordsReturnValueNotSupported:         "NotSupported",
	GetRecordsReturnValueInvalidRecordPointed: "InvalidRecordPointed",
	GetRecordsReturnValueNoRecordExistsInLog:  "NoRecordExistsInLog",
}

// String returns the string representation of the GetRecordsReturnValue value
func (g GetRecordsReturnValue) String() string {
	if value, exists := getRecordsReturnValueString[g]; exists {
		return value
	}
	return "Value not found in map"
}

const (
	PositionToFirstRecordReturnValueCompletedWithNoError PositionToFirstRecordReturnValue = 0
	PositionToFirstRecordReturnValueNotSupported         PositionToFirstRecordReturnValue = 1
	PositionToFirstRecordReturnValueNoRecordExists       PositionToFirstRecordReturnValue = 2
)

// positionToFirstRecordReturnValueString is a map of the PositionToFirstRecordReturnValue to their string representation
var positionToFirstRecordReturnValueString = map[PositionToFirstRecordReturnValue]string{
	PositionToFirstRecordReturnValueCompletedWithNoError: "CompletedWithNoError",
	PositionToFirstRecordReturnValueNotSupported:         "NotSupported",
	PositionToFirstRecordReturnValueNoRecordExists:       "NoRecordExists",
}

// String returns the string representation of the PositionToFirstRecordReturnValue value
func (p PositionToFirstRecordReturnValue) String() string {
	if value, exists := positionToFirstRecordReturnValueString[p]; exists {
		return value
	}
	return "Value not found in map"
}
