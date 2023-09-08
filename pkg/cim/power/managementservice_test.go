/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package power

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestPowerManagementService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewPowerManagementService(wsmanMessageCreator)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid cim_PowerManagementService Get wsman message", "CIM_PowerManagementService", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid cim_PowerManagementService Enumerate wsman message", "CIM_PowerManagementService", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid cim_PowerManagementService Pull wsman message", "CIM_PowerManagementService", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			{"Should create a valid cim_PowerManagementService ChangeBootOrder wsman message", "CIM_PowerManagementService", "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService/RequestPowerStateChange", "<h:RequestPowerStateChange_INPUT xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService\"><h:PowerState>8</h:PowerState><h:ManagedElement><Address xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\"><ResourceURI xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem</ResourceURI><SelectorSet xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\"><Selector Name=\"CreationClassName\">CIM_ComputerSystem</Selector><Selector Name=\"Name\">ManagedSystem</Selector></SelectorSet></ReferenceParameters></h:ManagedElement></h:RequestPowerStateChange_INPUT>", func() string {
				var powerState PowerState = PowerOffSoft
				return elementUnderTest.RequestPowerStateChange(powerState)
			}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				correctResponse := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response := test.responseFunc()
				if response != correctResponse {
					assert.Equal(t, correctResponse, response)
				}
			})
		}
	})
}
