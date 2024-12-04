package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "IP Tracker CLI Application",
		Long:  `IP Tracker CLI Application is a CLI application that allows you to track the location of an IP address.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
