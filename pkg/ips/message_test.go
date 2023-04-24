package ips

import (
	"fmt"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/stretchr/testify/assert"
)

const (
	xmlHeader               = `<?xml version="1.0" encoding="utf-8"?>`
	envelope                = `<Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>`
	enumerationContext      = `AC070000-0000-0000-0000-000000000000`
	operationTimeout        = `PT60S`
	serverCertificateIssuer = `serverCertificateIssuer`
	clientCertificate       = `clientCertificate`
	adminPassEncryptionType = 2
	adminPassword           = `bebb3497d69b544c732651365cc3462d`
	mcNonce                 = `ZxxE0cFy590zDBIR39q6QU6iuII=`
	signingAlgorithm        = 2
	digitalSignature        = `T0NvoR7RUkOpVULIcNL0VhpEK5rO3j5/TBpN82q1YgPM5sRBxqymu7fKBgAGGN49oD8xsqW4X0SWxjuB3q/TLHjNJJNxoHHlXZnb77HTwfXHp59E/TM10UvOX96qEgKU5Mp+8/IE9LnYxC1ajQostSRA/X+HA5F6kRctLiCK+ViWUCk4sAtPzHhhHSTB/98KDWuacPepScSpref532hpD2/g43nD3Wg0SjmOMExPLMMnijWE9KDkxE00+Bos28DD3Yclj4BMhkoXDw6k4EcTWKbGhtF/9meXXmSPwRmXEaWe8COIDrQks1mpyLblYu8yHHnUjhssdcCQHtAOu7t0RA==`
	GET                     = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	ENUMERATE               = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate"
	PULL                    = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull"
	DELETE                  = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete"
	ENUMERATE_BODY          = "<Enumerate xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" />"
	SET_CERTIFICATES        = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings/SetCertificates"
	ADD_NEXT_CERT_IN_CHAIN  = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/AddNextCertInChain"
	ADMIN_SETUP             = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/AdminSetup"
	SETUP                   = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/Setup"
	SEND_OPT_IN_CODE        = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService/SendOptInCode"
	START_OPT_IN            = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService/StartOptIn"
	CANCEL_OPT_IN           = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService/CancelOptIn"
)

func TestIPS(t *testing.T) {
	messageID := 0
	ipsClass := NewMessages()

	expectedResponse := func(method, action, headerExtra, body string) string {
		return fmt.Sprintf(`%s%s%s</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/%s</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>%s</w:OperationTimeout>%s</Header><Body>%s</Body></Envelope>`, xmlHeader, envelope, action, method, messageID, operationTimeout, headerExtra, body)
	}

	t.Run("ips_OptInService Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			headerExtra  string
			body         string
			responseFunc func() string
		}{
			// GET
			{"should create a valid ips_OptInService Get wsman message", "IPS_OptInService", GET, "", "", ipsClass.OptInService.Get},
			{"should create a valid ips_HostBasedSetupService Get wsman message", "IPS_HostBasedSetupService", GET, "", "", ipsClass.HostBasedSetupService.Get},
			{"should create a valid ips_AlarmClockOccurrence Get wsman message", "IPS_AlarmClockOccurrence", GET, "", "", ipsClass.AlarmClockOccurrence.Get},
			{"should create a valid ips_IEEE8021xCredentialContext Get wsman message", "IPS_8021xCredentialContext", GET, "", "", ipsClass.IEEE8021xCredentialContext.Get},
			{"should create a valid ips_IEEE8021xSettings Get wsman message", "IPS_IEEE8021xSettings", GET, "", "", ipsClass.IEEE8021xSettings.Get},

			// ENUMERATE
			{"should create a valid ips_OptInService Enumerate wsman message", "IPS_OptInService", ENUMERATE, "", ENUMERATE_BODY, ipsClass.OptInService.Enumerate},
			{"should create a valid ips_HostBasedSetupService Enumerate wsman message", "IPS_HostBasedSetupService", ENUMERATE, "", ENUMERATE_BODY, ipsClass.HostBasedSetupService.Enumerate},
			{"should create a valid ips_AlarmClockOccurrence Enumerate wsman message", "IPS_AlarmClockOccurrence", ENUMERATE, "", ENUMERATE_BODY, ipsClass.AlarmClockOccurrence.Enumerate},
			{"should create a valid ips_IEEE8021xCredentialContext Enumerate wsman message", "IPS_8021xCredentialContext", ENUMERATE, "", ENUMERATE_BODY, ipsClass.IEEE8021xCredentialContext.Enumerate},
			{"should create a valid ips_IEEE8021xSettings Enumerate wsman message", "IPS_IEEE8021xSettings", ENUMERATE, "", ENUMERATE_BODY, ipsClass.IEEE8021xSettings.Enumerate},

			// PULL
			{"should create a valid ips_OptInService Pull wsman message", "IPS_OptInService", PULL, "", fmt.Sprintf(`<Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull>`, enumerationContext), func() string { return ipsClass.OptInService.Pull(enumerationContext) }},
			{"should create a valid ips_HostBasedSetupService Pull wsman message", "IPS_HostBasedSetupService", PULL, "", fmt.Sprintf(`<Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull>`, enumerationContext), func() string { return ipsClass.HostBasedSetupService.Pull(enumerationContext) }},
			{"should create a valid ips_AlarmClockOccurrence Pull wsman message", "IPS_AlarmClockOccurrence", PULL, "", fmt.Sprintf(`<Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull>`, enumerationContext), func() string { return ipsClass.AlarmClockOccurrence.Pull(enumerationContext) }},
			{"should create a valid ips_IEEE8021xCredentialContext Pull wsman message", "IPS_8021xCredentialContext", PULL, "", fmt.Sprintf(`<Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull>`, enumerationContext), func() string { return ipsClass.IEEE8021xCredentialContext.Pull(enumerationContext) }},
			{"should create a valid ips_IEEE8021xSettings Pull wsman message", "IPS_IEEE8021xSettings", PULL, "", fmt.Sprintf(`<Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull>`, enumerationContext), func() string { return ipsClass.IEEE8021xSettings.Pull(enumerationContext) }},

			// DELETE
			{"should create a valid ips_AlarmClockOccurrence Delete wsman message", "IPS_AlarmClockOccurrence", DELETE, "<w:SelectorSet><w:Selector Name=\"Name\">Instance</w:Selector></w:SelectorSet>", "", func() string {
				selector := &wsman.Selector{
					Name:  "Name",
					Value: "Instance",
				}
				return ipsClass.AlarmClockOccurrence.Delete(selector)
			}},

			// SET CERTIFICATES
			{"should create a valid ips_IEEE8021xSettings set certificates wsman message", "IPS_IEEE8021xSettings", SET_CERTIFICATES, "", fmt.Sprintf(`<h:SetCertificates_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings"><h:ServerCertificateIssuer>%s</h:ServerCertificateIssuer><h:ClientCertificate>%s</h:ClientCertificate></h:SetCertificates_INPUT>`, serverCertificateIssuer, clientCertificate), func() string {
				return ipsClass.IEEE8021xSettings.SetCertificates(serverCertificateIssuer, clientCertificate)
			}},

			// ADD NEXT CERT IN CHAIN
			{"should create a valid ips_HostBasedSetupService AddNextCertInChain wsman message", "IPS_HostBasedSetupService", ADD_NEXT_CERT_IN_CHAIN, "", fmt.Sprintf(`<h:AddNextCertInChain_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NextCertificate>%s</h:NextCertificate><h:IsLeafCertificate>true</h:IsLeafCertificate><h:IsRootCertificate>false</h:IsRootCertificate></h:AddNextCertInChain_INPUT>`, clientCertificate), func() string {
				return ipsClass.HostBasedSetupService.AddNextCertInChain(clientCertificate, true, false)
			}},

			// AdminSetup
			{"should create a valid ips_HostBasedSetupService AdminSetup wsman message", "IPS_HostBasedSetupService", ADMIN_SETUP, "", fmt.Sprintf(`<h:AdminSetup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword><h:McNonce>%s</h:McNonce><h:SigningAlgorithm>%d</h:SigningAlgorithm><h:DigitalSignature>%s</h:DigitalSignature></h:AdminSetup_INPUT>`, adminPassEncryptionType, adminPassword, mcNonce, signingAlgorithm, digitalSignature), func() string {
				return ipsClass.HostBasedSetupService.AdminSetup(adminPassEncryptionType, adminPassword, mcNonce, signingAlgorithm, digitalSignature)
			}},

			//ADD NEXT CERT IN CHAIN
			{"should create a valid ips_HostBasedSetupService Setup wsman message", "IPS_HostBasedSetupService", SETUP, "", fmt.Sprintf(`<h:Setup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword></h:Setup_INPUT>`, adminPassEncryptionType, adminPassword), func() string {
				return ipsClass.HostBasedSetupService.Setup(adminPassEncryptionType, adminPassword)
			}},

			// SEND_OPT_IN_CODE
			{"should create a valid ips_OptInService send opt in code wsman message", "IPS_OptInService", SEND_OPT_IN_CODE, "", `<h:SendOptInCode_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"><h:OptInCode>1</h:OptInCode></h:SendOptInCode_INPUT>`, func() string {
				return ipsClass.OptInService.SendOptInCode(1)
			}},

			// START_OPT_IN
			{"should create a valid ips_OptInService start opt in code wsman message", "IPS_OptInService", START_OPT_IN, "", `<h:StartOptIn_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"></h:StartOptIn_INPUT>`, func() string {
				return ipsClass.OptInService.StartOptIn()
			}},

			// CANCEL_OPT_IN
			{"should create a valid ips_OptInService cancel opt in code wsman message", "IPS_OptInService", CANCEL_OPT_IN, "", `<h:CancelOptIn_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"></h:CancelOptIn_INPUT>`, func() string {
				return ipsClass.OptInService.CancelOptIn()
			}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				correctResponse := expectedResponse(test.method, test.action, test.headerExtra, test.body)
				messageID++
				response := test.responseFunc()
				if response != correctResponse {
					assert.Equal(t, correctResponse, response)
				}
			})
		}
	})
}
