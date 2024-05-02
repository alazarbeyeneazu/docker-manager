package docker

import (
	"net/http"

	"github.com/alazarbeyeneazu/docker-manager/internal/glue/routing"
	"github.com/alazarbeyeneazu/docker-manager/internal/handler"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(
	grp *gin.RouterGroup,
	log zap.Logger,
	user handler.Docker,

) {
	userRoute := []routing.Route{
		{
			Method:     http.MethodGet,
			Path:       "/list",
			Handler:    user.GetDockerStatus,
			Middleware: []gin.HandlerFunc{},
			Domains:    []string{"v1"},
		},
	}
	routing.RegisterRoute(grp, userRoute, log)

}
