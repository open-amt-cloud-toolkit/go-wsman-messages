/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type CredentialContext struct {
	base wsman.Base
}

const IPS_8021xCredentialContext = "IPS_8021xCredentialContext"

// NewIEEE8021xCredentialContext returns a new instance of the IPS_8021xCredentialContext struct.
func NewIEEE8021xCredentialContext(wsmanMessageCreator *wsman.WSManMessageCreator) CredentialContext {
	return CredentialContext{
		base: wsman.NewBase(wsmanMessageCreator, IPS_8021xCredentialContext),
	}
}

// Get retrieves the representation of the instance
func (b CredentialContext) Get() string {
	return b.base.Get(nil)
}

func (b CredentialContext) Enumerate() string {
	return b.base.Enumerate()
}

func (b CredentialContext) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
