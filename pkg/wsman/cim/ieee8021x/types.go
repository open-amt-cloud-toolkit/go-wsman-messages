/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips/ieee8021x"
)

type Settings struct {
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
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse `xml:"PullResponse"`
		GetResponse       IEEE8021xSettingsResponse
	}

	IEEE8021xSettingsResponse ieee8021x.IEEE8021xSettingsResponse // calls return IPS version of IEEE8021xSettings

	Time struct {
		DateTime string `xml:"Datetime"`
	}

	PullResponse struct {
		XMLName                xml.Name                    `xml:"PullResponse"`
		IEEE8021xSettingsItems []IEEE8021xSettingsResponse `xml:"Items>IPS_IEEE8021xSettings"`
	}
)

type (
	IEEE8021xSettingsRequest struct {
		H             string   `xml:"xmlns:w,attr"`
		XMLName       xml.Name `xml:"CIM_IEEE8021xSettings"`
		ElementName   string   `xml:"w:ElementName"`
		InstanceID    string   `xml:"w:InstanceID"`
		Enabled       int      `xml:"w:Enabled"`
		AvailableInS0 bool     `xml:"w:AvailableInS0"`
		PxeTimeout    int      `xml:"w:PxeTimeout"`
	}
)
