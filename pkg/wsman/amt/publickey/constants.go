/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package publickey


import (
	"encoding/json"
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

const (
	AMT_PublicKeyCertificate       string = "AMT_PublicKeyCertificate"
	AMT_PublicKeyManagementService string = "AMT_PublicKeyManagementService"
	GeneratePKCS10RequestEx        string = "GeneratePKCS10RequestEx"
	AddTrustedRootCertificate      string = "AddTrustedRootCertificate"
	AddCertificate                 string = "AddCertificate"
	GenerateKeyPair                string = "GenerateKeyPair"
	AddKey                         string = "AddKey"
)

type (
	Response struct {
		*client.Message
		XMLName  xml.Name       `xml:"Envelope"`
		Header   message.Header `xml:"Header"`
		Body 	 Body		    `xml:"Body"`
	}
	Body struct {
		XMLName 			xml.Name					`xml:"Body"`
		CertGetResponse 	PublicKeyCertificate 		`xml:"AMT_PublicKeyCertificate"`
		ManagementResponse  PublicKeyManagementService	`xml:"AMT_PublicKeyManagementService"`
		EnumerateResponse 	common.EnumerateResponse
		PullResponse  	  	PullResponse 				`xml:"PullResponse"`
	}
	PublicKeyCertificate struct {
		ElementName            string
		InstanceID             string
		X509Certificate        string
		ReadOnlyCertificate    bool
		TrustedRootCertificate bool
		Issuer                 string
		Subject                string
	}
	PublicKeyManagementService struct {
		CreationClassName       string
 		ElementName             string
 		EnabledDefault          int
		EnabledState            int
 		Name                    string
 		RequestedState          int
 		SystemCreationClassName string
 		SystemName              string
	}
	PullResponse struct {
		PublicKeyCertificateItems []PublicKeyCertificate `xml:"Items>AMT_PublicKeyCertificate"`
		PublicKeyManagementServiceItems []PublicKeyManagementService `xml:"Items>AMT_PublicKeyManagementService"`
	}
)

func (w *Response) JSON() string {
	jsonOutput, err := json.Marshal(w.Body)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}