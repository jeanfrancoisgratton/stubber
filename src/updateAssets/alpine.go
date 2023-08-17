package updateAssets

import (
	"fmt"
	"path/filepath"
	"stubber/helpers"
)

func updateAlpine() error {
	var err error
	goarch := helpers.Arch

	// alpine uses x86_64, not amd64
	arch := helpers.Arch
	if arch == "amd64" {
		arch = "x86_64"
		goarch = "amd64"
	}
	placeholders := map[string]string{
		"# Maintainer": "# Maintainer: " + helpers.Maintainer,
		"# Packager:":  "# Packager: " + helpers.Packager,
		"pkgver":       "pkgver=" + helpers.VersionNumber,
		"pkgrel":       "pkgrel=" + helpers.ReleaseNumber,
		"arch":         "arch=" + goarch,
	}

	fmt.Printf("Stub: %s\n", helpers.Yellow("Alpine"))
	paths := {"APKBUILD", "Makefile"}

	for _, path : range paths
	if err = replaceStrings(filepath.Join(helpers.RootDir, "__alpine", "APKBUILD"), placeholders); err == nil {
		err = replaceStrings(filepath.Join(helpers.RootDir, "__alpine", "Makefile"), placeholders)
	}
	return err
}
