/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package common

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
)

type EnumerationResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  message.Header
	Body    EnumerationBody
}

type EnumerationBody struct {
	EnumerateResponse EnumerateResponse
}

type EnumerateResponse struct {
	EnumerationContext string `xml:"EnumerationContext,omitempty"`
}

type ReturnValue struct {
	XMLName        xml.Name `xml:"RequestStateChange_OUTPUT,omitempty"`
	ReturnValue    int      `xml:"ReturnValue,omitempty"`
	ReturnValueStr string   `xml:"ReturnValueStr,omitempty"`
}

// AMT WSMAN Error Response Types.
type (
	AMTError struct {
		SubCode string
		Message string
		Detail  string
	}

	ErrorResponse struct {
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    ErrorBody      `xml:"Body"`
	}

	ErrorBody struct {
		XMLName xml.Name `xml:"Body"`
		Fault   Fault    `xml:"Fault"`
	}

	Fault struct {
		XMLName xml.Name `xml:"Fault"`
		Code    Code     `xml:"Code"`
		Reason  Reason   `xml:"Reason"`
		Detail  string   `xml:"Detail"`
	}

	Code struct {
		XMLName xml.Name `xml:"Code"`
		Value   string   `xml:"Value"`
		SubCode SubCode  `xml:"Subcode"`
	}

	SubCode struct {
		XMLName xml.Name `xml:"Subcode"`
		Value   string   `xml:"Value"`
	}

	Reason struct {
		XMLName xml.Name `xml:"Reason"`
		Text    string   `xml:"Text"`
	}
)
