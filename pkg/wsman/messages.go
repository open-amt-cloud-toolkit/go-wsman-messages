/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wsman

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips"
)

type Messages struct {
	client *client.Client
	AMT    amt.Messages
	CIM    cim.Messages
	IPS    ips.Messages
}

type ClientParameters struct {
	Target            string
	Username          string
	Password          string
	UseDigest         bool
	UseTLS            bool
	SelfSignedAllowed bool
}

func NewMessages(cp ClientParameters) Messages {
	client := client.NewWsmanClient(cp.Target, cp.Username, cp.Password, cp.UseDigest, cp.UseTLS, cp.SelfSignedAllowed)
	m := Messages{
		client: client,
	}
	m.AMT = amt.NewMessages(client)
	m.CIM = cim.NewMessages(client)
	m.IPS = ips.NewMessages(client)
	return m
}
