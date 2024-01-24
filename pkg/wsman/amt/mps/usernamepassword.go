/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package mps facilitiates communication with Intel® AMT devices to configure the username and password used to access an MPS.
package mps

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewMPSUsernamePasswordWithClient instantiates a new UsernamePassword
func NewMPSUsernamePasswordWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) UsernamePassword {
	return UsernamePassword{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_MPSUsernamePassword, client),
	}
}

// Get retrieves the representation of the instance
func (usernamePassword UsernamePassword) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: usernamePassword.base.Get(nil),
		},
	}
	// send the message to AMT
	err = usernamePassword.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (usernamePassword UsernamePassword) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: usernamePassword.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = usernamePassword.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pulls instances of this class, following an Enumerate operation
func (usernamePassword UsernamePassword) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: usernamePassword.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = usernamePassword.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Put will change properties of the selected instance
func (usernamePassword UsernamePassword) Put(mpsUsernamePassword MPSUsernamePasswordRequest) (response Response, err error) {
	mpsUsernamePassword.H = fmt.Sprintf("%s%s", message.AMTSchema, AMT_MPSUsernamePassword)
	response = Response{
		Message: &client.Message{
			XMLInput: usernamePassword.base.Put(mpsUsernamePassword, false, nil),
		},
	}
	// send the message to AMT
	err = usernamePassword.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
