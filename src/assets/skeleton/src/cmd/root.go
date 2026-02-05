// {{ SOFTWARE NAME }}
// src/cmd/root.go

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "{{ SOFTWARE NAME }}",
	Short:   "Add a short description here",
	Version: "{{ PACKAGE VERSION }}-{{ PACKAGE RELEASE }} ({{ RELEASE DATE }}), Go version = " + runtime.Version(),
	Long: `This tools allows you to create a software directory structure.
This follows my template and allows you to package your software with minimal effort once built`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
