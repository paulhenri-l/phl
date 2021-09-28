package template

import (
	"fmt"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use: "install",
	Short: "install the given code template in the current directory",
	Aliases: []string{"i"},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello there")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
