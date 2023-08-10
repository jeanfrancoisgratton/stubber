// stubber
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// creator.go, jfgratton : 2023-06-27

package executor

import (
	"os"
	"path/filepath"
	"stubber/helpers"
)

// Usage:
// stubber [-s stub rootdir] [-g "GO VERSION"] [-a] [-d] [-r] NAME

func CreateStub(softname string) error {
	var errcode error
	var currentdir string
	if helpers.BinaryName == "" {
		helpers.BinaryName = softname
	}
	// First, we need to switch directory to either currentdir, or whichever defined by the -p flag
	currentdir, err := os.Getwd()
	if err != nil {
		return nil
	}

	if helpers.RootDir == "." {
		helpers.RootDir = currentdir
	} else {
		if _, err := os.Stat(helpers.RootDir); os.IsNotExist(err) {
			if e := os.MkdirAll(helpers.RootDir, os.FileMode(0755)); e != nil {
				return e
			}
		}
	}
	if errcode = os.Chdir(helpers.RootDir); errcode != nil {
		return errcode
	}

	// Now we add the packaging stubs: APK, DEB, RPM
	if helpers.AlpineStub {
		if errcode = os.MkdirAll(filepath.Join(helpers.RootDir, "__alpine"), os.FileMode(0755)); errcode == nil {
			if errcode = stubAlpine(softname); errcode != nil {
				return errcode
			}
		}
	}

	if helpers.DebianStub {
		if errcode = stubDebian(); errcode != nil {
			return errcode
		}
	}
	if helpers.RedHatStub {
		if errcode = stubRedHat(); errcode != nil {
			return errcode
		}
	}

	return stubSkeleton()
}
