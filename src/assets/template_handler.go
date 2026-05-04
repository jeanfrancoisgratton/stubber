package assets

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v5/terminalfx"

	"stubber/helpers"
)

// replacePlaceholders replaces every occurrence of placeholder with value.
func replacePlaceholders(line, placeholder, value string) string {
	return strings.ReplaceAll(line, placeholder, value)
}

// ProcessEmbeddedAsset reads an embedded template asset, replaces placeholders,
// and writes the processed content to outputPath.
func ProcessEmbeddedAsset(inputPath, outputPath string, placeholders map[string]string) *cerr.CustomError {
	if !helpers.Quiet {
		fmt.Printf("File: %s -> %s ... ", hftx.White(inputPath), hftx.White(outputPath))
	}

	data, err := FS.ReadFile(filepath.ToSlash(inputPath))
	if err != nil {
		return &cerr.CustomError{Title: "Error reading the input file", Message: err.Error()}
	}

	dir := filepath.Dir(outputPath)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return &cerr.CustomError{Title: "Error creating directory", Message: err.Error()}
		}
	}

	content := string(data)
	for placeholder, value := range placeholders {
		content = replacePlaceholders(content, placeholder, value)
	}

	if err := os.WriteFile(outputPath, []byte(content), 0644); err != nil {
		return &cerr.CustomError{Title: "Error writing to the output file", Message: err.Error()}
	}

	if !helpers.Quiet {
		fmt.Printf("%s\n", hftx.Green("done"))
	}
	return nil
}
