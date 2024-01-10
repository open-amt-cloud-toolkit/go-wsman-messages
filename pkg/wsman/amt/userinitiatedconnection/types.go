/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Service struct {
	base message.Base
}

type RequestedState int

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
		XMLName                   xml.Name                  `xml:"Body"`
		RequestStateChange_OUTPUT RequestStateChange_OUTPUT `xml:"RequestStateChange_OUTPUT"`
		GetResponse               UserResponse
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse
	}
	UserResponse struct {
		XMLName                 xml.Name     `xml:"AMT_UserInitiatedConnectionService"`
		CreationClassName       string       `xml:"CreationClassName,omitempty"`
		ElementName             string       `xml:"ElementName,omitempty"`
		EnabledState            EnabledState `xml:"EnabledState"`
		Name                    string       `xml:"Name,omitempty"`
		SystemCreationClassName string       `xml:"SystemCreationClassName,omitempty"`
		SystemName              string       `xml:"SystemName,omitempty"`
	}
	PullResponse struct {
		XMLName   xml.Name       `xml:"PullResponse"`
		UserItems []UserResponse `xml:"Items>AMT_UserInitiatedConnectionService"`
	}
	RequestStateChange_OUTPUT struct {
		XMLName     xml.Name `xml:"RequestStateChange_OUTPUT"`
		ReturnValue int      `xml:"ReturnValue"`
	}
	EnabledState int
)
