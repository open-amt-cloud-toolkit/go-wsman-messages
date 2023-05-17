/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

type IEEE8021xProfile struct {
	ElementName                     string
	InstanceID                      string
	Enabled                         bool
	ActiveInS0                      bool
	AuthenticationProtocol          AuthenticationProtocol
	RoamingIdentity                 string
	ServerCertificateName           string
	ServerCertificateNameComparison ServerCertificateNameComparison
	Username                        string
	Password                        string
	Domain                          string
	ProtectedAccessCredential       []int
	PACPassword                     string
	ClientCertificate               string
	ServerCertificateIssue          string
	PxeTimeout                      int
}

type AuthenticationProtocol int

const AMT_IEEE8021xProfile = "AMT_8021XProfile"

const (
	TLS AuthenticationProtocol = iota
	TTLS_MSCHAPv2
	PEAP_MSCHAPv2
	EAP_GTC
	EAPFAST_MSCHAPv2
	EAPFAST_GTC
	EAPFAST_TLS
)

type ServerCertificateNameComparison int

const (
	FullName ServerCertificateNameComparison = iota
	DomainSuffix
)

type Profile struct {
	base wsman.Base
}

func NewIEEE8021xProfile(wsmanMessageCreator *wsman.WSManMessageCreator) Profile {
	return Profile{
		base: wsman.NewBase(wsmanMessageCreator, AMT_IEEE8021xProfile),
	}
}

// Get retrieves the representation of the instance
func (IEEE8021xProfile Profile) Get() string {
	return IEEE8021xProfile.base.Get(nil)
}

// Enumerates the instances of this class
func (IEEE8021xProfile Profile) Enumerate() string {
	return IEEE8021xProfile.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (IEEE8021xProfile Profile) Pull(enumerationContext string) string {
	return IEEE8021xProfile.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (IEEE8021xProfile Profile) Put(ieee8021xProfile IEEE8021xProfile) string {
	return IEEE8021xProfile.base.Put(ieee8021xProfile, false, nil)
}
