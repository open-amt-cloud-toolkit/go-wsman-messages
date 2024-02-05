/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

const (
	EnvelopeResponseApply = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBodyApply          = `<g:AMT_AuthorizationService><g:CreationClassName>AMT_AuthorizationService</g:CreationClassName><g:ElementName>Intel(r) AMT Authorization Service</g:ElementName><g:Name>Intel(r) AMT Alarm Clock Service</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_AuthorizationService>`
)

func TestPositiveAMT_RemoteAccessPolicyAppliesToMPS(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/remoteaccess/policyappliestomps",
	}
	elementUnderTest := NewRemoteAccessPolicyAppliesToMPSWithClient(wsmanMessageCreator, &client)
	t.Run("amt_RemoteAccessPolicyAppliesToMPS Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Get wsman message",
				AMT_RemoteAccessPolicyAppliesToMPS,
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RemoteAccessPolicyAppliesToMPSGetResponse: RemoteAccessPolicyAppliesToMPSResponse{
						XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "AMT_RemoteAccessPolicyAppliesToMPS"},
						ManagedElement: ManagedElementResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "ManagedElement"},
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParametersResponse{
								XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP",
								SelectorSet: SelectorSetResponse{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
									Selectors: []SelectorResponse{
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "CreationClassName",
											Text:    "AMT_ManagementPresenceRemoteSAP",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "Name",
											Text:    "Intel(r) AMT:Management Presence Server 0",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemCreationClassName",
											Text:    "CIM_ComputerSystem",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemName",
											Text:    "Intel(r) AMT",
										},
									},
								},
							},
						},
						MpsType:       2,
						OrderOfAccess: 0,
						PolicySet: PolicySetResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "PolicySet"},
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParametersResponse{
								XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule",
								SelectorSet: SelectorSetResponse{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
									Selectors: []SelectorResponse{
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "CreationClassName",
											Text:    "AMT_RemoteAccessPolicyRule",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "PolicyRuleName",
											Text:    "Periodic",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemCreationClassName",
											Text:    "CIM_ComputerSystem",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemName",
											Text:    "Intel(r) AMT",
										},
									},
								},
							},
						},
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Enumerate wsman message",
				AMT_RemoteAccessPolicyAppliesToMPS,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CE000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Pull wsman message",
				AMT_RemoteAccessPolicyAppliesToMPS,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						PolicyAppliesItems: []RemoteAccessPolicyAppliesToMPSResponse{
							{
								XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "AMT_RemoteAccessPolicyAppliesToMPS"},
								ManagedElement: ManagedElementResponse{
									XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "ManagedElement"},
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: ReferenceParametersResponse{
										XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
										ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP",
										SelectorSet: SelectorSetResponse{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selectors: []SelectorResponse{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "CreationClassName",
													Text:    "AMT_ManagementPresenceRemoteSAP",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "Name",
													Text:    "Intel(r) AMT:Management Presence Server 0",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "SystemCreationClassName",
													Text:    "CIM_ComputerSystem",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "SystemName",
													Text:    "Intel(r) AMT",
												},
											},
										},
									},
								},
								MpsType:       0,
								OrderOfAccess: 0,
								PolicySet: PolicySetResponse{
									XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "PolicySet"},
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: ReferenceParametersResponse{
										XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
										ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule",
										SelectorSet: SelectorSetResponse{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selectors: []SelectorResponse{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "CreationClassName",
													Text:    "AMT_RemoteAccessPolicyRule",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "PolicyRuleName",
													Text:    "Periodic",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "SystemCreationClassName",
													Text:    "CIM_ComputerSystem",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "SystemName",
													Text:    "Intel(r) AMT",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Put wsman message",
				AMT_RemoteAccessPolicyAppliesToMPS,
				wsmantesting.PUT,
				`<h:AMT_RemoteAccessPolicyAppliesToMPS xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS"><h:ManagedElement xmlns:b="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS"><b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address><b:ReferenceParameters xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><c:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</c:ResourceURI><c:SelectorSet><c:Selector Name="CreationClassName">AMT_ManagementPresenceRemoteSAP</c:Selector><c:Selector Name="Name">Intel(r) AMT:Management Presence Server 0</c:Selector><c:Selector Name="SystemCreationClassName">CIM_ComputerSystem</c:Selector><c:Selector Name="SystemName">Intel(r) AMT</c:Selector></c:SelectorSet></b:ReferenceParameters></h:ManagedElement><h:OrderOfAccess>0</h:OrderOfAccess><h:MpsType>2</h:MpsType><h:PolicySet xmlns:b="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS"><b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address><b:ReferenceParameters xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><c:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</c:ResourceURI><c:SelectorSet><c:Selector Name="CreationClassName">AMT_RemoteAccessPolicyRule</c:Selector><c:Selector Name="PolicyRuleName">Periodic</c:Selector><c:Selector Name="SystemCreationClassName">CIM_ComputerSystem</c:Selector><c:Selector Name="SystemName">Intel(r) AMT</c:Selector></c:SelectorSet></b:ReferenceParameters></h:PolicySet></h:AMT_RemoteAccessPolicyAppliesToMPS>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Put"
					rapatmps := RemoteAccessPolicyAppliesToMPSRequest{
						H: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS",
						ManagedElement: ManagedElement{
							B:       "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS",
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParameters{
								C:           "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP",
								SelectorSet: SelectorSet{
									Selectors: []Selector{
										{
											Name: "CreationClassName",
											Text: "AMT_ManagementPresenceRemoteSAP",
										},
										{
											Name: "Name",
											Text: "Intel(r) AMT:Management Presence Server 0",
										},
										{
											Name: "SystemCreationClassName",
											Text: "CIM_ComputerSystem",
										},
										{
											Name: "SystemName",
											Text: "Intel(r) AMT",
										},
									},
								},
							},
						},
						OrderOfAccess: 0,
						MPSType:       BothMPS,
						PolicySet: PolicySet{
							B:       "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS",
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParameters{
								C:           "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule",
								SelectorSet: SelectorSet{
									Selectors: []Selector{
										{
											Name: "CreationClassName",
											Text: "AMT_RemoteAccessPolicyRule",
										},
										{
											Name: "PolicyRuleName",
											Text: "Periodic",
										},
										{
											Name: "SystemCreationClassName",
											Text: "CIM_ComputerSystem",
										},
										{
											Name: "SystemName",
											Text: "Intel(r) AMT",
										},
									},
								},
							},
						},
					}
					return elementUnderTest.Put(&rapatmps)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
				},
			},
			//{"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Create wsman message", AMT_RemoteAccessPolicyAppliesToMPS, wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Delete wsman message",
				AMT_RemoteAccessPolicyAppliesToMPS, wsmantesting.DELETE,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">Instance</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					return elementUnderTest.Delete("Instance")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
func TestNegativeAMT_RemoteAccessPolicyAppliesToMPS(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/remoteaccess/policyappliestomps",
	}
	elementUnderTest := NewRemoteAccessPolicyAppliesToMPSWithClient(wsmanMessageCreator, &client)
	t.Run("amt_RemoteAccessPolicyAppliesToMPS Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Get wsman message",
				AMT_RemoteAccessPolicyAppliesToMPS,
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RemoteAccessPolicyAppliesToMPSGetResponse: RemoteAccessPolicyAppliesToMPSResponse{
						XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "AMT_RemoteAccessPolicyAppliesToMPS"},
						ManagedElement: ManagedElementResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "ManagedElement"},
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParametersResponse{
								XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP",
								SelectorSet: SelectorSetResponse{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
									Selectors: []SelectorResponse{
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "CreationClassName",
											Text:    "AMT_ManagementPresenceRemoteSAP",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "Name",
											Text:    "Intel(r) AMT:Management Presence Server 0",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemCreationClassName",
											Text:    "CIM_ComputerSystem",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemName",
											Text:    "Intel(r) AMT",
										},
									},
								},
							},
						},
						MpsType:       2,
						OrderOfAccess: 0,
						PolicySet: PolicySetResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "PolicySet"},
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParametersResponse{
								XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule",
								SelectorSet: SelectorSetResponse{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
									Selectors: []SelectorResponse{
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "CreationClassName",
											Text:    "AMT_RemoteAccessPolicyRule",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "PolicyRuleName",
											Text:    "Periodic",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemCreationClassName",
											Text:    "CIM_ComputerSystem",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemName",
											Text:    "Intel(r) AMT",
										},
									},
								},
							},
						},
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Enumerate wsman message",
				AMT_RemoteAccessPolicyAppliesToMPS,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CE000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Pull wsman message",
				AMT_RemoteAccessPolicyAppliesToMPS,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						PolicyAppliesItems: []RemoteAccessPolicyAppliesToMPSResponse{
							{
								XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "AMT_RemoteAccessPolicyAppliesToMPS"},
								ManagedElement: ManagedElementResponse{
									XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "ManagedElement"},
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: ReferenceParametersResponse{
										XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
										ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP",
										SelectorSet: SelectorSetResponse{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selectors: []SelectorResponse{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "CreationClassName",
													Text:    "AMT_ManagementPresenceRemoteSAP",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "Name",
													Text:    "Intel(r) AMT:Management Presence Server 0",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "SystemCreationClassName",
													Text:    "CIM_ComputerSystem",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "SystemName",
													Text:    "Intel(r) AMT",
												},
											},
										},
									},
								},
								MpsType:       0,
								OrderOfAccess: 0,
								PolicySet: PolicySetResponse{
									XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS", Local: "PolicySet"},
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: ReferenceParametersResponse{
										XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
										ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule",
										SelectorSet: SelectorSetResponse{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selectors: []SelectorResponse{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "CreationClassName",
													Text:    "AMT_RemoteAccessPolicyRule",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "PolicyRuleName",
													Text:    "Periodic",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "SystemCreationClassName",
													Text:    "CIM_ComputerSystem",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "SystemName",
													Text:    "Intel(r) AMT",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Put wsman message",
				AMT_RemoteAccessPolicyAppliesToMPS,
				wsmantesting.PUT,
				`<h:AMT_RemoteAccessPolicyAppliesToMPS xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS"><h:ManagedElement xmlns:b="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS"><b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address><b:ReferenceParameters xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><c:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</c:ResourceURI><c:SelectorSet><c:Selector Name="CreationClassName">AMT_ManagementPresenceRemoteSAP</c:Selector><c:Selector Name="Name">Intel(r) AMT:Management Presence Server 0</c:Selector><c:Selector Name="SystemCreationClassName">CIM_ComputerSystem</c:Selector><c:Selector Name="SystemName">Intel(r) AMT</c:Selector></c:SelectorSet></b:ReferenceParameters></h:ManagedElement><h:OrderOfAccess>0</h:OrderOfAccess><h:MpsType>2</h:MpsType><h:PolicySet xmlns:b="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS"><b:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:Address><b:ReferenceParameters xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><c:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</c:ResourceURI><c:SelectorSet><c:Selector Name="CreationClassName">AMT_RemoteAccessPolicyRule</c:Selector><c:Selector Name="PolicyRuleName">Periodic</c:Selector><c:Selector Name="SystemCreationClassName">CIM_ComputerSystem</c:Selector><c:Selector Name="SystemName">Intel(r) AMT</c:Selector></c:SelectorSet></b:ReferenceParameters></h:PolicySet></h:AMT_RemoteAccessPolicyAppliesToMPS>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					rapatmps := RemoteAccessPolicyAppliesToMPSRequest{
						H: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS",
						ManagedElement: ManagedElement{
							B:       "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS",
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParameters{
								C:           "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP",
								SelectorSet: SelectorSet{
									Selectors: []Selector{
										{
											Name: "CreationClassName",
											Text: "AMT_ManagementPresenceRemoteSAP",
										},
										{
											Name: "Name",
											Text: "Intel(r) AMT:Management Presence Server 0",
										},
										{
											Name: "SystemCreationClassName",
											Text: "CIM_ComputerSystem",
										},
										{
											Name: "SystemName",
											Text: "Intel(r) AMT",
										},
									},
								},
							},
						},
						OrderOfAccess: 0,
						MPSType:       BothMPS,
						PolicySet: PolicySet{
							B:       "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS",
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParameters{
								C:           "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule",
								SelectorSet: SelectorSet{
									Selectors: []Selector{
										{
											Name: "CreationClassName",
											Text: "AMT_RemoteAccessPolicyRule",
										},
										{
											Name: "PolicyRuleName",
											Text: "Periodic",
										},
										{
											Name: "SystemCreationClassName",
											Text: "CIM_ComputerSystem",
										},
										{
											Name: "SystemName",
											Text: "Intel(r) AMT",
										},
									},
								},
							},
						},
					}
					return elementUnderTest.Put(&rapatmps)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
				},
			},
			//{"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Create wsman message", AMT_RemoteAccessPolicyAppliesToMPS, wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			{
				"should create a valid AMT_RemoteAccessPolicyAppliesToMPS Delete wsman message",
				AMT_RemoteAccessPolicyAppliesToMPS, wsmantesting.DELETE,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">Instance</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Delete("Instance")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
