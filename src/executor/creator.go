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

// checkWhitespace() : quick and dirty way to ensure that a filename is properly enclosed by quotes
//func checkWhitespace(s string) bool {
//	for _, c := range s {
//		if c == ' ' || c == '\t' {
//			return true
//		}
//	}
//	return false
//}

func CreateStub(softname string) error {
	var errcode error
	var currentdir string

	//// Badly needed for filenames
	//if checkWhitespace(softname) {
	//	softname = "'" + softname + "'"
	//}

	if helpers.BinaryName == "" {
		helpers.BinaryName = softname
	}
	// First, we need to switch directory to either currentdir, or whichever defined by the -p flag
	currentdir, err := os.Getwd()
	if err != nil {
		return nil
	}

	// Second, We create the project root dir if -p is provided
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

	//
	// Now we add the packaging stubs: APK, DEB, RPM, and Skeleton
	//

	// Alpine ( -a )
	if helpers.AlpineStub {
		if errcode = os.MkdirAll(filepath.Join(helpers.RootDir, "__alpine"), os.FileMode(0755)); errcode == nil {
			if errcode = stubAlpine(softname); errcode != nil {
				return errcode
			}
		}
	}

	// Debian ( -d )
	if helpers.DebianStub {
		if errcode = os.MkdirAll(filepath.Join(helpers.RootDir, "__debian"), os.FileMode(0755)); errcode == nil {
			if errcode = stubDebian(softname); errcode != nil {
				return errcode
			}
		}
	}

	// Debian ( -r )
	if helpers.RedHatStub {
		if errcode = stubRedHat(softname); errcode != nil {
			return errcode
		}
	}

	return stubSkeleton()
}
