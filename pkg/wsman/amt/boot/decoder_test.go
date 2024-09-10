/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"testing"
)

func TestFirmwareVerbosity_String(t *testing.T) {
	tests := []struct {
		state    FirmwareVerbosity
		expected string
	}{
		{SystemDefault, "SystemDefault"},
		{QuietMinimal, "QuietMinimal"},
		{VerboseAll, "VerboseAll"},
		{ScreenBlank, "ScreenBlank"},
		{FirmwareVerbosity(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestIDERBootDevice_String(t *testing.T) {
	tests := []struct {
		state    IDERBootDevice
		expected string
	}{
		{FloppyBoot, "FloppyBoot"},
		{CDBoot, "CDBoot"},
		{IDERBootDevice(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
