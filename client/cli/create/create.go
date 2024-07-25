package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

var file string

var CreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "add"},
	Short:   "Create a book",
	Long:    "Book store cli application to manage operations",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating book")
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&file, "file", "f", "", "File spec to create a book")
}
