/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package optin

const IPS_OptInService string = "IPS_OptInService"

const (
	FALSE CanModifyOptInPolicy = 0
	TRUE  CanModifyOptInPolicy = 1
)

const (
	NotStarted OptInState = iota
	Requested
	Displayed
	Received
	InSession
)

const (
	None OptInRequired = 0
	KVM  OptInRequired = 1
	All  OptInRequired = 4294967295
)

const (
	ReturnValuePTStatusSuccess ReturnValue = iota
	ReturnValuePTStatusInternalError
	ReturnValuePTStatusInvalidState
	ReturnValuePTStatusBlocked
	ReturnValuePTStatusFailedFFS
)
