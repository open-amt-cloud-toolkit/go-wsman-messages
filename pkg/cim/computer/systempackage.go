package computer

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

const CIM_ComputerSystemPackage = "CIM_ComputerSystemPackage"

type SystemPackage struct {
	base wsman.Base
}

// NewComputerSystemPackage returns a new instance of the ComputerSystemPackage struct.
func NewComputerSystemPackage(wsmanMessageCreator *wsman.WSManMessageCreator) SystemPackage {
	return SystemPackage{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_ComputerSystemPackage)),
	}
}
func (b SystemPackage) Get() string {
	return b.base.Get(nil)
}

func (b SystemPackage) Enumerate() string {
	return b.base.Enumerate()
}
func (b SystemPackage) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}
