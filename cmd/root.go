package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "test",
	Short: "test",
	Long: "test",
}

func Execute() error{
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd,timeCmd,sqlCmd)
}