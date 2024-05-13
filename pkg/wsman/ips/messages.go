/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ips

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/ips/alarmclock"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/ips/hostbasedsetup"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/ips/ieee8021x"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/ips/optin"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

type Messages struct {
	wsmanMessageCreator        *message.WSManMessageCreator
	OptInService               optin.Service
	HostBasedSetupService      hostbasedsetup.Service
	AlarmClockOccurrence       alarmclock.Occurrence
	IEEE8021xCredentialContext ieee8021x.CredentialContext
	IEEE8021xSettings          ieee8021x.Settings
}

func NewMessages(client client.WSMan) Messages {
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	m := Messages{
		wsmanMessageCreator: wsmanMessageCreator,
	}
	m.OptInService = optin.NewOptInServiceWithClient(wsmanMessageCreator, client)
	m.HostBasedSetupService = hostbasedsetup.NewHostBasedSetupServiceWithClient(wsmanMessageCreator, client)
	m.AlarmClockOccurrence = alarmclock.NewAlarmClockOccurrenceWithClient(wsmanMessageCreator, client)
	m.IEEE8021xCredentialContext = ieee8021x.NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator, client)
	m.IEEE8021xSettings = ieee8021x.NewIEEE8021xSettingsWithClient(wsmanMessageCreator, client)

	return m
}
