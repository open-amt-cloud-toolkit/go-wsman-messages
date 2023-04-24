package wifi

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/actions"
)

type Port struct {
	base wsman.Base
}

const CIM_WiFiPort = "CIM_WiFiPort"

// NewWiFiPort returns a new instance of the WiFiPort struct.
func NewWiFiPort(wsmanMessageCreator *wsman.WSManMessageCreator) Port {
	return Port{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_WiFiPort)),
	}
}
func (w Port) RequestStateChange(requestedState int) string {
	return w.base.RequestStateChange(actions.RequestStateChange(string(CIM_WiFiPort)), requestedState)
}
func (b Port) Get() string {
	return b.base.Get(nil)
}

func (b Port) Enumerate() string {
	return b.base.Enumerate()
}
func (b Port) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
