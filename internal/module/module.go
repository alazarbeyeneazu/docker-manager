package module

import (
	"context"

	"github.com/alazarbeyeneazu/docker-manager/internal/constants"
)

type Docker interface {
	GetDockerStatus(ctx context.Context) ([]constants.GetContainersResponse, error)
}
