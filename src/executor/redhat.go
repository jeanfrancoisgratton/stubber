package executor

import (
	"fmt"
	"stubber/helpers"
	"stubber/templates"
)

func stubRedHat(softwarename string) error {
	var err error

	arch := helpers.Arch
	if arch == "amd64" {
		arch = "x86_64"
	}
	placeholders := map[string]string{
		"{{ SOFTWARE NAME }}":   softwarename,
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
	if err = templates.ProcessEmbeddedAsset("rpm/specfile", softwarename+".spec", placeholders); err == nil {
		err = templates.ProcessEmbeddedAsset("rpm/rpmbuild-deps.sh", "rpmbuild-deps.sh", placeholders)
	}
	return err
}
