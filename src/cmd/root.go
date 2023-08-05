// stubber : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/root.go

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"stubber/executor"
	"stubber/helpers"
)

// Usage:
// stubber [-s stub rootdir] [-g "GO VERSION"] [-a] [-d] [-r] NAME VERSION RELEASE

var version = "0.100-0 (2023.06.25)"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "stubber",
	Short:   "Creates your GOLANG software directory structure",
	Version: version,
	Long: `This tools allows you to a software directory structure.
This follows my template and allows you with minimal effort to package your software once built`,
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		helpers.Changelog()
	},
}

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Creates the directory structure (skeleton) for the new software",
	Example: "software_name",
	Run: func(cmd *cobra.Command, args []string) {
		if !helpers.AlpineStub && !helpers.DebianStub && !helpers.RedHatStub {
			fmt.Println("You need at least to disable one of the following: -a (alpine), -d (debian), -r (redhat)")
			os.Exit(1)
		}
		if len(args) < 3 {
			fmt.Println("Usage: stubber create [-a|-d|-r] $SOFTWARENAME $VERSIONNUMBER $RELEASENUMBER")
			os.Exit(2)
		}
		if err := executor.CreateStub(args[0], args[1], args[2]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm", "del"},
	Short:   "Deletes the directory structure (skeleton) for the new software",
	Run: func(cmd *cobra.Command, args []string) {
		if err := executor.DeleteStub(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"up"},
	Short:   "Updates the build scripts with new Version and Release numbers",
	Run: func(cmd *cobra.Command, args []string) {
		if err := executor.Update(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(clCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(updateCmd)
	createCmd.PersistentFlags().StringVarP(&helpers.StubRootDir, "stubdir", "s", ".", "Where to put the skeleton dir.")
	createCmd.PersistentFlags().StringVarP(&helpers.BinaryName, "binaryname", "b", "", "Outout binary name.")
	createCmd.PersistentFlags().StringVarP(&helpers.Description, "desc", "d", "", "Package description.")
	createCmd.PersistentFlags().StringVarP(&helpers.GoVersion, "gover", "g", "1.20.6", "Where to put the skeleton dir.")
	createCmd.PersistentFlags().StringVarP(&helpers.Platform, "platform", "p", "amd64", "Platform (architecture).")
	createCmd.PersistentFlags().BoolVarP(&helpers.AlpineStub, "alpine", "a", true, "Create an Alpine packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.DebianStub, "debian", "d", true, "Create a Debian packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.RedHatStub, "redhat", "r", true, "Create a RedHat packaging stub.")
	updateCmd.PersistentFlags().StringVarP(&helpers.VersionNumber, "versionnumber", "V", "", "Version number to use.")
	updateCmd.PersistentFlags().StringVarP(&helpers.ReleaseNumber, "releasenumber", "R", "", "Release number to use.")
	updateCmd.PersistentFlags().StringVarP(&helpers.GoVersion, "gover", "g", "1.20.6", "Where to put the skeleton dir.")
	updateCmd.PersistentFlags().StringVarP(&helpers.Platform, "platform", "p", "amd64", "Platform (architecture).")
	updateCmd.PersistentFlags().StringVarP(&helpers.StubRootDir, "stubdir", "s", ".", "Where to put the skeleton dir.")
	updateCmd.PersistentFlags().StringVarP(&helpers.Description, "desc", "d", "", "Package description.")
	updateCmd.PersistentFlags().StringVarP(&helpers.BinaryName, "binaryname", "b", "", "Outout binary name.")
}
