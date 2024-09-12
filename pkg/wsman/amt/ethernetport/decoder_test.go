/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ethernetport

import "testing"

func TestLinkPolicy_String(t *testing.T) {
	tests := []struct {
		state    LinkPolicy
		expected string
	}{
		{LinkPolicyS0AC, "S0 AC"},
		{LinkPolicySxAC, "Sx AC"},
		{LinkPolicyS0DC, "S0 DC"},
		{LinkPolicySxDC, "Sx DC"},
		{LinkPolicy(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestLinkPreference_String(t *testing.T) {
	tests := []struct {
		state    LinkPreference
		expected string
	}{
		{LinkPreferenceME, "Management Engine"},
		{LinkPreferenceHOST, "Host"},
		{LinkPreference(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestLinkControl_String(t *testing.T) {
	tests := []struct {
		state    LinkControl
		expected string
	}{
		{LinkControlME, "Management Engine"},
		{LinkControlHOST, "Host"},
		{LinkControl(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestWLANLinkProtectionLevel_String(t *testing.T) {
	tests := []struct {
		state    WLANLinkProtectionLevel
		expected string
	}{
		{WLANLinkProtectionLevelOverride, "Override"},
		{WLANLinkProtectionLevelNone, "None"},
		{WLANLinkProtectionLevelPassive, "Passive"},
		{WLANLinkProtectionLevelHigh, "High"},
		{WLANLinkProtectionLevel(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestPhysicalConnectionType_String(t *testing.T) {
	tests := []struct {
		state    PhysicalConnectionType
		expected string
	}{
		{PhysicalConnectionIntegratedLANNIC, "Integrated LAN NIC"},
		{PhysicalConnectionDiscreteLANNIC, "Discrete LAN NIC"},
		{PhysicalConnectionLANviaThunderboltDock, "LAN via Thunderbolt Dock"},
		{PhysicalConnectionWirelessLAN, "Wireless LAN"},
		{PhysicalConnectionType(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestPhysicalNicMediumToString(t *testing.T) {
	tests := []struct {
		state    PhysicalNicMedium
		expected string
	}{
		{PhysicalNicMediumSMBUS, "SMBUS"},
		{PhysicalNicMediumPCIe, "PCIe"},
		{PhysicalNicMedium(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
