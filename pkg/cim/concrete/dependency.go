/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package concrete

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

type DependencyPullResponseEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  message.Header
	Body    DependencyPullResponseBody
}

type DependencyPullResponseBody struct {
	PullResponse DependencyResponse
}

type DependencyResponse struct {
	EnumerationContext string
	Items              []Relationship `xml:"Items>CIM_ConcreteDependency"`
}

type Relationship struct {
	Antecedent models.AssociationReference
	Dependent  models.AssociationReference
}

type Dependency struct {
	base message.Base
}

const ClassName = "CIM_ConcreteDependency"

// NewDependency returns a new instance of the NewDependency struct.
// should be NewDependency() because concrete is scoped already as package name.
func NewDependency(wsmanMessageCreator *message.WSManMessageCreator) Dependency {
	return Dependency{
		base: message.NewBase(wsmanMessageCreator, ClassName),
	}
}

// Get the representation of the instance
func (b Dependency) Get() string {
	return b.base.Get(nil)
}

// Enumerate the instances of this class
func (b Dependency) Enumerate() string {
	return b.base.Enumerate()
}

// Pull instances of this class, following an Enumerate operation
func (b Dependency) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
