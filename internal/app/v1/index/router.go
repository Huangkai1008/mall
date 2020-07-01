package index

import (
	"github.com/gin-gonic/gin"

	"mall/internal/pkg/router"
)

func NewRouter(h *Handler) router.Group {
	return func(r *gin.RouterGroup) {
		indexAPI := r.Group("/")
		{
			indexAPI.GET("ping", Ping)
		}
	}
}
