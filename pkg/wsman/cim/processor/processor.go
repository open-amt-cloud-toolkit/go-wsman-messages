/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package processor

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewProcessor returns a new instance of the Processor struct.
func NewProcessorWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_Processor, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (processor Package) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: processor.base.Get(nil),
		},
	}

	err = processor.base.Execute(response.Message)
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
func (processor Package) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: processor.base.Enumerate(),
		},
	}

	err = processor.base.Execute(response.Message)
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
func (processor Package) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: processor.base.Pull(enumerationContext),
		},
	}
	err = processor.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
