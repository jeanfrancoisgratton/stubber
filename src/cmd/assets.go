// stubber
// Written by J.F.Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/assets.go
// Original timestamp: 2026/05/03 10:35:50

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"stubber/assets"
)

var assetsCmd = &cobra.Command{
	Use:     "assets",
	Short:   "Embedded assets (templates) commands",
	Example: "  stubber assets list",
}

var assetsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Lists all the embedded assets",
	Example: "  stubber assets list",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if _, e := assets.List(true); e != nil {
			fmt.Println(e.Error())
		}
	},
}

func init() {
	assetsCmd.AddCommand(assetsListCmd)
}
