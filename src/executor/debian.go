package executor

import (
	"stubber/helpers"
	"stubber/templates"
)

func stubDebian(softwarename string) error {
	var err error
	placeholders := map[string]string{
		"{{ GO VERSION }}":      helpers.GoVersion,
		"{{ ARCHITECTURE }}":    helpers.Arch,
		"{{ SOFTWARE NAME }}":   softwarename,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ MAINTAINER }}":      helpers.Maintainer,
		"{{ DESCRIPTION }}":     helpers.Description,
		"{{ PACKAGE SECTION }}": helpers.Section,
		"{{ DEPENDENCIES }}":    helpers.Dependencies,
	}
	paths := []string{"1.install-build-deps.sh", "2.build_binary.sh", "3.restore_repo.sh", "control"}

	for _, pathloop := range paths {
		if err = templates.ProcessEmbeddedAsset("deb/"+pathloop, "__debian/"+pathloop, placeholders); err != nil {
			return err
		}
	}
	return nil
}
