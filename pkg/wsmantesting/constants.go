package wsmantesting

import "fmt"

const (
	XMLHeader          = `<?xml version="1.0" encoding="utf-8"?>`
	Envelope           = `<Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>`
	EnumerationContext = `AC070000-0000-0000-0000-000000000000`
	OperationTimeout   = `PT60S`
	GET                = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	ENUMERATE          = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate"
	PULL               = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull"
	DELETE             = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete"
	ENUMERATE_BODY     = "<Enumerate xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" />"
)

var PULL_BODY = fmt.Sprintf(`<Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull>`, EnumerationContext)

var ExpectedResponse = func(messageID int, method, action, body string) string {
	return fmt.Sprintf(`%s%s%s</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/%s</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>%s</w:OperationTimeout></Header><Body>%s</Body></Envelope>`, XMLHeader, Envelope, action, method, messageID, OperationTimeout, body)
}
