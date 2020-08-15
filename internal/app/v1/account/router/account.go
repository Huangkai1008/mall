package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"mall/internal/app/v1/account/handler"
	"mall/internal/pkg/router"
)

func NewAccountRouter(h *handler.AccountHandler) router.Group {
	return func(r *gin.RouterGroup) {
		userApi := r.Group("/accounts")
		{
			userApi.POST("", h.Register)
		}
	}
}

var ProviderSet = wire.NewSet(NewAccountRouter)
