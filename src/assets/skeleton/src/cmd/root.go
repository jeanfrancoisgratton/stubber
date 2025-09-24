// {{ SOFTWARE NAME }}
// src/cmd/root.go

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "{{ SOFTWARE NAME }}",
	Short:   "Add a short description here",
	Version: "{{ PACKAGE VERSION }}-{{ PACKAGE RELEASE }} ({{ RELEASE DATE }})",
	Long: `This tools allows you to create a software directory structure.
This follows my template and allows you to package your software with minimal effort once built`,
}

// Shows changelog
var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows the Changelog",
	Run: func(cmd *cobra.Command, args []string) {
		changeLog()
	},
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

	rootCmd.AddCommand(clCmd)
}

func changeLog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Println("CHANGELOG")
	fmt.Println("=========")
	fmt.Println()

	fmt.Print(`
VERSION			DATE			COMMENT
-------			----			-------
{{ PACKAGE VERSION }}		{{ RELEASE DATE }}		Initial release
`)
}
