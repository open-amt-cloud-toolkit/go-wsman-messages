/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package general

import (
	"encoding/json"
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

const AMT_GeneralSettings = "AMT_GeneralSettings"

// OUTPUTS
type (
	Response struct {
		*wsman.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName            xml.Name        `xml:"Body"`
		AMTGeneralSettings GeneralSettings `xml:"AMT_GeneralSettings"`

		EnumerateResponse common.EnumerateResponse
		PullResponse 	  PullResponse
	}
	GeneralSettings struct {
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
		DHCPSyncRequiresHostname      int
	}
	PullResponse struct {
		Items []Item
	}
	Item struct {
		AMTGeneralSettings GeneralSettings `xml:"AMT_GeneralSettings"`
	}
)

type PreferredAddressFamily int

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

func (w *Response) JSON() string {
	jsonOutput, err := json.Marshal(w.Body)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}

type Settings struct {
	base   message.Base
	client wsman.WSManClient
}

func NewGeneralSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client wsman.WSManClient) Settings {
	return Settings{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_GeneralSettings, client),
		client: client,
	}
}

func NewGeneralSettings(wsmanMessageCreator *message.WSManMessageCreator) Settings {
	return Settings{
		base: message.NewBase(wsmanMessageCreator, AMT_GeneralSettings),
	}
}

// Get retrieves the representation of the instance
func (GeneralSettings Settings) Get() (response Response, err error) {
	response = Response{
		Message: &wsman.Message{
			XMLInput: GeneralSettings.base.Get(nil),
		},
	}
	// send the message to AMT
	err = GeneralSettings.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (GeneralSettings Settings) Enumerate() (response Response, err error) {
	response = Response{
		Message: &wsman.Message{
			XMLInput: GeneralSettings.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = GeneralSettings.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pulls instances of this class, following an Enumerate operation
func (GeneralSettings Settings) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &wsman.Message{
			XMLInput: GeneralSettings.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = GeneralSettings.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// // Put will change properties of the selected instance
// func (GeneralSettings Settings) Put(generalSettings GeneralSettings) string {
// 	return GeneralSettings.base.Put(generalSettings, false, nil)
// }
