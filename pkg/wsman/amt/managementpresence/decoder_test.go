/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package managementpresence

import "testing"

func TestInfoFormat_String(t *testing.T) {
	tests := []struct {
		state    InfoFormat
		expected string
	}{
		{InfoFormatOther, "Other"},
		{InfoFormatHostName, "HostName"},
		{InfoFormatIPv4Address, "IPv4Address"},
		{InfoFormatIPv6Address, "IPv6Address"},
		{InfoFormatIPXAddress, "IPXAddress"},
		{InfoFormatDECnetAddress, "DECnetAddress"},
		{InfoFormatSNAAddress, "SNAAddress"},
		{InfoFormatAutonomousSystemNumber, "AutonomousSystemNumber"},
		{InfoFormatMPLSLabel, "MPLSLabel"},
		{InfoFormatIPv4SubnetAddress, "IPv4SubnetAddress"},
		{InfoFormatIPv6SubnetAddress, "IPv6SubnetAddress"},
		{InfoFormatIPv4AddressRange, "IPv4AddressRange"},
		{InfoFormatIPv6AddressRange, "IPv6AddressRange"},
		{InfoFormatDialString, "DialString"},
		{InfoFormatEthernetAddress, "EthernetAddress"},
		{InfoFormatTokenRingAddress, "TokenRingAddress"},
		{InfoFormatATMAddress, "ATMAddress"},
		{InfoFormatFrameRelayAddress, "FrameRelayAddress"},
		{InfoFormatURL, "URL"},
		{InfoFormatFQDN, "FQDN"},
		{InfoFormatUserFQDN, "UserFQDN"},
		{InfoFormatDERASN1DN, "DERASN1DN"},
		{InfoFormatDERASN1GN, "DERASN1GN"},
		{InfoFormatKeyID, "KeyID"},
		{InfoFormatParameterizedURL, "ParameterizedURL"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
