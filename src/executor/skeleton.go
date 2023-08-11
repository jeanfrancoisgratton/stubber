package executor

import (
	"fmt"
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
	paths := []string{"FIXME.md", "go.version", "IN THIS BRANCH.md", "LICENSE", "PACKAGING.md", "README.md",
		"ROADMAP.md", "TODO.md", "src/build.sh", "src/go.mod", "src/main.go", "src/upgrade_pkgs.sh", "src/cmd/root.go",
		"src/helpers/misc.go"}

	for _, pathloop := range paths {
		if err = templates.ProcessEmbeddedAsset("skeleton/"+pathloop, pathloop, placeholders); err != nil {
			return err
		}
	}
	return nil
}
