/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

type Profile struct {
	base message.Base
}

func NewIEEE8021xProfileWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Profile {
	return Profile{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_IEEE8021xProfile, client),
	}
}

// Get retrieves the representation of the instance
func (profile Profile) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: profile.base.Get(nil),
		},
	}
	// send the message to AMT
	err = profile.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (profile Profile) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: profile.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = profile.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pulls instances of this class, following an Enumerate operation
func (profile Profile) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: profile.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = profile.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Put will change properties of the selected instance
func (profile Profile) Put(ieee8021xProfile ProfileRequest) (response Response, err error) {
	ieee8021xProfile.H = fmt.Sprintf("%s%s", message.AMTSchema, AMT_IEEE8021xProfile)
	response = Response{
		Message: &client.Message{
			XMLInput: profile.base.Put(ieee8021xProfile, false, nil),
		},
	}
	// send the message to AMT
	err = profile.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
