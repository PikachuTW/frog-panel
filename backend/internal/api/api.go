package api

import (
	"frog-panel/internal/api/info"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	info.Register(r.Group("/info"))

	return r
}
