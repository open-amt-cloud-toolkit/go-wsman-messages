/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package managementpresence

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type RemoteSAP struct {
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
		GetResponse       ManagementRemoteResponse
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	PullResponse struct {
		XMLName               xml.Name                   `xml:"PullResponse"`
		ManagementRemoteItems []ManagementRemoteResponse `xml:"Items>AMT_ManagementPresenceRemoteSAP"`
	}

	ManagementRemoteResponse struct {
		XMLName                 xml.Name   `xml:"AMT_ManagementPresenceRemoteSAP"`
		AccessInfo              string     `xml:"AccessInfo,omitempty"`              // Access or addressing information or a combination of this information for a remote connection. This information can be a host name, network address, or similar information.
		CN                      string     `xml:"CN,omitempty"`                      // A common name used when AccessInfo is an IP address.
		CreationClassName       string     `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName             string     `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		InfoFormat              InfoFormat `xml:"InfoFormat,omitempty"`              // An enumerated integer that describes the format and interpretation of the AccessInfo property.
		Name                    string     `xml:"Name,omitempty"`                    // The Name property uniquely identifies the ServiceAccessPoint and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		Port                    int        `xml:"Port,omitempty"`                    // The port to be used to establish a tunnel with the MPS.
		SystemCreationClassName string     `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		SystemName              string     `xml:"SystemName,omitempty"`              // The Name of the scoping System.
	}

	// An enumerated integer that describes the format and interpretation of the AccessInfo property. 206'Parameterized URL'- a URL containing ${parameterName} strings. Those strings are intended to be replaced in their entirety by the value of the named parameter. The interpretation of such parameters is not defined by this subclass. As an example use: If a parameter named 'CompanyURL' has a value of 'www.DMTF.org' and the value of AccessInfo was 'http:\${CompanyURL}', then the resultant URL is intended to be 'http:\www.dmtf.org'.
	//
	// The supported values are 3 (IPv4 address), 4 (IPv6 address) ,201 (FQDN).
	//
	// ValueMap={1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 100, 101, 102, 103, 104, 200, 201, 202, 203, 204, 205, 206, .., 32768..65535}
	//
	// Values={Other, Host Name, IPv4 Address, IPv6 Address, IPX Address, DECnet Address, SNA Address, Autonomous System Number, MPLS Label, IPv4 Subnet Address, IPv6 Subnet Address, IPv4 Address Range, IPv6 Address Range, Dial String, Ethernet Address, Token Ring Address, ATM Address, Frame Relay Address, URL, FQDN, User FQDN, DER ASN1 DN, DER ASN1 GN, Key ID, Parameterized URL, DMTF Reserved, Vendor Reserved}
	InfoFormat int
)
