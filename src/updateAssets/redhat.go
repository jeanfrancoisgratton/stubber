package updateAssets

import (
	"fmt"
	hf "github.com/jeanfrancoisgratton/helperFunctions/v3/terminalfx"
	"strings"
	"stubber/helpers"
)

func updateRedHat(softwarename string) error {
	var err error
	goarch := helpers.Arch

	// Debian uses amd64, not x86_64
	arch := strings.ToLower(helpers.Arch)
	if arch == "amd64" {
		arch = "x86_64"
		goarch = "amd64"
	}

	placeholders := map[string]string{
		// rpmbuild-deps.sh
		"grep ^BuildRequires \"": "grep ^BuildRequires \"" + softwarename + ".spec\" |awk -F\\: '{print \"sudo dnf install -y\"$2}'|sed -e 's/,/ /g' | sh",
		"sudo wget":              "sudo wget -q https://go.dev/dl/go" + helpers.GoVersion + ".linux-" + goarch + ".tar.gz -O /tmp/go.tar.gz",
		// specfile
		"%define _version":    "%define _version " + helpers.VersionNumber,
		"%define _rel":        "%define _rel  " + helpers.ReleaseNumber,
		"%define _arch":       "%define _arch " + arch,
		"%define _binaryname": "%define _binaryname " + helpers.BinaryName,
	}
	paths := []string{"rpmbuild-deps.sh", softwarename + ".spec"}

	fmt.Printf("Stub: %s\n", hf.Yellow("RedHat"))
	for _, pathloop := range paths {
		if err = replaceStrings(pathloop, placeholders); err != nil {
			return err
		}
	}
	return nil
}
