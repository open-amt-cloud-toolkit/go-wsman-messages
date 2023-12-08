/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

type SourceSetting struct {
	base   message.Base
	client client.WSMan
}

// NewBootSourceSetting returns a new instance of the BootSourceSetting struct.
func NewBootSourceSettingWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SourceSetting {
	return SourceSetting{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_BootSourceSetting, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (SourceSetting SourceSetting) Get(instanceID string) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: instanceID,
	}
	response = Response{
		Message: &client.Message{
			XMLInput: SourceSetting.base.Get(&selector),
		},
	}

	err = SourceSetting.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (SourceSetting SourceSetting) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: SourceSetting.base.Enumerate(),
		},
	}

	err = SourceSetting.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pulls instances of this class, following an Enumerate operation
func (SourceSetting SourceSetting) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: SourceSetting.base.Pull(enumerationContext),
		},
	}
	err = SourceSetting.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
