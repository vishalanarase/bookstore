package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vishalanarase/bookstore/client/cli/cmd/create"
	"github.com/vishalanarase/bookstore/client/cli/cmd/list"
)

// Bookstore root command
var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "Manage bookstore",
	Long:  "Create, update, delete, and list books",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Add create command
	bookCmd.AddCommand(create.CreateCmd)
	// Add list command
	bookCmd.AddCommand(list.ListCmd)
}
