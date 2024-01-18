/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/card"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Memory struct {
	base   message.Base
	client client.WSMan
}

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
		MemoryResponse    PhysicalMemory
	}

	PullResponse struct {
		XMLName xml.Name `xml:"PullResponse"`
		common.EnumerateResponse
		MemoryItems     []PhysicalMemory       `xml:"Items>CIM_PhysicalMemory"`
		PhysicalPackage []card.PackageResponse `xml:"Items>CIM_Card"` // This might need to be fixed if we get more than just CIM_Card back on a Pull call
	}
	PhysicalMemory struct {
		XMLName                    xml.Name                 `xml:"CIM_PhysicalMemory"`
		PartNumber                 string                   `xml:"PartNumber"`
		SerialNumber               string                   `xml:"SerialNumber"`
		Manufacturer               string                   `xml:"Manufacturer"`
		ElementName                string                   `xml:"ElementName"`
		CreationClassName          string                   `xml:"CreationClassName"`
		Tag                        string                   `xml:"Tag"`
		OperationalStatus          models.OperationalStatus `xml:"OperationalStatus"`
		FormFactor                 int                      `xml:"FormFactor,omitempty"`
		MemoryType                 MemoryType               `xml:"MemoryType,omitempty"`
		Speed                      int                      `xml:"Speed,omitempty"`
		Capacity                   int                      `xml:"Capacity,omitempty"`
		BankLabel                  string                   `xml:"BankLabel,omitempty"`
		ConfiguredMemoryClockSpeed int                      `xml:"ConfiguredMemoryClockSpeed,omitempty"`
		IsSpeedInMhz               bool                     `xml:"IsSpeedInMhz,omitempty"`
		MaxMemorySpeed             int                      `xml:"MaxMemorySpeed,omitempty"`
	}

	MemoryType int
)
