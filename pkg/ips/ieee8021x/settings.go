/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/ips/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/ips/methods"
)

type IEEE8021xSettings struct {
	models.IEEE8021xSettings
	Enabled       IEEE8021xSettingsEnabled `json:"Enabled,omitempty"`
	PxeTimeout    int                      `json:"PxeTimeout,omitempty"`
	AvailableInS0 bool                     `json:"AvailableInS0,omitempty"`
}

type IEEE8021xSettingsEnabled int

const (
	EnabledWithCertificates    IEEE8021xSettingsEnabled = 2
	Disabled                   IEEE8021xSettingsEnabled = 3
	EnabledWithoutCertificates IEEE8021xSettingsEnabled = 6
)

type Settings struct {
	base wsman.Base
}
type Certificate struct {
	XMLName                 xml.Name `xml:"h:SetCertificates_INPUT"`
	H                       string   `xml:"xmlns:h,attr"`
	ServerCertificateIssuer string   `xml:"h:ServerCertificateIssuer"`
	ClientCertificate       string   `xml:"h:ClientCertificate"`
}

const IPS_IEEE8021xSettings = "IPS_IEEE8021xSettings"

// NewIEEE8021xSettings returns a new instance of the IEEE8021xSettings struct.
func NewIEEE8021xSettings(wsmanMessageCreator *wsman.WSManMessageCreator) Settings {
	return Settings{
		base: wsman.NewBase(wsmanMessageCreator, IPS_IEEE8021xSettings),
	}
}

// Get retrieves the representation of the instance
func (b Settings) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Settings) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Settings) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (b Settings) Put(ieee8021xSettings IEEE8021xSettings) string {

	return b.base.Put(ieee8021xSettings, false, nil)
}

func (b Settings) SetCertificates(serverCertificateIssuer, clientCertificate string) string {
	header := b.base.WSManMessageCreator.CreateHeader(string(actions.SetCertificates), string(IPS_IEEE8021xSettings), nil, "", "")
	body := b.base.WSManMessageCreator.CreateBody(string(methods.SetCertificates_INPUT), string(IPS_IEEE8021xSettings),
		Certificate{
			H:                       "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings",
			ServerCertificateIssuer: serverCertificateIssuer,
			ClientCertificate:       clientCertificate,
		},
	)
	return b.base.WSManMessageCreator.CreateXML(header, body)
}
