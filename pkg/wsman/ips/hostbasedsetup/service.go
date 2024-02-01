/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbasedsetup

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

// NewHostBasedSetupService returns a new instance of the HostBasedSetupService struct.
func NewHostBasedSetupServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPS_HostBasedSetupService, client),
	}
}

// Get retrieves the representation of the instance
func (service Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Get(nil),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call
func (service Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Enumerate(),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (service Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Pull(enumerationContext),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Add a certificate to the provisioning certificate chain, to be used by AdminSetup or UpgradeClientToAdmin methods.
func (service Service) AddNextCertInChain(cert string, isLeaf bool, isRoot bool) (response Response, err error) {
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPS_HostBasedSetupService, AddNextCertInChain), IPS_HostBasedSetupService, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(AddNextCertInChain), IPS_HostBasedSetupService, AddNextCertInChain_INPUT{
		H:                 "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		NextCertificate:   cert,
		IsLeafCertificate: isLeaf,
		IsRootCertificate: isRoot,
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Setup Intel® AMT from the local host, resulting in Admin Setup Mode. Requires OS administrator rights, and moves Intel® AMT from "Pre Provisioned" state to "Post Provisioned" state. The control mode after this method is run will be "Admin".
func (service Service) AdminSetup(adminPassEncryptionType AdminPassEncryptionType, digestRealm string, adminPassword string, mcNonce string, signingAlgorithm SigningAlgorithm, digitalSignature string) (response Response, err error) {
	hashInHex := createMD5Hash(adminPassword, digestRealm)
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPS_HostBasedSetupService, AdminSetup), IPS_HostBasedSetupService, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(AdminSetup), IPS_HostBasedSetupService, AdminSetup_INPUT{
		H:                          "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		NetAdminPassEncryptionType: int(adminPassEncryptionType),
		NetworkAdminPassword:       string(hashInHex),
		McNonce:                    mcNonce,
		SigningAlgorithm:           int(signingAlgorithm),
		DigitalSignature:           digitalSignature,
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (service Service) Setup(adminPassEncryptionType AdminPassEncryptionType, digestRealm, adminPassword string) (response Response, err error) {
	hashInHex := createMD5Hash(adminPassword, digestRealm)
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPS_HostBasedSetupService, Setup), IPS_HostBasedSetupService, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(Setup), IPS_HostBasedSetupService, Setup_INPUT{
		H:                          "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		NetAdminPassEncryptionType: int(adminPassEncryptionType),
		NetworkAdminPassword:       string(hashInHex),
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func createMD5Hash(adminPassword string, digestRealm string) string {
	// Create an md5 hash.
	setupPassword := "admin:" + digestRealm + ":" + adminPassword
	hash := md5.New()
	_, _ = io.WriteString(hash, setupPassword)
	hashInHex := fmt.Sprintf("%x", hash.Sum(nil))
	return hashInHex
}

// Upgrade Intel® AMT from Client to Admin Control Mode.
func (service Service) UpgradeClientToAdmin(mcNonce string, signingAlgorithm SigningAlgorithm, digitalSignature string) (response Response, err error) {
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPS_HostBasedSetupService, UpgradeClientToAdmin), IPS_HostBasedSetupService, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(UpgradeClientToAdmin), IPS_HostBasedSetupService, UpgradeClientToAdmin_INPUT{
		H:                "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		McNonce:          mcNonce,
		SigningAlgorithm: int(signingAlgorithm),
		DigitalSignature: digitalSignature,
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
