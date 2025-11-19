package info

import (
	"github.com/spf13/cobra"
)

func init() {
	Info.AddCommand(write)
	Info.AddCommand(cat)
	Info.AddCommand(rm)
}

var Info = &cobra.Command{
	Use:   "info",
	Short: "Attach and manage descriptive notes for files and directories.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
