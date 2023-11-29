/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kerberos

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/methods"
)

type KerberosSettingData struct {
	base message.Base
}
type SetCredentialCacheState_INPUT struct {
	XMLName xml.Name `xml:"h:SetCredentialCacheState_INPUT"`
	H       string   `xml:"xmlns:h,attr"`
	Enabled bool     `xml:"h:Enabled"`
}

func NewKerberosSettingData(wsmanMessageCreator *message.WSManMessageCreator) KerberosSettingData {
	return KerberosSettingData{
		base: message.NewBase(wsmanMessageCreator, AMT_KerberosSettingData),
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
	header := k.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_KerberosSettingData, GetCredentialCacheState), AMT_KerberosSettingData, nil, "", "")
	body := k.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetCredentialCacheState), AMT_KerberosSettingData, nil)

	return k.base.WSManMessageCreator.CreateXML(header, body)
}

// func (k KerberosSettingData) SetCredentialCacheState(enabled bool) string {
// 	header := k.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_KerberosSettingData, SetCredentialCacheState), AMT_KerberosSettingData, nil, "", "")
// 	body := k.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetCredentialCacheState), AMT_KerberosSettingData, SetCredentialCacheState_INPUT{Enabled: enabled})

// 	return k.base.WSManMessageCreator.CreateXML(header, body)
// }
