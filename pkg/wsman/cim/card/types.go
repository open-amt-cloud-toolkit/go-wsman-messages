/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package card

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Package struct {
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
		XMLName           xml.Name `xml:"Body"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
		PackageResponse   PackageResponse
	}

	PullResponse struct {
		XMLName   xml.Name          `xml:"PullResponse"`
		CardItems []PackageResponse `xml:"Items>CIM_Card"`
	}

	PackageResponse struct {
		XMLName           xml.Name                 `xml:"CIM_Card"`
		CanBeFRUed        bool                     `xml:"CanBeFRUed"`
		CreationClassName string                   `xml:"CreationClassName"`
		ElementName       string                   `xml:"ElementName"`
		Manufacturer      string                   `xml:"Manufacturer"`
		Model             string                   `xml:"Model"`
		OperationalStatus models.OperationalStatus `xml:"OperationalStatus"`
		PackageType       models.PackageType       `xml:"PackageType"`
		SerialNumber      string                   `xml:"SerialNumber"`
		Tag               string                   `xml:"Tag"`
		Version           string                   `xml:"Version"`
	}
)
