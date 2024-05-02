package docker

import (
	"net/http"

	"github.com/alazarbeyeneazu/docker-manager/internal/handler"
	"github.com/alazarbeyeneazu/docker-manager/internal/module"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

type docker struct {
	dockerModule module.Docker
	logger       zap.Logger
}

func Init(dockerModule module.Docker, log zap.Logger) handler.Docker {
	return &docker{
		dockerModule: dockerModule,
		logger:       log,
	}
}
func (d *docker) GetDockerStatus(c *gin.Context) {
	containers, err := d.dockerModule.GetDockerStatus(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	data := gin.H{
		"Containers": containers,
	}
	c.HTML(http.StatusOK, "index.html", data)

}
