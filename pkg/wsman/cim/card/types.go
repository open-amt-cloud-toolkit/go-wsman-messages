/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package card

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Package struct {
	base   message.Base
	client client.WSMan
}

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
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
		PackageResponse   PackageResponse
	}

	PullResponse struct {
		XMLName   xml.Name          `xml:"PullResponse"`
		CardItems []PackageResponse `xml:"Items>CIM_Card"`
	}

	PackageResponse struct {
		XMLName           xml.Name          `xml:"CIM_Card"`
		CanBeFRUed        bool              `xml:"CanBeFRUed"`        // Boolean that indicates whether this PhysicalElement can be FRUed (TRUE) or not (FALSE).
		CreationClassName string            `xml:"CreationClassName"` // CreationClassName indicates the name of the class or the subclass used in the creation of an instance.
		ElementName       string            `xml:"ElementName"`
		Manufacturer      string            `xml:"Manufacturer"`      // The name of the organization responsible for producing the PhysicalElement.
		Model             string            `xml:"Model"`             // The name by which the PhysicalElement is generally known.
		OperationalStatus OperationalStatus `xml:"OperationalStatus"` // Indicates the current statuses of the element
		PackageType       PackageType       `xml:"PackageType"`       // Enumeration defining the type of the PhysicalPackage. Note that this enumeration expands on the list in the Entity MIB (the attribute, entPhysicalClass). The numeric values are consistent with CIM's enum numbering guidelines, but are slightly different than the MIB's values.
		SerialNumber      string            `xml:"SerialNumber"`      // A manufacturer-allocated number used to identify the Physical Element.
		Tag               string            `xml:"Tag"`               // An arbitrary string that uniquely identifies the Physical Element and serves as the key of the Element.
		Version           string            `xml:"Version"`           // A string that indicates the version of the PhysicalElement.
	}
)

// Indicates the current statuses of the element. Various operational statuses are defined. Many of the enumeration's values are self-explanatory. However, a few are not and are described here in more detail:
//
// - "Stressed" indicates that the element is functioning, but needs attention. Examples of "Stressed" states are overload, overheated, and so on.
//
// - "Predictive Failure" indicates that an element is functioning nominally but predicting a failure in the near future.
//
// - "In Service" describes an element being configured, maintained, cleaned, or otherwise administered.
//
// - "No Contact" indicates that the monitoring system has knowledge of this element, but has never been able to establish communications with it.
//
// - "Lost Communication" indicates that the ManagedSystem Element is known to exist and has been contacted successfully in the past, but is currently unreachable.
//
// - "Stopped" and "Aborted" are similar, although the former implies a clean and orderly stop, while the latter implies an abrupt stop where the state and configuration of the element might need to be updated.
//
// - "Dormant" indicates that the element is inactive or quiesced.
//
// - "Supporting Entity in Error" indicates that this element might be "OK" but that another element, on which it is dependent, is in error. An example is a network service or endpoint that cannot function due to lower-layer networking problems.
//
// - "Completed" indicates that the element has completed its operation. This value should be combined with either OK, Error, or Degraded so that a client can tell if the complete operation Completed with OK (passed), Completed with Error (failed), or Completed with Degraded (the operation finished, but it did not complete OK or did not report an error).
//
// - "Power Mode" indicates that the element has additional power model information contained in the Associated PowerManagementService association.
//
// - "Relocating" indicates the element is being relocated.
//
// OperationalStatus replaces the Status property on ManagedSystemElement to provide a consistent approach to enumerations, to address implementation needs for an array property, and to provide a migration path from today's environment to the future. This change was not made earlier because it required the deprecated qualifier. Due to the widespread use of the existing Status property in management applications, it is strongly recommended that providers or instrumentation provide both the Status and OperationalStatus properties. Further, the first value of OperationalStatus should contain the primary status for the element. When instrumented, Status (because it is single-valued) should also provide the primary status of the element.
//
// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, .., 0x8000..}
//
// Values={Unknown, Other, OK, Degraded, Stressed, Predictive Failure, Error, Non-Recoverable Error, Starting, Stopping, Stopped, In Service, No Contact, Lost Communication, Aborted, Dormant, Supporting Entity in Error, Completed, Power Mode, Relocating, DMTF Reserved, Vendor Reserved}
type OperationalStatus int

// Enumeration defining the type of the PhysicalPackage. Note that this enumeration expands on the list in the Entity MIB (the attribute, entPhysicalClass). The numeric values are consistent with CIM's enum numbering guidelines, but are slightly different than the MIB's values.
//
// Unknown - indicates that the package type is not known.
//
// Other - The package type does not correspond to an existing enumerated value. The value is specified using the OtherPackageType property.
//
// - "Rack" through "Port/Connector" are defined per the Entity-MIB (where the semantics of rack are equivalent to the MIB's 'stack' value).  The other values (for battery, processor, memory, power source/generator and storage media package) are self-explanatory.
//
// - "Blade" should be used when the PhysicalPackage contains the operational hardware aspects of a ComputerSystem, without the supporting mechanicals such as power and cooling. For example, a Blade Server includes processor(s) and memory, and relies on the containing chassis to supply power and cooling. In many respects, a Blade can be considered a "Module/Card". However, it is tracked differently by inventory systems and differs in terms of service philosophy. For example, a Blade is intended to be hot-plugged into a hosting enclosure without requiring additional cabling, and does not require a cover to be removed from the enclosure for installation.
//
// - "Blade Expansion" has characteristics of a "Blade" and a "Module/Card". However, it is distinct from both due to inventory tracking and service philosophy, and because of its hardware dependence on a Blade. A Blade Expansion must be attached to a Blade prior to inserting the resultant assembly into an enclosure.
//
// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
//
// Values={Unknown, Other, Rack, Chassis/Frame, Cross Connect/Backplane, Container/Frame Slot, Power Supply, Fan, Sensor, Module/Card, Port/Connector, Battery, Processor, Memory, Power Source/Generator, Storage Media Package (e.g., Disk or Tape Drive), Blade, Blade Expansion}

type PackageType int
