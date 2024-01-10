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

func TestPositiveAMT_BootCapabilities(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/boot/capabilities",
	}
	elementUnderTest := NewBootCapabilitiesWithClient(wsmanMessageCreator, &client)

	t.Run("amt_BootCapabilities Tests", func(t *testing.T) {
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
				"should create a valid AMT_BootCapabilities Get wsman message",
				AMT_BootCapabilities,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					BootCapabilitiesGetResponse: BootCapabilitiesResponse{
						XMLName:                    xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootCapabilities", Local: "AMT_BootCapabilities"},
						BIOSPause:                  false,
						BIOSReflash:                true,
						BIOSSecureBoot:             true,
						BIOSSetup:                  true,
						ConfigurationDataReset:     false,
						ElementName:                "Intel(r) AMT: Boot Capabilities",
						ForceCDorDVDBoot:           true,
						ForceDiagnosticBoot:        false,
						ForceHardDriveBoot:         true,
						ForceHardDriveSafeModeBoot: false,
						ForcePXEBoot:               true,
						ForcedProgressEvents:       true,
						IDER:                       true,
						InstanceID:                 "Intel(r) AMT:BootCapabilities 0",
						KeyboardLock:               true,
						PowerButtonLock:            false,
						ResetButtonLock:            false,
						SOL:                        true,
						SecureErase:                false,
						SleepButtonLock:            false,
						UserPasswordBypass:         true,
						VerbosityQuiet:             false,
						VerbosityScreenBlank:       false,
						VerbosityVerbose:           false,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_BootCapabilities Enumerate wsman message",
				AMT_BootCapabilities,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "E6070000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_BootCapabilities Pull wsman message",
				AMT_BootCapabilities,
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
						BootCapabilitiesItems: []BootCapabilitiesResponse{
							{
								XMLName:                    xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootCapabilities", Local: "AMT_BootCapabilities"},
								BIOSPause:                  false,
								BIOSReflash:                true,
								BIOSSecureBoot:             true,
								BIOSSetup:                  true,
								ConfigurationDataReset:     false,
								ElementName:                "Intel(r) AMT: Boot Capabilities",
								ForceCDorDVDBoot:           true,
								ForceDiagnosticBoot:        false,
								ForceHardDriveBoot:         true,
								ForceHardDriveSafeModeBoot: false,
								ForcePXEBoot:               true,
								ForcedProgressEvents:       true,
								IDER:                       true,
								InstanceID:                 "Intel(r) AMT:BootCapabilities 0",
								KeyboardLock:               true,
								PowerButtonLock:            false,
								ResetButtonLock:            false,
								SOL:                        true,
								SecureErase:                false,
								SleepButtonLock:            false,
								UserPasswordBypass:         true,
								VerbosityQuiet:             false,
								VerbosityScreenBlank:       false,
								VerbosityVerbose:           false,
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

func TestNegativeAMT_BootCapabilities(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/boot/capabilities",
	}
	elementUnderTest := NewBootCapabilitiesWithClient(wsmanMessageCreator, &client)

	t.Run("amt_BootCapabilities Tests", func(t *testing.T) {
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
				"should create a valid AMT_BootCapabilities Get wsman message",
				AMT_BootCapabilities,
				wsmantesting.GET,
				"<error></error>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					BootCapabilitiesGetResponse: BootCapabilitiesResponse{
						XMLName:                            xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootCapabilities", Local: "AMT_BootCapabilities"},
						InstanceID:                         "t",
						ElementName:                        "t",
						IDER:                               false,
						SOL:                                false,
						BIOSReflash:                        false,
						BIOSSetup:                          false,
						BIOSPause:                          false,
						ForcePXEBoot:                       false,
						ForceHardDriveBoot:                 false,
						ForceHardDriveSafeModeBoot:         false,
						ForceDiagnosticBoot:                false,
						ForceCDorDVDBoot:                   false,
						VerbosityScreenBlank:               false,
						PowerButtonLock:                    false,
						ResetButtonLock:                    false,
						KeyboardLock:                       false,
						SleepButtonLock:                    false,
						UserPasswordBypass:                 false,
						ForcedProgressEvents:               false,
						VerbosityVerbose:                   false,
						VerbosityQuiet:                     false,
						ConfigurationDataReset:             false,
						BIOSSecureBoot:                     false,
						SecureErase:                        false,
						ForceWinREBoot:                     false,
						ForceUEFILocalPBABoot:              false,
						ForceUEFIHTTPSBoot:                 false,
						AMTSecureBootControl:               false,
						UEFIWiFiCoExistenceAndProfileShare: false,
						PlatformErase:                      0,
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_BootCapabilities Enumerate wsman message",
				AMT_BootCapabilities,
				wsmantesting.ENUMERATE,
				"<error></error>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "error",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_BootCapabilities Pull wsman message",
				AMT_BootCapabilities,
				wsmantesting.PULL,
				"<error></error>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						BootCapabilitiesItems: []BootCapabilitiesResponse{
							{
								XMLName:                            xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootCapabilities", Local: "AMT_BootCapabilities"},
								InstanceID:                         "t",
								ElementName:                        "t",
								IDER:                               false,
								SOL:                                false,
								BIOSReflash:                        false,
								BIOSSetup:                          false,
								BIOSPause:                          false,
								ForcePXEBoot:                       false,
								ForceHardDriveBoot:                 false,
								ForceHardDriveSafeModeBoot:         false,
								ForceDiagnosticBoot:                false,
								ForceCDorDVDBoot:                   false,
								VerbosityScreenBlank:               false,
								PowerButtonLock:                    false,
								ResetButtonLock:                    false,
								KeyboardLock:                       false,
								SleepButtonLock:                    false,
								UserPasswordBypass:                 false,
								ForcedProgressEvents:               false,
								VerbosityVerbose:                   false,
								VerbosityQuiet:                     false,
								ConfigurationDataReset:             false,
								BIOSSecureBoot:                     false,
								SecureErase:                        false,
								ForceWinREBoot:                     false,
								ForceUEFILocalPBABoot:              false,
								ForceUEFIHTTPSBoot:                 false,
								AMTSecureBootControl:               false,
								UEFIWiFiCoExistenceAndProfileShare: false,
								PlatformErase:                      0,
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
				assert.Error(t, err)
				assert.NotEqual(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
