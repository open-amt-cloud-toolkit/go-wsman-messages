/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportconfiguration

import (
	"fmt"
	"html"

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

func NewWiFiPortConfigurationService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{
		base: wsman.NewBase(wsmanMessageCreator, AMT_WiFiPortConfigurationService),
	}
}
func (s Service) Get() string {
	return s.base.Get(nil)
}
func (s Service) Enumerate() string {
	return s.base.Enumerate()
}
func (s Service) Pull(enumerationContext string) string {
	return s.base.Pull(enumerationContext)
}
func (s Service) Put(wiFiPortConfigurationService WiFiPortConfigurationService) string {
	return s.base.Put(wiFiPortConfigurationService, false, nil)
}

// AddWiFiSettings atomically creates instances and associates them based on the input parameters.
func (s Service) AddWiFiSettings(wifiEndpointSettings models.WiFiEndpointSettings, selector wsman.Selector, ieee8021xSettingsInput *models.IEEE8021xSettings, clientCredential, caCredential string) string {
	header := s.base.WSManMessageCreator.CreateHeader(string(actions.AddWiFiSettings), AMT_WiFiPortConfigurationService, nil, "", "")
	//wifiEndpointSettings.PSKPassPhrase = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(wifiEndpointSettings.PSKPassPhrase, "&", "&amp;"), "<", "&lt;"), ">", "&gt;"), "\"", "&quot;"), "'", "&apos;")
	//needs testing
	wifiEndpointSettings.PSKPassPhrase = html.EscapeString(wifiEndpointSettings.PSKPassPhrase)

	dataArray := []interface{}{}
	wifiEndpointObject := map[string]interface{}{
		"WiFiEndpoint": map[string]interface{}{
			"Address": "/wsman",
			"ReferenceParameters": map[string]interface{}{
				"ResourceURI": fmt.Sprintf("http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/%s", wifi.CIM_WiFiEndpoint),
				"SelectorSet": s.base.WSManMessageCreator.CreateSelectorObjectForBody(selector),
			},
		},
	}
	dataArray = append(dataArray, wifiEndpointObject)

	wifiEndpointSettingInputObject := map[string]interface{}{
		"WiFiEndpointSettingsInput": wifiEndpointSettings,
		"namespace":                 "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings",
	}
	dataArray = append(dataArray, wifiEndpointSettingInputObject)

	if ieee8021xSettingsInput != nil {
		ieee8021xSettingsInputObject := map[string]interface{}{
			"ieee8021xSettingsInput": *ieee8021xSettingsInput,
			"namespace":              "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_IEEE8021xSettings",
		}
		dataArray = append(dataArray, ieee8021xSettingsInputObject)
	}

	if clientCredential != "" {
		clientCredentialObject := map[string]interface{}{
			"ClientCredential": clientCredential,
		}
		dataArray = append(dataArray, clientCredentialObject)
	}

	if caCredential != "" {
		caCredentialObject := map[string]interface{}{
			"CACredential": caCredential,
		}
		dataArray = append(dataArray, caCredentialObject)
	}

	body := s.base.WSManMessageCreator.CreateBody(string(methods.AddWiFiSettings)+"_INPUT", AMT_WiFiPortConfigurationService, dataArray)
	return s.base.WSManMessageCreator.CreateXML(header, body)
}
