package userinitiatedconnection

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/redirection"
)

type RequestedState int

const AMT_UserInitiatedConnectionService = "AMT_UserInitiatedConnectionService"

const (
	RequestedStateAllInterfacesDisabled      RequestedState = 32768
	RequestedStateBIOSInterfaceEnabled       RequestedState = 32769
	RequestedStateOSInterfaceEnabled         RequestedState = 32770
	RequestedStateBIOSandOSInterfacesEnabled RequestedState = 32771
)

type Service struct {
	base wsman.Base
}

func NewUserInitiatedConnectionService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{
		base: wsman.NewBase(wsmanMessageCreator, AMT_UserInitiatedConnectionService),
	}
}
func (UserInitiatedConnectionService Service) Get() string {
	return UserInitiatedConnectionService.base.Get(nil)
}
func (UserInitiatedConnectionService Service) Enumerate() string {
	return UserInitiatedConnectionService.base.Enumerate()
}
func (UserInitiatedConnectionService Service) Pull(enumerationContext string) string {
	return UserInitiatedConnectionService.base.Pull(enumerationContext)
}
func (UserInitiatedConnectionService Service) RequestStateChange(requestedState RequestedState) string {
	return UserInitiatedConnectionService.base.RequestStateChange(actions.RequestStateChange(redirection.AMT_RedirectionService), int(requestedState))

}
