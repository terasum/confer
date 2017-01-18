package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "confer",
	Short: "Confer is a simple config file reader tool",
	Long: `A simple config file reader tool built with love by terasum and friends in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
	},
}