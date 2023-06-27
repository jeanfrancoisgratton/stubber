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
	Use:   "create",
	Short: "Creates the directory structure (skeleton) for the new software",
	Run: func(cmd *cobra.Command, args []string) {
		if err := executor.CreateStub(args); err != nil {
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(clCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.PersistentFlags().StringVarP(&helpers.StubRootDir, "stubdir", "s", ".", "Where to put the skeleton dir.")
	rootCmd.PersistentFlags().StringVarP(&helpers.GoVersion, "gover", "g", "1.20.5", "Where to put the skeleton dir.")
	rootCmd.PersistentFlags().BoolVarP(&helpers.AlpineStub, "alpine", "a", true, "Create an Alpine packaging stub.")
	rootCmd.PersistentFlags().BoolVarP(&helpers.DebianStub, "debian", "d", true, "Create a Debian packaging stub.")
	rootCmd.PersistentFlags().BoolVarP(&helpers.RedHatStub, "redhat", "r", true, "Create a RedHat packaging stub.")
}
