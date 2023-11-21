/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package redirection

import (
	//"encoding/json"
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
		XMLName     xml.Name    `xml:"Body"`
		Redirection Redirection `xml:"AMT_RedirectionService"`

		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}
	Redirection struct {
		CreationClassName       string
		ElementName             string
		EnabledState            int
		ListenerEnabled         bool
		Name                    string
		SystemCreationClassName string
		SystemName              string
	}
	PullResponse struct {
		Items []Item
	}
	Item struct {
		Redirection Redirection `xml:"AMT_RedirectionService"`
	}
)
type RedirectionResponse struct {
	AMT_RedirectionService RedirectionService
}

type RedirectionService struct {
	Name                    string
	CreationClassName       string
	SystemName              string
	SystemCreationClassName string
	ElementName             string
	ListenerEnabled         bool
	AccessLog               []string
	EnabledState            EnabledState
}

type EnabledState int

const AMT_RedirectionService = "AMT_RedirectionService"

const (
	Unknown EnabledState = iota
	Other
	Enabled
	Disabled
	ShuttingDown
	NotApplicable
	EnabledButOffline
	InTest
	Deferred
	Quiesce
	Starting
	DMTFReserved
	IDERAndSOLAreDisabled         = 32768
	IDERIsEnabledAndSOLIsDisabled = 32769
	SOLIsEnabledAndIDERIsDisabled = 32770
	IDERAndSOLAreEnabled          = 32771
)

type RequestedState int

const (
	DisableIDERAndSOL       RequestedState = 32768
	EnableIDERAndDisableSOL RequestedState = 32769
	EnableSOLAndDisableIDER RequestedState = 32770
	EnableIDERAndSOL        RequestedState = 32771
)

type Service struct {
	base   message.Base
	client client.WSMan
}

// func (w *Response) JSON() string {
// 	jsonOutput, err := json.Marshal(w.Body)
// 	if err != nil {
// 		return ""
// 	}
// 	return string(jsonOutput)
// }

func NewRedirectionService(wsmanMessageCreator *message.WSManMessageCreator) Service {
	return Service{
		base: message.NewBase(wsmanMessageCreator, AMT_RedirectionService),
	}
}

func NewRedirectionServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base:   message.NewBaseWithClient(wsmanMessageCreator, AMT_RedirectionService, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (s Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Get(nil),
		},
	}

	// send the message to AMT
	err = s.base.Execute(response.Message)
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
func (s Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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

func (s Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
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

// // Put will change properties of the selected instance
// func (RedirectionService Service) Put(redirectionService RedirectionService) string {
// 	return RedirectionService.base.Put(redirectionService, false, nil)
// }

// // RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
// func (RedirectionService Service) RequestStateChange(requestedState RequestedState) string {
// 	return RedirectionService.base.RequestStateChange(actions.RequestStateChange(AMT_RedirectionService), int(requestedState))
// }
