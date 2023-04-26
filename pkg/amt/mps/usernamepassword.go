/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mps

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

const AMT_MPSUsernamePassword = "AMT_MPSUsernamePassword"

type MPSUsernamePassword struct {
	models.SharedCredential
}
type UsernamePassword struct {
	base wsman.Base
}

func NewMPSUsernamePassword(wsmanMessageCreator *wsman.WSManMessageCreator) UsernamePassword {
	return UsernamePassword{
		base: wsman.NewBase(wsmanMessageCreator, AMT_MPSUsernamePassword),
	}
}
func (MPSUsernamePassword UsernamePassword) Get() string {
	return MPSUsernamePassword.base.Get(nil)
}
func (MPSUsernamePassword UsernamePassword) Enumerate() string {
	return MPSUsernamePassword.base.Enumerate()
}
func (MPSUsernamePassword UsernamePassword) Pull(enumerationContext string) string {
	return MPSUsernamePassword.base.Pull(enumerationContext)
}
func (MPSUsernamePassword UsernamePassword) Put(mpsUsernamePassword MPSUsernamePassword) string {
	return MPSUsernamePassword.base.Put(mpsUsernamePassword, false, nil)
}
