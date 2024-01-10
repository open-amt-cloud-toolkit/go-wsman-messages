/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Service struct {
	base message.Base
}

// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                          xml.Name `xml:"Body"`
		GetResponse                      TimeSynchronizationServiceResponse
		EnumerateResponse                common.EnumerateResponse
		PullResponse                     PullResponse
		GetLowAccuracyTimeSynchResponse  GetLowAccuracyTimeSynchResponse
		SetHighAccuracyTimeSynchResponse SetHighAccuracyTimeSynchResponse
	}

	PullResponse struct {
		XMLName                         xml.Name                             `xml:"PullResponse"`
		TimeSynchronizationServiceItems []TimeSynchronizationServiceResponse `xml:"Items>AMT_TimeSynchronizationService"`
	}

	TimeSynchronizationServiceResponse struct {
		XMLName                 xml.Name             `xml:"AMT_TimeSynchronizationService"`
		Name                    string               `xml:"Name,omitempty"`
		CreationClassName       string               `xml:"CreationClassName,omitempty"`
		SystemName              string               `xml:"SystemName,omitempty"`
		SystemCreationClassName string               `xml:"SystemCreationClassName,omitempty"`
		ElementName             string               `xml:"ElementName,omitempty"`
		EnabledState            EnabledState         `xml:"EnabledState,omitempty"`
		RequestedState          RequestedState       `xml:"RequestedState,omitempty"`
		LocalTimeSyncEnabled    LocalTimeSyncEnabled `xml:"LocalTimeSyncEnabled,omitempty"`
		TimeSource              TimeSource           `xml:"TimeSource,omitempty"`
	}

	EnabledState         int
	RequestedState       int
	LocalTimeSyncEnabled int
	TimeSource           int

	GetLowAccuracyTimeSynchResponse struct {
		XMLName     xml.Name `xml:"GetLowAccuracyTimeSynch_OUTPUT"`
		Ta0         int64    `xml:"Ta0"`
		ReturnValue int      `xml:"ReturnValue"`
	}

	SetHighAccuracyTimeSynchResponse struct {
		XMLName     xml.Name `xml:"SetHighAccuracyTimeSynch_OUTPUT"`
		ReturnValue int      `xml:"ReturnValue"`
	}
)

// Request Types
type (
	SetHighAccuracyTimeSynch_INPUT struct {
		XMLName xml.Name `xml:"h:SetHighAccuracyTimeSynch_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Ta0     int64    `xml:"h:Ta0"`
		Tm1     int64    `xml:"h:Tm1"`
		Tm2     int64    `xml:"h:Tm2"`
	}
)
