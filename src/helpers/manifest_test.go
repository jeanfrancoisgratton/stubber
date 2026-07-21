// stubber
// Original name: src/helpers/manifest_test.go

package helpers

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestManifestFileName(t *testing.T) {
	if got := ManifestFileName("myapp"); got != "myapp.json" {
		t.Errorf("ManifestFileName(\"myapp\") = %q, want %q", got, "myapp.json")
	}
}

func TestSaveAndFindManifest(t *testing.T) {
	dir := t.TempDir()

	want := Manifest{
		SoftwareName:  "myapp",
		BinaryName:    "myapp",
		VersionNumber: "1.2.3",
		ReleaseNumber: "4",
		Section:       "utils",
		Dependencies:  "libc",
		Stubs:         Stubs{Debian: true, ArchLinux: true},
	}

	if e := SaveManifest(dir, want); e != nil {
		t.Fatalf("SaveManifest: %+v", e)
	}

	// The file must be named after the software.
	if _, err := os.Stat(filepath.Join(dir, "myapp.json")); err != nil {
		t.Fatalf("expected myapp.json to exist: %v", err)
	}

	got, e := FindManifest(dir)
	if e != nil {
		t.Fatalf("FindManifest: %+v", e)
	}
	if got.SoftwareName != want.SoftwareName || got.VersionNumber != want.VersionNumber ||
		got.ReleaseNumber != want.ReleaseNumber || got.Section != want.Section ||
		got.Dependencies != want.Dependencies {
		t.Errorf("round-trip mismatch:\n got  %+v\n want %+v", got, want)
	}
	if !got.Stubs.Debian || !got.Stubs.ArchLinux || got.Stubs.Alpine {
		t.Errorf("stubs not preserved: %+v", got.Stubs)
	}
}

// TestSaveManifestDoesNotHTMLEscape guards the human-readability of the manifest:
// angle brackets in fields like the maintainer must stay literal so that a
// `refresh` never rewrites a clean manifest into < / > escapes.
func TestSaveManifestDoesNotHTMLEscape(t *testing.T) {
	dir := t.TempDir()
	m := Manifest{
		SoftwareName: "myapp",
		Maintainer:   "J.F.Gratton <jean-francois@example.net>",
	}
	if e := SaveManifest(dir, m); e != nil {
		t.Fatalf("SaveManifest: %+v", e)
	}

	data, err := os.ReadFile(filepath.Join(dir, "myapp.json"))
	if err != nil {
		t.Fatal(err)
	}
	// HTML-escaped output would render the brackets as unicode escapes and so
	// would NOT contain the literal "<...>" substring.
	if !strings.Contains(string(data), "<jean-francois@example.net>") {
		t.Errorf("manifest did not keep literal angle brackets (HTML-escaped?):\n%s", data)
	}
}

func TestFindManifestNone(t *testing.T) {
	dir := t.TempDir()
	if _, e := FindManifest(dir); e == nil {
		t.Fatal("expected an error when no manifest is present")
	}
}

func TestFindManifestIgnoresUnrelatedJSON(t *testing.T) {
	dir := t.TempDir()

	// A plain JSON file that is not a manifest.
	if err := os.WriteFile(filepath.Join(dir, "package.json"), []byte(`{"name":"x"}`), 0644); err != nil {
		t.Fatal(err)
	}
	// A manifest-looking file whose name does NOT match its softwarename.
	if err := os.WriteFile(filepath.Join(dir, "other.json"), []byte(`{"softwarename":"different"}`), 0644); err != nil {
		t.Fatal(err)
	}

	if _, e := FindManifest(dir); e == nil {
		t.Fatal("expected no manifest to be recognised among unrelated JSON files")
	}

	// Now drop in a real one; it must be the one that is found.
	if e := SaveManifest(dir, Manifest{SoftwareName: "real"}); e != nil {
		t.Fatalf("SaveManifest: %+v", e)
	}
	got, e := FindManifest(dir)
	if e != nil {
		t.Fatalf("FindManifest: %+v", e)
	}
	if got.SoftwareName != "real" {
		t.Errorf("found %q, want %q", got.SoftwareName, "real")
	}
}

func TestFindManifestMultiple(t *testing.T) {
	dir := t.TempDir()
	if e := SaveManifest(dir, Manifest{SoftwareName: "alpha"}); e != nil {
		t.Fatalf("SaveManifest: %+v", e)
	}
	if e := SaveManifest(dir, Manifest{SoftwareName: "beta"}); e != nil {
		t.Fatalf("SaveManifest: %+v", e)
	}
	if _, e := FindManifest(dir); e == nil {
		t.Fatal("expected an error when several manifests are present")
	}
}
