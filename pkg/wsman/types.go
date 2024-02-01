/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wsman

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/ips"
)

// Messages implements client.WSMan, amt.Messages, cim.Messages, and ips.Messages
type Messages struct {
	Client client.WSMan
	AMT    amt.Messages
	CIM    cim.Messages
	IPS    ips.Messages
}

// ClientParameters struct defines the connection settings for wsman client
type ClientParameters struct {
	Target            string
	Username          string
	Password          string
	UseDigest         bool
	UseTLS            bool
	SelfSignedAllowed bool
}
