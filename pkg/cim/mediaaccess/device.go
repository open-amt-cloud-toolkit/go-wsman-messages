package mediaaccess

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

const CIM_MediaAccessDevice = "CIM_MediaAccessDevice"

type Device struct {
	base wsman.Base
}

// NewMediaAccessDevice returns a new instance of the MediaAccessDevice struct.
func NewMediaAccessDevice(wsmanMessageCreator *wsman.WSManMessageCreator) Device {
	return Device{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_MediaAccessDevice)),
	}
}
func (b Device) Get() string {
	return b.base.Get(nil)
}

func (b Device) Enumerate() string {
	return b.base.Enumerate()
}
func (b Device) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
