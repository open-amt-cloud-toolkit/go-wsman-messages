package wsman

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
