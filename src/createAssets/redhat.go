package createAssets

import (
	"fmt"
	"path/filepath"
	"stubber/helpers"
	"stubber/templates"
)

func stubRedHat(softwarename string) error {
	var err error

	placeholders := map[string]string{
		"{{ SOFTWARE NAME }}":   softwarename,
		"{{ GO VERSION }}":      helpers.GoVersion,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ BINARY NAME }}":     helpers.BinaryName,
		"{{ SECTION }}":         helpers.Section,
		"{{ DESCRIPTION }}":     helpers.Description,
		"{{ URL }}":             helpers.Url,
	}

	fmt.Printf("Stub: %s\n", helpers.Yellow("RedHat"))
	if err = templates.ProcessEmbeddedAsset(filepath.Join("rpm", "specfile"), softwarename+".spec", placeholders); err == nil {
		err = templates.ProcessEmbeddedAsset(filepath.Join("rpm", "rpmbuild-deps.sh"), "rpmbuild-deps.sh", placeholders)
	}
	return err
}
