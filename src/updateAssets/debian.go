package updateAssets

import (
	"fmt"
	"path/filepath"
	"strings"
	"stubber/helpers"
)

func updateDebian() error {
	var err error
	//var goarch string

	// Debian uses amd64, not x86_64
	arch := strings.ToLower(helpers.Arch)
	if arch == "x86_64" {
		arch = "amd64"
	}

	placeholders := map[string]string{
		"sudo /opt/bin/install_golang.sh": "sudo /opt/bin/install_golang.sh " + helpers.GoVersion + " " + arch,
		"{{ ARCHITECTURE }}":              arch,
		"{{ PACKAGE VERSION }}":           helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}":           helpers.ReleaseNumber,
		"{{ MAINTAINER }}":                helpers.Maintainer,
		"{{ DESCRIPTION }}":               helpers.Description,
		"{{ PACKAGE SECTION }}":           helpers.Section,
		"{{ DEPENDENCIES }}":              helpers.Dependencies,
		"{{ BINARY NAME }}":               helpers.BinaryName,
	}
	paths := []string{"1.install-build-deps.sh", "2.build_binary.sh", "control"}

	fmt.Printf("Stub: %s\n", helpers.Yellow("Debian"))
	for _, pathloop := range paths {
		// err = replaceStrings(filepath.Join(helpers.RootDir, "__alpine", "Makefile"), placeholders)
		if err = replaceStrings(filepath.Join(helpers.RootDir, "__debian", pathloop), placeholders); err != nil {
			//if err = templates.ProcessEmbeddedAsset(filepath.Join("deb", pathloop), filepath.Join(helpers.RootDir, "__debian", pathloop), placeholders); err != nil {
			return err
		}
	}
	return nil
}
