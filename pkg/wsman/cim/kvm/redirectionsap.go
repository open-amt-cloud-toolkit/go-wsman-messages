/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kvm

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewKVMRedirectionSAP returns a new instance of the KVMRedirectionSAP struct.
func NewKVMRedirectionSAPWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) RedirectionSAP {
	return RedirectionSAP{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_KVMRedirectionSAP, client),
		client: client,
	}
}

// RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
func (redirectionSAP RedirectionSAP) RequestStateChange(requestedState KVMRedirectionSAPRequestedStateInputs) string {
	return redirectionSAP.base.RequestStateChange(methods.RequestStateChange(CIM_KVMRedirectionSAP), int(requestedState))
}

// Get retrieves the representation of the instance
func (redirectionSAP RedirectionSAP) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: redirectionSAP.base.Get(nil),
		},
	}

	err = redirectionSAP.base.Execute(response.Message)
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
func (redirectionSAP RedirectionSAP) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: redirectionSAP.base.Enumerate(),
		},
	}

	err = redirectionSAP.base.Execute(response.Message)
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
func (redirectionSAP RedirectionSAP) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: redirectionSAP.base.Pull(enumerationContext),
		},
	}
	err = redirectionSAP.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
