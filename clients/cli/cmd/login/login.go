package login

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// Bookstore login command
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to bookstore",
	Long:  "Login to bookstore",

	Run: func(cmd *cobra.Command, args []string) {
		// Get username and password from flags also handle the error
		// If login is successful, print "Login successful"
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatal(err)
		}
		if username == "" {
			log.Fatal("Username is required")
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatal(err)
		}
		if password == "" {
			log.Fatal("Password is required")
		}

		// Call the login API

		fmt.Println("Login successful", username, password)
	},
}

func init() {
	// Add --username parameter
	LoginCmd.Flags().StringP("username", "u", "", "Username")
	// Add --password parameter
	LoginCmd.Flags().StringP("password", "p", "", "Password")
}
