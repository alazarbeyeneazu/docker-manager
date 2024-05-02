package initiator

import (
	"github.com/alazarbeyeneazu/docker-manager/internal/glue/docker"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouting(
	grp *gin.RouterGroup,
	log zap.Logger,
	handler Handler,
) {
	// user.Init(grp, log, handler.User)
	docker.Init(grp, log, handler.docker)
}
