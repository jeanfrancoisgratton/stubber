// stubber
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// execDriver.go, jfgratton : 2023-06-27

package createAssets

import (
	"fmt"
	"os"
	"path/filepath"
	"stubber/helpers"
)

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

	// Second, We create the project root dir if -p is provided
	if helpers.RootDir == "." {
		helpers.RootDir = filepath.Join(currentdir, softname)
		if err := os.MkdirAll(softname, os.FileMode(0755)); err != nil {
			return nil
		} else {
			os.Chdir(filepath.Join(currentdir, softname))
		}
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

	fmt.Printf("Creating stub for software %s in %s\n", helpers.Green(softname), helpers.Green(helpers.RootDir))

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
				os.Chdir(currentdir)
				return errcode
			}
		}
	}

	// Debian ( -r )
	if helpers.RedHatStub {
		if errcode = stubRedHat(softname); errcode != nil {
			os.Chdir(currentdir)
			return errcode
		}
	}

	// Skeleton ( -
	if helpers.SkeletonStub {
		if errcode = os.MkdirAll(filepath.Join(helpers.RootDir, "src", "cmd"), os.FileMode(0755)); errcode != nil {
			os.Chdir(currentdir)
			return errcode
		}
		if errcode := stubSkeleton(softname); errcode != nil {
			os.Chdir(currentdir)
			return errcode
		}
	}
	os.Chdir(currentdir)
	return nil
}
