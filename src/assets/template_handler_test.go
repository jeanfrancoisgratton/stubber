// stubber
// Original name: src/assets/template_handler_test.go

package assets

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"stubber/helpers"
)

func TestProcessEmbeddedAsset(t *testing.T) {
	helpers.Quiet = true

	// go.version is the simplest template: its entire content is one placeholder.
	out := filepath.Join(t.TempDir(), "nested", "go.version")
	if e := ProcessEmbeddedAsset("skeleton/go.version", out, map[string]string{
		"{{ GO VERSION }}": "9.9.9",
	}); e != nil {
		t.Fatalf("ProcessEmbeddedAsset: %+v", e)
	}

	// The output directory did not exist beforehand; it must have been created.
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("reading output: %v", err)
	}
	if got := strings.TrimSpace(string(data)); got != "9.9.9" {
		t.Errorf("substituted content = %q, want %q", got, "9.9.9")
	}
}

func TestProcessEmbeddedAssetMissingInput(t *testing.T) {
	helpers.Quiet = true
	out := filepath.Join(t.TempDir(), "x")
	if e := ProcessEmbeddedAsset("does/not/exist", out, nil); e == nil {
		t.Fatal("expected an error for a missing embedded input")
	}
}
