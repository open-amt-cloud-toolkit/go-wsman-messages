/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kerberos

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type SettingData struct {
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
		XMLName                        xml.Name `xml:"Body"`
		GetResponse                    KerberosSettingDataResponse
		EnumerateResponse              common.EnumerateResponse
		PullResponse                   PullResponse
		GetCredentialCacheState_OUTPUT GetCredentialCacheState_OUTPUT
	}

	PullResponse struct {
		XMLName                  xml.Name                      `xml:"PullResponse"`
		KerberosSettingDataItems []KerberosSettingDataResponse `xml:"Items>AMT_KerberosSettingData"`
	}
	KerberosSettingDataResponse struct {
		XMLName                        xml.Name                         `xml:"AMT_KerberosSettingData"`
		ElementName                    string                           `xml:"ElementName,omitempty"`
		InstanceID                     string                           `xml:"InstanceID,omitempty"`
		RealmName                      string                           `xml:"RealmName,omitempty"`
		ServicePrincipalName           []string                         `xml:"ServicePrincipalName,omitempty"`
		ServicePrincipalProtocol       []ServicePrincipalProtocol       `xml:"ServicePrincipalProtocol"`
		KeyVersion                     int                              `xml:"KeyVersion,omitempty"`
		EncryptionAlgorithm            EncryptionAlgorithm              `xml:"EncryptionAlgorithm,omitempty"`
		MasterKey                      []int                            `xml:"MasterKey"`
		MaximumClockTolerance          int                              `xml:"MaximumClockTolerance,omitempty"`
		KrbEnabled                     bool                             `xml:"KrbEnabled"`
		Passphrase                     string                           `xml:"Passphrase,omitempty"`
		Salt                           string                           `xml:"Salt,omitempty"`
		IterationCount                 int                              `xml:"IterationCount,omitempty"`
		SupportedEncryptionAlgorithms  []SupportedEncryptionAlgorithms  `xml:"SupportedEncryptionAlgorithms"`
		ConfiguredEncryptionAlgorithms []ConfiguredEncryptionAlgorithms `xml:"ConfiguredEncryptionAlgorithms"`
	}
	GetCredentialCacheState_OUTPUT struct {
		XMLName     xml.Name `xml:"GetCredentialCacheState_OUTPUT"`
		Enabled     bool     `xml:"Enabled"`
		ReturnValue int      `xml:"ReturnValue"`
	}
)

// INPUTS
// Request Types
type (
	SetCredentialCacheState_INPUT struct {
		XMLName xml.Name `xml:"h:SetCredentialCacheState_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Enabled bool     `xml:"h:Enabled"`
	}
)

type ServicePrincipalProtocol int
type SupportedEncryptionAlgorithms int
type ConfiguredEncryptionAlgorithms int
type EncryptionAlgorithm int
