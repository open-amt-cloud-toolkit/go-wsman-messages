/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type (
	SettingData struct {
		base message.Base
	}
	CredentialContext struct {
		base message.Base
	}
	Collection struct {
		base message.Base
	}
)

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
		XMLName                               xml.Name                           `xml:"Body"`
		SettingDataGetAndPutResponse          SettingDataResponse                `xml:"AMT_TLSSettingData"`
		CredentialContextGetResponse          CredentialContextResponse          `xml:"AMT_TLSCredentialContext"`
		ProtocolEndpointCollectionGetResponse ProtocolEndpointCollectionResponse `xml:"AMT_TLSProtocolEndpointCollection"`
		EnumerateResponse                     common.EnumerateResponse
		PullResponse                          PullResponse
	}
	SettingDataResponse struct {
		XMLName                    xml.Name `xml:"AMT_TLSSettingData"`
		AcceptNonSecureConnections bool     `xml:"AcceptNonSecureConnections"`
		ElementName                string   `xml:"ElementName"`
		Enabled                    bool     `xml:"Enabled"`
		InstanceID                 string   `xml:"InstanceID"`
		MutualAuthentication       bool     `xml:"MutualAuthentication"`
	}
	CredentialContextResponse struct {
		XMLName xml.Name `xml:"AMT_TLSCredentialContext"`
	}
	ProtocolEndpointCollectionResponse struct {
		XMLName     xml.Name `xml:"AMT_TLSProtocolEndpointCollection"`
		ElementName string   `xml:"ElementName"`
	}
	PullResponse struct {
		XMLName                         xml.Name                             `xml:"PullResponse"`
		SettingDataItems                []SettingDataResponse                `xml:"Items>AMT_TLSSettingData"`
		ProtocolEndpointCollectionItems []ProtocolEndpointCollectionResponse `xml:"Items>AMT_TLSProtocolEndpointCollection"`
		CredentialContextItems          []CredentialContextResponse          `xml:"Items>AMT_TLSCredentialContext"`
	}
)

type (
	SettingDataRequest struct {
		XMLName                       xml.Name `xml:"h:AMT_TLSSettingData"`
		H                             string   `xml:"xmlns:h,attr"`
		ElementName                   string   `xml:"h:ElementName,omitempty"`
		InstanceID                    string   `xml:"h:InstanceID,omitempty"`
		MutualAuthentication          bool     `xml:"h:MutualAuthentication"`
		Enabled                       bool     `xml:"h:Enabled"`
		TrustedCN                     []string `xml:"h:TrustedCN,omitempty"`
		AcceptNonSecureConnections    bool     `xml:"h:AcceptNonSecureConnections"`
		NonSecureConnectionsSupported bool     `xml:"h:NonSecureConnectionsSupported"`
	}
)
