/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package methods

type Methods string

const (
	Get                     Methods = "Get"
	Pull                    Methods = "Pull"
	Enumerate               Methods = "Enumerate"
	Put                     Methods = "Put"
	Delete                  Methods = "Delete"
	SetBootConfigRole       Methods = "SetBootConfigRole"
	ChangeBootOrder         Methods = "ChangeBootOrder"
	RequestPowerStateChange Methods = "RequestPowerStateChange"
	RequestStateChange      Methods = "RequestStateChange"
)
