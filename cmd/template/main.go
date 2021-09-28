package template

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "template",
	Short: "Code template related commands",
	Aliases: []string{"tpl"},
}

func InstallCommands(cmd *cobra.Command) {
	cmd.AddCommand(rootCmd)
}
