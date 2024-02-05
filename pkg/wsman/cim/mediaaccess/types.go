/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mediaaccess

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Device struct {
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
	}

	PullResponse struct {
		XMLName            xml.Name            `xml:"PullResponse"`
		MediaAccessDevices []MediaAccessDevice `xml:"Items>CIM_MediaAccessDevice"`
	}

	// TODO: Capabilities and OperationalStatus can return multiple items with the same tag, need to handle this
	MediaAccessDevice struct {
		Capabilities            CapabilitiesValues       `xml:"Capabilities,omitempty"` // ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} Values={Unknown, Other, Sequential Access, Random Access, Supports Writing, Encryption, Compression, Supports Removeable Media, Manual Cleaning, Automatic Cleaning, SMART Notification, Supports Dual Sided Media, Predismount Eject Not Required} ArrayType=Indexed
		CreationClassName       string                   `xml:"CreationClassName"`
		DeviceID                string                   `xml:"DeviceID"`
		ElementName             string                   `xml:"ElementName"`
		EnabledDefault          models.EnabledDefault    `xml:"EnabledDefault"`
		EnabledState            models.EnabledState      `xml:"EnabledState"`
		MaxMediaSize            int                      `xml:"MaxMediaSize,omitempty"`
		OperationalStatus       models.OperationalStatus `xml:"OperationalStatus"`
		RequestedState          models.RequestedState    `xml:"RequestedState"`
		Security                SecurityValues           `xml:"Security,omitempty"` // ValueMap={1, 2, 3, 4, 5, 6, 7} Values={Other, Unknown, None, Read Only, Locked Out, Boot Bypass, Boot Bypass and Read Only}
		SystemCreationClassName string                   `xml:"SystemCreationClassName"`
		SystemName              string                   `xml:"SystemName"`
	}

	CapabilitiesValues int
	SecurityValues     int
)
