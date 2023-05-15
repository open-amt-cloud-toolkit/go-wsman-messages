/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kerberos

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
)

const AMT_KerberosSettingData = "AMT_KerberosSettingData"

type KerberosSettingData struct {
	base wsman.Base
}
type SetCredentialCacheState_INPUT struct {
	XMLName xml.Name `xml:"h:SetCredentialCacheState_INPUT"`
	H       string   `xml:"xmlns:h,attr"`
	Enabled bool     `xml:"h:Enabled"`
}

func NewKerberosSettingData(wsmanMessageCreator *wsman.WSManMessageCreator) KerberosSettingData {
	return KerberosSettingData{
		base: wsman.NewBase(wsmanMessageCreator, AMT_KerberosSettingData),
	}
}

// Get retrieves the representation of the instance
func (KerberosSettingData KerberosSettingData) Get() string {
	return KerberosSettingData.base.Get(nil)
}

// Enumerates the instances of this class
func (KerberosSettingData KerberosSettingData) Enumerate() string {
	return KerberosSettingData.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (KerberosSettingData KerberosSettingData) Pull(enumerationContext string) string {
	return KerberosSettingData.base.Pull(enumerationContext)
}
func (k KerberosSettingData) GetCredentialCacheState() string {
	header := k.base.WSManMessageCreator.CreateHeader(string(actions.GetCredentialCacheState), AMT_KerberosSettingData, nil, "", "")
	body := k.base.WSManMessageCreator.CreateBody("GetCredentialCacheState_INPUT", AMT_KerberosSettingData, nil)

	return k.base.WSManMessageCreator.CreateXML(header, body)
}

// func (k KerberosSettingData) SetCredentialCacheState(enabled bool) string {
// 	header := k.base.WSManMessageCreator.CreateHeader(string(actions.SetCredentialCacheState), AMT_KerberosSettingData, nil, "", "")
// 	body := k.base.WSManMessageCreator.CreateBody("SetCredentialCacheState_INPUT", AMT_KerberosSettingData, SetCredentialCacheState_INPUT{Enabled: enabled})

// 	return k.base.WSManMessageCreator.CreateXML(header, body)
// }
