// stubber
// Written by J.F.Gratton <jean-francois@famillegratton.net>
// Original filename: src/createAssets/windows.go

package createAssets

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v5/terminalfx"
	"stubber/assets"
	"stubber/helpers"
)

// newGUID returns a random RFC 4122 version 4 UUID, formatted the way WiX
// expects (uppercase, unbraced, 8-4-4-4-12). crypto/rand rather than a
// dependency: this is the only thing stubber needs a UUID for.
func newGUID() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10

	s := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
	return strings.ToUpper(s), nil
}

// stubWindows renders __windows/Makefile and __windows/<softwarename>.wxs.
//
// Unlike every other distro, the .wxs carries two identifiers -- UpgradeCode
// and the main Component's Guid -- that MUST NOT change once a project has
// shipped a release under them (see the comment at the top of the rendered
// .wxs for why). So this is the one stub in stubber that is not
// unconditionally overwritten: the .wxs is rendered with freshly generated
// GUIDs only when it does not already exist. The Makefile carries no such
// state (it reads name/version/description from the manifest at build time,
// not from anything baked in here), so it is re-rendered every time, exactly
// like every other distro's Makefile.
func stubWindows(softwarename string) *cerr.CustomError {
	fmt.Printf("Stub: %s\n", hftx.Yellow("Windows"))

	makefilePlaceholders := map[string]string{
		"{{ SOFTWARE NAME }}": softwarename,
	}
	if err := assets.ProcessEmbeddedAsset(
		filepath.Join("windows", "Makefile"),
		filepath.Join("__windows", "Makefile"),
		makefilePlaceholders,
	); err != nil {
		return err
	}

	wxsPath := filepath.Join("__windows", softwarename+".wxs")
	if _, err := os.Stat(wxsPath); err == nil {
		fmt.Printf("File: %s already exists, leaving its UpgradeCode/Component Guid untouched\n", hftx.White(wxsPath))
		return nil
	} else if !os.IsNotExist(err) {
		return &cerr.CustomError{Title: "Unable to stat " + wxsPath, Message: err.Error()}
	}

	upgradeCode, err := newGUID()
	if err != nil {
		return &cerr.CustomError{Title: "Unable to generate UpgradeCode", Message: err.Error()}
	}
	componentGUID, err := newGUID()
	if err != nil {
		return &cerr.CustomError{Title: "Unable to generate Component Guid", Message: err.Error()}
	}

	wxsPlaceholders := map[string]string{
		"{{ SOFTWARE NAME }}":        softwarename,
		"{{ UPGRADE CODE }}":         upgradeCode,
		"{{ COMPONENT GUID }}":       componentGUID,
		"{{ INSTALL DIR PROPERTY }}": installDirProperty(softwarename, helpers.Target),
	}
	return assets.ProcessEmbeddedAsset(
		filepath.Join("windows", "product.wxs"),
		wxsPath,
		wxsPlaceholders,
	)
}

// installDirProperty returns the WiX <Property> line that pins INSTALLDIR to
// an absolute path, or an empty string when target is unset (leaving the
// Program Files default computed from the Directory table untouched). This is
// written directly into the .wxs source (unlike Manufacturer/Description,
// which flow through wixl's -D at build time), so target's path segments are
// XML-escaped here.
func installDirProperty(softwarename, target string) string {
	if target == "" {
		return ""
	}
	path := strings.TrimRight(strings.ReplaceAll(target, "/", `\`), `\`) + `\` + softwarename
	return fmt.Sprintf("    <Property Id='INSTALLDIR' Value='%s' />", escapeXMLAttr(path))
}

func escapeXMLAttr(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "'", "&apos;", `"`, "&quot;")
	return r.Replace(s)
}
