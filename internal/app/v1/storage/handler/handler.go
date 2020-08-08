package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go.uber.org/zap"
	"mall/internal/app/v1/storage/schema"
	"mall/internal/app/v1/storage/service"

	resp "mall/internal/pkg/util/response"
)

type Handler struct {
	logger  *zap.Logger
	service *service.Service
}

func NewHandler(logger *zap.Logger, service *service.Service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

// PutObject put object to storage.
func (h *Handler) PutObject(c *gin.Context) {
	var ocSchema schema.ObjectCreateSchema
	if err := c.ShouldBind(&ocSchema); err != nil {
		errs := err.(validator.ValidationErrors)
		resp.BadEntityRequest(c, ocSchema.Validate(errs))
		return
	}

	file := ocSchema.File
	if objectSchema, err := h.service.PutObject(file.Filename, file); err != nil {
		resp.BadRequest(c, err.Error())
		return
	} else {
		resp.Created(c, objectSchema)
	}
}

var ProviderSet = wire.NewSet(NewHandler)
