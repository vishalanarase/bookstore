/*
Copyright © 2024
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vishalanarase/bookstore/client/cli/cmd/book"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "cli to manage bookstore",
	Long:  "cli to manage bookstore",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// Flags or options for root command
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.cli.yaml)")

	// Add subcommands
	rootCmd.AddCommand(book.BookCmd)
}
