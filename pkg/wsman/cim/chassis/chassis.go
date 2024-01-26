/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package chassis

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewChassis returns a new instance of the Chassis struct.
func NewChassisWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_Chassis, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (chassis Package) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: chassis.base.Get(nil),
		},
	}

	err = chassis.base.Execute(response.Message)
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
func (chassis Package) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: chassis.base.Enumerate(),
		},
	}

	err = chassis.base.Execute(response.Message)
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
func (chassis Package) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: chassis.base.Pull(enumerationContext),
		},
	}
	err = chassis.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
