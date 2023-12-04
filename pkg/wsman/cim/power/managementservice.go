/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package power

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

type PowerState int

const (
	// Power On
	PowerOn PowerState = 2 // Verified Hardware Power On

	// Sleep - Light
	SleepLight PowerState = 3 // ?

	// Sleep - Deep
	SleepDeep PowerState = 4 // ?

	// Power Cycle (Off Soft)
	PowerCycleOffSoft PowerState = 6 // ?

	// Power Off - Hard
	PowerOffHard PowerState = 8 // Verfied Hardware Power Off

	// Hibernate
	Hibernate PowerState = 7 // ?

	// Power Off - Soft
	PowerOffSoft PowerState = 9 // ?

	// Power Cycle (Off Hard)
	PowerCycleOffHard PowerState = 5 // Verified Hardware Power Cycle (off then on)

	// Master Bus Reset
	MasterBusReset PowerState = 10 // Verified Hardware Reboot

	// Diagnostic Interrupt (NMI)
	DiagnosticInterruptNMI PowerState = 11 // ?

	// Power Off - Soft Graceful
	PowerOffSoftGraceful PowerState = 12 // ?

	// Power Off - Hard Graceful
	PowerOffHardGraceful PowerState = 13 // ?

	// Master Bus Reset Graceful
	MasterBusResetGraceful PowerState = 14 // ?

	// Power Cycle (Off - Soft Graceful)
	PowerCycleOffSoftGraceful PowerState = 15 // ?

	// Power Cycle (Off - Hard Graceful)
	PowerCycleOffHardGraceful PowerState = 16 // ?
)

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                        xml.Name                       `xml:"Body"`
		RequestPowerStateChange_OUTPUT RequestPowerStateChange_OUTPUT `xml:"RequestPowerStateChange_OUTPUT"`
	}

	RequestPowerStateChange_OUTPUT struct {
		ReturnValue int `xml:"ReturnValue"`
	}
)

type ManagementService struct {
	base   message.Base
	client client.WSMan
}

type PowerActionResponse struct {
	RequestPowerStateChange_OUTPUT message.ReturnValue
}

// NewPowerManagementService returns a new instance of the PowerManagementService struct.
func NewPowerManagementService(wsmanMessageCreator *message.WSManMessageCreator) ManagementService {
	return ManagementService{
		base: message.NewBase(wsmanMessageCreator, string(CIM_PowerManagementService)),
	}
}

func NewPowerManagementServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) ManagementService {
	return ManagementService{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_PowerManagementService, client),
		client: client,
	}
}

// RequestPowerStateChange defines the desired power state of the managed element, and when the element should be put into that state.
func (p ManagementService) RequestPowerStateChange(powerState PowerState) (response Response, err error) {
	header := p.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(CIM_PowerManagementService, RequestPowerStateChange), CIM_PowerManagementService, nil, "", "")
	body := fmt.Sprintf(`<Body><h:RequestPowerStateChange_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService"><h:PowerState>%d</h:PowerState><h:ManagedElement><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="CreationClassName">CIM_ComputerSystem</Selector><Selector Name="Name">ManagedSystem</Selector></SelectorSet></ReferenceParameters></h:ManagedElement></h:RequestPowerStateChange_INPUT></Body>`, powerState)
	response = Response{
		Message: &client.Message{
			XMLInput: p.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = p.base.Execute(response.Message)
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

// // Get retrieves the representation of the instance
// func (b ManagementService) Get() string {
// 	return b.base.Get(nil)
// }

// // Enumerates the instances of this class
// func (b ManagementService) Enumerate() string {
// 	return b.base.Enumerate()
// }

// // Pulls instances of this class, following an Enumerate operation
// func (b ManagementService) Pull(enumerationContext string) string {
// 	return b.base.Pull(enumerationContext)
// }
