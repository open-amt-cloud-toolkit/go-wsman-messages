/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package messagelog

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/methods"
)

type MessageLog struct {
	base message.Base
}
type GetRecords_INPUT struct {
	XMLName             xml.Name `xml:"h:GetRecords_INPUT"`
	H                   string   `xml:"xmlns:h,attr"`
	IterationIdentifier int      `xml:"h:IterationIdentifier"`
	MaxReadRecords      int      `xml:"h:MaxReadRecords"`
}

func NewMessageLog(wsmanMessageCreator *message.WSManMessageCreator) MessageLog {
	return MessageLog{
		base: message.NewBase(wsmanMessageCreator, AMT_MessageLog),
	}
}

// Get retrieves the representation of the instance
func (MessageLog MessageLog) Get() string {
	return MessageLog.base.Get(nil)
}

// Enumerates the instances of this class
func (MessageLog MessageLog) Enumerate() string {
	return MessageLog.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (MessageLog MessageLog) Pull(enumerationContext string) string {
	return MessageLog.base.Pull(enumerationContext)
}
func (MessageLog MessageLog) GetRecords(identifier int) string {
	if identifier < 1 {
		identifier = 1
	}

	header := MessageLog.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_MessageLog, GetRecords), AMT_MessageLog, nil, "", "")
	body := MessageLog.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetRecords), AMT_MessageLog, &GetRecords_INPUT{
		IterationIdentifier: identifier,
		MaxReadRecords:      390,
	})

	return MessageLog.base.WSManMessageCreator.CreateXML(header, body)
}

func (MessageLog MessageLog) PositionToFirstRecord() string {
	header := MessageLog.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_MessageLog, PositionToFirstRecord), AMT_MessageLog, nil, "", "")
	body := fmt.Sprintf(`<Body><h:PositionToFirstRecord_INPUT xmlns:h="%s%s" /></Body>`, MessageLog.base.WSManMessageCreator.ResourceURIBase, AMT_MessageLog)

	return MessageLog.base.WSManMessageCreator.CreateXML(header, body)
}
