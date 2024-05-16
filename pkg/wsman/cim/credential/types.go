/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package credential

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Context struct {
	base message.Base
}

// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name `xml:"Body"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
	}

	PullResponse struct {
		XMLName       xml.Name `xml:"PullResponse"`
		Items         Items    `xml:"Items"`
		EndOfSequence xml.Name `xml:"EndOfSequence"`
	}

	Items struct {
		CredentialContext      []CredentialContext      `xml:"CIM_CredentialContext"`
		CredentialContextTLS   []CredentialContextTLS   `xml:"AMT_TLSCredentialContext"`
		CredentialContext8021x []CredentialContext8021x `xml:"IPS_8021xCredentialContext"`
	}
	CredentialContext struct {
		ElementInContext        models.AssociationReference `xml:"ElementInContext"`        // A Credential whose context is defined.
		ElementProvidingContext models.AssociationReference `xml:"ElementProvidingContext"` // The ManagedElement that provides context or scope for the Credential.
	}
	CredentialContextTLS struct {
		ElementInContext        models.AssociationReference `xml:"ElementInContext"`        // A Credential whose context is defined.
		ElementProvidingContext models.AssociationReference `xml:"ElementProvidingContext"` // The ManagedElement that provides context or scope for the Credential.
	}
	CredentialContext8021x struct {
		ElementInContext        models.AssociationReference `xml:"ElementInContext"`        // A Credential whose context is defined.
		ElementProvidingContext models.AssociationReference `xml:"ElementProvidingContext"` // The ManagedElement that provides context or scope for the Credential.
	}
)
