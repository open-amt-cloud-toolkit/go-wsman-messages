/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
)

const AMT_RemoteAccessService = "AMT_RemoteAccessService"

type Service struct {
	base wsman.Base
}
type MPServer struct {
	AccessInfo string
	InfoFormat MPServerInfoFormat
	Port       int
	AuthMethod MPServerAuthMethod
	Username   string
	Password   string
	CommonName string
}

type MPServerInfoFormat uint8

const (
	IPv4Address MPServerInfoFormat = 3
	IPv6Address MPServerInfoFormat = 4
	FQDN        MPServerInfoFormat = 201
)

type MPServerAuthMethod uint8

const (
	MutualAuthentication           MPServerAuthMethod = 1
	UsernamePasswordAuthentication MPServerAuthMethod = 2
)

func NewRemoteAccessService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{
		base: wsman.NewBase(wsmanMessageCreator, AMT_RemoteAccessService),
	}
}

// Get retrieves the representation of the instance
func (RemoteAccessService Service) Get() string {
	return RemoteAccessService.base.Get(nil)
}
func (RemoteAccessService Service) Enumerate() string {
	return RemoteAccessService.base.Enumerate()
}
func (RemoteAccessService Service) Pull(enumerationContext string) string {
	return RemoteAccessService.base.Pull(enumerationContext)
}
func (r Service) AddMPS(mpServer MPServer) string {
	header := r.base.WSManMessageCreator.CreateHeader(string(actions.AddMps), AMT_RemoteAccessService, nil, "", "")
	body := fmt.Sprintf(`<Body><h:AddMpServer_INPUT xmlns:h="%s%s"><h:AccessInfo>%s</h:AccessInfo><h:InfoFormat>%d</h:InfoFormat><h:Port>%d</h:Port><h:AuthMethod>%d</h:AuthMethod><h:Username>%s</h:Username><h:Password>%s</h:Password><h:CN>%s</h:CN></h:AddMpServer_INPUT></Body>`, r.base.WSManMessageCreator.ResourceURIBase, AMT_RemoteAccessService, mpServer.AccessInfo, mpServer.InfoFormat, mpServer.Port, mpServer.AuthMethod, mpServer.Username, mpServer.Password, mpServer.CommonName)
	return r.base.WSManMessageCreator.CreateXML(header, body)
}

func (r Service) AddRemoteAccessPolicyRule(remoteAccessPolicyRule RemoteAccessPolicyRule, selector wsman.Selector) string {
	header := r.base.WSManMessageCreator.CreateHeader(string(actions.AddRemoteAccessPolicyRule), AMT_RemoteAccessService, nil, "", "")
	body := fmt.Sprintf(`<Body><h:AddRemoteAccessPolicyRule_INPUT xmlns:h="%s%s"><h:Trigger>%d</h:Trigger><h:TunnelLifeTime>%d</h:TunnelLifeTime><h:ExtendedData>%s</h:ExtendedData><h:MpServer><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">%s%s</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="%s">%s</Selector></SelectorSet></ReferenceParameters></h:MpServer></h:AddRemoteAccessPolicyRule_INPUT></Body>`, r.base.WSManMessageCreator.ResourceURIBase,
		AMT_RemoteAccessService,
		remoteAccessPolicyRule.Trigger,
		remoteAccessPolicyRule.TunnelLifeTime,
		remoteAccessPolicyRule.ExtendedData,
		r.base.WSManMessageCreator.ResourceURIBase,
		"AMT_ManagementPresenceRemoteSAP", selector.Name, selector.Value)
	return r.base.WSManMessageCreator.CreateXML(header, body)
}
