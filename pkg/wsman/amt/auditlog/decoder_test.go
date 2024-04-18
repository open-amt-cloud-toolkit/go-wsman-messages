/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package auditlog

import "testing"

func TestOverwritePolicy_String(t *testing.T) {
	tests := []struct {
		state    OverwritePolicy
		expected string
	}{
		{Unknown, "Unknown"},
		{WrapsWhenFull, "WrapsWhenFull"},
		{NeverOverwrites, "NeverOverwrites"},
		{PartialRestrictedRollover, "PartialRestrictedRollover"},
		{OverwritePolicy(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestStoragePolicy_String(t *testing.T) {
	tests := []struct {
		state    StoragePolicy
		expected string
	}{
		{NoRollOver, "NoRollOver"},
		{RollOver, "RollOver"},
		{RestrictedRollOver, "RestrictedRollOver"},
		{StoragePolicy(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
