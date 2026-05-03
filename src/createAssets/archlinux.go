// stubber
// Written by J.F.Gratton <jean-francois@famillegratton.net>
// Original filename: src/createAssets/archlinux.go
// Original timestamp: 2026/05/02 20:23:35

package createAssets

import (
	"fmt"
	"os"
	"path/filepath"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v3/terminalfx"
	"stubber/helpers"
	"stubber/templates"
)

func stubArchLinux(softwarename string) *cerr.CustomError {
	placeholders := map[string]string{
		"{{ SOFTWARE NAME }}":   softwarename,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ DESCRIPTION }}":     helpers.Description,
		"{{ BINARY NAME }}":     helpers.BinaryName,
	}

	fmt.Printf("Stub: %s\n", hftx.Yellow("ArchLinux"))
	paths := []string{"1.install-build-deps.sh", "2.build-package.sh", "PKGBUILD"}

	for _, pathloop := range paths {

		if err := templates.ProcessEmbeddedAsset(filepath.Join("arch", pathloop), filepath.Join("__archlinux", pathloop), placeholders); err != nil {
			return err
		}
		os.Chmod(filepath.Join("__archlinux", pathloop), os.FileMode(0755))
	}
	return nil
}
