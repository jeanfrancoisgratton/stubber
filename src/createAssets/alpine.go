package createAssets

import (
	"fmt"
	"path/filepath"
	"stubber/helpers"
	"stubber/templates"
)

func stubAlpine(softwarename string) error {
	var err error

	// alpine uses x86_64, not amd64
	arch := helpers.Arch
	if arch == "amd64" {
		arch = "x86_64"
	}
	placeholders := map[string]string{
		"{{ MAINTAINER }}":      helpers.Maintainer,
		"{{ PACKAGER }}":        helpers.Packager,
		"{{ SOFTWARE NAME }}":   softwarename,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ DESCRIPTION }}":     helpers.Description,
		"{{ ARCHITECTURE }}":    arch,
		"{{ BINARY NAME }}":     helpers.BinaryName,
		"{{ GO VERSION }}":      helpers.GoVersion,
	}

	fmt.Printf("Stub: %s\n", helpers.Yellow("Alpine"))
	//	if err = templates.ProcessEmbeddedAsset(filepath.Join(helpers.RootDir, "apk", "APKBUILD"), filepath.Join("__alpine", "APKBUILD"), placeholders); err == nil {
	if err = templates.ProcessEmbeddedAsset(filepath.Join("apk", "APKBUILD"), filepath.Join("__alpine", "APKBUILD"), placeholders); err == nil {
		// Alpine's Makefile takes amd64 for arch name, not x86_64
		arch = "amd64"
		err = templates.ProcessEmbeddedAsset(filepath.Join("apk", "Makefile"), filepath.Join("__alpine", "Makefile"), placeholders)
	}
	return err
}
