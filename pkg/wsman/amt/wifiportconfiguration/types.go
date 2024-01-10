/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportconfiguration

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/wifi"
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
		Body    Body
	}

	Body struct {
		XMLName                      xml.Name `xml:"Body"`
		WiFiPortConfigurationService WiFiPortConfigurationServiceResponse
		PullResponse                 PullResponse
		EnumerateResponse            common.EnumerateResponse
		AddWiFiSettings_OUTPUT       AddWiFiSettings_OUTPUT
	}
	WiFiPortConfigurationServiceResponse struct {
		XMLName                            xml.Name                           `xml:"AMT_WiFiPortConfigurationService"`
		RequestedState                     models.RequestedState              `xml:"RequestedState"`
		EnabledState                       models.EnabledState                `xml:"EnabledState"`
		HealthState                        models.HealthState                 `xml:"HealthState"`
		ElementName                        string                             `xml:"ElementName,omitempty"`
		SystemCreationClassName            string                             `xml:"SystemCreationClassName,omitempty"`
		SystemName                         string                             `xml:"SystemName,omitempty"`
		CreationClassName                  string                             `xml:"CreationClassName,omitempty"`
		Name                               string                             `xml:"Name,omitempty"`
		LocalProfileSynchronizationEnabled LocalProfileSynchronizationEnabled `xml:"localProfileSynchronizationEnabled"`
		LastConnectedSsidUnderMeControl    string                             `xml:"LastConnectedSsidUnderMeControl,omitempty"`
		NoHostCsmeSoftwarePolicy           NoHostCsmeSoftwarePolicy           `xml:"NoHostCsmeSoftwarePolicy"`
		UEFIWiFiProfileShareEnabled        UEFIWiFiProfileShareEnabled        `xml:"UEFIWiFiProfileShareEnabled"`
	}
	PullResponse struct {
		XMLName                    xml.Name                               `xml:"PullResponse"`
		WiFiPortConfigurationItems []WiFiPortConfigurationServiceResponse `xml:"Items>AMT_WiFiPortConfigurationService"`
	}
	AddWiFiSettings_OUTPUT struct {
		XMLName xml.Name `xml:"AddWiFiSettings_OUTPUT"`
		// not concerned with these entries on OUTPUT
		//IEEE8021xSettings    *models.IEEE8021xSettings `xml:"g:IEEE8021xSettingsInput,omitempty"`
		//ClientCredential     *ClientCredential         `xml:"g:ClientCredential,omitempty"`
		//CACredential         *CACredential             `xml:"g:CACredential,omitempty"`
		ReturnValue int `xml:"ReturnValue"`
	}
)

type (
	LocalProfileSynchronizationEnabled int
	NoHostCsmeSoftwarePolicy           int
	UEFIWiFiProfileShareEnabled        int
)

// INPUT
// Request Types
type (
	AddWiFiSettings_INPUT struct {
		XMLName              xml.Name `xml:"h:AddWiFiSettings_INPUT"`
		H                    string   `xml:"xmlns:h,attr"`
		WifiEndpoint         WiFiEndpoint
		WiFiEndpointSettings wifi.WiFiEndpointSettings_INPUT
		IEEE8021xSettings    *models.IEEE8021xSettings `xml:"h:IEEE8021xSettingsInput,omitempty"`
		ClientCredential     *ClientCredentialRequest  `xml:"h:ClientCredential,omitempty"`
		CACredential         *CACredentialRequest      `xml:"h:CACredential,omitempty"`
	}
	WiFiPortConfigurationServiceRequest struct {
		XMLName                            xml.Name                           `xml:"h:AMT_WiFiPortConfigurationService"`
		H                                  string                             `xml:"xmlns:h,attr"`
		RequestedState                     models.RequestedState              `xml:"h:RequestedState,omitempty"`
		EnabledState                       models.EnabledState                `xml:"h:EnabledState,omitempty"`
		HealthState                        models.HealthState                 `xml:"h:HealthState,omitempty"`
		ElementName                        string                             `xml:"h:ElementName,omitempty"`
		SystemCreationClassName            string                             `xml:"h:SystemCreationClassName,omitempty"`
		SystemName                         string                             `xml:"h:SystemName,omitempty"`
		CreationClassName                  string                             `xml:"h:CreationClassName,omitempty"`
		Name                               string                             `xml:"h:Name,omitempty"`
		LocalProfileSynchronizationEnabled LocalProfileSynchronizationEnabled `xml:"h:localProfileSynchronizationEnabled"`
		LastConnectedSsidUnderMeControl    string                             `xml:"h:LastConnectedSsidUnderMeControl,omitempty"`
		NoHostCsmeSoftwarePolicy           NoHostCsmeSoftwarePolicy           `xml:"h:NoHostCsmeSoftwarePolicy,omitempty"`
		UEFIWiFiProfileShareEnabled        UEFIWiFiProfileShareEnabled        `xml:"h:UEFIWiFiProfileShareEnabled,omitempty"`
	}
	CACredentialRequest struct {
		XMLName             xml.Name            `xml:"h:CACredential,omitempty"`
		H                   string              `xml:"xmlns:a,attr"`
		Address             string              `xml:"a:Address,omitempty"`
		ReferenceParameters ReferenceParameters `xml:"a:ReferenceParameters,omitempty"`
	}
	ClientCredentialRequest struct {
		XMLName             xml.Name            `xml:"h:ClientCredential,omitempty"`
		H                   string              `xml:"xmlns:a,attr"`
		Address             string              `xml:"a:Address,omitempty"`
		ReferenceParameters ReferenceParameters `xml:"a:ReferenceParameters,omitempty"`
	}

	ReferenceParameters struct {
		XMLName     xml.Name    `xml:"a:ReferenceParameters"`
		H           string      `xml:"xmlns:c,attr"`
		ResourceURI string      `xml:"c:ResourceURI,omitempty"`
		SelectorSet SelectorSet `xml:"c:SelectorSet,omitempty"`
	}

	SelectorSet struct {
		H        string   `xml:"xmlns:c,attr"`
		XMLName  xml.Name `xml:"c:SelectorSet,omitempty"`
		Selector []Selector
	}

	Selector struct {
		H       string   `xml:"xmlns:c,attr"`
		XMLName xml.Name `xml:"c:Selector,omitempty"`
		Name    string   `xml:"Name,attr"`
		Value   string   `xml:",chardata"`
	}
)

type WiFiEndpoint struct {
	XMLName             xml.Name            `xml:"h:WiFiEndpoint,omitempty"`
	Address             string              `xml:"a:Address,omitempty"`
	ReferenceParameters ReferenceParameters `xml:"a:ReferenceParameters,omitempty"`
}

type AddWiFiSettingsResponse struct {
	XMLName                xml.Name `xml:"Body"`
	AddWiFiSettings_OUTPUT AddWiFiSettings_OUTPUT
}
