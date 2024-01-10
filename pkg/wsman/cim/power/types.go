/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package power

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type ManagementService struct {
	base   message.Base
	client client.WSMan
}

type PowerState int

// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                         xml.Name            `xml:"Body"`
		RequestPowerStateChangeResponse PowerActionResponse `xml:"RequestPowerStateChange_OUTPUT"`
		GetResponse                     PowerManagementService
		EnumerateResponse               common.EnumerateResponse
		PullResponse                    PullResponse
	}

	PullResponse struct {
		XMLName                     xml.Name                 `xml:"PullResponse"`
		PowerManagementServiceItems []PowerManagementService `xml:"Items>CIM_PowerManagementService"`
	}

	PowerManagementService struct {
		XMLName                 xml.Name              `xml:"CIM_PowerManagementService"`
		CreationClassName       string                `xml:"CreationClassName,omitempty"`
		ElementName             string                `xml:"ElementName,omitempty"`
		EnabledState            models.EnabledState   `xml:"EnabledState,omitempty"`
		Name                    string                `xml:"Name,omitempty"`
		RequestedState          models.RequestedState `xml:"RequestedState,omitempty"`
		SystemCreationClassName string                `xml:"SystemCreationClassName,omitempty"`
		SystemName              string                `xml:"SystemName,omitempty"`
	}

	PowerActionResponse struct {
		ReturnValue int `xml:"ReturnValue"`
	}
)
