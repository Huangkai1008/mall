package http

import (
	"mall/internal/pkg/middleware"
	"mall/internal/pkg/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewRouter returns a new Gin router.
func NewRouter(o *Options, logger *zap.Logger, group router.Group) *gin.Engine {
	gin.SetMode(o.Mode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware(logger))
	apiGroup := r.Group("/api")
	v1Group := apiGroup.Group("/v1")
	{
		group(v1Group)
	}
	return r
}
