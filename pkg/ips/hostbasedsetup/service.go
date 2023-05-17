/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbasedsetup

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/ips/actions"
)

type AdminPassEncryptionType int

const (
	AdminPassEncryptionTypeNone AdminPassEncryptionType = iota
	AdminPassEncryptionTypeOther
	AdminPassEncryptionTypeHTTPDigestMD5A1
)

type SigningAlgorithm int

const (
	SigningAlgorithmNone SigningAlgorithm = iota
	SigningAlgorithmOther
	SigningAlgorithmRSASHA2256
)

type Service struct {
	base wsman.Base
}

const IPS_HostBasedSetupService = "IPS_HostBasedSetupService"

// NewHostBasedSetupService returns a new instance of the HostBasedSetupService struct.
func NewHostBasedSetupService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{
		base: wsman.NewBase(wsmanMessageCreator, string(IPS_HostBasedSetupService)),
	}
}

// Get retrieves the representation of the instance
func (b Service) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b Service) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b Service) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}

type AddNextCertInChain struct {
	XMLName           xml.Name `xml:"h:AddNextCertInChain_INPUT"`
	H                 string   `xml:"xmlns:h,attr"`
	NextCertificate   string   `xml:"h:NextCertificate"`
	IsLeafCertificate bool     `xml:"h:IsLeafCertificate"`
	IsRootCertificate bool     `xml:"h:IsRootCertificate"`
}

// Add a certificate to the provisioning certificate chain, to be used by AdminSetup or UpgradeClientToAdmin methods.
func (b Service) AddNextCertInChain(cert string, isLeaf bool, isRoot bool) string {
	header := b.base.WSManMessageCreator.CreateHeader(string(actions.AddNextCertInChain), string(IPS_HostBasedSetupService), nil, "", "")
	body := b.base.WSManMessageCreator.CreateBody("AddNextCertInChain_INPUT", string(IPS_HostBasedSetupService), AddNextCertInChain{
		H:                 "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		NextCertificate:   cert,
		IsLeafCertificate: isLeaf,
		IsRootCertificate: isRoot,
	})
	return b.base.WSManMessageCreator.CreateXML(header, body)
}

type AdminSetup struct {
	XMLName                    xml.Name `xml:"h:AdminSetup_INPUT"`
	H                          string   `xml:"xmlns:h,attr"`
	NetAdminPassEncryptionType int      `xml:"h:NetAdminPassEncryptionType"`
	NetworkAdminPassword       string   `xml:"h:NetworkAdminPassword"`
	McNonce                    string   `xml:"h:McNonce"`
	SigningAlgorithm           int      `xml:"h:SigningAlgorithm"`
	DigitalSignature           string   `xml:"h:DigitalSignature"`
}

// Setup Intel(R) AMT from the local host, resulting in Admin Setup Mode. Requires OS administrator rights, and moves Intel(R) AMT from "Pre Provisioned" state to "Post Provisioned" state. The control mode after this method is run will be "Admin".
func (b Service) AdminSetup(adminPassEncryptionType AdminPassEncryptionType, adminPassword string, mcNonce string, signingAlgorithm SigningAlgorithm, digitalSignature string) string {
	header := b.base.WSManMessageCreator.CreateHeader(string(actions.AdminSetup), string(IPS_HostBasedSetupService), nil, "", "")
	body := b.base.WSManMessageCreator.CreateBody("AdminSetup_INPUT", string(IPS_HostBasedSetupService), AdminSetup{
		H:                          "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		NetAdminPassEncryptionType: int(adminPassEncryptionType),
		NetworkAdminPassword:       adminPassword,
		McNonce:                    mcNonce,
		SigningAlgorithm:           int(signingAlgorithm),
		DigitalSignature:           digitalSignature,
	})
	return b.base.WSManMessageCreator.CreateXML(header, body)
}

type Setup struct {
	XMLName                    xml.Name `xml:"h:Setup_INPUT"`
	H                          string   `xml:"xmlns:h,attr"`
	NetAdminPassEncryptionType int      `xml:"h:NetAdminPassEncryptionType"`
	NetworkAdminPassword       string   `xml:"h:NetworkAdminPassword"`
}

func (b Service) Setup(adminPassEncryptionType AdminPassEncryptionType, adminPassword string) string {
	header := b.base.WSManMessageCreator.CreateHeader(string(actions.Setup), string(IPS_HostBasedSetupService), nil, "", "")
	body := b.base.WSManMessageCreator.CreateBody("Setup_INPUT", string(IPS_HostBasedSetupService), Setup{
		H:                          "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		NetAdminPassEncryptionType: int(adminPassEncryptionType),
		NetworkAdminPassword:       adminPassword,
	})
	return b.base.WSManMessageCreator.CreateXML(header, body)
}
