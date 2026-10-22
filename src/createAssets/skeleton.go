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
		"{{ SOFTWARE NAME }}":    softwarename,
		"{{ PACKAGE VERSION }}":  helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}":  helpers.ReleaseNumber,
		"{{ GO VERSION }}":       helpers.GoVersion,
		"{{ GO MAJOR MINOR }}":   helpers.ExtractMajorMinorVersionString(helpers.GoVersion),
		"{{ BINARY NAME }}":      helpers.BinaryName,
		"{{ SECTION }}":          helpers.Section,
		"{{ DESCRIPTION }}":      helpers.Description,
		"{{ RELEASE DATE }}":     helpers.ReleaseDate,
		"{{ RELEASE DATE ISO }}": helpers.ReleaseDateISO,
	}

	fmt.Printf("Stub: %s\n", hftx.Yellow("Skeleton"))

	paths := []string{
		"docs/ISSUES.md",
		"docs/ROADMAP.md",
		"docs/TODO.md",
		"docs/CHANGELOG.md",
		"images/.gitkeep",
		"go.version",
		"LICENSE",
		"README.md",
		"gitignore",
		"dontexec.sh",
		"version-bump.sh",
		"src/build.sh",
		"src/go.mod.tmpl",
		"src/main.go.tmpl",
		"src/updateBuildDeps.sh",
		"src/cmd/root.go.tmpl",
		"src/cmd/completion.go.tmpl",
	}

	// The skeleton assets mirror stubber's own layout (docs/, images/), so each
	// asset path is also its output path, bar the renames below.
	for _, pathloop := range paths {
		filename := pathloop

		switch {
		case pathloop == "gitignore":
			filename = ".gitignore"

		case strings.HasSuffix(pathloop, ".tmpl"):
			filename = strings.TrimSuffix(pathloop, ".tmpl")
		}

		if err := assets.ProcessEmbeddedAsset("skeleton/"+pathloop, filename, placeholders); err != nil {
			return err
		}

		if strings.HasSuffix(filename, ".sh") {
			_ = os.Chmod(filename, fs.FileMode(0755))
		}
	}

	return nil
}
