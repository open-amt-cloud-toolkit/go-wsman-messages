package physical

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type Memory struct {
	base wsman.Base
}

const CIM_PhysicalMemory = "CIM_PhysicalMemory"

// NewPhysicalMemory returns a new instance of the PhysicalMemory struct.
func NewPhysicalMemory(wsmanMessageCreator *wsman.WSManMessageCreator) Memory {
	return Memory{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_PhysicalMemory)),
	}
}
func (b Memory) Get() string {
	return b.base.Get(nil)
}

func (b Memory) Enumerate() string {
	return b.base.Enumerate()
}
func (b Memory) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
