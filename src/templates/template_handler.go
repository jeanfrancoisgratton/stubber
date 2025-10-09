//go:generate go-bindata -pkg templates -o assets.go -prefix ../assets ../assets/...

package templates

import (
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v3/terminalfx"
	"os"
	"path/filepath"
	"strings"
	"stubber/helpers"
)

// Function to replace placeholder variables in a line
func replacePlaceholders(line, placeholder, value string) string {
	return strings.ReplaceAll(line, placeholder, value)
}

// Function to process an embedded asset and replace placeholders
func ProcessEmbeddedAsset(inputPath, outputPath string, placeholders map[string]string) *cerr.CustomError {
	//var err error
	if !helpers.Quiet {
		fmt.Printf("File: %s -> %s ... ", hftx.White(inputPath), hftx.White(outputPath))
	}

	// Read the embedded input file
	data, aErr := Asset(inputPath)
	if aErr != nil {
		//fmt.Printf("error reading the input file: %s \n", hf.Red(aErr.Error()))
		return &cerr.CustomError{Title: "Error reading the input file", Message: aErr.Error()}
	}

	// Create the target directory if it does not yet exist
	dir := filepath.Dir(outputPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return &cerr.CustomError{Title: "Error creating directory", Message: err.Error()}
		}
	}
	// Create the output file
	output, err := os.Create(outputPath)
	if err != nil {
		return &cerr.CustomError{Title: "Error creating the output file", Message: err.Error()}
	}
	defer output.Close()

	// Convert the embedded data to a string
	content := string(data)

	// Replace placeholder variables in the content
	for placeholder, value := range placeholders {
		content = replacePlaceholders(content, placeholder, value)
	}

	// Write the modified content to the output file
	if _, err := output.WriteString(content); err != nil {
		return &cerr.CustomError{Title: "Error writing to the output file", Message: err.Error()}
	}

	if !helpers.Quiet {
		fmt.Printf("%s\n", hftx.Green("done"))
	}
	return nil
}
