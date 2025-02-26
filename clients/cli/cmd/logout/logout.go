package logout

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Bookstore logout command
var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout to bookstore",
	Long:  "Logout to bookstore",

	Run: func(cmd *cobra.Command, args []string) {
		// Logout from current session

		fmt.Println("Logout successful")
	},
}
