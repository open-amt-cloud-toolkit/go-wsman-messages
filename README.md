# go-wsman-messages

![CodeQL](https://github.com/open-amt-cloud-toolkit/go-wsman-messages/actions/workflows/codeql-analysis.yml/badge.svg?branch=main&event=push) 
![codecov.io](https://codecov.io/github/open-amt-cloud-toolkit/go-wsman-messages/coverage.svg?branch=main) 
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/open-amt-cloud-toolkit/go-wsman-messages/badge)](https://api.securityscorecards.dev/projects/github.com/open-amt-cloud-toolkit/go-wsman-messages) 
[![Discord Shield](https://discordapp.com/api/guilds/1063200098680582154/widget.png?style=shield)](https://discord.gg/yrcMp2kDWh)

> Disclaimer: Production viable releases are tagged and listed under 'Releases'.  All other check-ins should be considered 'in-development' and should not be used in production

This repository contains a Go library that implements APIs for communicating with Intel® Active Management Technology (AMT) capable platforms. These APIs are based on the AMT SDK documentation, which can be found [here](https://software.intel.com/content/www/us/en/develop/articles/intel-active-management-technology-software-development-kit-sdk.html).

## How to use it

To use this library, you need to import it in your Go project:
``` go
import "github.com/open-amt-cloud-toolkit/go-wsman-messages"
```

Then, you can create an instance of the wsman.Messages struct by passing in the client parameters using the client.Parameters struct. For example:

```go
clientParams := client.Parameters{
    Target:             "192.168.0.120",
    Username:           "admin",
    Password:           "amtP@ssw0rd",
    UseDigest:          true,
    UseTLS:             true,
    SelfSignedAllowed:  true,
    LogAMTMessages:     true,
}
wsmanMessages := wsman.NewMessages(clientParams)
```

Next, you can call the various methods of the wsman.Messages struct.  Go-wsman-messages will authenticate with AMT using the client parameters provided and send the message to the Intel® AMT device and handle the response, returning a package specific Response struct or error message.  For example, to get the general settings of an Intel® AMT device, you can do:

```go
response, err := wsmanMessages.AMT.GeneralSettings.Get()
if err != nil {
    // handle error
}
// process response
```

# Dev tips for passing CI Checks

- Ensure code is formatted correctly with `gofmt -s -w ./` 
- Ensure all unit tests pass with `go test ./...`
- Ensure code has been linted with `docker run --rm -v ${pwd}:/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run -v`