/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kvm

const CIM_KVMRedirectionSAP string = "CIM_KVMRedirectionSAP"

const (
	KVMRedirectionSAPEEnabledInput  KVMRedirectionSAPRequestedStateInputs = 2
	KVMRedirectionSAPEDisabledInput KVMRedirectionSAPRequestedStateInputs = 3
	KVMRedirectionSAPRShutDownInput KVMRedirectionSAPRequestedStateInputs = 4
	KVMRedirectionSAPOfflineInput   KVMRedirectionSAPRequestedStateInputs = 6
	KVMRedirectionSAPRTestInput     KVMRedirectionSAPRequestedStateInputs = 7
	KVMRedirectionSAPEDeferredInput KVMRedirectionSAPRequestedStateInputs = 8
	KVMRedirectionSAPEQuiesceInput  KVMRedirectionSAPRequestedStateInputs = 9
	KVMRedirectionSAPRRebootInput   KVMRedirectionSAPRequestedStateInputs = 10
	KVMRedirectionSAPRResetInput    KVMRedirectionSAPRequestedStateInputs = 11
)
const (
	KVMRedirectionSAPKUnknown KVMRedirectionSAPKVMProtocol = iota
	KVMRedirectionSAPKOther
	KVMRedirectionSAPKRaw
	KVMRedirectionSAPKRDP
	KVMRedirectionSAPKVNC_RFB
	KVMRedirectionSAPDMTFReserved
	KVMRedirectionSAPVendorSpecified
)
const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
	EnabledStateDMTFReserved
	EnabledStateVendorReserved
)
