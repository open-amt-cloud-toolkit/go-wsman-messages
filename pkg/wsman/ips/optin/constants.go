/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package optin

const IPS_OptInService string = "IPS_OptInService"

const (
	False int = 0
	True  int = 1
)

const (
	NotStarted int = iota
	Requested
	Displayed
	Received
	InSession
)

const (
	None uint32 = 0
	KVM  uint32 = 1
	All  uint32 = 4294967295
)

const (
	ReturnValuePTStatusSuccess int = iota
	ReturnValuePTStatusInternalError
	ReturnValuePTStatusInvalidState
	ReturnValuePTStatusBlocked
	ReturnValuePTStatusFailedFFS
)
