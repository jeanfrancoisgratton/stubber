// stubber
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original timestamp: 2025/09/15 08:35
// Original filename: src/cmd/completion.go
// Bash/Zsh completion scripts via Cobra.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates shell completion scripts",
	Long: `Generates completion scripts for your shell.

Bash:
  $ source <(stubber completion bash)
  # To persist:
  $ stubber completion bash | sudo tee /etc/bash_completion.d/stubber > /dev/null

Zsh:
  $ stubber completion zsh > ~/.zsh/stubber
  $ echo 'fpath=($HOME/.zsh $fpath)' >> ~/.zshrc
  $ echo 'autoload -Uz compinit && compinit' >> ~/.zshrc
  # Or, for current session:
  $ source <(stubber completion zsh)
`,
	Example: "  source <(stubber completion bash)\n  source <(stubber completion zsh)",
}

var completionBashCmd = &cobra.Command{
	Use:     "bash",
	Short:   "Generates a Bash completion script",
	Example: "  source <(stubber completion bash)",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// V2 is recommended; writes to stdout
		return rootCmd.GenBashCompletionV2(os.Stdout, true)
	},
}

var completionZshCmd = &cobra.Command{
	Use:     "zsh",
	Short:   "Generates a Zsh completion script",
	Example: "  stubber completion zsh > ~/.zsh/stubber",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Ensure the script is zsh-compatible
		return rootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	completionCmd.AddCommand(completionBashCmd, completionZshCmd)
}
