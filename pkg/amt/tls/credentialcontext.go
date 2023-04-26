/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_TLSCredentialContext = "AMT_TLSCredentialContext"

type CredentialContext struct {
	base wsman.Base
}

func NewTLSCredentialContext(wsmanMessageCreator *wsman.WSManMessageCreator) CredentialContext {
	return CredentialContext{
		base: wsman.NewBase(wsmanMessageCreator, AMT_TLSCredentialContext),
	}
}
func (TLSCredentialContext CredentialContext) Get() string {
	return TLSCredentialContext.base.Get(nil)
}
func (TLSCredentialContext CredentialContext) Enumerate() string {
	return TLSCredentialContext.base.Enumerate()
}
func (TLSCredentialContext CredentialContext) Pull(enumerationContext string) string {
	return TLSCredentialContext.base.Pull(enumerationContext)
}
func (TLSCredentialContext CredentialContext) Delete(selector *wsman.Selector) string {
	return TLSCredentialContext.base.Delete(selector)
}
func (TLSCredentialContext CredentialContext) Create(certHandle string) string {
	header := TLSCredentialContext.base.WSManMessageCreator.CreateHeader(string(wsman.BaseActionsCreate), AMT_TLSCredentialContext, nil, "", "")
	body := fmt.Sprintf(`<Body><h:AMT_TLSCredentialContext xmlns:h="%sAMT_TLSCredentialContext"><h:ElementInContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementInContext><h:ElementProvidingContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_TLSProtocolEndpointCollection</w:ResourceURI><w:SelectorSet><w:Selector Name="ElementName">TLSProtocolEndpointInstances Collection</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementProvidingContext></h:AMT_TLSCredentialContext></Body>`, TLSCredentialContext.base.WSManMessageCreator.ResourceURIBase, TLSCredentialContext.base.WSManMessageCreator.ResourceURIBase, certHandle, TLSCredentialContext.base.WSManMessageCreator.ResourceURIBase)
	return TLSCredentialContext.base.WSManMessageCreator.CreateXML(header, body)
}
