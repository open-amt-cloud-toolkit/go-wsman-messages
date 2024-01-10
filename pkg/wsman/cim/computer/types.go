/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package computer

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type SystemPackage struct {
	base   message.Base
	client client.WSMan
}

// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name              `xml:"Body"`
		GetResponse       ComputerSystemPackage `xml:"CIM_ComputerSystemPackage"`
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse `xml:"PullResponse"`
	}

	PullResponse struct {
		Items []ComputerSystemPackage `xml:"Items>CIM_ComputerSystemPackage"`
	}

	Antecedent struct {
		XMLName             xml.Name `xml:"Antecedent,omitempty"`
		Address             string   `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParameters
	}
	Dependent struct {
		XMLName             xml.Name `xml:"Dependent,omitempty"`
		Address             string   `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParameters
	}
	ComputerSystemPackage struct {
		Antecedent   Antecedent
		Dependent    Dependent
		PlatformGUID string `xml:"PlatformGUID,omitempty"`
	}
	ReferenceParameters struct {
		XMLName     xml.Name    `xml:"ReferenceParameters"`
		ResourceURI string      `xml:"ResourceURI,omitempty"`
		SelectorSet SelectorSet `xml:"SelectorSet,omitempty"`
	}
	SelectorSet struct {
		XMLName  xml.Name `xml:"SelectorSet,omitempty"`
		Selector []Selector
	}
	Selector struct {
		XMLName xml.Name `xml:"Selector,omitempty"`
		Name    string   `xml:"Name,attr"`
		Value   string   `xml:",chardata"`
	}
)
