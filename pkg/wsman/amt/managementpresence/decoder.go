/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package managementpresence

const (
	AMTManagementPresenceRemoteSAP string = "AMT_ManagementPresenceRemoteSAP"
)

const (
	InfoFormatOther InfoFormat = iota + 1
	InfoFormatHostName
	InfoFormatIPv4Address
	InfoFormatIPv6Address
	InfoFormatIPXAddress
	InfoFormatDECnetAddress
	InfoFormatSNAAddress
	InfoFormatAutonomousSystemNumber
	InfoFormatMPLSLabel
	InfoFormatIPv4SubnetAddress
	InfoFormatIPv6SubnetAddress
	InfoFormatIPv4AddressRange
	InfoFormatIPv6AddressRange
	InfoFormatDialString        InfoFormat = 100
	InfoFormatEthernetAddress   InfoFormat = 101
	InfoFormatTokenRingAddress  InfoFormat = 102
	InfoFormatATMAddress        InfoFormat = 103
	InfoFormatFrameRelayAddress InfoFormat = 104
	InfoFormatURL               InfoFormat = 200
	InfoFormatFQDN              InfoFormat = 201
	InfoFormatUserFQDN          InfoFormat = 202
	InfoFormatDERASN1DN         InfoFormat = 203
	InfoFormatDERASN1GN         InfoFormat = 204
	InfoFormatKeyID             InfoFormat = 205
	InfoFormatParameterizedURL  InfoFormat = 206
)

// infoFormatString is a map of InfoFormat to their string representation.
var infoFormatString = map[InfoFormat]string{
	InfoFormatOther:                  "Other",
	InfoFormatHostName:               "HostName",
	InfoFormatIPv4Address:            "IPv4Address",
	InfoFormatIPv6Address:            "IPv6Address",
	InfoFormatIPXAddress:             "IPXAddress",
	InfoFormatDECnetAddress:          "DECnetAddress",
	InfoFormatSNAAddress:             "SNAAddress",
	InfoFormatAutonomousSystemNumber: "AutonomousSystemNumber",
	InfoFormatMPLSLabel:              "MPLSLabel",
	InfoFormatIPv4SubnetAddress:      "IPv4SubnetAddress",
	InfoFormatIPv6SubnetAddress:      "IPv6SubnetAddress",
	InfoFormatIPv4AddressRange:       "IPv4AddressRange",
	InfoFormatIPv6AddressRange:       "IPv6AddressRange",
	InfoFormatDialString:             "DialString",
	InfoFormatEthernetAddress:        "EthernetAddress",
	InfoFormatTokenRingAddress:       "TokenRingAddress",
	InfoFormatATMAddress:             "ATMAddress",
	InfoFormatFrameRelayAddress:      "FrameRelayAddress",
	InfoFormatURL:                    "URL",
	InfoFormatFQDN:                   "FQDN",
	InfoFormatUserFQDN:               "UserFQDN",
	InfoFormatDERASN1DN:              "DERASN1DN",
	InfoFormatDERASN1GN:              "DERASN1GN",
	InfoFormatKeyID:                  "KeyID",
	InfoFormatParameterizedURL:       "ParameterizedURL",
}

// String returns a string representation of a InfoFormatString value.
func (i InfoFormat) String() string {
	if value, exists := infoFormatString[i]; exists {
		return value
	}

	return "Value not found in map"
}
