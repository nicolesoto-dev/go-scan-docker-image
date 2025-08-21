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
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return err
		}
		fmt.Printf("Serving on :%d\n", port)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error while executing dscan '%s'\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("i", "", "Image to scan")
}
