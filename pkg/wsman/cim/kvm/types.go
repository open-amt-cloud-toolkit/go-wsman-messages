/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kvm

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type RedirectionSAP struct {
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
		XMLName           xml.Name          `xml:"Body"`
		GetResponse       KVMRedirectionSAP `xml:"CIM_KVMRedirectionSAP"`
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse `xml:"PullResponse"`
	}

	KVMRedirectionSAP struct {
		CreationClassName       string                       `xml:"CreationClassName"`
		ElementName             string                       `xml:"ElementName"`
		Name                    string                       `xml:"Name"`
		SystemCreationClassName string                       `xml:"SystemCreationClassName"`
		SystemName              string                       `xml:"SystemName"`
		EnabledState            models.EnabledState          `xml:"EnabledState,omitempty"`
		RequestedState          models.RequestedState        `xml:"RequestedState,omitempty"`
		KVMProtocol             KVMRedirectionSAPKVMProtocol `xml:"KVMProtocol,omitempty"`
	}

	Time struct {
		DateTime string `xml:"Datetime"`
	}

	PullResponse struct {
		XMLName xml.Name            `xml:"PullResponse"`
		Items   []KVMRedirectionSAP `xml:"Items>CIM_KVMRedirectionSAP"`
	}
)

type KVMRedirectionSAPRequestedStateInputs int

type KVMRedirectionSAPKVMProtocol int
