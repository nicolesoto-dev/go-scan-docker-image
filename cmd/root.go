package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/nicolesoto-dev/go-scan-docker-image.git/internal/inspector"
	"github.com/spf13/cobra"
)

var (
	image    string
	platform string
)

var rootCmd = &cobra.Command{
	Use:   "dscan <image-name>",
	Short: "A basic Docker image metadata scanner.",
	Long:  "dscan is a tool that can help you to scan local and remote Docker images.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		imageName := args[0]
		if imageName == "" {
			return errors.New("image name cannot be empty")
		}

		dockerInspector, err := inspector.New()
		if err != nil {
			return fmt.Errorf("failed to create docker inspector: %w", err)
		}
		inspectedImage, err := dockerInspector.InspectImage(imageName)
		if err != nil {
			return fmt.Errorf("failed to inspect image '%s': %w", imageName, err)
		}
		fmt.Printf("Successfully inspected image: %s\n", imageName)
		fmt.Printf("  ID: %s\n", inspectedImage.ID)
		fmt.Printf("  Architecture: %s\n", inspectedImage.Architecture)
		fmt.Printf("  Operating System: %s\n", inspectedImage.Os)
		fmt.Printf("  Created: %s\n", inspectedImage.Created)

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
}
