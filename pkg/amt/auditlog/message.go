/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Represents an Audit Log in the Intel AMT subsystem.
package auditlog

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
)

const AMT_AuditLog = "AMT_AuditLog"

type AuditLog struct {
	base wsman.Base
}

type readRecords_INPUT struct {
	XMLName    xml.Name `xml:"h:ReadRecords_INPUT"`
	H          string   `xml:"xmlns:h,attr"`
	StartIndex int      `xml:"h:StartIndex" json:"StartIndex"`
}

func NewAuditLog(wsmanMessageCreator *wsman.WSManMessageCreator) AuditLog {
	return AuditLog{base: wsman.NewBase(wsmanMessageCreator, AMT_AuditLog)}
}

// Get retrieves the representation of the instance
func (AuditLog AuditLog) Get() string {
	return AuditLog.base.Get(nil)
}

// Enumerates the instances of this class
func (AuditLog AuditLog) Enumerate() string {
	return AuditLog.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (AuditLog AuditLog) Pull(enumerationContext string) string {
	return AuditLog.base.Pull(enumerationContext)
}

// ReadRecords returns a list of consecutive audit log records in chronological order:
// The first record in the returned array is the oldest record stored in the log.
// startIndex Identifies the position of the first record to retrieve. An index of 1 indicates the first record in the log.
func (a AuditLog) ReadRecords(startIndex int) string {
	if startIndex < 1 {
		startIndex = 0
	}
	header := a.base.WSManMessageCreator.CreateHeader(string(actions.ReadRecords), AMT_AuditLog, nil, "", "")
	body := a.base.WSManMessageCreator.CreateBody("ReadRecords_INPUT", AMT_AuditLog, &readRecords_INPUT{StartIndex: startIndex})

	return a.base.WSManMessageCreator.CreateXML(header, body)
}
