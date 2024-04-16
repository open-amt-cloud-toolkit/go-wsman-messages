/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package environmentdetection

import "testing"

func TestDetectionAlgorithm_String(t *testing.T) {
	tests := []struct {
		state    DetectionAlgorithm
		expected string
	}{
		{LocalDomains, "LocalDomains"},
		{RemoteURLs, "RemoteURLs"},
		{DetectionAlgorithm(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
