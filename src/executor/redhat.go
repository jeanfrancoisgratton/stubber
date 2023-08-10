package executor

import (
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
	if err = templates.ProcessEmbeddedAsset("skeleton/stubber.spec", "softwarename", placeholders); err == nil {
		err = templates.ProcessEmbeddedAsset("skeleton/rpm-install-build-deps.sh", "rpm-install-build-deps.sh", placeholders)
	}
	return err
}
