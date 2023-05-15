/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

import (
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

const AMT_SetupAndConfigurationService = "AMT_SetupAndConfigurationService"

type SetupAndConfigurationService struct {
	models.CredentialManagementService
	AMT_SetupAndConfigurationService struct {
		CreationClassName             string
		ElementName                   string
		EnabledState                  string
		Name                          string
		PasswordModel                 string
		ProvisioningMode              string
		ProvisioningServerOTP         string
		ProvisioningState             string
		RequestedState                string
		SystemCreationClassName       string
		SystemName                    string
		ZeroTouchConfigurationEnabled string
	}
}
type Service struct {
	base wsman.Base
}

func NewSetupAndConfigurationService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{
		base: wsman.NewBase(wsmanMessageCreator, AMT_SetupAndConfigurationService),
	}
}
func (s Service) Get() string {
	return s.base.Get(nil)
}

// Enumerates the instances of this class
func (s Service) Enumerate() string {
	return s.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (s Service) Pull(enumerationContext string) string {
	return s.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (s Service) Put(setupAndConfigurationService SetupAndConfigurationService) string {
	return s.base.Put(setupAndConfigurationService, false, nil)
}
func (s Service) CommitChanges() string {
	header := s.base.WSManMessageCreator.CreateHeader(string(actions.CommitChanges), AMT_SetupAndConfigurationService, nil, "", "")
	body := s.base.WSManMessageCreator.CreateBody("CommitChanges_INPUT", AMT_SetupAndConfigurationService, nil)
	return s.base.WSManMessageCreator.CreateXML(header, body)
}

func (s Service) GetUuid() string {
	header := s.base.WSManMessageCreator.CreateHeader(string(actions.GetUuid), AMT_SetupAndConfigurationService, nil, "", "")
	body := s.base.WSManMessageCreator.CreateBody("GetUuid_INPUT", AMT_SetupAndConfigurationService, nil)
	return s.base.WSManMessageCreator.CreateXML(header, body)
}

func (s Service) SetMEBXPassword(password string) string {
	header := s.base.WSManMessageCreator.CreateHeader(string(actions.SetMEBxPassword), AMT_SetupAndConfigurationService, nil, "", "")
	body := fmt.Sprintf(`<Body><h:SetMEBxPassword_INPUT xmlns:h="%s%s"><h:Password>%s</h:Password></h:SetMEBxPassword_INPUT></Body>`, s.base.WSManMessageCreator.ResourceURIBase, AMT_SetupAndConfigurationService, password)
	return s.base.WSManMessageCreator.CreateXML(header, body)
}

func (s Service) Unprovision(provisioningMode int) string {
	if provisioningMode == 0 {
		provisioningMode = 1
	}
	header := s.base.WSManMessageCreator.CreateHeader("Unprovision", AMT_SetupAndConfigurationService, nil, "", "")
	body := fmt.Sprintf(`<Body><h:Unprovision_INPUT xmlns:h="%s%s"><h:ProvisioningMode>%d</h:ProvisioningMode></h:Unprovision_INPUT></Body>`, s.base.WSManMessageCreator.ResourceURIBase, AMT_SetupAndConfigurationService, provisioningMode)
	return s.base.WSManMessageCreator.CreateXML(header, body)
}
