// stubber : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/root.go

package cmd

import (
	"fmt"
	"os"
	"stubber/createAssets"
	"stubber/helpers"
	"stubber/updateAssets"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "stubber",
	Short:   "Creates your GOLANG software directory structure",
	Version: "1.92.00 (2025.12.21)",
	Long: `This tools allows you to create a software directory structure.
This follows my template and allows you with minimal effort to package your software once built`,
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
		if err := updateAssets.UpdateVersions(args[0]); err != nil {
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
	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(completionCmd, createCmd, updateCmd)
	rootCmd.PersistentFlags().BoolVarP(&helpers.Quiet, "quiet", "q", false, "Silence non-essential output.")
	rootCmd.PersistentFlags().StringVarP(&helpers.RootDir, "projectrootdir", "p", ".", "Project root directory.")
	rootCmd.PersistentFlags().StringVarP(&helpers.BinaryName, "binaryname", "b", "", "Output binary name.")
	rootCmd.PersistentFlags().StringVarP(&helpers.GoVersion, "gover", "g", "1.25.5", "Where to put the skeleton dir.")

	createCmd.PersistentFlags().StringVarP(&helpers.VersionNumber, "packagever", "V", "", "Package version number.")
	createCmd.PersistentFlags().StringVarP(&helpers.ReleaseNumber, "packagerel", "R", "", "Package release number.")
	createCmd.PersistentFlags().StringVarP(&helpers.Description, "desc", "D", "", "Package description.")
	createCmd.PersistentFlags().BoolVarP(&helpers.AlpineStub, "alpine", "a", false, "Create an Alpine packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.DebianStub, "debian", "d", false, "Create a Debian packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.RedHatStub, "redhat", "r", false, "Create a RedHat packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.SkeletonStub, "skeleton", "k", false, "Create the skeleton stub in the project root directory.")
	//createCmd.PersistentFlags().BoolVarP(&helpers.EnableGithubActions, "gha", "w", false, "Copy gha files from .github/workflows/ .")
	createCmd.PersistentFlags().StringVarP(&helpers.Maintainer, "maintainer", "M", "", "Software maintainer.")
	createCmd.PersistentFlags().StringVarP(&helpers.Packager, "packager", "P", "", "Software packager.")
	createCmd.PersistentFlags().StringVarP(&helpers.Section, "section", "s", "Packaging tool", "Debian package section.")
	createCmd.PersistentFlags().StringVarP(&helpers.Dependencies, "depends", "e", "", "Package dependencies.")
	createCmd.PersistentFlags().StringVarP(&helpers.Url, "url", "u", "https://git.famillegratton.net:3000/ADD_URL_HERE", "Git repo URL.")

	updateCmd.PersistentFlags().StringVarP(&helpers.VersionNumber, "packagever", "V", "1.00.00", "Package version number.")
	updateCmd.PersistentFlags().StringVarP(&helpers.ReleaseNumber, "packagerel", "R", "0", "Package release number.")
	updateCmd.PersistentFlags().StringVarP(&helpers.Description, "desc", "D", "", "Package description.")
	updateCmd.PersistentFlags().StringVarP(&helpers.Arch, "arch", "A", "amd64", "Arch (architecture).")
	updateCmd.PersistentFlags().BoolVarP(&helpers.AlpineStub, "alpine", "a", false, "Create an Alpine packaging stub.")
	updateCmd.PersistentFlags().BoolVarP(&helpers.DebianStub, "debian", "d", false, "Create a Debian packaging stub.")
	updateCmd.PersistentFlags().BoolVarP(&helpers.RedHatStub, "redhat", "r", false, "Create a RedHat packaging stub.")
	updateCmd.PersistentFlags().BoolVarP(&helpers.SkeletonStub, "skeleton", "k", false, "Create the skeleton stub in the project root directory.")
	updateCmd.PersistentFlags().StringVarP(&helpers.Maintainer, "maintainer", "M", "", "Software maintainer.")
	updateCmd.PersistentFlags().StringVarP(&helpers.Packager, "packager", "P", "", "Software packager.")
	updateCmd.PersistentFlags().StringVarP(&helpers.Section, "section", "s", "Packaging tool", "Debian package section.")
	updateCmd.PersistentFlags().StringVarP(&helpers.Dependencies, "depends", "e", "", "Package dependencies.")
	updateCmd.PersistentFlags().StringVarP(&helpers.Url, "url", "u", "", "Github repo URL.")
}
