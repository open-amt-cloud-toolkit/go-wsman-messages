/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package messagelog

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

// OUTPUTS
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                       xml.Name `xml:"Body"`
		GetResponse                   MessageLogResponse
		EnumerateResponse             common.EnumerateResponse
		PullResponse                  PullResponse
		GetRecordsResponse            GetRecordsResponse
		PositionToFirstRecordResponse PositionToFirstRecordResponse
	}

	PullResponse struct {
		XMLName         xml.Name             `xml:"PullResponse"`
		MessageLogItems []MessageLogResponse `xml:"Items>AMT_MessageLog"`
	}

	MessageLogResponse struct {
		XMLName                xml.Name                 `xml:"AMT_MessageLog"`
		Capabilities           []int                    `xml:"Capabilities"`
		CharacterSet           CharacterSet             `xml:"CharacterSet"`
		CreationClassName      string                   `xml:"CreationClassName"`
		CurrentNumberOfRecords int                      `xml:"CurrentNumberOfRecords"`
		ElementName            string                   `xml:"ElementName"`
		EnabledDefault         models.EnabledDefault    `xml:"EnabledDefault"`
		EnabledState           models.EnabledState      `xml:"EnabledState"`
		HealthState            models.HealthState       `xml:"HealthState"`
		IsFrozen               bool                     `xml:"IsFrozen"`
		LastChange             LastChange               `xml:"LastChange"`
		LogState               LogState                 `xml:"LogState"`
		MaxLogSize             int                      `xml:"MaxLogSize"`
		MaxNumberOfRecords     int                      `xml:"MaxNumberOfRecords"`
		MaxRecordSize          int                      `xml:"MaxRecordSize"`
		Name                   string                   `xml:"Name"`
		OperationalStatus      models.OperationalStatus `xml:"OperationalStatus"`
		OverwritePolicy        OverwritePolicy          `xml:"OverwritePolicy"`
		PercentageNearFull     int                      `xml:"PercentageNearFull"`
		RequestedState         models.RequestedState    `xml:"RequestedState"`
		SizeOfHeader           int                      `xml:"SizeOfHeader"`
		SizeOfRecordHeader     int                      `xml:"SizeOfRecordHeader"`
		Status                 string                   `xml:"Status"`
	}

	Capabilities    int
	CharacterSet    int
	LastChange      int
	LogState        int
	OverwritePolicy int

	GetRecordsResponse struct {
		XMLName             xml.Name `xml:"GetRecords_OUTPUT"`
		IterationIdentifier int      `xml:"IterationIdentifier"`
		NoMoreRecords       bool     `xml:"NoMoreRecords"`
		RecordArray         []string `xml:"RecordArray"`
		ReturnValue         int      `xml:"ReturnValue"`
	}

	PositionToFirstRecordResponse struct {
		XMLName             xml.Name `xml:"PositionToFirstRecord_OUTPUT"`
		IterationIdentifier int      `xml:"IterationIdentifier"`
		ReturnValue         int      `xml:"ReturnValue"`
	}
)

// INPUTS
type GetRecords_INPUT struct {
	XMLName             xml.Name `xml:"h:GetRecords_INPUT"`
	H                   string   `xml:"xmlns:h,attr"`
	IterationIdentifier int      `xml:"h:IterationIdentifier"`
	MaxReadRecords      int      `xml:"h:MaxReadRecords"`
}
