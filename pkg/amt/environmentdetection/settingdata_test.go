package environmentdetection

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_EnvironmentDetectionSettingData(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := wsman.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewEnvironmentDetectionSettingData(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_EnvironmentDetectionSettingData Get wsman message", "AMT_EnvironmentDetectionSettingData", wsmantesting.GET, "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_EnvironmentDetectionSettingData Enumerate wsman message", "AMT_EnvironmentDetectionSettingData", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_EnvironmentDetectionSettingData Pull wsman message", "AMT_EnvironmentDetectionSettingData", wsmantesting.PULL, wsmantesting.PULL_BODY, func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
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
