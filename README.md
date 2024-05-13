# go-wsman-messages

![CodeQL](https://img.shields.io/github/actions/workflow/status/open-amt-cloud-toolkit/go-wsman-messages/codeql-analysis.yml?style=for-the-badge&label=CodeQL&logo=github)
![Build](https://img.shields.io/github/actions/workflow/status/open-amt-cloud-toolkit/go-wsman-messages/ci.yml?style=for-the-badge&logo=github)
![Codecov](https://img.shields.io/codecov/c/github/open-amt-cloud-toolkit/go-wsman-messages?style=for-the-badge&logo=codecov)
[![OSSF-Scorecard Score](https://img.shields.io/ossf-scorecard/github.com/open-amt-cloud-toolkit/go-wsman-messages?style=for-the-badge&label=OSSF%20Score)](https://api.securityscorecards.dev/projects/github.com/open-amt-cloud-toolkit/go-wsman-messages)
[![Discord](https://img.shields.io/discord/1063200098680582154?style=for-the-badge&label=Discord&logo=discord&logoColor=white&labelColor=%235865F2&link=https%3A%2F%2Fdiscord.gg%2FqmTWWFyA)](https://discord.gg/qmTWWFyA)


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
- Ensure code has been gci'd with `gci.exe write --skip-generated -s standard -s default .`
- Ensure code has been linted with `docker run --rm -v ${pwd}:/app -w /app golangci/golangci-lint:latest golangci-lint run -v`