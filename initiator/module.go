package initiator

import (
	"github.com/alazarbeyeneazu/docker-manager/internal/module"
	"github.com/alazarbeyeneazu/docker-manager/internal/module/docker"
	"go.uber.org/zap"
)

type Module struct {
	docker module.Docker
}

func InitModule(log *zap.Logger, persistenceDb Persistence) Module {
	return Module{
		docker: docker.Init(log),
	}
}
