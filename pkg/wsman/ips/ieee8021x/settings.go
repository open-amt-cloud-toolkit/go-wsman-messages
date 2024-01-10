/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips/methods"
)

// NewIEEE8021xSettings returns a new instance of the IEEE8021xSettings struct.
func NewIEEE8021xSettings(wsmanMessageCreator *message.WSManMessageCreator) Settings {
	return Settings{
		base: message.NewBase(wsmanMessageCreator, IPS_IEEE8021xSettings),
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
