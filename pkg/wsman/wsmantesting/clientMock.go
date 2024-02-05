package wsmantesting

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// MockClient is a mock implementation of the wsman.Client interface for testing.
type MockClient struct {
	CurrentMessage   string
	PackageUnderTest string
}

func (c *MockClient) Post(msg string) ([]byte, error) {
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/" + c.PackageUnderTest + "/" + strings.ToLower(c.CurrentMessage) + ".xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer xmlFile.Close()
	// read file into string
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	// strip carriage returns and new line characters
	xmlData = []byte(strings.ReplaceAll(string(xmlData), "\r\n", ""))

	// Simulate a successful response for testing.
	return []byte(xmlData), nil
}
