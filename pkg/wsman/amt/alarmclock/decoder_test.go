/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import "testing"

func TestReturnValue_String(t *testing.T) {
	tests := []struct {
		state    ReturnValue
		expected string
	}{
		{Success, "Success"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()

		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
