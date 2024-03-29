/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package message

import (
	"errors"
	"fmt"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	Err error
}

func (c *MockClient) Post(msg string) ([]byte, error) {
	var response []byte = nil
	return response, c.Err
}

func TestBaseWithClient(t *testing.T) {
	mockWsmanMessageCreator := NewWSManMessageCreator("test-uri")
	mockClient := MockClient{}
	MessageId := 0
	base := NewBaseWithClient(mockWsmanMessageCreator, "TestClass", &mockClient)
	t.Run("Execute with no error", func(t *testing.T) {
		mockClient.Err = nil
		message := client.Message{
			XMLInput: "TestMessage",
		}
		err := base.Execute(&message)
		assert.NoError(t, err)
	})
	t.Run("Execute returns error", func(t *testing.T) {
		mockClient.Err = errors.New("test error")
		message := client.Message{
			XMLInput: "TestMessage",
		}
		err := base.Execute(&message)
		assert.Error(t, err)
	})
	t.Run("Enumerate", func(t *testing.T) {
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" /></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Enumerate()
		assert.Equal(t, expected, actual)
	})

	t.Run("Get", func(t *testing.T) {
		selector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Get(selector)
		assert.Equal(t, expected, actual)
	})

	t.Run("Pull", func(t *testing.T) {
		enumerationContext := "test-context"
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\"><EnumerationContext>test-context</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Pull(enumerationContext)
		assert.Equal(t, expected, actual)
	})

	t.Run("Delete", func(t *testing.T) {
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Name\">Value</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Delete(Selector{Name: "Name", Value: "Value"})
		assert.Equal(t, expected, actual)
	})

	t.Run("Put", func(t *testing.T) {
		data := "test-data"
		customSelector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Put(data, true, customSelector)
		assert.Equal(t, expected, actual)

		expectedDefaultSelector := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"InstanceID\">test-data</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageId)
		MessageId++
		actualDefaultSelector := base.Put(data, true, nil)
		assert.Equal(t, expectedDefaultSelector, actualDefaultSelector)

		expectedNoSelector := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><string>test-data</string></Body></Envelope>", MessageId)
		MessageId++
		actualNoSelector := base.Put(data, false, customSelector)
		assert.Equal(t, expectedNoSelector, actualNoSelector)
	})

	t.Run("Create", func(t *testing.T) {
		data := "test-data"
		selector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Create</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Create(data, selector)
		assert.Equal(t, expected, actual)
	})

	t.Run("RequestStateChange", func(t *testing.T) {
		actionName := "test-action"
		requestedState := 2
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>test-action</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:RequestStateChange_INPUT xmlns:h=\"test-uriTestClass\"><h:RequestedState>2</h:RequestedState></h:RequestStateChange_INPUT></Body></Envelope>", MessageId)
		MessageId++
		actual := base.RequestStateChange(actionName, requestedState)
		assert.Equal(t, expected, actual)
	})
}

func TestBase(t *testing.T) {
	mockWsmanMessageCreator := NewWSManMessageCreator("test-uri")
	base := NewBase(mockWsmanMessageCreator, "TestClass")
	MessageId := 0

	t.Run("Enumerate", func(t *testing.T) {
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" /></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Enumerate()
		assert.Equal(t, expected, actual)
	})

	t.Run("Get", func(t *testing.T) {
		selector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Get(selector)
		assert.Equal(t, expected, actual)
	})

	t.Run("Pull", func(t *testing.T) {
		enumerationContext := "test-context"
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\"><EnumerationContext>test-context</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Pull(enumerationContext)
		assert.Equal(t, expected, actual)
	})

	t.Run("Delete", func(t *testing.T) {
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Name\">Value</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Delete(Selector{Name: "Name", Value: "Value"})
		assert.Equal(t, expected, actual)
	})

	t.Run("Put", func(t *testing.T) {
		data := "test-data"
		customSelector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Put(data, true, customSelector)
		assert.Equal(t, expected, actual)

		expectedDefaultSelector := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"InstanceID\">test-data</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageId)
		MessageId++
		actualDefaultSelector := base.Put(data, true, nil)
		assert.Equal(t, expectedDefaultSelector, actualDefaultSelector)

		expectedNoSelector := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><string>test-data</string></Body></Envelope>", MessageId)
		MessageId++
		actualNoSelector := base.Put(data, false, customSelector)
		assert.Equal(t, expectedNoSelector, actualNoSelector)
	})

	t.Run("Create", func(t *testing.T) {
		data := "test-data"
		selector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Create</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageId)
		MessageId++
		actual := base.Create(data, selector)
		assert.Equal(t, expected, actual)
	})

	t.Run("RequestStateChange", func(t *testing.T) {
		actionName := "test-action"
		requestedState := 2
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>test-action</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:RequestStateChange_INPUT xmlns:h=\"test-uriTestClass\"><h:RequestedState>2</h:RequestedState></h:RequestStateChange_INPUT></Body></Envelope>", MessageId)
		MessageId++
		actual := base.RequestStateChange(actionName, requestedState)
		assert.Equal(t, expected, actual)
	})
}
