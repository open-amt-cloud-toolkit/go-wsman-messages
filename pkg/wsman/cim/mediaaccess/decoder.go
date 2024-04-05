/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mediaaccess

const CIM_MediaAccessDevice string = "CIM_MediaAccessDevice"

const (
	Other Security = iota + 1
	Unknown
	None
	ReadOnly
	LockedOut
	BootBypass
	BootBypassAndReadOnly
)

// securityToString is a map of Security values to string
var securityToString = map[Security]string{
	Other:                 "Other",
	Unknown:               "Unknown",
	None:                  "None",
	ReadOnly:              "ReadOnly",
	LockedOut:             "LockedOut",
	BootBypass:            "BootBypass",
	BootBypassAndReadOnly: "BootBypassandReadOnly",
}

// String returns a human-readable string representation of the Security enumeration
func (e Security) String() string {
	if s, ok := securityToString[e]; ok {
		return s
	}
	return "Value not found in map"
}

const (
	UnknownCapabilities Capabilities = iota
	OtherCapabilities
	SequentialAccess
	RandomAccess
	SupportsWriting
	Encryption
	Compression
	SupportsRemovableMedia
	ManualCleaning
	AutomaticCleaning
	SmartNotification
	SupportsDualSidedMedia
	PreDismountEjectNotRequired
)

// capabilitiesToString is a map of Capabilities value to string
var capabilitiesToString = map[Capabilities]string{
	UnknownCapabilities:         "Unknown",
	OtherCapabilities:           "Other",
	SequentialAccess:            "SequentialAccess",
	RandomAccess:                "RandomAccess",
	SupportsWriting:             "SupportsWriting",
	Encryption:                  "Encryption",
	Compression:                 "Compression",
	SupportsRemovableMedia:      "SupportsRemovableMedia",
	ManualCleaning:              "ManualCleaning",
	AutomaticCleaning:           "AutomaticCleaning",
	SmartNotification:           "SmartNotification",
	SupportsDualSidedMedia:      "SupportsDualSidedMedia",
	PreDismountEjectNotRequired: "PreDismountEjectNotRequired",
}

// String returns a human-readable string representation of the Capabilities enumeration
func (e Capabilities) String() string {
	if s, ok := capabilitiesToString[e]; ok {
		return s
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

// enabledDefaultToString is a map of EnabledDefault value to string
var enabledDefaultToString = map[EnabledDefault]string{
	EnabledDefaultEnabled:           "Enabled",
	EnabledDefaultDisabled:          "Disabled",
	EnabledDefaultNotApplicable:     "NotApplicable",
	EnabledDefaultEnabledButOffline: "EnabledButOffline",
	EnabledDefaultNoDefault:         "NoDefault",
	EnabledDefaultQuiesce:           "Quiesce",
}

// String returns a human-readable string representation of the EnabledDefault enumeration
func (e EnabledDefault) String() string {
	if s, ok := enabledDefaultToString[e]; ok {
		return s
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

// enabledStateToString is a map of EnabledState value to string
var enabledStateToString = map[EnabledState]string{
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

// String returns a human-readable string representation of the EnabledState enumeration
func (e EnabledState) String() string {
	if s, ok := enabledStateToString[e]; ok {
		return s
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

// requestedStateToString is a map of RequestedState value to string
var requestedStateToString = map[RequestedState]string{
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

// String returns a human-readable string representation of the RequestedState enumeration
func (e RequestedState) String() string {
	if s, ok := requestedStateToString[e]; ok {
		return s
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

// operationalStatusToString is a map of OperationalStatus value to string
var operationalStatusToString = map[OperationalStatus]string{
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

// String returns a human-readable string representation of the OperationalStatus enumeration
func (e OperationalStatus) String() string {
	if s, ok := operationalStatusToString[e]; ok {
		return s
	}
	return "Value not found in map"
}
