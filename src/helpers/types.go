// stubber
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/helpers/types.go
// Original time: 2023/06/28 08:20

package helpers

import "time"

// Command-line flags
var RootDir string
var AlpineStub, DebianStub, RedHatStub, SkeletonStub, ArchLinuxStub bool
var GoVersion = "1.26.2"
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
var Quiet = false
var ReleaseDate = time.Now().Format("2006.01.02")

//var EnableGithubActions = false

// Stubs records which packaging/skeleton stubs a project owns.
type Stubs struct {
	Alpine    bool `json:"alpine"`
	Debian    bool `json:"debian"`
	RedHat    bool `json:"redhat"`
	ArchLinux bool `json:"archlinux"`
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
	Stubs         Stubs  `json:"stubs"`
}
