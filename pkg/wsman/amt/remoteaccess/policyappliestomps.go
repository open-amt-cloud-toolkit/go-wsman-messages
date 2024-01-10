/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

type PolicyAppliesToMPS struct {
	base message.Base
}

func NewRemoteAccessPolicyAppliesToMPSWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) PolicyAppliesToMPS {
	return PolicyAppliesToMPS{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_RemoteAccessPolicyAppliesToMPS, client),
	}
}

func NewRemoteAccessPolicyAppliesToMPS(wsmanMessageCreator *message.WSManMessageCreator) PolicyAppliesToMPS {
	return PolicyAppliesToMPS{
		base: message.NewBase(wsmanMessageCreator, AMT_RemoteAccessPolicyAppliesToMPS),
	}
}

// Get retrieves the representation of the instance
func (policyAppliesToMPS PolicyAppliesToMPS) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: policyAppliesToMPS.base.Get(nil),
		},
	}
	// send the message to AMT
	err = policyAppliesToMPS.base.Execute(response.Message)
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
func (policyAppliesToMPS PolicyAppliesToMPS) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: policyAppliesToMPS.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = policyAppliesToMPS.base.Execute(response.Message)
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
func (policyAppliesToMPS PolicyAppliesToMPS) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: policyAppliesToMPS.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = policyAppliesToMPS.base.Execute(response.Message)
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
func (policyAppliesToMPS PolicyAppliesToMPS) Put(remoteAccessPolicyAppliesToMPS *RemoteAccessPolicyAppliesToMPSRequest) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: policyAppliesToMPS.base.Put(remoteAccessPolicyAppliesToMPS, false, nil),
		},
	}
	// send the message to AMT
	err = policyAppliesToMPS.base.Execute(response.Message)
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
func (policyAppliesToMPS PolicyAppliesToMPS) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "Name", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: policyAppliesToMPS.base.Delete(selector),
		},
	}
	// send the message to AMT
	err = policyAppliesToMPS.base.Execute(response.Message)
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
