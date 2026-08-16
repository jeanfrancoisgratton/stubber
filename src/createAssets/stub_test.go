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

// dontexec.sh manages the _dontexec build-veto marker. It is the one skeleton
// file that is stubber-owned tooling rather than user content, so it is emitted
// by create *and* re-emitted by refresh, and it has to land executable -- a
// non-executable veto helper is useless.
func TestDontexecIsRenderedExecutableByCreateAndRefresh(t *testing.T) {
	assertExecutable := func(path string) {
		t.Helper()
		fi, err := os.Stat(path)
		if err != nil {
			t.Fatalf("expected %s: %v", path, err)
		}
		if fi.Mode().Perm()&0111 == 0 {
			t.Errorf("%s is not executable (mode %v)", path, fi.Mode().Perm())
		}
	}

	root := filepath.Join(t.TempDir(), "demo")
	setBaseline(root)
	helpers.SkeletonStub = true

	if e := CreateStub("demo"); e != nil {
		t.Fatalf("CreateStub: %+v", e)
	}
	script := filepath.Join(root, "dontexec.sh")
	assertExecutable(script)

	// It carries no project-specific value, so nothing should have been
	// substituted into it.
	if got := readFile(t, script); strings.Contains(got, "{{") {
		t.Errorf("dontexec.sh contains an unrendered placeholder:\n%s", got)
	}

	// Refresh is how an already-scaffolded project picks it up: delete it and
	// confirm a skeleton refresh puts it back, executable.
	if err := os.Remove(script); err != nil {
		t.Fatal(err)
	}
	setBaseline(root)
	helpers.SkeletonStub = true

	if e := RefreshStub("demo"); e != nil {
		t.Fatalf("RefreshStub: %+v", e)
	}
	assertExecutable(script)
}

// The --depends flag is free-form and Debian-shaped, but pacman needs a bash
// array of individually quoted elements, so stubArchLinux rewrites it.
func TestPacmanDepends(t *testing.T) {
	cases := []struct {
		name, in, want string
	}{
		{"empty stays empty", "", ""},
		{"single", "libc", "'libc'"},
		{"comma separated", "libc, bash-completion", "'libc' 'bash-completion'"},
		{"space separated", "libc bash-completion", "'libc' 'bash-completion'"},
		{"ragged separators", " libc ,,  bash-completion ,", "'libc' 'bash-completion'"},
		{"version constraint survives", "glibc>=2.38", "'glibc>=2.38'"},
		// A single quote cannot be escaped inside a bash single-quoted string.
		{"embedded quote", "we'ird", `'we'\''ird'`},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := pacmanDepends(tc.in); got != tc.want {
				t.Errorf("pacmanDepends(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

// The URL and dependency placeholders were both missing from stubArchLinux's
// replacement map, so PKGBUILD rendered url="{{ URL }}" literally and silently
// dropped whatever --depends was given.
func TestArchLinuxRendersUrlAndDepends(t *testing.T) {
	root := filepath.Join(t.TempDir(), "demo")
	setBaseline(root)
	helpers.ArchLinuxStub = true

	if e := CreateStub("demo"); e != nil {
		t.Fatalf("CreateStub: %+v", e)
	}

	pkgbuild := filepath.Join(root, "__archlinux", "PKGBUILD")
	assertContains(t, pkgbuild, `url="https://example.com/repo"`)
	assertContains(t, pkgbuild, `depends=('libc')`)

	if got := readFile(t, pkgbuild); strings.Contains(got, "{{") {
		t.Errorf("PKGBUILD contains an unrendered placeholder:\n%s", got)
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
