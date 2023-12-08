/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package remoteaccess

import (
	"encoding/xml"
	"encoding/json"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

const (
	AMT_RemoteAccessPolicyAppliesToMPS string = "AMT_RemoteAccessPolicyAppliesToMPS"
	AMT_RemoteAccessPolicyRule         string = "AMT_RemoteAccessPolicyRule"
	AMT_RemoteAccessService            string = "AMT_RemoteAccessService"
	AddMps                             string = "AddMpServer"
	AddRemoteAccessPolicyRule          string = "AddRemoteAccessPolicyRule"
)

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName      	   			  xml.Name     						`xml:"Body"`
		ServiceGetResponse 			  RemoteAccessService				`xml:"AMT_RemoteAccessService"`
		PolicyAppliesToMPSGetResponse RemoteAccessPolicyAppliesToMPS	`xml:"AMT_RemoteAccessPolicyAppliesToMPS"`
		PolicyRuleGetResponse		  RemoteAccessPolicyRule			`xml:"AMT_RemoteAccessPolicyRule"`
		EnumerateResponse 			  common.EnumerateResponse
		PullResponse      			  PullResponse						`xml:"PullResponse"`
	}
	RemoteAccessService struct {
		CreationClassName       string
		ElementName             string
		Name                    string
		SystemCreationClassName string
		SystemName              string
	}
	RemoteAccessPolicyAppliesToMPS struct {
		CreationClassName       string
 		Name                    string
		SystemCreationClassName string
 		SystemName              string
	}
	RemoteAccessPolicyRule struct {
		CreationClassName       string
		ElementName             string
		ExtendedData            string
		PolicyRuleName          string
		SystemCreationClassName string
		SystemName              string
		Trigger                 int
		TunnelLifeTime          int
	}
	PullResponse struct {
		RemoteAccessServiceItems			[]RemoteAccessService			 `xml:"Items>AMT_RemoteAccessService"`
		RemoteAccessPolicyAppliesToMPSItems []RemoteAccessPolicyAppliesToMPS `xml:"Items>AMT_RemoteAccessPolicyAppliesToMPS"`
		RemoteAccessPolicyRuleItems			[]RemoteAccessPolicyRule		 `xml:"Items>AMT_RemoteAccessPolicyRule"`
	}
)


func (w *Response) JSON() string {
	jsonOutput, err := json.Marshal(w.Body)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}