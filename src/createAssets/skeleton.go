package createAssets

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v5/terminalfx"
	"stubber/assets"
	"stubber/helpers"
)

func stubSkeleton(softwarename string) *cerr.CustomError {
	placeholders := map[string]string{
		"{{ SOFTWARE NAME }}":   softwarename,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ GO VERSION }}":      helpers.GoVersion,
		"{{ GO MAJOR MINOR }}":  helpers.ExtractMajorMinorVersionString(helpers.GoVersion),
		"{{ BINARY NAME }}":     helpers.BinaryName,
		"{{ SECTION }}":         helpers.Section,
		"{{ DESCRIPTION }}":     helpers.Description,
		"{{ RELEASE DATE }}":    helpers.ReleaseDate,
	}

	fmt.Printf("Stub: %s\n", hftx.Yellow("Skeleton"))
	paths := []string{"ISSUES.md", "go.version", "CHANGELOG.md", "LICENSE", "README.md",
		"gitignore", "src/build.sh", "src/go.mod.tmpl", "src/main.go", "src/updateBuildDeps.sh", "src/cmd/root.go"}

	for _, pathloop := range paths {
		// We have to add a special condition here because source and target filenames differ for some of the files
		filename := pathloop

		// We'll have to consider using a switch {} block here if it keeps growing...
		switch {
		case pathloop == "gitignore":
			filename = ".gitignore"
		case pathloop == "src/go.mod.tmpl.tmpl":
			filename = "src/go.mod"
		}

		if err := assets.ProcessEmbeddedAsset("skeleton/"+pathloop, filename, placeholders); err != nil {
			return err
		}
		if strings.HasSuffix(filename, ".sh") {
			os.Chmod(filename, fs.FileMode(0755))
		}
	}
	return nil
}
