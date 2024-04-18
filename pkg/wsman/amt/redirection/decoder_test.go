/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package redirection

import "testing"

func TestEnabledState_String(t *testing.T) {
	tests := []struct {
		state    EnabledState
		expected string
	}{
		{Unknown, "Unknown"},
		{Other, "Other"},
		{Enabled, "Enabled"},
		{Disabled, "Disabled"},
		{ShuttingDown, "ShuttingDown"},
		{NotApplicable, "NotApplicable"},
		{EnabledButOffline, "EnabledButOffline"},
		{InTest, "InTest"},
		{Deferred, "Deferred"},
		{Quiesce, "Quiesce"},
		{Starting, "Starting"},
		{DMTFReserved, "DMTFReserved"},
		{IDERAndSOLAreDisabled, "IDERAndSOLAreDisabled"},
		{IDERIsEnabledAndSOLIsDisabled, "IDERIsEnabledAndSOLIsDisabled"},
		{SOLIsEnabledAndIDERIsDisabled, "SOLIsEnabledAndIDERIsDisabled"},
		{IDERAndSOLAreEnabled, "IDERAndSOLAreEnabled"},
		{EnabledState(999), "Value not found in map"},
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
		{CompletedWithNoError, "CompletedWithNoError"},
		{NotSupported, "NotSupported"},
		{UnknownOrUnspecified, "UnknownOrUnspecified"},
		{CannotCompleteWithinTimeoutPeriod, "CannotCompleteWithinTimeoutPeriod"},
		{Failed, "Failed"},
		{InvalidParameter, "InvalidParameter"},
		{InUse, "InUse"},
		{MethodParametersCheckedJobStarted, "MethodParametersCheckedJobStarted"},
		{InvalidStateTransition, "InvalidStateTransition"},
		{UseOfTimeoutParameterNotSupported, "UseOfTimeoutParameterNotSupported"},
		{Busy, "Busy"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
