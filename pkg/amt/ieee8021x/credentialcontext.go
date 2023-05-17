/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_IEEE8021xCredentialContext = "AMT_8021xCredentialContext"

type CredentialContext struct {
	base wsman.Base
}

func NewIEEE8021xCredentialContext(wsmanMessageCreator *wsman.WSManMessageCreator) CredentialContext {
	return CredentialContext{
		base: wsman.NewBase(wsmanMessageCreator, AMT_IEEE8021xCredentialContext),
	}
}

// Get retrieves the representation of the instance
func (IEEE8021xCredentialContext CredentialContext) Get() string {
	return IEEE8021xCredentialContext.base.Get(nil)
}

// Enumerates the instances of this class
func (IEEE8021xCredentialContext CredentialContext) Enumerate() string {
	return IEEE8021xCredentialContext.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (IEEE8021xCredentialContext CredentialContext) Pull(enumerationContext string) string {
	return IEEE8021xCredentialContext.base.Pull(enumerationContext)
}
