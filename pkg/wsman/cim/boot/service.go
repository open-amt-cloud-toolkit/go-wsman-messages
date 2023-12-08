/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

type Service struct {
	base   message.Base
	client client.WSMan
}

// NewBootService returns a new instance of the BootService struct.
func NewBootServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_BootService, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (Service Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: Service.base.Get(nil),
		},
	}

	err = Service.base.Execute(response.Message)
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
func (Service Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: Service.base.Enumerate(),
		},
	}

	err = Service.base.Execute(response.Message)
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
func (Service Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: Service.base.Pull(enumerationContext),
		},
	}
	err = Service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
