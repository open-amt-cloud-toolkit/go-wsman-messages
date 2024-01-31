/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package wsman facilitates access to AMT, CIM, and IPS classes for communication with Intel® AMT devices.
package wsman

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips"
)

// NewMessages instantiates a new Messages class with client connection parameters
func NewMessages(wsman client.WSMan) Messages {
	m := Messages{}

	m.AMT = amt.NewMessages(m.Client)
	m.CIM = cim.NewMessages(m.Client)
	m.IPS = ips.NewMessages(m.Client)
	return m
}
