//go:generate go-bindata -pkg templates -o assets.go -prefix ../assets ../assets/...

package templates

import (
	"strings"
)

// Function to replace placeholder variables in a line
func replacePlaceholders(line, placeholder, value string) string {
	return strings.ReplaceAll(line, placeholder, value)
}

// Function to process an embedded asset and replace placeholders
func processEmbeddedAsset(inputPath string, placeholders map[string]string) (string, error) {
	// Read the embedded input file
	data, err := Asset(inputPath)
	if err != nil {
		return "", err
	}

	// Convert the embedded data to a string
	content := string(data)

	// Replace placeholder variables in the content
	for placeholder, value := range placeholders {
		content = replacePlaceholders(content, placeholder, value)
	}

	return content, nil
}
