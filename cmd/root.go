package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	image    string
	platform string
)

var rootCmd = &cobra.Command{
	Use:   "dscan",
	Long:  "dscan is a tool that can help you to scan local and remote Docker images'",
	Short: "Cool tool I guess",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		image, err := cmd.Flags().GetString("image")
		if err != nil {
			return err
		}
		fmt.Printf("Image :%s\n", image)
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
	rootCmd.PersistentFlags().StringVar(&image, "image", "", "Image to scan")
}
