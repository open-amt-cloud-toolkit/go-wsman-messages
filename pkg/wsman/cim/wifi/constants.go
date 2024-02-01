/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

const (
	CIM_WiFiEndpoint         string = "CIM_WiFiEndpoint"
	CIM_WiFiEndpointSettings string = "CIM_WiFiEndpointSettings"
	CIM_WiFiPort             string = "CIM_WiFiPort"
)

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

const (
	AuthenticationMethod_Other AuthenticationMethod = iota + 1
	AuthenticationMethod_OpenSystem
	AuthenticationMethod_SharedKey
	AuthenticationMethod_WPA_PSK
	AuthenticationMethod_WPA_IEEE8021x
	AuthenticationMethod_WPA2_PSK
	AuthenticationMethod_WPA2_IEEE8021x
	AuthenticationMethod_DMTFReserved
)
const (
	AuthenticationMethod_WPA3_SAE AuthenticationMethod = iota + 32768
	AuthenticationMethod_WPA3_OWE
	AuthenticationMethod_VendorReserved
)

const (
	BSSType_Unknown        BSSType = 0
	BSSType_Independent    BSSType = 2
	BSSType_Infrastructure BSSType = 3
)

const (
	EncryptionMethod_Other EncryptionMethod = iota + 1
	EncryptionMethod_WEP
	EncryptionMethod_TKIP
	EncryptionMethod_CCMP
	EncryptionMethod_None
	EncryptionMethod_DMTFReserved
)

const (
	CompletedwithNoError ReturnValue = iota
	NotSupported
	UnknownorUnspecifiedError
	CannotcompletewithinTimeoutPeriod
	Failed
	InvalidParameter
	InUse
)
const (
	MethodParametersCheckedJobStarted ReturnValue = iota + 4096
	InvalidStateTransition
	UseofTimeoutParameterNotSupported
	Busy
)
