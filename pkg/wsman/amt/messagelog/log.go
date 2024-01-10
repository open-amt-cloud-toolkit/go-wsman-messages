/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package messagelog

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

type MessageLog struct {
	base message.Base
}

func NewMessageLogWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) MessageLog {
	return MessageLog{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_MessageLog, client),
	}
}

// Get retrieves the representation of the instance
func (messageLog MessageLog) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: messageLog.base.Get(nil),
		},
	}
	// send the message to AMT
	err = messageLog.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (messageLog MessageLog) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: messageLog.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = messageLog.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pulls instances of this class, following an Enumerate operation
func (messageLog MessageLog) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: messageLog.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = messageLog.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (messageLog MessageLog) GetRecords(identifier int) (response Response, err error) {
	if identifier < 1 {
		identifier = 1
	}

	header := messageLog.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_MessageLog, GetRecords), AMT_MessageLog, nil, "", "")
	body := messageLog.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetRecords), AMT_MessageLog, &GetRecords_INPUT{
		IterationIdentifier: identifier,
		MaxReadRecords:      390,
	})

	response = Response{
		Message: &client.Message{
			XMLInput: messageLog.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = messageLog.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (messageLog MessageLog) PositionToFirstRecord() (response Response, err error) {
	header := messageLog.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_MessageLog, PositionToFirstRecord), AMT_MessageLog, nil, "", "")
	body := messageLog.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(PositionToFirstRecord), AMT_MessageLog, nil)
	response = Response{
		Message: &client.Message{
			XMLInput: messageLog.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = messageLog.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
