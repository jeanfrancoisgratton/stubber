package executor

import (
	"fmt"
	"strings"
	"stubber/helpers"
	"stubber/templates"
)

func stubDebian(softwarename string) error {
	var err error

	// Debian uses amd64, not x86_64
	arch := strings.ToLower(helpers.Arch)
	if arch == "x86_64" {
		arch = "amd64"
	}

	placeholders := map[string]string{
		"{{ GO VERSION }}":      helpers.GoVersion,
		"{{ ARCHITECTURE }}":    arch,
		"{{ SOFTWARE NAME }}":   softwarename,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ MAINTAINER }}":      helpers.Maintainer,
		"{{ DESCRIPTION }}":     helpers.Description,
		"{{ PACKAGE SECTION }}": helpers.Section,
		"{{ DEPENDENCIES }}":    helpers.Dependencies,
		"{{ BINARY NAME }}":     helpers.BinaryName,
	}
	paths := []string{"1.install-build-deps.sh", "2.build_binary.sh", "3.restore_repo.sh", "control", "preinst", "prerm", "postinst", "postrm"}

	fmt.Printf("Stub: %s\n", helpers.Yellow("Debian"))
	for _, pathloop := range paths {
		if err = templates.ProcessEmbeddedAsset("deb/"+pathloop, "__debian/"+pathloop, placeholders); err != nil {
			return err
		}
	}
	return nil
}
