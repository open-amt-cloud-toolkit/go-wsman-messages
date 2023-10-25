/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package message

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
)

type Base struct {
	WSManMessageCreator *WSManMessageCreator
	className           string
	client              wsman.WSManClient
}

type Header struct {
	XMLName     xml.Name `xml:"Header"`
	To          string   `xml:"To"`
	RelatesTo   int      `xml:"RelatesTo"`
	Action      Action   `xml:"Action"`
	MessageID   string   `xml:"MessageID"`
	ResourceURI string   `xml:"ResourceURI"`
}
type Action struct {
	XMLName        xml.Name `xml:"Action"`
	MustUnderstand string   `xml:"mustUnderstand,attr"`
	Value          string   `xml:",chardata"`
}

type EnumerationResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  Header
	Body    EnumerationBody
}

type EnumerationBody struct {
	EnumerateResponse EnumerateResponse
}

type EnumerateResponse struct {
	EnumerationContext string
}

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
