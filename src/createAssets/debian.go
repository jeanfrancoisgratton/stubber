package createAssets

import (
	"fmt"
	"os"
	"path/filepath"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v5/terminalfx"
	"stubber/assets"
	"stubber/helpers"
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
		"{{ URL }}":             helpers.Url,
		"{{ COPYRIGHT YEAR }}":  helpers.CopyrightYear(),
	}
	paths := []string{"install-build-deps.sh", "restore_repo.sh", "Makefile", "control", "copyright", "preinst", "prerm", "postinst", "postrm"}

	fmt.Printf("Stub: %s\n", hftx.Yellow("Debian"))
	for _, pathloop := range paths {
		if err := assets.ProcessEmbeddedAsset(filepath.Join("debian", pathloop), filepath.Join("__debian", pathloop), placeholders); err != nil {
			return err
		}
		// copyright is documentation, not something dpkg or the build ever
		// executes; the Makefile installs it 0644 into the package.
		if pathloop == "copyright" {
			continue
		}
		os.Chmod(filepath.Join("__debian", pathloop), os.FileMode(0755))
	}
	return nil
}
