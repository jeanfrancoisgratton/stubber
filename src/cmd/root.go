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
	Version: "1.81.02 (2025.10.08)",
	Long: `This tools allows you to create a software directory structure.
This follows my template and allows you with minimal effort to package your software once built`,
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		changelog()
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
	rootCmd.AddCommand(completionCmd, clCmd, createCmd, updateCmd)
	rootCmd.PersistentFlags().BoolVarP(&helpers.Quiet, "quiet", "q", false, "Silence non-essential output.")
	rootCmd.PersistentFlags().StringVarP(&helpers.RootDir, "projectrootdir", "p", ".", "Project root directory.")
	rootCmd.PersistentFlags().StringVarP(&helpers.BinaryName, "binaryname", "b", "", "Output binary name.")
	rootCmd.PersistentFlags().StringVarP(&helpers.GoVersion, "gover", "g", "1.25.3", "Where to put the skeleton dir.")

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
	createCmd.PersistentFlags().StringVarP(&helpers.Url, "url", "u", "https://url-not-set/", "Github repo URL.")

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

func changelog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Print(`
VERSION		DATE			COMMENT
-------		----			-------
1.81.02		2025.10.09		GO version bump, upgraded customError and helperFunctions, enabled shell completion
1.80.02		2025.09.23		Renamed the imports checker script
1.80.00		2025.09.12		GO version bump, added a new cyclic imports check script in src/
1.78.01		2025.07.09		fixed typo in CHANGELOG.md var placeholder. GO version bump
1.78.00		2025.07.01		disable CGO when building the tool and assets 
1.77.00		2025.06.14		added a new template variable to deal with version and changelog
1.76.00		2025.06.14		updated GO default version in templates, fixed broken asset generator in build.sh
1.75.01		2025.05.09		fixed rpm-builddeps.sh, broken for a long time
1.75.00		2025.03.29		build.sh script update
1.74.01		2025.03.15		GO version bump, url fix in APK packaging
1.73.00		2024.11.15		removed unneeded files, GO version bump
1.72.02		2024.10.19		modified buildDeps install process for go-bindata
1.71.00		2024.08.27		binary, assets and templates GO version bump; last version before integrating git support
1.70.00		2024.08.11		specfile now arch-independent, GO version bump
1.65.01		2024.08.05		added a flag (-w) to add github actions
1.62.00		2024.07.28		same as 1.61.01, updated assets
1.61.01		2024.05.25		GO version bump, updated updateBuildDeps.sh
1.60.00		2024.04.05		GO version bump, rewrite of build.sh
1.55.01		2024.03.15		APKBUILD now respects the -u parameter
1.55.00		2024.02.16		more packaging fixes
1.54.01		2024.02.15		fixes in APK and RPM packaging scripts
1.54.00		2024.02.14		GO version bump, arm64 is no longer supported, more post/pre install/upgrade/remove scripts
1.53.01		2024.01.11		GO version bump, release number bump
1.53.00		2024.01.09		misc fixes in assets generation, removed/renamed some files
1.52.02		2023.12.31		assets fixes
1.52.01		2023.12.29		output fix, added missing go.version file in assets/
1.52.00		2023.11.08		build.sh changes
1.51.00		2023.10.19		version scheme refactor, misc minor changes
1.505		2023.08.18		removed unresolved function call in cmd/root.go, doc update
1.500		2023.08.16		FEATURE: update asset command, to bump version of software
1.205		2023.08.13		fixed missing placeholder in RPM stub
1.201		2023.08.13		asset generation was silently broken in RPM/DEB/APK building
1.100		2023.08.12		added Debian packaging script, added missing placeholders, etc
1.010		2023.08.12		re-instated -V and -R flags, added CHANGELOG.md in assets, removed "IN THIS BRANCH" and src/helpers/
1.000		2023.08.11		final version
0.500		2023.08.11		completed apk, deb, rpm, skeleton
0.100		2023.06.25		stub
`)
}
