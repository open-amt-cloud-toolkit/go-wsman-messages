package common

import (
	"encoding/xml"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

type EnumerationResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  wsman.Header
	Body    EnumerationBody
}

type EnumerationBody struct {
	EnumerateResponse EnumerateResponse
}

type EnumerateResponse struct {
	EnumerationContext string
}
