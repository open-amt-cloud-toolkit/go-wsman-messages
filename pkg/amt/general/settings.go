/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package general

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

type Response struct {
	XMLName xml.Name     `xml:"Envelope"`
	Header  wsman.Header `xml:"Header"`
	Body    Body         `xml:"Body"`
}

type Body struct {
	XMLName            xml.Name        `xml:"Body"`
	AMTGeneralSettings GeneralSettings `xml:"AMT_GeneralSettings"`
}

type GeneralSettings struct {
	models.SettingData
	XMLName                       xml.Name `xml:"AMT_GeneralSettings"`
	NetworkInterfaceEnabled       bool
	DigestRealm                   string
	IdleWakeTimeout               int
	HostName                      string
	DomainName                    string
	PingResponseEnabled           bool
	WsmanOnlyMode                 bool
	PreferredAddressFamily        PreferredAddressFamily
	DHCPv6ConfigurationTimeout    int
	DDNSUpdateEnabled             bool
	DDNSUpdateByDHCPServerEnabled bool
	SharedFQDN                    bool
	HostOSFQDN                    string
	DDNSTTL                       int
	AMTNetworkEnabled             AMTNetworkEnabled
	RmcpPingResponseEnabled       bool
	DDNSPeriodicUpdateInterval    int
	PresenceNotificationInterval  int
	PrivacyLevel                  PrivacyLevel
	PowerSource                   PowerSource
	ThunderboltDockEnabled        ThunderboltDockEnabled
	OemID                         int
}

type PreferredAddressFamily int

const AMT_GeneralSettings = "AMT_GeneralSettings"

const (
	IPv4 PreferredAddressFamily = iota
	IPv6
)

type PrivacyLevel int

const (
	PrivacyLevelDefault PrivacyLevel = iota
	PrivacyLevelEnhanced
	PrivacyLevelExtreme
)

type PowerSource int

const (
	AC PowerSource = iota
	DC
)

type AMTNetworkEnabled FeatureEnabled
type ThunderboltDockEnabled FeatureEnabled
type FeatureEnabled int

const (
	Disabled FeatureEnabled = iota
	Enabled
)

type Settings struct {
	base wsman.Base
}

func NewGeneralSettings(wsmanMessageCreator *wsman.WSManMessageCreator) Settings {
	return Settings{
		base: wsman.NewBase(wsmanMessageCreator, AMT_GeneralSettings),
	}
}

// Get retrieves the representation of the instance
func (GeneralSettings Settings) Get() string {
	return GeneralSettings.base.Get(nil)
}

// Enumerates the instances of this class
func (GeneralSettings Settings) Enumerate() string {
	return GeneralSettings.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (GeneralSettings Settings) Pull(enumerationContext string) string {
	return GeneralSettings.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (GeneralSettings Settings) Put(generalSettings GeneralSettings) string {
	return GeneralSettings.base.Put(generalSettings, false, nil)
}
