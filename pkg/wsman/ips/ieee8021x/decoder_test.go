/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import "testing"

func TestEnabled_String(t *testing.T) {
	tests := []struct {
		state    Enabled
		expected string
	}{
		{EnabledWithCertificates, "EnabledWithCertificates"},
		{Disabled, "Disabled"},
		{EnabledWithoutCertificates, "EnabledWithoutCertificates"},
		{Enabled(999), "Value not found in map"},
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
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
