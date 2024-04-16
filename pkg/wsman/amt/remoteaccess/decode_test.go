/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import "testing"

func TestTrigger_String(t *testing.T) {
	tests := []struct {
		state    Trigger
		expected string
	}{
		{TriggerUserInitiated, "UserInitiated"},
		{TriggerAlert, "Alert"},
		{TriggerPeriodic, "Periodic"},
		{TriggerHomeProvisioning, "HomeProvisioning"},
		{Trigger(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestMPSType_String(t *testing.T) {
	tests := []struct {
		state    MPSType
		expected string
	}{
		{ExternalMPS, "ExternalMPS"},
		{InternalMPS, "InternalMPS"},
		{BothMPS, "BothMPS"},
		{MPSType(999), "Value not found in map"},
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
		{ReturnValueNotPermitted, "NotPermitted"},
		{ReturnValueMaxLimitReached, "MaxLimitReached"},
		{ReturnValueInvalidParameter, "InvalidParameter"},
		{ReturnValueFlashWriteLimitExceeded, "FlashWriteLimitExceeded"},
		{ReturnValueDuplicate, "Duplicate"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
