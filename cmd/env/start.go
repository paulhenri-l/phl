package env

import (
	"fmt"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
	Short: "Start the dev environment",
	Aliases: []string{"s"},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello there")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
