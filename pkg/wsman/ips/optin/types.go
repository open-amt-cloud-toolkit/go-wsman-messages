/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package optin

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Service struct {
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
		EnumerateResponse     common.EnumerateResponse
		GetResponse           OptInServiceResponse `xml:"IPS_OptInService"`
		PullResponse          PullResponse         `xml:"PullResponse"`
		StartOptInResponse    StartOptIn_OUTPUT    `xml:"StartOptIn_OUTPUT"`
		CancelOptInResponse   CancelOptIn_OUTPUT   `xml:"CancelOptIn_OUTPUT"`
		SendOptInCodeResponse SendOptInCode_OUTPUT `xml:"SendOptInCode_OUTPUT"`
	}

	OptInServiceResponse struct {
		XMLName                 xml.Name `xml:"IPS_OptInService"`
		Name                    string
		CreationClassName       string
		SystemName              string
		SystemCreationClassName string
		ElementName             string
		OptInCodeTimeout        int
		OptInRequired           int
		OptInState              int
		CanModifyOptInPolicy    int
		OptInDisplayTimeout     int
	}

	PullResponse struct {
		Items []OptInServiceResponse `xml:"Items>IPS_OptInService"`
	}
	StartOptIn_OUTPUT struct {
		XMLName     xml.Name `xml:"StartOptIn_OUTPUT"`
		ReturnValue int
	}

	CancelOptIn_OUTPUT struct {
		XMLName     xml.Name `xml:"CancelOptIn_OUTPUT"`
		ReturnValue int
	}

	SendOptInCode_OUTPUT struct {
		XMLName     xml.Name `xml:"SendOptInCode_OUTPUT"`
		ReturnValue int
	}
)

// INPUT
// Request Types
type (
	OptInCode struct {
		XMLName   xml.Name `xml:"h:SendOptInCode_INPUT"`
		H         string   `xml:"xmlns:h,attr"`
		OptInCode int      `xml:"h:OptInCode"`
	}
)

type OptInService struct {
	XMLName                 xml.Name             `xml:"h:IPS_OptInService"`
	H                       string               `xml:"xmlns:h,attr"`
	Name                    string               `xml:"h:Name,omitempty"`
	CreationClassName       string               `xml:"h:CreationClassName,omitempty"`
	SystemName              string               `xml:"h:SystemName,omitempty"`
	SystemCreationClassName string               `xml:"h:SystemCreationClassName,omitempty"`
	ElementName             string               `xml:"h:ElementName,omitempty"`
	OptInCodeTimeout        int                  `xml:"h:OptInCodeTimeout,omitempty"`
	OptInRequired           OptInRequired        `xml:"h:OptInRequired,omitempty"`
	OptInState              OptInState           `xml:"h:OptInState,omitempty"`
	CanModifyOptInPolicy    CanModifyOptInPolicy `xml:"h:CanModifyOptInPolicy,omitempty"`
	OptInDisplayTimeout     int                  `xml:"h:OptInDisplayTimeout,omitempty"`
}

type OptInRequired int
type OptInState int
type CanModifyOptInPolicy int
type ReturnValue int
