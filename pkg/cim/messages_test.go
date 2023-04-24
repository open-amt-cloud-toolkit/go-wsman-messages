package cim

import (
	"fmt"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/power"
	"github.com/stretchr/testify/assert"
)

const (
	xmlHeader          = `<?xml version="1.0" encoding="utf-8"?>`
	envelope           = `<Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>`
	enumerationContext = `AC070000-0000-0000-0000-000000000000`
	operationTimeout   = `PT60S`
	GET                = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	ENUMERATE          = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate"
	PULL               = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull"
	DELETE             = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete"
	ENUMERATE_BODY     = "<Enumerate xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" />"
)

var PULL_BODY = fmt.Sprintf(`<Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull>`, enumerationContext)

func TestCIM(t *testing.T) {
	messageID := 0
	cimClass := NewMessages()

	expectedResponse := func(method, action, body string) string {
		return fmt.Sprintf(`%s%s%s</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/%s</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>%s</w:OperationTimeout></Header><Body>%s</Body></Envelope>`, xmlHeader, envelope, action, method, messageID, operationTimeout, body)
	}

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid cim_BIOSElement Get wsman message", "CIM_BIOSElement", GET, "", cimClass.BIOSElement.Get},
			{"should create a valid cim_BootService Get wsman message", "CIM_BootService", GET, "", cimClass.BootService.Get},
			{"should create a valid cim_BootConfigSetting Get wsman message", "CIM_BootConfigSetting", GET, "", cimClass.BootConfigSetting.Get},
			{"should create a valid cim_BootSourceSetting Get wsman message", "CIM_BootSourceSetting", GET, "", cimClass.BootSourceSetting.Get},
			{"should create a valid cim_Card Get wsman message", "CIM_Card", GET, "", cimClass.Card.Get},
			{"should create a valid cim_Chassis Get wsman message", "CIM_Chassis", GET, "", cimClass.Chassis.Get},
			{"should create a valid cim_Chip Get wsman message", "CIM_Chip", GET, "", cimClass.Chip.Get},
			{"should create a valid cim_ComputerSystemPackage Get wsman message", "CIM_ComputerSystemPackage", GET, "", cimClass.ComputerSystemPackage.Get},
			{"should create a valid cim_IEEE8021xSettings Get wsman message", "CIM_IEEE8021xSettings", GET, "", cimClass.IEEE8021xSettings.Get},
			{"should create a valid cim_KVMRedirectionSAP Get wsman message", "CIM_KVMRedirectionSAP", GET, "", cimClass.KVMRedirectionSAP.Get},
			{"should create a valid cim_MediaAccessDevice Get wsman message", "CIM_MediaAccessDevice", GET, "", cimClass.MediaAccessDevice.Get},
			{"should create a valid cim_PhysicalMemory Get wsman message", "CIM_PhysicalMemory", GET, "", cimClass.PhysicalMemory.Get},
			{"should create a valid cim_PhysicalPackage Get wsman message", "CIM_PhysicalPackage", GET, "", cimClass.PhysicalPackage.Get},
			{"should create a valid cim_PowerManagementService Get wsman message", "CIM_PowerManagementService", GET, "", cimClass.PowerManagementService.Get},
			{"should create a valid cim_Processor Get wsman message", "CIM_Processor", GET, "", cimClass.Processor.Get},
			{"should create a valid cim_ServiceAvailableToElement Get wsman message", "CIM_ServiceAvailableToElement", GET, "", cimClass.ServiceAvailableToElement.Get},
			{"should create a valid cim_SoftwareIdentity Get wsman message", "CIM_SoftwareIdentity", GET, "", cimClass.SoftwareIdentity.Get},
			{"should create a valid cim_SystemPackaging Get wsman message", "CIM_SystemPackaging", GET, "", cimClass.SystemPackaging.Get},
			{"should create a valid cim_WiFiEndpointSettings Get wsman message", "CIM_WiFiEndpointSettings", GET, "", cimClass.WiFiEndpointSettings.Get},
			{"should create a valid cim_WiFiPort Get wsman message", "CIM_WiFiPort", GET, "", cimClass.WiFiPort.Get},
			//ENUMERATES
			{"should create a valid cim_BIOSElement Enumerate wsman message", "CIM_BIOSElement", ENUMERATE, ENUMERATE_BODY, cimClass.BIOSElement.Enumerate},
			{"should create a valid cim_BootService Enumerate wsman message", "CIM_BootService", ENUMERATE, ENUMERATE_BODY, cimClass.BootService.Enumerate},
			{"should create a valid cim_BootConfigSetting Enumerate wsman message", "CIM_BootConfigSetting", ENUMERATE, ENUMERATE_BODY, cimClass.BootConfigSetting.Enumerate},
			{"should create a valid cim_BootSourceSetting Enumerate wsman message", "CIM_BootSourceSetting", ENUMERATE, ENUMERATE_BODY, cimClass.BootSourceSetting.Enumerate},
			{"should create a valid cim_Card Enumerate wsman message", "CIM_Card", ENUMERATE, ENUMERATE_BODY, cimClass.Card.Enumerate},
			{"should create a valid cim_Chassis Enumerate wsman message", "CIM_Chassis", ENUMERATE, ENUMERATE_BODY, cimClass.Chassis.Enumerate},
			{"should create a valid cim_Chip Enumerate wsman message", "CIM_Chip", ENUMERATE, ENUMERATE_BODY, cimClass.Chip.Enumerate},
			{"should create a valid cim_ComputerSystemPackage Enumerate wsman message", "CIM_ComputerSystemPackage", ENUMERATE, ENUMERATE_BODY, cimClass.ComputerSystemPackage.Enumerate},
			{"should create a valid cim_IEEE8021xSettings Enumerate wsman message", "CIM_IEEE8021xSettings", ENUMERATE, ENUMERATE_BODY, cimClass.IEEE8021xSettings.Enumerate},
			{"should create a valid cim_KVMRedirectionSAP Enumerate wsman message", "CIM_KVMRedirectionSAP", ENUMERATE, ENUMERATE_BODY, cimClass.KVMRedirectionSAP.Enumerate},
			{"should create a valid cim_MediaAccessDevice Enumerate wsman message", "CIM_MediaAccessDevice", ENUMERATE, ENUMERATE_BODY, cimClass.MediaAccessDevice.Enumerate},
			{"should create a valid cim_PhysicalMemory Enumerate wsman message", "CIM_PhysicalMemory", ENUMERATE, ENUMERATE_BODY, cimClass.PhysicalMemory.Enumerate},
			{"should create a valid cim_PhysicalPackage Enumerate wsman message", "CIM_PhysicalPackage", ENUMERATE, ENUMERATE_BODY, cimClass.PhysicalPackage.Enumerate},
			{"should create a valid cim_PowerManagementService Enumerate wsman message", "CIM_PowerManagementService", ENUMERATE, ENUMERATE_BODY, cimClass.PowerManagementService.Enumerate},
			{"should create a valid cim_Processor Enumerate wsman message", "CIM_Processor", ENUMERATE, ENUMERATE_BODY, cimClass.Processor.Enumerate},
			{"should create a valid cim_ServiceAvailableToElement Enumerate wsman message", "CIM_ServiceAvailableToElement", ENUMERATE, ENUMERATE_BODY, cimClass.ServiceAvailableToElement.Enumerate},
			{"should create a valid cim_SoftwareIdentity Enumerate wsman message", "CIM_SoftwareIdentity", ENUMERATE, ENUMERATE_BODY, cimClass.SoftwareIdentity.Enumerate},
			{"should create a valid cim_SystemPackaging Enumerate wsman message", "CIM_SystemPackaging", ENUMERATE, ENUMERATE_BODY, cimClass.SystemPackaging.Enumerate},
			{"should create a valid cim_WiFiEndpointSettings Enumerate wsman message", "CIM_WiFiEndpointSettings", ENUMERATE, ENUMERATE_BODY, cimClass.WiFiEndpointSettings.Enumerate},
			{"should create a valid cim_WiFiPort Enumerate wsman message", "CIM_WiFiPort", ENUMERATE, ENUMERATE_BODY, cimClass.WiFiPort.Enumerate},
			//PULLS
			{"should create a valid cim_BIOSElement Pull wsman message", "CIM_BIOSElement", PULL, PULL_BODY, func() string { return cimClass.BIOSElement.Pull(enumerationContext) }},
			{"should create a valid cim_BootService Enumerate wsman message", "CIM_BootService", PULL, PULL_BODY, func() string { return cimClass.BootService.Pull(enumerationContext) }},
			{"should create a valid cim_BootConfigSetting Enumerate wsman message", "CIM_BootConfigSetting", PULL, PULL_BODY, func() string { return cimClass.BootConfigSetting.Pull(enumerationContext) }},
			{"should create a valid cim_BootSourceSetting Enumerate wsman message", "CIM_BootSourceSetting", PULL, PULL_BODY, func() string { return cimClass.BootSourceSetting.Pull(enumerationContext) }},
			{"should create a valid cim_Card Enumerate wsman message", "CIM_Card", PULL, PULL_BODY, func() string { return cimClass.Card.Pull(enumerationContext) }},
			{"should create a valid cim_Chassis Enumerate wsman message", "CIM_Chassis", PULL, PULL_BODY, func() string { return cimClass.Chassis.Pull(enumerationContext) }},
			{"should create a valid cim_Chip Enumerate wsman message", "CIM_Chip", PULL, PULL_BODY, func() string { return cimClass.Chip.Pull(enumerationContext) }},
			{"should create a valid cim_ComputerSystemPackage Enumerate wsman message", "CIM_ComputerSystemPackage", PULL, PULL_BODY, func() string { return cimClass.ComputerSystemPackage.Pull(enumerationContext) }},
			{"should create a valid cim_IEEE8021xSettings Enumerate wsman message", "CIM_IEEE8021xSettings", PULL, PULL_BODY, func() string { return cimClass.IEEE8021xSettings.Pull(enumerationContext) }},
			{"should create a valid cim_KVMRedirectionSAP Enumerate wsman message", "CIM_KVMRedirectionSAP", PULL, PULL_BODY, func() string { return cimClass.KVMRedirectionSAP.Pull(enumerationContext) }},
			{"should create a valid cim_MediaAccessDevice Enumerate wsman message", "CIM_MediaAccessDevice", PULL, PULL_BODY, func() string { return cimClass.MediaAccessDevice.Pull(enumerationContext) }},
			{"should create a valid cim_PhysicalMemory Enumerate wsman message", "CIM_PhysicalMemory", PULL, PULL_BODY, func() string { return cimClass.PhysicalMemory.Pull(enumerationContext) }},
			{"should create a valid cim_PhysicalPackage Enumerate wsman message", "CIM_PhysicalPackage", PULL, PULL_BODY, func() string { return cimClass.PhysicalPackage.Pull(enumerationContext) }},
			{"should create a valid cim_PowerManagementService Enumerate wsman message", "CIM_PowerManagementService", PULL, PULL_BODY, func() string { return cimClass.PowerManagementService.Pull(enumerationContext) }},
			{"should create a valid cim_Processor Enumerate wsman message", "CIM_Processor", PULL, PULL_BODY, func() string { return cimClass.Processor.Pull(enumerationContext) }},
			{"should create a valid cim_ServiceAvailableToElement Enumerate wsman message", "CIM_ServiceAvailableToElement", PULL, PULL_BODY, func() string { return cimClass.ServiceAvailableToElement.Pull(enumerationContext) }},
			{"should create a valid cim_SoftwareIdentity Enumerate wsman message", "CIM_SoftwareIdentity", PULL, PULL_BODY, func() string { return cimClass.SoftwareIdentity.Pull(enumerationContext) }},
			{"should create a valid cim_SystemPackaging Enumerate wsman message", "CIM_SystemPackaging", PULL, PULL_BODY, func() string { return cimClass.SystemPackaging.Pull(enumerationContext) }},
			{"should create a valid cim_WiFiEndpointSettings Enumerate wsman message", "CIM_WiFiEndpointSettings", PULL, PULL_BODY, func() string { return cimClass.WiFiEndpointSettings.Pull(enumerationContext) }},
			{"should create a valid cim_WiFiPort Enumerate wsman message", "CIM_WiFiPort", PULL, PULL_BODY, func() string { return cimClass.WiFiPort.Pull(enumerationContext) }},
			// DELETES
			// {"should create a valid cim_WiFiEndpointSettings Delete wsman message", "CIM_WiFiEndpointSettings", DELETE, "", func() string {
			// 	selector := &wsman.Selector{
			// 		Name:  "InstanceID",
			// 		Value: "Intel(r) AMT:WiFi Endpoint Settings home",
			// 	}
			// 	return cimClass.WiFiEndpointSettings.Delete(selector)
			// }},
			// REQUEST STATE CHANGE
			{"should create a valid cim_WiFiPort Request State Change wsman message", "CIM_WiFiPort", "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiPort/RequestStateChange", "<h:RequestStateChange_INPUT xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiPort\"><h:RequestedState>3</h:RequestedState></h:RequestStateChange_INPUT>", func() string { return cimClass.WiFiPort.RequestStateChange(3) }},
			// PowerManagementService
			{"Should create a valid cim_PowerManagementService ChangeBootOrder wsman message", "CIM_PowerManagementService", "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService/RequestPowerStateChange", "<h:RequestPowerStateChange_INPUT xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService\"><h:PowerState>8</h:PowerState><h:ManagedElement><Address xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\"><ResourceURI xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem</ResourceURI><SelectorSet xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\"><Selector Name=\"CreationClassName\">CIM_ComputerSystem</Selector><Selector Name=\"Name\">ManagedSystem</Selector></SelectorSet></ReferenceParameters></h:ManagedElement></h:RequestPowerStateChange_INPUT>", func() string {
				var powerState power.PowerState = power.PowerOffSoft
				return cimClass.PowerManagementService.RequestPowerStateChange(powerState)
			}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				correctResponse := expectedResponse(test.method, test.action, test.body)
				messageID++
				response := test.responseFunc()
				if response != correctResponse {
					assert.Equal(t, correctResponse, response)
				}
			})
		}
	})
}
