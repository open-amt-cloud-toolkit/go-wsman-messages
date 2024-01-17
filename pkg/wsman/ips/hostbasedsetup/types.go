/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbasedsetup

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
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
		Setup_OUTPUT              Setup_OUTPUT              `xml:"Setup_OUTPUT"`
		AdminSetup_OUTPUT         AdminSetup_OUTPUT         `xml:"AdminSetup_OUTPUT"`
		AddNextCertInChain_OUTPUT AddNextCertInChain_OUTPUT `xml:"AddNextCertInChain_OUTPUT"`
		IPS_HostBasedSetupService HostBasedSetupService     `xml:"IPS_HostBasedSetupService"`
	}

	HostBasedSetupService struct {
		XMLName                 xml.Name `xml:"IPS_HostBasedSetupService"`
		ElementName             string
		SystemCreationClassName string
		SystemName              string
		CreationClassName       string
		Name                    string
		CurrentControlMode      CurrentControlMode
		AllowedControlModes     AllowedControlModes
		ConfigurationNonce      string
		CertChainStatus         CertChainStatus
	}
	AddNextCertInChain_OUTPUT struct {
		XMLName     xml.Name `xml:"AddNextCertInChain_OUTPUT"`
		ReturnValue ReturnValue
	}

	AdminSetup_OUTPUT struct {
		XMLName     xml.Name `xml:"AdminSetup_OUTPUT"`
		ReturnValue ReturnValue
	}

	Setup_OUTPUT struct {
		XMLName     xml.Name `xml:"Setup_OUTPUT"`
		ReturnValue ReturnValue
	}
)

// INPUT
// Request Types
type (
	AddNextCertInChain_INPUT struct {
		XMLName           xml.Name `xml:"h:AddNextCertInChain_INPUT"`
		H                 string   `xml:"xmlns:h,attr"`
		NextCertificate   string   `xml:"h:NextCertificate"`
		IsLeafCertificate bool     `xml:"h:IsLeafCertificate"`
		IsRootCertificate bool     `xml:"h:IsRootCertificate"`
	}
	AdminSetup_INPUT struct {
		XMLName                    xml.Name `xml:"h:AdminSetup_INPUT"`
		H                          string   `xml:"xmlns:h,attr"`
		NetAdminPassEncryptionType int      `xml:"h:NetAdminPassEncryptionType"`
		NetworkAdminPassword       string   `xml:"h:NetworkAdminPassword"`
		McNonce                    string   `xml:"h:McNonce"`
		SigningAlgorithm           int      `xml:"h:SigningAlgorithm"`
		DigitalSignature           string   `xml:"h:DigitalSignature"`
	}
	Setup_INPUT struct {
		XMLName                    xml.Name `xml:"h:Setup_INPUT"`
		H                          string   `xml:"xmlns:h,attr"`
		NetAdminPassEncryptionType int      `xml:"h:NetAdminPassEncryptionType"`
		NetworkAdminPassword       string   `xml:"h:NetworkAdminPassword"`
	}
	UpgradeClientToAdmin_INPUT struct {
		XMLName          xml.Name `xml:"h:UpgradeClientToAdmin_INPUT"`
		H                string   `xml:"xmlns:h,attr"`
		McNonce          string   `xml:"h:McNonce"`
		SigningAlgorithm int      `xml:"h:SigningAlgorithm"`
		DigitalSignature string   `xml:"h:DigitalSignature"`
	}
)

type AdminPassEncryptionType int
type SigningAlgorithm int
type CurrentControlMode int
type CertChainStatus int
type AllowedControlModes int
type ReturnValue int
