package ieee8021x

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

type IEEE8021xProfile struct {
	ElementName                     string
	InstanceID                      string
	Enabled                         bool
	ActiveInS0                      bool
	AuthenticationProtocol          AuthenticationProtocol
	RoamingIdentity                 string
	ServerCertificateName           string
	ServerCertificateNameComparison ServerCertificateNameComparison
	Username                        string
	Password                        string
	Domain                          string
	ProtectedAccessCredential       []int
	PACPassword                     string
	ClientCertificate               string
	ServerCertificateIssue          string
	PxeTimeout                      int
}

type AuthenticationProtocol int

const AMT_IEEE8021xProfile = "AMT_8021XProfile"

const (
	AuthenticationProtocolTLS AuthenticationProtocol = iota
	AuthenticationProtocolTTLS_MSCHAPv2
	AuthenticationProtocolPEAP_MSCHAPv2
	AuthenticationProtocolEAP_GTC
	AuthenticationProtocolEAPFAST_MSCHAPv2
	AuthenticationProtocolEAPFAST_GTC
	AuthenticationProtocolEAPFAST_TLS
)

type ServerCertificateNameComparison int

const (
	ServerCertificateNameComparisonFullName ServerCertificateNameComparison = iota
	ServerCertificateNameComparisonDomainSuffix
)

type Profile struct {
	base wsman.Base
}

func NewIEEE8021xProfile(wsmanMessageCreator *wsman.WSManMessageCreator) Profile {
	return Profile{
		base: wsman.NewBase(wsmanMessageCreator, AMT_IEEE8021xProfile),
	}
}
func (IEEE8021xProfile Profile) Get() string {
	return IEEE8021xProfile.base.Get(nil)
}
func (IEEE8021xProfile Profile) Enumerate() string {
	return IEEE8021xProfile.base.Enumerate()
}
func (IEEE8021xProfile Profile) Pull(enumerationContext string) string {
	return IEEE8021xProfile.base.Pull(enumerationContext)
}
func (IEEE8021xProfile Profile) Put(ieee8021xProfile IEEE8021xProfile) string {
	return IEEE8021xProfile.base.Put(ieee8021xProfile, false, nil)
}
