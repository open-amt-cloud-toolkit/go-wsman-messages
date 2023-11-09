/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/json"
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

const AMT_RemoteAccessPolicyRule = "AMT_RemoteAccessPolicyRule"

type (
	ResponseRule struct {
		*client.Message
		XMLName  xml.Name       `xml:"Envelope"`
		Header   message.Header `xml:"Header"`
		BodyRule BodyRule       `xml:"Body"`
	}
	BodyRule struct {
		XMLName          xml.Name         `xml:"Body"`
		RemotePolicyRule RemotePolicyRule `xml:"AMT_RemoteAccessPolicyRule"`

		EnumerateResponse common.EnumerateResponse
	}
	RemotePolicyRule struct {
		CreationClassName       string
		ElementName             string
		ExtendedData            string
		PolicyRuleName          string
		SystemCreationClassName string
		SystemName              string
		Trigger                 int
		TunnelLifeTime          int
	}
)

type RemoteAccessPolicyRule struct {
	Trigger        Trigger
	TunnelLifeTime int
	ExtendedData   string
}

type Trigger uint8

const (
	UserInitiated Trigger = iota
	Alert
	Periodic
	HomeProvisioning
)

type PolicyRule struct {
	base         message.Base
	clientPolicy client.WSMan
}

func (w *ResponseRule) JSONRule() string {
	jsonOutput, err := json.Marshal(w.BodyRule)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}

func NewPolicyRuleWithClient(wsmanMessageCreator *message.WSManMessageCreator, clientPolicy client.WSMan) PolicyRule {
	return PolicyRule{
		base:         message.NewBaseWithClient(wsmanMessageCreator, AMT_RemoteAccessPolicyRule, clientPolicy),
		clientPolicy: clientPolicy,
	}
}

func NewRemoteAccessPolicyRule(wsmanMessageCreator *message.WSManMessageCreator) PolicyRule {
	return PolicyRule{
		base: message.NewBase(wsmanMessageCreator, AMT_RemoteAccessPolicyRule),
	}
}

// Get retrieves the representation of the instance
func (RemoteAccessPolicyRule PolicyRule) Get() (response ResponseRule, err error) {
	response = ResponseRule{
		Message: &client.Message{
			XMLInput: RemoteAccessPolicyRule.base.Get(nil),
		},
	}
	// send the message to AMT
	err = RemoteAccessPolicyRule.base.Execute(response.Message)
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
func (RemoteAccessPolicyRule PolicyRule) Enumerate() (response ResponseRule, err error) {
	response = ResponseRule{
		Message: &client.Message{
			XMLInput: RemoteAccessPolicyRule.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = RemoteAccessPolicyRule.base.Execute(response.Message)
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
// func (RemoteAccessPolicyRule PolicyRule) Pull(enumerationContext string) string {
// 	return RemoteAccessPolicyRule.base.Pull(enumerationContext)
// }

// // Put will change properties of the selected instance
// func (RemoteAccessPolicyRule PolicyRule) Put(remoteAccessPolicyRule RemoteAccessPolicyRule) string {
// 	return RemoteAccessPolicyRule.base.Put(remoteAccessPolicyRule, false, nil)
// }

// // Delete removes a the specified instance
// func (RemoteAccessPolicyRule PolicyRule) Delete(handle string) string {
// 	selector := message.Selector{Name: "PolicyRuleName", Value: handle}
// 	return RemoteAccessPolicyRule.base.Delete(selector)
// }
