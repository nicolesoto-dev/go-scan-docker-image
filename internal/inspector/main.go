package inspector

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Client interface {
	ImageInspectWithRaw(ctx context.Context, imageID string) (types.ImageInspect, []byte, error)
}

type DockerInspector struct {
	cli Client
}

func New() (*DockerInspector, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return &DockerInspector{cli: cli}, nil
}

func (i *DockerInspector) InspectImage(imageID string) (types.ImageInspect, error) {
	ctx := context.Background()
	result, _, err := i.cli.ImageInspectWithRaw(ctx, imageID)
	if err != nil {
		return types.ImageInspect{}, err
	}
	return result, nil
}

//(ctx context.Context, imageID string) (image.InspectResponse, []byte, error)
