package handler

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"mall/internal/app/v1/storage/schema"
	"mall/internal/app/v1/storage/service"

	resp "mall/internal/pkg/util/response"
)

type StorageHandler struct {
	logger  *zap.Logger
	service *service.StorageService
}

func NewHandler(logger *zap.Logger, service *service.StorageService) *StorageHandler {
	return &StorageHandler{
		logger:  logger,
		service: service,
	}
}

// PutObject put object to storage.
func (h *StorageHandler) PutObject(c echo.Context) (err error){
	s := new(schema.ObjectCreateSchema)
	if err = c.Bind(s); err != nil {
		return err
	}
	if err = c.Validate(s); err != nil {
		return err
	}
	file := s.File
	if objectSchema, err := h.service.PutObject(file.Filename, file); err != nil {
		return err
	} else {
		return resp.Created(c, objectSchema)
	}
}

var ProviderSet = wire.NewSet(NewHandler)
