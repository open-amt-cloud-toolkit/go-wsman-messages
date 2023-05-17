/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package power

import (
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/actions"
)

type PowerState int

const (
	// Power On
	PowerOn PowerState = 2 + iota

	// Sleep - Light
	SleepLight

	// Sleep - Deep
	SleepDeep

	// Power Cycle (Off Soft)
	PowerCycleOffSoft

	// Power Off - Hard
	PowerOffHard

	// Hibernate
	Hibernate

	// Power Off - Soft
	PowerOffSoft

	// Power Cycle (Off Hard)
	PowerCycleOffHard

	// Master Bus Reset
	MasterBusReset

	// Diagnostic Interrupt (NMI)
	DiagnosticInterruptNMI

	// Power Off - Soft Graceful
	PowerOffSoftGraceful

	// Power Off - Hard Graceful
	PowerOffHardGraceful

	// Master Bus Reset Graceful
	MasterBusResetGraceful

	// Power Cycle (Off - Soft Graceful)
	PowerCycleOffSoftGraceful

	// Power Cycle (Off - Hard Graceful)
	PowerCycleOffHardGraceful
)

type ManagementService struct {
	base wsman.Base
}

const CIM_PowerManagementService = "CIM_PowerManagementService"

// NewPowerManagementService returns a new instance of the PowerManagementService struct.
func NewPowerManagementService(wsmanMessageCreator *wsman.WSManMessageCreator) ManagementService {
	return ManagementService{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_PowerManagementService)),
	}
}

// RequestPowerStateChange defines the desired power state of the managed element, and when the element should be put into that state.
func (p ManagementService) RequestPowerStateChange(powerState PowerState) string {
	header := p.base.WSManMessageCreator.CreateHeader(string(actions.RequestPowerStateChange), string(CIM_PowerManagementService), nil, "", "")
	body := fmt.Sprintf(`<Body><h:RequestPowerStateChange_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService"><h:PowerState>%d</h:PowerState><h:ManagedElement><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="CreationClassName">CIM_ComputerSystem</Selector><Selector Name="Name">ManagedSystem</Selector></SelectorSet></ReferenceParameters></h:ManagedElement></h:RequestPowerStateChange_INPUT></Body>`, powerState)
	return p.base.WSManMessageCreator.CreateXML(header, body)
}

// Get retrieves the representation of the instance
func (b ManagementService) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b ManagementService) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b ManagementService) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
