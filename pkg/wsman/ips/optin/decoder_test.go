/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package optin

import "testing"

func TestOptInRequired_String(t *testing.T) {
	tests := []struct {
		state    OptInRequired
		expected string
	}{
		{OptInRequiredNone, "None"},
		{OptInRequiredKVM, "KVM"},
		{OptInRequiredAll, "All"},
		{OptInRequired(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestOptInState_String(t *testing.T) {
	tests := []struct {
		state    OptInState
		expected string
	}{
		{NotStarted, "NotStarted"},
		{Requested, "Requested"},
		{Displayed, "Displayed"},
		{Received, "Received"},
		{InSession, "InSession"},
		{OptInState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestCanModifyOptInPolicy_String(t *testing.T) {
	tests := []struct {
		state    CanModifyOptInPolicy
		expected string
	}{
		{CanModifyOptInPolicyFalse, "False"},
		{CanModifyOptInPolicyTrue, "True"},
		{CanModifyOptInPolicy(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestReturnValue_String(t *testing.T) {
	tests := []struct {
		state    ReturnValue
		expected string
	}{
		{ReturnValueSuccess, "Success"},
		{ReturnValueInternalError, "InternalError"},
		{ReturnValueInvalidState, "InvalidState"},
		{ReturnValueBlocked, "Blocked"},
		{ReturnValueFailedFFS, "FailedFFS"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
