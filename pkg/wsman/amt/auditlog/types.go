/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package auditlog

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

// INPUTS
type Service struct {
	base message.Base
}

type readRecords_INPUT struct {
	XMLName    xml.Name `xml:"h:ReadRecords_INPUT"`
	H          string   `xml:"xmlns:h,attr"`
	StartIndex int      `xml:"h:StartIndex" json:"StartIndex"`
}

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
		XMLName             xml.Name `xml:"Body"`
		EnumerateResponse   common.EnumerateResponse
		GetResponse         AuditLog
		PullResponse        PullResponse
		ReadRecordsResponse ReadRecords_OUTPUT
	}
	PullResponse struct {
		XMLName       xml.Name   `xml:"PullResponse"`
		AuditLogItems []AuditLog `xml:"Items>AMT_AuditLog"`
	}

	AuditLog struct {
		XMLName                xml.Name        `xml:"AMT_AuditLog"`
		OverwritePolicy        OverwritePolicy `xml:"OverwritePolicy,omitempty"`
		CurrentNumberOfRecords int             `xml:"CurrentNumberOfRecords,omitempty"`
		MaxNumberOfRecords     int             `xml:"MaxNumberOfRecords,omitempty"`
		ElementName            string          `xml:"ElementName,omitempty"`
		EnabledState           EnabledState    `xml:"EnabledState,omitempty"`
		RequestedState         RequestedState  `xml:"RequestedState,omitempty"`
		PercentageFree         int             `xml:"PercentageFree,omitempty"`
		Name                   string          `xml:"Name,omitempty"`
		TimeOfLastRecord       Datetime        `xml:"TimeOfLastRecord"`
		AuditState             int             `xml:"AuditState,omitempty"`
		MaxAllowedAuditors     int             `xml:"MaxAllowedAuditors,omitempty"`
		StoragePolicy          StoragePolicy   `xml:"StoragePolicy,omitempty"`
		MinDaysToKeep          int             `xml:"MinDaysToKeep,omitempty"`
	}

	Datetime struct {
		Datetime string `xml:"Datetime,omitempty"`
	}

	ReadRecords_OUTPUT struct {
		XMLName          xml.Name `xml:"ReadRecords_OUTPUT,omitempty"`
		TotalRecordCount int      `xml:"TotalRecordCount,omitempty"`
		RecordsReturned  int      `xml:"RecordsReturned,omitempty"`
		EventRecords     []string `xml:"EventRecords,omitempty"`
		ReturnValue      int      `xml:"ReturnValue,omitempty"`
	}

	OverwritePolicy int
	EnabledState    int
	RequestedState  int
	StoragePolicy   int
)
