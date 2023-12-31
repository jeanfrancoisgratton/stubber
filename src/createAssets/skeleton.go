package createAssets

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"stubber/helpers"
	"stubber/templates"
)

func stubSkeleton(softwarename string) error {
	var err error

	placeholders := map[string]string{
		"{{ SOFTWARE NAME }}":   softwarename,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ GO VERSION }}":      helpers.GoVersion,
		"{{ ARCHITECTURE }}":    helpers.Arch,
		"{{ GO MAJOR MINOR }}":  helpers.ExtractMajorMinorVersionString(helpers.GoVersion),
		"{{ BINARY NAME }}":     helpers.BinaryName,
		"{{ SECTION }}":         helpers.Section,
		"{{ DESCRIPTION }}":     helpers.Description,
	}

	fmt.Printf("Stub: %s\n", helpers.Yellow("Skeleton"))
	paths := []string{"FIXME.md", "go.version", "CHANGELOG.md", "LICENSE", "PACKAGING.md", "README.md",
		"TODO.md", "gitignore", "src/build.sh", "src/go.mod", "src/main.go", "src/upgrade_pkgs.sh", "src/cmd/root.go"}

	for _, pathloop := range paths {
		// We have to add a special condition here because source and target filenames differ for gitignore
		filename := pathloop
		if pathloop == "gitignore" {
			filename = ".gitignore"
		}
		if err = templates.ProcessEmbeddedAsset("skeleton/"+pathloop, filename, placeholders); err != nil {
			return err
		}
		if strings.HasSuffix(filename, ".sh") {
			os.Chmod(filename, fs.FileMode(0755))

		}
	}
	return nil
}
