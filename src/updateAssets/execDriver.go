// stubber
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// execDriver.go, jfgratton : 2023-06-27

package updateAssets

import (
	"fmt"
	"os"
	"strings"
	"stubber/helpers"
)

func UpdateVersions(softwarename string) error {
	var errcode error
	var currentdir string

	// First, we need to switch directory to either currentdir, or whichever defined by the -p flag
	currentdir, err := os.Getwd()
	if err != nil {
		return nil
	}

	if errcode = os.Chdir(helpers.RootDir); errcode != nil {
		return errcode
	}

	//
	// Now we update the packaging stubs: APK, DEB and RPM
	//

	// Alpine ( -a )
	if helpers.AlpineStub {
		if errcode = updateAlpine(); errcode != nil {
			os.Chdir(currentdir)
			return errcode
		}
	}

	// Debian ( -d )
	if helpers.DebianStub {
		if errcode = updateDebian(); errcode != nil {
			os.Chdir(currentdir)
			return errcode
		}
	}

	// Debian ( -r )
	if helpers.RedHatStub {
		if errcode = updateRedHat(softwarename); errcode != nil {
			os.Chdir(currentdir)
			return errcode
		}
	}

	if helpers.SkeletonStub {
		if errcode = updateSkeleton(); errcode != nil {
			os.Chdir(currentdir)
			return errcode
		}
	}

	os.Chdir(currentdir)
	return nil
}

func replaceStrings(filein string, placeholders map[string]string) error {
	var err error
	var content []byte

	if content, err = os.ReadFile(filein); err != nil {
		return err
	}
	contentStr := string(content)

	for orgstr, newstr := range placeholders {
		contentStr = strings.ReplaceAll(contentStr, orgstr, newstr)
	}

	fmt.Printf("Modified content:\n%s\n", contentStr) // Debug print

	if err = os.WriteFile(filein, []byte(contentStr), 0644); err != nil {
		return err
	}
	return nil
}
