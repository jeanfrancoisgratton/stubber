// stubber
// Original name: src/createAssets/coverage_test.go

package createAssets

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"stubber/assets"
	"stubber/helpers"
)

// TestEveryEmbeddedSkeletonAssetIsRendered is the same drift guard for the
// skeleton, which had none: ROADMAP.md and TODO.md sat embedded but absent from
// skeleton.go's paths slice, so they were never written into a scaffold and
// nothing noticed.
//
// It walks the whole subtree rather than one directory, because skeleton assets
// nest (src/, src/cmd/), and mirrors stubSkeleton's two renaming rules.
func TestEveryEmbeddedSkeletonAssetIsRendered(t *testing.T) {
	root := filepath.Join(t.TempDir(), "demo")
	setBaseline(root)
	helpers.SkeletonStub = true

	if e := CreateStub("demo"); e != nil {
		t.Fatalf("CreateStub: %+v", e)
	}

	err := fs.WalkDir(assets.FS, "skeleton", func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil || d.IsDir() {
			return walkErr
		}

		// Paths are relative to skeleton/, and stubSkeleton renames as it goes:
		// gitignore -> .gitignore, *.tmpl loses the suffix, and
		// ISSUES/ROADMAP/TODO/CHANGELOG.md move under docs/.
		rel := strings.TrimPrefix(path, "skeleton/")
		switch {
		case rel == "gitignore":
			rel = ".gitignore"
		case strings.HasSuffix(rel, ".tmpl"):
			rel = strings.TrimSuffix(rel, ".tmpl")
		}
		switch rel {
		case "ISSUES.md", "ROADMAP.md", "TODO.md", "CHANGELOG.md":
			rel = "docs/" + rel
		}

		if _, err := os.Stat(filepath.Join(root, rel)); err != nil {
			t.Errorf("embedded asset %s was not rendered (expected %s): %v", path, rel, err)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking embedded skeleton: %v", err)
	}
}

// TestEveryEmbeddedPackagingAssetIsRendered is a drift guard: every file
// embedded under a packaging asset directory (alpine/, debian/, archlinux/, redhat/) must be
// produced by CreateStub. The per-stub `paths` slices are maintained by hand, so
// adding, renaming or removing an asset without updating them would otherwise go
// unnoticed until a user's scaffold silently misses a file (e.g. a new Makefile).
func TestEveryEmbeddedPackagingAssetIsRendered(t *testing.T) {
	// rendered maps an embedded asset basename to the filename create writes,
	// mirroring each stub's own renaming rules.
	cases := []struct {
		name     string
		assetDir string
		outDir   string
		enable   func()
		rendered func(binary, software, base string) string
	}{
		{
			name: "alpine", assetDir: "alpine", outDir: "__alpine",
			enable: func() { helpers.AlpineStub = true },
			rendered: func(binary, software, base string) string {
				// Everything except APKBUILD and the Makefile is an install
				// script and is prefixed with the binary name.
				if base == "APKBUILD" || base == "Makefile" {
					return base
				}
				return binary + "." + base
			},
		},
		{
			name: "debian", assetDir: "debian", outDir: "__debian",
			enable:   func() { helpers.DebianStub = true },
			rendered: func(binary, software, base string) string { return base },
		},
		{
			name: "archlinux", assetDir: "archlinux", outDir: "__archlinux",
			enable:   func() { helpers.ArchLinuxStub = true },
			rendered: func(binary, software, base string) string { return base },
		},
		{
			name: "redhat", assetDir: "redhat", outDir: "__redhat",
			enable: func() { helpers.RedHatStub = true },
			rendered: func(binary, software, base string) string {
				if base == "specfile" {
					return software + ".spec"
				}
				return base
			},
		},
		{
			name: "windows", assetDir: "windows", outDir: "__windows",
			enable: func() { helpers.WindowsStub = true },
			rendered: func(binary, software, base string) string {
				if base == "product.wxs" {
					return software + ".wxs"
				}
				return base
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := filepath.Join(t.TempDir(), "demo")
			setBaseline(root)
			tc.enable()

			if e := CreateStub("demo"); e != nil {
				t.Fatalf("CreateStub: %+v", e)
			}

			embedded, err := fs.ReadDir(assets.FS, tc.assetDir)
			if err != nil {
				t.Fatalf("reading embedded %s: %v", tc.assetDir, err)
			}

			for _, de := range embedded {
				if de.IsDir() {
					continue
				}
				want := tc.rendered("demo", "demo", de.Name())
				if _, err := os.Stat(filepath.Join(root, tc.outDir, want)); err != nil {
					t.Errorf("embedded asset %s/%s was not rendered (expected %s/%s): %v",
						tc.assetDir, de.Name(), tc.outDir, want, err)
				}
			}
		})
	}
}
