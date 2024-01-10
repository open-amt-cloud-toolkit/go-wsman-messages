/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

func NewTLSCredentialContextWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) CredentialContext {
	return CredentialContext{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_TLSCredentialContext, client),
	}
}

// Get retrieves the representation of the instance
func (credentialContext CredentialContext) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Get(nil),
		},
	}
	// send the message to AMT
	err = credentialContext.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (credentialContext CredentialContext) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = credentialContext.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pulls instances of this class, following an Enumerate operation
func (credentialContext CredentialContext) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = credentialContext.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// // Delete removes a the specified instance
// func (credentialContext CredentialContext) Delete(handle string) (response Response, err error) {
// 	selector := message.Selector{Name: "Name", Value: handle}
// 	response = Response{
// 		Message: &client.Message{
// 			XMLInput: credentialContext.base.Delete(selector),
// 		},
// 	}
// 	// send the message to AMT
// 	err = credentialContext.base.Execute(response.Message)
// 	if err != nil {
// 		return
// 	}
// 	// put the xml response into the go struct
// 	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// // Creates a new instance of this class
// func (credentialContext CredentialContext) Create(certHandle string) (response Response, err error) {
// 	header := credentialContext.base.WSManMessageCreator.CreateHeader(string(actions.Create), AMT_TLSCredentialContext, nil, "", "")
// 	body := fmt.Sprintf(`<Body><h:AMT_TLSCredentialContext xmlns:h="%sAMT_TLSCredentialContext"><h:ElementInContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementInContext><h:ElementProvidingContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_TLSProtocolEndpointCollection</w:ResourceURI><w:SelectorSet><w:Selector Name="ElementName">TLSProtocolEndpointInstances Collection</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementProvidingContext></h:AMT_TLSCredentialContext></Body>`, credentialContext.base.WSManMessageCreator.ResourceURIBase, credentialContext.base.WSManMessageCreator.ResourceURIBase, certHandle, credentialContext.base.WSManMessageCreator.ResourceURIBase)
// 	response = Response{
// 		Message: &client.Message{
// 			XMLInput: credentialContext.base.WSManMessageCreator.CreateXML(header, body),
// 		},
// 	}
// 	// send the message to AMT
// 	err = credentialContext.base.Execute(response.Message)
// 	if err != nil {
// 		return
// 	}
// 	// put the xml response into the go struct
// 	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
// 	if err != nil {
// 		return
// 	}
// 	return
// }
