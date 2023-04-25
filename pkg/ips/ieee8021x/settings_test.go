package ieee8021x

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestIPS_IEEE8021xSettings(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/ips-schema/1/"
	wsmanMessageCreator := wsman.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewIEEE8021xSettings(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid IPS_IEEE8021xSettings Get wsman message", "IPS_IEEE8021xSettings", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid IPS_IEEE8021xSettings Enumerate wsman message", "IPS_IEEE8021xSettings", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid IPS_IEEE8021xSettings Pull wsman message", "IPS_IEEE8021xSettings", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			// SET CERTIFICATES
			{"should create a valid ips_IEEE8021xSettings set certificates wsman message", "IPS_IEEE8021xSettings", wsmantesting.SET_CERTIFICATES, fmt.Sprintf(`<h:SetCertificates_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings"><h:ServerCertificateIssuer>%s</h:ServerCertificateIssuer><h:ClientCertificate>%s</h:ClientCertificate></h:SetCertificates_INPUT>`, wsmantesting.ServerCertificateIssuer, wsmantesting.ClientCertificate), func() string {
				return elementUnderTest.SetCertificates(wsmantesting.ServerCertificateIssuer, wsmantesting.ClientCertificate)
			}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				correctResponse := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response := test.responseFunc()
				if response != correctResponse {
					assert.Equal(t, correctResponse, response)
				}
			})
		}
	})
}
