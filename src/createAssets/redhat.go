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

func stubRedHat(softwarename string) *cerr.CustomError {
	var err *cerr.CustomError

	placeholders := map[string]string{
		"{{ SOFTWARE NAME }}":   softwarename,
		"{{ GO VERSION }}":      helpers.GoVersion,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ BINARY NAME }}":     helpers.BinaryName,
		"{{ SECTION }}":         helpers.Section,
		"{{ DESCRIPTION }}":     helpers.Description,
		"{{ URL }}":             helpers.Url,
		"{{ RELEASE DATE }}":    helpers.ReleaseDate,
	}

	paths := []string{"specfile", "rpmbuild-deps.sh", "Makefile", "updateChangelog.sh"}

	fmt.Printf("Stub: %s\n", hftx.Yellow("RedHat"))
	for _, pathloop := range paths {
		filename := pathloop
		if pathloop == "specfile" {
			filename = softwarename + ".spec"
		}
		if err = assets.ProcessEmbeddedAsset(filepath.Join("rpm", pathloop), filepath.Join("__redhat", filename), placeholders); err != nil {
			return err
		}
	}

	os.Chmod(filepath.Join("__redhat", "rpmbuild-deps.sh"), os.FileMode(0755))
	os.Chmod(filepath.Join("__redhat", "updateChangelog.sh"), os.FileMode(0755))
	os.Chmod("__redhat", os.FileMode(0755))
	return err
}
