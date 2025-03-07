/*
Copyright © 2024
*/
package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vishalanarase/bookstore/clients/cli/cmd/book"
	"github.com/vishalanarase/bookstore/clients/cli/cmd/login"
	"github.com/vishalanarase/bookstore/clients/cli/cmd/logout"
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

func init() {
	// Flags or options for root command
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.cli.yaml)")

	// Add subcommands
	rootCmd.AddCommand(book.BookCmd)
	// Add login commnad
	rootCmd.AddCommand(login.LoginCmd)
	// Add logout command
	rootCmd.AddCommand(logout.LogoutCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
