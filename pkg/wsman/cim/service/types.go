/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package service

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type AvailableToElement struct {
	base   message.Base
	client client.WSMan
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
		PullResponse                     PullResponse
		EnumerateResponse                common.EnumerateResponse
		AssociatedPowerManagementService CIM_AssociatedPowerManagementService `xml:"CIM_AssociatedPowerManagementService"`
	}
	PullResponse struct {
		XMLName                          xml.Name                               `xml:"PullResponse"`
		AssociatedPowerManagementService []CIM_AssociatedPowerManagementService `xml:"Items>CIM_AssociatedPowerManagementService"`
	}
	CIM_AssociatedPowerManagementService struct {
		AvailableRequestedPowerStates int `xml:"AvailableReuestedPowerStates,omitempty"`
		PowerState                    int `xml:"PowerState,omitempty"`
		ServiceProvided               ServiceProvided
		UserOfService                 UserOfService
	}
	ServiceProvided struct {
		XMLName             xml.Name `xml:"ServiceProvided,omitempty"`
		Address             string   `xml:"Address"`
		ReferenceParameters ReferenceParameters
	}
	UserOfService struct {
		XMLName             xml.Name `xml:"UserOfService,omitempty"`
		Address             string   `xml:"Address"`
		ReferenceParameters ReferenceParameters
	}
	ReferenceParameters struct {
		XMLName      xml.Name `xml:"ReferenceParameters,omitempty"`
		ResourceURI  string   `xml:"ResourceURI,omitempty"`
		SelelctorSet message.SelectorSet
	}
)
