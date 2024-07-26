package list

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/vishalanarase/bookstore/openapiclient"
)

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "all"},
	Short:   "List a books",
	Long:    "List a books",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing book")
		// Create a new configuration object
		config := openapiclient.NewConfiguration()
		config.Servers[0].URL = "http://localhost:8080/v1"
		//config.BasePath = "https://api.bookstore.com/v1"

		// Set the API key in the headers
		//config.AddDefaultHeader("Authorization", "Bearer YOUR_API_KEY")

		// Create a new API client with the configuration
		client := openapiclient.NewAPIClient(config)

		listRequest := client.BooksAPI.ListBooks(context.Background())
		// Example: Call an endpoint
		books, resp, err := client.BooksAPI.ListBooksExecute(listRequest)
		if err != nil {
			log.Fatalf("Error calling API: %v", err)
		}

		fmt.Printf("Response code: %d\n", resp.StatusCode)
		for _, book := range books {
			fmt.Printf("Book: %s\n", *book.Id)
			fmt.Printf("Book: %s\n", *book.Title)
		}
	},
}
