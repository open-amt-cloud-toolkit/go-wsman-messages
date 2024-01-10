/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mps

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type UsernamePassword struct {
	base message.Base
}

// OUTPUTS
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
		GetResponse       MPSUsernamePasswordResponse
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	MPSUsernamePasswordResponse struct {
		XMLName    xml.Name `xml:"AMT_MPSUsernamePassword"`
		InstanceID string   `xml:"InstanceID,omitempty"`
		RemoteID   string   `xml:"RemoteID,omitempty"`
		Secret     string   `xml:"Secret,omitempty"`
		Algorithm  string   `xml:"Algorithm,omitempty"`
		Protocol   string   `xml:"Protocol,omitempty"`
	}

	PullResponse struct {
		XMLName                  xml.Name                      `xml:"PullResponse"`
		MPSUsernamePasswordItems []MPSUsernamePasswordResponse `xml:"Items>AMT_MPSUsernamePassword"`
	}
)

// INPUTS
// Request Types
type (
	MPSUsernamePasswordRequest struct {
		XMLName    xml.Name `xml:"h:AMT_MPSUsernamePassword"`
		H          string   `xml:"xmlns:h,attr"`
		InstanceID string   `xml:"h:InstanceID,omitempty"`
		RemoteID   string   `xml:"h:RemoteID,omitempty"`
		Secret     string   `xml:"h:Secret,omitempty"`
		Algorithm  string   `xml:"h:Algorithm,omitempty"`
		Protocol   string   `xml:"h:Protocol,omitempty"`
	}
)
