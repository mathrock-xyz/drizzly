package main

import (
	"github.com/mathrock-xyz/drizzly/info"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(info.Info)
}

var root = &cobra.Command{
	Use:   "dz",
	Short: "A collection of lightweight terminal utilities",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
