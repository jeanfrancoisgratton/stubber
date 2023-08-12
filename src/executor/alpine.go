package executor

import (
	"fmt"
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
	}

	fmt.Printf("Stub: %s\n", helpers.Yellow("Alpine"))
	if err = templates.ProcessEmbeddedAsset("apk/APKBUILD", "__alpine/APKBUILD", placeholders); err == nil {
		err = templates.ProcessEmbeddedAsset("apk/Makefile", "__alpine/Makefile", placeholders)
	}
	return err
}
