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
)

// Package Types
type (
	Settings struct {
		base message.Base
	}
	CredentialContext struct {
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
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
		IEEE8021xSettings IEEE8021xSettings
	}

	IEEE8021xSettingsResponse struct {
		XMLName       xml.Name `xml:"IPS_IEEE8021xSettings"`
		ElementName   string   `xml:"ElementName"`
		InstanceID    string   `xml:"InstanceID"`
		Enabled       int      `xml:"Enabled"`
		AvailableInS0 bool     `xml:"AvailableInS0"`
		PxeTimeout    int      `xml:"PxeTimeout"`
	}

	PullResponse struct {
		Items []IEEE8021xSettings `xml:"Items>IPS_IEEE8021xSetings"`
	}
)

// INPUT
// Request Types
type (
	IEEE8021xSettings struct {
		XMLName                         xml.Name                        `xml:"h:IEEE8021xSettingsInput,omitempty"`
		H                               string                          `xml:"xmlns:q,attr"`
		ElementName                     string                          `xml:"q:ElementName,omitempty"`
		InstanceID                      string                          `xml:"q:InstanceID,omitempty"`
		AuthenticationProtocol          AuthenticationProtocol          `xml:"q:AuthenticationProtocol"`
		RoamingIdentity                 string                          `xml:"q:RoamingIdentity,omitempty"`
		ServerCertificateName           string                          `xml:"q:ServerCertificateName,omitempty"`
		ServerCertificateNameComparison ServerCertificateNameComparison `xml:"q:ServerCertificateNameComparison,omitempty"`
		Username                        string                          `xml:"q:Username,omitempty"`
		Password                        string                          `xml:"q:Password,omitempty"`
		Domain                          string                          `xml:"q:Domain,omitempty"`
		ProtectedAccessCredential       string                          `xml:"q:ProtectedAccessCredential,omitempty"`
		PACPassword                     string                          `xml:"q:PACPassword,omitempty"`
		PSK                             string                          `xml:"q:PSK,omitempty"`
		Enabled                         IEEE8021xSettingsEnabled        `json:"Enabled,omitempty"`
		PxeTimeout                      int                             `json:"PxeTimeout,omitempty"`
		AvailableInS0                   bool                            `json:"AvailableInS0,omitempty"`
	}
	Certificate struct {
		XMLName                 xml.Name `xml:"h:SetCertificates_INPUT"`
		H                       string   `xml:"xmlns:h,attr"`
		ServerCertificateIssuer string   `xml:"h:ServerCertificateIssuer"`
		ClientCertificate       string   `xml:"h:ClientCertificate"`
	}
)

// ServerCertificateNameComparison represents the ServerCertificateNameComparison type for IEEE8021xProfile.
type ServerCertificateNameComparison int
type IEEE8021xSettingsEnabled int
type AuthenticationProtocol int
