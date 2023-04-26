/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wsman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	mockWsmanMessageCreator := NewWSManMessageCreator("test-uri")
	base := NewBase(mockWsmanMessageCreator, "TestClass")

	t.Run("Enumerate", func(t *testing.T) {
		expected := "<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>0</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" /></Body></Envelope>"
		actual := base.Enumerate()
		assert.Equal(t, expected, actual)
	})

	t.Run("Get", func(t *testing.T) {
		selector := &Selector{Name: "Key", Value: "Value"}
		expected := "<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>1</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>"
		actual := base.Get(selector)
		assert.Equal(t, expected, actual)
	})

	t.Run("Pull", func(t *testing.T) {
		enumerationContext := "test-context"
		expected := "<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>2</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\"><EnumerationContext>test-context</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>"
		actual := base.Pull(enumerationContext)
		assert.Equal(t, expected, actual)
	})

	t.Run("Delete", func(t *testing.T) {
		selector := &Selector{Name: "Key", Value: "Value"}
		expected := "<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>3</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>"
		actual := base.Delete(selector)
		assert.Equal(t, expected, actual)
	})

	t.Run("Put", func(t *testing.T) {
		data := "test-data"
		customSelector := &Selector{Name: "Key", Value: "Value"}
		expected := "<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>4</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>"
		actual := base.Put(data, true, customSelector)
		assert.Equal(t, expected, actual)

		expectedNoSelector := "<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>5</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><string>test-data</string></Body></Envelope>"
		actualNoSelector := base.Put(data, false, customSelector)
		assert.Equal(t, expectedNoSelector, actualNoSelector)
	})

	t.Run("Create", func(t *testing.T) {
		data := "test-data"
		selector := &Selector{Name: "Key", Value: "Value"}
		expected := "<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Create</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>6</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>"
		actual := base.Create(data, selector)
		assert.Equal(t, expected, actual)
	})

	t.Run("RequestStateChange", func(t *testing.T) {
		actionName := "test-action"
		requestedState := 2
		expected := "<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>test-action</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>7</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:RequestStateChange_INPUT xmlns:h=\"test-uriTestClass\"><h:RequestedState>2</h:RequestedState></h:RequestStateChange_INPUT></Body></Envelope>"
		actual := base.RequestStateChange(actionName, requestedState)
		assert.Equal(t, expected, actual)
	})
}
