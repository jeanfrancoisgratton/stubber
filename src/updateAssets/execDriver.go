// stubber
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// execDriver.go, jfgratton : 2023-06-27

package updateAssets

import (
	"bufio"
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
		if errcode = updateDebian(softwarename); errcode != nil {
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

func replaceStrings(filePath string, placeholders map[string]string) error {
	tempFilePath := filePath + ".tmp"

	inputFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		replaced := false

		for key, value := range placeholders {
			if strings.HasPrefix(line, key) {
				line = value
				replaced = true
				break
			}
		}

		if replaced {
			_, err := outputFile.WriteString(line + "\n")
			if err != nil {
				return err
			}
		} else {
			_, err := outputFile.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := inputFile.Close(); err != nil {
		return err
	}
	if err := outputFile.Close(); err != nil {
		return err
	}

	if err := os.Rename(tempFilePath, filePath); err != nil {
		return err
	}

	return nil
}
