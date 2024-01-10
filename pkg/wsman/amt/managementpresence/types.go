/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package managementpresence

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

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
		GetResponse       ManagementRemoteResponse
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	PullResponse struct {
		XMLName               xml.Name                   `xml:"PullResponse"`
		ManagementRemoteItems []ManagementRemoteResponse `xml:"Items>AMT_ManagementPresenceRemoteSAP"`
	}

	ManagementRemoteResponse struct {
		XMLName                 xml.Name   `xml:"AMT_ManagementPresenceRemoteSAP"`
		AccessInfo              string     `xml:"AccessInfo,omitempty"`
		CN                      string     `xml:"CN,omitempty"`
		CreationClassName       string     `xml:"CreationClassName,omitempty"`
		ElementName             string     `xml:"ElementName,omitempty"`
		InfoFormat              InfoFormat `xml:"InfoFormat,omitempty"`
		Name                    string     `xml:"Name,omitempty"`
		Port                    int        `xml:"Port,omitempty"`
		SystemCreationClassName string     `xml:"SystemCreationClassName,omitempty"`
		SystemName              string     `xml:"SystemName,omitempty"`
	}

	InfoFormat int
)
