package router

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

	"github.com/Huangkai1008/micro-kit/pkg/transport/http"

	"github.com/Huangkai1008/mall/internal/app/v1/account/handler"
)

func NewAccountRouter(h *handler.AccountHandler) http.Group {
	return func(e *echo.Group) {
		userApi := e.Group("/accounts")
		{
			userApi.POST("/", h.Register)
			userApi.POST("/authentication", h.Login)
		}
	}
}

var ProviderSet = wire.NewSet(NewAccountRouter)
