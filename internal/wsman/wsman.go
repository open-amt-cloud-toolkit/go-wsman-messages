/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wsman

import (
	"encoding/xml"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Selector struct {
	XMLName xml.Name `xml:"w:Selector,omitempty"`
	Name    string   `xml:"Name,attr"`
	Value   string   `xml:",chardata"`
}
type Selector_OUTPUT struct {
	XMLName xml.Name `xml:"Selector,omitempty"`
	Name    string   `xml:"Name,attr"`
	Value   string   `xml:",chardata"`
}

type ReturnValue struct {
	ReturnValue    int
	ReturnValueStr string
}
type WSManMessageCreator struct {
	MessageID        int
	XmlCommonPrefix  string
	XmlCommonEnd     string
	AnonymousAddress string
	DefaultTimeout   string
	ResourceURIBase  string
}

func NewWSManMessageCreator(resourceUriBase string) *WSManMessageCreator {
	return &WSManMessageCreator{
		MessageID:        0,
		XmlCommonPrefix:  `<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope">`,
		XmlCommonEnd:     `</Envelope>`,
		AnonymousAddress: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
		DefaultTimeout:   "PT60S",
		ResourceURIBase:  resourceUriBase,
	}
}
func (w WSManMessageCreator) CreateXML(header, body string) string {
	return w.XmlCommonPrefix + header + body + w.XmlCommonEnd
}

func (w *WSManMessageCreator) CreateHeader(action string, wsmanClass string, selector *Selector, address string, timeout string) string {
	header := "<Header>"
	header += fmt.Sprintf(`<a:Action>%s</a:Action><a:To>/wsman</a:To><w:ResourceURI>%s%s</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo>`, action, w.ResourceURIBase, wsmanClass, w.MessageID)
	w.MessageID++
	if address != "" {
		header += fmt.Sprintf(`<a:Address>%s</a:Address>`, address)
	} else {
		header += fmt.Sprintf(`<a:Address>%s</a:Address>`, w.AnonymousAddress)
	}
	header += "</a:ReplyTo>"
	if timeout != "" {
		header += fmt.Sprintf(`<w:OperationTimeout>%s</w:OperationTimeout>`, timeout)
	} else {
		header += fmt.Sprintf(`<w:OperationTimeout>%s</w:OperationTimeout>`, w.DefaultTimeout)
	}
	if selector != nil {
		header += w.createSelector(*selector)
	}
	header += "</Header>"
	return header
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}
func (w WSManMessageCreator) namespaceMe(subj interface{}, wsmanClass string) {
	ifaceValue := reflect.ValueOf(subj)
	// Check if the interface value is a pointer
	if ifaceValue.Kind() == reflect.Ptr {
		// Get the underlying value of the interface
		stype := ifaceValue.Elem()
		if stype.Kind() == reflect.Struct {
			field := stype.FieldByName("H")
			if field.IsValid() && field.CanAddr() {
				fieldPtr := field.Addr().Interface().(*string)
				*fieldPtr = fmt.Sprintf("%s%s", w.ResourceURIBase, wsmanClass)
			}
		}
	}
}
func (w WSManMessageCreator) CreateBody(method string, wsmanClass string, data interface{}) string {
	var str strings.Builder
	str.WriteString("<Body>")
	if data != nil {
		w.namespaceMe(data, wsmanClass)
		xmlString, err := xml.Marshal(data)
		if err != nil {
			log.Println(err)
		}
		str.WriteString(string(xmlString))
	} else {
		str.WriteString(fmt.Sprintf(`<h:%s xmlns:h="%s%s">`, method, w.ResourceURIBase, wsmanClass))
		str.WriteString(fmt.Sprintf(`</h:%s>`, method))
	}
	str.WriteString("</Body>")
	return str.String()
}

// createSelector creates a WSMAN string based on Selector Set information provided.
// It can be used in the header or body.
// selectorSet is the selector data being passed in. It could take many forms depending on the WSMAN call.
func (w *WSManMessageCreator) createSelector(selectorSet Selector) string {

	if selectorSet.Name != "" {
		return fmt.Sprintf(`<w:SelectorSet><w:Selector Name="%s">%s</w:Selector></w:SelectorSet>`, selectorSet.Name, selectorSet.Value)
	}
	return ""

	// result := "<w:SelectorSet>"
	// for propName, propValue := range v {
	// 	propValueMap := propValue.(map[string]interface{})
	// 	result += fmt.Sprintf(`<w:Selector Name="%s">`, propName)
	// 	if refParams, ok := propValueMap["ReferenceParameters"].(map[string]interface{}); ok {
	// 		address := propValueMap["Address"].(string)
	// 		resourceURI := refParams["ResourceURI"].(string)
	// 		selectorSet := refParams["SelectorSet"].(map[string]interface{})
	// 		selectorArray := selectorSet["Selector"].(map[string]interface{})
	// 		result += "<a:EndpointReference>"
	// 		result += fmt.Sprintf(`<a:Address>%s</a:Address><a:ReferenceParameters><w:ResourceURI>%s</w:ResourceURI><w:SelectorSet>`, address, resourceURI)

	// 		if name, ok := selectorArray["$.Name"].(string); ok {
	// 			value := selectorArray["_"].(string)
	// 			result += fmt.Sprintf(`<w:Selector Name="%s">%s</w:Selector>`, name, value)
	// 		}
	// 		result += "</w:SelectorSet></a:ReferenceParameters></a:EndpointReference>"
	// 	}
	// 	result += "</w:Selector>"
	// }
	// result += "</w:SelectorSet>"
	// return result

}

// createSelectorObjectForBody creates an object for the body using the given selector.
func (w *WSManMessageCreator) CreateSelectorObjectForBody(selector Selector) map[string]interface{} {
	obj := map[string]interface{}{
		"Selector": []map[string]interface{}{
			{
				"_": selector.Value,
				"$": map[string]string{
					"Name": selector.Name,
				},
			},
		},
	}
	return obj
}

func createCommonBodyPull(enumerationContext string, maxElements, maxCharacters int) string {
	if maxElements == 0 {
		maxElements = 999
	}
	if maxCharacters == 0 {
		maxCharacters = 99999
	}
	return fmt.Sprintf(`<Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>%d</MaxElements><MaxCharacters>%d</MaxCharacters></Pull></Body>`, enumerationContext, maxElements, maxCharacters)
}

func (w WSManMessageCreator) createCommonBodyCreateOrPut(wsmanClass string, data interface{}) string {
	return w.CreateBody(wsmanClass, wsmanClass, data)
}

func createCommonBodyRequestStateChange(input string, requestedState int) string {
	return fmt.Sprintf(`<Body><h:RequestStateChange_INPUT xmlns:h="%s"><h:RequestedState>%d</h:RequestedState></h:RequestStateChange_INPUT></Body>`, input, requestedState)
}
