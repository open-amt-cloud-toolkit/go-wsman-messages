/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type ManagementService struct {
	base message.Base
}

type Certificate struct {
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
		XMLName                               xml.Name                         `xml:"Body"`
		AddTrustedRootCertificate_OUTPUT      AddTrustedRootCertificate_OUTPUT `xml:"AddTrustedRootCertificate_OUTPUT,omitempty"`
		AddCertificate_OUTPUT                 AddCertificate_OUTPUT            `xml:"AddCertificate_OUTPUT,omitempty"`
		AddKey_OUTPUT                         AddKey_OUTPUT                    `xml:"AddKey_OUTPUT,omitempty"`
		GenerateKeyPair_OUTPUT                GenerateKeyPair_OUTPUT           `xml:"GenerateKeyPair_OUTPUT,omitempty"`
		KeyManagementGetResponse              KeyManagementResponse            `xml:"AMT_PublicKeyManagementService,omitempty"`
		PublicKeyCertificateGetAndPutResponse PublicKeyCertificateResponse     `xml:"AMT_PublicKeyCertificate,omitempty"`
		EnumerateResponse                     common.EnumerateResponse
		PullResponse                          PullResponse
	}
	PullResponse struct {
		XMLName                   xml.Name                       `xml:"PullResponse,omitempty"`
		KeyManagementItems        []KeyManagementResponse        `xml:"Items>AMT_PublicKeyManagementService,omitempty"`
		PublicKeyCertificateItems []PublicKeyCertificateResponse `xml:"Items>AMT_PublicKeyCertificate,omitempty"`
	}
	KeyManagementResponse struct {
		XMLName                 xml.Name            `xml:"AMT_PublicKeyManagementService,omitempty"`
		CreationClassName       string              `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName             string              `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledDefault          EnabledDefault      `xml:"EnabledDefault"`                    // An enumerated value indicating an administrator's default or startup configuration for the Enabled State of an element. By default, the element is "Enabled" (value=2).
		EnabledState            EnabledState        `xml:"EnabledState"`                      // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		Name                    string              `xml:"Name,omitempty"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		OperationalStatus       []OperationalStatus `xml:"OperationalStatus,omitempty"`       // Indicates the current statuses of the element.
		RequestedState          RequestedState      `xml:"RequestedState"`                    // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		SystemCreationClassName string              `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		SystemName              string              `xml:"SystemName,omitempty"`              // The Name of the scoping System.
	}
	PublicKeyCertificateResponse struct {
		XMLName                xml.Name `xml:"AMT_PublicKeyCertificate,omitempty"`
		ElementName            string   `xml:"ElementName,omitempty"`     // A user-friendly name for the object . . .
		InstanceID             string   `xml:"InstanceID,omitempty"`      // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		X509Certificate        string   `xml:"X509Certificate,omitempty"` // uint8[4100] // The X.509 Certificate blob.
		TrustedRootCertificate bool     `xml:"TrustedRootCertficate"`     // For root certificate [that were added by AMT_PublicKeyManagementService.AddTrustedRootCertificate()]this property will be true. FYI Certificate is spelled wrong comimg from AMT.
		Issuer                 string   `xml:"Issuer,omitempty"`          // The Issuer field of this certificate.
		Subject                string   `xml:"Subject,omitempty"`         // The Subject field of this certificate.
		ReadOnlyCertificate    bool     `xml:"ReadOnlyCertificate"`       // Indicates whether the certificate is an Intel AMT self-signed certificate. If True, the certificate cannot be deleted.
	}
	AddTrustedRootCertificate_OUTPUT struct {
		XMLName            xml.Name                   `xml:"AddTrustedRootCertificate_OUTPUT"`
		CreatedCertificate CreatedCertificateResponse `xml:"CreatedCertificate,omitempty"`
		ReturnValue        int                        `xml:"ReturnValue,omitempty"`
	}
	AddCertificate_OUTPUT struct {
		XMLName            xml.Name                   `xml:"AddCertificate_OUTPUT"`
		CreatedCertificate CreatedCertificateResponse `xml:"CreatedCertificate,omitempty"`
		ReturnValue        int                        `xml:"ReturnValue,omitempty"`
	}
	AddKey_OUTPUT struct {
		XMLName     xml.Name           `xml:"AddKey_OUTPUT,omitempty"`
		CreatedKey  CreatedKeyResponse `xml:"CreatedKey,omitempty"`
		ReturnValue int                `xml:"ReturnValue,omitempty"`
	}
	GenerateKeyPair_OUTPUT struct {
		XMLName     xml.Name        `xml:"GenerateKeyPair_OUTPUT,omitempty"`
		KeyPair     KeyPairResponse `xml:"KeyPair,omitempty"`
		ReturnValue int             `xml:"ReturnValue,omitempty"`
	}
	KeyPairResponse struct {
		XMLName             xml.Name                    `xml:"KeyPair,omitempty"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	CreatedKeyResponse struct {
		XMLName             xml.Name                    `xml:"CreatedKey,omitempty"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	CreatedCertificateResponse struct {
		XMLName             xml.Name                    `xml:"CreatedCertificate,omitempty"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	ReferenceParametersResponse struct {
		XMLName     xml.Name `xml:"ReferenceParameters,omitempty"`
		ResourceURI string
		SelectorSet SelectorSetResponse `xml:"SelectorSet,omitempty"`
	}
	SelectorSetResponse struct {
		XMLName   xml.Name           `xml:"SelectorSet,omitempty"`
		Selectors []SelectorResponse `xml:"Selector"`
	}
	SelectorResponse struct {
		XMLName xml.Name `xml:"Selector,omitempty"`
		Name    string   `xml:"Name,attr"`
		Text    string   `xml:",chardata"`
	}

	// Indicates the current statuses of the element. Various operational statuses are defined. Many of the enumeration's values are self-explanatory. However, a few are not and are described here in more detail.
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
	// OperationalStatus replaces the Status property on ManagedSystemElement to provide a consistent approach to enumerations, to address implementation needs for an array property, and to provide a migration path from today's environment to the future. This change was not made earlier because it required the deprecated qualifier. Due to the widespread use of the existing Status property in management applications, it is strongly recommended that providers or instrumentation provide both the Status and OperationalStatus properties. Further, the first value of OperationalStatus should contain the primary status for the element. When instrumented, Status (because it is single-valued) should also provide the primary status of the element.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, .., 0x8000..}
	//
	// Values={Unknown, Other, OK, Degraded, Stressed, Predictive Failure, Error, Non-Recoverable Error, Starting, Stopping, Stopped, In Service, No Contact, Lost Communication, Aborted, Dormant, Supporting Entity in Error, Completed, Power Mode, Relocating, DMTF Reserved, Vendor Reserved}
	OperationalStatus int
	// An enumerated value indicating an administrator's default or startup configuration for the Enabled State of an element. By default, the element is "Enabled" (value=2).
	//
	// ValueMap={2, 3, 5, 6, 7, 9, .., 32768..65535}
	//
	// Values={Enabled, Disabled, Not Applicable, Enabled but Offline, No Default, Quiesce, DMTF Reserved, Vendor Reserved}
	EnabledDefault int
	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states. For example, shutting down (value=4) and starting (value=10) are transient states between enabled and disabled.
	//
	// The following text briefly summarizes the various enabled and disabled states:
	//
	// Enabled (2) indicates that the element is or could be executing commands, will process any queued commands, and queues new requests.
	//
	// Disabled (3) indicates that the element will not execute commands and will drop any new requests.
	//
	// Shutting Down (4) indicates that the element is in the process of going to a Disabled state.
	//
	// Not Applicable (5) indicates the element does not support being enabled or disabled.
	//
	// Enabled but Offline (6) indicates that the element might be completing commands, and will drop any new requests.
	//
	// Test (7) indicates that the element is in a test state.
	//
	// Deferred (8) indicates that the element might be completing commands, but will queue any new requests.
	//
	// Quiesce (9) indicates that the element is enabled but in a restricted mode.
	//
	// Starting (10) indicates that the element is in the process of going to an Enabled state. New requests are queued.
	//
	// Value 6 ("Enabled but Offline") can be recieved also if the Audit Log is in locked state.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11..32767, 32768..65535}
	//
	// Values={Unknown, Other, Enabled, Disabled, Shutting Down, Not Applicable, Enabled but Offline, In Test, Deferred, Quiesce, Starting, DMTF Reserved, Vendor Reserved}
	EnabledState int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested. The actual state of the element is represented by EnabledState. This property is provided to compare the last requested and current enabled or disabled states. Note that when EnabledState is set to 5 ("Not Applicable"), then this property has no meaning. Refer to the EnabledState property description for explanations of the values in the RequestedState enumeration.
	//
	// "Unknown" (0) indicates the last requested state for the element is unknown.
	//
	// Note that the value "No Change" (5) has been deprecated in lieu of indicating the last requested state is "Unknown" (0). If the last requested or desired state is unknown, RequestedState should have the value "Unknown" (0), but may have the value "No Change" (5).Offline (6) indicates that the element has been requested to transition to the Enabled but Offline EnabledState.	It should be noted that there are two new values in RequestedState that build on the statuses of EnabledState. These are "Reboot" (10) and "Reset" (11). Reboot refers to doing a "Shut Down" and then moving to an "Enabled" state. Reset indicates that the element is first "Disabled" and then "Enabled". The distinction between requesting "Shut Down" and "Disabled" should also be noted. Shut Down requests an orderly transition to the Disabled state, and might involve removing power, to completely erase any existing state. The Disabled state requests an immediate disabling of the element, such that it will not execute or accept any commands or processing requests.
	//
	// This property is set as the result of a method invocation (such as Start or StopService on CIM_Service), or can be overridden and defined as WRITEable in a subclass. The method approach is considered superior to a WRITEable property, because it allows an explicit invocation of the operation and the return of a result code.
	//
	// If knowledge of the last RequestedState is not supported for the EnabledLogicalElement, the property shall be NULL or have the value 12 "Not Applicable".
	//
	// ValueMap={0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, .., 32768..65535}
	//
	// Values={Unknown, Enabled, Disabled, Shut Down, No Change, Offline, Test, Deferred, Quiesce, Reboot, Reset, Not Applicable, DMTF Reserved, Vendor Reserved}
	RequestedState int
)

// INPUTS
// Request Types
type (
	PublicKeyCertificateRequest struct {
		XMLName                xml.Name `xml:"h:AMT_PublicKeyCertificate"`
		H                      string   `xml:"xmlns:h,attr"`
		ElementName            string   `xml:"h:ElementName"`            // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		InstanceID             string   `xml:"h:InstanceID"`             // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		X509Certificate        string   `xml:"h:X509Certificate"`        // uint8[4100] // The X.509 Certificate blob.
		TrustedRootCertificate bool     `xml:"h:TrustedRootCertificate"` // For root certificate [that were added by AMT_PublicKeyManagementService.AddTrustedRootCertificate()]this property will be true.
		Issuer                 string   `xml:"h:Issuer"`                 // The Issuer field of this certificate.
		Subject                string   `xml:"h:Subject"`                // The Subject field of this certificate.
		ReadOnlyCertificate    bool     `xml:"h:ReadOnlyCertificate"`    // Indicates whether the certificate is an Intel AMT self-signed certificate. If True, the certificate cannot be deleted.
	}
	AddCertificate_INPUT struct {
		XMLName         xml.Name `xml:"h:AddCertificate_INPUT"`
		H               string   `xml:"xmlns:h,attr"`
		CertificateBlob string   `xml:"h:CertificateBlob"` // The use of ECC 192/224 is blocked starting from Intel® CSME 18.0.
	}
	AddTrustedRootCertificate_INPUT struct {
		XMLName         xml.Name `xml:"h:AddTrustedRootCertificate_INPUT"`
		H               string   `xml:"xmlns:h,attr"`
		CertificateBlob string   `xml:"h:CertificateBlob"` // The use of ECC 192/224 is blocked starting from Intel® CSME 18.0.
	}
	AddKey_INPUT struct {
		XMLName xml.Name `xml:"h:AddKey_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		KeyBlob string   `xml:"h:KeyBlob"` // The use of ECC 192/224 is blocked starting from Intel® CSME 18.0.
	}
	GenerateKeyPair_INPUT struct {
		XMLName      xml.Name     `xml:"h:GenerateKeyPair_INPUT"`
		H            string       `xml:"xmlns:h,attr"`
		KeyAlgorithm KeyAlgorithm `xml:"h:KeyAlgorithm"` // The algorithm of the generated key.
		KeyLength    KeyLength    `xml:"h:KeyLength"`    // The length of the generatd key in bits.
	}
	PKCS10Request struct {
		XMLName                      xml.Name         `xml:"h:GeneratePKCS10RequestEx_INPUT"`
		H                            string           `xml:"xmlns:h,attr"`
		KeyPair                      string           `xml:"h:KeyPair"`
		NullSignedCertificateRequest string           `xml:"h:NullSignedCertificateRequest"` // A binary representation of the null-signed PKCS#10 request.the request must include a valid PKCS10RequestInfo, that will be signed by AMT FW. The Public Key specified in the request must match the public key of the referenced KeyPair parameter.
		SigningAlgorithm             SigningAlgorithm `xml:"h:SigningAlgorithm"`             // The signing algorithm that the FW should use for signing the certificate request
	}

	// The signing algorithm that the FW should use for signing the certificate request
	//
	// Required
	//
	// ValueMap={0, 1, ..}
	//
	// Values={SHA1-RSA, SHA256-RSA, Reserved} Note: SHA1 is no longer available starting from Intel CSME 18.0.
	SigningAlgorithm int

	// The algorithm of the generated key.
	//
	// Required
	//
	// ValueMap={0, ..}
	//
	// Values={RSA, Reserved}
	KeyAlgorithm int

	// The length of the generatd key in bits
	KeyLength int
)
