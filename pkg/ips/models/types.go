/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package models

import (
	base "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	cim "github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

type CurrentControlMode int

const (
	NotProvisioned CurrentControlMode = 0
	Client         CurrentControlMode = 1
	Admin          CurrentControlMode = 2
)

type CertChainStatus int

const (
	CertChainStatusNotStarted      CertChainStatus = 0
	CertChainStatusChainInProgress CertChainStatus = 1
	CertChainStatusChainComplete   CertChainStatus = 2
)

type OptInRequired int

const (
	None OptInRequired = 0
	KVM  OptInRequired = 1
	All  OptInRequired = 4294967295
)

type OptInState int

const (
	NotStarted OptInState = 0
	Requested  OptInState = 1
	Displayed  OptInState = 2
	Received   OptInState = 3
	InSession  OptInState = 4
)

type CanModifyOptInPolicy int

const (
	FALSE CanModifyOptInPolicy = 0
	TRUE  CanModifyOptInPolicy = 1
)

type HostBasedSetupService struct {
	cim.SecurityService
	CurrentControlMode  CurrentControlMode `json:"currentControlMode,omitempty"`
	AllowedControlModes []int              `json:"allowedControlModes,omitempty"`
	ConfigurationNonce  []int              `json:"configurationNonce,omitempty"`
	CertChainStatus     CertChainStatus    `json:"certChainStatus,omitempty"`
}

type OptInService struct {
	cim.Service
	OptInCodeTimeout     int                  `json:"optInCodeTimeout,omitempty"`
	OptInRequired        OptInRequired        `json:"optInRequired,omitempty"`
	OptInState           OptInState           `json:"optInState,omitempty"`
	CanModifyOptInPolicy CanModifyOptInPolicy `json:"canModifyOptInPolicy,omitempty"`
	OptInDisplayTimeout  int                  `json:"optInDisplayTimeout,omitempty"`
}

type OptInServiceResponse struct {
	IPS_OptInService OptInService `json:"IPS_OptInService"`
}

type StartOptIn_OUTPUT struct {
	StartOptIn_OUTPUT base.ReturnValue `json:"StartOptIn_OUTPUT"`
}

type CancelOptIn_OUTPUT struct {
	CancelOptIn_OUTPUT base.ReturnValue `json:"CancelOptIn_OUTPUT"`
}

type SendOptInCode_OUTPUT struct {
	SendOptInCode_OUTPUT base.ReturnValue `json:"SendOptInCode_OUTPUT"`
}
