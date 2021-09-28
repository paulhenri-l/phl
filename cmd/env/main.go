package env

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "env",
	Short: "Dev env related commands",
}

func InstallCommands(cmd *cobra.Command) {
	cmd.AddCommand(rootCmd)
}
