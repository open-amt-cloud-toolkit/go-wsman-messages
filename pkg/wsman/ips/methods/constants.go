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
	Setup                   Methods = "Setup"
	AdminSetup              Methods = "AdminSetup"
	StartOptIn              Methods = "StartOptIn"
	CancelOptIn             Methods = "CancelOptIn"
	SendOptInCode           Methods = "SendOptInCode"
	RequestPowerStateChange Methods = "RequestPowerStateChange"
	AddNextCertInChain      Methods = "AddNextCertInChain"
	SetCertificates         Methods = "SetCertificates"
	SetCertificates_INPUT   Methods = "SetCertificates_INPUT"
	UpgradeClientToAdmin    Methods = "UpgradeClientToAdmin"
)
