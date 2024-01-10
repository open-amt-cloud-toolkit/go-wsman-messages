/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package environmentdetection

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type SettingData struct {
	base message.Base
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
		XMLName           xml.Name `xml:"Body"`
		GetAndPutResponse EnvironmentDetectionSettingDataResponse
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}
	EnvironmentDetectionSettingDataResponse struct {
		XMLName                    xml.Name           `xml:"AMT_EnvironmentDetectionSettingData"`
		ElementName                string             `xml:"ElementName,omitempty"`
		InstanceID                 string             `xml:"InstanceID,omitempty"`
		DetectionAlgorithm         DetectionAlgorithm `xml:"DetectionAlgorithm,omitempty"`
		DetectionStrings           []string           `xml:"DetectionStrings,omitempty"`
		DetectionIPv6LocalPrefixes []string           `xml:"DetectionIPv6LocalPrefixes,omitempty"`
	}
	PullResponse struct {
		XMLName                              xml.Name                                  `xml:"PullResponse"`
		EnvironmentDetectionSettingDataItems []EnvironmentDetectionSettingDataResponse `xml:"Items>AMT_EnvironmentDetectionSettingData"`
	}
)

// INPUTS
// Request Types
type EnvironmentDetectionSettingDataRequest struct {
	XMLName                    xml.Name           `xml:"h:AMT_EnvironmentDetectionSettingData"`
	H                          string             `xml:"xmlns:h,attr"`
	ElementName                string             `xml:"h:ElementName"`
	InstanceID                 string             `xml:"h:InstanceID"`
	DetectionAlgorithm         DetectionAlgorithm `xml:"h:DetectionAlgorithm"`
	DetectionStrings           []string           `xml:"h:DetectionStrings,omitempty"`
	DetectionIPv6LocalPrefixes []string           `xml:"h:DetectionIPv6LocalPrefixes,omitempty"`
}
type DetectionAlgorithm int
