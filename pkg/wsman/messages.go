/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package wsman facilitates access to AMT, CIM, and IPS classes for communication with Intel® AMT devices.
package wsman

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/ips"
)

// NewMessages instantiates a new Messages class with client connection parameters
func NewMessages(cp client.Parameters) Messages {
	client := client.NewWsman(cp)
	m := Messages{
		client: client,
	}
	m.AMT = amt.NewMessages(client)
	m.CIM = cim.NewMessages(client)
	m.IPS = ips.NewMessages(client)
	return m
}
