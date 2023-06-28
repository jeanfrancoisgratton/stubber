// stubber
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// creator.go, jfgratton : 2023-06-27

package executor

import (
	"os"
	"stubber/helpers"
)

// Usage:
// stubber [-s stub rootdir] [-g "GO VERSION"] [-a] [-d] [-r] NAME VERSION RELEASE

func CreateStub(softname string, version string, release string) error {
	var errcode error
	var currentdir string

	// First, we need to switch directory to either currentdir, or whichever defined by the -s flag
	currentdir, err := os.Getwd()
	if err != nil {
		return nil
	}

	if helpers.StubRootDir == "." {
		helpers.StubRootDir = currentdir
	}
	if errcode = os.Chdir(helpers.StubRootDir); errcode != nil {
		return nil
	}

	// Now we add the packaging stubs: APK, DEB, RPM
	if helpers.AlpineStub {
		if errcode = stubAlpine(); errcode != nil {
			return errcode
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
