/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package chassis

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Package struct {
	base   message.Base
	client client.WSMan
}

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
		XMLName      xml.Name          `xml:"PullResponse"`
		PackageItems []PackageResponse `xml:"Items>CIM_Chassis"`
	}

	PackageResponse struct {
		XMLName            xml.Name                  `xml:"CIM_Chassis"`
		Version            string                    `xml:"Version"`
		SerialNumber       string                    `xml:"SerialNumber"`
		Model              string                    `xml:"Model"`
		Manufacturer       string                    `xml:"Manufacturer"`
		ElementName        string                    `xml:"ElementName"`
		CreationClassName  string                    `xml:"CreationClassName"`
		Tag                string                    `xml:"Tag"`
		OperationalStatus  models.OperationalStatus  `xml:"OperationalStatus"`
		PackageType        models.PackageType        `xml:"PackageType"`
		ChassisPackageType models.ChassisPackageType `xml:"ChassisPackageType"`
	}
)
