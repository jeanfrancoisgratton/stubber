// stubber : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/root.go

package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"stubber/createAssets"
	"stubber/helpers"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "stubber",
	Short:   "Creates your GOLANG software directory structure",
	Version: "2.7.0 (2026.08.16), Go version : v" + strings.TrimPrefix(runtime.Version(), "go"),
	Long: `This tools allows you to create a software directory structure.
This follows my template and allows you with minimal effort to package your software once built`,
}

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Creates the directory structure (skeleton) for the new software",
	Example: "software_name",
	Run: func(cmd *cobra.Command, args []string) {
		if !helpers.ArchLinuxStub && !helpers.AlpineStub && !helpers.DebianStub && !helpers.RedHatStub && !helpers.SkeletonStub {
			fmt.Println("You need to enable at least one of the following: -A (archlinux), -a (alpine), -d (debian), -r (redhat) or -k (skeleton)")
			os.Exit(1)
		}
		if len(args) != 1 {
			fmt.Println("Usage: stubber create [-A|-a|-d|-r|-k] $SOFTWARENAME")
			os.Exit(2)
		}
		if err := createAssets.CreateStub(args[0]); err != nil {

			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Updates an existing stub, keeping unset values and updating the ones you pass",
	Args:  cobra.NoArgs,
	Long: `Re-renders the stubs of an existing project.

Run this from the project root, i.e. the directory that holds the
<softwarename>.json manifest; stubber discovers the software name from it, so no
positional argument is needed. Any flag you pass overrides its stored value,
while every flag you omit is preserved as-is.

Only the packaging stubs you explicitly request (-A -a -d -r -k) are refreshed;
passing none refreshes nothing. -k (skeleton) refreshes go.version only, so your
source code and documentation are never overwritten.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !helpers.ArchLinuxStub && !helpers.AlpineStub && !helpers.DebianStub && !helpers.RedHatStub && !helpers.SkeletonStub {
			fmt.Println("Nothing to refresh: pass at least one of -A (archlinux), -a (alpine), -d (debian), -r (redhat) or -k (skeleton)")
			os.Exit(0)
		}

		// The project root (holding <softwarename>.json) is the current dir, or -p
		root := helpers.RootDir
		if root == "." {
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			root = cwd
		}

		m, err := helpers.FindManifest(root)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Override stored values only for the flags actually passed on the command line
		if cmd.Flags().Changed("gover") {
			m.GoVersion = helpers.GoVersion
		}
		if cmd.Flags().Changed("binaryname") {
			m.BinaryName = helpers.BinaryName
		}
		if cmd.Flags().Changed("packagever") {
			m.VersionNumber = helpers.VersionNumber
		}
		if cmd.Flags().Changed("packagerel") {
			m.ReleaseNumber = helpers.ReleaseNumber
		}
		if cmd.Flags().Changed("desc") {
			m.Description = helpers.Description
		}
		if cmd.Flags().Changed("maintainer") {
			m.Maintainer = helpers.Maintainer
		}
		if cmd.Flags().Changed("packager") {
			m.Packager = helpers.Packager
		}
		if cmd.Flags().Changed("section") {
			m.Section = helpers.Section
		}
		if cmd.Flags().Changed("depends") {
			m.Dependencies = helpers.Dependencies
		}

		// Push the merged values back into the globals used by the renderers
		helpers.RootDir = root
		helpers.BinaryName = m.BinaryName
		helpers.GoVersion = m.GoVersion
		helpers.VersionNumber = m.VersionNumber
		helpers.ReleaseNumber = m.ReleaseNumber
		helpers.Description = m.Description
		helpers.Maintainer = m.Maintainer
		helpers.Packager = m.Packager
		helpers.Section = m.Section
		helpers.Dependencies = m.Dependencies
		helpers.Url = m.Url

		// A refreshed stub type now exists; never clear the ones we did not touch
		m.Stubs.Alpine = m.Stubs.Alpine || helpers.AlpineStub
		m.Stubs.Debian = m.Stubs.Debian || helpers.DebianStub
		m.Stubs.RedHat = m.Stubs.RedHat || helpers.RedHatStub
		m.Stubs.ArchLinux = m.Stubs.ArchLinux || helpers.ArchLinuxStub
		m.Stubs.Skeleton = m.Stubs.Skeleton || helpers.SkeletonStub

		if err := createAssets.RefreshStub(m.SoftwareName); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := helpers.SaveManifest(root, m); err != nil {
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
	rootCmd.AddCommand(completionCmd, createCmd, refreshCmd, assetsCmd)
	rootCmd.PersistentFlags().BoolVarP(&helpers.Quiet, "quiet", "q", false, "Silence non-essential output.")
	rootCmd.PersistentFlags().StringVarP(&helpers.RootDir, "projectrootdir", "p", ".", "Project root directory.")
	rootCmd.PersistentFlags().StringVarP(&helpers.BinaryName, "binaryname", "b", "", "Output binary name.")
	rootCmd.PersistentFlags().StringVarP(&helpers.GoVersion, "gover", "g", "1.26.4", "Where to put the skeleton dir.")

	createCmd.PersistentFlags().StringVarP(&helpers.VersionNumber, "packagever", "V", "", "Package version number.")
	createCmd.PersistentFlags().StringVarP(&helpers.ReleaseNumber, "packagerel", "R", "", "Package release number.")
	createCmd.PersistentFlags().StringVarP(&helpers.Description, "desc", "D", "", "Package description.")
	createCmd.PersistentFlags().BoolVarP(&helpers.ArchLinuxStub, "archlinux", "A", false, "Create an Archlinux packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.AlpineStub, "alpine", "a", false, "Create an Alpine packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.DebianStub, "debian", "d", false, "Create a Debian packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.RedHatStub, "redhat", "r", false, "Create a RedHat packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.SkeletonStub, "skeleton", "k", false, "Create the skeleton stub in the project root directory.")
	createCmd.PersistentFlags().StringVarP(&helpers.Maintainer, "maintainer", "M", "", "Software maintainer.")
	createCmd.PersistentFlags().StringVarP(&helpers.Packager, "packager", "P", "", "Software packager.")
	createCmd.PersistentFlags().StringVarP(&helpers.Section, "section", "s", "Packaging tool", "Debian package section.")
	createCmd.PersistentFlags().StringVarP(&helpers.Dependencies, "depends", "e", "", "Package dependencies.")
	createCmd.PersistentFlags().StringVarP(&helpers.Url, "url", "u", "https://git.famillegratton.net:3000/ADD_URL_HERE", "Git repo URL.")

	// These must be supplied explicitly on `create` (they seed the manifest)
	_ = createCmd.MarkPersistentFlagRequired("desc")
	_ = createCmd.MarkPersistentFlagRequired("section")
	_ = createCmd.MarkPersistentFlagRequired("depends")

	// refresh reuses the same value flags as create (minus -u); unpassed flags keep their stored value
	refreshCmd.PersistentFlags().StringVarP(&helpers.VersionNumber, "packagever", "V", "", "Package version number.")
	refreshCmd.PersistentFlags().StringVarP(&helpers.ReleaseNumber, "packagerel", "R", "", "Package release number.")
	refreshCmd.PersistentFlags().StringVarP(&helpers.Description, "desc", "D", "", "Package description.")
	refreshCmd.PersistentFlags().BoolVarP(&helpers.ArchLinuxStub, "archlinux", "A", false, "Refresh the Archlinux packaging stub.")
	refreshCmd.PersistentFlags().BoolVarP(&helpers.AlpineStub, "alpine", "a", false, "Refresh the Alpine packaging stub.")
	refreshCmd.PersistentFlags().BoolVarP(&helpers.DebianStub, "debian", "d", false, "Refresh the Debian packaging stub.")
	refreshCmd.PersistentFlags().BoolVarP(&helpers.RedHatStub, "redhat", "r", false, "Refresh the RedHat packaging stub.")
	refreshCmd.PersistentFlags().BoolVarP(&helpers.SkeletonStub, "skeleton", "k", false, "Refresh the skeleton stub (go.version only).")
	refreshCmd.PersistentFlags().StringVarP(&helpers.Maintainer, "maintainer", "M", "", "Software maintainer.")
	refreshCmd.PersistentFlags().StringVarP(&helpers.Packager, "packager", "P", "", "Software packager.")
	refreshCmd.PersistentFlags().StringVarP(&helpers.Section, "section", "s", "Packaging tool", "Debian package section.")
	refreshCmd.PersistentFlags().StringVarP(&helpers.Dependencies, "depends", "e", "", "Package dependencies.")
}
