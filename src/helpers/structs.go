// stubber
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/helpers/structs.go
// Original time: 2023/06/28 08:20

package helpers

import "time"

// Command-line flags
var RootDir string
var AlpineStub, DebianStub, RedHatStub, SkeletonStub bool
var GoVersion = "1.25.2"
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
