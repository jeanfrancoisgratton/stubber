// stubber
// Written by J.F.Gratton <jean-francois@famillegratton.net>
// Original filename: src/assets/list.go
// Original timestamp: 2026/05/03 09:42:07

package assets

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"sort"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hf "github.com/jeanfrancoisgratton/helperFunctions/v5"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v5/terminalfx"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Entry describes one item embedded in FS.
type Entry struct {
	Path  string
	IsDir bool
	Size  int64
}

// List returns every embedded asset entry, including directories.
//
// Paths are relative to the embedded filesystem root, for example:
//
//	alpine/APKBUILD
//	archlinux/PKGBUILD
//	skeleton/src/go.mod
//
// The root "." entry is intentionally skipped.
func List(displayOutput bool) ([]Entry, *cerr.CustomError) {
	var entries []Entry

	err := fs.WalkDir(FS, ".", func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if path == "." {
			return nil
		}

		entry := Entry{
			Path:  path,
			IsDir: d.IsDir(),
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			entry.Size = info.Size()
		}

		entries = append(entries, entry)
		return nil
	})

	if err != nil {
		return nil, &cerr.CustomError{Title: "Error listing files", Message: err.Error()}
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Path < entries[j].Path
	})

	if displayOutput {
		showEntries(entries)
	}
	return entries, nil
}

// ListFiles returns only embedded regular files.
//
// Directories are excluded.
func ListFiles(displayOutput bool) ([]string, error) {
	entries, err := List(displayOutput)
	if err != nil {
		return nil, err
	}

	files := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir {
			files = append(files, entry.Path)
		}
	}

	return files, nil
}

// ListDirs returns only embedded directories.
//
// The root "." directory is excluded.
func ListDirs(displayOutput bool) ([]string, error) {
	entries, err := List(displayOutput)
	if err != nil {
		return nil, err
	}

	dirs := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir {
			dirs = append(dirs, entry.Path)
		}
	}

	return dirs, nil
}

// Exists reports whether path exists in the embedded filesystem.
//
// This checks the embedded virtual filesystem, not the host filesystem.
func Exists(path string) (bool, error) {
	_, err := fs.Stat(FS, path)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	return false, err
}

func showEntries(entries []Entry) {

	if len(entries) == 0 {
		fmt.Println(hftx.WarningSign("No entries found"))
		return
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Path", "Directory ?", "Size"})

	for _, entry := range entries {
		isd := hftx.ErrorSign("")
		if entry.IsDir {
			isd = hftx.EnabledSign("")
		}
		t.AppendRow(table.Row{entry.Path, isd, hf.SI(entry.Size)})
	}

	t.SortBy([]table.SortBy{
		{Name: "Path", Mode: table.Asc},
	})
	t.SetStyle(table.StyleColoredYellowWhiteOnBlack)
	t.Style().Format.Header = text.FormatDefault
	t.Render()
}
