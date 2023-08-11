//go:generate go-bindata -pkg templates -o assets.go -prefix ../assets ../assets/...

package templates

import (
	"fmt"
	"os"
	"strings"
	"stubber/helpers"
)

// Function to replace placeholder variables in a line
func replacePlaceholders(line, placeholder, value string) string {
	return strings.ReplaceAll(line, placeholder, value)
}

// Function to process an embedded asset and replace placeholders
func ProcessEmbeddedAsset(inputPath, outputPath string, placeholders map[string]string) error {
	if !helpers.Quiet {
		fmt.Printf("File: %s -> %s ... ", helpers.White(inputPath), helpers.White(outputPath))
	}

	// Read the embedded input file
	data, err := Asset(inputPath)
	if err != nil {
		fmt.Printf("error creating the input file: %s \n", helpers.Red(err.Error()))
		//return fmt.Errorf("error reading embedded input file '%s': %s", inputPath, err)
		return err
	}

	// Create the output file
	output, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("error creating the input file: %s \n", helpers.Red(err.Error()))
		return err
		//return fmt.Errorf("error creating output file '%s': %s", outputPath, err)
	}
	defer output.Close()

	// Convert the embedded data to a string
	content := string(data)

	// Replace placeholder variables in the content
	for placeholder, value := range placeholders {
		content = replacePlaceholders(content, placeholder, value)
	}

	// Write the modified content to the output file
	_, err = output.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to output file '%s': %s", outputPath, err)
	}

	if !helpers.Quiet {
		fmt.Printf("%s\n", helpers.Green("done"))
	}
	return nil
}
