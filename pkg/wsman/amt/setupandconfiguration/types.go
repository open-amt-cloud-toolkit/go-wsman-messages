/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Service struct {
	base message.Base
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
		XMLName                xml.Name `xml:"Body"`
		GetResponse            SetupAndConfigurationServiceResponse
		EnumerateResponse      common.EnumerateResponse
		PullResponse           PullResponse
		GetUuid_OUTPUT         GetUuid_OUTPUT         `xml:"GetUuid_OUTPUT"`
		Unprovision_OUTPUT     Unprovision_OUTPUT     `xml:"Unprovision_OUTPUT"`
		CommitChanges_OUTPUT   CommitChanges_OUTPUT   `xml:"CommitChanges_OUTPUT"`
		SetMEBxPassword_OUTPUT SetMEBxPassword_OUTPUT `xml:"SetMEBxPassword_OUTPUT"`
	}

	SetupAndConfigurationServiceResponse struct {
		XMLName                       xml.Name               `xml:"AMT_SetupAndConfigurationService"`
		RequestedState                RequestedState         `xml:"RequestedState,omitempty"`
		EnabledState                  EnabledState           `xml:"EnabledState,omitempty"`
		ElementName                   string                 `xml:"ElementName,omitempty"`
		SystemCreationClassName       string                 `xml:"SystemCreationClassName,omitempty"`
		SystemName                    string                 `xml:"SystemName,omitempty"`
		CreationClassName             string                 `xml:"CreationClassName,omitempty"`
		Name                          string                 `xml:"Name,omitempty"`
		ProvisioningMode              ProvisioningModeValue  `xml:"ProvisioningMode,omitempty"`
		ProvisioningState             ProvisioningStateValue `xml:"ProvisioningState,omitempty"`
		ZeroTouchConfigurationEnabled bool                   `xml:"ZeroTouchConfigurationEnabled,omitempty"`
		ProvisioningServerOTP         string                 `xml:"ProvisioningServerOTP,omitempty"`
		ConfigurationServerFQDN       string                 `xml:"ConfigurationServerFQDN,omitempty"`
		PasswordModel                 int                    `xml:"PasswordModel,omitempty"`
		DhcpDNSSuffix                 string                 `xml:"DhcpDNSSuffix,omitempty"`
		TrustedDNSSuffix              string                 `xml:"TrustedDNSSuffix,omitempty"`
	}
	PullResponse struct {
		XMLName                           xml.Name                               `xml:"PullResponse"`
		SetupAndConfigurationServiceItems []SetupAndConfigurationServiceResponse `xml:"Items>AMT_SetupAndConfigurationService"`
	}

	GetUuid_OUTPUT struct {
		XMLName xml.Name `xml:"GetUuid_OUTPUT"`
		UUID    string   `xml:"UUID"`
	}

	Unprovision_OUTPUT struct {
		XMLName     xml.Name `xml:"Unprovision_OUTPUT"`
		ReturnValue int
	}

	CommitChanges_OUTPUT struct {
		XMLName     xml.Name `xml:"CommitChanges_OUTPUT"`
		ReturnValue int
	}

	SetMEBxPassword_OUTPUT struct {
		XMLName     xml.Name `xml:"SetMEBxPassword_OUTPUT"`
		ReturnValue int
	}
)

// Request Types
type (
	SetupAndConfigurationServiceRequest struct {
		XMLName                       xml.Name               `xml:"h:AMT_SetupAndConfigurationService"`
		H                             string                 `xml:"xmlns:h,attr"`
		RequestedState                RequestedState         `xml:"h:RequestedState,omitempty"`
		EnabledState                  EnabledState           `xml:"h:EnabledState,omitempty"`
		ElementName                   string                 `xml:"h:ElementName,omitempty"`
		SystemCreationClassName       string                 `xml:"h:SystemCreationClassName,omitempty"`
		SystemName                    string                 `xml:"h:SystemName,omitempty"`
		CreationClassName             string                 `xml:"h:CreationClassName,omitempty"`
		Name                          string                 `xml:"h:Name,omitempty"`
		ProvisioningMode              ProvisioningModeValue  `xml:"h:ProvisioningMode,omitempty"`
		ProvisioningState             ProvisioningStateValue `xml:"h:ProvisioningState,omitempty"`
		ZeroTouchConfigurationEnabled bool                   `xml:"h:ZeroTouchConfigurationEnabled,omitempty"`
		ProvisioningServerOTP         string                 `xml:"h:ProvisioningServerOTP,omitempty"`
		ConfigurationServerFQDN       string                 `xml:"h:ConfigurationServerFQDN,omitempty"`
		PasswordModel                 int                    `xml:"h:PasswordModel,omitempty"`
		DhcpDNSSuffix                 string                 `xml:"h:DhcpDNSSuffix,omitempty"`
		TrustedDNSSuffix              string                 `xml:"h:TrustedDNSSuffix,omitempty"`
	}

	MEBXPassword struct {
		XMLName  xml.Name `xml:"h:SetMEBxPassword_INPUT"`
		H        string   `xml:"xmlns:h,attr"`
		Password string   `xml:"h:Password,omitempty"`
	}

	ProvisioningMode struct {
		XMLName          xml.Name `xml:"h:Unprovision_INPUT"`
		H                string   `xml:"xmlns:h,attr"`
		ProvisioningMode int      `xml:"h:ProvisioningMode,omitempty"`
	}

	EnabledState           int
	RequestedState         int
	ProvisioningModeValue  int
	ProvisioningStateValue int
)
