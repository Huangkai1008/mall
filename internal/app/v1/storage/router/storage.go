package router

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

	"github.com/Huangkai1008/micro-kit/pkg/transport/http"

	"github.com/Huangkai1008/mall/internal/app/v1/storage/handler"
)

func NewStorageRouter(h *handler.StorageHandler) http.Group {
	return func(e *echo.Group) {
		storageApi := e.Group("/storage")
		{
			storageApi.POST("objects", h.PutObject)
		}
	}
}

var ProviderSet = wire.NewSet(NewStorageRouter)
