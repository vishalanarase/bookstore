package book

import (
	"github.com/spf13/cobra"
)

// Bookstore root command
var BookCmd = &cobra.Command{
	Use:   "book",
	Short: "Manage bookstore",
	Long:  "Create, update, delete, and list books",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Add create command
	BookCmd.AddCommand(createCmd)
	// Add list command
	BookCmd.AddCommand(listCmd)
}
