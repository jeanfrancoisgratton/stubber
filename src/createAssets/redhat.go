package createAssets

import (
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
	"path/filepath"
	"stubber/helpers"
	"stubber/templates"
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

	paths := []string{"specfile", "rpmbuild-deps.sh" /*, "tito.props"*/}

	for _, pathloop := range paths {
		filename := pathloop
		if pathloop == "specfile" {
			filename = softwarename + ".spec"
		}
		//if pathloop == "tito.props" {
		//	filename = filepath.Join(".tito", "tito.props")
		//}
		if err = templates.ProcessEmbeddedAsset(filepath.Join("rpm", pathloop), filename, placeholders); err != nil {
			return err
		}
	}
	fmt.Printf("Stub: %s\n", hf.Yellow("RedHat"))
	if err = templates.ProcessEmbeddedAsset(filepath.Join("rpm", "specfile"), softwarename+".spec", placeholders); err == nil {
		err = templates.ProcessEmbeddedAsset(filepath.Join("rpm", "rpmbuild-deps.sh"), "rpmbuild-deps.sh", placeholders)
	}
	return err
}
