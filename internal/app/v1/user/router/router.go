package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"mall/internal/app/v1/user/handler"
	"mall/internal/pkg/router"
)

func NewRouter(h *handler.Handler) router.Group {
	return func(r *gin.RouterGroup) {
		userApi := r.Group("/users")
		{
			userApi.POST("", h.Create)
		}
	}
}

var ProviderSet = wire.NewSet(NewRouter)
