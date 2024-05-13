/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package optin

const (
	IPSOptInService string = "IPS_OptInService"
	ValueNotFound   string = "Value not found in map"
)

const (
	OptInRequiredNone OptInRequired = iota
	OptInRequiredKVM
	OptInRequiredAll OptInRequired = 4294967295
)

// optInRequiredToString is a map of OptInRequired value to string.
var optInRequiredToString = map[OptInRequired]string{
	OptInRequiredNone: "None",
	OptInRequiredKVM:  "KVM",
	OptInRequiredAll:  "All",
}

// String returns a human-readable string representation of the OptInRequired enumeration.
func (o OptInRequired) String() string {
	if s, ok := optInRequiredToString[o]; ok {
		return s
	}

	return ValueNotFound
}

const (
	NotStarted OptInState = iota
	Requested
	Displayed
	Received
	InSession
)

// optInStateToString is a map of OptInState value to string.
var optInStateToString = map[OptInState]string{
	NotStarted: "NotStarted",
	Requested:  "Requested",
	Displayed:  "Displayed",
	Received:   "Received",
	InSession:  "InSession",
}

// String returns a human-readable string representation of the OptInState enumeration.
func (o OptInState) String() string {
	if s, ok := optInStateToString[o]; ok {
		return s
	}

	return ValueNotFound
}

const (
	CanModifyOptInPolicyFalse CanModifyOptInPolicy = 0
	CanModifyOptInPolicyTrue  CanModifyOptInPolicy = 1
)

// canModifyOptInPolicyToString is a map of CanModifyOptInPolicy value to string.
var canModifyOptInPolicyToString = map[CanModifyOptInPolicy]string{
	CanModifyOptInPolicyFalse: "False",
	CanModifyOptInPolicyTrue:  "True",
}

// String returns a human-readable string representation of the CanModifyOptInPolicy enumeration.
func (c CanModifyOptInPolicy) String() string {
	if s, ok := canModifyOptInPolicyToString[c]; ok {
		return s
	}

	return ValueNotFound
}

const (
	ReturnValueSuccess ReturnValue = iota
	ReturnValueInternalError
	ReturnValueInvalidState
	ReturnValueBlocked
	ReturnValueFailedFFS
)

// returnValueToString is a map of ReturnValue value to string.
var returnValueToString = map[ReturnValue]string{
	ReturnValueSuccess:       "Success",
	ReturnValueInternalError: "InternalError",
	ReturnValueInvalidState:  "InvalidState",
	ReturnValueBlocked:       "Blocked",
	ReturnValueFailedFFS:     "FailedFFS",
}

// String returns a human-readable string representation of the ReturnValue enumeration.
func (r ReturnValue) String() string {
	if s, ok := returnValueToString[r]; ok {
		return s
	}

	return ValueNotFound
}
