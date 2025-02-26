package book

import (
	"fmt"

	"github.com/spf13/cobra"
)

var file string

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "add"},
	Short:   "Create a book",
	Long:    "Book store cli application to manage operations",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating book")
	},
}

func init() {
	createCmd.Flags().StringVarP(&file, "file", "f", "", "File to create a book")
}
