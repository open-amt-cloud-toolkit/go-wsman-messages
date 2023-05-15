/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package authorization

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
)

type AddUserAclEntry struct {
	XMLName          xml.Name         `xml:"h:AddUserAclEntryEx_INPUT"`
	H                string           `xml:"xmlns:h,attr"`
	Handle           int              `xml:"h:Handle,omitempty"`
	DigestUsername   string           `xml:"h:DigestUsername"`
	DigestPassword   string           `xml:"h:DigestPassword"`
	AccessPermission AccessPermission `xml:"h:AccessPermission"`
	Realms           []RealmValues    `xml:"h:Realms>h:RealmValue,omitempty"`
	KerberosUserSid  string           `xml:"h:KerberosUserSid"`
}
type UpdateUserAclEntry struct {
	XMLName          xml.Name         `xml:"h:UpdateUserAclEntry_INPUT"`
	H                string           `xml:"xmlns:h,attr"`
	Handle           int              `xml:"h:Handle,omitempty"`
	DigestUsername   string           `xml:"h:DigestUsername"`
	DigestPassword   string           `xml:"h:DigestPassword"`
	AccessPermission AccessPermission `xml:"h:AccessPermission"`
	Realms           []RealmValues    `xml:"h:Realms>h:RealmValue,omitempty"`
	KerberosUserSid  string           `xml:"h:KerberosUserSid"`
}

type AccessPermission int

const (
	LocalAccessOnly AccessPermission = iota
	NetworkAccessOnly
	LocalAndNetworkAccess
)

type RealmValues int

const AMT_AuthorizationService = "AMT_AuthorizationService"

const (
	RedirectionRealm                RealmValues = 2
	PTAdministrationRealm           RealmValues = 3
	HardwareAssetRealm              RealmValues = 4
	RemoteControlRealm              RealmValues = 5
	StorageRealm                    RealmValues = 6
	EventManagerRealm               RealmValues = 7
	StorageAdminRealm               RealmValues = 8
	AgentPresenceLocalRealm         RealmValues = 9
	AgentPresenceRemoteRealm        RealmValues = 10
	CircuitBreakerRealm             RealmValues = 11
	NetworkTimeRealm                RealmValues = 12
	GeneralInfoRealm                RealmValues = 13
	EndpointAccessControlRealm      RealmValues = 17
	EndpointAccessControlAdminRealm RealmValues = 18
	EventLogReaderRealm             RealmValues = 19
	AuditLogRealm                   RealmValues = 20
	ACLRealm                        RealmValues = 21
	LocalSystemRealm                RealmValues = 24
)

type AuthorizationService struct {
	base wsman.Base
}
type EnumerateUserAclEntries_INPUT struct {
	XMLName    xml.Name `xml:"h:EnumerateUserAclEntries_INPUT"`
	H          string   `xml:"xmlns:h,attr"`
	StartIndex int      `xml:"h:StartIndex"`
}

type GetAclEnabledState_INPUT struct {
	XMLName xml.Name `xml:"h:GetAclEnabledState_INPUT"`
	H       string   `xml:"xmlns:h,attr"`
	Handle  int      `xml:"h:Handle"`
}
type GetUserAclEntryEx_INPUT struct {
	XMLName xml.Name `xml:"h:GetUserAclEntryEx_INPUT"`
	H       string   `xml:"xmlns:h,attr"`
	Handle  int      `xml:"h:Handle"`
}
type RemoveUserAclEntry_INPUT struct {
	XMLName xml.Name `xml:"h:RemoveUserAclEntry_INPUT"`
	H       string   `xml:"xmlns:h,attr"`
	Handle  int      `xml:"h:Handle"`
}

type SetAclEnabledState_INPUT struct {
	XMLName xml.Name `xml:"h:SetAclEnabledState_INPUT"`
	H       string   `xml:"xmlns:h,attr"`
	Handle  int      `xml:"h:Handle"`
	Enabled bool     `xml:"h:Enabled"`
}

type SetAdminACLEntryEx_INPUT struct {
	XMLName        xml.Name `xml:"h:SetAdminACLEntryEx_INPUT"`
	H              string   `xml:"xmlns:h,attr"`
	Username       string   `xml:"h:Username"`
	DigestPassword string   `xml:"h:DigestPassword"`
}

func NewAuthorizationService(wsmanMessageCreator *wsman.WSManMessageCreator) AuthorizationService {
	return AuthorizationService{
		base: wsman.NewBase(wsmanMessageCreator, AMT_AuthorizationService),
	}
}

// Get retrieves the representation of the instance
func (AuthorizationService AuthorizationService) Get() string {
	return AuthorizationService.base.Get(nil)
}
func (AuthorizationService AuthorizationService) Enumerate() string {
	return AuthorizationService.base.Enumerate()
}
func (AuthorizationService AuthorizationService) Pull(enumerationContext string) string {
	return AuthorizationService.base.Pull(enumerationContext)
}

// func (as AuthorizationService) AddUserAclEntryEx(accessPermission AccessPermission, realms []RealmValues, digestUsername, digestPassword, kerberosUserSid string) string {
// 	if (digestUsername == "" || digestPassword == "") && kerberosUserSid == "" {
// 		panic("Missing user ACL entry information")
// 	}
// 	if digestUsername != "" && len(digestUsername) > 16 {
// 		panic("Username too long")
// 	}
// 	header := as.base.WSManMessageCreator.CreateHeader(string(actions.AddUserAclEntryEx), AMT_AuthorizationService, nil, "", "")
// 	aclObject := AddUserAclEntry{
// 		DigestUsername:   digestUsername,
// 		DigestPassword:   digestPassword,
// 		KerberosUserSid:  kerberosUserSid,
// 		AccessPermission: accessPermission,
// 		Realms:           realms,
// 	}
// 	body := as.base.WSManMessageCreator.CreateBody("AddUserAclEntryEx_INPUT", AMT_AuthorizationService, &aclObject)
// 	return as.base.WSManMessageCreator.CreateXML(header, body)
// }

func (as AuthorizationService) EnumerateUserAclEntries(startIndex int) string {
	if startIndex == 0 {
		startIndex = 1
	}
	header := as.base.WSManMessageCreator.CreateHeader(string(actions.EnumerateUserAclEntries), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody("EnumerateUserAclEntries_INPUT", AMT_AuthorizationService, &EnumerateUserAclEntries_INPUT{StartIndex: startIndex})
	return as.base.WSManMessageCreator.CreateXML(header, body)
}

func (as AuthorizationService) GetAclEnabledState(handle int) string {
	header := as.base.WSManMessageCreator.CreateHeader(string(actions.GetAclEnabledState), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody("GetAclEnabledState_INPUT", AMT_AuthorizationService, &GetAclEnabledState_INPUT{Handle: handle})
	return as.base.WSManMessageCreator.CreateXML(header, body)
}

func (as AuthorizationService) GetAdminAclEntry() string {
	header := as.base.WSManMessageCreator.CreateHeader(string(actions.GetAdminAclEntry), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody("GetAdminAclEntry_INPUT", AMT_AuthorizationService, nil)
	return as.base.WSManMessageCreator.CreateXML(header, body)
}

func (as AuthorizationService) GetAdminAclEntryStatus() string {
	header := as.base.WSManMessageCreator.CreateHeader(string(actions.GetAdminAclEntryStatus), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody("GetAdminAclEntryStatus_INPUT", AMT_AuthorizationService, nil)
	return as.base.WSManMessageCreator.CreateXML(header, body)
}

func (as AuthorizationService) GetAdminNetAclEntryStatus() string {
	header := as.base.WSManMessageCreator.CreateHeader(string(actions.GetAdminNetAclEntryStatus), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody("GetAdminNetAclEntryStatus_INPUT", AMT_AuthorizationService, nil)
	return as.base.WSManMessageCreator.CreateXML(header, body)
}
func (as AuthorizationService) GetUserAclEntryEx(handle int) string {
	header := as.base.WSManMessageCreator.CreateHeader(string(actions.GetUserAclEntryEx), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody("GetUserAclEntryEx_INPUT", AMT_AuthorizationService, &GetUserAclEntryEx_INPUT{Handle: handle})
	return as.base.WSManMessageCreator.CreateXML(header, body)
}

func (as AuthorizationService) RemoveUserAclEntry(handle int) string {
	header := as.base.WSManMessageCreator.CreateHeader(string(actions.RemoveUserAclEntry), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody("RemoveUserAclEntry_INPUT", AMT_AuthorizationService, &RemoveUserAclEntry_INPUT{Handle: handle})
	return as.base.WSManMessageCreator.CreateXML(header, body)
}

func (as AuthorizationService) SetAclEnabledState(handle int, enabled bool) string {
	header := as.base.WSManMessageCreator.CreateHeader(string(actions.SetAclEnabledState), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody("SetAclEnabledState_INPUT", AMT_AuthorizationService, &SetAclEnabledState_INPUT{Handle: handle, Enabled: enabled})
	return as.base.WSManMessageCreator.CreateXML(header, body)
}

func (as AuthorizationService) SetAdminACLEntryEx(username, digestPassword string) string {
	header := as.base.WSManMessageCreator.CreateHeader(string(actions.SetAdminAclEntryEx), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody("SetAdminAclEntryEx_INPUT", AMT_AuthorizationService, &SetAdminACLEntryEx_INPUT{Username: username, DigestPassword: digestPassword})
	return as.base.WSManMessageCreator.CreateXML(header, body)
}

// func (as AuthorizationService) UpdateUserAclEntryEx(handle int, accessPermission AccessPermission, realms []RealmValues, digestUsername, digestPassword, kerberosUserSid string) string {
// 	if (digestUsername == "" || digestPassword == "") && kerberosUserSid == "" {
// 		panic("Missing user ACL entry information")
// 	}
// 	if digestUsername != "" && len(digestUsername) > 16 {
// 		panic("Username too long")
// 	}
// 	header := as.base.WSManMessageCreator.CreateHeader(string(actions.UpdateUserAclEntryEx), AMT_AuthorizationService, nil, "", "")
// 	aclObject := &UpdateUserAclEntry{
// 		Handle:           handle,
// 		DigestUsername:   digestUsername,
// 		DigestPassword:   digestPassword,
// 		KerberosUserSid:  kerberosUserSid,
// 		AccessPermission: accessPermission,
// 		Realms:           realms,
// 	}
// 	body := as.base.WSManMessageCreator.CreateBody("UpdateUserAclEntryEx_INPUT", AMT_AuthorizationService, aclObject)
// 	return as.base.WSManMessageCreator.CreateXML(header, body)
// }
