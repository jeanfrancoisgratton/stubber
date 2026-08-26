// stubber
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// execDriver.go, jfgratton : 2023-06-27

package createAssets

import (
	"fmt"
	"os"
	"path/filepath"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v5/terminalfx"
	"stubber/helpers"
)

func CreateStub(softname string) *cerr.CustomError {
	//var err error
	var currentdir string

	if helpers.BinaryName == "" {
		helpers.BinaryName = softname
	}
	// First, we need to switch directory to either currentdir, or whichever defined by the -p flag
	currentdir, err := os.Getwd()
	if err != nil {
		return &cerr.CustomError{Title: "Unable to getcwd", Message: err.Error()}
	}

	// Second, We create the project root dir if -p is provided
	if helpers.RootDir == "." {
		helpers.RootDir = filepath.Join(currentdir, softname)
		if err = os.MkdirAll(softname, os.FileMode(0755)); err != nil {
			return &cerr.CustomError{Title: "Unable to mkdir", Message: err.Error()}
		} else {
			os.Chdir(filepath.Join(currentdir, softname))
		}
	} else {
		if _, err = os.Stat(helpers.RootDir); os.IsNotExist(err) {
			if e := os.MkdirAll(helpers.RootDir, os.FileMode(0755)); e != nil {
				return &cerr.CustomError{Title: "Unable to mkdir", Message: e.Error()}
			}
		}
	}
	if err = os.Chdir(helpers.RootDir); err != nil {
		return &cerr.CustomError{Title: "Unable to mkdir", Message: err.Error()}
	}

	fmt.Printf("Creating stub for software %s in %s\n", hftx.Green(softname), hftx.Green(helpers.RootDir))

	// Alpine ( -a )
	if helpers.AlpineStub {
		if err = os.MkdirAll(filepath.Join(helpers.RootDir, "__alpine"), os.FileMode(0755)); err == nil {
			if err := stubAlpine(softname); err != nil {
				//return &cerr.CustomError{Title: "Unable to create the __alpine build dir", Message: err.Error()}
				return err
			}
		}
	}

	// Debian ( -d )
	if helpers.DebianStub {
		if err = os.MkdirAll(filepath.Join(helpers.RootDir, "__debian"), os.FileMode(0755)); err == nil {
			if err := stubDebian(softname); err != nil {
				os.Chdir(currentdir)
				return err
			}
		}
	}

	// RedHat ( -r )
	if helpers.RedHatStub {
		if err = os.MkdirAll(filepath.Join(helpers.RootDir, "__redhat"), os.FileMode(0755)); err == nil {
			if err := stubRedHat(softname); err != nil {
				os.Chdir(currentdir)
				return err
			}
		}
	}

	// ArchLinux -A
	if helpers.ArchLinuxStub {
		if err = os.MkdirAll(filepath.Join(helpers.RootDir, "__archlinux"), os.FileMode(0755)); err == nil {
			if err := stubArchLinux(softname); err != nil {
				os.Chdir(currentdir)
				return err
			}
		}
	}

	// Windows ( -w )
	if helpers.WindowsStub {
		if err = os.MkdirAll(filepath.Join(helpers.RootDir, "__windows"), os.FileMode(0755)); err == nil {
			if err := stubWindows(softname); err != nil {
				os.Chdir(currentdir)
				return err
			}
		}
	}

	// Skeleton ( -
	if helpers.SkeletonStub {
		if err = os.MkdirAll(filepath.Join(helpers.RootDir, "src", "cmd"), os.FileMode(0755)); err != nil {
			os.Chdir(currentdir)
			return &cerr.CustomError{Title: "Unable to create the skeleton structure", Message: err.Error()}
		}

		if errcode := stubSkeleton(softname); errcode != nil {
			os.Chdir(currentdir)
			return errcode
		}
	}

	// Record every value so that `refresh` can preserve the flags not passed to it
	if err := helpers.SaveManifest(helpers.RootDir, manifestFromFlags(softname)); err != nil {
		os.Chdir(currentdir)
		return err
	}

	os.Chdir(currentdir)
	return nil
}

// manifestFromFlags snapshots the current flag/global state into a Manifest.
func manifestFromFlags(softname string) helpers.Manifest {
	return helpers.Manifest{
		SoftwareName:  softname,
		BinaryName:    helpers.BinaryName,
		GoVersion:     helpers.GoVersion,
		VersionNumber: helpers.VersionNumber,
		ReleaseNumber: helpers.ReleaseNumber,
		Description:   helpers.Description,
		Maintainer:    helpers.Maintainer,
		Packager:      helpers.Packager,
		Section:       helpers.Section,
		Dependencies:  helpers.Dependencies,
		Url:           helpers.Url,
		Manufacturer:  helpers.Manufacturer,
		Target:        helpers.Target,
		Stubs: helpers.Stubs{
			Alpine:    helpers.AlpineStub,
			Debian:    helpers.DebianStub,
			RedHat:    helpers.RedHatStub,
			ArchLinux: helpers.ArchLinuxStub,
			Windows:   helpers.WindowsStub,
			Skeleton:  helpers.SkeletonStub,
		},
	}
}
