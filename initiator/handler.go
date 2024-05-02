package initiator

import (
	"github.com/alazarbeyeneazu/docker-manager/internal/handler"
	"github.com/alazarbeyeneazu/docker-manager/internal/handler/docker"
	"go.uber.org/zap"
)

type Handler struct {
	docker handler.Docker
}

func InitHandler(module Module, log zap.Logger) Handler {
	return Handler{
		docker: docker.Init(module.docker, log),
	}
}
