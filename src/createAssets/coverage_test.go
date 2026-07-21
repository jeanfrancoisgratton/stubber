// stubber
// Original name: src/createAssets/coverage_test.go

package createAssets

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"stubber/assets"
	"stubber/helpers"
)

// TestEveryEmbeddedPackagingAssetIsRendered is a drift guard: every file
// embedded under a packaging asset directory (apk/, deb/, arch/, rpm/) must be
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
			name: "alpine", assetDir: "apk", outDir: "__alpine",
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
			name: "debian", assetDir: "deb", outDir: "__debian",
			enable:   func() { helpers.DebianStub = true },
			rendered: func(binary, software, base string) string { return base },
		},
		{
			name: "archlinux", assetDir: "arch", outDir: "__archlinux",
			enable:   func() { helpers.ArchLinuxStub = true },
			rendered: func(binary, software, base string) string { return base },
		},
		{
			name: "redhat", assetDir: "rpm", outDir: "__redhat",
			enable: func() { helpers.RedHatStub = true },
			rendered: func(binary, software, base string) string {
				if base == "specfile" {
					return software + ".spec"
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
