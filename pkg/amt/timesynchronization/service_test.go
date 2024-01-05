/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_TimeSynchronizationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewTimeSynchronizationService(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_TimeSynchronizationService Get wsman message", "AMT_TimeSynchronizationService", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_TimeSynchronizationService Enumerate wsman message", "AMT_TimeSynchronizationService", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_TimeSynchronizationService Pull wsman message", "AMT_TimeSynchronizationService", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			{"should create a valid AMT_TimeSynchronizationService GetLowAccuracyTimeSynch wsman message", "AMT_TimeSynchronizationService", string(actions.GetLowAccuracyTimeSynch), `<h:GetLowAccuracyTimeSynch_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService"></h:GetLowAccuracyTimeSynch_INPUT>`, elementUnderTest.GetLowAccuracyTimeSynch},
			{"should create a valid AMT_TimeSynchronizationService SetHighAccuracyTimeSynch wsman message", "AMT_TimeSynchronizationService", string(actions.SetHighAccuracyTimeSynch), "<h:SetHighAccuracyTimeSynch_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService\"><h:Ta0>1644240911</h:Ta0><h:Tm1>1644240943</h:Tm1><h:Tm2>1644240943</h:Tm2></h:SetHighAccuracyTimeSynch_INPUT>", func() string { return elementUnderTest.SetHighAccuracyTimeSynch(1644240911, 1644240943, 1644240943) }},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				correctResponse := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response := test.responseFunc()
				if response != correctResponse {
					assert.Equal(t, correctResponse, response)
				}
			})
		}
	})

	t.Run("should parse GetLowAccuracyTimeSynch response", func(t *testing.T) {
		// change the return value from 0 as that is default and doesn't
		// prove that xml was parsed correctly
		xmlWithErr := strings.Replace(getLowAccuracyTimeSynchXMLResponse,
			`<g:ReturnValue>0</g:ReturnValue>`,
			`<g:ReturnValue>1</g:ReturnValue>`, 1)
		var rsp Response
		err := xml.Unmarshal([]byte(xmlWithErr), &rsp)
		assert.Nil(t, err)
		assert.Equal(t, int64(1704394160), rsp.Body.GetLowAccuracyTimeSynch_OUTPUT.Ta0)
		assert.Equal(t, 1, rsp.Body.GetLowAccuracyTimeSynch_OUTPUT.ReturnValue)
	})

	t.Run("should parse SetHighAccuracyTimeSynch response", func(t *testing.T) {
		// change the return value from 0 as that is default and doesn't
		// prove that xml was parsed correctly
		xmlWithErr := strings.Replace(setHighAccuracyTimeSynchXMLResponse,
			`<g:ReturnValue>0</g:ReturnValue>`,
			`<g:ReturnValue>1</g:ReturnValue>`, 1)
		var rsp Response
		err := xml.Unmarshal([]byte(xmlWithErr), &rsp)
		assert.Nil(t, err)
		assert.Equal(t, 1, rsp.Body.SetHighAccuracyTimeSynch_OUTPUT.ReturnValue)
	})
}

const getLowAccuracyTimeSynchXMLResponse = `<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope"
            xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing"
            xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"
            xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust"
            xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
            xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd"
            xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <a:Header>
        <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
        <b:RelatesTo>0</b:RelatesTo>
        <b:Action a:mustUnderstand="true">
            http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService/GetLowAccuracyTimeSynchResponse
        </b:Action>
        <b:MessageID>uuid:00000000-8086-8086-8086-000000011E1F</b:MessageID>
        <c:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService</c:ResourceURI>
    </a:Header>
    <a:Body>
        <g:GetLowAccuracyTimeSynch_OUTPUT>
            <g:Ta0>1704394160</g:Ta0>
            <g:ReturnValue>0</g:ReturnValue>
        </g:GetLowAccuracyTimeSynch_OUTPUT>
    </a:Body>
</a:Envelope>`

const setHighAccuracyTimeSynchXMLResponse = `
<?xml version="1.0" encoding="UTF-8"?>
<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing"
            xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"
            xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust"
            xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
            xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd"
            xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <a:Header>
        <b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To>
        <b:RelatesTo>9</b:RelatesTo>
        <b:Action a:mustUnderstand="true">
            http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService/SetHighAccuracyTimeSynchResponse
        </b:Action>
        <b:MessageID>uuid:00000000-8086-8086-8086-000000000061</b:MessageID>
        <c:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService</c:ResourceURI>
    </a:Header>
    <a:Body>
        <g:SetHighAccuracyTimeSynch_OUTPUT>
            <g:ReturnValue>0</g:ReturnValue>
        </g:SetHighAccuracyTimeSynch_OUTPUT>
    </a:Body>
</a:Envelope>
`
