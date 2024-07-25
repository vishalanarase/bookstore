package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vishalanarase/bookstore/client/cli/create"
)

var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "Manage bookstore",
	Long:  "Create, update, delete, and list books",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	bookCmd.AddCommand(create.CreateCmd)
}
