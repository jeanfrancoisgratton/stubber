// stubber
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Original name: src/createAssets/refresh.go

package createAssets

import (
	"fmt"
	"os"
	"path/filepath"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v5/terminalfx"
	"stubber/assets"
	"stubber/helpers"
)

// RefreshStub re-renders only the stubs that were explicitly requested on the
// command line. It reuses the very same rendering routines as `create`, so any
// value carried in the helpers globals (already merged from the manifest and
// the flags that were actually passed) is applied uniformly.
//
// helpers.RootDir is expected to already point at the resolved project root.
func RefreshStub(softname string) *cerr.CustomError {
	if helpers.BinaryName == "" {
		helpers.BinaryName = softname
	}

	currentdir, err := os.Getwd()
	if err != nil {
		return &cerr.CustomError{Title: "Unable to getcwd", Message: err.Error()}
	}
	if err = os.Chdir(helpers.RootDir); err != nil {
		return &cerr.CustomError{Title: "Unable to enter the project root", Message: err.Error()}
	}
	defer os.Chdir(currentdir)

	fmt.Printf("Refreshing stub for software %s in %s\n", hftx.Green(softname), hftx.Green(helpers.RootDir))

	// Alpine ( -a )
	if helpers.AlpineStub {
		if err = os.MkdirAll(filepath.Join(helpers.RootDir, "__alpine"), os.FileMode(0755)); err == nil {
			if e := stubAlpine(softname); e != nil {
				return e
			}
		}
	}

	// Debian ( -d )
	if helpers.DebianStub {
		if err = os.MkdirAll(filepath.Join(helpers.RootDir, "__debian"), os.FileMode(0755)); err == nil {
			if e := stubDebian(softname); e != nil {
				return e
			}
		}
	}

	// RedHat ( -r )
	if helpers.RedHatStub {
		if err = os.MkdirAll(filepath.Join(helpers.RootDir, "__redhat"), os.FileMode(0755)); err == nil {
			if e := stubRedHat(softname); e != nil {
				return e
			}
		}
	}

	// ArchLinux ( -A )
	if helpers.ArchLinuxStub {
		if err = os.MkdirAll(filepath.Join(helpers.RootDir, "__archlinux"), os.FileMode(0755)); err == nil {
			if e := stubArchLinux(softname); e != nil {
				return e
			}
		}
	}

	// Skeleton ( -k ): metadata only, so we never clobber real source or docs
	if helpers.SkeletonStub {
		if e := refreshSkeleton(); e != nil {
			return e
		}
	}

	return nil
}

// refreshSkeleton re-renders go.version and dontexec.sh. Every other skeleton
// file is either user-authored source (src/...) or user-edited content (README,
// CHANGELOG, ...) and must never be overwritten by a refresh.
//
// dontexec.sh qualifies because it is stubber-owned tooling, not content: it
// carries no project-specific value and nobody edits it by hand, so refreshing
// is how already-scaffolded projects pick it up.
func refreshSkeleton() *cerr.CustomError {
	placeholders := map[string]string{
		"{{ GO VERSION }}": helpers.GoVersion,
	}

	fmt.Printf("Stub: %s (metadata only)\n", hftx.Yellow("Skeleton"))
	if e := assets.ProcessEmbeddedAsset("skeleton/go.version", "go.version", placeholders); e != nil {
		return e
	}
	if e := assets.ProcessEmbeddedAsset("skeleton/dontexec.sh", "dontexec.sh", nil); e != nil {
		return e
	}
	_ = os.Chmod("dontexec.sh", os.FileMode(0755))

	return nil
}
