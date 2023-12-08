/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/json"
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
)

const (
	CIM_BootConfigSetting string = "CIM_BootConfigSetting"
	CIM_BootSourceSetting string = "CIM_BootSourceSetting"
	CIM_BootService       string = "CIM_BootService"
	ChangeBootOrder       string = "ChangeBootOrder"
)

var currentMessage string

type Source string

const (
	HardDrive             Source = "CIM:Hard-Disk:1"
	CD                    Source = "CIM:CD/DVD:1"
	PXE                   Source = "CIM:Network:1"
	OCR_UEFI_HTTPS        Source = "Intel(r)AMT:OCR-UEFI-Boot-Option-HTTPS:1"
	OCR_UEFI_BootOption1  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:1"
	OCR_UEFI_BootOption2  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:2"
	OCR_UEFI_BootOption3  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:3"
	OCR_UEFI_BootOption4  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:4"
	OCR_UEFI_BootOption5  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:5"
	OCR_UEFI_BootOption6  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:6"
	OCR_UEFI_BootOption7  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:7"
	OCR_UEFI_BootOption8  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:8"
	OCR_UEFI_BootOption9  Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:9"
	OCR_UEFI_BootOption10 Source = "Intel(r)AMT:OCR-UEFI-Boot-Option:10"
)

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                  xml.Name          `xml:"Body"`
		ConfigSettingGetResponse BootConfigSetting `xml:"CIM_BootConfigSetting"`
		SourceSettingGetResponse BootSourceSetting `xml:"CIM_BootSourceSetting"`
		ServiceGetResponse       BootService       `xml:"CIM_BootService"`
		EnumerateResponse        common.EnumerateResponse
		PullResponse             PullResponse           `xml:"PullResponse"`
		ChangeBootOrder_OUTPUT   ChangeBootOrder_OUTPUT `xml:"ChangeBootOrder_OUTPUT"`
	}

	BootConfigSetting struct {
		XMLName     xml.Name `xml:"CIM_BootConfigSetting"`
		InstanceID  string   `xml:"InstanceID"`
		ElementName string   `xml:"ElementName"`
	}

	BootSourceSetting struct {
		XMLName              xml.Name `xml:"CIM_BootSourceSetting"`
		ElementName          string   `xml:"ElementName"`
		InstanceID           string   `xml:"InstanceID"`
		StructuredBootString string   `xml:"StructuredBootString"`
		BIOSBootString       string   `xml:"BIOSBootString"`
		BootString           string   `xml:"BootString"`
		FailThroughSupported int      `xml:"FailThroughSupported"`
	}

	BootService struct {
		XMLName                 xml.Name `xml:"CIM_BootService"`
		Name                    string   `xml:"Name"`
		CreationClassName       string   `xml:"CreationClassName"`
		SystemName              string   `xml:"SystemName"`
		SystemCreationClassName string   `xml:"SystemCreationClassName"`
		ElementName             string   `xml:"ElementName"`
		OperationalStatus       []int    `xml:"OperationalStatus"`
		EnabledState            int      `xml:"EnabledState"`
		RequestedState          int      `xml:"RequestedState"`
	}

	PullResponse struct {
		BootSourceSettingItems []BootSourceSetting `xml:"Items>CIM_BootSourceSetting"`
		BootConfigSettingItems []BootConfigSetting `xml:"Items>CIM_BootConfigSetting"`
		BootServiceItems       []BootService       `xml:"Items>CIM_BootService"`
	}

	ChangeBootOrder_OUTPUT struct {
		message.ReturnValue
	}
)

func (w *Response) JSON() string {
	jsonOutput, err := json.Marshal(w.Body)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}
