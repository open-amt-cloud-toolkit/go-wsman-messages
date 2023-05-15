/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package environmentdetection

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

const AMT_EnvironmentDetectionSettingData = "AMT_EnvironmentDetectionSettingData"

type EnvironmentDetectionSettingData struct {
	models.SettingData
	DetectionAlgorithm         DetectionAlgorithm
	DetectionStrings           []string
	DetectionIPv6LocalPrefixes []string
}

type DetectionAlgorithm uint8

const (
	LocalDomains DetectionAlgorithm = iota
	RemoteURLs
)

type SettingData struct {
	base wsman.Base
}

func NewEnvironmentDetectionSettingData(wsmanMessageCreator *wsman.WSManMessageCreator) SettingData {
	return SettingData{
		base: wsman.NewBase(wsmanMessageCreator, AMT_EnvironmentDetectionSettingData),
	}
}

// Get retrieves the representation of the instance
func (EnvironmentDetectionSettingData SettingData) Get() string {
	return EnvironmentDetectionSettingData.base.Get(nil)
}

// Enumerates the instances of this class
func (EnvironmentDetectionSettingData SettingData) Enumerate() string {
	return EnvironmentDetectionSettingData.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (EnvironmentDetectionSettingData SettingData) Pull(enumerationContext string) string {
	return EnvironmentDetectionSettingData.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (EnvironmentDetectionSettingData SettingData) Put(environmentDetectionSettingData EnvironmentDetectionSettingData) string {
	return EnvironmentDetectionSettingData.base.Put(environmentDetectionSettingData, false, nil)
}
