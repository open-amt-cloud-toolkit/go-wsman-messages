# wsman-go

# Dev tips for passing CI Checks

- Ensure code is formatted correctly with `gofmt -s -w ./` 
- Ensure all unit tests pass with `go test ./...`
- Ensure code has been linted with `docker run --rm -v ${pwd}:/app -w /app golangci/golangci-lint:v1.52.2 golangci-lint run -v`