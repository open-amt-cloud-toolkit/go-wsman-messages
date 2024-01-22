/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package optin

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips/actions"
)

// NewOptInService returns a new instance of the OptInService struct.
func NewOptInServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPS_OptInService, client),
	}
}

// Gets the representation of OptInService.
func (service Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Get(nil),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (service Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Enumerate(),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pulls instances of this class, following an Enumerate operation
func (service Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Pull(enumerationContext),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Send the opt-in code to Intel® AMT.
func (service Service) SendOptInCode(optInCode int) (response Response, err error) {
	header := service.base.WSManMessageCreator.CreateHeader(string(actions.SendOptInCode), string(IPS_OptInService), nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody("SendOptInCode_INPUT", string(IPS_OptInService), OptInCode{
		H:         "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService",
		OptInCode: optInCode,
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Request an opt-in code.
func (service Service) StartOptIn() (response Response, err error) {
	header := service.base.WSManMessageCreator.CreateHeader(string(actions.StartOptIn), string(IPS_OptInService), nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody("StartOptIn_INPUT", string(IPS_OptInService), nil)
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Cancel a previous opt-in code request.
func (service Service) CancelOptIn() (response Response, err error) {
	header := service.base.WSManMessageCreator.CreateHeader(string(actions.CancelOptIn), string(IPS_OptInService), nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody("CancelOptIn_INPUT", string(IPS_OptInService), nil)
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
