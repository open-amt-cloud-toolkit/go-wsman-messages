/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

const AMT_RemoteAccessPolicyAppliesToMPS = "AMT_RemoteAccessPolicyAppliesToMPS"

type RemoteAccessPolicyAppliesToMPS struct {
	XMLName xml.Name `xml:"h:AMT_RemoteAccessPolicyAppliesToMPS"`
	H       string   `xml:"xmlns:h,attr"`
	PolicySetAppliesToElement
	OrderOfAccess int     `xml:"h:OrderOfAccess"`
	MPSType       MPSType `xml:"h:MpsType"`
}
type Policy struct {
	models.ManagedElement
	CommonName     string   `xml:"h:CommonName"`
	PolicyKeywords []string `xml:"h:PolicyKeywords"`
}

type PolicySet struct {
	XMLName xml.Name `xml:"h:PolicySet"`
	Policy
	PolicyDecisionStrategy PolicyDecisionStrategy `xml:"h:PolicyDecisionStrategy"` // ValueMap={1, 2} Values={First Matching, All}
	PolicyRoles            []string               `xml:"h:PolicyRoles"`            // MaxLen=256
	Enabled                models.Enabled         `xml:"h:Enabled"`                // ValueMap={1, 2, 3} Values={Enabled, Disabled, Enabled For Debug}
}

type PolicySetAppliesToElement struct {
	PolicySet      PolicySet
	ManagedElement models.ManagedElement
}

/**
 * First Matching:1 | All:2
 */
type PolicyDecisionStrategy uint8

const (
	PolicyDecisionStrategyFirstMatching PolicyDecisionStrategy = 1
	PolicyDecisionStrategyAll           PolicyDecisionStrategy = 2
)

type MPSType int

const (
	ExternalMPS MPSType = iota
	InternalMPS
	BothMPS
)

type PolicyAppliesToMPS struct {
	base wsman.Base
}

func NewRemoteAccessPolicyAppliesToMPS(wsmanMessageCreator *wsman.WSManMessageCreator) PolicyAppliesToMPS {
	return PolicyAppliesToMPS{
		base: wsman.NewBase(wsmanMessageCreator, AMT_RemoteAccessPolicyAppliesToMPS),
	}
}

// Get retrieves the representation of the instance
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Get() string {
	return RemoteAccessPolicyAppliesToMPS.base.Get(nil)
}

// Enumerates the instances of this class
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Enumerate() string {
	return RemoteAccessPolicyAppliesToMPS.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Pull(enumerationContext string) string {
	return RemoteAccessPolicyAppliesToMPS.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Put(remoteAccessPolicyAppliesToMPS *RemoteAccessPolicyAppliesToMPS) string {
	return RemoteAccessPolicyAppliesToMPS.base.Put(remoteAccessPolicyAppliesToMPS, false, nil)
}

// Delete removes a the specified instance
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Delete(handle string) string {
	selector := wsman.Selector{Name: "Name", Value: handle}
	return RemoteAccessPolicyAppliesToMPS.base.Delete(selector)
}

// Creates a new instance of this class
func (RemoteAccessPolicyAppliesToMPS PolicyAppliesToMPS) Create(remoteAccessPolicyAppliesToMPS RemoteAccessPolicyAppliesToMPS) string {
	return RemoteAccessPolicyAppliesToMPS.base.Create(remoteAccessPolicyAppliesToMPS, nil)
}
