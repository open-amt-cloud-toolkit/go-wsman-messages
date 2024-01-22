/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewRemoteAccessServiceWithClient instantiates a new Service
func NewRemoteAccessServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_RemoteAccessService, client),
	}
}

// Get retrieves the representation of the instance
func (service Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Get(nil),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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
func (service Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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
func (service Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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

// AddMPS adds a Management Presence Server to the Intel® AMT subsystem.
// Creates an AMT_ManagementPresenceRemoteSAP instance and an AMT_RemoteAccessCredentialContext association to a credential.
// This credential may be an existing AMT_PublicKeyCertificate instance (if the created MPS is configured to use mutual authentication).
// If the created MpServer is configured to use username password authentication, an AMT_MPSUsernamePassword instance is created and used as the associated credential.
func (service Service) AddMPS(mpServer AddMpServerRequest) (response Response, err error) {
	mpServer.H = fmt.Sprintf("%s%s", message.AMTSchema, AMT_RemoteAccessService)
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_RemoteAccessService, AddMps), AMT_RemoteAccessService, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(AddMps), AMT_RemoteAccessService, mpServer)
	// body := fmt.Sprintf(`<Body><h:AddMpServer_INPUT xmlns:h="%s%s"><h:AccessInfo>%s</h:AccessInfo><h:InfoFormat>%d</h:InfoFormat><h:Port>%d</h:Port><h:AuthMethod>%d</h:AuthMethod><h:Username>%s</h:Username><h:Password>%s</h:Password><h:CN>%s</h:CN></h:AddMpServer_INPUT></Body>`, service.base.WSManMessageCreator.ResourceURIBase, AMT_RemoteAccessService, mpServer.AccessInfo, mpServer.InfoFormat, mpServer.Port, mpServer.AuthMethod, mpServer.Username, mpServer.Password, mpServer.CommonName)
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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

// AddRemoteAccessPolicyRule adds a Remote Access policy to the Intel® AMT subsystem.
// The policy defines an event that will trigger an establishment of a tunnel between AMT and a pre-configured MPS.
// Creates an AMT_RemoteAccessPolicyRule instance and associates it to a given list of AMT_ManagementPresenceRemoteSAP instances with AMT_PolicySetAppliesToElement association instances.
// Returns an XML string representing the WS-Management message to be sent to the Intel® AMT subsystem.
func (service Service) AddRemoteAccessPolicyRule(remoteAccessPolicyRule RemoteAccessPolicyRuleRequest, name string) (response Response, err error) {
	selector := message.Selector{
		Name:  "Name",
		Value: name,
	}
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_RemoteAccessService, AddRemoteAccessPolicyRule), AMT_RemoteAccessService, nil, "", "")
	body := fmt.Sprintf(`<Body><h:AddRemoteAccessPolicyRule_INPUT xmlns:h="%s%s"><h:Trigger>%d</h:Trigger><h:TunnelLifeTime>%d</h:TunnelLifeTime><h:ExtendedData>%s</h:ExtendedData><h:MpServer><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">%s%s</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="%s">%s</Selector></SelectorSet></ReferenceParameters></h:MpServer></h:AddRemoteAccessPolicyRule_INPUT></Body>`,
		service.base.WSManMessageCreator.ResourceURIBase,
		AMT_RemoteAccessService,
		remoteAccessPolicyRule.Trigger,
		remoteAccessPolicyRule.TunnelLifeTime,
		remoteAccessPolicyRule.ExtendedData,
		service.base.WSManMessageCreator.ResourceURIBase,
		"AMT_ManagementPresenceRemoteSAP", selector.Name, selector.Value)
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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
