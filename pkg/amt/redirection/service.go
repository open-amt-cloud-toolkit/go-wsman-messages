/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

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
