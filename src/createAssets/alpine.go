package createAssets

import (
	"fmt"
	"os"
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
		"{{ BINARY NAME }}":     helpers.BinaryName,
		"{{ GO VERSION }}":      helpers.GoVersion,
	}

	fmt.Printf("Stub: %s\n", helpers.Yellow("Alpine"))
	paths := []string{"APKBUILD", "Makefile", "post-install", "pre-install", "pre-upgrade", "post-upgrade", "pre-deinstall", "post-deinstall"}
	//	if err = templates.ProcessEmbeddedAsset(filepath.Join(helpers.RootDir, "apk", "APKBUILD"), filepath.Join("__alpine", "APKBUILD"), placeholders); err == nil {

	for _, pathloop := range paths {
		targetFname := ""
		// target filename is different when dealing with the install scripts (that is, everything except APKBUILD and the Makefile)
		if pathloop != "APKBUILD" && pathloop != "Makefile" {
			targetFname = helpers.BinaryName + "." + pathloop
		} else {
			targetFname = pathloop
		}
		if err = templates.ProcessEmbeddedAsset(filepath.Join("apk", pathloop), filepath.Join("__alpine", targetFname), placeholders); err != nil {
			return err
		}
		os.Chmod(filepath.Join("__alpine", targetFname), os.FileMode(0755))
	}
	return nil
}
