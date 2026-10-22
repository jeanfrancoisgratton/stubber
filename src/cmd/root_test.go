// stubber
// Original name: src/cmd/root_test.go

package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestSubcommandsRegistered(t *testing.T) {
	got := map[string]bool{}
	for _, c := range rootCmd.Commands() {
		got[c.Name()] = true
	}
	for _, want := range []string{"create", "refresh", "assets", "completion", "conn", "version"} {
		if !got[want] {
			t.Errorf("subcommand %q is not registered", want)
		}
	}
}

func TestCreateRequiredFlags(t *testing.T) {
	for _, name := range []string{"desc", "section", "url"} {
		f := createCmd.PersistentFlags().Lookup(name)
		if f == nil {
			t.Fatalf("create flag %q not defined", name)
		}
		req := f.Annotations[cobra.BashCompOneRequiredFlag]
		if len(req) == 0 || req[0] != "true" {
			t.Errorf("create flag %q is not marked required", name)
		}
	}
}

func TestRefreshFlagsAreOptional(t *testing.T) {
	// The same value flags must stay optional on refresh (values come from the manifest).
	for _, name := range []string{"desc", "section", "depends"} {
		f := refreshCmd.PersistentFlags().Lookup(name)
		if f == nil {
			t.Fatalf("refresh flag %q not defined", name)
		}
		if req := f.Annotations[cobra.BashCompOneRequiredFlag]; len(req) > 0 && req[0] == "true" {
			t.Errorf("refresh flag %q should not be required", name)
		}
	}
}

func TestCreateDefaultsSurviveRefreshRegistration(t *testing.T) {
	// create and refresh bind the same variables, and the last registered default wins:
	// both commands must declare the same defaults
	for _, name := range []string{"maintainer", "packager", "section", "manufacturer", "depends", "desc", "packagever", "packagerel"} {
		c, r := createCmd.PersistentFlags().Lookup(name), refreshCmd.PersistentFlags().Lookup(name)
		if c == nil || r == nil {
			t.Fatalf("flag %q not defined on both create and refresh", name)
		}
		if c.DefValue != r.DefValue {
			t.Errorf("flag %q: create default %q, refresh default %q", name, c.DefValue, r.DefValue)
		}
	}
	if f := createCmd.PersistentFlags().Lookup("maintainer"); f.DefValue == "" {
		t.Error("create flag \"maintainer\" lost its default")
	}
}
