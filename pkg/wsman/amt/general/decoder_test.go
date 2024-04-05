/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package general

import "testing"

func TestPrivacyLevel_String(t *testing.T) {
	tests := []struct {
		state    PrivacyLevel
		expected string
	}{
		{Default, "Default"},
		{Enhanced, "Enhanced"},
		{Extreme, "Extreme"},
		{PrivacyLevel(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestPowerSource_String(t *testing.T) {
	tests := []struct {
		state    PowerSource
		expected string
	}{
		{AC, "AC"},
		{DC, "DC"},
		{PowerSource(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestPreferredAddressFamily_String(t *testing.T) {
	tests := []struct {
		state    PreferredAddressFamily
		expected string
	}{
		{IPv4, "IPv4"},
		{IPv6, "IPv6"},
		{PreferredAddressFamily(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestAMTNetwork_String(t *testing.T) {
	tests := []struct {
		state    AMTNetwork
		expected string
	}{
		{AMTNetworkDisabled, "Disabled"},
		{AMTNetworkEnabled, "Enabled"},
		{AMTNetwork(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestThunderboltDock_String(t *testing.T) {
	tests := []struct {
		state    ThunderboltDock
		expected string
	}{
		{ThunderboltDockDisabled, "Disabled"},
		{ThunderboltDockEnabled, "Enabled"},
		{ThunderboltDock(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestDHCPSyncRequiresHostname_String(t *testing.T) {
	tests := []struct {
		state    DHCPSyncRequiresHostname
		expected string
	}{
		{DHCPSyncRequiresHostnameDisabled, "Disabled"},
		{DHCPSyncRequiresHostnameEnabled, "Enabled"},
		{DHCPSyncRequiresHostname(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
