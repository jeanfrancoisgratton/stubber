package updateAssets

import (
	"fmt"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v3/terminalfx"
	"os"
	"path/filepath"
	"stubber/helpers"
)

// updateSkeleton2 is a placeholder function in case it is needed.
// The "real" function, the one that is effective, is below... for now
func updateSkeleton2() error {
	var err error

	placeholders := map[string]string{
		"{{ GO VERSION }}": helpers.GoVersion,
	}

	fmt.Printf("Stub: %s\n", hftx.Yellow("Skeleton"))
	// This for..loop seems like overkill, but I plan on further, potential file modifications, so here we are...
	paths := []string{"go.version"}

	for _, pathloop := range paths {
		//filename := pathloop
		if err = replaceStrings(pathloop, placeholders); err != nil {
			return err
		}
	}
	return nil
}

func updateSkeleton() error {
	if err := os.WriteFile(filepath.Join(helpers.RootDir, "go.version"), []byte(helpers.GoVersion), 0644); err != nil {
		return err
	}
	return nil
}
