/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportconfiguration

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/wifi"
)

const AMT_WiFiPortConfigurationService = "AMT_WiFiPortConfigurationService"

type WiFiPortConfigurationService struct {
	models.NetworkPortConfigurationService
	RequestedState                     RequestedState
	EnabledState                       EnabledState
	HealthState                        HealthState
	ElementName                        string
	SystemCreationClassName            string
	SystemName                         string
	CreationClassName                  string
	Name                               string
	LocalProfileSynchronizationEnabled LocalProfileSynchronizationEnabled
	LastConnectedSsidUnderMeControl    string
	NoHostCsmeSoftwarePolicy           NoHostCsmeSoftwarePolicy
	UEFIWiFiProfileShareEnabled        UEFIWiFiProfileShareEnabled
}

type RequestedState int

const (
	RequestedStateUnknown       RequestedState = 0
	RequestedStateEnabled       RequestedState = 2
	RequestedStateDisabled      RequestedState = 3
	RequestedStateShutDown      RequestedState = 4
	RequestedStateNoChange      RequestedState = 5
	RequestedStateOffline       RequestedState = 6
	RequestedStateTest          RequestedState = 7
	RequestedStateDeferred      RequestedState = 8
	RequestedStateQuiesce       RequestedState = 9
	RequestedStateReboot        RequestedState = 10
	RequestedStateReset         RequestedState = 11
	RequestedStateNotApplicable RequestedState = 12
)

type EnabledState int

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)

type HealthState int

const (
	Unknown             HealthState = 0
	OK                  HealthState = 5
	DegradedWarning     HealthState = 10
	MinorFailure        HealthState = 15
	MajorFailure        HealthState = 20
	CriticalFailure     HealthState = 25
	NonRecoverableError HealthState = 30
)

type LocalProfileSynchronizationEnabled int

const (
	LocalSyncDisabled LocalProfileSynchronizationEnabled = 0
	UnrestrictedSync  LocalProfileSynchronizationEnabled = 3
)

type NoHostCsmeSoftwarePolicy int

const (
	RelaxedPolicy NoHostCsmeSoftwarePolicy = iota
	AggressivePolicy
	Reserved
)

type UEFIWiFiProfileShareEnabled int

const (
	Enabled UEFIWiFiProfileShareEnabled = iota
	Disabled
)

type Service struct {
	base wsman.Base
}
type AddWiFiSettings_INPUT struct {
	XMLName              xml.Name `xml:"h:AddWiFiSettings_INPUT"`
	H                    string   `xml:"xmlns:h,attr"`
	WifiEndpoint         WiFiEndpoint
	WiFiEndpointSettings models.WiFiEndpointSettings
	IEEE8021xSettings    *models.IEEE8021xSettings `xml:"h:IEEE8021xSettingsInput,omitempty"`
	ClientCredential     *ClientCredential         `xml:"h:ClientCredential,omitempty"`
	CACredential         *CACredential             `xml:"h:CACredential,omitempty"`
}
type WiFiEndpoint struct {
	XMLName             xml.Name                   `xml:"h:WiFiEndpoint,omitempty"`
	Address             string                     `xml:"a:Address,omitempty"`
	ReferenceParameters models.ReferenceParameters `xml:"a:ReferenceParameters,omitempty"`
}
type CACredential struct {
	XMLName             xml.Name                   `xml:"h:CACredential,omitempty"`
	Address             string                     `xml:"a:Address,omitempty"`
	ReferenceParameters models.ReferenceParameters `xml:"a:ReferenceParameters,omitempty"`
}
type ClientCredential struct {
	XMLName             xml.Name                   `xml:"h:ClientCredential,omitempty"`
	Address             string                     `xml:"a:Address,omitempty"`
	ReferenceParameters models.ReferenceParameters `xml:"a:ReferenceParameters,omitempty"`
}

func NewWiFiPortConfigurationService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{
		base: wsman.NewBase(wsmanMessageCreator, AMT_WiFiPortConfigurationService),
	}
}

// Get retrieves the representation of the instance
func (s Service) Get() string {
	return s.base.Get(nil)
}

// Enumerates the instances of this class
func (s Service) Enumerate() string {
	return s.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (s Service) Pull(enumerationContext string) string {
	return s.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (s Service) Put(wiFiPortConfigurationService WiFiPortConfigurationService) string {
	return s.base.Put(wiFiPortConfigurationService, false, nil)
}

// AddWiFiSettings atomically creates instances and associates them based on the input parameters.
func (s Service) AddWiFiSettings(wifiEndpointSettings models.WiFiEndpointSettings, ieee8021xSettingsInput *models.IEEE8021xSettings, wifiEndpoint, clientCredential, caCredential string) string {
	header := s.base.WSManMessageCreator.CreateHeader(string(actions.AddWiFiSettings), AMT_WiFiPortConfigurationService, nil, "", "")

	input := AddWiFiSettings_INPUT{
		WifiEndpoint: WiFiEndpoint{
			Address: "/wsman",
			ReferenceParameters: models.ReferenceParameters{
				ResourceURI: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/" + wifi.CIM_WiFiEndpoint,
				SelectorSet: models.SelectorSet{
					Selector: []wsman.Selector{
						{
							Name:  "Name",
							Value: wifiEndpoint,
						},
					},
				},
			},
		},
		WiFiEndpointSettings: wifiEndpointSettings,
	}
	input.WiFiEndpointSettings.H = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings"
	if ieee8021xSettingsInput != nil {
		input.IEEE8021xSettings = ieee8021xSettingsInput
		input.IEEE8021xSettings.H = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_IEEE8021xSettings"
		input.CACredential = &CACredential{
			Address: "default",
			ReferenceParameters: models.ReferenceParameters{
				ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
				SelectorSet: models.SelectorSet{
					Selector: []wsman.Selector{
						{
							Name:  "InstanceID",
							Value: caCredential,
						},
					},
				},
			},
		}
		input.ClientCredential = &ClientCredential{
			Address: "default",
			ReferenceParameters: models.ReferenceParameters{
				ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
				SelectorSet: models.SelectorSet{
					Selector: []wsman.Selector{
						{
							Name:  "InstanceID",
							Value: clientCredential,
						},
					},
				},
			},
		}
	}

	body := s.base.WSManMessageCreator.CreateBody(string(methods.AddWiFiSettings)+"_INPUT", AMT_WiFiPortConfigurationService, &input)
	return s.base.WSManMessageCreator.CreateXML(header, body)
}
