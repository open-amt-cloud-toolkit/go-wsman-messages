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
		{LinkPolicyS0AC, "LinkPolicyS0AC"},
		{LinkPolicySxAC, "LinkPolicySxAC"},
		{LinkPolicyS0DC, "LinkPolicyS0DC"},
		{LinkPolicySxDC, "LinkPolicySxDC"},
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
		{LinkPreferenceME, "LinkPreferenceME"},
		{LinkPreferenceHOST, "LinkPreferenceHOST"},
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
		{LinkControlME, "LinkControlME"},
		{LinkControlHOST, "LinkControlHOST"},
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
		{WLANLinkProtectionLevelOverride, "WLANLinkProtectionLevelOverride"},
		{WLANLinkProtectionLevelNone, "WLANLinkProtectionLevelNone"},
		{WLANLinkProtectionLevelPassive, "WLANLinkProtectionLevelPassive"},
		{WLANLinkProtectionLevelHigh, "WLANLinkProtectionLevelHigh"},
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
		{PhysicalConnectionIntegratedLANNIC, "PhysicalConnectionIntegratedLANNIC"},
		{PhysicalConnectionDiscreteLANNIC, "PhysicalConnectionDiscreteLANNIC"},
		{PhysicalConnectionLANviaThunderboltDock, "PhysicalConnectionLANviaThunderboltDock"},
		{PhysicalConnectionWirelessLAN, "PhysicalConnectionWirelessLAN"},
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
		{PhysicalNicMediumSMBUS, "PhysicalNicMediumSMBUS"},
		{PhysicalNicMediumPCIe, "PhysicalNicMediumPCIe"},
		{PhysicalNicMedium(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
