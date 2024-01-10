/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ethernetport

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
		GetAndPutResponse SettingsResponse
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	PullResponse struct {
		XMLName           xml.Name           `xml:"PullResponse"`
		EthernetPortItems []SettingsResponse `xml:"Items>AMT_EthernetPortSettings"`
	}

	SettingsResponse struct {
		XMLName                      xml.Name `xml:"AMT_EthernetPortSettings"`
		ElementName                  string   `xml:"ElementName,omitempty"`
		InstanceID                   string   `xml:"InstanceID,omitempty"`
		VLANTag                      int      `xml:"VLANTag,omitempty"`
		SharedMAC                    bool     `xml:"SharedMAC,omitempty"`
		MACAddress                   string   `xml:"MACAddress,omitempty"`
		LinkIsUp                     bool     `xml:"LinkIsUp,omitempty"`
		LinkPolicy                   []int    `xml:"LinkPolicy,omitempty"`
		LinkPreference               int      `xml:"LinkPreference,omitempty"`
		LinkControl                  int      `xml:"LinkControl,omitempty"`
		SharedStaticIp               bool     `xml:"SharedStaticIp,omitempty"`
		SharedDynamicIP              bool     `xml:"SharedDynamicIP,omitempty"`
		IpSyncEnabled                bool     `xml:"IpSyncEnabled,omitempty"`
		DHCPEnabled                  bool     `xml:"DHCPEnabled,omitempty"`
		IPAddress                    string   `xml:"IPAddress,omitempty"`
		SubnetMask                   string   `xml:"SubnetMask,omitempty"`
		DefaultGateway               string   `xml:"DefaultGateway,omitempty"`
		PrimaryDNS                   string   `xml:"PrimaryDNS,omitempty"`
		SecondaryDNS                 string   `xml:"SecondaryDNS,omitempty"`
		ConsoleTcpMaxRetransmissions int      `xml:"ConsoleTcpMaxRetransmissions,omitempty"`
		WLANLinkProtectionLevel      int      `xml:"WLANLinkProtectionLevel,omitempty"`
		PhysicalConnectionType       int      `xml:"PhysicalConnectionType,omitempty"`
		PhysicalNicMedium            int      `xml:"PhysicalNicMedium,omitempty"`
	}
)

type SettingsRequest struct {
	XMLName                      xml.Name                     `xml:"h:AMT_EthernetPortSettings"`
	H                            string                       `xml:"xmlns:h,attr"`
	ElementName                  string                       `xml:"h:ElementName,omitempty"` // In Intel AMT Release 6.0 and later releases value is 'Intel(r) AMT Ethernet Port Settings'
	InstanceID                   string                       `xml:"h:InstanceID,omitempty"`  // In Intel AMT Release 6.0 and later releases value is 'Intel(r) AMT Ethernet Port Settings 0' for wired instance and 'Intel(r) AMT Ethernet Port Settings 1' for wireless instance
	VLANTag                      int                          `xml:"h:VLANTag,omitempty"`     // Not supported in AMT 4.x, 6.0 and later releases
	SharedMAC                    bool                         `xml:"h:SharedMAC"`             // Required property in AMT 3.0 and 3.2
	MACAddress                   string                       `xml:"h:MACAddress,omitempty"`  // Readonly
	LinkIsUp                     bool                         `xml:"h:LinkIsUp"`
	LinkPolicy                   []LinkPolicy                 `xml:"h:LinkPolicy,omitempty"`
	LinkPreference               LinkPreference               `xml:"h:LinkPreference,omitempty"`
	LinkControl                  LinkControl                  `xml:"h:LinkControl,omitempty"` // Readonly
	SharedStaticIp               bool                         `xml:"h:SharedStaticIp"`
	SharedDynamicIP              bool                         `xml:"h:SharedDynamicIP,omitempty"` // Readonly
	IpSyncEnabled                bool                         `xml:"h:IpSyncEnabled"`
	DHCPEnabled                  bool                         `xml:"h:DHCPEnabled"`
	IPAddress                    string                       `xml:"h:IPAddress,omitempty"`
	SubnetMask                   string                       `xml:"h:SubnetMask,omitempty"`
	DefaultGateway               string                       `xml:"h:DefaultGateway,omitempty"`
	PrimaryDNS                   string                       `xml:"h:PrimaryDNS,omitempty"`
	SecondaryDNS                 string                       `xml:"h:SecondaryDNS,omitempty"`
	ConsoleTcpMaxRetransmissions ConsoleTcpMaxRetransmissions `xml:"h:ConsoleTcpMaxRetransmissions,omitempty"`
	WLANLinkProtectionLevel      WLANLinkProtectionLevel      `xml:"h:WLANLinkProtectionLevel,omitempty"` // Readonly
	PhysicalConnectionType       PhysicalConnectionType       `xml:"h:PhysicalConnectionType,omitempty"`  // applicable in AMT 15.0 and later
	PhysicalNicMedium            PhysicalNicMedium            `xml:"h:PhysicalNicMedium,omitempty"`       // applicable in AMT 15.0 and later
}

type Selector message.Selector
type LinkPolicy int
type ConsoleTcpMaxRetransmissions int
type LinkPreference int
type LinkControl int
type WLANLinkProtectionLevel int
type PhysicalConnectionType int
type PhysicalNicMedium int
