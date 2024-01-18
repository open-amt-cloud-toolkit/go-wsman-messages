/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportconfiguration

const (
	AMT_WiFiPortConfigurationService string = "AMT_WiFiPortConfigurationService"
	AddWiFiSettings                  string = "AddWiFiSettings"
)

const (
	LocalSyncDisabled LocalProfileSynchronizationEnabled = 0
	UnrestrictedSync  LocalProfileSynchronizationEnabled = 3
)

const (
	RelaxedPolicy NoHostCsmeSoftwarePolicy = iota
	AggressivePolicy
	Reserved
)

const (
	Disabled UEFIWiFiProfileShareEnabled = iota
	Enabled
)
