/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package optin

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/ips/actions"
)

type Service struct {
	base wsman.Base
}

const IPS_OptInService = "IPS_OptInService"

// NewOptInService returns a new instance of the OptInService struct.
func NewOptInService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{
		base: wsman.NewBase(wsmanMessageCreator, string(IPS_OptInService)),
	}
}

// Gets the representation of OptInService.
func (b Service) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Service) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Service) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}

type OptInCode struct {
	XMLName   xml.Name `xml:"h:SendOptInCode_INPUT"`
	H         string   `xml:"xmlns:h,attr"`
	OptInCode int      `xml:"h:OptInCode"`
}

// Send the opt-in code to Intel(R) AMT.
func (b Service) SendOptInCode(optInCode int) string {
	header := b.base.WSManMessageCreator.CreateHeader(string(actions.SendOptInCode), string(IPS_OptInService), nil, "", "")
	body := b.base.WSManMessageCreator.CreateBody("SendOptInCode_INPUT", string(IPS_OptInService), OptInCode{
		H:         "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService",
		OptInCode: optInCode,
	})
	return b.base.WSManMessageCreator.CreateXML(header, body)
}

// Request an opt-in code.
func (b Service) StartOptIn() string {
	header := b.base.WSManMessageCreator.CreateHeader(string(actions.StartOptIn), string(IPS_OptInService), nil, "", "")
	body := b.base.WSManMessageCreator.CreateBody("StartOptIn_INPUT", string(IPS_OptInService), nil)
	return b.base.WSManMessageCreator.CreateXML(header, body)
}

// Cancel a previous opt-in code request.
func (b Service) CancelOptIn() string {
	header := b.base.WSManMessageCreator.CreateHeader(string(actions.CancelOptIn), string(IPS_OptInService), nil, "", "")
	body := b.base.WSManMessageCreator.CreateBody("CancelOptIn_INPUT", string(IPS_OptInService), nil)
	return b.base.WSManMessageCreator.CreateXML(header, body)
}
