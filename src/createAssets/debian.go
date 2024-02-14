package createAssets

import (
	"fmt"
	"os"
	"path/filepath"
	"stubber/helpers"
	"stubber/templates"
)

func stubDebian(softwarename string) error {
	var err error

	placeholders := map[string]string{
		"{{ GO VERSION }}":      helpers.GoVersion,
		"{{ ARCHITECTURE }}":    "amd64",
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
		if err = templates.ProcessEmbeddedAsset(filepath.Join("deb", pathloop), filepath.Join("__debian", pathloop), placeholders); err != nil {
			return err
		}
		os.Chmod(filepath.Join("__debian", pathloop), os.FileMode(0755))
	}
	return nil
}
