/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publicprivate

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

type KeyPair struct {
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
		GetResponse       PublicPrivateKeyPair
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	PullResponse struct {
		XMLName                   xml.Name               `xml:"PullResponse"`
		PublicPrivateKeyPairItems []PublicPrivateKeyPair `xml:"Items>AMT_PublicPrivateKeyPair"`
	}

	PublicPrivateKeyPair struct {
		XMLName     xml.Name `xml:"AMT_PublicPrivateKeyPair"`
		ElementName string   // A user-friendly name for the object . . .
		InstanceID  string   // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		//DERKey      [1210]uint8 // RSA Key encoded as DES PKCS#1.
		DERKey string
	}

	PublicPrivateSelector message.Selector
)
