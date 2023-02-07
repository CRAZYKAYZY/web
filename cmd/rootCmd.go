package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mcli",
	Short: "mcli-tool",
	Long:  `entry point of the cobra command tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the cobra cli tool")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
