/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package software

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Identity struct {
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
		XMLName                  xml.Name `xml:"Body"`
		PullResponse             PullResponse
		EnumerateResponse        common.EnumerateResponse
		SoftwareIdentityResponse SoftwareIdentity
	}

	PullResponse struct {
		XMLName               xml.Name           `xml:"PullResponse"`
		SoftwareIdentityItems []SoftwareIdentity `xml:"Items>CIM_SoftwareIdentity"`
	}

	SoftwareIdentity struct {
		XMLName       xml.Name `xml:"CIM_SoftwareIdentity"`
		InstanceID    string   `xml:"InstanceID"`
		VersionString string   `xml:"VersionString"`
		IsEntity      bool     `xml:"IsEntity"`
	}
)
