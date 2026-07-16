// stubber
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Original name: src/helpers/manifest.go

package helpers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
)

// ManifestFileName returns the per-project manifest filename for a given
// software: "<softname>.json". This file is written at `create` time and
// records every value used to render the stubs, so that `refresh` can keep the
// value of any flag that is NOT passed.
func ManifestFileName(softname string) string {
	return softname + ".json"
}

// FindManifest locates and reads the single stubber manifest ("<softname>.json")
// living directly in rootdir. `refresh` is meant to be run from the project root,
// so it discovers the software name from the manifest rather than taking it as an
// argument. A .json file only qualifies if it parses as a manifest whose
// SoftwareName matches its own filename; that keeps unrelated .json files
// (package.json, tsconfig.json, ...) from ever being mistaken for a manifest.
func FindManifest(rootdir string) (Manifest, *cerr.CustomError) {
	var zero Manifest

	entries, err := os.ReadDir(rootdir)
	if err != nil {
		return zero, &cerr.CustomError{Title: "Unable to read the project directory", Message: err.Error()}
	}

	var found []Manifest
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		data, rerr := os.ReadFile(filepath.Join(rootdir, e.Name()))
		if rerr != nil {
			continue
		}
		var m Manifest
		if json.Unmarshal(data, &m) != nil {
			continue
		}
		if m.SoftwareName != "" && e.Name() == ManifestFileName(m.SoftwareName) {
			found = append(found, m)
		}
	}

	switch len(found) {
	case 1:
		return found[0], nil
	case 0:
		return zero, &cerr.CustomError{
			Title:   "No manifest found",
			Message: "no stubber manifest (<softwarename>.json) in " + rootdir + "; run `refresh` from the project root, or `stubber create` first",
		}
	default:
		names := make([]string, 0, len(found))
		for _, m := range found {
			names = append(names, ManifestFileName(m.SoftwareName))
		}
		return zero, &cerr.CustomError{
			Title:   "Multiple manifests found",
			Message: "several stubber manifests in " + rootdir + " (" + strings.Join(names, ", ") + "); keep only one",
		}
	}
}

// SaveManifest writes m to the root of rootdir.
func SaveManifest(rootdir string, m Manifest) *cerr.CustomError {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return &cerr.CustomError{Title: "Unable to encode the manifest", Message: err.Error()}
	}

	path := filepath.Join(rootdir, ManifestFileName(m.SoftwareName))
	if err := os.WriteFile(path, append(data, '\n'), 0644); err != nil {
		return &cerr.CustomError{Title: "Unable to write the manifest", Message: err.Error()}
	}
	return nil
}
