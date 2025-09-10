package createAssets

import (
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"os"
	"path/filepath"
	"stubber/helpers"
	"stubber/templates"
)

func stubDebian(softwarename string) *cerr.CustomError {

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
		"{{ RELEASE DATE }}":    helpers.ReleaseDate,
	}
	paths := []string{"1.install-build-deps.sh", "2.build_binary.sh", "3.restore_repo.sh", "control", "preinst", "prerm", "postinst", "postrm"}

	fmt.Printf("Stub: %s\n", hf.Yellow("Debian"))
	for _, pathloop := range paths {
		if err := templates.ProcessEmbeddedAsset(filepath.Join("deb", pathloop), filepath.Join("__debian", pathloop), placeholders); err != nil {
			return err
		}
		os.Chmod(filepath.Join("__debian", pathloop), os.FileMode(0755))
	}
	return nil
}
