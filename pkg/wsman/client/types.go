package client

// Parameters struct defines the connection settings for wsman client
type Parameters struct {
	Target            string
	Username          string
	Password          string
	UseDigest         bool
	UseTLS            bool
	SelfSignedAllowed bool
	LogAMTMessages    bool
	IsRedirection     bool
}
