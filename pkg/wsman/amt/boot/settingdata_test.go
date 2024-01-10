/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/wsmantesting"
)

func TestPositiveAMT_BootSettingData(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/boot/settingdata",
	}
	elementUnderTest := NewBootSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_BootSettingData Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_BootSettingData Get wsman message",
				AMT_BootSettingData,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					BootSettingDataGetResponse: BootSettingDataResponse{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootSettingData", Local: "AMT_BootSettingData"},
						InstanceID:             "Intel(r) AMT:BootSettingData 0",
						ElementName:            "Intel(r) AMT Boot Configuration Settings",
						OwningEntity:           "Intel(r) AMT",
						UseSOL:                 false,
						UseSafeMode:            false,
						ReflashBIOS:            false,
						BIOSSetup:              false,
						BIOSPause:              false,
						LockPowerButton:        false,
						LockResetButton:        false,
						LockKeyboard:           false,
						LockSleepButton:        false,
						UserPasswordBypass:     false,
						ForcedProgressEvents:   false,
						FirmwareVerbosity:      0,
						ConfigurationDataReset: false,
						IDERBootDevice:         0,
						UseIDER:                false,
						EnforceSecureBoot:      false,
						BootMediaIndex:         0,
						SecureErase:            false,
						RSEPassword:            "",
						OptionsCleared:         false,
						BIOSLastStatus:         []int{2, 0},
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_BootSettingData Enumerate wsman message",
				AMT_BootSettingData,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "DD070000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_BootSettingData Pull wsman message",
				AMT_BootSettingData,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						BootSettingDataItems: []BootSettingDataResponse{
							{
								XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootSettingData", Local: "AMT_BootSettingData"},
								InstanceID:             "Intel(r) AMT:BootSettingData 0",
								ElementName:            "Intel(r) AMT Boot Configuration Settings",
								OwningEntity:           "Intel(r) AMT",
								UseSOL:                 false,
								UseSafeMode:            false,
								ReflashBIOS:            false,
								BIOSSetup:              false,
								BIOSPause:              false,
								LockPowerButton:        false,
								LockResetButton:        false,
								LockKeyboard:           false,
								LockSleepButton:        false,
								UserPasswordBypass:     false,
								ForcedProgressEvents:   false,
								FirmwareVerbosity:      0,
								ConfigurationDataReset: false,
								IDERBootDevice:         0,
								UseIDER:                false,
								EnforceSecureBoot:      false,
								BootMediaIndex:         0,
								SecureErase:            false,
								RSEPassword:            "",
								OptionsCleared:         false,
								BIOSLastStatus:         []int{2, 0},
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}

func TestNegativeAMT_BootSettingData(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/boot/settingdata",
	}
	elementUnderTest := NewBootSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_BootSettingData Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_BootSettingData Get wsman message",
				AMT_BootSettingData,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					BootSettingDataGetResponse: BootSettingDataResponse{
						XMLName:                  xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootSettingData", Local: "AMT_BootSettingData"},
						InstanceID:               "t",
						ElementName:              "t",
						OwningEntity:             "t",
						UseSOL:                   false,
						UseSafeMode:              false,
						ReflashBIOS:              false,
						BIOSSetup:                false,
						BIOSPause:                false,
						LockPowerButton:          false,
						LockResetButton:          false,
						LockKeyboard:             false,
						LockSleepButton:          false,
						UserPasswordBypass:       false,
						ForcedProgressEvents:     false,
						FirmwareVerbosity:        0,
						ConfigurationDataReset:   false,
						IDERBootDevice:           1,
						UseIDER:                  false,
						EnforceSecureBoot:        false,
						BootMediaIndex:           0,
						SecureErase:              false,
						RSEPassword:              "",
						OptionsCleared:           false,
						WinREBootEnabled:         false,
						UEFILocalPBABootEnabled:  false,
						UEFIHTTPSBootEnabled:     false,
						SecureBootControlEnabled: false,
						BootguardStatus:          false,
						BIOSLastStatus:           []int{0, 0},
						UEFIBootParametersArray:  []int{0},
						UEFIBootNumberOfParams:   []int{0},
						RPEEnabled:               false,
						PlatformErase:            false,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_BootSettingData Enumerate wsman message",
				AMT_BootSettingData,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "5C000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_BootSettingData Pull wsman message",
				AMT_BootSettingData,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						BootSettingDataItems: []BootSettingDataResponse{
							{
								XMLName:                  xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootSettingData", Local: "AMT_BootSettingData"},
								InstanceID:               "t",
								ElementName:              "t",
								OwningEntity:             "t",
								UseSOL:                   false,
								UseSafeMode:              false,
								ReflashBIOS:              false,
								BIOSSetup:                false,
								BIOSPause:                false,
								LockPowerButton:          false,
								LockResetButton:          false,
								LockKeyboard:             false,
								LockSleepButton:          false,
								UserPasswordBypass:       false,
								ForcedProgressEvents:     false,
								FirmwareVerbosity:        0,
								ConfigurationDataReset:   false,
								IDERBootDevice:           1,
								UseIDER:                  false,
								EnforceSecureBoot:        false,
								BootMediaIndex:           0,
								SecureErase:              false,
								RSEPassword:              "",
								OptionsCleared:           false,
								WinREBootEnabled:         false,
								UEFILocalPBABootEnabled:  false,
								UEFIHTTPSBootEnabled:     false,
								SecureBootControlEnabled: false,
								BootguardStatus:          false,
								BIOSLastStatus:           []int{0, 0},
								UEFIBootParametersArray:  []int{0},
								UEFIBootNumberOfParams:   []int{0},
								RPEEnabled:               false,
								PlatformErase:            false},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
