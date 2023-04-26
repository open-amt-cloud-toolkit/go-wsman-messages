# go-wsman-messages

> Disclaimer: Production viable releases are tagged and listed under 'Releases'.  All other check-ins should be considered 'in-development' and should not be used in production

This repository contains a Go library that creates properly formatted wsman messages to send to Intel® Active Management Technology (AMT) capable platforms. These messages are based on the AMT SDK documentation, which can be found [here](https://software.intel.com/content/www/us/en/develop/articles/intel-active-management-technology-software-development-kit-sdk.html).

This library provides an easy to use API that returns an XML string that is ready to be sent to an Intel® AMT device. It supports calls into AMT, IPS, and CIM classes supported by Intel® AMT devices.

## How to use it

To use this library, you need to import it in your Go project:
``` go
import "github.com/open-amt-cloud-toolkit/go-wsman-messages"
```

Then, you can create an instance of the message class you want to use, such as `amt.NewMessages()`, `ips.NewMessages()`, or `cim.NewMessages()`. For example:

```go
amtClass := amt.NewMessages()
```

Next, you can call the methods of the message class to get the XML string for the desired operation. For example, to get the general settings of an Intel® AMT device, you can do:

```go
message := amtClass.GeneralSettings.Get()
```

Finally, you can send the message to the Intel® AMT device using the wsman HTTP Client. For example:

```go
client := wsman.NewClient("http://localhost:16992/wsman", username, password, true)
response, err := client.Post(message) 
if err != nil {
// handle error
}
// process response
```


# Dev tips for passing CI Checks

- Ensure code is formatted correctly with `gofmt -s -w ./` 
- Ensure all unit tests pass with `go test ./...`
- Ensure code has been linted with `docker run --rm -v ${pwd}:/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run -v`