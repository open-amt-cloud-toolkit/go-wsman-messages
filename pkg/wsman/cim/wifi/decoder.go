/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

const (
	CIMWiFiEndpoint         string = "CIM_WiFiEndpoint"
	CIMWiFiEndpointSettings string = "CIM_WiFiEndpointSettings"
	CIMWiFiPort             string = "CIM_WiFiPort"
	ValueNotFound           string = "Value not found in map"
)

const (
	AuthenticationMethodOther AuthenticationMethod = iota + 1
	AuthenticationMethodOpenSystem
	AuthenticationMethodSharedKey
	AuthenticationMethodWPAPSK
	AuthenticationMethodWPAIEEE8021x
	AuthenticationMethodWPA2PSK
	AuthenticationMethodWPA2IEEE8021x
	AuthenticationMethodWPA3SAE AuthenticationMethod = 32768
	AuthenticationMethodWPA3OWE AuthenticationMethod = 32769
)

// authenticationMethodMap is a map of the AuthenticationMethod enumeration.
var authenticationMethodMap = map[AuthenticationMethod]string{
	AuthenticationMethodOther:         "Other",
	AuthenticationMethodOpenSystem:    "OpenSystem",
	AuthenticationMethodSharedKey:     "SharedKey",
	AuthenticationMethodWPAPSK:        "WPAPSK",
	AuthenticationMethodWPAIEEE8021x:  "WPAIEEE8021x",
	AuthenticationMethodWPA2PSK:       "WPA2PSK",
	AuthenticationMethodWPA2IEEE8021x: "WPA2IEEE8021x",
	AuthenticationMethodWPA3SAE:       "WPA3SAE",
	AuthenticationMethodWPA3OWE:       "WPA3OWE",
}

// String returns a human-readable string representation of the AuthenticationMethod enumeration.
func (e AuthenticationMethod) String() string {
	if s, ok := authenticationMethodMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	BSSTypeUnknown        BSSType = 0
	BSSTypeIndependent    BSSType = 2
	BSSTypeInfrastructure BSSType = 3
)

// bssTypeMap is a map of the BSSType enumeration.
var bssTypeMap = map[BSSType]string{
	BSSTypeUnknown:        "Unknown",
	BSSTypeIndependent:    "Independent",
	BSSTypeInfrastructure: "Infrastructure",
}

// String returns a human-readable string representation of the BSSType enumeration.
func (e BSSType) String() string {
	if s, ok := bssTypeMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	EnabledStateWifiDisabled      EnabledState = 3
	EnabledStateWifiEnabledS0     EnabledState = 32768
	EnabledStateWifiEnabledS0SxAC EnabledState = 32769
)

// enabledStateMap is a map of the EnabledState enumeration.
var enabledStateMap = map[EnabledState]string{
	EnabledStateWifiDisabled:      "WifiDisabled",
	EnabledStateWifiEnabledS0:     "WifiEnabledS0",
	EnabledStateWifiEnabledS0SxAC: "WifiEnabledS0SxAC",
}

// String returns a human-readable string representation of the EnabledState enumeration.
func (e EnabledState) String() string {
	if s, ok := enabledStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	EncryptionMethod_Other EncryptionMethod = iota + 1
	EncryptionMethod_WEP
	EncryptionMethod_TKIP
	EncryptionMethod_CCMP
	EncryptionMethod_None
)

// encryptionMethodMap is a map of the EncryptionMethod enumeration.
var encryptionMethodMap = map[EncryptionMethod]string{
	EncryptionMethod_Other: "Other",
	EncryptionMethod_WEP:   "WEP",
	EncryptionMethod_TKIP:  "TKIP",
	EncryptionMethod_CCMP:  "CCMP",
	EncryptionMethod_None:  "None",
}

// String returns a human-readable string representation of the EncryptionMethod enumeration.
func (e EncryptionMethod) String() string {
	if s, ok := encryptionMethodMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	HealthStateUnknown             HealthState = 0
	HealthStateOK                  HealthState = 5
	HealthStateDegraded            HealthState = 10
	HealthStateMinorFailure        HealthState = 15
	HealthStateMajorFailure        HealthState = 20
	HealthStateCriticalFailure     HealthState = 25
	HealthStateNonRecoverableError HealthState = 30
)

// healthStateMap is a map of the HealthState enumeration.
var healthStateMap = map[HealthState]string{
	HealthStateOK:                  "OK",
	HealthStateDegraded:            "Degraded",
	HealthStateMinorFailure:        "MinorFailure",
	HealthStateMajorFailure:        "MajorFailure",
	HealthStateCriticalFailure:     "CriticalFailure",
	HealthStateNonRecoverableError: "NonRecoverableError",
}

// String returns a human-readable string representation of the HealthState enumeration.
func (e HealthState) String() string {
	if s, ok := healthStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	LinkTechnologyUnknown LinkTechnology = iota
	LinkTechnologyOther
	LinkTechnologyEthernet
	LinkTechnologyIB
	LinkTechnologyFC
	LinkTechnologyFDDI
	LinkTechnologyATM
	LinkTechnologyTokenRing
	LinkTechnologyFrameRelay
	LinkTechnologyInfrared
	LinkTechnologyBlueTooth
	LinkTechnologyWirelessLAN
)

// linkTechnologyMap is a map of the LinkTechnology enumeration.
var linkTechnologyMap = map[LinkTechnology]string{
	LinkTechnologyUnknown:     "Unknown",
	LinkTechnologyOther:       "Other",
	LinkTechnologyEthernet:    "Ethernet",
	LinkTechnologyIB:          "IB",
	LinkTechnologyFC:          "FC",
	LinkTechnologyFDDI:        "FDDI",
	LinkTechnologyATM:         "ATM",
	LinkTechnologyTokenRing:   "TokenRing",
	LinkTechnologyFrameRelay:  "FrameRelay",
	LinkTechnologyInfrared:    "Infrared",
	LinkTechnologyBlueTooth:   "BlueTooth",
	LinkTechnologyWirelessLAN: "WirelessLAN",
}

// String returns a human-readable string representation of the LinkTechnology enumeration.
func (e LinkTechnology) String() string {
	if s, ok := linkTechnologyMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	RequestedStateWifiDisabled      RequestedState = 3
	RequestedStateWifiEnabledS0     RequestedState = 32768
	RequestedStateWifiEnabledS0SxAC RequestedState = 32769
)

// requestedStateMap is a map of the RequestedState enumeration.
var requestedStateMap = map[RequestedState]string{
	RequestedStateWifiDisabled:      "WifiDisabled",
	RequestedStateWifiEnabledS0:     "WifiEnabledS0",
	RequestedStateWifiEnabledS0SxAC: "WifiEnabledS0SxAC",
}

// String returns a human-readable string representation of the RequestedState enumeration.
func (e RequestedState) String() string {
	if s, ok := requestedStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	CompletedWithNoError ReturnValue = iota
	NotSupported
	UnknownOrUnspecifiedError
	CannotCompleteWithinTimeoutPeriod
	Failed
	InvalidParameter
	InUse
	MethodParametersCheckedJobStarted ReturnValue = 4096
	InvalidStateTransition            ReturnValue = 4097
	UseOfTimeoutParameterNotSupported ReturnValue = 4098
	Busy                              ReturnValue = 4099
)

// returnValueMap is a map of the ReturnValue enumeration.
var returnValueMap = map[ReturnValue]string{
	CompletedWithNoError:              "CompletedWithNoError",
	NotSupported:                      "NotSupported",
	UnknownOrUnspecifiedError:         "UnknownOrUnspecifiedError",
	CannotCompleteWithinTimeoutPeriod: "CannotCompleteWithinTimeoutPeriod",
	Failed:                            "Failed",
	InvalidParameter:                  "InvalidParameter",
	InUse:                             "InUse",
	MethodParametersCheckedJobStarted: "MethodParametersCheckedJobStarted",
	InvalidStateTransition:            "InvalidStateTransition",
	UseOfTimeoutParameterNotSupported: "UseOfTimeoutParameterNotSupported",
	Busy:                              "Busy",
}

// String returns a human-readable string representation of the ReturnValue enumeration.
func (e ReturnValue) String() string {
	if s, ok := returnValueMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	PortTypeUnknown PortType = 0
	PortTypeOther   PortType = 1
	PortType80211a  PortType = 70
	PortType80211b  PortType = 71
	PortType80211g  PortType = 72
	PortType80211n  PortType = 73
)

// portTypeMap is a map of the PortType enumeration.
var portTypeMap = map[PortType]string{
	PortTypeUnknown: "Unknown",
	PortTypeOther:   "Other",
	PortType80211a:  "802.11a",
	PortType80211b:  "802.11b",
	PortType80211g:  "802.11g",
	PortType80211n:  "802.11n",
}

// String returns a human-readable string representation of the PortType enumeration.
func (e PortType) String() string {
	if s, ok := portTypeMap[e]; ok {
		return s
	}

	return ValueNotFound
}
