package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"phl/cmd/env"
	"phl/cmd/template"
)

var rootCmd = &cobra.Command{
	Use:   "phl",
	Short: "PHL is Paul-Henri's development and management console",
}

func init() {
	template.InstallCommands(rootCmd)
	env.InstallCommands(rootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
