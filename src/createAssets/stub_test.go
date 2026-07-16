// stubber
// Original name: src/createAssets/stub_test.go

package createAssets

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"stubber/helpers"
)

// setBaseline resets the package-level flag globals to a known state before each
// test. CreateStub/RefreshStub read every value from these globals, so tests must
// drive them explicitly rather than rely on leftover state.
func setBaseline(root string) {
	helpers.Quiet = true
	helpers.RootDir = root
	helpers.BinaryName = ""
	helpers.GoVersion = "1.30.0"
	helpers.VersionNumber = "1.0.0"
	helpers.ReleaseNumber = "1"
	helpers.Description = "test tool"
	helpers.Maintainer = "Tester <t@example.com>"
	helpers.Packager = "Pkgr <p@example.com>"
	helpers.Section = "utils"
	helpers.Dependencies = "libc"
	helpers.Url = "https://example.com/repo"
	helpers.AlpineStub = false
	helpers.DebianStub = false
	helpers.RedHatStub = false
	helpers.ArchLinuxStub = false
	helpers.SkeletonStub = false
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	return string(data)
}

func assertContains(t *testing.T, path, want string) {
	t.Helper()
	if got := readFile(t, path); !strings.Contains(got, want) {
		t.Errorf("%s does not contain %q\n---\n%s\n---", path, want, got)
	}
}

func TestCreateStubWritesFilesAndManifest(t *testing.T) {
	root := filepath.Join(t.TempDir(), "demo")
	setBaseline(root)
	helpers.DebianStub = true
	helpers.ArchLinuxStub = true
	helpers.SkeletonStub = true

	if e := CreateStub("demo"); e != nil {
		t.Fatalf("CreateStub: %+v", e)
	}

	// Rendered packaging files carry the values we supplied.
	assertContains(t, filepath.Join(root, "__debian", "control"), "Version: 1.0.0-1")
	assertContains(t, filepath.Join(root, "__debian", "control"), "Depends: libc")
	assertContains(t, filepath.Join(root, "__archlinux", "PKGBUILD"), "pkgver=1.0.0")

	// Skeleton go.version holds the Go version, and templates are stripped of .tmpl.
	if got := strings.TrimSpace(readFile(t, filepath.Join(root, "go.version"))); got != "1.30.0" {
		t.Errorf("go.version = %q, want %q", got, "1.30.0")
	}
	if _, err := os.Stat(filepath.Join(root, "src", "main.go")); err != nil {
		t.Errorf("expected rendered src/main.go: %v", err)
	}

	// The manifest is written, named after the software, and records the stubs.
	m, e := helpers.FindManifest(root)
	if e != nil {
		t.Fatalf("FindManifest: %+v", e)
	}
	if m.SoftwareName != "demo" || m.BinaryName != "demo" || m.VersionNumber != "1.0.0" {
		t.Errorf("manifest values wrong: %+v", m)
	}
	if !m.Stubs.Debian || !m.Stubs.ArchLinux || !m.Stubs.Skeleton || m.Stubs.Alpine || m.Stubs.RedHat {
		t.Errorf("manifest stubs wrong: %+v", m.Stubs)
	}
}

func TestRefreshOnlyTouchesRequestedStubs(t *testing.T) {
	root := filepath.Join(t.TempDir(), "demo")
	setBaseline(root)
	helpers.DebianStub = true
	helpers.ArchLinuxStub = true

	if e := CreateStub("demo"); e != nil {
		t.Fatalf("CreateStub: %+v", e)
	}
	archBefore := readFile(t, filepath.Join(root, "__archlinux", "PKGBUILD"))

	// Refresh only Debian, bumping the version. (In the CLI the merge of stored +
	// changed values happens in the command; here we drive the merged globals.)
	setBaseline(root)
	helpers.VersionNumber = "2.0.0"
	helpers.DebianStub = true

	if e := RefreshStub("demo"); e != nil {
		t.Fatalf("RefreshStub: %+v", e)
	}

	// Debian was re-rendered with the new version, other values preserved.
	assertContains(t, filepath.Join(root, "__debian", "control"), "Version: 2.0.0-1")
	assertContains(t, filepath.Join(root, "__debian", "control"), "Depends: libc")

	// ArchLinux was not requested, so it must be byte-for-byte unchanged.
	if archAfter := readFile(t, filepath.Join(root, "__archlinux", "PKGBUILD")); archAfter != archBefore {
		t.Errorf("__archlinux/PKGBUILD changed despite not being refreshed")
	}
}

func TestRefreshSkeletonOnlyUpdatesGoVersion(t *testing.T) {
	root := filepath.Join(t.TempDir(), "demo")
	setBaseline(root)
	helpers.SkeletonStub = true

	if e := CreateStub("demo"); e != nil {
		t.Fatalf("CreateStub: %+v", e)
	}

	// Simulate real user code living where the skeleton template was rendered.
	mainPath := filepath.Join(root, "src", "main.go")
	const realCode = "package main // REAL USER CODE"
	if err := os.WriteFile(mainPath, []byte(realCode), 0644); err != nil {
		t.Fatal(err)
	}

	// Refresh the skeleton with a new Go version.
	setBaseline(root)
	helpers.SkeletonStub = true
	helpers.GoVersion = "2.5.5"

	if e := RefreshStub("demo"); e != nil {
		t.Fatalf("RefreshStub: %+v", e)
	}

	// go.version is updated ...
	if got := strings.TrimSpace(readFile(t, filepath.Join(root, "go.version"))); got != "2.5.5" {
		t.Errorf("go.version = %q, want %q", got, "2.5.5")
	}
	// ... but the user's real source is never clobbered.
	if got := readFile(t, mainPath); got != realCode {
		t.Errorf("src/main.go was overwritten by refresh: %q", got)
	}
}
