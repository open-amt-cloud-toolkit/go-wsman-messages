/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package credential

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

type ContextPullResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  message.Header
	Body    ContextPullResponseBody
}

type ContextPullResponseBody struct {
	PullResponse ContextResponse
}

type ContextResponse struct {
	EnumerationContext string
	Items              []Relationship `xml:"Items>CIM_CredentialContext"`
}

type Relationship struct {
	ElementInContext        models.AssociationReference
	ElementProvidingContext models.AssociationReference
}

type Context struct {
	base message.Base
}

const ClassName = "CIM_CredentialContext"

// NewContext returns a new instance of the NewContext struct.
func NewContext(wsmanMessageCreator *message.WSManMessageCreator) Context {
	return Context{
		base: message.NewBase(wsmanMessageCreator, ClassName),
	}
}

// Get the representation of the instance
func (b Context) Get() string {
	return b.base.Get(nil)
}

// Enumerate the instances of this class
func (b Context) Enumerate() string {
	return b.base.Enumerate()
}

// Pull instances of this class, following an Enumerate operation
func (b Context) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
