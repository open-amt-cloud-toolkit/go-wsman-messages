/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package ethernetport

const (
	AMT_EthernetPortSettings string = "AMT_EthernetPortSettings"
)

const (
	S0AC LinkPolicy = 1   // available on S0 AC
	SxAC LinkPolicy = 14  // available on Sx AC
	S0DC LinkPolicy = 16  // available on S0 DC
	SxDC LinkPolicy = 224 // available on Sx DC
)

const (
	LinkPreferenceME LinkPreference = iota + 1
	LinkPreferenceHOST
)

const (
	LinkControlME LinkControl = iota + 1
	LinkControlHOST
)

const (
	ConsoleTcpMaxRetransmissions5 ConsoleTcpMaxRetransmissions = iota + 5
	ConsoleTcpMaxRetransmissions6
	ConsoleTcpMaxRetransmissions7
)

const (
	OVERRIDE WLANLinkProtectionLevel = iota
	NONE
	PASSIVE
	HIGH
)

const (
	IntegratedLANNIC PhysicalConnectionType = iota
	DiscreteLANNIC
	LANviaThunderboldDock
	WirelessLAN
)

const (
	SMBUS PhysicalNicMedium = iota
	PCIe
)
