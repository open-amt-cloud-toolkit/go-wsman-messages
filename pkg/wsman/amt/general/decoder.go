/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package general

const (
	AMTGeneralSettings string = "AMT_GeneralSettings"
	ValueNotFound      string = "Value not found in map"
)

const (
	Default PrivacyLevel = iota
	Enhanced
	Extreme
)

// privacyLevelToString is a map of PrivacyLevel to their string representation.
var privacyLevelToString = map[PrivacyLevel]string{
	Default:  "Default",
	Enhanced: "Enhanced",
	Extreme:  "Extreme",
}

// String returns the string representation of the PrivacyLevel value.
func (p PrivacyLevel) String() string {
	if value, exists := privacyLevelToString[p]; exists {
		return value
	}

	return ValueNotFound
}

const (
	AC PowerSource = iota
	DC
)

// powerSourceToString is a map of PowerSource to their string representation.
var powerSourceToString = map[PowerSource]string{
	AC: "AC",
	DC: "DC",
}

// String returns the string representation of the PowerSource value.
func (p PowerSource) String() string {
	if value, exists := powerSourceToString[p]; exists {
		return value
	}

	return ValueNotFound
}

const (
	IPv4 PreferredAddressFamily = iota
	IPv6
)

// preferredAddressFamilyToString is a map of PreferredAddressFamily to their string representation.
var preferredAddressFamilyToString = map[PreferredAddressFamily]string{
	IPv4: "IPv4",
	IPv6: "IPv6",
}

// String returns the string representation of the PreferredAddressFamily value.
func (p PreferredAddressFamily) String() string {
	if value, exists := preferredAddressFamilyToString[p]; exists {
		return value
	}

	return ValueNotFound
}

const (
	AMTNetworkDisabled AMTNetwork = iota
	AMTNetworkEnabled
)

// amtNetworkEnabledToString is a map of FeatureEnabled values to their string representations.
var amtNetworkEnabledToString = map[AMTNetwork]string{
	AMTNetworkDisabled: "Disabled",
	AMTNetworkEnabled:  "Enabled",
}

// String returns the string representation of the FeatureEnabled value.
func (a AMTNetwork) String() string {
	if value, exists := amtNetworkEnabledToString[a]; exists {
		return value
	}

	return ValueNotFound
}

const (
	ThunderboltDockDisabled ThunderboltDock = iota
	ThunderboltDockEnabled
)

// thunderboltDockEnabledToString is a map of FeatureEnabled values to their string representations.
var thunderboltDockEnabledToString = map[ThunderboltDock]string{
	ThunderboltDockDisabled: "Disabled",
	ThunderboltDockEnabled:  "Enabled",
}

// String returns the string representation of the FeatureEnabled value.
func (t ThunderboltDock) String() string {
	if value, exists := thunderboltDockEnabledToString[t]; exists {
		return value
	}

	return ValueNotFound
}

const (
	DHCPSyncRequiresHostnameDisabled DHCPSyncRequiresHostname = iota
	DHCPSyncRequiresHostnameEnabled
)

// dhcpSyncRequiresHostnameToString is a map of DHCPSyncRequiresHostname values to their string representations.
var dhcpSyncRequiresHostnameToString = map[DHCPSyncRequiresHostname]string{
	DHCPSyncRequiresHostnameDisabled: "Disabled",
	DHCPSyncRequiresHostnameEnabled:  "Enabled",
}

// String returns the string representation of the DHCPSyncRequiresHostname value.
func (d DHCPSyncRequiresHostname) String() string {
	if value, exists := dhcpSyncRequiresHostnameToString[d]; exists {
		return value
	}

	return ValueNotFound
}
