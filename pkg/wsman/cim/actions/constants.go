/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package actions holds generic action constants for various AMT commands
package actions

type Actions string

const (
	Enumerate               Actions = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate"
	Pull                    Actions = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull"
	Get                     Actions = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	Put                     Actions = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Put"
	Delete                  Actions = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete"
	Create                  Actions = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Create"
	SetBootConfigRole       Actions = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService/SetBootConfigRole"
	ChangeBootOrder         Actions = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting/ChangeBootOrder"
	RequestPowerStateChange Actions = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService/RequestPowerStateChange"
)
