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

// Enumerates the instances of this class
func (ManagementPresenceRemoteSAP RemoteSAP) Enumerate() string {
	return ManagementPresenceRemoteSAP.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (ManagementPresenceRemoteSAP RemoteSAP) Pull(enumerationContext string) string {
	return ManagementPresenceRemoteSAP.base.Pull(enumerationContext)
}

// Delete removes a the specified instance
func (ManagementPresenceRemoteSAP RemoteSAP) Delete(handle string) string {
	selector := wsman.Selector{Name: "Name", Value: handle}
	return ManagementPresenceRemoteSAP.base.Delete(selector)
}
