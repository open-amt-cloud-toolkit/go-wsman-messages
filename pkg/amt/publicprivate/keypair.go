package publicprivate

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_PublicPrivateKeyPair = "AMT_PublicPrivateKeyPair"

type KeyPair struct {
	base wsman.Base
}

func NewPublicPrivateKeyPair(wsmanMessageCreator *wsman.WSManMessageCreator) KeyPair {
	return KeyPair{
		base: wsman.NewBase(wsmanMessageCreator, AMT_PublicPrivateKeyPair),
	}
}
func (PublicPrivateKeyPair KeyPair) Get() string {
	return PublicPrivateKeyPair.base.Get(nil)
}
func (PublicPrivateKeyPair KeyPair) Enumerate() string {
	return PublicPrivateKeyPair.base.Enumerate()
}
func (PublicPrivateKeyPair KeyPair) Pull(enumerationContext string) string {
	return PublicPrivateKeyPair.base.Pull(enumerationContext)
}
