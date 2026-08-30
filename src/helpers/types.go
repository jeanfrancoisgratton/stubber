// stubber
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/helpers/types.go
// Original time: 2023/06/28 08:20

package helpers

import "time"

// Command-line flags
var RootDir string
var AlpineStub, DebianStub, RedHatStub, SkeletonStub, ArchLinuxStub, WindowsStub bool
var GoVersion = "1.26.6"
var Arch string
var VersionNumber string
var ReleaseNumber string
var BinaryName = ""
var Description = ""
var Maintainer = "Jean-Francois Gratton <jean-francois@famillegratton.net>"
var Packager = "APK Builder <builder@famillegratton.net>"
var Section = ""
var Dependencies = ""
var Url = ""
var Manufacturer = "famillegratton.net"
var Quiet = false
var ReleaseDate = time.Now().Format("2006.01.02")

// CopyrightYear is the year stamped into debian/copyright's Copyright line.
// It is the year alone: ReleaseDate is too precise to read as a copyright
// assertion.
//
// Deliberately a function rather than one of the flag-backed vars above: it is
// always the current year. There is no flag to override it and it is not
// recorded in the manifest, so a refresh years later re-stamps the year it runs
// in instead of replaying a stale one.
func CopyrightYear() string { return time.Now().Format("2006") }

//var EnableGithubActions = false

// Stubs records which packaging/skeleton stubs a project owns.
type Stubs struct {
	Alpine    bool `json:"alpine"`
	Debian    bool `json:"debian"`
	RedHat    bool `json:"redhat"`
	ArchLinux bool `json:"archlinux"`
	Windows   bool `json:"windows"`
	Skeleton  bool `json:"skeleton"`
}

// Manifest is the full, lossless value set for a stubbed project.
type Manifest struct {
	SoftwareName  string `json:"softwarename"`
	BinaryName    string `json:"binaryname"`
	GoVersion     string `json:"goversion"`
	VersionNumber string `json:"versionnumber"`
	ReleaseNumber string `json:"releasenumber"`
	Description   string `json:"description"`
	Maintainer    string `json:"maintainer"`
	Packager      string `json:"packager"`
	Section       string `json:"section"`
	Dependencies  string `json:"dependencies"`
	Url           string `json:"url"`
	Manufacturer  string `json:"manufacturer"`
	Stubs         Stubs  `json:"stubs"`
}
