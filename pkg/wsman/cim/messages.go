/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package cim

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/bios"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/boot"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/computer"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/concrete"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/credential"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/ieee8021x"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/kvm"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/mediaaccess"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/physical"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/power"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/service"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/software"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/system"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/wifi"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

type Messages struct {
	wsmanMessageCreator       *message.WSManMessageCreator
	BIOSElement               bios.Element
	BootConfigSetting         boot.ConfigSetting
	BootService               boot.Service
	BootSourceSetting         boot.SourceSetting
	Card                      physical.Card
	Chassis                   physical.Chassis
	Chip                      physical.Chip
	ComputerSystemPackage     computer.SystemPackage
	ConcreteDependency        concrete.Dependency
	CredentialContext         credential.Context
	IEEE8021xSettings         ieee8021x.Settings
	KVMRedirectionSAP         kvm.RedirectionSAP
	MediaAccessDevice         mediaaccess.Device
	PhysicalMemory            physical.Memory
	PhysicalPackage           physical.Package
	PowerManagementService    power.ManagementService
	Processor                 physical.Processor
	ServiceAvailableToElement service.AvailableToElement
	SoftwareIdentity          software.Identity
	SystemPackaging           system.Packaging
	WiFiEndpointSettings      wifi.EndpointSettings
	WiFiPort                  wifi.Port
}

func NewMessages(client *client.Client) Messages {
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	m := Messages{
		wsmanMessageCreator: wsmanMessageCreator,
	}
	m.BIOSElement = bios.NewBIOSElement(wsmanMessageCreator)
	m.BootConfigSetting = boot.NewBootConfigSetting(wsmanMessageCreator)
	m.BootService = boot.NewBootService(wsmanMessageCreator)
	m.BootSourceSetting = boot.NewBootSourceSetting(wsmanMessageCreator)
	m.Card = physical.NewCard(wsmanMessageCreator)
	m.Chassis = physical.NewChassis(wsmanMessageCreator)
	m.Chip = physical.NewChip(wsmanMessageCreator)
	m.ComputerSystemPackage = computer.NewComputerSystemPackage(wsmanMessageCreator)
	m.ConcreteDependency = concrete.NewDependency(wsmanMessageCreator)
	m.CredentialContext = credential.NewContext(wsmanMessageCreator)
	m.IEEE8021xSettings = ieee8021x.NewIEEE8021xSettings(wsmanMessageCreator)
	m.KVMRedirectionSAP = kvm.NewKVMRedirectionSAP(wsmanMessageCreator)
	m.MediaAccessDevice = mediaaccess.NewMediaAccessDevice(wsmanMessageCreator)
	m.PhysicalMemory = physical.NewPhysicalMemory(wsmanMessageCreator)
	m.PhysicalPackage = physical.NewPhysicalPackage(wsmanMessageCreator)
	m.PowerManagementService = power.NewPowerManagementService(wsmanMessageCreator)
	m.Processor = physical.NewProcessor(wsmanMessageCreator)
	m.ServiceAvailableToElement = service.NewServiceAvailableToElement(wsmanMessageCreator)
	m.SoftwareIdentity = software.NewSoftwareIdentity(wsmanMessageCreator)
	m.SystemPackaging = system.NewSystemPackaging(wsmanMessageCreator)
	m.WiFiEndpointSettings = wifi.NewWiFiEndpointSettings(wsmanMessageCreator)
	m.WiFiPort = wifi.NewWiFiPort(wsmanMessageCreator)
	return m
}
