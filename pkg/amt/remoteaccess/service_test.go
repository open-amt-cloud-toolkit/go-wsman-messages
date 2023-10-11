/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	//"fmt"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	//"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_AlarmClockService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	// client := wsmantesting.MockClient{
	// 	PackageUnderTest: "amt/general",
	// }
	client := wsman.NewClient("http://localhost:16992/wsman", "admin", "Intel123!", true)

	elementUnderTest := NewRemoteAccessServiceWithClient(wsmanMessageCreator, client)
	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc     func() (ResponseStomps, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_RemoteAccessService Get wsman message", 
				"AMT_RemoteAccessService", 
				wsmantesting.GET, 
				"", 
				func() (ResponseStomps, error) {
					//client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				BodyStomps{
					XMLName: xml.Name{Space: "", Local: ""},
					RemoteAccess: RemoteAccess{
	
					},
				},
			},
			//ENUMERATES
			/*{
				"should create a valid AMT_RemoteAccessService Enumerate wsman message", 
				"AMT_RemoteAccessService", 
				wsmantesting.ENUMERATE, 
				wsmantesting.ENUMERATE_BODY, 
				func() (ResponseStomps, error) {
					//client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				BodyStomps{},
			},*/
			//PULLS
		// 	{"should create a valid AMT_RemoteAccessService Pull wsman message", "AMT_RemoteAccessService", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
		// 	{"should create a valid AMT_RemoteAccessService AddMPS wsman message", "AMT_RemoteAccessService", string(actions.AddMps), fmt.Sprintf(`<h:AddMpServer_INPUT xmlns:h="%s%s"><h:AccessInfo>%s</h:AccessInfo><h:InfoFormat>%d</h:InfoFormat><h:Port>%d</h:Port><h:AuthMethod>%d</h:AuthMethod><h:Username>%s</h:Username><h:Password>%s</h:Password><h:CN>%s</h:CN></h:AddMpServer_INPUT>`, resourceUriBase, AMT_RemoteAccessService, "AccessInfo", 1, 2, 3, "Username", "Password", "CommonName"), func() string {
		// 		mpsServer := MPServer{
		// 			AccessInfo: "AccessInfo",
		// 			InfoFormat: 1,
		// 			Port:       2,
		// 			AuthMethod: 3,
		// 			Username:   "Username",
		// 			Password:   "Password",
		// 			CommonName: "CommonName",
		// 		}
		// 		return elementUnderTest.AddMPS(mpsServer)
		// 	}},
		// 	{"should create a valid AMT_RemoteAccessPolicyRule wsman message", "AMT_RemoteAccessService", string(actions.AddRemoteAccessPolicyRule), fmt.Sprintf(`<h:AddRemoteAccessPolicyRule_INPUT xmlns:h="%s%s"><h:Trigger>%d</h:Trigger><h:TunnelLifeTime>%d</h:TunnelLifeTime><h:ExtendedData>%s</h:ExtendedData><h:MpServer><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">%s%s</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="myselector">true</Selector></SelectorSet></ReferenceParameters></h:MpServer></h:AddRemoteAccessPolicyRule_INPUT>`, resourceUriBase, AMT_RemoteAccessService, 2, 0, "0300", "http://intel.com/wbem/wscim/1/amt-schema/1/", "AMT_ManagementPresenceRemoteSAP"), func() string {
		// 		remoteAccessPolicyRule := RemoteAccessPolicyRule{
		// 			Trigger:        2,
		// 			TunnelLifeTime: 0,
		// 			ExtendedData:   "0300",
		// 		}
		// 		selector := message.Selector{
		// 			Name:  "myselector",
		// 			Value: "true",
		// 		}
		// 		return elementUnderTest.AddRemoteAccessPolicyRule(remoteAccessPolicyRule, selector)
		// 	}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				println(response.XMLOutput)
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.BodyStomps)
			})
		}
	})
}
