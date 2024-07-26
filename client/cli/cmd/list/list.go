package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "all"},
	Short:   "List a books",
	Long:    "List a books",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing book")
	},
}
