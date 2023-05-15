/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ethernetport

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

type EthernetPortSettings struct {
	models.SettingData
	VLANTag                      int
	SharedMAC                    bool
	MACAddress                   string
	LinkIsUp                     bool
	LinkPolicy                   LinkPolicy
	LinkPreference               LinkPreference
	LinkControl                  LinkControl
	SharedStaticIp               bool
	SharedDynamicIP              bool
	IpSyncEnabled                bool
	DHCPEnabled                  bool
	IPAddress                    string
	SubnetMask                   string
	DefaultGateway               string
	PrimaryDNS                   string
	SecondaryDNS                 string
	ConsoleTcpMaxRetransmissions ConsoleTcpMaxRetransmissions
	WLANLinkProtectionLevel      WLANLinkProtectionLevel
	PhysicalConnectionType       PhysicalConnectionType
	PhysicalNicMedium            PhysicalNicMedium
}

type LinkPolicyValues int

const AMT_EthernetPortSettings = "AMT_EthernetPortSettings"

const (
	S0AC LinkPolicyValues = 1
	SxAC LinkPolicyValues = 14
	S0DC LinkPolicyValues = 16
	SxDC LinkPolicyValues = 224
)

type LinkPolicy []LinkPolicyValues

type LinkPreference int

const (
	LinkPreferenceME LinkPreference = iota + 1
	LinkPreferenceHOST
)

type LinkControl int

const (
	LinkControlME LinkControl = iota + 1
	LinkControlHOST
)

type ConsoleTcpMaxRetransmissions int

const (
	ConsoleTcpMaxRetransmissions5 ConsoleTcpMaxRetransmissions = iota + 5
	ConsoleTcpMaxRetransmissions6
	ConsoleTcpMaxRetransmissions7
)

type WLANLinkProtectionLevel int

const (
	OVERRIDE WLANLinkProtectionLevel = iota
	NONE
	PASSIVE
	HIGH
)

type PhysicalConnectionType int

const (
	IntegratedLANNIC PhysicalConnectionType = iota
	DiscreteLANNIC
	LANviaThunderboldDock
	WirelessLAN
)

type PhysicalNicMedium int

const (
	SMBUS PhysicalNicMedium = iota
	PCIe
)

type Settings struct {
	base wsman.Base
}

func NewEthernetPortSettings(wsmanMessageCreator *wsman.WSManMessageCreator) Settings {
	return Settings{
		base: wsman.NewBase(wsmanMessageCreator, AMT_EthernetPortSettings),
	}
}

// Get retrieves the representation of the instance
func (EthernetPortSettings Settings) Get() string {
	return EthernetPortSettings.base.Get(nil)
}

// Enumerates the instances of this class
func (EthernetPortSettings Settings) Enumerate() string {
	return EthernetPortSettings.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (EthernetPortSettings Settings) Pull(enumerationContext string) string {
	return EthernetPortSettings.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (EthernetPortSettings Settings) Put(ethernetPortSettings EthernetPortSettings) string {
	return EthernetPortSettings.base.Put(ethernetPortSettings, false, nil)
}
