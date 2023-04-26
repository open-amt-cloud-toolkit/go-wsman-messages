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
func (TLSSettingData SettingData) Get() string {
	return TLSSettingData.base.Get(nil)
}
func (TLSSettingData SettingData) Enumerate() string {
	return TLSSettingData.base.Enumerate()
}
func (TLSSettingData SettingData) Pull(enumerationContext string) string {
	return TLSSettingData.base.Pull(enumerationContext)
}
func (TLSSettingData SettingData) Put(tlsSettingData TLSSettingData) string {
	return TLSSettingData.base.Put(tlsSettingData, false, nil)
}
