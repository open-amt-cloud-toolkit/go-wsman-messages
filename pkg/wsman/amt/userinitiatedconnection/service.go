/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

import (
	"encoding/json"
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type RequestedState int

const AMT_UserInitiatedConnectionService = "AMT_UserInitiatedConnectionService"

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName xml.Name `xml:"Body"`
		User    User     `xml:"AMT_UserInitiatedConnectionService"`

		EnumerateResponse common.EnumerateResponse
	}
	User struct {
		CreationClassName       string
		ElementName             string
		EnabledState            int
		Name                    string
		SystemCreationClassName string
		SystemName              string
	}
)

const (
	AllInterfacesDisabled      RequestedState = 32768
	BIOSInterfaceEnabled       RequestedState = 32769
	OSInterfaceEnabled         RequestedState = 32770
	BIOSandOSInterfacesEnabled RequestedState = 32771
)

func (w *Response) JSON() string {
	jsonOutput, err := json.Marshal(w.Body)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}

type Service struct {
	base   message.Base
	client client.WSMan
}

func NewUserInitiatedConnectionService(wsmanMessageCreator *message.WSManMessageCreator) Service {
	return Service{
		base: message.NewBase(wsmanMessageCreator, AMT_UserInitiatedConnectionService),
	}
}

func NewUserInitiatedConnectionServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_UserInitiatedConnectionService, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (UserInitiatedConnectionService Service) Get() (response Response, err error) {

	response = Response{
		Message: &client.Message{
			XMLInput: UserInitiatedConnectionService.base.Get(nil),
		},
	}

	// send the message to AMT
	err = UserInitiatedConnectionService.base.Execute(response.Message)
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
func (UserInitiatedConnectionService Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: UserInitiatedConnectionService.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = UserInitiatedConnectionService.base.Execute(response.Message)
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
// func (UserInitiatedConnectionService Service) Pull(enumerationContext string) string {
// 	return UserInitiatedConnectionService.base.Pull(enumerationContext)
// }

// // RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
// func (UserInitiatedConnectionService Service) RequestStateChange(requestedState RequestedState) string {
// 	return UserInitiatedConnectionService.base.RequestStateChange(actions.RequestStateChange(redirection.AMT_RedirectionService), int(requestedState))

// }
