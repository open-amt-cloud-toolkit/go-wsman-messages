/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/json"
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName      xml.Name     `xml:"Body"`
		RemoteAccess RemoteAccess `xml:"AMT_RemoteAccessService"`

		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}
	RemoteAccess struct {
		CreationClassName       string
		ElementName             string
		Name                    string
		SystemCreationClassName string
		SystemName              string
	}
	PullResponse struct {
		Items []Item
	}
	Item struct {
		RemoteAccess RemoteAccess `xml:"AMT_RemoteAccessService"`
	}
)
type Service struct {
	base   message.Base
	client client.WSMan
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

func (w *Response) JSONStomps() string {
	jsonOutput, err := json.Marshal(w.Body)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}

func NewRemoteAccessService(wsmanMessageCreator *message.WSManMessageCreator) Service {
	return Service{
		base: message.NewBase(wsmanMessageCreator, AMT_RemoteAccessService),
	}
}

func NewRemoteAccessServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_RemoteAccessService, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (RemoteAccessService Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: RemoteAccessService.base.Get(nil),
		},
	}
	// send the message to AMT
	err = RemoteAccessService.base.Execute(response.Message)
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
func (RemoteAccessService Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: RemoteAccessService.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = RemoteAccessService.base.Execute(response.Message)
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
func (RemoteAccessService Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: RemoteAccessService.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = RemoteAccessService.base.Execute(response.Message)
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

/*func (r Service) AddMPS(mpServer MPServer) string {
	header := r.base.WSManMessageCreator.CreateHeader(string(actions.AddMps), AMT_RemoteAccessService, nil, "", "")
	body := fmt.Sprintf(`<Body><h:AddMpServer_INPUT xmlns:h="%s%s"><h:AccessInfo>%s</h:AccessInfo><h:InfoFormat>%d</h:InfoFormat><h:Port>%d</h:Port><h:AuthMethod>%d</h:AuthMethod><h:Username>%s</h:Username><h:Password>%s</h:Password><h:CN>%s</h:CN></h:AddMpServer_INPUT></Body>`, r.base.WSManMessageCreator.ResourceURIBase, AMT_RemoteAccessService, mpServer.AccessInfo, mpServer.InfoFormat, mpServer.Port, mpServer.AuthMethod, mpServer.Username, mpServer.Password, mpServer.CommonName)
	return r.base.WSManMessageCreator.CreateXML(header, body)
}

func (r Service) AddRemoteAccessPolicyRule(remoteAccessPolicyRule RemoteAccessPolicyRule, selector message.Selector) string {
	header := r.base.WSManMessageCreator.CreateHeader(string(actions.AddRemoteAccessPolicyRule), AMT_RemoteAccessService, nil, "", "")
	body := fmt.Sprintf(`<Body><h:AddRemoteAccessPolicyRule_INPUT xmlns:h="%s%s"><h:Trigger>%d</h:Trigger><h:TunnelLifeTime>%d</h:TunnelLifeTime><h:ExtendedData>%s</h:ExtendedData><h:MpServer><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">%s%s</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="%s">%s</Selector></SelectorSet></ReferenceParameters></h:MpServer></h:AddRemoteAccessPolicyRule_INPUT></Body>`, r.base.WSManMessageCreator.ResourceURIBase,
		AMT_RemoteAccessService,
		remoteAccessPolicyRule.Trigger,
		remoteAccessPolicyRule.TunnelLifeTime,
		remoteAccessPolicyRule.ExtendedData,
		r.base.WSManMessageCreator.ResourceURIBase,
		"AMT_ManagementPresenceRemoteSAP", selector.Name, selector.Value)
	return r.base.WSManMessageCreator.CreateXML(header, body)
}*/
