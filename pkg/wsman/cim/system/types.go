/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package system

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Package struct {
	base message.Base
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
		XMLName           xml.Name `xml:"Body"`
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}
	PullResponse struct {
		XMLName            xml.Name        `xml:"PullResponse"`
		SystemPackageItems []SystemPackage `xml:"Items>CIM_ComputerSystemPackage"`
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
	SystemPackage struct {
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
