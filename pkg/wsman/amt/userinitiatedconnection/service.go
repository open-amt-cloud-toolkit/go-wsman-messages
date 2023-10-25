/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/redirection"
)

type RequestedState int

const AMT_UserInitiatedConnectionService = "AMT_UserInitiatedConnectionService"

const (
	AllInterfacesDisabled      RequestedState = 32768
	BIOSInterfaceEnabled       RequestedState = 32769
	OSInterfaceEnabled         RequestedState = 32770
	BIOSandOSInterfacesEnabled RequestedState = 32771
)

type Service struct {
	base message.Base
}

func NewUserInitiatedConnectionService(wsmanMessageCreator *message.WSManMessageCreator) Service {
	return Service{
		base: message.NewBase(wsmanMessageCreator, AMT_UserInitiatedConnectionService),
	}
}

// Get retrieves the representation of the instance
func (UserInitiatedConnectionService Service) Get() string {
	return UserInitiatedConnectionService.base.Get(nil)
}

// Enumerates the instances of this class
func (UserInitiatedConnectionService Service) Enumerate() string {
	return UserInitiatedConnectionService.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (UserInitiatedConnectionService Service) Pull(enumerationContext string) string {
	return UserInitiatedConnectionService.base.Pull(enumerationContext)
}

// RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
func (UserInitiatedConnectionService Service) RequestStateChange(requestedState RequestedState) string {
	return UserInitiatedConnectionService.base.RequestStateChange(actions.RequestStateChange(redirection.AMT_RedirectionService), int(requestedState))

}
