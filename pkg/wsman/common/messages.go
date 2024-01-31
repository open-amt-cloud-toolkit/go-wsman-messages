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
