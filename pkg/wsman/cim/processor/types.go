/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package processor

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Package struct {
	base   message.Base
	client client.WSMan
}

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name `xml:"Body"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
		PackageResponse   PackageResponse
	}

	PullResponse struct {
		XMLName      xml.Name          `xml:"PullResponse"`
		PackageItems []PackageResponse `xml:"Items>CIM_Processor"`
	}

	PackageResponse struct {
		XMLName                 xml.Name                 `xml:"CIM_Processor"`
		DeviceID                string                   `xml:"DeviceID,omitempty"`
		CreationClassName       string                   `xml:"CreationClassName,omitempty"`
		SystemName              string                   `xml:"SystemName,omitempty"`
		SystemCreationClassName string                   `xml:"SystemCreationClassName,omitempty"`
		ElementName             string                   `xml:"ElementName,omitempty"`
		OperationalStatus       models.OperationalStatus `xml:"OperationalStatus,omitempty"`
		HealthState             models.HealthState       `xml:"HealthState,omitempty"`
		EnabledState            models.EnabledState      `xml:"EnabledState,omitempty"`
		RequestedState          models.RequestedState    `xml:"RequestedState,omitempty"`
		Role                    string                   `xml:"Role,omitempty"`
		Family                  int                      `xml:"Family,omitempty"`
		OtherFamilyDescription  string                   `xml:"OtherFamilyDescription,omitempty"`
		UpgradeMethod           models.UpgradeMethod     `xml:"UpgradeMethod,omitempty"`
		MaxClockSpeed           int                      `xml:"MaxClockSpeed,omitempty"`
		CurrentClockSpeed       int                      `xml:"CurrentClockSpeed,omitempty"`
		Stepping                string                   `xml:"Stepping,omitempty"`
		CPUStatus               models.CPUStatus         `xml:"CPUStatus,omitempty"`
		ExternalBusClockSpeed   int                      `xml:"ExternalBusClockSpeed,omitempty"`
	}
)
