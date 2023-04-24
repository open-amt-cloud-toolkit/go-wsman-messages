package methods

type Methods string

const (
	MethodGet                     Methods = "Get"
	MethodPull                    Methods = "Pull"
	MethodEnumerate               Methods = "Enumerate"
	MethodPut                     Methods = "Put"
	MethodDelete                  Methods = "Delete"
	MethodSetBootConfigRole       Methods = "SetBootConfigRole"
	MethodChangeBootOrder         Methods = "ChangeBootOrder"
	MethodRequestPowerStateChange Methods = "RequestPowerStateChange"
	MethodRequestStateChange      Methods = "RequestStateChange"
)
