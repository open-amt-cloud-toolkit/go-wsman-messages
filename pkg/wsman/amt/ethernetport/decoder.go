/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ethernetport

const (
	AMT_EthernetPortSettings string = "AMT_EthernetPortSettings"
)

const (
	ConsoleTcpMaxRetransmissions_5 ConsoleTcpMaxRetransmissions = iota + 5
	ConsoleTcpMaxRetransmissions_6
	ConsoleTcpMaxRetransmissions_7
)

const (
	LinkPolicyS0AC LinkPolicy = 1   // available on S0 AC
	LinkPolicySxAC LinkPolicy = 14  // available on Sx AC
	LinkPolicyS0DC LinkPolicy = 16  // available on S0 DC
	LinkPolicySxDC LinkPolicy = 224 // available on Sx DC
)

// linkPolicyToString is a map of LinkPolicy values to their string representations
var linkPolicyToString = map[LinkPolicy]string{
	LinkPolicyS0AC: "LinkPolicyS0AC",
	LinkPolicySxAC: "LinkPolicySxAC",
	LinkPolicyS0DC: "LinkPolicyS0DC",
	LinkPolicySxDC: "LinkPolicySxDC",
}

// String returns the string representation of the LinkPolicy value
func (l LinkPolicy) String() string {
	if value, exists := linkPolicyToString[l]; exists {
		return value
	}

	return "Value not found in map"
}

const (
	LinkPreferenceME LinkPreference = iota + 1
	LinkPreferenceHOST
)

// linkPreferenceToString is a map of LinkPreference values to their string representations
var linkPreferenceToString = map[LinkPreference]string{
	LinkPreferenceME:   "LinkPreferenceME",
	LinkPreferenceHOST: "LinkPreferenceHOST",
}

// String returns the string representation of the LinkPreference value
func (l LinkPreference) String() string {
	if value, exists := linkPreferenceToString[l]; exists {
		return value
	}

	return "Value not found in map"
}

const (
	LinkControlME LinkControl = iota + 1
	LinkControlHOST
)

// linkControlToString is a map of LinkControl values to their string representations
var linkControlToString = map[LinkControl]string{
	LinkControlME:   "LinkControlME",
	LinkControlHOST: "LinkControlHOST",
}

// String returns the string representation of the LinkControl value
func (l LinkControl) String() string {
	if value, exists := linkControlToString[l]; exists {
		return value
	}

	return "Value not found in map"
}

const (
	WLANLinkProtectionLevelOverride WLANLinkProtectionLevel = iota
	WLANLinkProtectionLevelNone
	WLANLinkProtectionLevelPassive
	WLANLinkProtectionLevelHigh
)

// wlanLinkProtectionLevelToString is a map of WLANLinkProtectionLevel values to their string representations
var wlanLinkProtectionLevelToString = map[WLANLinkProtectionLevel]string{
	WLANLinkProtectionLevelOverride: "WLANLinkProtectionLevelOverride",
	WLANLinkProtectionLevelNone:     "WLANLinkProtectionLevelNone",
	WLANLinkProtectionLevelPassive:  "WLANLinkProtectionLevelPassive",
	WLANLinkProtectionLevelHigh:     "WLANLinkProtectionLevelHigh",
}

// String returns the string representation of the WLANLinkProtectionLevel value
func (w WLANLinkProtectionLevel) String() string {
	if value, exists := wlanLinkProtectionLevelToString[w]; exists {
		return value
	}

	return "Value not found in map"
}

const (
	PhysicalConnectionIntegratedLANNIC PhysicalConnectionType = iota
	PhysicalConnectionDiscreteLANNIC
	PhysicalConnectionLANviaThunderboltDock
	PhysicalConnectionWirelessLAN
)

// PhysicalConnectionTypeToString is a map of PhysicalConnectionType values to their string representations
var PhysicalConnectionTypeToString = map[PhysicalConnectionType]string{
	PhysicalConnectionIntegratedLANNIC:      "PhysicalConnectionIntegratedLANNIC",
	PhysicalConnectionDiscreteLANNIC:        "PhysicalConnectionDiscreteLANNIC",
	PhysicalConnectionLANviaThunderboltDock: "PhysicalConnectionLANviaThunderboltDock",
	PhysicalConnectionWirelessLAN:           "PhysicalConnectionWirelessLAN",
}

// String returns the string representation of the PhysicalConnectionType value
func (p PhysicalConnectionType) String() string {
	if value, exists := PhysicalConnectionTypeToString[p]; exists {
		return value
	}

	return "Value not found in map"
}

const (
	PhysicalNicMediumSMBUS PhysicalNicMedium = iota
	PhysicalNicMediumPCIe
)

// PhysicalNicMediumToString is a map of PhysicalNicMedium values to their string representations
var PhysicalNicMediumToString = map[PhysicalNicMedium]string{
	PhysicalNicMediumSMBUS: "PhysicalNicMediumSMBUS",
	PhysicalNicMediumPCIe:  "PhysicalNicMediumPCIe",
}

// String returns the string representation of the PhysicalNicMedium value
func (p PhysicalNicMedium) String() string {
	if value, exists := PhysicalNicMediumToString[p]; exists {
		return value
	}

	return "Value not found in map"
}
