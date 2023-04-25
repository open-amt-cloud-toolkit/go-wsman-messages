package actions

type Actions string

const (
	Enumerate          Actions = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate"
	Pull               Actions = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull"
	Get                Actions = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	Put                Actions = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Put"
	Delete             Actions = "http://schemas.xmlsoap.org/ws/2004/09/transfeeleteete"
	Setup              Actions = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/Setup"
	AdminSetup         Actions = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/AdminSetup"
	AddNextCertInChain Actions = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/AddNextCertInChain"
	StartOptIn         Actions = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService/StartOptIn"
	CancelOptIn        Actions = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService/CancelOptIn"
	SendOptInCode      Actions = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService/SendOptInCode"
	SetCertificates    Actions = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings/SetCertificates"
)
