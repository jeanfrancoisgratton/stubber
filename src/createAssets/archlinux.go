// stubber
// Written by J.F.Gratton <jean-francois@famillegratton.net>
// Original filename: src/createAssets/archlinux.go
// Original timestamp: 2026/05/02 20:23:35

package createAssets

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v5/terminalfx"
	"stubber/assets"
	"stubber/helpers"
)

// pacmanDepends renders helpers.Dependencies as the body of a PKGBUILD bash
// array. The flag is free-form and Debian-shaped ("libc, bash-completion"),
// which pacman cannot consume: it wants each element quoted separately, so
// "libc, bash-completion" becomes 'libc' 'bash-completion'.
//
// An empty list yields an empty string, leaving depends=() as it was.
func pacmanDepends(deps string) string {
	fields := strings.FieldsFunc(deps, func(r rune) bool {
		return r == ',' || unicode.IsSpace(r)
	})

	quoted := make([]string, 0, len(fields))
	for _, f := range fields {
		// A single quote cannot be escaped inside a bash single-quoted string;
		// close, insert an escaped quote, reopen.
		quoted = append(quoted, "'"+strings.ReplaceAll(f, "'", `'\''`)+"'")
	}

	return strings.Join(quoted, " ")
}

func stubArchLinux(softwarename string) *cerr.CustomError {
	placeholders := map[string]string{
		"{{ SOFTWARE NAME }}":   softwarename,
		"{{ PACKAGE VERSION }}": helpers.VersionNumber,
		"{{ PACKAGE RELEASE }}": helpers.ReleaseNumber,
		"{{ DESCRIPTION }}":     helpers.Description,
		"{{ BINARY NAME }}":     helpers.BinaryName,
		"{{ URL }}":             helpers.Url,
		"{{ DEPENDENCIES }}":    pacmanDepends(helpers.Dependencies),
	}

	fmt.Printf("Stub: %s\n", hftx.Yellow("ArchLinux"))
	paths := []string{"1.install-build-deps.sh", "2.build-package.sh", "Makefile", "PKGBUILD"}

	for _, path := range paths {

		if err := assets.ProcessEmbeddedAsset(filepath.Join("archlinux", path), filepath.Join("__archlinux", path), placeholders); err != nil {
			return err
		}
		os.Chmod(filepath.Join("__archlinux", path), os.FileMode(0755))
	}
	return nil
}
