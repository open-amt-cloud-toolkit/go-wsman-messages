/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package message

import (
	"encoding/xml"
)

type Header struct {
	XMLName     xml.Name `xml:"Header"`
	To          string   `xml:"To"`
	RelatesTo   int      `xml:"RelatesTo"`
	Action      Action   `xml:"Action"`
	MessageID   string   `xml:"MessageID"`
	ResourceURI string   `xml:"ResourceURI"`
}
type Action struct {
	XMLName        xml.Name `xml:"Action"`
	MustUnderstand string   `xml:"mustUnderstand,attr"`
	Value          string   `xml:",chardata"`
}
