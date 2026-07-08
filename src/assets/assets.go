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
//	apk/APKBUILD
//	deb/control
//	rpm/specfile
//	skeleton/src/go.mod
//
// That preserves the old go-bindata names produced by:
//
//	go-bindata -prefix ../assets ../assets/...
//

//go:embed all:apk all:arch all:deb all:rpm all:skeleton
var FS embed.FS
