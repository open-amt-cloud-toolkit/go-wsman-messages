/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

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
		XMLName                      xml.Name `xml:"Body"`
		ProfileGetAndPutResponse     ProfileResponse
		CredentialContextGetResponse CredentialContextResponse
		EnumerateResponse            common.EnumerateResponse
		PullResponse                 PullResponse
	}

	PullResponse struct {
		XMLName                xml.Name                    `xml:"PullResponse"`
		ProfileItems           []ProfileResponse           `xml:"Items>AMT_8021XProfile"`
		CredentialContextItems []CredentialContextResponse `xml:"Items>AMT_8021xCredentialContext"`
	}

	ProfileResponse struct {
		XMLName                         xml.Name                        `xml:"AMT_8021XProfile"`
		ElementName                     string                          `xml:"ElementName,omitempty"`
		InstanceID                      string                          `xml:"InstanceID,omitempty"`
		Enabled                         bool                            `xml:"Enabled,omitempty"`
		ActiveInS0                      bool                            `xml:"ActiveInS0,omitempty"`
		AuthenticationProtocol          AuthenticationProtocol          `xml:"AuthenticationProtocol,omitempty"`
		RoamingIdentity                 string                          `xml:"RoamingIdentity,omitempty"`
		ServerCertificateName           string                          `xml:"ServerCertificateName,omitempty"`
		ServerCertificateNameComparison ServerCertificateNameComparison `xml:"ServerCertificateNameComparison,omitempty"`
		Username                        string                          `xml:"Username,omitempty"`
		Password                        string                          `xml:"Password,omitempty"`
		Domain                          string                          `xml:"Domain,omitempty"`
		ProtectedAccessCredential       []int                           `xml:"ProtectedAccessCredential,omitempty"`
		PACPassword                     string                          `xml:"PACPassword,omitempty"`
		ClientCertificate               string                          `xml:"ClientCertificate,omitempty"`
		ServerCertificateIssue          string                          `xml:"ServerCertificateIssue,omitempty"`
		PxeTimeout                      int                             `xml:"PxeTimeout,omitempty"`
	}

	CredentialContextResponse struct {
		XMLName xml.Name `xml:"AMT_8021XCredentialContext"`
	}
)

type ProfileRequest struct {
	XMLName                         xml.Name                        `xml:"h:AMT_8021XProfile"`
	H                               string                          `xml:"xmlns:h,attr"`
	ElementName                     string                          `xml:"h:ElementName"` // Required
	InstanceID                      string                          `xml:"h:InstanceID,omitempty"`
	Enabled                         bool                            `xml:"h:Enabled"`    // Required
	ActiveInS0                      bool                            `xml:"h:ActiveInS0"` //Functionality: when FALSE, AMT is not accessible (over 802.1x enabled port) in case the host is in S0 but fails to authenticate to the server.  When TRUE, AMT handles the authentication in this case (but the host still can't be accessed until it authenticates successfully).  If 802.1X is not configured, this API may still succeed as the setting may be stored for future use.  The default factory setting is TRUE.
	AuthenticationProtocol          AuthenticationProtocol          `xml:"h:AuthenticationProtocol"`
	RoamingIdentity                 string                          `xml:"h:RoamingIdentity,omitempty"`       // This string, if defined, is sent in response to 802.1x "request identity" as clear text. If empty, the username is sent.
	ServerCertificateName           string                          `xml:"h:ServerCertificateName,omitempty"` // This field is optional. If not defined, the name is not checked. The authenticity of the certificate is always verified.
	ServerCertificateNameComparison ServerCertificateNameComparison `xml:"h:ServerCertificateNameComparison"` // Required if ServerCertificateName is defined
	Username                        string                          `xml:"h:Username,omitempty"`
	Password                        string                          `xml:"h:Password,omitempty"`
	Domain                          string                          `xml:"h:Domain,omitempty"`                    // The Domain string shouldn't contain the suffix, so the user name (Domain\user) will be correct.  If the Domain string contains a suffix (e.g. Domain = intel.com), the user trying to authenticate will be of the form intel.com\user (instead of intel\user) and thus authentication will fail
	ProtectedAccessCredential       []int                           `xml:"h:ProtectedAccessCredential,omitempty"` // This field is relevant for EAP-FAST only. It is not required if the server is configured for "PAC provisioning".
	PACPassword                     string                          `xml:"h:PACPassword,omitempty"`               // Optional password to extract the PAC (Protected Access Credential)information from the PAC data.
	ClientCertificate               string                          `xml:"h:ClientCertificate,omitempty"`         // The client certificate should be specified in the Put request while configuring the profile. This will delete the existing instance of AMT_8021xCredentialContext that represents the client certificate, and create a new instance if a client certificate EPR is provided.  This property will never be returned in Get response.
	ServerCertificateIssue          string                          `xml:"h:ServerCertificateIssue,omitempty"`    // The root certificate should be specified in the Put request while configuring the profile. This will delete the existing instance of AMT_8021xCredentialContext that represents the trusted root certificate, and create a new instance if a trusted root certificate EPR is provided.  This property will never be returned in Get response.  This field is optional. If not defined, AMT looks for a matching root certidicate in its repository.
	PxeTimeout                      int                             `xml:"h:PxeTimeout,omitempty"`                // Timeout in seconds, in which the Intel(R) AMT will hold an authenticated 802.1X session. During the defined period, Intel(R) AMT manages the 802.1X negotiation while a PXE boot takes place. After the timeout, control of the negotiation passes to the host.  The maximum value is 86400 seconds (one day).  A value of 0 disables the feature.  If you do not set a profile, the value of PxeTimeout is 0. If you set a profile without specifying a value for PxeTimeout, the firmware sets it to 120
}

type AuthenticationProtocol int
type ServerCertificateNameComparison int
