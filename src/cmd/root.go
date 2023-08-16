// stubber : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/root.go

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"stubber/createAssets"
	"stubber/helpers"
)

var version = "1.500-0 (2023.08.16)"

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
		if !helpers.AlpineStub && !helpers.DebianStub && !helpers.RedHatStub && !helpers.SkeletonStub {
			fmt.Println("You need to enable at least one of the following: -a (alpine), -d (debian), -r (redhat) or -k (skeleton)")
			os.Exit(1)
		}
		if len(args) != 1 {
			fmt.Println("Usage: stubber create [-a|-d|-r|-k] $SOFTWARENAME")
			os.Exit(2)
		}
		if err := createAssets.CreateStub(args[0]); err != nil {
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
		if !helpers.AlpineStub && !helpers.DebianStub && !helpers.RedHatStub && !helpers.SkeletonStub {
			fmt.Println("You need to enable at least one of the following: -a (alpine), -d (debian), -r (redhat) or -k (skeleton)")
			os.Exit(1)
		}
		if len(args) != 1 {
			fmt.Println("Usage: stubber update [-a|-d|-r|-k] $SOFTWARENAME")
			os.Exit(2)
		}
		if err := createAssets.CreateStub(args[0]); err != nil {
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
	rootCmd.AddCommand(updateCmd)
	rootCmd.PersistentFlags().BoolVarP(&helpers.Quiet, "quiet", "q", false, "Silence non-essential output.")
	rootCmd.PersistentFlags().StringVarP(&helpers.RootDir, "projectrootdir", "p", ".", "Project root directory.")
	rootCmd.PersistentFlags().StringVarP(&helpers.BinaryName, "binaryname", "b", "", "Output binary name.")
	rootCmd.PersistentFlags().StringVarP(&helpers.VersionNumber, "packagever", "V", "", "Package version number.")
	rootCmd.PersistentFlags().StringVarP(&helpers.ReleaseNumber, "packagerel", "R", "", "Package release number.")
	rootCmd.PersistentFlags().StringVarP(&helpers.Description, "desc", "D", "", "Package description.")
	rootCmd.PersistentFlags().StringVarP(&helpers.GoVersion, "gover", "g", "1.21.0", "Where to put the skeleton dir.")
	rootCmd.PersistentFlags().StringVarP(&helpers.Arch, "arch", "A", "amd64", "Arch (architecture).")
	rootCmd.PersistentFlags().BoolVarP(&helpers.AlpineStub, "alpine", "a", false, "Create an Alpine packaging stub.")
	rootCmd.PersistentFlags().BoolVarP(&helpers.DebianStub, "debian", "d", false, "Create a Debian packaging stub.")
	rootCmd.PersistentFlags().BoolVarP(&helpers.RedHatStub, "redhat", "r", false, "Create a RedHat packaging stub.")
	rootCmd.PersistentFlags().BoolVarP(&helpers.SkeletonStub, "skeleton", "k", false, "Create the skeleton stub in the project root directory.")
	rootCmd.PersistentFlags().StringVarP(&helpers.Maintainer, "maintainer", "M", "", "Software maintainer.")
	rootCmd.PersistentFlags().StringVarP(&helpers.Packager, "packager", "P", "", "Software packager.")
	rootCmd.PersistentFlags().StringVarP(&helpers.Section, "section", "s", "", "Debian package section.")
	rootCmd.PersistentFlags().StringVarP(&helpers.Dependencies, "depends", "e", "", "Package dependencies.")
	rootCmd.PersistentFlags().StringVarP(&helpers.Url, "url", "u", "", "Github repo URL.")
}
