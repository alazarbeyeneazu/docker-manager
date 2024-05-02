package handler

import "github.com/gin-gonic/gin"

type Docker interface {
	GetDockerStatus(c *gin.Context)
}
