/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/common"
)

const AMT_TLSSettingData = "AMT_TLSSettingData"

type (
	Response struct {
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName           xml.Name   `xml:"Body"`
		TlsSetting        TlsSetting `xml:"AMT_TLSSettingData"`
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}
	TlsSetting struct {
		AcceptNonSecureConnections    bool
		ElementName                   string
		Enabled                       bool
		InstanceID                    string
		MutualAuthentication          bool
		NonSecureConnectionsSupported *bool
	}
	PullResponse struct {
		TlsSettingItems []TlsSetting `xml:"Items>AMT_TLSSettingData"`
	}
)

// TLSSettingData supports aliased namespace for the PUT call to AMT
// TLSSettings responses from AMT have different aliases, and the PUT call requires namespace
// PUT response  -- <a:Body><g:AMT_TLSSettingData>
// PULL response -- <a:Body><g:PullResponse><g:Items><h:AMT_TLSSettingData>
type TLSSettingData struct {
	XMLName                    xml.Name `xml:"h:AMT_TLSSettingData"`
	H                          string   `xml:"xmlns:h,attr"`
	AcceptNonSecureConnections bool     `xml:"h:AcceptNonSecureConnections"`
	ElementName                string   `xml:"h:ElementName,omitempty"`
	Enabled                    bool     `xml:"h:Enabled"`
	InstanceID                 string   `xml:"h:InstanceID,omitempty"`
	MutualAuthentication       bool     `xml:"h:MutualAuthentication"`
	TrustedCN                  string   `xml:"h:TrustedCN,omitempty"`
}

type SettingData struct {
	base message.Base
}

func NewTLSSettingData(wsmanMessageCreator *message.WSManMessageCreator) SettingData {
	return SettingData{
		base: message.NewBase(wsmanMessageCreator, AMT_TLSSettingData),
	}
}

// Get retrieves the representation of the instance
func (TLSSettingData SettingData) Get() string {
	return TLSSettingData.base.Get(nil)
}

// Enumerate the instances of this class
func (TLSSettingData SettingData) Enumerate() string {
	return TLSSettingData.base.Enumerate()
}

// Pull instances of this class, following an Enumerate operation
func (TLSSettingData SettingData) Pull(enumerationContext string) string {
	return TLSSettingData.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (TLSSettingData SettingData) Put(data TLSSettingData) string {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: data.InstanceID,
	}
	return TLSSettingData.base.Put(&data, true, &selector)
}
