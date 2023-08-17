package updateAssets

import (
	"fmt"
	"path/filepath"
	"stubber/helpers"
	"stubber/templates"
)

func updateRedHat(softwarename string) error {
	var err error
	arch := helpers.Arch
	if arch == "amd64" {
		arch = "x86_64"
	}
	placeholders := map[string]string{
		"{{ GO VERSION }}":      helpers.GoVersion,
		"{{ ARCHITECTURE }}":    arch,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ BINARY NAME }}":     helpers.BinaryName,
		"{{ SECTION }}":         helpers.Section,
		"{{ DESCRIPTION }}":     helpers.Description,
		"{{ URL }}":             helpers.Url,
	}

	fmt.Printf("Stub: %s\n", helpers.Yellow("RedHat"))
	if err = templates.ProcessEmbeddedAsset(filepath.Join("rpm", "specfile"), softwarename+".spec", placeholders); err == nil {
		// The dependency script takes amd64 as an arch, not x86_64
		if arch == "x86_64" {
			placeholders["{{ ARCHITECTURE }} "] = "arch=amd64"
		}
		err = templates.ProcessEmbeddedAsset(filepath.Join("rpm", "rpmbuild-deps.sh"), filepath.Join(helpers.RootDir, "rpmbuild-deps.sh"), placeholders)
	}
	return nil
}
