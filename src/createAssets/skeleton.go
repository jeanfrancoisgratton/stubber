// stubber
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// execDriver.go, jfgratton : 2023-06-27

package createAssets

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"stubber/assets"
	"stubber/helpers"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v5/terminalfx"
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

	paths := []string{
		"ISSUES.md",
		"ROADMAP.md",
		"TODO.md",
		"go.version",
		"CHANGELOG.md",
		"LICENSE",
		"README.md",
		"gitignore",
		"dontexec.sh",
		"src/build.sh",
		"src/go.mod.tmpl",
		"src/main.go.tmpl",
		"src/updateBuildDeps.sh",
		"src/cmd/root.go.tmpl",
		"src/cmd/completion.go.tmpl",
	}

	// CHANGELOG/ISSUES/ROADMAP/TODO live under docs/, not the project root --
	// stubber's own layout (see docs/, images/) is the template for this.
	docs := map[string]bool{
		"ISSUES.md":    true,
		"ROADMAP.md":   true,
		"TODO.md":      true,
		"CHANGELOG.md": true,
	}

	for _, pathloop := range paths {
		filename := pathloop

		switch {
		case pathloop == "gitignore":
			filename = ".gitignore"

		case strings.HasSuffix(pathloop, ".tmpl"):
			filename = strings.TrimSuffix(pathloop, ".tmpl")
		}

		if docs[filename] {
			filename = "docs/" + filename
		}

		if err := assets.ProcessEmbeddedAsset("skeleton/"+pathloop, filename, placeholders); err != nil {
			return err
		}

		if strings.HasSuffix(filename, ".sh") {
			_ = os.Chmod(filename, fs.FileMode(0755))
		}
	}

	// images/ has no stub content of its own, so ProcessEmbeddedAsset never
	// creates it; carry it into every new project regardless.
	if err := os.MkdirAll("images", fs.FileMode(0755)); err != nil {
		return &cerr.CustomError{Title: "Unable to create the images directory", Message: err.Error()}
	}

	return nil
}
