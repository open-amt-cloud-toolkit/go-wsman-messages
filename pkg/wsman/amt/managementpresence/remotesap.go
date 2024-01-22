/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package managementpresence facilitiates communication with Intel® AMT devices to configure Management Presence Remote Service Access Points (or an MPS) to be accessed by the Intel® AMT subsystem from remote.
package managementpresence

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewManagementPresenceRemoteSAPWithClient instantiates a new RemoteSAP
func NewManagementPresenceRemoteSAPWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) RemoteSAP {
	return RemoteSAP{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_ManagementPresenceRemoteSAP, client),
	}
}

// Get retrieves the representation of the instance
func (remoteSAP RemoteSAP) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: remoteSAP.base.Get(nil),
		},
	}
	// send the message to AMT
	err = remoteSAP.base.Execute(response.Message)
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
func (remoteSAP RemoteSAP) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: remoteSAP.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = remoteSAP.base.Execute(response.Message)
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
func (remoteSAP RemoteSAP) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: remoteSAP.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = remoteSAP.base.Execute(response.Message)
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

// Delete removes a the specified instance
func (remoteSAP RemoteSAP) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "Name", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: remoteSAP.base.Delete(selector),
		},
	}
	// send the message to AMT
	err = remoteSAP.base.Execute(response.Message)
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
