package executor

import (
	"fmt"
	"stubber/helpers"
	"stubber/templates"
)

func stubRedHat(softwarename string) error {
	var err error
	placeholders := map[string]string{
		"{{ SOFTWARE NAME }}":   softwarename,
		"{{ GO VERSION }}":      helpers.GoVersion,
		"{{ ARCHITECTURE }}":    helpers.Arch,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ BINARY NAME }}":     helpers.BinaryName,
		"{{ SECTION }}":         helpers.Section,
		"{{ DESCRIPTION }}":     helpers.Description,
	}

	fmt.Printf("Stub: %s\n", helpers.Yellow("RedHat"))
	if err = templates.ProcessEmbeddedAsset("rpm/stubber.spec", softwarename+".spec", placeholders); err == nil {
		err = templates.ProcessEmbeddedAsset("rpm/rpm-install-build-deps.sh", "rpm-install-build-deps.sh", placeholders)
	}
	return err
}
