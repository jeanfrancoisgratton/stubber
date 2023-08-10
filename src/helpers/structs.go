// stubber
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/helpers/structs.go
// Original time: 2023/06/28 08:20

package helpers

var RootDir string
var AlpineStub, DebianStub, RedHatStub, SkeletonStub bool
var GoVersion string
var Arch string
var VersionNumber, ReleaseNumber string
var BinaryName = ""
var Description = ""
var Maintainer = "Jean-Francois Gratton <jean-francois@famillegratton.net>"
var Packager = "APK Builder <builder@famillegratton.net>"

type StubParamsStruct struct {
	RootDir, GoVersion, Platform                   string
	AlpineStub, DebianStub, RedHatStub, UpdateOnly bool
}

var StubDefault = StubParamsStruct{
	RootDir: ".", GoVersion: "1.21.0", AlpineStub: true, DebianStub: true,
	RedHatStub: true, UpdateOnly: false, Platform: "amd64"}
