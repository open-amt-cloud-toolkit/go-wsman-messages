/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ethernetport

import (
	"encoding/json"
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name     `xml:"Body"`
		EthernetPort      EthernetPort `xml:"AMT_EthernetPortSettings"`
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	EthernetPort struct {
		DHCPEnabled            bool
		DefaultGateway         string
		ElementName            string
		InstanceID             string
		IpSyncEnabled          bool
		LinkIsUp               bool
		LinkPolicy             int
		MACAddress             string
		PhysicalConnectionType int
		PrimaryDNS             string
		SecondaryDNS           string
		SharedDynamicIP        bool
		SharedMAC              bool
		SharedStaticIp         bool
		SubnetMask             string
	}

	PullResponse struct {
		Items []Item
	}
	Item struct {
		EthernetPort EthernetPort `xml:"AMT_EthernetPortSettings"`
	}
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

type Selector message.Selector

type LinkPolicyValues int

func (w *Response) JSON() string {
	jsonOutput, err := json.Marshal(w.Body)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}

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
	base   message.Base
	client client.WSMan
}

func NewEthernetPortSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_EthernetPortSettings, client),
		client: client,
	}
}

func NewEthernetPortSettings(wsmanMessageCreator *message.WSManMessageCreator) Settings {
	return Settings{
		base: message.NewBase(wsmanMessageCreator, AMT_EthernetPortSettings),
	}
}

// Get retrieves the representation of the instance
func (s Settings) Get(selector Selector) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Get((*message.Selector)(&selector)),
		},
	}

	// send the message to AMT
	err = s.base.Execute(response.Message)
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
func (s Settings) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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

// // Pulls instances of this class, following an Enumerate operation
func (s Settings) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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
// func (EthernetPortSettings Settings) Put(ethernetPortSettings EthernetPortSettings) string {
// 	return EthernetPortSettings.base.Put(ethernetPortSettings, false, nil)
// }
