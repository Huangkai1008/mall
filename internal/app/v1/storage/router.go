package storage

import (
	"github.com/gin-gonic/gin"

	"mall/internal/pkg/router"
)

func NewRouter(h *Handler) router.Group {
	return func(r *gin.RouterGroup) {
		storageApi := r.Group("/storage")
		{
			storageApi.POST("objects", h.PutObject)
		}
	}
}
