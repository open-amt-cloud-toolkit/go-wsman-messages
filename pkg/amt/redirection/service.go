package redirection

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
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
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
	EnabledStateDMTFReserved
	EnabledStateIDERAndSOLAreDisabled         = 32768
	EnabledStateIDERIsEnabledAndSOLIsDisabled = 32769
	EnabledStateSOLIsEnabledAndIDERIsDisabled = 32770
	EnabledStateIDERAndSOLAreEnabled          = 32771
)

type RequestedState int

const (
	RequestedStateDisableIDERAndSOL       RequestedState = 32768
	RequestedStateEnableIDERAndDisableSOL RequestedState = 32769
	RequestedStateEnableSOLAndDisableIDER RequestedState = 32770
	RequestedStateEnableIDERAndSOL        RequestedState = 32771
)

type Service struct {
	base wsman.Base
}

func NewRedirectionService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{
		base: wsman.NewBase(wsmanMessageCreator, AMT_RedirectionService),
	}
}
func (RedirectionService Service) Get() string {
	return RedirectionService.base.Get(nil)
}
func (RedirectionService Service) Enumerate() string {
	return RedirectionService.base.Enumerate()
}
func (RedirectionService Service) Pull(enumerationContext string) string {
	return RedirectionService.base.Pull(enumerationContext)
}
func (RedirectionService Service) Put(redirectionService RedirectionService) string {
	return RedirectionService.base.Put(redirectionService, false, nil)
}
func (RedirectionService Service) RequestStateChange(requestedState RequestedState) string {
	return RedirectionService.base.RequestStateChange(actions.RequestStateChange(AMT_RedirectionService), int(requestedState))
}
