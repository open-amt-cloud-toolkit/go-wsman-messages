package publickey

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_PublicKeyCertificate = "AMT_PublicKeyCertificate"

type PublicKeyCertificate struct {
	ElementName           string      // A user-friendly name for the object . . .
	InstanceID            string      // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
	X509Certificate       [4100]uint8 // uint8[4100] // The X.509 Certificate blob.
	TrustedRootCertficate bool        // For root certificate [that were added by AMT_PublicKeyManagementService.AddTrustedRootCertificate()]this property will be true.
	Issuer                string      // The Issuer field of this certificate.
	Subject               string      // The Subject field of this certificate.
	ReadOnlyCertificate   bool        // Indicates whether the certificate is an Intel AMT self-signed certificate. If True, the certificate cannot be deleted.
}
type Certificate struct {
	base wsman.Base
}

func NewPublicKeyCertificate(wsmanMessageCreator *wsman.WSManMessageCreator) Certificate {
	return Certificate{
		base: wsman.NewBase(wsmanMessageCreator, AMT_PublicKeyCertificate),
	}
}
func (PublicKeyCertificate Certificate) Get() string {
	return PublicKeyCertificate.base.Get(nil)
}
func (PublicKeyCertificate Certificate) Enumerate() string {
	return PublicKeyCertificate.base.Enumerate()
}
func (PublicKeyCertificate Certificate) Pull(enumerationContext string) string {
	return PublicKeyCertificate.base.Pull(enumerationContext)
}
func (PublicKeyCertificate Certificate) Put(publicKeyCertificate PublicKeyCertificate) string {
	return PublicKeyCertificate.base.Put(publicKeyCertificate, false, nil)
}
func (PublicKeyCertificate Certificate) Delete(selector *wsman.Selector) string {
	return PublicKeyCertificate.base.Delete(selector)
}
