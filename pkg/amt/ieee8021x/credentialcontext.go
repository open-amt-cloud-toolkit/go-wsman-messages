package ieee8021x

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_IEEE8021xCredentialContext = "AMT_8021xCredentialContext"

type CredentialContext struct {
	base wsman.Base
}

func NewIEEE8021xCredentialContext(wsmanMessageCreator *wsman.WSManMessageCreator) CredentialContext {
	return CredentialContext{
		base: wsman.NewBase(wsmanMessageCreator, AMT_IEEE8021xCredentialContext),
	}
}
func (IEEE8021xCredentialContext CredentialContext) Get() string {
	return IEEE8021xCredentialContext.base.Get(nil)
}
func (IEEE8021xCredentialContext CredentialContext) Enumerate() string {
	return IEEE8021xCredentialContext.base.Enumerate()
}
func (IEEE8021xCredentialContext CredentialContext) Pull(enumerationContext string) string {
	return IEEE8021xCredentialContext.base.Pull(enumerationContext)
}
