package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	resp "mall/internal/pkg/util/response"
)

type Handler struct {
	logger  *zap.Logger
	service *Service
}

func NewHandler(logger *zap.Logger, service *Service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

// PutObject put object to storage.
func (h *Handler) PutObject(c *gin.Context) {
	var ocSchema ObjectCreateSchema
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
