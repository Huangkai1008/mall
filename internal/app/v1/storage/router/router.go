package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"mall/internal/app/v1/storage/handler"

	"mall/internal/pkg/router"
)

func NewRouter(h *handler.Handler) router.Group {
	return func(r *gin.RouterGroup) {
		storageApi := r.Group("/storage")
		{
			storageApi.POST("objects", h.PutObject)
		}
	}
}

var ProviderSet = wire.NewSet(NewRouter)
