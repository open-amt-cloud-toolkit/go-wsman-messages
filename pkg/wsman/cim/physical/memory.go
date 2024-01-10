/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewPhysicalMemory returns a new instance of the PhysicalMemory struct.
func NewPhysicalMemoryWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Memory {
	return Memory{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_PhysicalMemory, client),
		client: client,
	}
}

// TODO: Figure out how to call GET requiring resourceURIs and Selectors
// Get retrieves the representation of the instance

// Enumerates the instances of this class
func (memory Memory) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: memory.base.Enumerate(),
		},
	}

	err = memory.base.Execute(response.Message)
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
func (memory Memory) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: memory.base.Pull(enumerationContext),
		},
	}
	err = memory.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
