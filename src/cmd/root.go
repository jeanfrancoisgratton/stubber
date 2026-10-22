// stubber : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/cmd/root.go

package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"stubber/git"

	"stubber/createAssets"
	"stubber/helpers"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v5/terminalfx"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stubber",
	Short: "Creates your Go software directory structure",
	Long: `stubber creates the directory structure of a new Go software: a buildable skeleton, and
packaging stubs for Alpine, Debian, RedHat, Arch Linux and Windows.

It can also create the software's git repository on its forge (Forgejo, GitHub or Bitbucket)
and push the scaffolded code to it; see "stubber create --help" and "stubber conn --help".`,
	Example: `  stubber create -k -d -r -D "My tool" -s utils -u https://git.example.net:3000/me/mytool mytool
  stubber refresh -d -V 1.1.0 -p ./mytool
  stubber conn add myforge`,
}

// buildVersion and buildDate are injected at link time via
// `-ldflags "-X stubber/cmd.buildVersion=... -X stubber/cmd.buildDate=..."`
// by every build path (see version-bump.sh and each __*/{Makefile,APKBUILD,
// PKGBUILD,spec}); a plain `go build` leaves them at these fallbacks.
var buildVersion = "dev"
var buildDate = "unknown"

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Shows the software version",
	Example: "  stubber version",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(hftx.White("stubber v" + buildVersion + " (" + buildDate + "), Go version = v" + strings.TrimPrefix(runtime.Version(), "go") + " (" + runtime.GOARCH + ")"))
	},
}

var createCmd = &cobra.Command{
	Use:   "create [flags] SOFTWARENAME",
	Short: "Creates the directory structure (skeleton and packaging stubs) for a new software",
	Long: `Creates the directory structure for SOFTWARENAME, in ./SOFTWARENAME (or in the directory given by -p).

At least one stub flag is needed: -A (archlinux), -a (alpine), -d (debian), -r (redhat), -w (windows)
or -k (skeleton). -D, -s and -u are mandatory. Every value used is recorded in SOFTWARENAME.json,
at the project root, so that "stubber refresh" can preserve it later.

With -G, stubber also creates the git repository given by -u (https://SERVER[:PORT]/OWNER/REPO) on its
forge, then commits the scaffolded code and pushes it on the develop branch. The repository is always
created private; changing its visibility is left to you. The credentials are those of the connection
whose server hosts the URL (see "stubber conn --help"); when several connections point to that server,
choose one with -c CONNECTION.`,
	Example: `  # Go skeleton only
  stubber create -k -D "My awesome CLI tool" -s utils -u https://git.example.net:3000/me/mytool mytool

  # Skeleton and every packaging stub, in a custom directory
  stubber create -A -a -d -r -w -k -V 1.0.0 -R 1 -D "My tool" -s utils -e "ca-certificates" \
    -u https://git.example.net:3000/me/mytool -p ~/dev/mytool mytool

  # Same, and create the (private) repository on GitHub, then push the code on develop
  stubber create -k -d -r -D "My tool" -s utils -u https://github.com/me/mytool -G mytool

  # Create the repository in the "myorg" organization, with the "work" connection
  stubber create -k -D "My tool" -s utils -u https://git.example.net:3000/myorg/mytool -G -c work mytool`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !helpers.ArchLinuxStub && !helpers.AlpineStub && !helpers.DebianStub && !helpers.RedHatStub && !helpers.WindowsStub && !helpers.SkeletonStub {
			fmt.Println("You need to enable at least one of the following: -A (archlinux), -a (alpine), -d (debian), -r (redhat), -w (windows) or -k (skeleton)")
			os.Exit(1)
		}
		if cmd.Flags().Changed("connection") && !helpers.CreateGitRepo {
			fmt.Println("-c (--connection) is only used with -G (--git)")
			os.Exit(1)
		}
		// With -G, everything that can be checked without side effects is checked before scaffolding anything
		var repoPlan *git.RepoPlan
		if helpers.CreateGitRepo {
			var e *cerr.CustomError
			if repoPlan, e = git.PrepareRepo(args[0]); e != nil {
				fmt.Println(e)
				fmt.Println("Nothing was created")
				os.Exit(1)
			}
		}
		if err := createAssets.CreateStub(args[0]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if repoPlan != nil {
			if e := repoPlan.Create(); e != nil {
				fmt.Println(e)
				os.Exit(1)
			}
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

Only the packaging stubs you explicitly request (-A -a -d -r -w -k) are refreshed;
passing none refreshes nothing. -k (skeleton) refreshes go.version only, so your
source code and documentation are never overwritten. -w (windows) refreshes the
Makefile but never re-renders an existing .wxs, so its UpgradeCode/Component
Guid are never disturbed by a refresh.

-u (the URL) cannot be changed by refresh; edit SOFTWARENAME.json if needed.`,
	Example: `  # From the project root: bump the version and refresh the Debian and RedHat stubs
  stubber refresh -d -r -V 1.1.0 -R 1

  # Change the Debian dependencies only, pointing -p at the project root
  stubber refresh -d -e "docker.io, ca-certificates" -p ./mytool

  # Move to a new Go release, without touching any source or packaging file
  stubber refresh -k -g 1.27.3`,
	Run: func(cmd *cobra.Command, args []string) {
		if !helpers.ArchLinuxStub && !helpers.AlpineStub && !helpers.DebianStub && !helpers.RedHatStub && !helpers.WindowsStub && !helpers.SkeletonStub {
			fmt.Println("Nothing to refresh: pass at least one of -A (archlinux), -a (alpine), -d (debian), -r (redhat), -w (windows) or -k (skeleton)")
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
			m.DefaultReleaseNumber = helpers.ReleaseNumber
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
		if cmd.Flags().Changed("manufacturer") {
			m.Manufacturer = helpers.Manufacturer
		}

		// Push the merged values back into the globals used by the renderers
		helpers.RootDir = root
		helpers.BinaryName = m.BinaryName
		helpers.GoVersion = m.GoVersion
		helpers.VersionNumber = m.VersionNumber
		helpers.ReleaseNumber = m.DefaultReleaseNumber
		helpers.Description = m.Description
		helpers.Maintainer = m.Maintainer
		helpers.Packager = m.Packager
		helpers.Section = m.Section
		helpers.Dependencies = m.Dependencies
		helpers.Url = m.Url
		helpers.Manufacturer = m.Manufacturer

		// A refreshed stub type now exists; never clear the ones we did not touch
		m.Stubs.Alpine = m.Stubs.Alpine || helpers.AlpineStub
		m.Stubs.Debian = m.Stubs.Debian || helpers.DebianStub
		m.Stubs.RedHat = m.Stubs.RedHat || helpers.RedHatStub
		m.Stubs.ArchLinux = m.Stubs.ArchLinux || helpers.ArchLinuxStub
		m.Stubs.Windows = m.Stubs.Windows || helpers.WindowsStub
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
	rootCmd.SilenceUsage = true

	rootCmd.AddCommand(completionCmd, createCmd, refreshCmd, assetsCmd, versionCmd, conn_cfgCmd)
	rootCmd.PersistentFlags().BoolVarP(&helpers.Quiet, "quiet", "q", false, "Silence non-essential output.")
	rootCmd.PersistentFlags().StringVarP(&helpers.RootDir, "projectrootdir", "p", ".", "Project root directory (create: ./SOFTWARENAME if left to \".\"; refresh: the current directory).")
	rootCmd.PersistentFlags().StringVarP(&helpers.BinaryName, "binaryname", "b", "", "Output binary name (default: the software name).")
	rootCmd.PersistentFlags().StringVarP(&helpers.GoVersion, "gover", "g", helpers.GoVersion, "Go version to build against.")

	createCmd.PersistentFlags().StringVarP(&helpers.VersionNumber, "packagever", "V", "", "Package version number.")
	createCmd.PersistentFlags().StringVarP(&helpers.ReleaseNumber, "packagerel", "R", "", "Package release number.")
	createCmd.PersistentFlags().StringVarP(&helpers.Description, "desc", "D", "", "Package description (mandatory).")
	createCmd.PersistentFlags().BoolVarP(&helpers.ArchLinuxStub, "archlinux", "A", false, "Create an Archlinux packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.AlpineStub, "alpine", "a", false, "Create an Alpine packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.DebianStub, "debian", "d", false, "Create a Debian packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.RedHatStub, "redhat", "r", false, "Create a RedHat packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.WindowsStub, "windows", "w", false, "Create a Windows packaging stub.")
	createCmd.PersistentFlags().BoolVarP(&helpers.SkeletonStub, "skeleton", "k", false, "Create the skeleton stub in the project root directory.")
	createCmd.PersistentFlags().StringVarP(&helpers.Maintainer, "maintainer", "M", helpers.Maintainer, "Software maintainer (\"Name <email>\").")
	createCmd.PersistentFlags().StringVarP(&helpers.Packager, "packager", "P", helpers.Packager, "Software packager (\"Name <email>\").")
	createCmd.PersistentFlags().StringVarP(&helpers.Section, "section", "s", "Packaging tool", "Debian package section / RPM group (mandatory).")
	createCmd.PersistentFlags().StringVarP(&helpers.Dependencies, "depends", "e", "", "Package dependencies, comma-separated.")
	createCmd.PersistentFlags().StringVarP(&helpers.Url, "url", "u", "", "Git repository URL, https://SERVER[:PORT]/OWNER/REPO (mandatory).")
	createCmd.PersistentFlags().StringVarP(&helpers.Manufacturer, "manufacturer", "m", helpers.Manufacturer, "Windows .msi Manufacturer/Publisher field.")
	createCmd.Flags().BoolVarP(&helpers.CreateGitRepo, "git", "G", false, "Create the --url repository (private) on its forge, then push the scaffolded code on the develop branch.")
	createCmd.Flags().StringVarP(&helpers.ConnectionName, "connection", "c", "", "Connection to use with -G (needed only when several connections point to the --url server).")

	_ = createCmd.MarkPersistentFlagRequired("url")
	// These must be supplied explicitly on `create` (they seed the manifest)
	_ = createCmd.MarkPersistentFlagRequired("desc")
	_ = createCmd.MarkPersistentFlagRequired("section")
	//_ = createCmd.MarkPersistentFlagRequired("depends")

	// refresh reuses the same value flags as create (minus -u and -G, which are create-only); unpassed flags keep their stored value
	// Both commands bind the same variables, and cobra writes a flag's default in its variable at registration: the defaults
	// below must stay the same as create's, otherwise create would end up with refresh's
	refreshCmd.PersistentFlags().StringVarP(&helpers.VersionNumber, "packagever", "V", "", "Package version number.")
	refreshCmd.PersistentFlags().StringVarP(&helpers.ReleaseNumber, "packagerel", "R", "", "Package release number.")
	refreshCmd.PersistentFlags().StringVarP(&helpers.Description, "desc", "D", "", "Package description.")
	refreshCmd.PersistentFlags().BoolVarP(&helpers.ArchLinuxStub, "archlinux", "A", false, "Refresh the Archlinux packaging stub.")
	refreshCmd.PersistentFlags().BoolVarP(&helpers.AlpineStub, "alpine", "a", false, "Refresh the Alpine packaging stub.")
	refreshCmd.PersistentFlags().BoolVarP(&helpers.DebianStub, "debian", "d", false, "Refresh the Debian packaging stub.")
	refreshCmd.PersistentFlags().BoolVarP(&helpers.RedHatStub, "redhat", "r", false, "Refresh the RedHat packaging stub.")
	refreshCmd.PersistentFlags().BoolVarP(&helpers.WindowsStub, "windows", "w", false, "Refresh the Windows packaging stub (Makefile only; the .wxs is created if missing, never overwritten).")
	refreshCmd.PersistentFlags().BoolVarP(&helpers.SkeletonStub, "skeleton", "k", false, "Refresh the skeleton stub (go.version only).")
	refreshCmd.PersistentFlags().StringVarP(&helpers.Maintainer, "maintainer", "M", helpers.Maintainer, "Software maintainer (\"Name <email>\").")
	refreshCmd.PersistentFlags().StringVarP(&helpers.Packager, "packager", "P", helpers.Packager, "Software packager (\"Name <email>\").")
	refreshCmd.PersistentFlags().StringVarP(&helpers.Section, "section", "s", "Packaging tool", "Debian package section / RPM group.")
	refreshCmd.PersistentFlags().StringVarP(&helpers.Dependencies, "depends", "e", "", "Package dependencies, comma-separated.")
	refreshCmd.PersistentFlags().StringVarP(&helpers.Manufacturer, "manufacturer", "m", helpers.Manufacturer, "Windows .msi Manufacturer/Publisher field.")
}
