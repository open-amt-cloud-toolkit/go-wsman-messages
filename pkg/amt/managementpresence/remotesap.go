/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package managementpresence

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_ManagementPresenceRemoteSAP = "AMT_ManagementPresenceRemoteSAP"

type RemoteSAP struct {
	base wsman.Base
}

func NewManagementPresenceRemoteSAP(wsmanMessageCreator *wsman.WSManMessageCreator) RemoteSAP {
	return RemoteSAP{
		base: wsman.NewBase(wsmanMessageCreator, AMT_ManagementPresenceRemoteSAP),
	}
}

// Get retrieves the representation of the instance
func (ManagementPresenceRemoteSAP RemoteSAP) Get() string {
	return ManagementPresenceRemoteSAP.base.Get(nil)
}
func (ManagementPresenceRemoteSAP RemoteSAP) Enumerate() string {
	return ManagementPresenceRemoteSAP.base.Enumerate()
}
func (ManagementPresenceRemoteSAP RemoteSAP) Pull(enumerationContext string) string {
	return ManagementPresenceRemoteSAP.base.Pull(enumerationContext)
}
func (ManagementPresenceRemoteSAP RemoteSAP) Delete(selector *wsman.Selector) string {
	return ManagementPresenceRemoteSAP.base.Delete(selector)
}
