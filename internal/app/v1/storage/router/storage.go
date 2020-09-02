package router

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

	"mall/internal/app/v1/storage/handler"
	"mall/internal/pkg/router"
)

func NewStorageRouter(h *handler.StorageHandler) router.Group {
	return func(e *echo.Group) {
		storageApi := e.Group("/storage")
		{
			storageApi.POST("objects", h.PutObject)
		}
	}
}

var ProviderSet = wire.NewSet(NewStorageRouter)
