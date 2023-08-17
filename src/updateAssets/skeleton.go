package updateAssets

import (
	"fmt"
	"os"
	"path/filepath"
	"stubber/helpers"
	"stubber/templates"
)

func updateSkeleton() error {
	var err error

	placeholders := map[string]string{
		"{{ GO VERSION }}": helpers.GoVersion,
	}

	fmt.Printf("Stub: %s\n", helpers.Yellow("Skeleton"))
	// This for..loop seems like overkill, but I plan on further, potential file modifications, so here we are...
	paths := []string{"go.version"}

	for _, pathloop := range paths {
		// We have to add a special condition here because source and target filenames differ for gitignore
		filename := pathloop
		if err = templates.ProcessEmbeddedAsset(filepath.Join(helpers.RootDir, pathloop), filepath.Join(helpers.RootDir, filename+".new"), placeholders); err != nil {
			return err
		} else {
			if err = os.Rename(filename+".new", filename); err != nil {
				return err
			}
		}
	}
	return nil
}
