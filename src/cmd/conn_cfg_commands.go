// stubber
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/conn_cfg_commands.go
// Original timestamp : 2026.10.18 11:04:26

package cmd

import (
	"fmt"
	"stubber/conn_cfg"

	"github.com/spf13/cobra"
)

var conn_cfgCmd = &cobra.Command{
	Use:     "conn",
	Aliases: []string{"connection"},
	Example: "  stubber conn list\n" +
		"  stubber conn add CONNECTION\n" +
		"  stubber conn info CONNECTION [CONNECTION...]\n" +
		"  stubber conn edit CONNECTION\n" +
		"  stubber conn rm CONNECTION",
	Short: "Connections registry commands",
	Long: `Manages the connections registry, $HOME/.config/JFG/stubber/connections.json

A connection tells stubber how to reach a forge (Forgejo, GitHub or Bitbucket): its type, its server URL,
and the credentials to use (an API token, or a username and password; the token wins when both are set).
Secrets are stored encrypted.

There is no default connection: "stubber create -G" uses the connection whose server hosts the -u URL.
When several connections point to the same server, choose one with "stubber create -G -c CONNECTION".`,
}

var conn_cfgListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Example: "  stubber conn list",
	Args:    cobra.NoArgs,
	Short:   "Lists all connections in the registry",
	Run: func(cmd *cobra.Command, args []string) {
		if err := conn_cfg.ListConnections(); err != nil {
			fmt.Println(err.Error())
		}
	},
}

var conn_cfgInfoCmd = &cobra.Command{
	Use:     "info CONNECTION [CONNECTION...]",
	Aliases: []string{"explain"},
	Example: "  stubber conn info myforgejo\n  stubber conn info myforgejo github bitbucket",
	Short:   "Prints the details of the given connections",
	Long:    `Prints the details of the given connections, as many as you wish; secrets are never shown`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := conn_cfg.ExplainConnections(args); err != nil {
			fmt.Println(err.Error())
		}
	},
}

var conn_cfgAddCmd = &cobra.Command{
	Use:     "add CONNECTION",
	Aliases: []string{"create"},
	Example: "  stubber conn add myforgejo",
	Short:   "Adds CONNECTION to the registry",
	Long: `Adds CONNECTION to the registry; you will be prompted for its values.

The server URL is the forge's root, as found in your repository URLs (ex: https://git.example.net:3000,
https://github.com, https://bitbucket.org); a GitHub connection needs an API token.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := conn_cfg.AddConnection(args[0]); err != nil {
			fmt.Println(err.Error())
		}
	},
}

var conn_cfgRmCmd = &cobra.Command{
	Use:     "rm CONNECTION",
	Aliases: []string{"remove"},
	Example: "  stubber conn rm myforgejo",
	Short:   "Removes CONNECTION from the registry",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := conn_cfg.RemoveConnection(args[0]); err != nil {
			fmt.Println(err.Error())
		}
	},
}

var conn_cfgEditCmd = &cobra.Command{
	Use:     "edit CONNECTION",
	Aliases: []string{"update"},
	Example: "  stubber conn edit myforgejo",
	Short:   "Edits CONNECTION in-place",
	Long:    `Pressing ENTER at a prompt keeps the current value`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := conn_cfg.EditConnection(args[0]); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	conn_cfgCmd.AddCommand(conn_cfgListCmd, conn_cfgInfoCmd, conn_cfgAddCmd, conn_cfgRmCmd, conn_cfgEditCmd)
}
