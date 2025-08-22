// Package inspector provides functionality to inspect local Docker images.
package inspector

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Client is an interface to abstract the Docker client for testing purposes.
// It only includes the methods we need from the official Docker client.
type Client interface {
	ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error)
}

// DockerInspector holds the Docker client.
type DockerInspector struct {
	cli Client
}

// New creates a new DockerInspector with a connection to the Docker daemon.
func New() (*DockerInspector, error) {
	// client.FromEnv configures the client from environment variables
	// like DOCKER_HOST, DOCKER_TLS_VERIFY, etc.
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return &DockerInspector{cli: cli}, nil
}

// Inspect takes an image name (e.g., "golang:1.21-alpine") and returns the
// detailed inspection data provided by the Docker daemon.
func (i *DockerInspector) Inspect(imageName string) (types.ImageInspect, error) {
	// The context is used to manage cancellation of the API request.
	ctx := context.Background()

	// cli.ImageInspectWithRaw is the SDK function that calls the "docker inspect" API endpoint.
	// It returns the parsed struct, the raw JSON body, and an error.
	// We only need the struct for now.
	inspectionData, _, err := i.cli.ImageInspectWithRaw(ctx, imageName)
	if err != nil {
		return types.ImageInspect{}, err
	}

	return inspectionData, nil
}
