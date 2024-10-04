package book

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/vishalanarase/bookstore/openapiclient"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "all"},
	Short:   "List a books",
	Long:    "List a books",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing book")

		// Create a new configuration object
		config := openapiclient.NewConfiguration()
		config.UserAgent = "cli"
		config.Servers[0].URL = "http://localhost:8080/v1"
		// Set the API key in the headers
		//config.AddDefaultHeader("Authorization", "Bearer YOUR_API_KEY")
		// Create a new API client with the configuration
		client := openapiclient.NewAPIClient(config)
		// Example: Call an endpoint
		listRequest := client.BooksAPI.ListBooks(context.Background())
		books, _, err := client.BooksAPI.ListBooksExecute(listRequest)
		if err != nil {
			log.Fatalf("Error calling API: %v", err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "Author", "Isbn", "Year", "Edition", "Rating"})
		// Set table borders
		table.SetBorder(true)  // Set the border around the table
		table.SetRowLine(true) // Set line between rows

		for _, book := range books {
			table.Append([]string{*book.Id, *book.Title, *book.Author, *book.Isbn,
				fmt.Sprintf("%d", *book.Edition), // Convert int to string
				fmt.Sprintf("%d", *book.Rating),  // Convert int to string
			})
		}

		table.Render() // Send output
	},
}
