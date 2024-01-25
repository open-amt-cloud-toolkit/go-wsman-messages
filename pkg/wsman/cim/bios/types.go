/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package bios

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type Element struct {
	base message.Base
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
		GetResponse       BiosElement
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	BiosElement struct {
		XMLName               xml.Name              `xml:"CIM_BIOSElement"`
		TargetOperatingSystem TargetOperatingSystem `xml:"TargetOperatingSystem"` // The TargetOperatingSystem property specifies the element's operating system environment.
		SoftwareElementID     string                `xml:"SoftwareElementID"`     // This is an identifier for the SoftwareElement and is designed to be used in conjunction with other keys to create a unique representation of the element.
		SoftwareElementState  SoftwareElementState  `xml:"SoftwareElementState"`  // The SoftwareElementState is defined in this model to identify various states of a SoftwareElement's life cycle.
		Name                  string                `xml:"Name"`                  // The name used to identify this SoftwareElement.
		OperationalStatus     OperationalStatus     `xml:"OperationalStatus"`     // Indicates the current statuses of the element.
		ElementName           string                `xml:"ElementName"`           // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		Version               string                `xml:"Version"`               // The version of the BIOS software image.
		Manufacturer          string                `xml:"Manufacturer"`          // The manufacturer of the BIOS software image.
		PrimaryBIOS           bool                  `xml:"PrimaryBIOS"`           // If true, this is the primary BIOS of the ComputerSystem.
		ReleaseDate           Time                  `xml:"ReleaseDate"`           // Date that this BIOS was released.
	}

	Time struct {
		DateTime string `xml:"Datetime"`
	}

	PullResponse struct {
		XMLName          xml.Name      `xml:"PullResponse"`
		BiosElementItems []BiosElement `xml:"Items>CIM_BIOSElement"`
	}
)

type (
	// The TargetOperatingSystem property specifies the element's operating system environment.
	// The value of this property does not ensure that it is binary executable.
	// Two other pieces of information are needed.
	// First, the version of the OS needs to be specified using the class, CIM_OSVersion Check.
	// The second piece of information is the architecture that the OS runs on.
	// This information is verified using CIM_ArchitectureCheck.
	// The combination of these constructs clearly identifies the level of OS required for a particular SoftwareElement.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113}
	//
	// Values={Unknown, Other, MACOS, ATTUNIX, DGUX, DECNT, Tru64 UNIX, OpenVMS, HPUX, AIX, MVS, OS400, OS/2, JavaVM, MSDOS, WIN3x, WIN95, WIN98, WINNT, WINCE, NCR3000, NetWare, OSF, DC/OS, Reliant UNIX, SCO UnixWare, SCO OpenServer, Sequent, IRIX, Solaris, SunOS, U6000, ASERIES, HP NonStop OS, HP NonStop OSS, BS2000, LINUX, Lynx, XENIX, VM, Interactive UNIX, BSDUNIX, FreeBSD, NetBSD, GNU Hurd, OS9, MACH Kernel, Inferno, QNX, EPOC, IxWorks, VxWorks, MiNT, BeOS, HP MPE, NextStep, PalmPilot, Rhapsody, Windows 2000, Dedicated, OS/390, VSE, TPF, Windows (R) Me, Caldera Open UNIX, OpenBSD, Not Applicable, Windows XP, z/OS, Microsoft Windows Server 2003, Microsoft Windows Server 2003 64-Bit, Windows XP 64-Bit, Windows XP Embedded, Windows Vista, Windows Vista 64-Bit, Windows Embedded for Point of Service, Microsoft Windows Server 2008, Microsoft Windows Server 2008 64-Bit, FreeBSD 64-Bit, RedHat Enterprise Linux, RedHat Enterprise Linux 64-Bit, Solaris 64-Bit, SUSE, SUSE 64-Bit, SLES, SLES 64-Bit, Novell OES, Novell Linux Desktop, Sun Java Desktop System, Mandriva, Mandriva 64-Bit, TurboLinux, TurboLinux 64-Bit, Ubuntu, Ubuntu 64-Bit, Debian, Debian 64-Bit, Linux 2.4.x, Linux 2.4.x 64-Bit, Linux 2.6.x, Linux 2.6.x 64-Bit, Linux 64-Bit, Other 64-Bit, Microsoft Windows Server 2008 R2, VMware ESXi, Microsoft Windows 7, CentOS 32-bit, CentOS 64-bit, Oracle Enterprise Linux 32-bit, Oracle Enterprise Linux 64-bit, eComStation 32-bitx, Microsoft Windows Server 2011, Microsoft Windows Server 2011 64-Bit, Microsoft Windows Server 8}
	TargetOperatingSystem int

	// The SoftwareElementState is defined in this model to identify various states of a SoftwareElement's life cycle.
	//
	// - A SoftwareElement in the deployable state describes the details necessary to successfully distribute it and the details (Checks and Actions) required to move it to the installable state (i.e, the next state).
	//
	// - A SoftwareElement in the installable state describes the details necessary to successfully install it and the details (Checks and Actions) required to create an element in the executable state (i.e., the next state).
	//
	// - A SoftwareElement in the executable state describes the details necessary to successfully start it and the details (Checks and Actions) required to move it to the running state (i.e., the next state).
	//
	// - A SoftwareElement in the running state describes the details necessary to manage the started element.
	//
	// ValueMap={0, 1, 2, 3}
	//
	// Values={Deployable, Installable, Executable, Running}
	SoftwareElementState int

	// Indicates the current statuses of the element.
	// Various operational statuses are defined.
	// Many of the enumeration's values are self-explanatory.
	// However, a few are not and are described here in more detail.
	//
	// "Stressed" indicates that the element is functioning, but needs attention. Examples of "Stressed" states are overload, overheated, and so on.
	//
	// "Predictive Failure" indicates that an element is functioning nominally but predicting a failure in the near future.
	//
	// "In Service" describes an element being configured, maintained, cleaned, or otherwise administered.
	//
	// "No Contact" indicates that the monitoring system has knowledge of this element, but has never been able to establish communications with it.
	//
	// "Lost Communication" indicates that the ManagedSystem Element is known to exist and has been contacted successfully in the past, but is currently unreachable.
	//
	// "Stopped" and "Aborted" are similar, although the former implies a clean and orderly stop, while the latter implies an abrupt stop where the state and configuration of the element might need to be updated.
	//
	// "Dormant" indicates that the element is inactive or quiesced.
	//
	// "Supporting Entity in Error" indicates that this element might be "OK" but that another element, on which it is dependent, is in error. An example is a network service or endpoint that cannot function due to lower-layer networking problems.
	//
	// "Completed" indicates that the element has completed its operation. This value should be combined with either OK, Error, or Degraded so that a client can tell if the complete operation Completed with OK (passed), Completed with Error (failed), or Completed with Degraded (the operation finished, but it did not complete OK or did not report an error).
	//
	// "Power Mode" indicates that the element has additional power model information contained in the Associated PowerManagementService association.
	//
	// "Relocating" indicates the element is being relocated.
	//
	// OperationalStatus replaces the Status property on ManagedSystemElement to provide a consistent approach to enumerations, to address implementation needs for an array property, and to provide a migration path from today's environment to the future.
	// This change was not made earlier because it required the deprecated qualifier.
	// Due to the widespread use of the existing Status property in management applications, it is strongly recommended that providers or instrumentation provide both the Status and OperationalStatus properties.
	// Further, the first value of OperationalStatus should contain the primary status for the element.
	// When instrumented, Status (because it is single-valued) should also provide the primary status of the element.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, .., 0x8000..}
	//
	// Values={Unknown, Other, OK, Degraded, Stressed, Predictive Failure, Error, Non-Recoverable Error, Starting, Stopping, Stopped, In Service, No Contact, Lost Communication, Aborted, Dormant, Supporting Entity in Error, Completed, Power Mode, Relocating, DMTF Reserved, Vendor Reserved}
	OperationalStatus int
)
