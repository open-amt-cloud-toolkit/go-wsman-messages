/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package bios

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Element struct {
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
		GetResponse       BiosElement
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	BiosElement struct {
		XMLName               xml.Name                 `xml:"CIM_BIOSElement"`
		TargetOperatingSystem int                      `xml:"TargetOperatingSystem"`
		SoftwareElementID     string                   `xml:"SoftwareElementID"`
		SoftwareElementState  int                      `xml:"SoftwareElementState"`
		Name                  string                   `xml:"Name"`
		OperationalStatus     models.OperationalStatus `xml:"OperationalStatus"`
		ElementName           string                   `xml:"ElementName"`
		Version               string                   `xml:"Version"`
		Manufacturer          string                   `xml:"Manufacturer"`
		PrimaryBIOS           bool                     `xml:"PrimaryBIOS"`
		ReleaseDate           Time                     `xml:"ReleaseDate"`
	}

	Time struct {
		DateTime string `xml:"Datetime"`
	}

	PullResponse struct {
		XMLName          xml.Name      `xml:"PullResponse"`
		BiosElementItems []BiosElement `xml:"Items>CIM_BIOSElement"`
	}
)
