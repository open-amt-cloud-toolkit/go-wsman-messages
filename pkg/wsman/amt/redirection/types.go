/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package redirection

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Service struct {
	base message.Base
}

// OUTPUT
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
		GetAndPutResponse         RedirectionResponse       `xml:"AMT_RedirectionService"`
		RequestStateChange_OUTPUT RequestStateChange_OUTPUT `xml:"RequestStateChange_OUTPUT"`
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse
	}
	RedirectionResponse struct {
		XMLName                 xml.Name     `xml:"AMT_RedirectionService"`
		CreationClassName       string       `xml:"CreationClassName,omitempty"`
		ElementName             string       `xml:"ElementName,omitempty"`
		EnabledState            EnabledState `xml:"EnabledState"`
		ListenerEnabled         bool         `xml:"ListenerEnabled"`
		Name                    string       `xml:"Name,omitempty"`
		SystemCreationClassName string       `xml:"SystemCreationClassName,omitempty"`
		SystemName              string       `xml:"SystemName,omitempty"`
	}
	PullResponse struct {
		XMLName          xml.Name              `xml:"PullResponse"`
		RedirectionItems []RedirectionResponse `xml:"Items>AMT_RedirectionService"`
	}
	RequestStateChange_OUTPUT struct {
		XMLName     xml.Name `xml:"RequestStateChange_OUTPUT"`
		ReturnValue int      `xml:"ReturnValue"`
	}
)

// INPUT
// Request Types
type (
	RedirectionRequest struct {
		XMLName                 xml.Name     `xml:"h:AMT_RedirectionService"`
		H                       string       `xml:"xmlns:h,attr"`
		CreationClassName       string       `xml:"h:CreationClassName,omitempty"`
		ElementName             string       `xml:"h:ElementName,omitempty"`
		EnabledState            EnabledState `xml:"h:EnabledState"`
		ListenerEnabled         bool         `xml:"h:ListenerEnabled"`
		Name                    string       `xml:"h:Name,omitempty"`
		SystemCreationClassName string       `xml:"h:SystemCreationClassName,omitempty"`
		SystemName              string       `xml:"h:SystemName,omitempty"`
	}
)

type EnabledState int
type RequestedState int
