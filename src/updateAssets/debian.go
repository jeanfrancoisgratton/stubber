package updateAssets

import (
	"fmt"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"path/filepath"
	"strings"
	"stubber/helpers"
)

func updateDebian(softwarename string) error {
	var err error
	//var goarch string

	// Debian uses amd64, not x86_64
	arch := strings.ToLower(helpers.Arch)
	if arch == "x86_64" {
		arch = "amd64"
	}

	placeholders := map[string]string{
		// Control
		"Version:":      "Version: " + helpers.VersionNumber,
		"Architecture:": "Architecture: " + arch,
		"Maintainer:":   "Mainteainer: " + helpers.Maintainer,
		"Description:":  "Description: " + helpers.Description,
		"Section:":      "Section: " + helpers.Section,
		"Depends:":      "Depends: " + helpers.Dependencies,
		// 2.build_binary.sh
		"PKGDIR=": "PKGDIR=" + softwarename + "-" + helpers.VersionNumber + "-" + helpers.ReleaseNumber + "_" + arch,
		// 1.install-build-deps.sh
		"sudo /opt/bin/install_golang.sh": "sudo /opt/bin/install_golang.sh " + helpers.GoVersion + " " + arch,
	}
	paths := []string{"1.install-build-deps.sh", "2.build_binary.sh", "control"}

	fmt.Printf("Stub: %s\n", hf.Yellow("Debian"))
	for _, pathloop := range paths {
		// err = replaceStrings(filepath.Join(helpers.RootDir, "__alpine", "Makefile"), placeholders)
		if err = replaceStrings(filepath.Join(helpers.RootDir, "__debian", pathloop), placeholders); err != nil {
			//if err = templates.ProcessEmbeddedAsset(filepath.Join("deb", pathloop), filepath.Join(helpers.RootDir, "__debian", pathloop), placeholders); err != nil {
			return err
		}
	}
	return nil
}
