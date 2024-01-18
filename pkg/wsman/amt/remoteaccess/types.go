/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

// OUTPUTS
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                     xml.Name `xml:"Body"`
		RemoteAccessGetResponse     RemoteAccessResponse
		RemotePolicyRuleGetResponse RemotePolicyRuleResponse
		PolicyAppliesGetResponse    PolicyAppliesResponse
		EnumerateResponse           common.EnumerateResponse
		PullResponse                PullResponse
		AddMpServerResponse         AddMpServerResponse
		AddRemotePolicyRuleResponse AddRemotePolicyRuleResponse
	}
	RemoteAccessResponse struct {
		XMLName                 xml.Name `xml:"AMT_RemoteAccessService"`
		CreationClassName       string   `xml:"CreationClassName,omitempty"`
		ElementName             string   `xml:"ElementName,omitempty"`
		Name                    string   `xml:"Name,omitempty"`
		SystemCreationClassName string   `xml:"SystemCreationClassName,omitempty"`
		SystemName              string   `xml:"SystemName,omitempty"`
	}
	RemotePolicyRuleResponse struct {
		XMLName                 xml.Name `xml:"AMT_RemoteAccessPolicyRule"`
		CreationClassName       string   `xml:"CreationClassName,omitempty"`
		ElementName             string   `xml:"ElementName,omitempty"`
		ExtendedData            string   `xml:"ExtendedData,omitempty"`
		PolicyRuleName          string   `xml:"PolicyRuleName,omitempty"`
		SystemCreationClassName string   `xml:"SystemCreationClassName,omitempty"`
		SystemName              string   `xml:"SystemName,omitempty"`
		Trigger                 Trigger  `xml:"Trigger"`
		TunnelLifeTime          int      `xml:"TunnelLifeTime"`
	}
	PolicyAppliesResponse struct {
		XMLName        xml.Name               `xml:"AMT_RemoteAccessPolicyAppliesToMPS"`
		ManagedElement ManagedElementResponse `xml:"ManagedElement"`
		MpsType        MPSType                `xml:"MpsType"`
		OrderOfAccess  int                    `xml:"OrderOfAccess"`
		PolicySet      PolicySetResponse      `xml:"PolicySet"`
	}
	PullResponse struct {
		XMLName               xml.Name                   `xml:"PullResponse"`
		RemoteAccessItems     []RemoteAccessResponse     `xml:"Items>AMT_RemoteAccessService"`
		RemotePolicyRuleItems []RemotePolicyRuleResponse `xml:"Items>AMT_RemoteAccessPolicyRule"`
		PolicyAppliesItems    []PolicyAppliesResponse    `xml:"Items>AMT_RemoteAccessPolicyAppliesToMPS"`
	}
	AddMpServerResponse struct {
		XMLName  xml.Name `xml:"AddMpServer_OUTPUT"`
		MpServer MpServer `xml:"MpServer"`
	}
	AddRemotePolicyRuleResponse struct {
		XMLName            xml.Name           `xml:"AddRemoteAccessPolicyRule_OUTPUT"`
		PolicyRuleResponse PolicyRuleResponse `xml:"PolicyRule"`
		ReturnValue        int                `xml:"ReturnValue"`
	}
	MpServer struct {
		XMLName             xml.Name                    `xml:"MpServer"`
		Address             string                      `xml:"Address"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters"`
	}
	ManagedElementResponse struct {
		XMLName             xml.Name                    `xml:"ManagedElement"`
		Address             string                      `xml:"Address"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters"`
	}
	PolicySetResponse struct {
		XMLName             xml.Name                    `xml:"PolicySet"`
		Address             string                      `xml:"Address"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters"`
	}
	PolicyRuleResponse struct {
		XMLName             xml.Name                    `xml:"PolicyRule"`
		Address             string                      `xml:"Address"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters"`
	}
	ReferenceParametersResponse struct {
		XMLName     xml.Name            `xml:"ReferenceParameters"`
		ResourceURI string              `xml:"ResourceURI"`
		SelectorSet SelectorSetResponse `xml:"SelectorSet"`
	}
	SelectorSetResponse struct {
		XMLName   xml.Name           `xml:"SelectorSet"`
		Selectors []SelectorResponse `xml:"Selector"`
	}
	SelectorResponse struct {
		XMLName xml.Name `xml:"Selector"`
		Name    string   `xml:"Name,attr"`
		Text    string   `xml:",chardata"`
	}
)

// INPUTS
// Request Types
type (
	RemoteAccessPolicyRuleSelector message.Selector
	PolicyDecisionStrategy         int // First Matching:1 | All:2
	MPServerInfoFormat             int
	MPServerAuthMethod             int
	Trigger                        int
	MPSType                        int
	Enabled                        int

	AddMpServerRequest struct {
		XMLName     xml.Name           `xml:"h:AddMpServer_INPUT"`
		H           string             `xml:"xmlns:h,attr"`
		AccessInfo  string             `xml:"h:AccessInfo"`
		InfoFormat  MPServerInfoFormat `xml:"h:InfoFormat"`
		Port        int                `xml:"h:Port"`
		AuthMethod  MPServerAuthMethod `xml:"h:AuthMethod"`
		Username    string             `xml:"h:Username,omitempty"`
		Password    string             `xml:"h:Password,omitempty"`
		CommonName  string             `xml:"h:CN"`
		Certificate string             `xml:"h:Certificate,omitempty"`
	}
	RemoteAccessPolicyRuleRequest struct {
		XMLName        xml.Name `xml:"h:AddRemoteAccessPolicyRule_INPUT"`
		H              string   `xml:"xmlns:h,attr"`
		Trigger        Trigger  `xml:"h:Trigger"`
		TunnelLifeTime int      `xml:"h:TunnelLifeTime"`
		ExtendedData   string   `xml:"h:ExtendedData"`
	}
	RemoteAccessPolicyAppliesToMPSRequest struct {
		XMLName        xml.Name       `xml:"h:AMT_RemoteAccessPolicyAppliesToMPS"`
		H              string         `xml:"xmlns:h,attr"`
		ManagedElement ManagedElement `xml:"h:ManagedElement"`
		OrderOfAccess  int            `xml:"h:OrderOfAccess"`
		MPSType        MPSType        `xml:"h:MpsType"`
		PolicySet      PolicySet      `xml:"h:PolicySet"`
	}
	ManagedElement struct {
		Address             string              `xml:"b:Address"`
		B                   string              `xml:"xmlns:b,attr"`
		ReferenceParameters ReferenceParameters `xml:"b:ReferenceParameters"`
	}
	ReferenceParameters struct {
		ResourceURI string      `xml:"c:ResourceURI"`
		C           string      `xml:"xmlns:c,attr"`
		SelectorSet SelectorSet `xml:"c:SelectorSet"`
	}
	SelectorSet struct {
		Selectors []Selector `xml:"c:Selector"`
	}
	Selector struct {
		Name string `xml:"Name,attr"`
		Text string `xml:",chardata"`
	}
	PolicySet struct {
		Address             string              `xml:"b:Address"`
		B                   string              `xml:"xmlns:b,attr"`
		ReferenceParameters ReferenceParameters `xml:"b:ReferenceParameters"`
	}
)
