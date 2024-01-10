/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package general

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Settings struct {
	base message.Base
}

// OUTPUTS
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name `xml:"Body"`
		GetResponse       GeneralSettingsResponse
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
		PutResponse       PutResponse
	}

	PullResponse struct {
		XMLName              xml.Name                  `xml:"PullResponse"`
		GeneralSettingsItems []GeneralSettingsResponse `xml:"Items>AMT_GeneralSettings"`
	}
	GeneralSettingsResponse struct {
		XMLName                       xml.Name               `xml:"AMT_GeneralSettings"`
		ElementName                   string                 `xml:"ElementName,omitempty"`
		InstanceID                    string                 `xml:"InstanceID,omitempty"`
		NetworkInterfaceEnabled       bool                   `xml:"NetworkInterfaceEnabled,omitempty"`
		DigestRealm                   string                 `xml:"DigestRealm,omitempty"`
		IdleWakeTimeout               int                    `xml:"IdleWakeTimeout,omitempty"`
		HostName                      string                 `xml:"HostName,omitempty"`
		DomainName                    string                 `xml:"DomainName,omitempty"`
		PingResponseEnabled           bool                   `xml:"PingResponseEnabled,omitempty"`
		WsmanOnlyMode                 bool                   `xml:"WsmanOnlyMode,omitempty"`
		PreferredAddressFamily        PreferredAddressFamily `xml:"PreferredAddressFamily,omitempty"`
		DHCPv6ConfigurationTimeout    int                    `xml:"DHCPv6ConfigurationTimeout,omitempty"`
		DDNSUpdateEnabled             bool                   `xml:"DDNSUpdateEnabled,omitempty"`
		DDNSUpdateByDHCPServerEnabled bool                   `xml:"DDNSUpdateByDHCPServerEnabled,omitempty"`
		SharedFQDN                    bool                   `xml:"SharedFQDN,omitempty"`
		HostOSFQDN                    string                 `xml:"HostOSFQDN,omitempty"`
		DDNSTTL                       int                    `xml:"DDNSTTL,omitempty"`
		AMTNetworkEnabled             AMTNetworkEnabled      `xml:"AMTNetworkEnabled,omitempty"`
		RmcpPingResponseEnabled       bool                   `xml:"RmcpPingResponseEnabled,omitempty"`
		DDNSPeriodicUpdateInterval    int                    `xml:"DDNSPeriodicUpdateInterval,omitempty"`
		PresenceNotificationInterval  int                    `xml:"PresenceNotificationInterval,omitempty"`
		PrivacyLevel                  PrivacyLevel           `xml:"PrivacyLevel,omitempty"`
		PowerSource                   PowerSource            `xml:"PowerSource,omitempty"`
		ThunderboltDockEnabled        ThunderboltDockEnabled `xml:"ThunderboltDockEnabled,omitempty"`
		OemID                         int                    `xml:"OemID,omitempty"`
		DHCPSyncRequiresHostname      int                    `xml:"DHCPSyncRequiresHostname,omitempty"` // AMT SDK Documentation Missing
	}

	PutResponse struct {
	}
)

type (
	GeneralSettingsRequest struct {
		XMLName                       xml.Name               `xml:"h:AMT_GeneralSettings"`
		H                             string                 `xml:"h:xmlns:h,attr"`
		ElementName                   string                 `xml:"h:ElementName,omitempty"`
		InstanceID                    string                 `xml:"h:InstanceID,omitempty"`
		NetworkInterfaceEnabled       bool                   `xml:"h:NetworkInterfaceEnabled,omitempty"`
		DigestRealm                   string                 `xml:"h:DigestRealm,omitempty"`
		IdleWakeTimeout               int                    `xml:"h:IdleWakeTimeout,omitempty"`
		HostName                      string                 `xml:"h:HostName,omitempty"`
		DomainName                    string                 `xml:"h:DomainName,omitempty"`
		PingResponseEnabled           bool                   `xml:"h:PingResponseEnabled,omitempty"`
		WsmanOnlyMode                 bool                   `xml:"h:WsmanOnlyMode,omitempty"`
		PreferredAddressFamily        PreferredAddressFamily `xml:"h:PreferredAddressFamily,omitempty"`
		DHCPv6ConfigurationTimeout    int                    `xml:"h:DHCPv6ConfigurationTimeout,omitempty"`
		DDNSUpdateEnabled             bool                   `xml:"h:DDNSUpdateEnabled,omitempty"`
		DDNSUpdateByDHCPServerEnabled bool                   `xml:"h:DDNSUpdateByDHCPServerEnabled,omitempty"`
		SharedFQDN                    bool                   `xml:"h:SharedFQDN,omitempty"`
		HostOSFQDN                    string                 `xml:"h:HostOSFQDN,omitempty"`
		DDNSTTL                       int                    `xml:"h:DDNSTTL,omitempty"`
		AMTNetworkEnabled             AMTNetworkEnabled      `xml:"h:AMTNetworkEnabled,omitempty"`
		RmcpPingResponseEnabled       bool                   `xml:"h:RmcpPingResponseEnabled,omitempty"`
		DDNSPeriodicUpdateInterval    int                    `xml:"h:DDNSPeriodicUpdateInterval,omitempty"`
		PresenceNotificationInterval  int                    `xml:"h:PresenceNotificationInterval,omitempty"`
		PrivacyLevel                  PrivacyLevel           `xml:"h:PrivacyLevel,omitempty"`
		PowerSource                   PowerSource            `xml:"h:PowerSource,omitempty"`
		ThunderboltDockEnabled        ThunderboltDockEnabled `xml:"h:ThunderboltDockEnabled,omitempty"`
		OemID                         int                    `xml:"h:OemID,omitempty"`
		DHCPSyncRequiresHostname      int                    `xml:"h:DHCPSyncRequiresHostname,omitempty"` // AMT SDK Documentation Missing
	}
)

type PreferredAddressFamily int
type PrivacyLevel int
type PowerSource int
type AMTNetworkEnabled FeatureEnabled
type ThunderboltDockEnabled FeatureEnabled
type FeatureEnabled int
