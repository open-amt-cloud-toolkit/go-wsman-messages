/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

const AMT_TLSSettingData = "AMT_TLSSettingData"

type TLSSettingData struct {
	models.SettingData
	MutualAuthentication          bool
	Enabled                       bool
	TrustedCN                     string
	AcceptNonSecureConnections    bool
	NonSecureConnectionsSupported bool
}
type SettingData struct {
	base wsman.Base
}

func NewTLSSettingData(wsmanMessageCreator *wsman.WSManMessageCreator) SettingData {
	return SettingData{
		base: wsman.NewBase(wsmanMessageCreator, AMT_TLSSettingData),
	}
}

// Get retrieves the representation of the instance
func (TLSSettingData SettingData) Get() string {
	return TLSSettingData.base.Get(nil)
}

// Enumerates the instances of this class
func (TLSSettingData SettingData) Enumerate() string {
	return TLSSettingData.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (TLSSettingData SettingData) Pull(enumerationContext string) string {
	return TLSSettingData.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (TLSSettingData SettingData) Put(tlsSettingData TLSSettingData) string {
	return TLSSettingData.base.Put(tlsSettingData, false, nil)
}
