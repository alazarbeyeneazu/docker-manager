package docker

import (
	"context"
	"strings"

	"github.com/alazarbeyeneazu/docker-manager/internal/constants"
	"github.com/alazarbeyeneazu/docker-manager/internal/module"
	dockercli "github.com/fsouza/go-dockerclient"
	"go.uber.org/zap"
)

type dockerhandler struct {
}

func Init(log *zap.Logger) module.Docker {
	return &dockerhandler{}
}
func (d *dockerhandler) GetDockerStatus(ctx context.Context) ([]constants.GetContainersResponse, error) {
	names := []constants.GetContainersResponse{}
	client, err := dockercli.NewClientFromEnv()
	if err != nil {
		return []constants.GetContainersResponse{}, err
	}
	imags, err := client.ListContainers(dockercli.ListContainersOptions{
		All: true,
	})
	if err != nil {
		return []constants.GetContainersResponse{}, err
	}
	for _, cont := range imags {
		p := cont.Names[0]
		names = append(names, constants.GetContainersResponse{
			ID:     cont.ID,
			Name:   strings.Trim(p, "/"),
			Status: cont.Status,
		})
	}
	return names, nil

}
