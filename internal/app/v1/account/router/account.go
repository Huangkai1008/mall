package router

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

	"mall/internal/app/v1/account/handler"
	"mall/internal/pkg/router"
)

func NewAccountRouter(h *handler.AccountHandler) router.Group {
	return func(e *echo.Group) {
		userApi := e.Group("/accounts")
		{
			userApi.POST("/", h.Register)
			userApi.POST("/authentication", h.Login)
		}
	}
}

var ProviderSet = wire.NewSet(NewAccountRouter)
