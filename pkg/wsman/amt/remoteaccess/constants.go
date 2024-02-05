/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

const (
	AMT_RemoteAccessPolicyAppliesToMPS string = "AMT_RemoteAccessPolicyAppliesToMPS"
	AMT_RemoteAccessPolicyRule         string = "AMT_RemoteAccessPolicyRule"
	AMT_RemoteAccessService            string = "AMT_RemoteAccessService"
	AddMps                             string = "AddMpServer"
	AddRemoteAccessPolicyRule          string = "AddRemoteAccessPolicyRule"
)

const (
	IPv4Address MPServerInfoFormat = 3
	IPv6Address MPServerInfoFormat = 4
	FQDN        MPServerInfoFormat = 201
)

const (
	MutualAuthentication           MPServerAuthMethod = 1
	UsernamePasswordAuthentication MPServerAuthMethod = 2
)

const (
	UserInitiated Trigger = iota
	Alert
	Periodic
	HomeProvisioning
)

const (
	PolicyDecisionStrategyFirstMatching PolicyDecisionStrategy = 1
	PolicyDecisionStrategyAll           PolicyDecisionStrategy = 2
)

const (
	ExternalMPS MPSType = iota
	InternalMPS
	BothMPS
)
