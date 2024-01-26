/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips/methods"
)

// NewIEEE8021xSettings returns a new instance of the IEEE8021xSettings struct.
func NewIEEE8021xSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPS_IEEE8021xSettings, client),
	}
}

// Get retrieves the representation of the instance
func (settings Settings) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Get(nil),
		},
	}
	err = settings.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call
func (settings Settings) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Enumerate(),
		},
	}
	err = settings.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (settings Settings) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Pull(enumerationContext),
		},
	}
	err = settings.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Put will change properties of the selected instance
func (settings Settings) Put(ieee8021xSettings IEEE8021xSettingsRequest) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Put(ieee8021xSettings, false, nil),
		},
	}
	err = settings.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (settings Settings) SetCertificates(serverCertificateIssuer, clientCertificate string) (response Response, err error) {
	header := settings.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPS_IEEE8021xSettings, SetCertificates), IPS_IEEE8021xSettings, nil, "", "")
	body := settings.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetCertificates), IPS_IEEE8021xSettings,
		Certificate{
			H:                       "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings",
			ServerCertificateIssuer: serverCertificateIssuer,
			ClientCertificate:       clientCertificate,
		},
	)
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = settings.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
