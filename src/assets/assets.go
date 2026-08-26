// stubber
// Written by J.F.Gratton <jean-francois@famillegratton.net>
// Original filename: src/assets/assets.go
// Original timestamp: 2026/05/03 08:05:21

package assets

import "embed"

// FS contains every template asset used by stubber.
//
// The all: prefix is intentional: some skeleton templates live under hidden
// directories such as .github/ and some filenames may begin with an underscore.
// A plain //go:embed skeleton would silently skip those entries.
//
// Asset names are relative to this directory, for example:
//
//	alpine/APKBUILD
//	debian/control
//	redhat/specfile
//	windows/product.wxs
//	skeleton/src/go.mod
//
// Each packaging directory is named after the distribution family it targets,
// matching both the __* directory it is rendered into and the stub function
// that renders it (alpine -> __alpine -> stubAlpine, and so on). The older
// packaging-format names (apk, deb, arch, rpm) were renamed for that symmetry.
// windows has no native package format at all, so there is no equivalent
// renaming history -- it targets a .msi from day one.
//

//go:embed all:alpine all:archlinux all:debian all:redhat all:skeleton all:windows
var FS embed.FS
