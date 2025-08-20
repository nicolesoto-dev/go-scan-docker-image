package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dscan",
	Long:  "dscan is a tool that can help you to scan local and remote Docker images'",
	Short: "Cool tool I guess",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error while executing dscan '%s'\n", err)
		os.Exit(1)
	}
}
