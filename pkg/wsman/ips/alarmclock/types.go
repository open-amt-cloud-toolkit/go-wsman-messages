/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Occurrence struct {
	base message.Base
}

// OUTPUT
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name     `xml:"Body"`
		PullResponse      PullResponse `xml:"PullResponse"`
		EnumerateResponse common.EnumerateResponse
		GetResponse       AlarmClockOccurrence `xml:"IPS_AlarmClockOccurrence"`
	}

	AlarmClockOccurrence struct {
		XMLName            xml.Name `xml:"IPS_AlarmClockOccurrence"`
		ElementName        string
		InstanceID         string
		StartTime          string
		Interval           string
		DeleteOnCompletion bool
	}

	PullResponse struct {
		XMLName xml.Name               `xml:"PullResponse"`
		Items   []AlarmClockOccurrence `xml:"Items>IPS_AlarmClockOccurrence"`
	}
)
