package general

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

type GeneralSettings struct {
	models.SettingData
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

type Response struct {
	AMT_GeneralSettings Settings
}

type Settings struct {
	base wsman.Base
}

func NewGeneralSettings(wsmanMessageCreator *wsman.WSManMessageCreator) Settings {
	return Settings{
		base: wsman.NewBase(wsmanMessageCreator, AMT_GeneralSettings),
	}
}
func (GeneralSettings Settings) Get() string {
	return GeneralSettings.base.Get(nil)
}
func (GeneralSettings Settings) Enumerate() string {
	return GeneralSettings.base.Enumerate()
}
func (GeneralSettings Settings) Pull(enumerationContext string) string {
	return GeneralSettings.base.Pull(enumerationContext)
}
func (GeneralSettings Settings) Put(generalSettings GeneralSettings) string {
	return GeneralSettings.base.Put(generalSettings, false, nil)
}
