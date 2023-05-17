/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/actions"
)

const CIM_BootService = "CIM_BootService"

type BootServiceRole int

const (
	// IsNext corresponds to the value 0.
	IsNext BootServiceRole = iota

	// IsNextSingleUse corresponds to the value 1.
	IsNextSingleUse

	// IsDefault corresponds to the value 2.
	IsDefault
)

type Service struct {
	base wsman.Base
}

// NewBootService returns a new instance of the BootService struct.
func NewBootService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_BootService)),
	}
}

// Get retrieves the representation of the instance
func (b Service) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Service) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Service) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}

// SetBootConfigRole sets the role of the BootConfigSetting that is directly or indirectly associated to one or more ComputerSystems.
func (b Service) SetBootConfigRole(bootSource string, role BootServiceRole) string {
	header := b.base.WSManMessageCreator.CreateHeader(string(actions.SetBootConfigRole), string(CIM_BootService), nil, "", "")
	body := fmt.Sprintf(`<Body><h:SetBootConfigRole_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService"><h:BootConfigSetting><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="InstanceID">%s</Selector></SelectorSet></ReferenceParameters></h:BootConfigSetting><h:Role>%d</h:Role></h:SetBootConfigRole_INPUT></Body>`, bootSource, role)
	return b.base.WSManMessageCreator.CreateXML(header, body)
}
