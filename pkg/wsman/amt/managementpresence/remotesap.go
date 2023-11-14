/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package managementpresence

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

const AMT_ManagementPresenceRemoteSAP = "AMT_ManagementPresenceRemoteSAP"

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName          xml.Name         `xml:"Body"`
		ManagementRemote ManagementRemote `xml:"AMT_EnvironmentDetectionSettingData"`

		EnumerateResponse common.EnumerateResponse
	}
	ManagementRemote struct {
		AccessInfo              string
		CN                      string
		CreationClassName       string
		ElementName             string
		InfoFormat              int
		Name                    string
		Port                    int
		SystemCreationClassName string
		SystemName              string
	}
)
type RemoteSAP struct {
	base   message.Base
	client client.WSMan
}

func NewManagementPresenceRemoteSAPWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) RemoteSAP {
	return RemoteSAP{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_ManagementPresenceRemoteSAP, client),
		client: client,
	}
}

func NewManagementPresenceRemoteSAP(wsmanMessageCreator *message.WSManMessageCreator) RemoteSAP {
	return RemoteSAP{
		base: message.NewBase(wsmanMessageCreator, AMT_ManagementPresenceRemoteSAP),
	}
}

// Get retrieves the representation of the instance
func (ManagementPresenceRemoteSAP RemoteSAP) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: ManagementPresenceRemoteSAP.base.Get(nil),
		},
	}

	// send the message to AMT
	err = ManagementPresenceRemoteSAP.base.Execute(response.Message)
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
func (ManagementPresenceRemoteSAP RemoteSAP) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: ManagementPresenceRemoteSAP.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = ManagementPresenceRemoteSAP.base.Execute(response.Message)
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
// func (ManagementPresenceRemoteSAP RemoteSAP) Pull(enumerationContext string) string {
// 	return ManagementPresenceRemoteSAP.base.Pull(enumerationContext)
// }

// // Delete removes a the specified instance
// func (ManagementPresenceRemoteSAP RemoteSAP) Delete(handle string) string {
// 	selector := message.Selector{Name: "Name", Value: handle}
// 	return ManagementPresenceRemoteSAP.base.Delete(selector)
// }
