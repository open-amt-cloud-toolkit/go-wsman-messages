package physical

import "github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"

type Card struct {
	base wsman.Base
}

const CIM_Card = "CIM_Card"

// NewCard returns a new instance of the Card struct.
func NewCard(wsmanMessageCreator *wsman.WSManMessageCreator) Card {
	return Card{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_Card)),
	}
}
func (b Card) Get() string {
	return b.base.Get(nil)
}

func (b Card) Enumerate() string {
	return b.base.Enumerate()
}
func (b Card) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}