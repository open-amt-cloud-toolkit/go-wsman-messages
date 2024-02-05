/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package general

const (
	AMT_GeneralSettings string = "AMT_GeneralSettings"
)

const (
	IPv4 PreferredAddressFamily = iota
	IPv6
)

const (
	PrivacyLevelDefault PrivacyLevel = iota
	PrivacyLevelEnhanced
	PrivacyLevelExtreme
)

const (
	AC PowerSource = iota
	DC
)

const (
	Disabled FeatureEnabled = iota
	Enabled
)
