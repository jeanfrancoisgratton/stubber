// stubber
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/helpers/structs.go
// Original time: 2023/06/28 08:20

package helpers

var StubRootDir string
var AlpineStub, DebianStub, RedHatStub bool
var GoVersion string
var Platform string
var VersionNumber, ReleaseNumber string
var BinaryName = ""
var Description = ""

type StubParamsStruct struct {
	StubRootDir, GoVersion, Platform               string
	AlpineStub, DebianStub, RedHatStub, UpdateOnly bool
}

var StubDefault = StubParamsStruct{
	StubRootDir: ".", GoVersion: "1.20.6", AlpineStub: true, DebianStub: true,
	RedHatStub: true, UpdateOnly: false, Platform: "amd64"}
