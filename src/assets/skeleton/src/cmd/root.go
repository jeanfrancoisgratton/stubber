// {{ SOFTWARE NAME }}
// src/cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "{{ SOFTWARE NAME }}",
	Short:   "Add a short description here",
	Version: "{{ PACKAGE VERSION }}-{{ PACKAGE RELEASE }} (2024.xx.yy)",
	Long: `This tools allows you to a software directory structure.
This follows my template and allows you with minimal effort to package your software once built`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
