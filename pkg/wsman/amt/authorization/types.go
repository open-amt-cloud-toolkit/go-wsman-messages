/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package authorization

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type AuthorizationService struct {
	base message.Base
}

// OUTPUTS
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName           xml.Name `xml:"Body"`
		GetResponse       AuthorizationOccurrence
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	AuthorizationOccurrence struct {
		XMLName                 xml.Name       `xml:"AMT_AuthorizationService"`
		AllowHttpQopAuthOnly    int            `xml:"AllowHttpQopAuthOnly"`
		CreationClassName       string         `xml:"CreationClassName"`
		ElementName             string         `xml:"ElementName"`
		EnabledState            EnabledState   `xml:"EnabledState"`
		Name                    string         `xml:"Name"`
		RequestedState          RequestedState `xml:"RequestedState"`
		SystemCreationClassName string         `xml:"SystemCreationClassName"`
		SystemName              string         `xml:"SystemName"`
	}
	PullResponse struct {
		XMLName                      xml.Name                  `xml:"PullResponse"`
		AuthorizationOccurrenceItems []AuthorizationOccurrence `xml:"Items>AMT_AuthorizationService"`
	}
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
type RealmValues int
type EnabledState int
type RequestedState int
type PTStatus int

// INPUTS
type (
	EnumerateUserAclEntries_INPUT struct {
		XMLName    xml.Name `xml:"h:EnumerateUserAclEntries_INPUT"`
		H          string   `xml:"xmlns:h,attr"`
		StartIndex int      `xml:"h:StartIndex"`
	}

	GetAclEnabledState_INPUT struct {
		XMLName xml.Name `xml:"h:GetAclEnabledState_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"`
	}
	GetUserAclEntryEx_INPUT struct {
		XMLName xml.Name `xml:"h:GetUserAclEntryEx_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"`
	}
	RemoveUserAclEntry_INPUT struct {
		XMLName xml.Name `xml:"h:RemoveUserAclEntry_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"`
	}

	SetAclEnabledState_INPUT struct {
		XMLName xml.Name `xml:"h:SetAclEnabledState_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"`
		Enabled bool     `xml:"h:Enabled"`
	}

	SetAdminACLEntryEx_INPUT struct {
		XMLName        xml.Name `xml:"h:SetAdminACLEntryEx_INPUT"`
		H              string   `xml:"xmlns:h,attr"`
		Username       string   `xml:"h:Username"`
		DigestPassword string   `xml:"h:DigestPassword"`
	}
)
