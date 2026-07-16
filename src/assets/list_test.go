// stubber
// Original name: src/assets/list_test.go

package assets

import "testing"

func TestEmbeddedAssetsPresent(t *testing.T) {
	// A representative file from every embedded stub set, including the renamed
	// skeleton templates (which must still be embedded via the all: prefix).
	for _, p := range []string{
		"apk/APKBUILD",
		"arch/PKGBUILD",
		"deb/control",
		"rpm/specfile",
		"skeleton/go.version",
		"skeleton/src/main.go.tmpl",
		"skeleton/src/cmd/root.go.tmpl",
		"skeleton/src/cmd/completion.go.tmpl",
	} {
		ok, err := Exists(p)
		if err != nil {
			t.Fatalf("Exists(%q): %v", p, err)
		}
		if !ok {
			t.Errorf("embedded asset %q is missing", p)
		}
	}
}

func TestListFilesNonEmpty(t *testing.T) {
	files, err := ListFiles(false)
	if err != nil {
		t.Fatalf("ListFiles: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("expected at least one embedded file")
	}
}
