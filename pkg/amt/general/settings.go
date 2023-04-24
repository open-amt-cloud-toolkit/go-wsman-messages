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
	PreferredAddressFamilyIPv4 PreferredAddressFamily = 0
	PreferredAddressFamilyIPv6 PreferredAddressFamily = 1
)

type AMTNetworkEnabled int

const (
	AMTNetworkEnabledDisabled AMTNetworkEnabled = 0
	AMTNetworkEnabledEnabled  AMTNetworkEnabled = 1
)

type PrivacyLevel int

const (
	PrivacyLevelDefault  PrivacyLevel = 0
	PrivacyLevelEnhanced PrivacyLevel = 1
	PrivacyLevelExtreme  PrivacyLevel = 2
)

type PowerSource int

const (
	PowerSourceAC PowerSource = 0
	PowerSourceDC PowerSource = 1
)

type ThunderboltDockEnabled int

const (
	ThunderboltDockEnabledDisabled ThunderboltDockEnabled = 0
	ThunderboltDockEnabledEnabled  ThunderboltDockEnabled = 1
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
