/*********************************************************************
 * Copyright (c) Intel Corporation 2024
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package amterror

import (
	"encoding/xml"
	"testing"
)

func TestDecodeWSMANError(t *testing.T) {
	tests := []struct {
		input    string
		expected error
	}{
		{
			"<?xml version=\"1.0\" encoding=\"UTF-8\"?><a:Envelope xmlns:g=\"http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd\" xmlns:f=\"http://schemas.xmlsoap.org/ws/2004/08/eventing\" xmlns:e=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns:d=\"http://schemas.xmlsoap.org/ws/2004/09/transfer\" xmlns:c=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" xmlns:b=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:a=\"http://www.w3.org/2003/05/soap-envelope\" xmlns:h=\"http://schemas.xmlsoap.org/ws/2005/02/trust\" xmlns:i=\"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand=\"true\">http://schemas.xmlsoap.org/ws/2004/08/addressing/fault</b:Action><b:MessageID>uuid:00000000-8086-8086-8086-000000000061</b:MessageID></a:Header><a:Body><a:Fault><a:Code><a:Value>a:Sender</a:Value><a:Subcode><a:Value>b:DestinationUnreachable</a:Value></a:Subcode></a:Code><a:Reason><a:Text xml:lang=\"en-US\">No route can be determined to reach the destination role defined by the WSAddressing To.</a:Text></a:Reason><a:Detail></a:Detail></a:Fault></a:Body></a:Envelope>",
			NewAMTError("b:DestinationUnreachable", "No route can be determined to reach the destination role defined by the WSAddressing To.", ""),
		},
		{
			"bad xml",
			xml.Unmarshal([]byte("bad xml"), &ErrorResponse{}),
		},
	}

	for _, test := range tests {
		result := DecodeAMTErrorString(test.input)
		if result == nil {
			if test.expected != nil {
				t.Errorf("Expected %s, but got nil", test.expected)
			}
		} else {
			if result.Error() != test.expected.Error() {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		}
	}
}
