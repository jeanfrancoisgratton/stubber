package createAssets

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"stubber/helpers"
	"stubber/templates"

	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
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

	fmt.Printf("Stub: %s\n", hf.Yellow("Skeleton"))
	paths := []string{"ISSUES.md", "go.version", "CHANGELOG.md", "LICENSE", "README.md",
		"gitignore", "src/build.sh", "src/go.mod", "src/main.go", "src/updateBuildDeps.sh", "src/checkImports.sh", "src/cmd/root.go",
		/*".github/workflows/publish_release.yaml.disabled"*/}

	for _, pathloop := range paths {
		// We have to add a special condition here because source and target filenames differ for some of the files
		filename := pathloop

		// We'll have to consider using a switch {} block here if it keeps growing...
		if pathloop == "gitignore" {
			filename = ".gitignore"
		}
		//if pathloop == "publish_release.yaml.disabled" {
		//	filename = filepath.Join(".github", "workflows", "publish_release.yaml.disabled")
		//}

		if err := templates.ProcessEmbeddedAsset("skeleton/"+pathloop, filename, placeholders); err != nil {
			return err
		}
		if strings.HasSuffix(filename, ".sh") {
			os.Chmod(filename, fs.FileMode(0755))
		}
	}
	return nil
}
