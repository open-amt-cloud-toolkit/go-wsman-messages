/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewPolicyRuleWithClient instantiates a new PolicyRule
func NewPolicyRuleWithClient(wsmanMessageCreator *message.WSManMessageCreator, clientPolicy client.WSMan) PolicyRule {
	return PolicyRule{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_RemoteAccessPolicyRule, clientPolicy),
	}
}

// Get retrieves the representation of the instance
func (policyRule PolicyRule) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: policyRule.base.Get(nil),
		},
	}
	// send the message to AMT
	err = policyRule.base.Execute(response.Message)
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
func (policyRule PolicyRule) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: policyRule.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = policyRule.base.Execute(response.Message)
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
func (policyRule PolicyRule) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: policyRule.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = policyRule.base.Execute(response.Message)
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
func (policyRule PolicyRule) Put(remoteAccessPolicyRule RemoteAccessPolicyRuleRequest) (response Response, err error) {
	remoteAccessPolicyRule.H = fmt.Sprintf("%s%s", message.AMTSchema, AMT_RemoteAccessPolicyRule)
	response = Response{
		Message: &client.Message{
			XMLInput: policyRule.base.Put(remoteAccessPolicyRule, false, nil),
		},
	}
	// send the message to AMT
	err = policyRule.base.Execute(response.Message)
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
func (policyRule PolicyRule) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "PolicyRuleName", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: policyRule.base.Delete(selector),
		},
	}
	// send the message to AMT
	err = policyRule.base.Execute(response.Message)
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
