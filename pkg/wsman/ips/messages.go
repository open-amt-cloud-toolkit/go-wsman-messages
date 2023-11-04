/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ips

import (
	// "reflect"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips/alarmclock"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips/hostbasedsetup"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips/ieee8021x"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips/optin"
)

type Messages struct {
	wsmanMessageCreator        *message.WSManMessageCreator
	OptInService               optin.Service
	HostBasedSetupService      hostbasedsetup.Service
	AlarmClockOccurrence       alarmclock.Occurrence
	IEEE8021xCredentialContext ieee8021x.CredentialContext
	IEEE8021xSettings          ieee8021x.Settings
}

func NewMessages(client *client.Client) Messages {
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	m := Messages{
		wsmanMessageCreator: wsmanMessageCreator,
	}
	m.OptInService = optin.NewOptInService(wsmanMessageCreator)
	m.HostBasedSetupService = hostbasedsetup.NewHostBasedSetupService(wsmanMessageCreator)
	m.AlarmClockOccurrence = alarmclock.NewAlarmClockOccurrence(wsmanMessageCreator)
	m.IEEE8021xCredentialContext = ieee8021x.NewIEEE8021xCredentialContext(wsmanMessageCreator)
	m.IEEE8021xSettings = ieee8021x.NewIEEE8021xSettings(wsmanMessageCreator)
	return m
}
